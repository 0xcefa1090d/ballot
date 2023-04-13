// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {ERC20} from "transmissions11/solmate/tokens/ERC20.sol";
import {IVoting} from "./interfaces/IVoting.sol";
import {ReceiptProofVerifier} from "./libs/ReceiptProofVerifier.sol";
import {RLPReader} from "hamdiallam/Solidity-RLP/RLPReader.sol";
import {SafeTransferLib} from "transmissions11/solmate/utils/SafeTransferLib.sol";

/**
 * @title New Vote Bounty System
 * @author 0xcefa1090d
 * @notice Trustless bounty system for vote creation.
 */
contract NewVoteBounty {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;
    using SafeTransferLib for ERC20;

    uint256 public constant OPEN_BOUNTY_COST = 2e16;
    bytes32 constant START_VOTE_SIG =
        0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394;

    address public immutable VOTING;

    mapping(uint256 => bytes32) public cache;
    mapping(bytes32 => uint256) public getCommitCloseBountyBlockNumber;
    mapping(address => uint256) public getRefundAmount;
    mapping(bytes32 => uint256) public getRewardAmount;

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

    event IncreaseBounty(bytes32 indexed identifier, uint256 addedAmount);

    event CommitCloseBounty(bytes32 indexed identifier);

    event ApplyCloseBounty(bytes32 indexed identifier);

    constructor(address voting) {
        VOTING = voting;
    }

    /**
     * @notice Claim an existing bounty.
     * @param creator The account which created the bounty being claimed.
     * @param createdAt The timestamp of the block which the bounty was created in.
     * @param rewardToken The reward token for the bounty.
     * @param digest The keccak256 hash of the concatenated metadata keccak256 hash and script keccak256 hash.
     * @param receiptIndex The index of the receipt which has the appropriate StartVote log.
     * @param headerRlpBytes The rlp encoded header of the block which the StartVote log was emitted.
     * @param proofRlpBytes The rlp encoded MPT proof of the receipt from the receipts trie of the block the StartVote log was emitted in.
     */
    function claimBounty(
        address creator,
        uint256 createdAt,
        address rewardToken,
        bytes32 digest,
        uint256 receiptIndex,
        bytes memory headerRlpBytes,
        bytes memory proofRlpBytes
    ) external {
        bytes32 identifier = keccak256(
            abi.encode(creator, createdAt, rewardToken, digest)
        );
        uint256 rewardAmount = getRewardAmount[identifier];
        require(rewardAmount != 0);

        ReceiptProofVerifier.BlockHeader memory header = ReceiptProofVerifier
            .parseBlockHeader(headerRlpBytes);
        require(
            header.hash == blockhash(header.number) ||
                header.hash == cache[header.number]
        );
        require(createdAt < header.timestamp);

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

            getRewardAmount[identifier] = 0;
            getRefundAmount[creator] += OPEN_BOUNTY_COST;

            address claimant = address(
                uint160(uint256(receipt.logs[i].topics[2]))
            );
            emit ClaimBounty(identifier, claimant, voteId);
            ERC20(rewardToken).safeTransfer(claimant, rewardAmount);
            return;
        }
        revert();
    }

    /**
     * @notice Open a new bounty.
     * @dev OPEN_BOUNTY_COST value of ETH must be attached for successfuly creation of a bounty.
     * @param rewardToken The reward token for the bounty.
     * @param rewardAmount The reward amount for the bounty.
     * @param metadata The metadata of the desired vote to be started.
     * @param script The script of the desired vote to be started.
     */
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
                block.timestamp,
                rewardToken,
                keccak256(
                    bytes.concat(keccak256(bytes(metadata)), keccak256(script))
                )
            )
        );
        require(getRewardAmount[identifier] == 0);

        getRewardAmount[identifier] = rewardAmount;

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

    /**
     * @notice Increase the bounty reward.
     * @param createdAt The timestamp of the block which the bounty was created.
     * @param rewardToken The reward token of the bounty.
     * @param digest The keccak256 hash of the concatenated metadata keccak256 hash and script keccak256 hash.
     * @param increaseAmount The amount to increase the bounty reward by.
     */
    function increaseBounty(
        uint256 createdAt,
        address rewardToken,
        bytes32 digest,
        uint256 increaseAmount
    ) external {
        require(increaseAmount != 0);

        bytes32 identifier = keccak256(
            abi.encode(msg.sender, createdAt, rewardToken, digest)
        );
        uint256 rewardAmount = getRewardAmount[identifier];
        require(rewardAmount != 0);

        getRewardAmount[identifier] = rewardAmount + increaseAmount;

        emit IncreaseBounty(identifier, increaseAmount);

        ERC20(rewardToken).safeTransferFrom(
            msg.sender,
            address(this),
            increaseAmount
        );
    }

    /**
     * @notice Refund a bounty creator their open bounty costs.
     * @param creator The bounty creator account.
     */
    function refund(address creator) external {
        uint256 refundAmount = getRefundAmount[creator];
        require(refundAmount != 0);

        getRefundAmount[creator] = 0;

        SafeTransferLib.safeTransferETH(creator, refundAmount);
    }

    /**
     * @notice Commit to closing an active bounty.
     * @param createdAt The timestamp of the block which the bounty was created.
     * @param rewardToken The reward token of the bounty.
     * @param digest The keccak256 hash of the concatenated metadata keccak256 hash and script keccak256 hash.
     */
    function commitCloseBounty(
        uint256 createdAt,
        address rewardToken,
        bytes32 digest
    ) external {
        bytes32 identifier = keccak256(
            abi.encode(msg.sender, createdAt, rewardToken, digest)
        );
        require(getRewardAmount[identifier] != 0);

        getCommitCloseBountyBlockNumber[identifier] = block.number;
        emit CommitCloseBounty(identifier);
    }

    /**
     * @notice Apply a previously committed close an active bounty request.
     * @dev The commitment to close a bounty must have occurred between 256 and 128 block ago.
     * @param createdAt The timestamp of the block which the bounty was created.
     * @param rewardToken The reward token of the bounty.
     * @param digest The keccak256 hash of the concatenated metadata keccak256 hash and script keccak256 hash.
     */
    function applyCloseBounty(
        uint256 createdAt,
        address rewardToken,
        bytes32 digest
    ) external {
        bytes32 identifier = keccak256(
            abi.encode(msg.sender, createdAt, rewardToken, digest)
        );
        uint256 rewardAmount = getRewardAmount[identifier];
        require(rewardAmount != 0);

        uint256 commitBlockNumber = getCommitCloseBountyBlockNumber[identifier];
        require(
            block.number - 256 < commitBlockNumber &&
                commitBlockNumber < block.number - 128
        );

        getRewardAmount[identifier] = 0;
        emit ApplyCloseBounty(identifier);

        ERC20(rewardToken).safeTransfer(msg.sender, rewardAmount);
        SafeTransferLib.safeTransferETH(msg.sender, OPEN_BOUNTY_COST);
    }

    /**
     * @notice Cache a block hash for later use when claiming a bounty.
     * @dev Useful for when gas costs are prohibitively high.
     * @param blockNumber The block number to cache.
     */
    function cacheBlockHash(uint256 blockNumber) external {
        bytes32 value = blockhash(blockNumber);
        require(value != bytes32(0));

        cache[blockNumber] = value;
    }

    /**
     * @notice Utility function for calculating the identifier of a bounty.
     * @param creator The account which created the bounty being claimed.
     * @param createdAt The timestamp of the block which the bounty was created in.
     * @param rewardToken The reward token for the bounty.
     * @param metadata The metadata of the desired vote to be started.
     * @param script The script of the desired vote to be started.
     */
    function calculateIdentifier(
        address creator,
        uint256 createdAt,
        address rewardToken,
        string memory metadata,
        bytes memory script
    ) external pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    creator,
                    createdAt,
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
