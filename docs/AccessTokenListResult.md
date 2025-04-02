# AccessTokenListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]UserAccessToken**](UserAccessToken.md) |  | [optional] 

## Methods

### NewAccessTokenListResult

`func NewAccessTokenListResult() *AccessTokenListResult`

NewAccessTokenListResult instantiates a new AccessTokenListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenListResultWithDefaults

`func NewAccessTokenListResultWithDefaults() *AccessTokenListResult`

NewAccessTokenListResultWithDefaults instantiates a new AccessTokenListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AccessTokenListResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccessTokenListResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccessTokenListResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AccessTokenListResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *AccessTokenListResult) GetData() []UserAccessToken`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccessTokenListResult) GetDataOk() (*[]UserAccessToken, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccessTokenListResult) SetData(v []UserAccessToken)`

SetData sets Data field to given value.

### HasData

`func (o *AccessTokenListResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AccessTokenListResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AccessTokenListResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


