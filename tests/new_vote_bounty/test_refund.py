import ape
import pytest
from eth_hash.auto import keccak

AMOUNT = 10**21
METADATA = "Hello, World!"
SCRIPT = b"\xde\xad\xbe\xef"
DIGEST = keccak(keccak(METADATA.encode()) + keccak(SCRIPT))


@pytest.fixture(scope="module", autouse=True)
def setup(
    alice,
    bob,
    get_block_header_rlp,
    get_receipt_proof_rlp,
    new_vote_bounty,
    token_mock,
    voting_mock,
):
    token_mock.approve(new_vote_bounty, 2**256 - 1, sender=alice)
    receipt = new_vote_bounty.openBounty(
        token_mock, AMOUNT, METADATA, SCRIPT, sender=alice, value=new_vote_bounty.OPEN_BOUNTY_COST()
    )
    created_at = receipt.timestamp

    receipt = voting_mock.new_vote(METADATA, SCRIPT, sender=bob)
    header_rlp = get_block_header_rlp(receipt.block_number)
    index, proof_rlp = get_receipt_proof_rlp(receipt.txn_hash)

    new_vote_bounty.claimBounty(
        alice, created_at, token_mock, DIGEST, index, header_rlp, proof_rlp, sender=bob
    )


def test_refund_success(alice, new_vote_bounty):
    balance_before = alice.balance

    new_vote_bounty.refund(alice, sender=alice)

    assert alice.balance == balance_before + new_vote_bounty.OPEN_BOUNTY_COST()


def test_refund_fails(bob, new_vote_bounty):
    with ape.reverts():
        new_vote_bounty.refund(bob, sender=bob)
