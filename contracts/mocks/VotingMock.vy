# @version 0.3.7
"""
@title Voting Mock
@license MIT
"""

event StartVote:
    vote_id: indexed(uint256)
    creator: indexed(address)
    metadata: String[1024]
    min_balance: uint256
    min_time: uint256
    total_supply: uint256
    creator_voting_power: uint256


counter: uint256
scripts: HashMap[uint256, Bytes[1024]]


@external
def new_vote(_metadata: String[1024], _script: Bytes[1024]):
    vote_id: uint256 = self.counter

    self.counter = vote_id + 1
    self.scripts[vote_id] = _script

    log StartVote(vote_id, msg.sender, _metadata, 0, 0, 0, 0)


@view
@external
def getVote(_vote_id: uint256) -> (bool, bool, uint64, uint64, uint64, uint64, uint256, uint256, uint256, Bytes[1024]):
    return (False, False, 0, 0, 0, 0, 0, 0, 0, self.scripts[_vote_id])
