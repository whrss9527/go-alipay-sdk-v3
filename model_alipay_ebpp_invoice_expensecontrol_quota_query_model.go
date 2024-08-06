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

// checks if the AlipayEbppInvoiceExpensecontrolQuotaQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecontrolQuotaQueryModel{}

// AlipayEbppInvoiceExpensecontrolQuotaQueryModel struct for AlipayEbppInvoiceExpensecontrolQuotaQueryModel
type AlipayEbppInvoiceExpensecontrolQuotaQueryModel struct {
	// 企业共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 额度所属者ID
	OwnerId *string `json:"owner_id,omitempty"`
	// 额度所属者开放ID
	OwnerOpenId *string `json:"owner_open_id,omitempty"`
	// 额度所属者类型
	OwnerType *string `json:"owner_type,omitempty"`
	// 页码
	PageNum *int32 `json:"page_num,omitempty"`
	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`
	// 额度ID列表
	QuotaIdList []string `json:"quota_id_list,omitempty"`
	// 额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// 额度维度ID
	TargetId *string `json:"target_id,omitempty"`
	// 额度维度
	TargetType *string `json:"target_type,omitempty"`
}

// NewAlipayEbppInvoiceExpensecontrolQuotaQueryModel instantiates a new AlipayEbppInvoiceExpensecontrolQuotaQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecontrolQuotaQueryModel() *AlipayEbppInvoiceExpensecontrolQuotaQueryModel {
	this := AlipayEbppInvoiceExpensecontrolQuotaQueryModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecontrolQuotaQueryModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecontrolQuotaQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecontrolQuotaQueryModelWithDefaults() *AlipayEbppInvoiceExpensecontrolQuotaQueryModel {
	this := AlipayEbppInvoiceExpensecontrolQuotaQueryModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerOpenId returns the OwnerOpenId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetOwnerOpenId() string {
	if o == nil || IsNil(o.OwnerOpenId) {
		var ret string
		return ret
	}
	return *o.OwnerOpenId
}

// GetOwnerOpenIdOk returns a tuple with the OwnerOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetOwnerOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerOpenId) {
		return nil, false
	}
	return o.OwnerOpenId, true
}

// HasOwnerOpenId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasOwnerOpenId() bool {
	if o != nil && !IsNil(o.OwnerOpenId) {
		return true
	}

	return false
}

// SetOwnerOpenId gets a reference to the given string and assigns it to the OwnerOpenId field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetOwnerOpenId(v string) {
	o.OwnerOpenId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetQuotaIdList returns the QuotaIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetQuotaIdList() []string {
	if o == nil || IsNil(o.QuotaIdList) {
		var ret []string
		return ret
	}
	return o.QuotaIdList
}

// GetQuotaIdListOk returns a tuple with the QuotaIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetQuotaIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.QuotaIdList) {
		return nil, false
	}
	return o.QuotaIdList, true
}

// HasQuotaIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasQuotaIdList() bool {
	if o != nil && !IsNil(o.QuotaIdList) {
		return true
	}

	return false
}

// SetQuotaIdList gets a reference to the given []string and assigns it to the QuotaIdList field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetQuotaIdList(v []string) {
	o.QuotaIdList = v
}

// GetQuotaType returns the QuotaType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetQuotaType() string {
	if o == nil || IsNil(o.QuotaType) {
		var ret string
		return ret
	}
	return *o.QuotaType
}

// GetQuotaTypeOk returns a tuple with the QuotaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetQuotaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaType) {
		return nil, false
	}
	return o.QuotaType, true
}

// HasQuotaType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasQuotaType() bool {
	if o != nil && !IsNil(o.QuotaType) {
		return true
	}

	return false
}

// SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetQuotaType(v string) {
	o.QuotaType = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) SetTargetType(v string) {
	o.TargetType = &v
}

func (o AlipayEbppInvoiceExpensecontrolQuotaQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecontrolQuotaQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.OwnerOpenId) {
		toSerialize["owner_open_id"] = o.OwnerOpenId
	}
	if !IsNil(o.OwnerType) {
		toSerialize["owner_type"] = o.OwnerType
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.QuotaIdList) {
		toSerialize["quota_id_list"] = o.QuotaIdList
	}
	if !IsNil(o.QuotaType) {
		toSerialize["quota_type"] = o.QuotaType
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel struct {
	value *AlipayEbppInvoiceExpensecontrolQuotaQueryModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel) Get() *AlipayEbppInvoiceExpensecontrolQuotaQueryModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel) Set(val *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel(val *AlipayEbppInvoiceExpensecontrolQuotaQueryModel) *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel {
	return &NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


