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

// checks if the EnterpriseAgreementDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterpriseAgreementDTO{}

// EnterpriseAgreementDTO struct for EnterpriseAgreementDTO
type EnterpriseAgreementDTO struct {
	// 协议id
	AgreementId *string `json:"agreement_id,omitempty"`
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// 签约时间
	SignDate *string `json:"sign_date,omitempty"`
	// 签约状态
	SignStatus *string `json:"sign_status,omitempty"`
}

// NewEnterpriseAgreementDTO instantiates a new EnterpriseAgreementDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAgreementDTO() *EnterpriseAgreementDTO {
	this := EnterpriseAgreementDTO{}
	return &this
}

// NewEnterpriseAgreementDTOWithDefaults instantiates a new EnterpriseAgreementDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAgreementDTOWithDefaults() *EnterpriseAgreementDTO {
	this := EnterpriseAgreementDTO{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *EnterpriseAgreementDTO) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgreementDTO) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *EnterpriseAgreementDTO) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *EnterpriseAgreementDTO) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *EnterpriseAgreementDTO) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgreementDTO) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *EnterpriseAgreementDTO) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *EnterpriseAgreementDTO) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetEnterpriseName returns the EnterpriseName field value if set, zero value otherwise.
func (o *EnterpriseAgreementDTO) GetEnterpriseName() string {
	if o == nil || IsNil(o.EnterpriseName) {
		var ret string
		return ret
	}
	return *o.EnterpriseName
}

// GetEnterpriseNameOk returns a tuple with the EnterpriseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgreementDTO) GetEnterpriseNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseName) {
		return nil, false
	}
	return o.EnterpriseName, true
}

// HasEnterpriseName returns a boolean if a field has been set.
func (o *EnterpriseAgreementDTO) HasEnterpriseName() bool {
	if o != nil && !IsNil(o.EnterpriseName) {
		return true
	}

	return false
}

// SetEnterpriseName gets a reference to the given string and assigns it to the EnterpriseName field.
func (o *EnterpriseAgreementDTO) SetEnterpriseName(v string) {
	o.EnterpriseName = &v
}

// GetSignDate returns the SignDate field value if set, zero value otherwise.
func (o *EnterpriseAgreementDTO) GetSignDate() string {
	if o == nil || IsNil(o.SignDate) {
		var ret string
		return ret
	}
	return *o.SignDate
}

// GetSignDateOk returns a tuple with the SignDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgreementDTO) GetSignDateOk() (*string, bool) {
	if o == nil || IsNil(o.SignDate) {
		return nil, false
	}
	return o.SignDate, true
}

// HasSignDate returns a boolean if a field has been set.
func (o *EnterpriseAgreementDTO) HasSignDate() bool {
	if o != nil && !IsNil(o.SignDate) {
		return true
	}

	return false
}

// SetSignDate gets a reference to the given string and assigns it to the SignDate field.
func (o *EnterpriseAgreementDTO) SetSignDate(v string) {
	o.SignDate = &v
}

// GetSignStatus returns the SignStatus field value if set, zero value otherwise.
func (o *EnterpriseAgreementDTO) GetSignStatus() string {
	if o == nil || IsNil(o.SignStatus) {
		var ret string
		return ret
	}
	return *o.SignStatus
}

// GetSignStatusOk returns a tuple with the SignStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgreementDTO) GetSignStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SignStatus) {
		return nil, false
	}
	return o.SignStatus, true
}

// HasSignStatus returns a boolean if a field has been set.
func (o *EnterpriseAgreementDTO) HasSignStatus() bool {
	if o != nil && !IsNil(o.SignStatus) {
		return true
	}

	return false
}

// SetSignStatus gets a reference to the given string and assigns it to the SignStatus field.
func (o *EnterpriseAgreementDTO) SetSignStatus(v string) {
	o.SignStatus = &v
}

func (o EnterpriseAgreementDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseAgreementDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.EnterpriseName) {
		toSerialize["enterprise_name"] = o.EnterpriseName
	}
	if !IsNil(o.SignDate) {
		toSerialize["sign_date"] = o.SignDate
	}
	if !IsNil(o.SignStatus) {
		toSerialize["sign_status"] = o.SignStatus
	}
	return toSerialize, nil
}

type NullableEnterpriseAgreementDTO struct {
	value *EnterpriseAgreementDTO
	isSet bool
}

func (v NullableEnterpriseAgreementDTO) Get() *EnterpriseAgreementDTO {
	return v.value
}

func (v *NullableEnterpriseAgreementDTO) Set(val *EnterpriseAgreementDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAgreementDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAgreementDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAgreementDTO(val *EnterpriseAgreementDTO) *NullableEnterpriseAgreementDTO {
	return &NullableEnterpriseAgreementDTO{value: val, isSet: true}
}

func (v NullableEnterpriseAgreementDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAgreementDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
