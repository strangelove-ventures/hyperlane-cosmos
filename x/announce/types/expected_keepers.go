package types

import context "context"

type MailboxKeeper interface {
	GetMailboxAddress() []byte
	GetDomain(context.Context) uint32
}
