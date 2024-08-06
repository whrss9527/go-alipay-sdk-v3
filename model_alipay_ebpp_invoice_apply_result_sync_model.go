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

// checks if the AlipayEbppInvoiceApplyResultSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceApplyResultSyncModel{}

// AlipayEbppInvoiceApplyResultSyncModel struct for AlipayEbppInvoiceApplyResultSyncModel
type AlipayEbppInvoiceApplyResultSyncModel struct {
	// 支付宝发起开票申请的id，该id具有唯一性，该字段由支付宝向税控发起申请的时候带过去，作为支付宝向税控开票申请的唯一标志
	ApplyId *string `json:"apply_id,omitempty"`
	// 支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果。  取值如下：  SUCCESS:成功;FAIL:失败
	Result *string `json:"result,omitempty"`
	// 结果码，支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果进行详细说明的结果码。  取值如下：  成功(SUCCESS),  参数不合法(PARAMETER_ILLEGAL),  商户税控设备异常(MERCHANT_TAX_DEVICE_ERROR).
	ResultCode *string `json:"result_code,omitempty"`
	// 结果描述，支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果描述信息。
	ResultMsg *string `json:"result_msg,omitempty"`
	// 该字段是税控商或ISV收到支付宝开票请求后生成的申请id，由税控商或isv自己生成，该id具有唯一性  当ISV接入时是按照tax_apply_id来查询发票信息时，该字段必填。
	TaxApplyId *string `json:"tax_apply_id,omitempty"`
}

// NewAlipayEbppInvoiceApplyResultSyncModel instantiates a new AlipayEbppInvoiceApplyResultSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceApplyResultSyncModel() *AlipayEbppInvoiceApplyResultSyncModel {
	this := AlipayEbppInvoiceApplyResultSyncModel{}
	return &this
}

// NewAlipayEbppInvoiceApplyResultSyncModelWithDefaults instantiates a new AlipayEbppInvoiceApplyResultSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceApplyResultSyncModelWithDefaults() *AlipayEbppInvoiceApplyResultSyncModel {
	this := AlipayEbppInvoiceApplyResultSyncModel{}
	return &this
}

// GetApplyId returns the ApplyId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetApplyId() string {
	if o == nil || IsNil(o.ApplyId) {
		var ret string
		return ret
	}
	return *o.ApplyId
}

// GetApplyIdOk returns a tuple with the ApplyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetApplyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyId) {
		return nil, false
	}
	return o.ApplyId, true
}

// HasApplyId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) HasApplyId() bool {
	if o != nil && !IsNil(o.ApplyId) {
		return true
	}

	return false
}

// SetApplyId gets a reference to the given string and assigns it to the ApplyId field.
func (o *AlipayEbppInvoiceApplyResultSyncModel) SetApplyId(v string) {
	o.ApplyId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *AlipayEbppInvoiceApplyResultSyncModel) SetResult(v string) {
	o.Result = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *AlipayEbppInvoiceApplyResultSyncModel) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMsg returns the ResultMsg field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetResultMsg() string {
	if o == nil || IsNil(o.ResultMsg) {
		var ret string
		return ret
	}
	return *o.ResultMsg
}

// GetResultMsgOk returns a tuple with the ResultMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetResultMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMsg) {
		return nil, false
	}
	return o.ResultMsg, true
}

// HasResultMsg returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) HasResultMsg() bool {
	if o != nil && !IsNil(o.ResultMsg) {
		return true
	}

	return false
}

// SetResultMsg gets a reference to the given string and assigns it to the ResultMsg field.
func (o *AlipayEbppInvoiceApplyResultSyncModel) SetResultMsg(v string) {
	o.ResultMsg = &v
}

// GetTaxApplyId returns the TaxApplyId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetTaxApplyId() string {
	if o == nil || IsNil(o.TaxApplyId) {
		var ret string
		return ret
	}
	return *o.TaxApplyId
}

// GetTaxApplyIdOk returns a tuple with the TaxApplyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) GetTaxApplyIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxApplyId) {
		return nil, false
	}
	return o.TaxApplyId, true
}

// HasTaxApplyId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyResultSyncModel) HasTaxApplyId() bool {
	if o != nil && !IsNil(o.TaxApplyId) {
		return true
	}

	return false
}

// SetTaxApplyId gets a reference to the given string and assigns it to the TaxApplyId field.
func (o *AlipayEbppInvoiceApplyResultSyncModel) SetTaxApplyId(v string) {
	o.TaxApplyId = &v
}

func (o AlipayEbppInvoiceApplyResultSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceApplyResultSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyId) {
		toSerialize["apply_id"] = o.ApplyId
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.ResultCode) {
		toSerialize["result_code"] = o.ResultCode
	}
	if !IsNil(o.ResultMsg) {
		toSerialize["result_msg"] = o.ResultMsg
	}
	if !IsNil(o.TaxApplyId) {
		toSerialize["tax_apply_id"] = o.TaxApplyId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceApplyResultSyncModel struct {
	value *AlipayEbppInvoiceApplyResultSyncModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceApplyResultSyncModel) Get() *AlipayEbppInvoiceApplyResultSyncModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceApplyResultSyncModel) Set(val *AlipayEbppInvoiceApplyResultSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceApplyResultSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceApplyResultSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceApplyResultSyncModel(val *AlipayEbppInvoiceApplyResultSyncModel) *NullableAlipayEbppInvoiceApplyResultSyncModel {
	return &NullableAlipayEbppInvoiceApplyResultSyncModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceApplyResultSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceApplyResultSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
