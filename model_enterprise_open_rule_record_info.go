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

// checks if the EnterpriseOpenRuleRecordInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterpriseOpenRuleRecordInfo{}

// EnterpriseOpenRuleRecordInfo struct for EnterpriseOpenRuleRecordInfo
type EnterpriseOpenRuleRecordInfo struct {
	// 开票规则账单日
	BillMonthDay *int32 `json:"bill_month_day,omitempty"`
	// 开票规则生效日期
	EffectiveStart *string `json:"effective_start,omitempty"`
	// 企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// 开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
	// 开票规则记录ID
	InvoiceRuleRecordId *string `json:"invoice_rule_record_id,omitempty"`
	// 发票抬头ID
	InvoiceTitleId *string `json:"invoice_title_id,omitempty"`
	// 开票申请方
	OpenApplyer *string `json:"open_applyer,omitempty"`
	// 开票模式
	OpenMode *string `json:"open_mode,omitempty"`
	// 开票申请类型
	OpenType *string `json:"open_type,omitempty"`
	// 所有者ID（企业情况下即为企业ID）。
	OwnerId *string `json:"owner_id,omitempty"`
	// 开票规则标记
	Tag *string `json:"tag,omitempty"`
}

// NewEnterpriseOpenRuleRecordInfo instantiates a new EnterpriseOpenRuleRecordInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseOpenRuleRecordInfo() *EnterpriseOpenRuleRecordInfo {
	this := EnterpriseOpenRuleRecordInfo{}
	return &this
}

// NewEnterpriseOpenRuleRecordInfoWithDefaults instantiates a new EnterpriseOpenRuleRecordInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseOpenRuleRecordInfoWithDefaults() *EnterpriseOpenRuleRecordInfo {
	this := EnterpriseOpenRuleRecordInfo{}
	return &this
}

// GetBillMonthDay returns the BillMonthDay field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetBillMonthDay() int32 {
	if o == nil || IsNil(o.BillMonthDay) {
		var ret int32
		return ret
	}
	return *o.BillMonthDay
}

// GetBillMonthDayOk returns a tuple with the BillMonthDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetBillMonthDayOk() (*int32, bool) {
	if o == nil || IsNil(o.BillMonthDay) {
		return nil, false
	}
	return o.BillMonthDay, true
}

// HasBillMonthDay returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasBillMonthDay() bool {
	if o != nil && !IsNil(o.BillMonthDay) {
		return true
	}

	return false
}

// SetBillMonthDay gets a reference to the given int32 and assigns it to the BillMonthDay field.
func (o *EnterpriseOpenRuleRecordInfo) SetBillMonthDay(v int32) {
	o.BillMonthDay = &v
}

// GetEffectiveStart returns the EffectiveStart field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetEffectiveStart() string {
	if o == nil || IsNil(o.EffectiveStart) {
		var ret string
		return ret
	}
	return *o.EffectiveStart
}

// GetEffectiveStartOk returns a tuple with the EffectiveStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetEffectiveStartOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStart) {
		return nil, false
	}
	return o.EffectiveStart, true
}

// HasEffectiveStart returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasEffectiveStart() bool {
	if o != nil && !IsNil(o.EffectiveStart) {
		return true
	}

	return false
}

// SetEffectiveStart gets a reference to the given string and assigns it to the EffectiveStart field.
func (o *EnterpriseOpenRuleRecordInfo) SetEffectiveStart(v string) {
	o.EffectiveStart = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *EnterpriseOpenRuleRecordInfo) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *EnterpriseOpenRuleRecordInfo) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *EnterpriseOpenRuleRecordInfo) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetInvoiceRuleId returns the InvoiceRuleId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetInvoiceRuleId() string {
	if o == nil || IsNil(o.InvoiceRuleId) {
		var ret string
		return ret
	}
	return *o.InvoiceRuleId
}

// GetInvoiceRuleIdOk returns a tuple with the InvoiceRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetInvoiceRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceRuleId) {
		return nil, false
	}
	return o.InvoiceRuleId, true
}

// HasInvoiceRuleId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasInvoiceRuleId() bool {
	if o != nil && !IsNil(o.InvoiceRuleId) {
		return true
	}

	return false
}

// SetInvoiceRuleId gets a reference to the given string and assigns it to the InvoiceRuleId field.
func (o *EnterpriseOpenRuleRecordInfo) SetInvoiceRuleId(v string) {
	o.InvoiceRuleId = &v
}

// GetInvoiceRuleRecordId returns the InvoiceRuleRecordId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetInvoiceRuleRecordId() string {
	if o == nil || IsNil(o.InvoiceRuleRecordId) {
		var ret string
		return ret
	}
	return *o.InvoiceRuleRecordId
}

// GetInvoiceRuleRecordIdOk returns a tuple with the InvoiceRuleRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetInvoiceRuleRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceRuleRecordId) {
		return nil, false
	}
	return o.InvoiceRuleRecordId, true
}

// HasInvoiceRuleRecordId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasInvoiceRuleRecordId() bool {
	if o != nil && !IsNil(o.InvoiceRuleRecordId) {
		return true
	}

	return false
}

// SetInvoiceRuleRecordId gets a reference to the given string and assigns it to the InvoiceRuleRecordId field.
func (o *EnterpriseOpenRuleRecordInfo) SetInvoiceRuleRecordId(v string) {
	o.InvoiceRuleRecordId = &v
}

