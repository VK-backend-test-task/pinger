# PingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerIp** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Success** | **bool** |  | 

## Methods

### NewPingInfo

`func NewPingInfo(containerIp string, timestamp time.Time, success bool, ) *PingInfo`

NewPingInfo instantiates a new PingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingInfoWithDefaults

`func NewPingInfoWithDefaults() *PingInfo`

NewPingInfoWithDefaults instantiates a new PingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerIp

`func (o *PingInfo) GetContainerIp() string`

GetContainerIp returns the ContainerIp field if non-nil, zero value otherwise.

### GetContainerIpOk

`func (o *PingInfo) GetContainerIpOk() (*string, bool)`

GetContainerIpOk returns a tuple with the ContainerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIp

`func (o *PingInfo) SetContainerIp(v string)`

SetContainerIp sets ContainerIp field to given value.


### GetTimestamp

`func (o *PingInfo) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PingInfo) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PingInfo) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSuccess

`func (o *PingInfo) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PingInfo) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PingInfo) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


