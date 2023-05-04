import ape
import pytest

AMOUNT = 10**21
METADATA = "Hello, World!"
SCRIPT = b"\xde\xad\xbe\xef"


@pytest.fixture(scope="module", autouse=True)
def setup(alice, start_vote_bounty, token_mock):
    token_mock.approve(start_vote_bounty, 2**256 - 1, sender=alice)


def test_open_bounty_success(alice, start_vote_bounty, token_mock):
    receipt = start_vote_bounty.openBounty(
        token_mock,
        AMOUNT,
        METADATA,
        SCRIPT,
        sender=alice,
        value=start_vote_bounty.OPEN_BOUNTY_COST(),
    )
    identifier = start_vote_bounty.calculateIdentifier(
        alice, receipt.timestamp, token_mock, METADATA, SCRIPT
    )

    # storage
    assert start_vote_bounty.getRewardAmount(identifier) == AMOUNT

    # event
    open_bounty_event = next(iter(start_vote_bounty.OpenBounty.from_receipt(receipt)))

    assert open_bounty_event.identifier == identifier
    assert open_bounty_event.creator == alice
    assert open_bounty_event.rewardToken == token_mock
    assert open_bounty_event.rewardAmount == AMOUNT
    assert open_bounty_event.startTime == receipt.timestamp
    assert open_bounty_event.metadata == METADATA
    assert open_bounty_event.script == SCRIPT

    # interaction
    assert token_mock.balanceOf(start_vote_bounty) == AMOUNT

    # retval
    assert receipt.return_value == identifier


def test_open_bounty_scheduled_success(alice, chain, start_vote_bounty, token_mock):
    start_time = chain.pending_timestamp + 86400
    receipt = start_vote_bounty.openBounty(
        start_time,
        token_mock,
        AMOUNT,
        METADATA,
        SCRIPT,
        sender=alice,
        value=start_vote_bounty.OPEN_BOUNTY_COST(),
    )
    identifier = start_vote_bounty.calculateIdentifier(
        alice, start_time, token_mock, METADATA, SCRIPT
    )

    assert start_vote_bounty.getRewardAmount(identifier) == AMOUNT

    # event
    open_bounty_event = next(iter(start_vote_bounty.OpenBounty.from_receipt(receipt)))

    assert open_bounty_event.startTime == start_time


@pytest.mark.parametrize("value_dx", [-1, 1])
def test_open_bounty_fails_invalid_txn_value(alice, start_vote_bounty, token_mock, value_dx):
    with ape.reverts():
        start_vote_bounty.openBounty(
            token_mock,
            AMOUNT,
            METADATA,
            SCRIPT,
            sender=alice,
            value=start_vote_bounty.OPEN_BOUNTY_COST() + value_dx,
        )


def test_open_bounty_fails_invalid_reward_amount(alice, start_vote_bounty, token_mock):
    with ape.reverts():
        start_vote_bounty.openBounty(
            token_mock,
            0,
            METADATA,
            SCRIPT,
            sender=alice,
            value=start_vote_bounty.OPEN_BOUNTY_COST(),
        )


def test_open_bounty_fails_invalid_start_time(alice, start_vote_bounty, token_mock):
    with ape.reverts():
        start_vote_bounty.openBounty(
            30,
            token_mock,
            0,
            METADATA,
            SCRIPT,
            sender=alice,
            value=start_vote_bounty.OPEN_BOUNTY_COST(),
        )
