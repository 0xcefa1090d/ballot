# bounty

A truly trustless bounty system for incentivizing governance participation

## Development

The following system dependencies are required:

- [Python (3.10 or greater)](https://www.python.org/)
- [Node.js LTS](https://nodejs.org/en)
- [npm](https://www.npmjs.com/)

To initialize a development environment, complete the following steps:

1. Clone the repository and change the working directory.

```bash
$ git clone https://github.com/0xcefa1090d/bounty.git
$ cd bounty
```

2. Create a virtualenv and activate it.

```bash
$ python -m venv venv
$ source venv/bin/activate
```

3. Install dependencies.

```bash
$ pip install -r requirements.txt
$ npm ci
```
4. Optionally, install pre-commit hooks.

```bash
$ pre-commit install
```
