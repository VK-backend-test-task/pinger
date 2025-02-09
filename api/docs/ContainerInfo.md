# ContainerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**LastPing** | Pointer to **time.Time** |  | [optional] 
**LastSuccess** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContainerInfo

`func NewContainerInfo(ip string, ) *ContainerInfo`

NewContainerInfo instantiates a new ContainerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInfoWithDefaults

`func NewContainerInfoWithDefaults() *ContainerInfo`

NewContainerInfoWithDefaults instantiates a new ContainerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ContainerInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ContainerInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ContainerInfo) SetIp(v string)`

SetIp sets Ip field to given value.


### GetLastPing

`func (o *ContainerInfo) GetLastPing() time.Time`

GetLastPing returns the LastPing field if non-nil, zero value otherwise.

### GetLastPingOk

`func (o *ContainerInfo) GetLastPingOk() (*time.Time, bool)`

GetLastPingOk returns a tuple with the LastPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPing

`func (o *ContainerInfo) SetLastPing(v time.Time)`

SetLastPing sets LastPing field to given value.

### HasLastPing

`func (o *ContainerInfo) HasLastPing() bool`

HasLastPing returns a boolean if a field has been set.

### GetLastSuccess

`func (o *ContainerInfo) GetLastSuccess() time.Time`

GetLastSuccess returns the LastSuccess field if non-nil, zero value otherwise.

### GetLastSuccessOk

`func (o *ContainerInfo) GetLastSuccessOk() (*time.Time, bool)`

GetLastSuccessOk returns a tuple with the LastSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccess

`func (o *ContainerInfo) SetLastSuccess(v time.Time)`

SetLastSuccess sets LastSuccess field to given value.

### HasLastSuccess

`func (o *ContainerInfo) HasLastSuccess() bool`

HasLastSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


