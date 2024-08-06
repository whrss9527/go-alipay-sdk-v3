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

// checks if the UserRiskPrediction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskPrediction{}

// UserRiskPrediction struct for UserRiskPrediction
type UserRiskPrediction struct {
	// 用户绑定手机号被二次放号风险等级。 NO_RESULT：手机号风险未入库 NO_RISK：用户绑定手机无二次放号风险，高置信度 LOW_RISK：用户绑定手机二次放号风险较小，商户自行决策是否相信。 HIGH_RISK：用户绑定手机已被放号，高置信度
	PhoneRecycleRiskLeve *string `json:"phone_recycle_risk_leve,omitempty"`
	// 用户拒付风险等级。 NO_SIGN：商户未签约。 NO_RESULT：未查询到账户信息。 LOW_RISK：用户拒付风险为低；处理建议：用户可以先享受服务，再进行支付。 MEDIUM_RISK：用户拒付风险为中；处理建议：根据业务场景客户自行判断提供或者不提供。 HIGH_RISK：用户拒付风险为高；处理建议：不建议先提供给用户服务。
	RefusedPaymentRiskLevel *string `json:"refused_payment_risk_level,omitempty"`
}

// NewUserRiskPrediction instantiates a new UserRiskPrediction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskPrediction() *UserRiskPrediction {
	this := UserRiskPrediction{}
	return &this
}

// NewUserRiskPredictionWithDefaults instantiates a new UserRiskPrediction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskPredictionWithDefaults() *UserRiskPrediction {
	this := UserRiskPrediction{}
	return &this
}

// GetPhoneRecycleRiskLeve returns the PhoneRecycleRiskLeve field value if set, zero value otherwise.
func (o *UserRiskPrediction) GetPhoneRecycleRiskLeve() string {
	if o == nil || IsNil(o.PhoneRecycleRiskLeve) {
		var ret string
		return ret
	}
	return *o.PhoneRecycleRiskLeve
}

// GetPhoneRecycleRiskLeveOk returns a tuple with the PhoneRecycleRiskLeve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskPrediction) GetPhoneRecycleRiskLeveOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneRecycleRiskLeve) {
		return nil, false
	}
	return o.PhoneRecycleRiskLeve, true
}

// HasPhoneRecycleRiskLeve returns a boolean if a field has been set.
func (o *UserRiskPrediction) HasPhoneRecycleRiskLeve() bool {
	if o != nil && !IsNil(o.PhoneRecycleRiskLeve) {
		return true
	}

	return false
}

// SetPhoneRecycleRiskLeve gets a reference to the given string and assigns it to the PhoneRecycleRiskLeve field.
func (o *UserRiskPrediction) SetPhoneRecycleRiskLeve(v string) {
	o.PhoneRecycleRiskLeve = &v
}

// GetRefusedPaymentRiskLevel returns the RefusedPaymentRiskLevel field value if set, zero value otherwise.
func (o *UserRiskPrediction) GetRefusedPaymentRiskLevel() string {
	if o == nil || IsNil(o.RefusedPaymentRiskLevel) {
		var ret string
		return ret
	}
	return *o.RefusedPaymentRiskLevel
}

// GetRefusedPaymentRiskLevelOk returns a tuple with the RefusedPaymentRiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskPrediction) GetRefusedPaymentRiskLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RefusedPaymentRiskLevel) {
		return nil, false
	}
	return o.RefusedPaymentRiskLevel, true
}

// HasRefusedPaymentRiskLevel returns a boolean if a field has been set.
func (o *UserRiskPrediction) HasRefusedPaymentRiskLevel() bool {
	if o != nil && !IsNil(o.RefusedPaymentRiskLevel) {
		return true
	}

	return false
}

// SetRefusedPaymentRiskLevel gets a reference to the given string and assigns it to the RefusedPaymentRiskLevel field.
func (o *UserRiskPrediction) SetRefusedPaymentRiskLevel(v string) {
	o.RefusedPaymentRiskLevel = &v
}

func (o UserRiskPrediction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskPrediction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhoneRecycleRiskLeve) {
		toSerialize["phone_recycle_risk_leve"] = o.PhoneRecycleRiskLeve
	}
	if !IsNil(o.RefusedPaymentRiskLevel) {
		toSerialize["refused_payment_risk_level"] = o.RefusedPaymentRiskLevel
	}
	return toSerialize, nil
}

type NullableUserRiskPrediction struct {
	value *UserRiskPrediction
	isSet bool
}

func (v NullableUserRiskPrediction) Get() *UserRiskPrediction {
	return v.value
}

func (v *NullableUserRiskPrediction) Set(val *UserRiskPrediction) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskPrediction) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskPrediction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskPrediction(val *UserRiskPrediction) *NullableUserRiskPrediction {
	return &NullableUserRiskPrediction{value: val, isSet: true}
}

func (v NullableUserRiskPrediction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskPrediction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


