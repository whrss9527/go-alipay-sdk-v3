/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AlipayOpenInstantdeliveryAccountQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenInstantdeliveryAccountQueryResponseModel{}

// AlipayOpenInstantdeliveryAccountQueryResponseModel struct for AlipayOpenInstantdeliveryAccountQueryResponseModel
type AlipayOpenInstantdeliveryAccountQueryResponseModel struct {
	// 账户余额
	Balance *string `json:"balance,omitempty"`
	// 企业营业执照图片正反面
	BusinessLicense *string `json:"business_license,omitempty"`
	// 经营范围
	BusinessScope *string `json:"business_scope,omitempty"`
	// 统一社会信用代码
	CreditCode *string `json:"credit_code,omitempty"`
	// 联系人邮箱
	Email *string `json:"email,omitempty"`
	// 企业地址
	EnterpriseAddress *string `json:"enterprise_address,omitempty"`
	// 企业所在地，市编码
	EnterpriseCity *string `json:"enterprise_city,omitempty"`
	// 企业所在地，区编码
	EnterpriseDistrict *string `json:"enterprise_district,omitempty"`
	// 企业全称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// 企业所在地，省编码
	EnterpriseProvince *string `json:"enterprise_province,omitempty"`
	// 企业类型
	EnterpriseType *string `json:"enterprise_type,omitempty"`
	// 在配送公司的账户状态信息
	LogisticsAccountStatus []LogisticsAccountStatusDTO `json:"logistics_account_status,omitempty"`
	// 联系人手机号
	Phone *string `json:"phone,omitempty"`
}

// NewAlipayOpenInstantdeliveryAccountQueryResponseModel instantiates a new AlipayOpenInstantdeliveryAccountQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenInstantdeliveryAccountQueryResponseModel() *AlipayOpenInstantdeliveryAccountQueryResponseModel {
	this := AlipayOpenInstantdeliveryAccountQueryResponseModel{}
	return &this
}

// NewAlipayOpenInstantdeliveryAccountQueryResponseModelWithDefaults instantiates a new AlipayOpenInstantdeliveryAccountQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenInstantdeliveryAccountQueryResponseModelWithDefaults() *AlipayOpenInstantdeliveryAccountQueryResponseModel {
	this := AlipayOpenInstantdeliveryAccountQueryResponseModel{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetBalance(v string) {
	o.Balance = &v
}

// GetBusinessLicense returns the BusinessLicense field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetBusinessLicense() string {
	if o == nil || IsNil(o.BusinessLicense) {
		var ret string
		return ret
	}
	return *o.BusinessLicense
}

// GetBusinessLicenseOk returns a tuple with the BusinessLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetBusinessLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLicense) {
		return nil, false
	}
	return o.BusinessLicense, true
}

// HasBusinessLicense returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasBusinessLicense() bool {
	if o != nil && !IsNil(o.BusinessLicense) {
		return true
	}

	return false
}

// SetBusinessLicense gets a reference to the given string and assigns it to the BusinessLicense field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetBusinessLicense(v string) {
	o.BusinessLicense = &v
}

// GetBusinessScope returns the BusinessScope field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetBusinessScope() string {
	if o == nil || IsNil(o.BusinessScope) {
		var ret string
		return ret
	}
	return *o.BusinessScope
}

// GetBusinessScopeOk returns a tuple with the BusinessScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetBusinessScopeOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessScope) {
		return nil, false
	}
	return o.BusinessScope, true
}

// HasBusinessScope returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasBusinessScope() bool {
	if o != nil && !IsNil(o.BusinessScope) {
		return true
	}

	return false
}

// SetBusinessScope gets a reference to the given string and assigns it to the BusinessScope field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetBusinessScope(v string) {
	o.BusinessScope = &v
}

// GetCreditCode returns the CreditCode field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetCreditCode() string {
	if o == nil || IsNil(o.CreditCode) {
		var ret string
		return ret
	}
	return *o.CreditCode
}

