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

// checks if the AlipayEbppPdeductSignAddResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppPdeductSignAddResponseModel{}

// AlipayEbppPdeductSignAddResponseModel struct for AlipayEbppPdeductSignAddResponseModel
type AlipayEbppPdeductSignAddResponseModel struct {
	// 支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// 支付宝协议状态。签约成功则返回success
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// 扩展参数，可为空
	ExtendField *string `json:"extend_field,omitempty"`
	// 通知方式设置。
	NotifyConfig *string `json:"notify_config,omitempty"`
	// 商户生成的代扣协议ID
	OutAgreementId *string `json:"out_agreement_id,omitempty"`
	// 支付方式设置
	PayConfig []string `json:"pay_config,omitempty"`
	// 签约时间
	SignDate *string `json:"sign_date,omitempty"`
}

// NewAlipayEbppPdeductSignAddResponseModel instantiates a new AlipayEbppPdeductSignAddResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppPdeductSignAddResponseModel() *AlipayEbppPdeductSignAddResponseModel {
	this := AlipayEbppPdeductSignAddResponseModel{}
	return &this
}

// NewAlipayEbppPdeductSignAddResponseModelWithDefaults instantiates a new AlipayEbppPdeductSignAddResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppPdeductSignAddResponseModelWithDefaults() *AlipayEbppPdeductSignAddResponseModel {
	this := AlipayEbppPdeductSignAddResponseModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignAddResponseModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayEbppPdeductSignAddResponseModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetAgreementStatus returns the AgreementStatus field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignAddResponseModel) GetAgreementStatus() string {
	if o == nil || IsNil(o.AgreementStatus) {
		var ret string
		return ret
	}
	return *o.AgreementStatus
}

// GetAgreementStatusOk returns a tuple with the AgreementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) GetAgreementStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementStatus) {
		return nil, false
	}
	return o.AgreementStatus, true
}

// HasAgreementStatus returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) HasAgreementStatus() bool {
	if o != nil && !IsNil(o.AgreementStatus) {
		return true
	}

	return false
}

// SetAgreementStatus gets a reference to the given string and assigns it to the AgreementStatus field.
func (o *AlipayEbppPdeductSignAddResponseModel) SetAgreementStatus(v string) {
	o.AgreementStatus = &v
}

// GetExtendField returns the ExtendField field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignAddResponseModel) GetExtendField() string {
	if o == nil || IsNil(o.ExtendField) {
		var ret string
		return ret
	}
	return *o.ExtendField
}

// GetExtendFieldOk returns a tuple with the ExtendField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) GetExtendFieldOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendField) {
		return nil, false
	}
	return o.ExtendField, true
}

// HasExtendField returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) HasExtendField() bool {
	if o != nil && !IsNil(o.ExtendField) {
		return true
	}

	return false
}

// SetExtendField gets a reference to the given string and assigns it to the ExtendField field.
func (o *AlipayEbppPdeductSignAddResponseModel) SetExtendField(v string) {
	o.ExtendField = &v
}

// GetNotifyConfig returns the NotifyConfig field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignAddResponseModel) GetNotifyConfig() string {
	if o == nil || IsNil(o.NotifyConfig) {
		var ret string
		return ret
	}
	return *o.NotifyConfig
}

// GetNotifyConfigOk returns a tuple with the NotifyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) GetNotifyConfigOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyConfig) {
		return nil, false
	}
	return o.NotifyConfig, true
}

// HasNotifyConfig returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) HasNotifyConfig() bool {
	if o != nil && !IsNil(o.NotifyConfig) {
		return true
	}

	return false
}

// SetNotifyConfig gets a reference to the given string and assigns it to the NotifyConfig field.
func (o *AlipayEbppPdeductSignAddResponseModel) SetNotifyConfig(v string) {
	o.NotifyConfig = &v
}

// GetOutAgreementId returns the OutAgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignAddResponseModel) GetOutAgreementId() string {
	if o == nil || IsNil(o.OutAgreementId) {
		var ret string
		return ret
	}
	return *o.OutAgreementId
}

// GetOutAgreementIdOk returns a tuple with the OutAgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) GetOutAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutAgreementId) {
		return nil, false
	}
	return o.OutAgreementId, true
}

// HasOutAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) HasOutAgreementId() bool {
	if o != nil && !IsNil(o.OutAgreementId) {
		return true
	}

	return false
}

// SetOutAgreementId gets a reference to the given string and assigns it to the OutAgreementId field.
func (o *AlipayEbppPdeductSignAddResponseModel) SetOutAgreementId(v string) {
	o.OutAgreementId = &v
}

// GetPayConfig returns the PayConfig field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignAddResponseModel) GetPayConfig() []string {
	if o == nil || IsNil(o.PayConfig) {
		var ret []string
		return ret
	}
	return o.PayConfig
}

// GetPayConfigOk returns a tuple with the PayConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) GetPayConfigOk() ([]string, bool) {
	if o == nil || IsNil(o.PayConfig) {
		return nil, false
	}
	return o.PayConfig, true
}

// HasPayConfig returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) HasPayConfig() bool {
	if o != nil && !IsNil(o.PayConfig) {
		return true
	}

	return false
}

// SetPayConfig gets a reference to the given []string and assigns it to the PayConfig field.
func (o *AlipayEbppPdeductSignAddResponseModel) SetPayConfig(v []string) {
	o.PayConfig = v
}

// GetSignDate returns the SignDate field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignAddResponseModel) GetSignDate() string {
	if o == nil || IsNil(o.SignDate) {
		var ret string
		return ret
	}
	return *o.SignDate
}

// GetSignDateOk returns a tuple with the SignDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) GetSignDateOk() (*string, bool) {
	if o == nil || IsNil(o.SignDate) {
		return nil, false
	}
	return o.SignDate, true
}

// HasSignDate returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignAddResponseModel) HasSignDate() bool {
	if o != nil && !IsNil(o.SignDate) {
		return true
	}

	return false
}

// SetSignDate gets a reference to the given string and assigns it to the SignDate field.
func (o *AlipayEbppPdeductSignAddResponseModel) SetSignDate(v string) {
	o.SignDate = &v
}

func (o AlipayEbppPdeductSignAddResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppPdeductSignAddResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.AgreementStatus) {
		toSerialize["agreement_status"] = o.AgreementStatus
	}
	if !IsNil(o.ExtendField) {
		toSerialize["extend_field"] = o.ExtendField
	}
	if !IsNil(o.NotifyConfig) {
		toSerialize["notify_config"] = o.NotifyConfig
	}
	if !IsNil(o.OutAgreementId) {
		toSerialize["out_agreement_id"] = o.OutAgreementId
	}
	if !IsNil(o.PayConfig) {
		toSerialize["pay_config"] = o.PayConfig
	}
	if !IsNil(o.SignDate) {
		toSerialize["sign_date"] = o.SignDate
	}
	return toSerialize, nil
}

type NullableAlipayEbppPdeductSignAddResponseModel struct {
	value *AlipayEbppPdeductSignAddResponseModel
	isSet bool
}

func (v NullableAlipayEbppPdeductSignAddResponseModel) Get() *AlipayEbppPdeductSignAddResponseModel {
	return v.value
}

func (v *NullableAlipayEbppPdeductSignAddResponseModel) Set(val *AlipayEbppPdeductSignAddResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppPdeductSignAddResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppPdeductSignAddResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppPdeductSignAddResponseModel(val *AlipayEbppPdeductSignAddResponseModel) *NullableAlipayEbppPdeductSignAddResponseModel {
	return &NullableAlipayEbppPdeductSignAddResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppPdeductSignAddResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppPdeductSignAddResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
