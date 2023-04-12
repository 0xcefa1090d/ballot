import pytest
from eth_hash.auto import keccak

AMOUNT = 10**21
METADATA = "Hello, World!"
SCRIPT = b"\xde\xad\xbe\xef"
DIGEST = keccak(keccak(METADATA.encode()) + keccak(SCRIPT))


@pytest.fixture(scope="module", autouse=True)
def setup(alice, new_vote_bounty, token_mock):
    token_mock.approve(new_vote_bounty, 2**256 - 1, sender=alice)
    new_vote_bounty.openBounty(
        token_mock, AMOUNT, METADATA, SCRIPT, sender=alice, value=new_vote_bounty.OPEN_BOUNTY_COST()
    )


def test_increase_bounty_success(alice, new_vote_bounty, token_mock):
    receipt = new_vote_bounty.increaseBounty(alice, token_mock, DIGEST, AMOUNT, sender=alice)
    identifier = new_vote_bounty.calculateIdentifier(alice, token_mock, METADATA, SCRIPT)

    # storage
    bounty = new_vote_bounty.getBounty(identifier)

    assert bounty.amount == 2 * AMOUNT
    assert bounty.timestamp < receipt.timestamp

    # event
    increase_bounty_event = next(iter(new_vote_bounty.IncreaseBounty.from_receipt(receipt)))

    assert increase_bounty_event.identifier == identifier
    assert increase_bounty_event.depositor == alice
    assert increase_bounty_event.addedAmount == AMOUNT

    # interaction
    assert token_mock.balanceOf(new_vote_bounty) == 2 * AMOUNT
