# Protobufs

This is the public protocol buffers API for [Hyperlane](https://github.com/strangelove-ventures/hyperlane-cosmos).

## Download

The `buf` CLI comes with an export command. Use `buf export -h` for details

#### Examples:

Download hyperlane-cosmos protos for a commit:
```bash
buf export buf.build/strangelove-ventures/hyperlane-cosmos:${commit} --output ./tmp
```

Download all project protos:
```bash
buf export . --output ./tmp
```