# MySqlGeometry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XCoordinate** | Pointer to **NullableFloat64** |  | [optional] [readonly] 
**YCoordinate** | Pointer to **NullableFloat64** |  | [optional] [readonly] 
**Srid** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**IsNull** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewMySqlGeometry

`func NewMySqlGeometry() *MySqlGeometry`

NewMySqlGeometry instantiates a new MySqlGeometry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySqlGeometryWithDefaults

`func NewMySqlGeometryWithDefaults() *MySqlGeometry`

NewMySqlGeometryWithDefaults instantiates a new MySqlGeometry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXCoordinate

`func (o *MySqlGeometry) GetXCoordinate() float64`

GetXCoordinate returns the XCoordinate field if non-nil, zero value otherwise.

### GetXCoordinateOk

`func (o *MySqlGeometry) GetXCoordinateOk() (*float64, bool)`

GetXCoordinateOk returns a tuple with the XCoordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXCoordinate

`func (o *MySqlGeometry) SetXCoordinate(v float64)`

SetXCoordinate sets XCoordinate field to given value.

### HasXCoordinate

`func (o *MySqlGeometry) HasXCoordinate() bool`

HasXCoordinate returns a boolean if a field has been set.

### SetXCoordinateNil

`func (o *MySqlGeometry) SetXCoordinateNil(b bool)`

 SetXCoordinateNil sets the value for XCoordinate to be an explicit nil

### UnsetXCoordinate
`func (o *MySqlGeometry) UnsetXCoordinate()`

UnsetXCoordinate ensures that no value is present for XCoordinate, not even an explicit nil
### GetYCoordinate

`func (o *MySqlGeometry) GetYCoordinate() float64`

GetYCoordinate returns the YCoordinate field if non-nil, zero value otherwise.

### GetYCoordinateOk

`func (o *MySqlGeometry) GetYCoordinateOk() (*float64, bool)`

GetYCoordinateOk returns a tuple with the YCoordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYCoordinate

`func (o *MySqlGeometry) SetYCoordinate(v float64)`

SetYCoordinate sets YCoordinate field to given value.

### HasYCoordinate

`func (o *MySqlGeometry) HasYCoordinate() bool`

HasYCoordinate returns a boolean if a field has been set.

### SetYCoordinateNil

`func (o *MySqlGeometry) SetYCoordinateNil(b bool)`

 SetYCoordinateNil sets the value for YCoordinate to be an explicit nil

### UnsetYCoordinate
`func (o *MySqlGeometry) UnsetYCoordinate()`

UnsetYCoordinate ensures that no value is present for YCoordinate, not even an explicit nil
### GetSrid

`func (o *MySqlGeometry) GetSrid() int32`

GetSrid returns the Srid field if non-nil, zero value otherwise.

### GetSridOk

`func (o *MySqlGeometry) GetSridOk() (*int32, bool)`

GetSridOk returns a tuple with the Srid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrid

`func (o *MySqlGeometry) SetSrid(v int32)`

SetSrid sets Srid field to given value.

### HasSrid

`func (o *MySqlGeometry) HasSrid() bool`

HasSrid returns a boolean if a field has been set.

### SetSridNil

`func (o *MySqlGeometry) SetSridNil(b bool)`

 SetSridNil sets the value for Srid to be an explicit nil

### UnsetSrid
`func (o *MySqlGeometry) UnsetSrid()`

UnsetSrid ensures that no value is present for Srid, not even an explicit nil
### GetIsNull

`func (o *MySqlGeometry) GetIsNull() bool`

GetIsNull returns the IsNull field if non-nil, zero value otherwise.

### GetIsNullOk

`func (o *MySqlGeometry) GetIsNullOk() (*bool, bool)`

GetIsNullOk returns a tuple with the IsNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNull

`func (o *MySqlGeometry) SetIsNull(v bool)`

SetIsNull sets IsNull field to given value.

### HasIsNull

`func (o *MySqlGeometry) HasIsNull() bool`

HasIsNull returns a boolean if a field has been set.

### GetValue

`func (o *MySqlGeometry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MySqlGeometry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MySqlGeometry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MySqlGeometry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MySqlGeometry) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MySqlGeometry) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