// GetInvoiceTitleId returns the InvoiceTitleId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetInvoiceTitleId() string {
	if o == nil || IsNil(o.InvoiceTitleId) {
		var ret string
		return ret
	}
	return *o.InvoiceTitleId
}

// GetInvoiceTitleIdOk returns a tuple with the InvoiceTitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetInvoiceTitleIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceTitleId) {
		return nil, false
	}
	return o.InvoiceTitleId, true
}

// HasInvoiceTitleId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasInvoiceTitleId() bool {
	if o != nil && !IsNil(o.InvoiceTitleId) {
		return true
	}

	return false
}

// SetInvoiceTitleId gets a reference to the given string and assigns it to the InvoiceTitleId field.
func (o *EnterpriseOpenRuleRecordInfo) SetInvoiceTitleId(v string) {
	o.InvoiceTitleId = &v
}

// GetOpenApplyer returns the OpenApplyer field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetOpenApplyer() string {
	if o == nil || IsNil(o.OpenApplyer) {
		var ret string
		return ret
	}
	return *o.OpenApplyer
}

// GetOpenApplyerOk returns a tuple with the OpenApplyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetOpenApplyerOk() (*string, bool) {
	if o == nil || IsNil(o.OpenApplyer) {
		return nil, false
	}
	return o.OpenApplyer, true
}

// HasOpenApplyer returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasOpenApplyer() bool {
	if o != nil && !IsNil(o.OpenApplyer) {
		return true
	}

	return false
}

// SetOpenApplyer gets a reference to the given string and assigns it to the OpenApplyer field.
func (o *EnterpriseOpenRuleRecordInfo) SetOpenApplyer(v string) {
	o.OpenApplyer = &v
}

// GetOpenMode returns the OpenMode field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetOpenMode() string {
	if o == nil || IsNil(o.OpenMode) {
		var ret string
		return ret
	}
	return *o.OpenMode
}

// GetOpenModeOk returns a tuple with the OpenMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetOpenModeOk() (*string, bool) {
	if o == nil || IsNil(o.OpenMode) {
		return nil, false
	}
	return o.OpenMode, true
}

// HasOpenMode returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasOpenMode() bool {
	if o != nil && !IsNil(o.OpenMode) {
		return true
	}

	return false
}

// SetOpenMode gets a reference to the given string and assigns it to the OpenMode field.
func (o *EnterpriseOpenRuleRecordInfo) SetOpenMode(v string) {
	o.OpenMode = &v
}

// GetOpenType returns the OpenType field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetOpenType() string {
	if o == nil || IsNil(o.OpenType) {
		var ret string
		return ret
	}
	return *o.OpenType
}

// GetOpenTypeOk returns a tuple with the OpenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetOpenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OpenType) {
		return nil, false
	}
	return o.OpenType, true
}

// HasOpenType returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasOpenType() bool {
	if o != nil && !IsNil(o.OpenType) {
		return true
	}

	return false
}

// SetOpenType gets a reference to the given string and assigns it to the OpenType field.
func (o *EnterpriseOpenRuleRecordInfo) SetOpenType(v string) {
	o.OpenType = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *EnterpriseOpenRuleRecordInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRecordInfo) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRecordInfo) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRecordInfo) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *EnterpriseOpenRuleRecordInfo) SetTag(v string) {
	o.Tag = &v
}

func (o EnterpriseOpenRuleRecordInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseOpenRuleRecordInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillMonthDay) {
		toSerialize["bill_month_day"] = o.BillMonthDay
	}
	if !IsNil(o.EffectiveStart) {
		toSerialize["effective_start"] = o.EffectiveStart
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmt_create"] = o.GmtCreate
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.InvoiceRuleId) {
		toSerialize["invoice_rule_id"] = o.InvoiceRuleId
	}
	if !IsNil(o.InvoiceRuleRecordId) {
		toSerialize["invoice_rule_record_id"] = o.InvoiceRuleRecordId
	}
	if !IsNil(o.InvoiceTitleId) {
		toSerialize["invoice_title_id"] = o.InvoiceTitleId
	}
	if !IsNil(o.OpenApplyer) {
		toSerialize["open_applyer"] = o.OpenApplyer
	}
	if !IsNil(o.OpenMode) {
		toSerialize["open_mode"] = o.OpenMode
	}
	if !IsNil(o.OpenType) {
		toSerialize["open_type"] = o.OpenType
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableEnterpriseOpenRuleRecordInfo struct {
	value *EnterpriseOpenRuleRecordInfo
	isSet bool
}

func (v NullableEnterpriseOpenRuleRecordInfo) Get() *EnterpriseOpenRuleRecordInfo {
	return v.value
}

func (v *NullableEnterpriseOpenRuleRecordInfo) Set(val *EnterpriseOpenRuleRecordInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseOpenRuleRecordInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseOpenRuleRecordInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseOpenRuleRecordInfo(val *EnterpriseOpenRuleRecordInfo) *NullableEnterpriseOpenRuleRecordInfo {
	return &NullableEnterpriseOpenRuleRecordInfo{value: val, isSet: true}
}

func (v NullableEnterpriseOpenRuleRecordInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseOpenRuleRecordInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
