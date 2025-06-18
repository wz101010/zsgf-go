# UserLocationsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalLocations** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]GeoLocationResponseModel**](GeoLocationResponseModel.md) |  | [optional] 

## Methods

### NewUserLocationsResult

`func NewUserLocationsResult() *UserLocationsResult`

NewUserLocationsResult instantiates a new UserLocationsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLocationsResultWithDefaults

`func NewUserLocationsResultWithDefaults() *UserLocationsResult`

NewUserLocationsResultWithDefaults instantiates a new UserLocationsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalLocations

`func (o *UserLocationsResult) GetTotalLocations() int64`

GetTotalLocations returns the TotalLocations field if non-nil, zero value otherwise.

### GetTotalLocationsOk

`func (o *UserLocationsResult) GetTotalLocationsOk() (*int64, bool)`

GetTotalLocationsOk returns a tuple with the TotalLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocations

`func (o *UserLocationsResult) SetTotalLocations(v int64)`

SetTotalLocations sets TotalLocations field to given value.

### HasTotalLocations

`func (o *UserLocationsResult) HasTotalLocations() bool`

HasTotalLocations returns a boolean if a field has been set.

### GetData

`func (o *UserLocationsResult) GetData() []GeoLocationResponseModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserLocationsResult) GetDataOk() (*[]GeoLocationResponseModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserLocationsResult) SetData(v []GeoLocationResponseModel)`

SetData sets Data field to given value.

### HasData

`func (o *UserLocationsResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UserLocationsResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserLocationsResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


