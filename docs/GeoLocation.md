# GeoLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 唯一标识 | [optional] 
**BizCode** | **string** | 业务代码 | 
**BizID** | **int64** | 业务ID | 
**Coordinates** | [**MySqlGeometry**](MySqlGeometry.md) |  | 
**LocationName** | Pointer to **NullableString** | 地点的名称 | [optional] 
**LocationType** | Pointer to **NullableString** | 地点类型 | [optional] 
**RecipientName** | Pointer to **NullableString** | 收货人姓名 | [optional] 
**PhoneNumber** | Pointer to **NullableString** | 收货人联系电话 | [optional] 
**Email** | Pointer to **NullableString** | 收货人电子邮件 | [optional] 
**Country** | Pointer to **NullableString** | 国家 | [optional] 
**State** | Pointer to **NullableString** | 州/省 | [optional] 
**City** | Pointer to **NullableString** | 城市 | [optional] 
**District** | Pointer to **NullableString** | 区/县 | [optional] 
**Street** | Pointer to **NullableString** | 街道 | [optional] 
**ZipCode** | Pointer to **NullableString** | 邮政编码 | [optional] 
**Address** | Pointer to **NullableString** | 详细的地址信息 | [optional] 
**MapType** | Pointer to **NullableString** | 地址类型 | [optional] 
**Remark** | Pointer to **NullableString** | 备注 | [optional] 
**Tags** | Pointer to **NullableString** | 标签 | [optional] 
**Enable** | Pointer to **bool** | 是否启用 | [optional] 
**ShowIndex** | Pointer to **int32** | 排序索引 | [optional] 
**CreateDate** | Pointer to **time.Time** | 创建时间 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 最后更新时间 | [optional] 

## Methods

### NewGeoLocation

`func NewGeoLocation(bizCode string, bizID int64, coordinates MySqlGeometry, ) *GeoLocation`

NewGeoLocation instantiates a new GeoLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoLocationWithDefaults

`func NewGeoLocationWithDefaults() *GeoLocation`

NewGeoLocationWithDefaults instantiates a new GeoLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GeoLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeoLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeoLocation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GeoLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBizCode

`func (o *GeoLocation) GetBizCode() string`

GetBizCode returns the BizCode field if non-nil, zero value otherwise.

### GetBizCodeOk

`func (o *GeoLocation) GetBizCodeOk() (*string, bool)`

GetBizCodeOk returns a tuple with the BizCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizCode

`func (o *GeoLocation) SetBizCode(v string)`

SetBizCode sets BizCode field to given value.


### GetBizID

`func (o *GeoLocation) GetBizID() int64`

GetBizID returns the BizID field if non-nil, zero value otherwise.

### GetBizIDOk

`func (o *GeoLocation) GetBizIDOk() (*int64, bool)`

GetBizIDOk returns a tuple with the BizID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizID

`func (o *GeoLocation) SetBizID(v int64)`

SetBizID sets BizID field to given value.


### GetCoordinates

`func (o *GeoLocation) GetCoordinates() MySqlGeometry`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoLocation) GetCoordinatesOk() (*MySqlGeometry, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoLocation) SetCoordinates(v MySqlGeometry)`

SetCoordinates sets Coordinates field to given value.


### GetLocationName

`func (o *GeoLocation) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *GeoLocation) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *GeoLocation) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *GeoLocation) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *GeoLocation) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *GeoLocation) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetLocationType

`func (o *GeoLocation) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *GeoLocation) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *GeoLocation) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *GeoLocation) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *GeoLocation) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *GeoLocation) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetRecipientName

`func (o *GeoLocation) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *GeoLocation) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *GeoLocation) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *GeoLocation) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *GeoLocation) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *GeoLocation) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetPhoneNumber

`func (o *GeoLocation) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GeoLocation) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GeoLocation) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GeoLocation) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *GeoLocation) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *GeoLocation) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *GeoLocation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GeoLocation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GeoLocation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GeoLocation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GeoLocation) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GeoLocation) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCountry

`func (o *GeoLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeoLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeoLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GeoLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *GeoLocation) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *GeoLocation) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetState

`func (o *GeoLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GeoLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GeoLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GeoLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *GeoLocation) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *GeoLocation) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCity

`func (o *GeoLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GeoLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GeoLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GeoLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *GeoLocation) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GeoLocation) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDistrict

`func (o *GeoLocation) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *GeoLocation) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *GeoLocation) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *GeoLocation) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *GeoLocation) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *GeoLocation) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetStreet

`func (o *GeoLocation) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GeoLocation) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GeoLocation) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GeoLocation) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *GeoLocation) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *GeoLocation) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetZipCode

`func (o *GeoLocation) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GeoLocation) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GeoLocation) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *GeoLocation) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *GeoLocation) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *GeoLocation) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetAddress

`func (o *GeoLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GeoLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GeoLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GeoLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *GeoLocation) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *GeoLocation) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetMapType

`func (o *GeoLocation) GetMapType() string`

GetMapType returns the MapType field if non-nil, zero value otherwise.

### GetMapTypeOk

`func (o *GeoLocation) GetMapTypeOk() (*string, bool)`

GetMapTypeOk returns a tuple with the MapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapType

`func (o *GeoLocation) SetMapType(v string)`

SetMapType sets MapType field to given value.

### HasMapType

`func (o *GeoLocation) HasMapType() bool`

HasMapType returns a boolean if a field has been set.

### SetMapTypeNil

`func (o *GeoLocation) SetMapTypeNil(b bool)`

 SetMapTypeNil sets the value for MapType to be an explicit nil

### UnsetMapType
`func (o *GeoLocation) UnsetMapType()`

UnsetMapType ensures that no value is present for MapType, not even an explicit nil
### GetRemark

`func (o *GeoLocation) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *GeoLocation) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *GeoLocation) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *GeoLocation) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *GeoLocation) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *GeoLocation) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetTags

`func (o *GeoLocation) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GeoLocation) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GeoLocation) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GeoLocation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GeoLocation) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GeoLocation) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetEnable

`func (o *GeoLocation) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GeoLocation) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GeoLocation) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GeoLocation) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetShowIndex

`func (o *GeoLocation) GetShowIndex() int32`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *GeoLocation) GetShowIndexOk() (*int32, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *GeoLocation) SetShowIndex(v int32)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *GeoLocation) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.

### GetCreateDate

`func (o *GeoLocation) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *GeoLocation) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *GeoLocation) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *GeoLocation) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *GeoLocation) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *GeoLocation) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *GeoLocation) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *GeoLocation) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


