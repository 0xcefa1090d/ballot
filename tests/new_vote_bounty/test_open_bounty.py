import ape
import pytest

AMOUNT = 10**21
METADATA = "Hello, World!"
SCRIPT = b"\xde\xad\xbe\xef"


@pytest.fixture(scope="module", autouse=True)
def setup(alice, new_vote_bounty, token_mock):
    token_mock.approve(new_vote_bounty, 2**256 - 1, sender=alice)


def test_open_bounty_success(alice, new_vote_bounty, token_mock):
    receipt = new_vote_bounty.openBounty(
        token_mock,
        AMOUNT,
        METADATA,
        SCRIPT,
        sender=alice,
        value=new_vote_bounty.OPEN_BOUNTY_COST(),
    )
    identifier = new_vote_bounty.calculateIdentifier(alice, token_mock, METADATA, SCRIPT)

    # storage
    bounty = new_vote_bounty.getBounty(identifier)

    assert bounty.amount == AMOUNT
    assert bounty.timestamp == receipt.timestamp

    # event
    open_bounty_event = next(iter(new_vote_bounty.OpenBounty.from_receipt(receipt)))

    assert open_bounty_event.identifier == identifier
    assert open_bounty_event.creator == alice
    assert open_bounty_event.rewardToken == token_mock
    assert open_bounty_event.rewardAmount == AMOUNT
    assert open_bounty_event.metadata == METADATA
    assert open_bounty_event.script == SCRIPT

    # interaction
    assert token_mock.balanceOf(new_vote_bounty) == AMOUNT

    # retval
    assert receipt.return_value == identifier


@pytest.mark.parametrize("value_dx", [-1, 1])
def test_open_bounty_fails_invalid_txn_value(alice, new_vote_bounty, token_mock, value_dx):
    with ape.reverts():
        new_vote_bounty.openBounty(
            token_mock,
            AMOUNT,
            METADATA,
            SCRIPT,
            sender=alice,
            value=new_vote_bounty.OPEN_BOUNTY_COST() + value_dx,
        )
