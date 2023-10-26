<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [hyperlane/mailbox/v1/types.proto](#hyperlane/mailbox/v1/types.proto)
    - [MessageDelivered](#hyperlane.mailbox.v1.MessageDelivered)
    - [Tree](#hyperlane.mailbox.v1.Tree)
  
- [hyperlane/mailbox/v1/genesis.proto](#hyperlane/mailbox/v1/genesis.proto)
    - [GenesisState](#hyperlane.mailbox.v1.GenesisState)
  
- [hyperlane/mailbox/v1/query.proto](#hyperlane/mailbox/v1/query.proto)
    - [QueryCurrentTreeMetadataRequest](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataRequest)
    - [QueryCurrentTreeMetadataResponse](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataResponse)
    - [QueryCurrentTreeRequest](#hyperlane.mailbox.v1.QueryCurrentTreeRequest)
    - [QueryCurrentTreeResponse](#hyperlane.mailbox.v1.QueryCurrentTreeResponse)
    - [QueryDomainRequest](#hyperlane.mailbox.v1.QueryDomainRequest)
    - [QueryDomainResponse](#hyperlane.mailbox.v1.QueryDomainResponse)
    - [QueryMsgDeliveredRequest](#hyperlane.mailbox.v1.QueryMsgDeliveredRequest)
    - [QueryMsgDeliveredResponse](#hyperlane.mailbox.v1.QueryMsgDeliveredResponse)
  
    - [Query](#hyperlane.mailbox.v1.Query)
  
- [hyperlane/mailbox/v1/tx.proto](#hyperlane/mailbox/v1/tx.proto)
    - [MsgDispatch](#hyperlane.mailbox.v1.MsgDispatch)
    - [MsgDispatchResponse](#hyperlane.mailbox.v1.MsgDispatchResponse)
    - [MsgProcess](#hyperlane.mailbox.v1.MsgProcess)
    - [MsgProcessResponse](#hyperlane.mailbox.v1.MsgProcessResponse)
  
    - [Msg](#hyperlane.mailbox.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="hyperlane/mailbox/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/mailbox/v1/types.proto



<a name="hyperlane.mailbox.v1.MessageDelivered"></a>

### MessageDelivered
Mailbox delivered message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  | Message id (hash) |






<a name="hyperlane.mailbox.v1.Tree"></a>

### Tree
Hyperlane's tree


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `count` | [uint32](#uint32) |  | Count of items inserted to tree |
| `branch` | [bytes](#bytes) | repeated | Each item inserted |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/mailbox/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/mailbox/v1/genesis.proto



<a name="hyperlane.mailbox.v1.GenesisState"></a>

### GenesisState
Hyperlane mailbox's keeper genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tree` | [Tree](#hyperlane.mailbox.v1.Tree) |  | Each genesis tree entry |
| `delivered_messages` | [MessageDelivered](#hyperlane.mailbox.v1.MessageDelivered) | repeated | Each message that has been delivered |
| `domain` | [uint32](#uint32) |  | The domain of this chain module, assigned by hyperlane |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/mailbox/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/mailbox/v1/query.proto



<a name="hyperlane.mailbox.v1.QueryCurrentTreeMetadataRequest"></a>

### QueryCurrentTreeMetadataRequest
QueryCurrentTreeMetadataRequest is the request type for the Query/Tree
metadata RPC method.






<a name="hyperlane.mailbox.v1.QueryCurrentTreeMetadataResponse"></a>

### QueryCurrentTreeMetadataResponse
QueryCurrentTreeResponseResponse is the response type for the Query/Tree
metadata RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `root` | [bytes](#bytes) |  |  |
| `count` | [uint32](#uint32) |  |  |






<a name="hyperlane.mailbox.v1.QueryCurrentTreeRequest"></a>

### QueryCurrentTreeRequest
QueryCurrentTreeRequest is the request type for the Query/Tree RPC method






<a name="hyperlane.mailbox.v1.QueryCurrentTreeResponse"></a>

### QueryCurrentTreeResponse
QueryCurrentTreeResponse is the response type for the Query/Tree RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `branches` | [bytes](#bytes) | repeated |  |
| `count` | [uint32](#uint32) |  |  |






<a name="hyperlane.mailbox.v1.QueryDomainRequest"></a>

### QueryDomainRequest
QueryDomain is the request type for the Query/Domain RPC
method.






<a name="hyperlane.mailbox.v1.QueryDomainResponse"></a>

### QueryDomainResponse
QueryDomainResponse is the response type for the Query/Domain RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `domain` | [uint32](#uint32) |  |  |






<a name="hyperlane.mailbox.v1.QueryMsgDeliveredRequest"></a>

### QueryMsgDeliveredRequest
QueryMsgDeliveredRequest is the request type to check if message was
delivered


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message_id` | [bytes](#bytes) |  |  |






<a name="hyperlane.mailbox.v1.QueryMsgDeliveredResponse"></a>

### QueryMsgDeliveredResponse
QueryMsgDeliveredResponse is the response type if message was delivered


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delivered` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.mailbox.v1.Query"></a>

### Query
Query service for hyperlane mailbox module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CurrentTreeMetadata` | [QueryCurrentTreeMetadataRequest](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataRequest) | [QueryCurrentTreeMetadataResponse](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataResponse) | Get current tree metadata | GET|/hyperlane/mailbox/v1/tree_metadata|
| `Domain` | [QueryDomainRequest](#hyperlane.mailbox.v1.QueryDomainRequest) | [QueryDomainResponse](#hyperlane.mailbox.v1.QueryDomainResponse) | Get domain | GET|/hyperlane/mailbox/v1/domain|
| `CurrentTree` | [QueryCurrentTreeRequest](#hyperlane.mailbox.v1.QueryCurrentTreeRequest) | [QueryCurrentTreeResponse](#hyperlane.mailbox.v1.QueryCurrentTreeResponse) | Get current tree | GET|/hyperlane/mailbox/v1/tree|
| `MsgDelivered` | [QueryMsgDeliveredRequest](#hyperlane.mailbox.v1.QueryMsgDeliveredRequest) | [QueryMsgDeliveredResponse](#hyperlane.mailbox.v1.QueryMsgDeliveredResponse) | Check if message was delivered | GET|/hyperlane/mailbox/v1/delivered|

 <!-- end services -->



<a name="hyperlane/mailbox/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/mailbox/v1/tx.proto



<a name="hyperlane.mailbox.v1.MsgDispatch"></a>

### MsgDispatch
MsgDispatch defines the request type for the Dispatch rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `destination_domain` | [uint32](#uint32) |  |  |
| `recipient_address` | [string](#string) |  |  |
| `message_body` | [string](#string) |  |  |






<a name="hyperlane.mailbox.v1.MsgDispatchResponse"></a>

### MsgDispatchResponse
MsgDispatchResponse defines the Dispatch response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message_id` | [string](#string) |  |  |






<a name="hyperlane.mailbox.v1.MsgProcess"></a>

### MsgProcess
MsgProcess defines the request type for the Process rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `metadata` | [string](#string) |  |  |
| `message` | [string](#string) |  |  |






<a name="hyperlane.mailbox.v1.MsgProcessResponse"></a>

### MsgProcessResponse
MsgProcessResponse defines the Process response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.mailbox.v1.Msg"></a>

### Msg
Msg defines the hyperlane mailbox Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Dispatch` | [MsgDispatch](#hyperlane.mailbox.v1.MsgDispatch) | [MsgDispatchResponse](#hyperlane.mailbox.v1.MsgDispatchResponse) | Dispatch sends interchain messages | |
| `Process` | [MsgProcess](#hyperlane.mailbox.v1.MsgProcess) | [MsgProcessResponse](#hyperlane.mailbox.v1.MsgProcessResponse) | Process delivers interchain messages | |

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

