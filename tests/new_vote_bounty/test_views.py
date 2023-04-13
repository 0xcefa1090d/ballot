import eth_abi
from eth_hash.auto import keccak


def test_calculate_identifier(alice, new_vote_bounty, token_mock):
    metadata = "Hello, World!"
    script = b"\xde\xad\xbe\xef"
    digest = keccak(keccak(metadata.encode()) + keccak(script))

    assert new_vote_bounty.calculateIdentifier(alice, 0, token_mock, metadata, script) == keccak(
        eth_abi.encode(
            ["address", "uint256", "address", "bytes32"],
            [alice.address, 0, token_mock.address, digest],
        )
    )