// GetCreditCodeOk returns a tuple with the CreditCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetCreditCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CreditCode) {
		return nil, false
	}
	return o.CreditCode, true
}

// HasCreditCode returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasCreditCode() bool {
	if o != nil && !IsNil(o.CreditCode) {
		return true
	}

	return false
}

// SetCreditCode gets a reference to the given string and assigns it to the CreditCode field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetCreditCode(v string) {
	o.CreditCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetEmail(v string) {
	o.Email = &v
}

// GetEnterpriseAddress returns the EnterpriseAddress field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseAddress() string {
	if o == nil || IsNil(o.EnterpriseAddress) {
		var ret string
		return ret
	}
	return *o.EnterpriseAddress
}

// GetEnterpriseAddressOk returns a tuple with the EnterpriseAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseAddress) {
		return nil, false
	}
	return o.EnterpriseAddress, true
}

// HasEnterpriseAddress returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasEnterpriseAddress() bool {
	if o != nil && !IsNil(o.EnterpriseAddress) {
		return true
	}

	return false
}

// SetEnterpriseAddress gets a reference to the given string and assigns it to the EnterpriseAddress field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetEnterpriseAddress(v string) {
	o.EnterpriseAddress = &v
}

// GetEnterpriseCity returns the EnterpriseCity field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseCity() string {
	if o == nil || IsNil(o.EnterpriseCity) {
		var ret string
		return ret
	}
	return *o.EnterpriseCity
}

// GetEnterpriseCityOk returns a tuple with the EnterpriseCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseCityOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseCity) {
		return nil, false
	}
	return o.EnterpriseCity, true
}

// HasEnterpriseCity returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasEnterpriseCity() bool {
	if o != nil && !IsNil(o.EnterpriseCity) {
		return true
	}

	return false
}

// SetEnterpriseCity gets a reference to the given string and assigns it to the EnterpriseCity field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetEnterpriseCity(v string) {
	o.EnterpriseCity = &v
}

// GetEnterpriseDistrict returns the EnterpriseDistrict field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseDistrict() string {
	if o == nil || IsNil(o.EnterpriseDistrict) {
		var ret string
		return ret
	}
	return *o.EnterpriseDistrict
}

// GetEnterpriseDistrictOk returns a tuple with the EnterpriseDistrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseDistrictOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseDistrict) {
		return nil, false
	}
	return o.EnterpriseDistrict, true
}

// HasEnterpriseDistrict returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasEnterpriseDistrict() bool {
	if o != nil && !IsNil(o.EnterpriseDistrict) {
		return true
	}

	return false
}

// SetEnterpriseDistrict gets a reference to the given string and assigns it to the EnterpriseDistrict field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetEnterpriseDistrict(v string) {
	o.EnterpriseDistrict = &v
}

// GetEnterpriseName returns the EnterpriseName field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseName() string {
	if o == nil || IsNil(o.EnterpriseName) {
		var ret string
		return ret
	}
	return *o.EnterpriseName
}

// GetEnterpriseNameOk returns a tuple with the EnterpriseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseName) {
		return nil, false
	}
	return o.EnterpriseName, true
}

// HasEnterpriseName returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasEnterpriseName() bool {
	if o != nil && !IsNil(o.EnterpriseName) {
		return true
	}

	return false
}

// SetEnterpriseName gets a reference to the given string and assigns it to the EnterpriseName field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetEnterpriseName(v string) {
	o.EnterpriseName = &v
}

// GetEnterpriseProvince returns the EnterpriseProvince field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseProvince() string {
	if o == nil || IsNil(o.EnterpriseProvince) {
		var ret string
		return ret
	}
	return *o.EnterpriseProvince
}

// GetEnterpriseProvinceOk returns a tuple with the EnterpriseProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseProvince) {
		return nil, false
	}
	return o.EnterpriseProvince, true
}

// HasEnterpriseProvince returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasEnterpriseProvince() bool {
	if o != nil && !IsNil(o.EnterpriseProvince) {
		return true
	}

	return false
}

