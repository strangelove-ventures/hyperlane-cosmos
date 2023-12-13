package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var (
	_ codectypes.UnpackInterfacesMessage = QueryOriginsDefaultIsmResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryCustomIsmResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryAllDefaultIsmsResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryAllCustomIsmsResponse{}
)

func (r QueryOriginsDefaultIsmResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if r.DefaultIsm != nil {
		var ism AbstractIsm
		return unpacker.UnpackAny(r.DefaultIsm, &ism)
	}
	return nil
}

func (r QueryCustomIsmResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if r.CustomIsm != nil {
		var ism AbstractIsm
		return unpacker.UnpackAny(r.CustomIsm, &ism)
	}
	return nil
}

func (r QueryAllDefaultIsmsResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, defaultIsm := range r.DefaultIsms {
		ism := defaultIsm.AbstractIsm
		if ism != nil {
			var aIsm AbstractIsm
			err := unpacker.UnpackAny(ism, &aIsm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r QueryAllCustomIsmsResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, customIsm := range r.CustomIsms {
		ism := customIsm.AbstractIsm
		if ism != nil {
			var aIsm AbstractIsm
			err := unpacker.UnpackAny(ism, &aIsm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
