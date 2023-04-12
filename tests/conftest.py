import functools

import pytest
import rlp


@functools.singledispatch
def unhexlify(obj):
    raise NotImplementedError()


@unhexlify.register(str)
def _(hexstr):
    match hexstr:
        case "0x":
            return b""
        case "0x0":
            return b"\x00"
        case _ if len(hexstr := hexstr[2:]) % 2 != 0:
            hexstr = "0" + hexstr

    return bytes.fromhex(hexstr)


@unhexlify.register(list)
def _(obj):
    return list(map(unhexlify, obj))


@unhexlify.register(tuple)
def _(obj):
    return tuple(map(unhexlify, obj))


def encode_block(block):
    header = [
        block["parentHash"],
        block["sha3Uncles"],
        block["miner"],
        block["stateRoot"],
        block["transactionsRoot"],
        block["receiptsRoot"],
        block["logsBloom"],
        block["difficulty"],
        block["number"],
        block["gasLimit"],
        block["gasUsed"],
        block["timestamp"],
        block["extraData"],
        block["mixHash"],
        block["nonce"],
    ]

    tmp = []
    if withdrawals_root := block.get("withdrawlsRoot"):
        tmp.insert(0, withdrawals_root or "0x")

    if (base_fee_per_gas := block.get("baseFeePerGas")) or withdrawals_root:
        tmp.insert(0, base_fee_per_gas or "0x")

    return rlp.encode(unhexlify(header + tmp))


def encode_receipt(receipt):
    fields = [
        receipt.get("root") or receipt["status"],
        receipt["cumulativeGasUsed"],
        receipt["logsBloom"],
        [(log["address"], log["topics"], log["data"]) for log in receipt["logs"]],
    ]

    return unhexlify(receipt.get("type", "0x")) + rlp.encode(unhexlify(fields))


@pytest.fixture(scope="session")
def alice(accounts):
    yield accounts[0]


@pytest.fixture(scope="session")
def bob(accounts):
    yield accounts[1]


@pytest.fixture(scope="session")
def charlie(accounts):
    yield accounts[2]


@pytest.fixture(scope="module")
def token_mock(alice, project):
    yield project.TokenMock.deploy(10**24, sender=alice)


@pytest.fixture(scope="module")
def voting_mock(alice, project):
    yield project.VotingMock.deploy(sender=alice)


@pytest.fixture(scope="module")
def new_vote_bounty(alice, project, voting_mock):
    yield project.NewVoteBounty.deploy(voting_mock, sender=alice)
