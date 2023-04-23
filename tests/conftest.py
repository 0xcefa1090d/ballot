import functools
import itertools

import pytest
import requests
import rlp
import trie

REQUEST_COUNTER = itertools.count()


@functools.singledispatch
def unhexlify(obj):
    raise NotImplementedError()


@unhexlify.register(str)
def _(hexstr):
    if hexstr in ("0x", "0x0"):
        return b""

    if len(hexstr := hexstr[2:]) % 2 != 0:
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

    tmp = [block.get("baseFeePerGas"), block.get("withdrawalsRoot")]
    for i, elem in enumerate(tmp[::-1]):
        if elem is not None:
            break
    else:
        i += 1
    tmp = [elem or "0x0" for elem in tmp[:-i]] if i != 0 else tmp

    return rlp.encode(unhexlify(header + tmp))


def encode_receipt(receipt):
    fields = [
        receipt.get("root") or receipt["status"],
        receipt["cumulativeGasUsed"],
        receipt["logsBloom"],
        [(log["address"], log["topics"], log["data"]) for log in receipt["logs"]],
    ]

    return unhexlify(receipt.get("type", "0x")) + rlp.encode(unhexlify(fields))


def encode_request(method, params):
    return {"jsonrpc": "2.0", "method": method, "params": params, "id": next(REQUEST_COUNTER)}


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
def start_vote_bounty(alice, project, voting_mock):
    yield project.StartVoteBounty.deploy(voting_mock, sender=alice)


@pytest.fixture(scope="session")
def get_block_header_rlp(chain):
    def _get_block_header(number):
        r = requests.post(
            chain.provider.uri, json=encode_request("eth_getBlockByNumber", [hex(number), False])
        )
        r.raise_for_status()
        return encode_block(r.json()["result"])

    return _get_block_header


@pytest.fixture(scope="session")
def get_receipt_proof_rlp(chain):
    def _get_receipt_proof_rlp(txn_hash):
        with requests.Session() as s:
            resp = s.post(
                chain.provider.uri, json=encode_request("eth_getTransactionReceipt", [txn_hash])
            )
            resp.raise_for_status()
            receipt = resp.json()["result"]
            index = int(receipt["transactionIndex"], 16)

            resp = s.post(
                chain.provider.uri,
                json=encode_request("eth_getBlockByHash", [receipt["blockHash"], False]),
            )
            resp.raise_for_status()
            block = resp.json()["result"]

            receipts_trie = trie.HexaryTrie({})
            resp = s.post(
                chain.provider.uri,
                json=[
                    encode_request("eth_getTransactionReceipt", [tx_hash])
                    for tx_hash in block["transactions"]
                ],
            )
            resp.raise_for_status()

            for receipt in (r["result"] for r in resp.json()):
                key = rlp.encode(unhexlify(receipt["transactionIndex"]))
                receipts_trie[key] = encode_receipt(receipt)

            return index, rlp.encode(receipts_trie.get_proof(rlp.encode(index)))

    return _get_receipt_proof_rlp
