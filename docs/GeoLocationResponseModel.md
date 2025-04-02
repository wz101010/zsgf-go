# GeoLocationResponseModel

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
**Id** | Pointer to **int32** | ID | [optional] 

## Methods

### NewGeoLocationResponseModel

`func NewGeoLocationResponseModel() *GeoLocationResponseModel`

NewGeoLocationResponseModel instantiates a new GeoLocationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoLocationResponseModelWithDefaults

`func NewGeoLocationResponseModelWithDefaults() *GeoLocationResponseModel`

NewGeoLocationResponseModelWithDefaults instantiates a new GeoLocationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *GeoLocationResponseModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GeoLocationResponseModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GeoLocationResponseModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GeoLocationResponseModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *GeoLocationResponseModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GeoLocationResponseModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GeoLocationResponseModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GeoLocationResponseModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLocationName

`func (o *GeoLocationResponseModel) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *GeoLocationResponseModel) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *GeoLocationResponseModel) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *GeoLocationResponseModel) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *GeoLocationResponseModel) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *GeoLocationResponseModel) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetLocationType

`func (o *GeoLocationResponseModel) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *GeoLocationResponseModel) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *GeoLocationResponseModel) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *GeoLocationResponseModel) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *GeoLocationResponseModel) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *GeoLocationResponseModel) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetRecipientName

`func (o *GeoLocationResponseModel) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *GeoLocationResponseModel) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *GeoLocationResponseModel) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *GeoLocationResponseModel) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *GeoLocationResponseModel) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *GeoLocationResponseModel) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetPhoneNumber

`func (o *GeoLocationResponseModel) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GeoLocationResponseModel) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GeoLocationResponseModel) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GeoLocationResponseModel) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *GeoLocationResponseModel) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *GeoLocationResponseModel) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *GeoLocationResponseModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GeoLocationResponseModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GeoLocationResponseModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GeoLocationResponseModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GeoLocationResponseModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GeoLocationResponseModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCountry

`func (o *GeoLocationResponseModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeoLocationResponseModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeoLocationResponseModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GeoLocationResponseModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *GeoLocationResponseModel) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *GeoLocationResponseModel) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetState

`func (o *GeoLocationResponseModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GeoLocationResponseModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GeoLocationResponseModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GeoLocationResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *GeoLocationResponseModel) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *GeoLocationResponseModel) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCity

`func (o *GeoLocationResponseModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GeoLocationResponseModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GeoLocationResponseModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GeoLocationResponseModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *GeoLocationResponseModel) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GeoLocationResponseModel) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDistrict

`func (o *GeoLocationResponseModel) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *GeoLocationResponseModel) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *GeoLocationResponseModel) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *GeoLocationResponseModel) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *GeoLocationResponseModel) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *GeoLocationResponseModel) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil
### GetStreet

`func (o *GeoLocationResponseModel) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GeoLocationResponseModel) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GeoLocationResponseModel) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GeoLocationResponseModel) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *GeoLocationResponseModel) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *GeoLocationResponseModel) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetZipCode

`func (o *GeoLocationResponseModel) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GeoLocationResponseModel) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GeoLocationResponseModel) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *GeoLocationResponseModel) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *GeoLocationResponseModel) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *GeoLocationResponseModel) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetAddress

`func (o *GeoLocationResponseModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GeoLocationResponseModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GeoLocationResponseModel) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GeoLocationResponseModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *GeoLocationResponseModel) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *GeoLocationResponseModel) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetMapType

`func (o *GeoLocationResponseModel) GetMapType() string`

GetMapType returns the MapType field if non-nil, zero value otherwise.

### GetMapTypeOk

`func (o *GeoLocationResponseModel) GetMapTypeOk() (*string, bool)`

GetMapTypeOk returns a tuple with the MapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapType

`func (o *GeoLocationResponseModel) SetMapType(v string)`

SetMapType sets MapType field to given value.

### HasMapType

`func (o *GeoLocationResponseModel) HasMapType() bool`

HasMapType returns a boolean if a field has been set.

### SetMapTypeNil

`func (o *GeoLocationResponseModel) SetMapTypeNil(b bool)`

 SetMapTypeNil sets the value for MapType to be an explicit nil

### UnsetMapType
`func (o *GeoLocationResponseModel) UnsetMapType()`

UnsetMapType ensures that no value is present for MapType, not even an explicit nil
### GetRemark

`func (o *GeoLocationResponseModel) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *GeoLocationResponseModel) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *GeoLocationResponseModel) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *GeoLocationResponseModel) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *GeoLocationResponseModel) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *GeoLocationResponseModel) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetTags

`func (o *GeoLocationResponseModel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GeoLocationResponseModel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GeoLocationResponseModel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GeoLocationResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GeoLocationResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GeoLocationResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetEnable

`func (o *GeoLocationResponseModel) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GeoLocationResponseModel) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GeoLocationResponseModel) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GeoLocationResponseModel) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetShowIndex

`func (o *GeoLocationResponseModel) GetShowIndex() int32`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *GeoLocationResponseModel) GetShowIndexOk() (*int32, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *GeoLocationResponseModel) SetShowIndex(v int32)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *GeoLocationResponseModel) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.

### GetCreateDate

`func (o *GeoLocationResponseModel) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *GeoLocationResponseModel) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *GeoLocationResponseModel) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *GeoLocationResponseModel) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *GeoLocationResponseModel) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *GeoLocationResponseModel) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *GeoLocationResponseModel) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *GeoLocationResponseModel) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetId

`func (o *GeoLocationResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeoLocationResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeoLocationResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GeoLocationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


