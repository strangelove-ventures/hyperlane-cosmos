<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [hyperlane/ism/v1/multisig.proto](#hyperlane/ism/v1/multisig.proto)
    - [MultiSigIsm](#hyperlane.ism.v1.MultiSigIsm)
    - [OriginsMultiSigIsm](#hyperlane.ism.v1.OriginsMultiSigIsm)
  
- [hyperlane/ism/v1/genesis.proto](#hyperlane/ism/v1/genesis.proto)
    - [GenesisState](#hyperlane.ism.v1.GenesisState)
  
- [hyperlane/ism/v1/query.proto](#hyperlane/ism/v1/query.proto)
    - [QueryAllDefaultIsmsRequest](#hyperlane.ism.v1.QueryAllDefaultIsmsRequest)
    - [QueryAllDefaultIsmsResponse](#hyperlane.ism.v1.QueryAllDefaultIsmsResponse)
    - [QueryOriginsDefaultIsmRequest](#hyperlane.ism.v1.QueryOriginsDefaultIsmRequest)
    - [QueryOriginsDefaultIsmResponse](#hyperlane.ism.v1.QueryOriginsDefaultIsmResponse)
  
    - [Query](#hyperlane.ism.v1.Query)
  
- [hyperlane/ism/v1/tx.proto](#hyperlane/ism/v1/tx.proto)
    - [MsgSetDefaultIsm](#hyperlane.ism.v1.MsgSetDefaultIsm)
    - [MsgSetDefaultIsmResponse](#hyperlane.ism.v1.MsgSetDefaultIsmResponse)
  
    - [Msg](#hyperlane.ism.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="hyperlane/ism/v1/multisig.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/ism/v1/multisig.proto



<a name="hyperlane.ism.v1.MultiSigIsm"></a>

### MultiSigIsm
MultiSig ISM for a specific origin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_pub_keys` | [bytes](#bytes) | repeated | Validator pub keys |
| `threshold` | [uint32](#uint32) |  | number of validators required |






<a name="hyperlane.ism.v1.OriginsMultiSigIsm"></a>

### OriginsMultiSigIsm
Hyperlane's default ISM


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `origin` | [uint32](#uint32) |  |  |
| `ism` | [MultiSigIsm](#hyperlane.ism.v1.MultiSigIsm) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/ism/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/ism/v1/genesis.proto



<a name="hyperlane.ism.v1.GenesisState"></a>

### GenesisState
Hyperlane ISM's keeper genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `default_ism` | [OriginsMultiSigIsm](#hyperlane.ism.v1.OriginsMultiSigIsm) | repeated | Genesis default ISM |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/ism/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/ism/v1/query.proto



<a name="hyperlane.ism.v1.QueryAllDefaultIsmsRequest"></a>

### QueryAllDefaultIsmsRequest
QueryAllDefaultIsmRequest is the request type for the AllDefaultIsms RPC
method.






<a name="hyperlane.ism.v1.QueryAllDefaultIsmsResponse"></a>

### QueryAllDefaultIsmsResponse
QueryAllDefaultIsmResponse is the response type for the AllDefaultIsms RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `default_isms` | [OriginsMultiSigIsm](#hyperlane.ism.v1.OriginsMultiSigIsm) | repeated |  |






<a name="hyperlane.ism.v1.QueryOriginsDefaultIsmRequest"></a>

### QueryOriginsDefaultIsmRequest
QueryDefaultIsmRequest is the request type for the DefaultIsm RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `origin` | [uint32](#uint32) |  |  |






<a name="hyperlane.ism.v1.QueryOriginsDefaultIsmResponse"></a>

### QueryOriginsDefaultIsmResponse
QueryDefaultIsmResponse is the response type for the DefaultIsm RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `default_ism` | [MultiSigIsm](#hyperlane.ism.v1.MultiSigIsm) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.ism.v1.Query"></a>

### Query
Query service for hyperlane ISM module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `OriginsDefaultIsm` | [QueryOriginsDefaultIsmRequest](#hyperlane.ism.v1.QueryOriginsDefaultIsmRequest) | [QueryOriginsDefaultIsmResponse](#hyperlane.ism.v1.QueryOriginsDefaultIsmResponse) | Get current default multisig ISM for an origin | GET|/hyperlane/ism/v1/origins_default_ism|
| `AllDefaultIsms` | [QueryAllDefaultIsmsRequest](#hyperlane.ism.v1.QueryAllDefaultIsmsRequest) | [QueryAllDefaultIsmsResponse](#hyperlane.ism.v1.QueryAllDefaultIsmsResponse) | Get all current default multisig ISMs | GET|/hyperlane/ism/v1/all_default_isms|

 <!-- end services -->



<a name="hyperlane/ism/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/ism/v1/tx.proto



<a name="hyperlane.ism.v1.MsgSetDefaultIsm"></a>

### MsgSetDefaultIsm
MsgSetDefaultIsm defines the request type for the SetDefaultIsm rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer` | [string](#string) |  |  |
| `isms` | [OriginsMultiSigIsm](#hyperlane.ism.v1.OriginsMultiSigIsm) | repeated |  |






<a name="hyperlane.ism.v1.MsgSetDefaultIsmResponse"></a>

### MsgSetDefaultIsmResponse
MsgSetDefaultIsmResponse defines the Msg/SetDefaultIsm response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.ism.v1.Msg"></a>

### Msg
Msg defines the hyperlane ISM Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetDefaultIsm` | [MsgSetDefaultIsm](#hyperlane.ism.v1.MsgSetDefaultIsm) | [MsgSetDefaultIsmResponse](#hyperlane.ism.v1.MsgSetDefaultIsmResponse) | SetDefaultIsm defines a rpc handler method for MsgSetDefaultIsm. | |

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

