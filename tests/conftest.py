import pytest


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
