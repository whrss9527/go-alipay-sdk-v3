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

// checks if the AlipayEbppInvoiceDetailOutputQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceDetailOutputQueryModel{}

// AlipayEbppInvoiceDetailOutputQueryModel struct for AlipayEbppInvoiceDetailOutputQueryModel
type AlipayEbppInvoiceDetailOutputQueryModel struct {
	// 发票代码 长度限制（10-12位），全电票则为空
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// 发票号码 长度限制（8-20位）
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// 支付宝用户 id
	OpenId *string `json:"open_id,omitempty"`
	// 获取发票明细应用场景。固定为 INVOICE_EXPENSE  表示发票报销。
	Scene *string `json:"scene,omitempty"`
	// 是否跳过发票报销状态同步；当为true时，跳过报销状态同步校验。默认为false，需要先做报销状态同步
	SkipExpenseProgressSync *bool `json:"skip_expense_progress_sync,omitempty"`
	// 发票归属用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayEbppInvoiceDetailOutputQueryModel instantiates a new AlipayEbppInvoiceDetailOutputQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceDetailOutputQueryModel() *AlipayEbppInvoiceDetailOutputQueryModel {
	this := AlipayEbppInvoiceDetailOutputQueryModel{}
	return &this
}

// NewAlipayEbppInvoiceDetailOutputQueryModelWithDefaults instantiates a new AlipayEbppInvoiceDetailOutputQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceDetailOutputQueryModelWithDefaults() *AlipayEbppInvoiceDetailOutputQueryModel {
	this := AlipayEbppInvoiceDetailOutputQueryModel{}
	return &this
}

// GetInvoiceCode returns the InvoiceCode field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetInvoiceCode() string {
	if o == nil || IsNil(o.InvoiceCode) {
		var ret string
		return ret
	}
	return *o.InvoiceCode
}

// GetInvoiceCodeOk returns a tuple with the InvoiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetInvoiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCode) {
		return nil, false
	}
	return o.InvoiceCode, true
}

// HasInvoiceCode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) HasInvoiceCode() bool {
	if o != nil && !IsNil(o.InvoiceCode) {
		return true
	}

	return false
}

// SetInvoiceCode gets a reference to the given string and assigns it to the InvoiceCode field.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) SetInvoiceCode(v string) {
	o.InvoiceCode = &v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetInvoiceNo() string {
	if o == nil || IsNil(o.InvoiceNo) {
		var ret string
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetInvoiceNoOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNo) {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) HasInvoiceNo() bool {
	if o != nil && !IsNil(o.InvoiceNo) {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given string and assigns it to the InvoiceNo field.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) SetInvoiceNo(v string) {
	o.InvoiceNo = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetScene returns the Scene field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetScene() string {
	if o == nil || IsNil(o.Scene) {
		var ret string
		return ret
	}
	return *o.Scene
}

// GetSceneOk returns a tuple with the Scene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetSceneOk() (*string, bool) {
	if o == nil || IsNil(o.Scene) {
		return nil, false
	}
	return o.Scene, true
}

// HasScene returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) HasScene() bool {
	if o != nil && !IsNil(o.Scene) {
		return true
	}

	return false
}

// SetScene gets a reference to the given string and assigns it to the Scene field.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) SetScene(v string) {
	o.Scene = &v
}

// GetSkipExpenseProgressSync returns the SkipExpenseProgressSync field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetSkipExpenseProgressSync() bool {
	if o == nil || IsNil(o.SkipExpenseProgressSync) {
		var ret bool
		return ret
	}
	return *o.SkipExpenseProgressSync
}

// GetSkipExpenseProgressSyncOk returns a tuple with the SkipExpenseProgressSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetSkipExpenseProgressSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipExpenseProgressSync) {
		return nil, false
	}
	return o.SkipExpenseProgressSync, true
}

// HasSkipExpenseProgressSync returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) HasSkipExpenseProgressSync() bool {
	if o != nil && !IsNil(o.SkipExpenseProgressSync) {
		return true
	}

	return false
}

// SetSkipExpenseProgressSync gets a reference to the given bool and assigns it to the SkipExpenseProgressSync field.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) SetSkipExpenseProgressSync(v bool) {
	o.SkipExpenseProgressSync = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayEbppInvoiceDetailOutputQueryModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayEbppInvoiceDetailOutputQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceDetailOutputQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceCode) {
		toSerialize["invoice_code"] = o.InvoiceCode
	}
	if !IsNil(o.InvoiceNo) {
		toSerialize["invoice_no"] = o.InvoiceNo
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.Scene) {
		toSerialize["scene"] = o.Scene
	}
	if !IsNil(o.SkipExpenseProgressSync) {
		toSerialize["skip_expense_progress_sync"] = o.SkipExpenseProgressSync
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceDetailOutputQueryModel struct {
	value *AlipayEbppInvoiceDetailOutputQueryModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryModel) Get() *AlipayEbppInvoiceDetailOutputQueryModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryModel) Set(val *AlipayEbppInvoiceDetailOutputQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceDetailOutputQueryModel(val *AlipayEbppInvoiceDetailOutputQueryModel) *NullableAlipayEbppInvoiceDetailOutputQueryModel {
	return &NullableAlipayEbppInvoiceDetailOutputQueryModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
