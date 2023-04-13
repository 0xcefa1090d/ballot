import ape
import pytest
from eth_hash.auto import keccak

AMOUNT = 10**21
METADATA = "Hello, World!"
SCRIPT = b"\xde\xad\xbe\xef"
DIGEST = keccak(keccak(METADATA.encode()) + keccak(SCRIPT))
CREATION_TIME = None


@pytest.fixture(scope="module", autouse=True)
def setup(alice, new_vote_bounty, token_mock):
    global CREATION_TIME

    token_mock.approve(new_vote_bounty, 2**256 - 1, sender=alice)
    receipt = new_vote_bounty.openBounty(
        token_mock, AMOUNT, METADATA, SCRIPT, sender=alice, value=new_vote_bounty.OPEN_BOUNTY_COST()
    )
    CREATION_TIME = receipt.timestamp


def test_commit_close_bounty_success(alice, new_vote_bounty, token_mock):
    receipt = new_vote_bounty.commitCloseBounty(CREATION_TIME, token_mock, DIGEST, sender=alice)
    identifier = new_vote_bounty.calculateIdentifier(
        alice, CREATION_TIME, token_mock, METADATA, SCRIPT
    )

    # storage
    assert new_vote_bounty.getCommitCloseBountyBlockNumber(identifier) == receipt.block_number

    # event
    commit_close_bounty_event = next(iter(new_vote_bounty.CommitCloseBounty.from_receipt(receipt)))
    assert commit_close_bounty_event.identifier == identifier


def test_commit_close_bounty_fails_invalid_bounty(bob, new_vote_bounty, token_mock):
    with ape.reverts():
        new_vote_bounty.commitCloseBounty(CREATION_TIME, token_mock, DIGEST, sender=bob)
