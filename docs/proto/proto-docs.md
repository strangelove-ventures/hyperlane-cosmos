<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [hyperlane/announce/v1/types.proto](#hyperlane/announce/v1/types.proto)
    - [GenesisAnnouncement](#hyperlane.announce.v1.GenesisAnnouncement)
    - [StorageMetadata](#hyperlane.announce.v1.StorageMetadata)
    - [StoredAnnouncement](#hyperlane.announce.v1.StoredAnnouncement)
    - [StoredAnnouncements](#hyperlane.announce.v1.StoredAnnouncements)
  
- [hyperlane/announce/v1/genesis.proto](#hyperlane/announce/v1/genesis.proto)
    - [GenesisState](#hyperlane.announce.v1.GenesisState)
  
- [hyperlane/announce/v1/query.proto](#hyperlane/announce/v1/query.proto)
    - [GetAnnouncedStorageLocationsRequest](#hyperlane.announce.v1.GetAnnouncedStorageLocationsRequest)
    - [GetAnnouncedStorageLocationsResponse](#hyperlane.announce.v1.GetAnnouncedStorageLocationsResponse)
    - [GetAnnouncedValidatorsRequest](#hyperlane.announce.v1.GetAnnouncedValidatorsRequest)
    - [GetAnnouncedValidatorsResponse](#hyperlane.announce.v1.GetAnnouncedValidatorsResponse)
  
    - [Query](#hyperlane.announce.v1.Query)
  
- [hyperlane/announce/v1/tx.proto](#hyperlane/announce/v1/tx.proto)
    - [MsgAnnouncement](#hyperlane.announce.v1.MsgAnnouncement)
    - [MsgAnnouncementResponse](#hyperlane.announce.v1.MsgAnnouncementResponse)
  
    - [Msg](#hyperlane.announce.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="hyperlane/announce/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/announce/v1/types.proto



<a name="hyperlane.announce.v1.GenesisAnnouncement"></a>

### GenesisAnnouncement
Genesis helper type for Hyperlane's Announcement.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `announcement` | [StoredAnnouncement](#hyperlane.announce.v1.StoredAnnouncement) |  |  |
| `validator` | [bytes](#bytes) |  | The validator (in eth address format) that announced |






<a name="hyperlane.announce.v1.StorageMetadata"></a>

### StorageMetadata
Helper type for Hyperlane's getAnnouncedStorageLocations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadata` | [string](#string) | repeated |  |






<a name="hyperlane.announce.v1.StoredAnnouncement"></a>

### StoredAnnouncement
Helper type for Hyperlane's Announcement.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `storage_location` | [string](#string) |  | location where signatures will be stored |






<a name="hyperlane.announce.v1.StoredAnnouncements"></a>

### StoredAnnouncements
Helper type for Hyperlane's Announcement.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `announcement` | [StoredAnnouncement](#hyperlane.announce.v1.StoredAnnouncement) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/announce/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/announce/v1/genesis.proto



<a name="hyperlane.announce.v1.GenesisState"></a>

### GenesisState
Hyperlane Announce's keeper genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `announcements` | [GenesisAnnouncement](#hyperlane.announce.v1.GenesisAnnouncement) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/announce/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/announce/v1/query.proto



<a name="hyperlane.announce.v1.GetAnnouncedStorageLocationsRequest"></a>

### GetAnnouncedStorageLocationsRequest
GetAnnouncedStorageLocationsRequest is the request type for the
GetAnnouncedStorageLocations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [bytes](#bytes) | repeated | list of validators where each validator is in hex-encoded eth address format (20 bytes) |






<a name="hyperlane.announce.v1.GetAnnouncedStorageLocationsResponse"></a>

### GetAnnouncedStorageLocationsResponse
GetAnnouncedStorageLocationsResponse is the response type for the
GetAnnouncedStorageLocations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadata` | [StorageMetadata](#hyperlane.announce.v1.StorageMetadata) | repeated |  |






<a name="hyperlane.announce.v1.GetAnnouncedValidatorsRequest"></a>

### GetAnnouncedValidatorsRequest
GetAnnouncedValidatorsRequest is the request type for the
GetAnnouncedValidators RPC method.






<a name="hyperlane.announce.v1.GetAnnouncedValidatorsResponse"></a>

### GetAnnouncedValidatorsResponse
GetAnnouncedValidatorsResponse is the response type for the
GetAnnouncedValidators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [string](#string) | repeated | list of validators where each validator is in hex-encoded eth address format (20 bytes) |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.announce.v1.Query"></a>

### Query
Query service for hyperlane announce module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GetAnnouncedStorageLocations` | [GetAnnouncedStorageLocationsRequest](#hyperlane.announce.v1.GetAnnouncedStorageLocationsRequest) | [GetAnnouncedStorageLocationsResponse](#hyperlane.announce.v1.GetAnnouncedStorageLocationsResponse) | Gets the announced storage locations (where signatures are stored) for the requested validators | GET|/hyperlane/announce/v1/get_announced_storage_locations|
| `GetAnnouncedValidators` | [GetAnnouncedValidatorsRequest](#hyperlane.announce.v1.GetAnnouncedValidatorsRequest) | [GetAnnouncedValidatorsResponse](#hyperlane.announce.v1.GetAnnouncedValidatorsResponse) | Gets a list of validators that have made announcements | GET|/hyperlane/announce/v1/get_announced_validators|

 <!-- end services -->



<a name="hyperlane/announce/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/announce/v1/tx.proto



<a name="hyperlane.announce.v1.MsgAnnouncement"></a>

### MsgAnnouncement
MsgAnnouncement Announces a validator signature storage location


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `validator` | [bytes](#bytes) |  | The validator (in eth address format) that is announcing its storage location |
| `storage_location` | [string](#string) |  | location where signatures will be stored |
| `signature` | [bytes](#bytes) |  | signed validator announcement |






<a name="hyperlane.announce.v1.MsgAnnouncementResponse"></a>

### MsgAnnouncementResponse
MsgAnnouncementResponse defines the MsgAnnouncementResponse response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.announce.v1.Msg"></a>

### Msg
Msg defines the hyperlane announce Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Announcement` | [MsgAnnouncement](#hyperlane.announce.v1.MsgAnnouncement) | [MsgAnnouncementResponse](#hyperlane.announce.v1.MsgAnnouncementResponse) | Announces a validator signature storage location | |

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

