// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {ERC20} from "transmissions11/solmate/tokens/ERC20.sol";
import {SafeTransferLib} from "transmissions11/solmate/utils/SafeTransferLib.sol";

contract NewVoteBounty {
    using SafeTransferLib for ERC20;

    uint256 public constant OPEN_BOUNTY_COST = 1e16;

    struct Bounty {
        uint256 amount;
        uint256 timestamp;
    }

    mapping(bytes32 => Bounty) public getBounty;
    mapping(address => uint256) public getRefundAmount;

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
}
