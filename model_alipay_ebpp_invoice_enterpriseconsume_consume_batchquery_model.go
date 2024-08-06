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

// checks if the AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel{}

// AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel struct for AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel
type AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel struct {
	// 企业ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 交易结束时间，格式：yyyy-MM-dd HH:mm:ss
	ConsumptionEnd *string `json:"consumption_end,omitempty"`
	// 交易开始时间，格式：yyyy-MM-dd HH:mm:ss
	ConsumptionStart *string `json:"consumption_start,omitempty"`
	// 员工支付宝UID列表，单次传入最大员工数量为10
	EmployeeList []string `json:"employee_list,omitempty"`
	// 员工支付宝UID列表，单次传入最大员工数量为10
	EmployeeOpenIds []string `json:"employee_open_ids,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel instantiates a new AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel() *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel {
	this := AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModelWithDefaults() *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel {
	this := AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetConsumptionEnd returns the ConsumptionEnd field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetConsumptionEnd() string {
	if o == nil || IsNil(o.ConsumptionEnd) {
		var ret string
		return ret
	}
	return *o.ConsumptionEnd
}

// GetConsumptionEndOk returns a tuple with the ConsumptionEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetConsumptionEndOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumptionEnd) {
		return nil, false
	}
	return o.ConsumptionEnd, true
}

// HasConsumptionEnd returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) HasConsumptionEnd() bool {
	if o != nil && !IsNil(o.ConsumptionEnd) {
		return true
	}

	return false
}

// SetConsumptionEnd gets a reference to the given string and assigns it to the ConsumptionEnd field.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) SetConsumptionEnd(v string) {
	o.ConsumptionEnd = &v
}

// GetConsumptionStart returns the ConsumptionStart field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetConsumptionStart() string {
	if o == nil || IsNil(o.ConsumptionStart) {
		var ret string
		return ret
	}
	return *o.ConsumptionStart
}

// GetConsumptionStartOk returns a tuple with the ConsumptionStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetConsumptionStartOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumptionStart) {
		return nil, false
	}
	return o.ConsumptionStart, true
}

// HasConsumptionStart returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) HasConsumptionStart() bool {
	if o != nil && !IsNil(o.ConsumptionStart) {
		return true
	}

	return false
}

// SetConsumptionStart gets a reference to the given string and assigns it to the ConsumptionStart field.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) SetConsumptionStart(v string) {
	o.ConsumptionStart = &v
}

// GetEmployeeList returns the EmployeeList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetEmployeeList() []string {
	if o == nil || IsNil(o.EmployeeList) {
		var ret []string
		return ret
	}
	return o.EmployeeList
}

// GetEmployeeListOk returns a tuple with the EmployeeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetEmployeeListOk() ([]string, bool) {
	if o == nil || IsNil(o.EmployeeList) {
		return nil, false
	}
	return o.EmployeeList, true
}

// HasEmployeeList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) HasEmployeeList() bool {
	if o != nil && !IsNil(o.EmployeeList) {
		return true
	}

	return false
}

// SetEmployeeList gets a reference to the given []string and assigns it to the EmployeeList field.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) SetEmployeeList(v []string) {
	o.EmployeeList = v
}

// GetEmployeeOpenIds returns the EmployeeOpenIds field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetEmployeeOpenIds() []string {
	if o == nil || IsNil(o.EmployeeOpenIds) {
		var ret []string
		return ret
	}
	return o.EmployeeOpenIds
}

// GetEmployeeOpenIdsOk returns a tuple with the EmployeeOpenIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) GetEmployeeOpenIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EmployeeOpenIds) {
		return nil, false
	}
	return o.EmployeeOpenIds, true
}

// HasEmployeeOpenIds returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) HasEmployeeOpenIds() bool {
	if o != nil && !IsNil(o.EmployeeOpenIds) {
		return true
	}

	return false
}

// SetEmployeeOpenIds gets a reference to the given []string and assigns it to the EmployeeOpenIds field.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) SetEmployeeOpenIds(v []string) {
	o.EmployeeOpenIds = v
}

func (o AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.ConsumptionEnd) {
		toSerialize["consumption_end"] = o.ConsumptionEnd
	}
	if !IsNil(o.ConsumptionStart) {
		toSerialize["consumption_start"] = o.ConsumptionStart
	}
	if !IsNil(o.EmployeeList) {
		toSerialize["employee_list"] = o.EmployeeList
	}
	if !IsNil(o.EmployeeOpenIds) {
		toSerialize["employee_open_ids"] = o.EmployeeOpenIds
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel struct {
	value *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) Get() *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) Set(val *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel(val *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel {
	return &NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


