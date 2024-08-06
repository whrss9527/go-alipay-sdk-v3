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

// checks if the AlipayOpenSpOperationResultQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSpOperationResultQueryResponseModel{}

// AlipayOpenSpOperationResultQueryResponseModel struct for AlipayOpenSpOperationResultQueryResponseModel
type AlipayOpenSpOperationResultQueryResponseModel struct {
	// 商户支付宝pid。仅间连场景，且存在绑定关系时返回值。
	BindUserId *string `json:"bind_user_id,omitempty"`
	// 代运营操作结果。 SUCCESS：代表成功。 PROCESS：待商家确认中。 NO_PERMISSION：表示当前商家支付宝账号无权限操作。需要提醒商家切换成发起授权时指定的支付宝账号。 NONE：表示不存在代运营绑定或授权关系。 NONE_ACCOUNT：间连商家推荐支付宝账号列表为空。
	HandleStatus *string `json:"handle_status,omitempty"`
	// 支付宝商户号。间连场景为商户smid，直连场景为商户支付宝pid
	MerchantNo *string `json:"merchant_no,omitempty"`
}

// NewAlipayOpenSpOperationResultQueryResponseModel instantiates a new AlipayOpenSpOperationResultQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSpOperationResultQueryResponseModel() *AlipayOpenSpOperationResultQueryResponseModel {
	this := AlipayOpenSpOperationResultQueryResponseModel{}
	return &this
}

// NewAlipayOpenSpOperationResultQueryResponseModelWithDefaults instantiates a new AlipayOpenSpOperationResultQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSpOperationResultQueryResponseModelWithDefaults() *AlipayOpenSpOperationResultQueryResponseModel {
	this := AlipayOpenSpOperationResultQueryResponseModel{}
	return &this
}

// GetBindUserId returns the BindUserId field value if set, zero value otherwise.
func (o *AlipayOpenSpOperationResultQueryResponseModel) GetBindUserId() string {
	if o == nil || IsNil(o.BindUserId) {
		var ret string
		return ret
	}
	return *o.BindUserId
}

// GetBindUserIdOk returns a tuple with the BindUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpOperationResultQueryResponseModel) GetBindUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.BindUserId) {
		return nil, false
	}
	return o.BindUserId, true
}

// HasBindUserId returns a boolean if a field has been set.
func (o *AlipayOpenSpOperationResultQueryResponseModel) HasBindUserId() bool {
	if o != nil && !IsNil(o.BindUserId) {
		return true
	}

	return false
}

// SetBindUserId gets a reference to the given string and assigns it to the BindUserId field.
func (o *AlipayOpenSpOperationResultQueryResponseModel) SetBindUserId(v string) {
	o.BindUserId = &v
}

// GetHandleStatus returns the HandleStatus field value if set, zero value otherwise.
func (o *AlipayOpenSpOperationResultQueryResponseModel) GetHandleStatus() string {
	if o == nil || IsNil(o.HandleStatus) {
		var ret string
		return ret
	}
	return *o.HandleStatus
}

// GetHandleStatusOk returns a tuple with the HandleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpOperationResultQueryResponseModel) GetHandleStatusOk() (*string, bool) {
	if o == nil || IsNil(o.HandleStatus) {
		return nil, false
	}
	return o.HandleStatus, true
}

// HasHandleStatus returns a boolean if a field has been set.
func (o *AlipayOpenSpOperationResultQueryResponseModel) HasHandleStatus() bool {
	if o != nil && !IsNil(o.HandleStatus) {
		return true
	}

	return false
}

// SetHandleStatus gets a reference to the given string and assigns it to the HandleStatus field.
func (o *AlipayOpenSpOperationResultQueryResponseModel) SetHandleStatus(v string) {
	o.HandleStatus = &v
}

// GetMerchantNo returns the MerchantNo field value if set, zero value otherwise.
func (o *AlipayOpenSpOperationResultQueryResponseModel) GetMerchantNo() string {
	if o == nil || IsNil(o.MerchantNo) {
		var ret string
		return ret
	}
	return *o.MerchantNo
}

// GetMerchantNoOk returns a tuple with the MerchantNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpOperationResultQueryResponseModel) GetMerchantNoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantNo) {
		return nil, false
	}
	return o.MerchantNo, true
}

// HasMerchantNo returns a boolean if a field has been set.
func (o *AlipayOpenSpOperationResultQueryResponseModel) HasMerchantNo() bool {
	if o != nil && !IsNil(o.MerchantNo) {
		return true
	}

	return false
}

// SetMerchantNo gets a reference to the given string and assigns it to the MerchantNo field.
func (o *AlipayOpenSpOperationResultQueryResponseModel) SetMerchantNo(v string) {
	o.MerchantNo = &v
}

func (o AlipayOpenSpOperationResultQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSpOperationResultQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BindUserId) {
		toSerialize["bind_user_id"] = o.BindUserId
	}
	if !IsNil(o.HandleStatus) {
		toSerialize["handle_status"] = o.HandleStatus
	}
	if !IsNil(o.MerchantNo) {
		toSerialize["merchant_no"] = o.MerchantNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenSpOperationResultQueryResponseModel struct {
	value *AlipayOpenSpOperationResultQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenSpOperationResultQueryResponseModel) Get() *AlipayOpenSpOperationResultQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSpOperationResultQueryResponseModel) Set(val *AlipayOpenSpOperationResultQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpOperationResultQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpOperationResultQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpOperationResultQueryResponseModel(val *AlipayOpenSpOperationResultQueryResponseModel) *NullableAlipayOpenSpOperationResultQueryResponseModel {
	return &NullableAlipayOpenSpOperationResultQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSpOperationResultQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpOperationResultQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
