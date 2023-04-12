// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {ERC20} from "transmissions11/solmate/tokens/ERC20.sol";
import {IVoting} from "./interfaces/IVoting.sol";
import {ReceiptProofVerifier} from "./libs/ReceiptProofVerifier.sol";
import {RLPReader} from "hamdiallam/Solidity-RLP/RLPReader.sol";
import {SafeTransferLib} from "transmissions11/solmate/utils/SafeTransferLib.sol";

contract NewVoteBounty {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;
    using SafeTransferLib for ERC20;

    uint256 public constant OPEN_BOUNTY_COST = 1e16;
    bytes32 constant START_VOTE_SIG =
        0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394;

    address public immutable VOTING;

    struct Bounty {
        uint256 amount;
        uint256 timestamp;
    }

    mapping(bytes32 => Bounty) public getBounty;
    mapping(address => uint256) public getRefundAmount;

    event ClaimBounty(
        bytes32 indexed identifier,
        address indexed claimant,
        uint256 indexed voteId
    );

    event OpenBounty(
        bytes32 indexed identifier,
        address indexed creator,
        address indexed rewardToken,
        uint256 rewardAmount,
        string metadata,
        bytes script
    );

    event IncreaseBounty(
        bytes32 indexed identifier,
        address indexed depositor,
        uint256 newRewardAmount
    );

    constructor(address voting) {
        VOTING = voting;
    }

    function claimBounty(
        address creator,
        address rewardToken,
        bytes32 digest,
        uint256 receiptIndex,
        bytes memory headerRlpBytes,
        bytes memory proofRlpBytes
    ) external {
        bytes32 identifier = keccak256(
            abi.encode(creator, rewardToken, digest)
        );
        Bounty memory bounty = getBounty[identifier];
        require(bounty.amount != 0);

        ReceiptProofVerifier.BlockHeader memory header = ReceiptProofVerifier
            .verifyBlockHeader(headerRlpBytes);
        require(bounty.timestamp < header.timestamp);

        ReceiptProofVerifier.Receipt memory receipt = ReceiptProofVerifier
            .extractReceiptFromProof(
                receiptIndex,
                header.receiptsRootHash,
                proofRlpBytes.toRlpItem().toList()
            );
        require(receipt.status);

        for (uint256 i = 0; i < receipt.logs.length; i++) {
            if (receipt.logs[i].logger != VOTING) continue;
            if (receipt.logs[i].topics[0] != START_VOTE_SIG) continue;

            (string memory metadata, , , , ) = abi.decode(
                receipt.logs[i].data,
                (string, uint256, uint256, uint256, uint256)
            );

            uint256 voteId = uint256(receipt.logs[i].topics[1]);
            (, , , , , , , , , bytes memory script) = IVoting(VOTING).getVote(
                voteId
            );

            if (
                keccak256(
                    bytes.concat(keccak256(bytes(metadata)), keccak256(script))
                ) != digest
            ) continue;

            delete getBounty[identifier];
            getRefundAmount[creator] += OPEN_BOUNTY_COST;

            address claimant = address(
                uint160(uint256(receipt.logs[i].topics[2]))
            );
            ERC20(rewardToken).safeTransfer(claimant, bounty.amount);
            emit ClaimBounty(identifier, claimant, voteId);
            return;
        }
        revert();
    }

    function openBounty(
        address rewardToken,
        uint256 rewardAmount,
        string memory metadata,
        bytes memory script
    ) external payable returns (bytes32) {
        require(msg.value == OPEN_BOUNTY_COST && rewardAmount != 0);

        bytes32 identifier = keccak256(
            abi.encode(
                msg.sender,
                rewardToken,
                keccak256(
                    bytes.concat(keccak256(bytes(metadata)), keccak256(script))
                )
            )
        );
        require(getBounty[identifier].amount == 0);

        getBounty[identifier] = Bounty(rewardAmount, block.timestamp);

        emit OpenBounty(
            identifier,
            msg.sender,
            rewardToken,
            rewardAmount,
            metadata,
            script
        );

        ERC20(rewardToken).safeTransferFrom(
            msg.sender,
            address(this),
            rewardAmount
        );
        return identifier;
    }

    function increaseBounty(
        address creator,
        address rewardToken,
        bytes32 digest,
        uint256 additionalAmount
    ) external {
        require(additionalAmount != 0);

        bytes32 identifier = keccak256(
            abi.encode(creator, rewardToken, digest)
        );
        uint256 rewardAmount = getBounty[identifier].amount;
        require(rewardAmount != 0);

        rewardAmount += additionalAmount;
        getBounty[identifier].amount = rewardAmount;

        emit IncreaseBounty(identifier, msg.sender, rewardAmount);

        ERC20(rewardToken).safeTransferFrom(
            msg.sender,
            address(this),
            additionalAmount
        );
    }

    function refund(address creator) external {
        uint256 refundAmount = getRefundAmount[creator];
        require(refundAmount != 0);

        getRefundAmount[creator] = 0;

        SafeTransferLib.safeTransferETH(creator, refundAmount);
    }

    function calculateIdentifier(
        address creator,
        address rewardToken,
        string memory metadata,
        bytes memory script
    ) external pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    creator,
                    rewardToken,
                    keccak256(
                        bytes.concat(
                            keccak256(bytes(metadata)),
                            keccak256(script)
                        )
                    )
                )
            );
    }
}
