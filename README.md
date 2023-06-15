# hyperlane-cosmos
[![Conforms to README.lint](https://img.shields.io/badge/README.lint-conforming-brightgreen)](https://github.com/strangelove-ventures/readme-dot-lint)

🌌 Why use hyperlane-cosmos
=============================
 * The [Cosmos-SDK](https://github.com/cosmos/cosmos-sdk) ecosystem is the internet of blockchains.
 * [Hyperlane](https://www.hyperlane.xyz/) is another protocol for blockchain interoperability, like [IBC](https://github.com/cosmos/ibc).
 * This repo provides a reference implementation for integrating and maintaining Hyperlane modules.

🌌🌌 Who benefits from it?
=============================
Cosmos-SDK based chain builders and integrators looking to add hyperlane connectivity to their chain will find this repository useful.

🌌🌌🌌 What does hyperlane-cosmos do?
=============================
[Hyperlane](https://www.hyperlane.xyz/) is an API to communicate between blockchains easily and securely, leverage the Hyperlane SDK to quickly build interchain applications, or bring interoperability to any new chain.

This repository provides modules integrated into a reference chain to demonstrate mailboxes, isms, and other hyperlane features.

🌌🌌🌌🌌 How do I use it
=============================
At present, `hyperlane-cosmos` serves as the collaboration space for developing the modules.

In the near future, this repository will include a module integration guide for chain maintainers to be able to introduce hyperlane into their chain.

🌌🌌🌌🌌🌌 Extras
=============================

## Geting Started

You will need go v1.19 or higher.

See go's documentation on how to [download and install](https://go.dev/doc/install)

### Clone the repository
```
cd $HOME
git clone https://github.com/strangelove-ventures/hyperlane-cosmos
cd hyperlane-cosmos
```

## Running Interchain Tests

To run interchain tests you will need to install [Docker](https://docs.docker.com/engine/install/)

Additionally, you will need heighliner to build the `hyperlane-simd` image.

### Install Heighliner
To build the `hyperlane-simd` image for interchain test you will need heighliner:

```
cd $HOME
git clone https://github.com/strangelove-ventures/heighliner
cd heighliner
go install ./...
```

### Build the hyperlane-simd image:
From within the `hyperlane-cosmos` directory, `heighliner` can build a `hyperlane-simd:local` image:

```
cd $HOME/hyperlane-cosmos
heighliner build --chain hyperlane-simd --local
```

### Run tests
From within the `hyperlane-cosmos` directory, the unit tests and interchain tests may be executed:

```
cd $HOME/hyperlane-cosmos
make test
```