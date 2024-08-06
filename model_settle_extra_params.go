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

// checks if the SettleExtraParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettleExtraParams{}

// SettleExtraParams struct for SettleExtraParams
type SettleExtraParams struct {
	// quit_type为USER_CANCEL_QUIT或者SETTLE_APPLY_QUIT
	QuitType *string `json:"quit_type,omitempty"`
	// 商链通权益抵扣信息,  当前只有运营商使用,  未使用权益抵扣忽略该字段  fundRuleList为对应的出资信息，目前该列表不支持多对象，长度限制为1。其中fundAgreementNo为本笔交易实际出资的协议，由(商链通二方通用版本接入文档)2.3.2获取，fundAmount为该出资协议出资的金额。  bizScene和subBizScene为(商链通二方通用版本接入文档)2.1.3中支付宝侧分配的信息。  bizMode=MERCHANT_ORDER为固定值。
	ScenePayLinkInfo *string `json:"scene_pay_link_info,omitempty"`
	// action_type选择PAY_TO_ZERO时必填， \"SERVICE_CANCELED\":\"服务已取消\", \"OTHER_CHANNEL_PERFORMANCE\":\"户已通过其他方式履约\"
	SettleAdjustReason *string `json:"settle_adjust_reason,omitempty"`
}

// NewSettleExtraParams instantiates a new SettleExtraParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettleExtraParams() *SettleExtraParams {
	this := SettleExtraParams{}
	return &this
}

// NewSettleExtraParamsWithDefaults instantiates a new SettleExtraParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettleExtraParamsWithDefaults() *SettleExtraParams {
	this := SettleExtraParams{}
	return &this
}

// GetQuitType returns the QuitType field value if set, zero value otherwise.
func (o *SettleExtraParams) GetQuitType() string {
	if o == nil || IsNil(o.QuitType) {
		var ret string
		return ret
	}
	return *o.QuitType
}

// GetQuitTypeOk returns a tuple with the QuitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleExtraParams) GetQuitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuitType) {
		return nil, false
	}
	return o.QuitType, true
}

// HasQuitType returns a boolean if a field has been set.
func (o *SettleExtraParams) HasQuitType() bool {
	if o != nil && !IsNil(o.QuitType) {
		return true
	}

	return false
}

// SetQuitType gets a reference to the given string and assigns it to the QuitType field.
func (o *SettleExtraParams) SetQuitType(v string) {
	o.QuitType = &v
}

// GetScenePayLinkInfo returns the ScenePayLinkInfo field value if set, zero value otherwise.
func (o *SettleExtraParams) GetScenePayLinkInfo() string {
	if o == nil || IsNil(o.ScenePayLinkInfo) {
		var ret string
		return ret
	}
	return *o.ScenePayLinkInfo
}

// GetScenePayLinkInfoOk returns a tuple with the ScenePayLinkInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleExtraParams) GetScenePayLinkInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ScenePayLinkInfo) {
		return nil, false
	}
	return o.ScenePayLinkInfo, true
}

// HasScenePayLinkInfo returns a boolean if a field has been set.
func (o *SettleExtraParams) HasScenePayLinkInfo() bool {
	if o != nil && !IsNil(o.ScenePayLinkInfo) {
		return true
	}

	return false
}

// SetScenePayLinkInfo gets a reference to the given string and assigns it to the ScenePayLinkInfo field.
func (o *SettleExtraParams) SetScenePayLinkInfo(v string) {
	o.ScenePayLinkInfo = &v
}

// GetSettleAdjustReason returns the SettleAdjustReason field value if set, zero value otherwise.
func (o *SettleExtraParams) GetSettleAdjustReason() string {
	if o == nil || IsNil(o.SettleAdjustReason) {
		var ret string
		return ret
	}
	return *o.SettleAdjustReason
}

// GetSettleAdjustReasonOk returns a tuple with the SettleAdjustReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettleExtraParams) GetSettleAdjustReasonOk() (*string, bool) {
	if o == nil || IsNil(o.SettleAdjustReason) {
		return nil, false
	}
	return o.SettleAdjustReason, true
}

// HasSettleAdjustReason returns a boolean if a field has been set.
func (o *SettleExtraParams) HasSettleAdjustReason() bool {
	if o != nil && !IsNil(o.SettleAdjustReason) {
		return true
	}

	return false
}

// SetSettleAdjustReason gets a reference to the given string and assigns it to the SettleAdjustReason field.
func (o *SettleExtraParams) SetSettleAdjustReason(v string) {
	o.SettleAdjustReason = &v
}

func (o SettleExtraParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettleExtraParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuitType) {
		toSerialize["quit_type"] = o.QuitType
	}
	if !IsNil(o.ScenePayLinkInfo) {
		toSerialize["scene_pay_link_info"] = o.ScenePayLinkInfo
	}
	if !IsNil(o.SettleAdjustReason) {
		toSerialize["settle_adjust_reason"] = o.SettleAdjustReason
	}
	return toSerialize, nil
}

type NullableSettleExtraParams struct {
	value *SettleExtraParams
	isSet bool
}

func (v NullableSettleExtraParams) Get() *SettleExtraParams {
	return v.value
}

func (v *NullableSettleExtraParams) Set(val *SettleExtraParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSettleExtraParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSettleExtraParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettleExtraParams(val *SettleExtraParams) *NullableSettleExtraParams {
	return &NullableSettleExtraParams{value: val, isSet: true}
}

func (v NullableSettleExtraParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettleExtraParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

