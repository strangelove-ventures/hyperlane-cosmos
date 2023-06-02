package bindings

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	mailboxkeeper "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/keeper"

)

func RegisterCustomPlugins(
	mailbox *mailboxkeeper.Keeper,
) []wasmkeeper.Option {
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(mailbox),
	)

	return []wasm.Option{
		messengerDecoratorOpt,
	}
}