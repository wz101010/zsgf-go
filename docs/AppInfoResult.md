# AppInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**AppInfoItem**](AppInfoItem.md) |  | [optional] 
**Props** | Pointer to [**[]AppProperty**](AppProperty.md) |  | [optional] 

## Methods

### NewAppInfoResult

`func NewAppInfoResult() *AppInfoResult`

NewAppInfoResult instantiates a new AppInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoResultWithDefaults

`func NewAppInfoResultWithDefaults() *AppInfoResult`

NewAppInfoResultWithDefaults instantiates a new AppInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *AppInfoResult) GetInfo() AppInfoItem`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AppInfoResult) GetInfoOk() (*AppInfoItem, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AppInfoResult) SetInfo(v AppInfoItem)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AppInfoResult) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetProps

`func (o *AppInfoResult) GetProps() []AppProperty`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *AppInfoResult) GetPropsOk() (*[]AppProperty, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *AppInfoResult) SetProps(v []AppProperty)`

SetProps sets Props field to given value.

### HasProps

`func (o *AppInfoResult) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetPropsNil

`func (o *AppInfoResult) SetPropsNil(b bool)`

 SetPropsNil sets the value for Props to be an explicit nil

### UnsetProps
`func (o *AppInfoResult) UnsetProps()`

UnsetProps ensures that no value is present for Props, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


