<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [hyperlane/igp/v1/types.proto](#hyperlane/igp/v1/types.proto)
    - [GasOracleConfig](#hyperlane.igp.v1.GasOracleConfig)
  
- [hyperlane/igp/v1/genesis.proto](#hyperlane/igp/v1/genesis.proto)
    - [GenesisState](#hyperlane.igp.v1.GenesisState)
  
- [hyperlane/igp/v1/query.proto](#hyperlane/igp/v1/query.proto)
    - [GetBeneficiaryRequest](#hyperlane.igp.v1.GetBeneficiaryRequest)
    - [GetBeneficiaryResponse](#hyperlane.igp.v1.GetBeneficiaryResponse)
    - [GetExchangeRateAndGasPriceRequest](#hyperlane.igp.v1.GetExchangeRateAndGasPriceRequest)
    - [GetExchangeRateAndGasPriceResponse](#hyperlane.igp.v1.GetExchangeRateAndGasPriceResponse)
    - [QuoteGasPaymentRequest](#hyperlane.igp.v1.QuoteGasPaymentRequest)
    - [QuoteGasPaymentResponse](#hyperlane.igp.v1.QuoteGasPaymentResponse)
  
    - [Query](#hyperlane.igp.v1.Query)
  
- [hyperlane/igp/v1/tx.proto](#hyperlane/igp/v1/tx.proto)
    - [MsgClaim](#hyperlane.igp.v1.MsgClaim)
    - [MsgClaimResponse](#hyperlane.igp.v1.MsgClaimResponse)
    - [MsgPayForGas](#hyperlane.igp.v1.MsgPayForGas)
    - [MsgPayForGasResponse](#hyperlane.igp.v1.MsgPayForGasResponse)
    - [MsgSetBeneficiary](#hyperlane.igp.v1.MsgSetBeneficiary)
    - [MsgSetBeneficiaryResponse](#hyperlane.igp.v1.MsgSetBeneficiaryResponse)
    - [MsgSetDestinationGasOverhead](#hyperlane.igp.v1.MsgSetDestinationGasOverhead)
    - [MsgSetDestinationGasOverheadResponse](#hyperlane.igp.v1.MsgSetDestinationGasOverheadResponse)
    - [MsgSetGasOracles](#hyperlane.igp.v1.MsgSetGasOracles)
    - [MsgSetGasOraclesResponse](#hyperlane.igp.v1.MsgSetGasOraclesResponse)
  
    - [Msg](#hyperlane.igp.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="hyperlane/igp/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/igp/v1/types.proto



<a name="hyperlane.igp.v1.GasOracleConfig"></a>

### GasOracleConfig
Hyperlane's tree


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `remote_domain` | [uint32](#uint32) |  | Count of items inserted to tree |
| `gas_oracle` | [string](#string) |  | Address of the oracle |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/igp/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/igp/v1/genesis.proto



<a name="hyperlane.igp.v1.GenesisState"></a>

### GenesisState
Hyperlane InterchainGasPaymaster's keeper genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_oracles` | [GasOracleConfig](#hyperlane.igp.v1.GasOracleConfig) | repeated | Every gas oracle that has been set |
| `beneficiary` | [string](#string) |  | Beneficiary |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/igp/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/igp/v1/query.proto



<a name="hyperlane.igp.v1.GetBeneficiaryRequest"></a>

### GetBeneficiaryRequest
GetBeneficiaryRequest is the request type for the Query/Tree RPC method.






<a name="hyperlane.igp.v1.GetBeneficiaryResponse"></a>

### GetBeneficiaryResponse
GetBeneficiaryResponse is the response type for the Query/Tree RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="hyperlane.igp.v1.GetExchangeRateAndGasPriceRequest"></a>

### GetExchangeRateAndGasPriceRequest
GetExchangeRateAndGasPriceRequest is the request type for the Query/Tree RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `destination_domain` | [uint32](#uint32) |  |  |






<a name="hyperlane.igp.v1.GetExchangeRateAndGasPriceResponse"></a>

### GetExchangeRateAndGasPriceResponse
GetExchangeRateAndGasPriceResponse is the response type for the Query/Tree
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exchange_rate` | [string](#string) |  |  |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="hyperlane.igp.v1.QuoteGasPaymentRequest"></a>

### QuoteGasPaymentRequest
QuoteGasPaymentRequest is the request type for the Query/Tree RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `destination_domain` | [uint32](#uint32) |  |  |
| `gas_amount` | [string](#string) |  |  |






<a name="hyperlane.igp.v1.QuoteGasPaymentResponse"></a>

### QuoteGasPaymentResponse
QuoteGasPaymentResponse is the response type for the Query/Tree RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.igp.v1.Query"></a>

### Query
Query service for hyperlane igp module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GetBeneficiary` | [GetBeneficiaryRequest](#hyperlane.igp.v1.GetBeneficiaryRequest) | [GetBeneficiaryResponse](#hyperlane.igp.v1.GetBeneficiaryResponse) | Gets the beneficiary | GET|/hyperlane/igp/v1/get_beneficiary|
| `QuoteGasPayment` | [QuoteGasPaymentRequest](#hyperlane.igp.v1.QuoteGasPaymentRequest) | [QuoteGasPaymentResponse](#hyperlane.igp.v1.QuoteGasPaymentResponse) | Quotes the amount of native tokens to pay for interchain gas. | GET|/hyperlane/igp/v1/quote_gas_payment|
| `GetExchangeRateAndGasPrice` | [GetExchangeRateAndGasPriceRequest](#hyperlane.igp.v1.GetExchangeRateAndGasPriceRequest) | [GetExchangeRateAndGasPriceResponse](#hyperlane.igp.v1.GetExchangeRateAndGasPriceResponse) | Gets the token exchange rate and gas price from the configured gas oracle for a given destination domain. | GET|/hyperlane/igp/v1/get_exchange_rate_and_gas_price|

 <!-- end services -->



<a name="hyperlane/igp/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/igp/v1/tx.proto



<a name="hyperlane.igp.v1.MsgClaim"></a>

### MsgClaim
MsgClaim defines the request type for the Claim rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |






<a name="hyperlane.igp.v1.MsgClaimResponse"></a>

### MsgClaimResponse
MsgClaimResponse defines the Claim response type.






<a name="hyperlane.igp.v1.MsgPayForGas"></a>

### MsgPayForGas
MsgPayForGas defines the request type for the Dispatch rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `message_id` | [string](#string) |  |  |
| `destination_domain` | [uint32](#uint32) |  |  |
| `gas_amount` | [string](#string) |  |  |
| `refund_address` | [string](#string) |  |  |
| `relayer_address` | [string](#string) |  |  |






<a name="hyperlane.igp.v1.MsgPayForGasResponse"></a>

### MsgPayForGasResponse
MsgPayForGasResponse defines the PayForGas response type.






<a name="hyperlane.igp.v1.MsgSetBeneficiary"></a>

### MsgSetBeneficiary
MsgSetBeneficiary defines the request type for the SetBeneficiary rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="hyperlane.igp.v1.MsgSetBeneficiaryResponse"></a>

### MsgSetBeneficiaryResponse
MsgSetBeneficiaryResponse defines the Claim response type.






<a name="hyperlane.igp.v1.MsgSetDestinationGasOverhead"></a>

### MsgSetDestinationGasOverhead
MsgSetDestinationGasOverhead defines the overhead gas amount for the given
destination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `destination_domain` | [uint32](#uint32) |  |  |
| `gas_overhead` | [string](#string) |  |  |






<a name="hyperlane.igp.v1.MsgSetDestinationGasOverheadResponse"></a>

### MsgSetDestinationGasOverheadResponse
MsgSetDestinationGasOverheadResponse defines the SetDestinationGasOverhead
response type.






<a name="hyperlane.igp.v1.MsgSetGasOracles"></a>

### MsgSetGasOracles
MsgSetGasOracles defines the request type for the Claim rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `configs` | [GasOracleConfig](#hyperlane.igp.v1.GasOracleConfig) | repeated |  |






<a name="hyperlane.igp.v1.MsgSetGasOraclesResponse"></a>

### MsgSetGasOraclesResponse
MsgSetGasOraclesResponse defines the Claim response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.igp.v1.Msg"></a>

### Msg
Msg defines the hyperlane igp Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PayForGas` | [MsgPayForGas](#hyperlane.igp.v1.MsgPayForGas) | [MsgPayForGasResponse](#hyperlane.igp.v1.MsgPayForGasResponse) | Deposits a payment for the relaying of a message to its destination chain. | |
| `Claim` | [MsgClaim](#hyperlane.igp.v1.MsgClaim) | [MsgClaimResponse](#hyperlane.igp.v1.MsgClaimResponse) | Transfers the entire native token balance to the beneficiary. | |
| `SetGasOracles` | [MsgSetGasOracles](#hyperlane.igp.v1.MsgSetGasOracles) | [MsgSetGasOraclesResponse](#hyperlane.igp.v1.MsgSetGasOraclesResponse) | Sets the gas oracles for remote domains specified in the config array. | |
| `SetBeneficiary` | [MsgSetBeneficiary](#hyperlane.igp.v1.MsgSetBeneficiary) | [MsgSetBeneficiaryResponse](#hyperlane.igp.v1.MsgSetBeneficiaryResponse) | Sets the beneficiary. | |
| `SetDestinationGasOverhead` | [MsgSetDestinationGasOverhead](#hyperlane.igp.v1.MsgSetDestinationGasOverhead) | [MsgSetDestinationGasOverheadResponse](#hyperlane.igp.v1.MsgSetDestinationGasOverheadResponse) | Sets the overhead gas for the destination domain. This is in the destination gas denom and will be added to the required payForGas payment. | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

