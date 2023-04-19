import ape
import pytest
from eth_hash.auto import keccak

AMOUNT = 10**21
METADATA = "Hello, World!"
SCRIPT = b"\xde\xad\xbe\xef"
DIGEST = keccak(keccak(METADATA.encode()) + keccak(SCRIPT))
CREATION_TIME = None


@pytest.fixture(scope="module", autouse=True)
def setup(alice, start_vote_bounty, token_mock):
    global CREATION_TIME

    token_mock.approve(start_vote_bounty, 2**256 - 1, sender=alice)
    receipt = start_vote_bounty.openBounty(
        token_mock,
        AMOUNT,
        METADATA,
        SCRIPT,
        sender=alice,
        value=start_vote_bounty.OPEN_BOUNTY_COST(),
    )
    CREATION_TIME = receipt.timestamp


def test_increase_bounty_success(alice, start_vote_bounty, token_mock):
    receipt = start_vote_bounty.increaseBounty(
        CREATION_TIME, token_mock, DIGEST, AMOUNT, sender=alice
    )
    identifier = start_vote_bounty.calculateIdentifier(
        alice, CREATION_TIME, token_mock, METADATA, SCRIPT
    )

    # storage
    assert start_vote_bounty.getRewardAmount(identifier) == 2 * AMOUNT

    # event
    increase_bounty_event = next(iter(start_vote_bounty.IncreaseBounty.from_receipt(receipt)))

    assert increase_bounty_event.identifier == identifier
    assert increase_bounty_event.addedAmount == AMOUNT

    # interaction
    assert token_mock.balanceOf(start_vote_bounty) == 2 * AMOUNT


def test_increase_bounty_fails_invalid_increase_amount(alice, start_vote_bounty, token_mock):
    with ape.reverts():
        start_vote_bounty.increaseBounty(CREATION_TIME, token_mock, DIGEST, 0, sender=alice)


def test_increase_bounty_fails_bounty_does_not_exist(bob, start_vote_bounty, token_mock):
    with ape.reverts():
        start_vote_bounty.increaseBounty(CREATION_TIME, token_mock, DIGEST, AMOUNT, sender=bob)
