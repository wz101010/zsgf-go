# GeoLocationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | Pointer to **float64** | 纬度 | [optional] 
**Longitude** | Pointer to **float64** | 经度 | [optional] 
**LocationName** | Pointer to **NullableString** | 地点的名称 | [optional] 
**LocationType** | Pointer to **NullableString** | 地点的类型，如家庭、工作、学校等 | [optional] 
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
**MapType** | Pointer to **NullableString** | 地址类型，百度、高德、谷歌 | [optional] 
**Remark** | Pointer to **NullableString** | 备注 | [optional] 
**Tags** | Pointer to **NullableString** | 标签 | [optional] 
**Enable** | Pointer to **bool** | 启用 | [optional] 
**ShowIndex** | Pointer to **int32** | 排序 | [optional] 
**CreateDate** | Pointer to **time.Time** | 创建时间 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 最后更新的时间 | [optional] 

## Methods

### NewGeoLocationModel

`func NewGeoLocationModel() *GeoLocationModel`

NewGeoLocationModel instantiates a new GeoLocationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoLocationModelWithDefaults

`func NewGeoLocationModelWithDefaults() *GeoLocationModel`

NewGeoLocationModelWithDefaults instantiates a new GeoLocationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *GeoLocationModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GeoLocationModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GeoLocationModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GeoLocationModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *GeoLocationModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GeoLocationModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GeoLocationModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GeoLocationModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLocationName

`func (o *GeoLocationModel) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *GeoLocationModel) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *GeoLocationModel) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *GeoLocationModel) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *GeoLocationModel) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *GeoLocationModel) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetLocationType

`func (o *GeoLocationModel) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *GeoLocationModel) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *GeoLocationModel) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *GeoLocationModel) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *GeoLocationModel) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *GeoLocationModel) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetRecipientName

`func (o *GeoLocationModel) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *GeoLocationModel) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *GeoLocationModel) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *GeoLocationModel) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *GeoLocationModel) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *GeoLocationModel) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetPhoneNumber

`func (o *GeoLocationModel) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GeoLocationModel) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GeoLocationModel) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GeoLocationModel) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *GeoLocationModel) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *GeoLocationModel) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *GeoLocationModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GeoLocationModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GeoLocationModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GeoLocationModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GeoLocationModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GeoLocationModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCountry

`func (o *GeoLocationModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeoLocationModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeoLocationModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GeoLocationModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *GeoLocationModel) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *GeoLocationModel) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetState

`func (o *GeoLocationModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GeoLocationModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GeoLocationModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GeoLocationModel) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *GeoLocationModel) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *GeoLocationModel) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCity

`func (o *GeoLocationModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GeoLocationModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GeoLocationModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GeoLocationModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *GeoLocationModel) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GeoLocationModel) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDistrict

`func (o *GeoLocationModel) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *GeoLocationModel) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *GeoLocationModel) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *GeoLocationModel) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *GeoLocationModel) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *GeoLocationModel) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetStreet

`func (o *GeoLocationModel) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GeoLocationModel) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GeoLocationModel) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GeoLocationModel) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *GeoLocationModel) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *GeoLocationModel) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetZipCode

`func (o *GeoLocationModel) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GeoLocationModel) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GeoLocationModel) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *GeoLocationModel) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *GeoLocationModel) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *GeoLocationModel) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetAddress

`func (o *GeoLocationModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GeoLocationModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GeoLocationModel) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GeoLocationModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *GeoLocationModel) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *GeoLocationModel) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetMapType

`func (o *GeoLocationModel) GetMapType() string`

GetMapType returns the MapType field if non-nil, zero value otherwise.

### GetMapTypeOk

`func (o *GeoLocationModel) GetMapTypeOk() (*string, bool)`

GetMapTypeOk returns a tuple with the MapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapType

`func (o *GeoLocationModel) SetMapType(v string)`

SetMapType sets MapType field to given value.

### HasMapType

`func (o *GeoLocationModel) HasMapType() bool`

HasMapType returns a boolean if a field has been set.

### SetMapTypeNil

`func (o *GeoLocationModel) SetMapTypeNil(b bool)`

 SetMapTypeNil sets the value for MapType to be an explicit nil

### UnsetMapType
`func (o *GeoLocationModel) UnsetMapType()`

UnsetMapType ensures that no value is present for MapType, not even an explicit nil
### GetRemark

`func (o *GeoLocationModel) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *GeoLocationModel) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *GeoLocationModel) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *GeoLocationModel) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *GeoLocationModel) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *GeoLocationModel) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetTags

`func (o *GeoLocationModel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GeoLocationModel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GeoLocationModel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GeoLocationModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GeoLocationModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GeoLocationModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetEnable

`func (o *GeoLocationModel) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GeoLocationModel) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GeoLocationModel) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GeoLocationModel) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetShowIndex

`func (o *GeoLocationModel) GetShowIndex() int32`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *GeoLocationModel) GetShowIndexOk() (*int32, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *GeoLocationModel) SetShowIndex(v int32)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *GeoLocationModel) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.

### GetCreateDate

`func (o *GeoLocationModel) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *GeoLocationModel) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *GeoLocationModel) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *GeoLocationModel) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *GeoLocationModel) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *GeoLocationModel) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *GeoLocationModel) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *GeoLocationModel) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


