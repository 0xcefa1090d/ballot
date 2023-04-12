// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12 <=0.8.18;

interface IVoting {
    function getVote(
        uint256 voteId
    )
        external
        view
        returns (
            bool open,
            bool executed,
            uint64 startDate,
            uint64 snapshotBlock,
            uint64 supportRequired,
            uint64 minAcceptQuorum,
            uint256 yea,
            uint256 nay,
            uint256 votingPower,
            bytes memory script
        );
}