// SetEnterpriseProvince gets a reference to the given string and assigns it to the EnterpriseProvince field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetEnterpriseProvince(v string) {
	o.EnterpriseProvince = &v
}

// GetEnterpriseType returns the EnterpriseType field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseType() string {
	if o == nil || IsNil(o.EnterpriseType) {
		var ret string
		return ret
	}
	return *o.EnterpriseType
}

// GetEnterpriseTypeOk returns a tuple with the EnterpriseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetEnterpriseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseType) {
		return nil, false
	}
	return o.EnterpriseType, true
}

// HasEnterpriseType returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasEnterpriseType() bool {
	if o != nil && !IsNil(o.EnterpriseType) {
		return true
	}

	return false
}

// SetEnterpriseType gets a reference to the given string and assigns it to the EnterpriseType field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetEnterpriseType(v string) {
	o.EnterpriseType = &v
}

// GetLogisticsAccountStatus returns the LogisticsAccountStatus field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetLogisticsAccountStatus() []LogisticsAccountStatusDTO {
	if o == nil || IsNil(o.LogisticsAccountStatus) {
		var ret []LogisticsAccountStatusDTO
		return ret
	}
	return o.LogisticsAccountStatus
}

// GetLogisticsAccountStatusOk returns a tuple with the LogisticsAccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetLogisticsAccountStatusOk() ([]LogisticsAccountStatusDTO, bool) {
	if o == nil || IsNil(o.LogisticsAccountStatus) {
		return nil, false
	}
	return o.LogisticsAccountStatus, true
}

// HasLogisticsAccountStatus returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasLogisticsAccountStatus() bool {
	if o != nil && !IsNil(o.LogisticsAccountStatus) {
		return true
	}

	return false
}

// SetLogisticsAccountStatus gets a reference to the given []LogisticsAccountStatusDTO and assigns it to the LogisticsAccountStatus field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetLogisticsAccountStatus(v []LogisticsAccountStatusDTO) {
	o.LogisticsAccountStatus = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AlipayOpenInstantdeliveryAccountQueryResponseModel) SetPhone(v string) {
	o.Phone = &v
}

func (o AlipayOpenInstantdeliveryAccountQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenInstantdeliveryAccountQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.BusinessLicense) {
		toSerialize["business_license"] = o.BusinessLicense
	}
	if !IsNil(o.BusinessScope) {
		toSerialize["business_scope"] = o.BusinessScope
	}
	if !IsNil(o.CreditCode) {
		toSerialize["credit_code"] = o.CreditCode
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EnterpriseAddress) {
		toSerialize["enterprise_address"] = o.EnterpriseAddress
	}
	if !IsNil(o.EnterpriseCity) {
		toSerialize["enterprise_city"] = o.EnterpriseCity
	}
	if !IsNil(o.EnterpriseDistrict) {
		toSerialize["enterprise_district"] = o.EnterpriseDistrict
	}
	if !IsNil(o.EnterpriseName) {
		toSerialize["enterprise_name"] = o.EnterpriseName
	}
	if !IsNil(o.EnterpriseProvince) {
		toSerialize["enterprise_province"] = o.EnterpriseProvince
	}
	if !IsNil(o.EnterpriseType) {
		toSerialize["enterprise_type"] = o.EnterpriseType
	}
	if !IsNil(o.LogisticsAccountStatus) {
		toSerialize["logistics_account_status"] = o.LogisticsAccountStatus
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}

type NullableAlipayOpenInstantdeliveryAccountQueryResponseModel struct {
	value *AlipayOpenInstantdeliveryAccountQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryAccountQueryResponseModel) Get() *AlipayOpenInstantdeliveryAccountQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryAccountQueryResponseModel) Set(val *AlipayOpenInstantdeliveryAccountQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryAccountQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryAccountQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryAccountQueryResponseModel(val *AlipayOpenInstantdeliveryAccountQueryResponseModel) *NullableAlipayOpenInstantdeliveryAccountQueryResponseModel {
	return &NullableAlipayOpenInstantdeliveryAccountQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryAccountQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryAccountQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


