# hyperlane-cosmos
[![Conforms to README.lint](https://img.shields.io/badge/README.lint-conforming-brightgreen)](https://github.com/strangelove-ventures/readme-dot-lint)

🌌 Why did we build hyperlane-cosmos?
=============================
 - How come [IBC](https://github.com/cosmos/ibc) gets to have all the Cosmos fun? This repo provides a reference implementation for integrating and maintaining [Hyperlane](https://www.hyperlane.xyz/) modules in [Cosmos](https://github.com/cosmos/cosmos-sdk) blockchains.

🌌🌌 Who benefits from it?
=============================
Cosmos-SDK based chain builders and integrators looking to add hyperlane connectivity to their chain.

🌌🌌🌌 What does hyperlane-cosmos do?
=============================
This repository provides modules integrated into a reference chain to demonstrate mailboxes, isms, and other hyperlane features.

🌌🌌🌌🌌 How do I use it?
=============================

## Getting Started

You will need `go v1.19` or higher.

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

🌌🌌🌌🌌🌌 Extras
=============================

### Roadmap
- [x] Build `hyperlane-cosmos` to serve as the collaboration space for developing the modules.
- [ ] include a module integration guide for chain maintainers to be able to introduce hyperlane into their chain.
