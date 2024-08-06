/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the AlipayCommerceEcEnterpriseAddressAddModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcEnterpriseAddressAddModel{}

// AlipayCommerceEcEnterpriseAddressAddModel struct for AlipayCommerceEcEnterpriseAddressAddModel
type AlipayCommerceEcEnterpriseAddressAddModel struct {
	// 通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。
	AccountId *string `json:"account_id,omitempty"`
	// 详细地址最长50个字符
	Address *string `json:"address,omitempty"`
	// 可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 市(国家统一行政规划编码)
	CityCode *string `json:"city_code,omitempty"`
	// 市
	CityName *string `json:"city_name,omitempty"`
	// 小区/楼宇
	Community *string `json:"community,omitempty"`
	// 通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 纬度
	Latitude *string `json:"latitude,omitempty"`
	// 经度
	Longitude *string `json:"longitude,omitempty"`
	// 备注最长50个字符
	Mark *string `json:"mark,omitempty"`
	// 高德地图poi
	PoiId *string `json:"poi_id,omitempty"`
}

// NewAlipayCommerceEcEnterpriseAddressAddModel instantiates a new AlipayCommerceEcEnterpriseAddressAddModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcEnterpriseAddressAddModel() *AlipayCommerceEcEnterpriseAddressAddModel {
	this := AlipayCommerceEcEnterpriseAddressAddModel{}
	return &this
}

// NewAlipayCommerceEcEnterpriseAddressAddModelWithDefaults instantiates a new AlipayCommerceEcEnterpriseAddressAddModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcEnterpriseAddressAddModelWithDefaults() *AlipayCommerceEcEnterpriseAddressAddModel {
	this := AlipayCommerceEcEnterpriseAddressAddModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetAddress(v string) {
	o.Address = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetCityCode returns the CityCode field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetCityCode() string {
	if o == nil || IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

// GetCityCodeOk returns a tuple with the CityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetCityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

// HasCityCode returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

// SetCityCode gets a reference to the given string and assigns it to the CityCode field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetCityCode(v string) {
	o.CityCode = &v
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetCityName(v string) {
	o.CityName = &v
}

// GetCommunity returns the Community field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetCommunity() string {
	if o == nil || IsNil(o.Community) {
		var ret string
		return ret
	}
	return *o.Community
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetCommunityOk() (*string, bool) {
	if o == nil || IsNil(o.Community) {
		return nil, false
	}
	return o.Community, true
}

// HasCommunity returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasCommunity() bool {
	if o != nil && !IsNil(o.Community) {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given string and assigns it to the Community field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetCommunity(v string) {
	o.Community = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetLongitude(v string) {
	o.Longitude = &v
}

// GetMark returns the Mark field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetMark() string {
	if o == nil || IsNil(o.Mark) {
		var ret string
		return ret
	}
	return *o.Mark
}

// GetMarkOk returns a tuple with the Mark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetMarkOk() (*string, bool) {
	if o == nil || IsNil(o.Mark) {
		return nil, false
	}
	return o.Mark, true
}

// HasMark returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasMark() bool {
	if o != nil && !IsNil(o.Mark) {
		return true
	}

	return false
}

// SetMark gets a reference to the given string and assigns it to the Mark field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetMark(v string) {
	o.Mark = &v
}

// GetPoiId returns the PoiId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetPoiId() string {
	if o == nil || IsNil(o.PoiId) {
		var ret string
		return ret
	}
	return *o.PoiId
}

// GetPoiIdOk returns a tuple with the PoiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) GetPoiIdOk() (*string, bool) {
	if o == nil || IsNil(o.PoiId) {
		return nil, false
	}
	return o.PoiId, true
}

// HasPoiId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) HasPoiId() bool {
	if o != nil && !IsNil(o.PoiId) {
		return true
	}

	return false
}

// SetPoiId gets a reference to the given string and assigns it to the PoiId field.
func (o *AlipayCommerceEcEnterpriseAddressAddModel) SetPoiId(v string) {
	o.PoiId = &v
}

func (o AlipayCommerceEcEnterpriseAddressAddModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcEnterpriseAddressAddModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.CityCode) {
		toSerialize["city_code"] = o.CityCode
	}
	if !IsNil(o.CityName) {
		toSerialize["city_name"] = o.CityName
	}
	if !IsNil(o.Community) {
		toSerialize["community"] = o.Community
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Mark) {
		toSerialize["mark"] = o.Mark
	}
	if !IsNil(o.PoiId) {
		toSerialize["poi_id"] = o.PoiId
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcEnterpriseAddressAddModel struct {
	value *AlipayCommerceEcEnterpriseAddressAddModel
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseAddressAddModel) Get() *AlipayCommerceEcEnterpriseAddressAddModel {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseAddressAddModel) Set(val *AlipayCommerceEcEnterpriseAddressAddModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseAddressAddModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseAddressAddModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseAddressAddModel(val *AlipayCommerceEcEnterpriseAddressAddModel) *NullableAlipayCommerceEcEnterpriseAddressAddModel {
	return &NullableAlipayCommerceEcEnterpriseAddressAddModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseAddressAddModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseAddressAddModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
