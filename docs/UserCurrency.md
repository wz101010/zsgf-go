# UserCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 用户资产的唯一标识符。 | [optional] 
**UserID** | Pointer to **int64** | 与用户资产关联的用户ID。 | [optional] 
**CurrencyCode** | Pointer to **NullableString** | 用户资产的货币代码，例如 &#39;USD&#39;, &#39;CNY&#39; 等。 | [optional] 
**Balance** | Pointer to **int64** | 用户的账户余额，表示当前持有的货币数量。 | [optional] 

## Methods

### NewUserCurrency

`func NewUserCurrency() *UserCurrency`

NewUserCurrency instantiates a new UserCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCurrencyWithDefaults

`func NewUserCurrencyWithDefaults() *UserCurrency`

NewUserCurrencyWithDefaults instantiates a new UserCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserCurrency) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCurrency) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCurrency) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserCurrency) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *UserCurrency) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *UserCurrency) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *UserCurrency) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *UserCurrency) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *UserCurrency) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *UserCurrency) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *UserCurrency) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *UserCurrency) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *UserCurrency) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *UserCurrency) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetBalance

`func (o *UserCurrency) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *UserCurrency) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *UserCurrency) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *UserCurrency) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


