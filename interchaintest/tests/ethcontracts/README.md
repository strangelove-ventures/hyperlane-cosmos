# Purpose

The files in the 'ethcontracts' folder are Go bindings to Solidity smart contracts.
We deploy these smart contracts to an Avalanche C-Chain in our Hyperlane test network.

## Generating Go bindings for hyperlane solidity contracts

1. Open the hyperlane monorepo in VSCode, e.g. from github.com/strangelove-ventures/hyperlane-monorepo.
2. Install the solidity extension for VSCode.
3. Right click on a hyperlane solidity contract you wish to generate bindings for, e.g. Mailbox.sol.
4. Compile the contract with the menu option "Solidity: Compile contract".
5. The contract ABI should show up in the monorepo under bin/contracts.
6. Run e.g. abigen --abi bin/contracts/Mailbox.abi --bin bin/contracts/Mailbox.bin --pkg mailbox --out mailbox.go.