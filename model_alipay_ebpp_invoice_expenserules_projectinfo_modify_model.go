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

// checks if the AlipayEbppInvoiceExpenserulesProjectinfoModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpenserulesProjectinfoModifyModel{}

// AlipayEbppInvoiceExpenserulesProjectinfoModifyModel struct for AlipayEbppInvoiceExpenserulesProjectinfoModifyModel
type AlipayEbppInvoiceExpenserulesProjectinfoModifyModel struct {
	// 企业ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 有效期截止（yyyy-MM-dd HH:mm:ss） 特殊说明： 1）与起始时间必须同时传，且大于起始时间 2）不传则默认不修改
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// 有效期起始（yyyy-MM-dd HH:mm:ss） 特殊说明： 1）与截止时间必须同时传，且小于截止时间 2）不传入则默认不修改
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`
	// 项目名 特殊说明：不传入则默认不修改
	ProjectName *string `json:"project_name,omitempty"`
}

// NewAlipayEbppInvoiceExpenserulesProjectinfoModifyModel instantiates a new AlipayEbppInvoiceExpenserulesProjectinfoModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpenserulesProjectinfoModifyModel() *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel {
	this := AlipayEbppInvoiceExpenserulesProjectinfoModifyModel{}
	return &this
}

// NewAlipayEbppInvoiceExpenserulesProjectinfoModifyModelWithDefaults instantiates a new AlipayEbppInvoiceExpenserulesProjectinfoModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpenserulesProjectinfoModifyModelWithDefaults() *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel {
	this := AlipayEbppInvoiceExpenserulesProjectinfoModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEffectiveEndDate returns the EffectiveEndDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetEffectiveEndDate() string {
	if o == nil || IsNil(o.EffectiveEndDate) {
		var ret string
		return ret
	}
	return *o.EffectiveEndDate
}

// GetEffectiveEndDateOk returns a tuple with the EffectiveEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetEffectiveEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveEndDate) {
		return nil, false
	}
	return o.EffectiveEndDate, true
}

// HasEffectiveEndDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) HasEffectiveEndDate() bool {
	if o != nil && !IsNil(o.EffectiveEndDate) {
		return true
	}

	return false
}

// SetEffectiveEndDate gets a reference to the given string and assigns it to the EffectiveEndDate field.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) SetEffectiveEndDate(v string) {
	o.EffectiveEndDate = &v
}

// GetEffectiveStartDate returns the EffectiveStartDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetEffectiveStartDate() string {
	if o == nil || IsNil(o.EffectiveStartDate) {
		var ret string
		return ret
	}
	return *o.EffectiveStartDate
}

// GetEffectiveStartDateOk returns a tuple with the EffectiveStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetEffectiveStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStartDate) {
		return nil, false
	}
	return o.EffectiveStartDate, true
}

// HasEffectiveStartDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) HasEffectiveStartDate() bool {
	if o != nil && !IsNil(o.EffectiveStartDate) {
		return true
	}

	return false
}

// SetEffectiveStartDate gets a reference to the given string and assigns it to the EffectiveStartDate field.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) SetEffectiveStartDate(v string) {
	o.EffectiveStartDate = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) SetProjectName(v string) {
	o.ProjectName = &v
}

func (o AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.EffectiveEndDate) {
		toSerialize["effective_end_date"] = o.EffectiveEndDate
	}
	if !IsNil(o.EffectiveStartDate) {
		toSerialize["effective_start_date"] = o.EffectiveStartDate
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.ProjectName) {
		toSerialize["project_name"] = o.ProjectName
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel struct {
	value *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel) Get() *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel) Set(val *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel(val *AlipayEbppInvoiceExpenserulesProjectinfoModifyModel) *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel {
	return &NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


