#!/usr/bin/env bash

mockgen_cmd="mockgen"
$mockgen_cmd -source=x/announce/types/expected_keepers.go -package testutil -destination x/announce/testutil/expected_keepers_mocks.go