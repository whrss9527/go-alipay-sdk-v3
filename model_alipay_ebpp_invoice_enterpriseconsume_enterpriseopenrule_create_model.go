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

// checks if the AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel{}

// AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel struct for AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel
type AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel struct {
	// 共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 账期结束日期
	BillMonthDay *int32 `json:"bill_month_day,omitempty"`
	// 企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 开票规则名称
	InvoiceRuleName *string `json:"invoice_rule_name,omitempty"`
	// 发票抬头
	InvoiceTitleId *string `json:"invoice_title_id,omitempty"`
	// 开票模式
	OpenMode *string `json:"open_mode,omitempty"`
	// 收件人地址
	ReceiveAddress *string `json:"receive_address,omitempty"`
	// 收件人姓名
	ReceiveName *string `json:"receive_name,omitempty"`
	// 收件人手机号
	ReceivePhone *string `json:"receive_phone,omitempty"`
	// 销方类型
	SellerType *string `json:"seller_type,omitempty"`
	// 开票规则标记
	Tag *string `json:"tag,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModelWithDefaults() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBillMonthDay returns the BillMonthDay field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetBillMonthDay() int32 {
	if o == nil || IsNil(o.BillMonthDay) {
		var ret int32
		return ret
	}
	return *o.BillMonthDay
}

// GetBillMonthDayOk returns a tuple with the BillMonthDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetBillMonthDayOk() (*int32, bool) {
	if o == nil || IsNil(o.BillMonthDay) {
		return nil, false
	}
	return o.BillMonthDay, true
}

// HasBillMonthDay returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasBillMonthDay() bool {
	if o != nil && !IsNil(o.BillMonthDay) {
		return true
	}

	return false
}

// SetBillMonthDay gets a reference to the given int32 and assigns it to the BillMonthDay field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetBillMonthDay(v int32) {
	o.BillMonthDay = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetInvoiceRuleName returns the InvoiceRuleName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetInvoiceRuleName() string {
	if o == nil || IsNil(o.InvoiceRuleName) {
		var ret string
		return ret
	}
	return *o.InvoiceRuleName
}

// GetInvoiceRuleNameOk returns a tuple with the InvoiceRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetInvoiceRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceRuleName) {
		return nil, false
	}
	return o.InvoiceRuleName, true
}

// HasInvoiceRuleName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasInvoiceRuleName() bool {
	if o != nil && !IsNil(o.InvoiceRuleName) {
		return true
	}

	return false
}

// SetInvoiceRuleName gets a reference to the given string and assigns it to the InvoiceRuleName field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetInvoiceRuleName(v string) {
	o.InvoiceRuleName = &v
}

// GetInvoiceTitleId returns the InvoiceTitleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetInvoiceTitleId() string {
	if o == nil || IsNil(o.InvoiceTitleId) {
		var ret string
		return ret
	}
	return *o.InvoiceTitleId
}

// GetInvoiceTitleIdOk returns a tuple with the InvoiceTitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetInvoiceTitleIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceTitleId) {
		return nil, false
	}
	return o.InvoiceTitleId, true
}

// HasInvoiceTitleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasInvoiceTitleId() bool {
	if o != nil && !IsNil(o.InvoiceTitleId) {
		return true
	}

	return false
}

// SetInvoiceTitleId gets a reference to the given string and assigns it to the InvoiceTitleId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetInvoiceTitleId(v string) {
	o.InvoiceTitleId = &v
}

// GetOpenMode returns the OpenMode field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetOpenMode() string {
	if o == nil || IsNil(o.OpenMode) {
		var ret string
		return ret
	}
	return *o.OpenMode
}

// GetOpenModeOk returns a tuple with the OpenMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetOpenModeOk() (*string, bool) {
	if o == nil || IsNil(o.OpenMode) {
		return nil, false
	}
	return o.OpenMode, true
}

// HasOpenMode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasOpenMode() bool {
	if o != nil && !IsNil(o.OpenMode) {
		return true
	}

	return false
}

// SetOpenMode gets a reference to the given string and assigns it to the OpenMode field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetOpenMode(v string) {
	o.OpenMode = &v
}

// GetReceiveAddress returns the ReceiveAddress field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetReceiveAddress() string {
	if o == nil || IsNil(o.ReceiveAddress) {
		var ret string
		return ret
	}
	return *o.ReceiveAddress
}

// GetReceiveAddressOk returns a tuple with the ReceiveAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetReceiveAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveAddress) {
		return nil, false
	}
	return o.ReceiveAddress, true
}

// HasReceiveAddress returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasReceiveAddress() bool {
	if o != nil && !IsNil(o.ReceiveAddress) {
		return true
	}

	return false
}

// SetReceiveAddress gets a reference to the given string and assigns it to the ReceiveAddress field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetReceiveAddress(v string) {
	o.ReceiveAddress = &v
}

// GetReceiveName returns the ReceiveName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetReceiveName() string {
	if o == nil || IsNil(o.ReceiveName) {
		var ret string
		return ret
	}
	return *o.ReceiveName
}

// GetReceiveNameOk returns a tuple with the ReceiveName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetReceiveNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveName) {
		return nil, false
	}
	return o.ReceiveName, true
}

// HasReceiveName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasReceiveName() bool {
	if o != nil && !IsNil(o.ReceiveName) {
		return true
	}

	return false
}

// SetReceiveName gets a reference to the given string and assigns it to the ReceiveName field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetReceiveName(v string) {
	o.ReceiveName = &v
}

// GetReceivePhone returns the ReceivePhone field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetReceivePhone() string {
	if o == nil || IsNil(o.ReceivePhone) {
		var ret string
		return ret
	}
	return *o.ReceivePhone
}

// GetReceivePhoneOk returns a tuple with the ReceivePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetReceivePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivePhone) {
		return nil, false
	}
	return o.ReceivePhone, true
}

// HasReceivePhone returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasReceivePhone() bool {
	if o != nil && !IsNil(o.ReceivePhone) {
		return true
	}

	return false
}

// SetReceivePhone gets a reference to the given string and assigns it to the ReceivePhone field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetReceivePhone(v string) {
	o.ReceivePhone = &v
}

// GetSellerType returns the SellerType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetSellerType() string {
	if o == nil || IsNil(o.SellerType) {
		var ret string
		return ret
	}
	return *o.SellerType
}

// GetSellerTypeOk returns a tuple with the SellerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetSellerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SellerType) {
		return nil, false
	}
	return o.SellerType, true
}

// HasSellerType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasSellerType() bool {
	if o != nil && !IsNil(o.SellerType) {
		return true
	}

	return false
}

// SetSellerType gets a reference to the given string and assigns it to the SellerType field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetSellerType(v string) {
	o.SellerType = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) SetTag(v string) {
	o.Tag = &v
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BillMonthDay) {
		toSerialize["bill_month_day"] = o.BillMonthDay
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.InvoiceRuleName) {
		toSerialize["invoice_rule_name"] = o.InvoiceRuleName
	}
	if !IsNil(o.InvoiceTitleId) {
		toSerialize["invoice_title_id"] = o.InvoiceTitleId
	}
	if !IsNil(o.OpenMode) {
		toSerialize["open_mode"] = o.OpenMode
	}
	if !IsNil(o.ReceiveAddress) {
		toSerialize["receive_address"] = o.ReceiveAddress
	}
	if !IsNil(o.ReceiveName) {
		toSerialize["receive_name"] = o.ReceiveName
	}
	if !IsNil(o.ReceivePhone) {
		toSerialize["receive_phone"] = o.ReceivePhone
	}
	if !IsNil(o.SellerType) {
		toSerialize["seller_type"] = o.SellerType
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel struct {
	value *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) Get() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) Set(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel {
	return &NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
