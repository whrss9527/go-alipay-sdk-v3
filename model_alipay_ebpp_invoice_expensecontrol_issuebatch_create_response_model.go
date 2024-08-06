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

// checks if the AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel{}

// AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel struct for AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel
type AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel struct {
	// 发放批次id
	IssueBatchId *string `json:"issue_batch_id,omitempty"`
	// 校验失败的数据
	IssueQuotaCheckFailedList []IssueQuotaCheckInfo `json:"issue_quota_check_failed_list,omitempty"`
}

// NewAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel instantiates a new AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel() *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel {
	this := AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModelWithDefaults() *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel {
	this := AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel{}
	return &this
}

// GetIssueBatchId returns the IssueBatchId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) GetIssueBatchId() string {
	if o == nil || IsNil(o.IssueBatchId) {
		var ret string
		return ret
	}
	return *o.IssueBatchId
}

// GetIssueBatchIdOk returns a tuple with the IssueBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) GetIssueBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.IssueBatchId) {
		return nil, false
	}
	return o.IssueBatchId, true
}

// HasIssueBatchId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) HasIssueBatchId() bool {
	if o != nil && !IsNil(o.IssueBatchId) {
		return true
	}

	return false
}

// SetIssueBatchId gets a reference to the given string and assigns it to the IssueBatchId field.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) SetIssueBatchId(v string) {
	o.IssueBatchId = &v
}

// GetIssueQuotaCheckFailedList returns the IssueQuotaCheckFailedList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) GetIssueQuotaCheckFailedList() []IssueQuotaCheckInfo {
	if o == nil || IsNil(o.IssueQuotaCheckFailedList) {
		var ret []IssueQuotaCheckInfo
		return ret
	}
	return o.IssueQuotaCheckFailedList
}

// GetIssueQuotaCheckFailedListOk returns a tuple with the IssueQuotaCheckFailedList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) GetIssueQuotaCheckFailedListOk() ([]IssueQuotaCheckInfo, bool) {
	if o == nil || IsNil(o.IssueQuotaCheckFailedList) {
		return nil, false
	}
	return o.IssueQuotaCheckFailedList, true
}

// HasIssueQuotaCheckFailedList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) HasIssueQuotaCheckFailedList() bool {
	if o != nil && !IsNil(o.IssueQuotaCheckFailedList) {
		return true
	}

	return false
}

// SetIssueQuotaCheckFailedList gets a reference to the given []IssueQuotaCheckInfo and assigns it to the IssueQuotaCheckFailedList field.
func (o *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) SetIssueQuotaCheckFailedList(v []IssueQuotaCheckInfo) {
	o.IssueQuotaCheckFailedList = v
}

func (o AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IssueBatchId) {
		toSerialize["issue_batch_id"] = o.IssueBatchId
	}
	if !IsNil(o.IssueQuotaCheckFailedList) {
		toSerialize["issue_quota_check_failed_list"] = o.IssueQuotaCheckFailedList
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel struct {
	value *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) Get() *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) Set(val *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel(val *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel {
	return &NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
