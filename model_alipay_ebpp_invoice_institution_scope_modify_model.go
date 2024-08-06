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

// checks if the AlipayEbppInvoiceInstitutionScopeModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceInstitutionScopeModifyModel{}

// AlipayEbppInvoiceInstitutionScopeModifyModel struct for AlipayEbppInvoiceInstitutionScopeModifyModel
type AlipayEbppInvoiceInstitutionScopeModifyModel struct {
	// 企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// 制度适用范围类型
	AdapterType *string `json:"adapter_type,omitempty"`
	// 增加适配id列表
	AddOwnerIdList []string `json:"add_owner_id_list,omitempty"`
	// 增加适配开放id列表
	AddOwnerOpenIdList []string `json:"add_owner_open_id_list,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// 归属类型
	OwnerType *string `json:"owner_type,omitempty"`
	// 删除适配id列表
	RemoveOwnerIdList []string `json:"remove_owner_id_list,omitempty"`
	// 删除适配开放id列表
	RemoveOwnerOpenIdList []string `json:"remove_owner_open_id_list,omitempty"`
}

// NewAlipayEbppInvoiceInstitutionScopeModifyModel instantiates a new AlipayEbppInvoiceInstitutionScopeModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceInstitutionScopeModifyModel() *AlipayEbppInvoiceInstitutionScopeModifyModel {
	this := AlipayEbppInvoiceInstitutionScopeModifyModel{}
	return &this
}

// NewAlipayEbppInvoiceInstitutionScopeModifyModelWithDefaults instantiates a new AlipayEbppInvoiceInstitutionScopeModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceInstitutionScopeModifyModelWithDefaults() *AlipayEbppInvoiceInstitutionScopeModifyModel {
	this := AlipayEbppInvoiceInstitutionScopeModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAdapterType returns the AdapterType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAdapterType() string {
	if o == nil || IsNil(o.AdapterType) {
		var ret string
		return ret
	}
	return *o.AdapterType
}

// GetAdapterTypeOk returns a tuple with the AdapterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAdapterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AdapterType) {
		return nil, false
	}
	return o.AdapterType, true
}

// HasAdapterType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasAdapterType() bool {
	if o != nil && !IsNil(o.AdapterType) {
		return true
	}

	return false
}

// SetAdapterType gets a reference to the given string and assigns it to the AdapterType field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetAdapterType(v string) {
	o.AdapterType = &v
}

// GetAddOwnerIdList returns the AddOwnerIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAddOwnerIdList() []string {
	if o == nil || IsNil(o.AddOwnerIdList) {
		var ret []string
		return ret
	}
	return o.AddOwnerIdList
}

// GetAddOwnerIdListOk returns a tuple with the AddOwnerIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAddOwnerIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.AddOwnerIdList) {
		return nil, false
	}
	return o.AddOwnerIdList, true
}

// HasAddOwnerIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasAddOwnerIdList() bool {
	if o != nil && !IsNil(o.AddOwnerIdList) {
		return true
	}

	return false
}

// SetAddOwnerIdList gets a reference to the given []string and assigns it to the AddOwnerIdList field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetAddOwnerIdList(v []string) {
	o.AddOwnerIdList = v
}

// GetAddOwnerOpenIdList returns the AddOwnerOpenIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAddOwnerOpenIdList() []string {
	if o == nil || IsNil(o.AddOwnerOpenIdList) {
		var ret []string
		return ret
	}
	return o.AddOwnerOpenIdList
}

// GetAddOwnerOpenIdListOk returns a tuple with the AddOwnerOpenIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAddOwnerOpenIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.AddOwnerOpenIdList) {
		return nil, false
	}
	return o.AddOwnerOpenIdList, true
}

// HasAddOwnerOpenIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasAddOwnerOpenIdList() bool {
	if o != nil && !IsNil(o.AddOwnerOpenIdList) {
		return true
	}

	return false
}

// SetAddOwnerOpenIdList gets a reference to the given []string and assigns it to the AddOwnerOpenIdList field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetAddOwnerOpenIdList(v []string) {
	o.AddOwnerOpenIdList = v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetInstitutionId() string {
	if o == nil || IsNil(o.InstitutionId) {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetInstitutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionId) {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasInstitutionId() bool {
	if o != nil && !IsNil(o.InstitutionId) {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetRemoveOwnerIdList returns the RemoveOwnerIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetRemoveOwnerIdList() []string {
	if o == nil || IsNil(o.RemoveOwnerIdList) {
		var ret []string
		return ret
	}
	return o.RemoveOwnerIdList
}

// GetRemoveOwnerIdListOk returns a tuple with the RemoveOwnerIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetRemoveOwnerIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoveOwnerIdList) {
		return nil, false
	}
	return o.RemoveOwnerIdList, true
}

// HasRemoveOwnerIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasRemoveOwnerIdList() bool {
	if o != nil && !IsNil(o.RemoveOwnerIdList) {
		return true
	}

	return false
}

// SetRemoveOwnerIdList gets a reference to the given []string and assigns it to the RemoveOwnerIdList field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetRemoveOwnerIdList(v []string) {
	o.RemoveOwnerIdList = v
}

// GetRemoveOwnerOpenIdList returns the RemoveOwnerOpenIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetRemoveOwnerOpenIdList() []string {
	if o == nil || IsNil(o.RemoveOwnerOpenIdList) {
		var ret []string
		return ret
	}
	return o.RemoveOwnerOpenIdList
}

// GetRemoveOwnerOpenIdListOk returns a tuple with the RemoveOwnerOpenIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) GetRemoveOwnerOpenIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoveOwnerOpenIdList) {
		return nil, false
	}
	return o.RemoveOwnerOpenIdList, true
}

// HasRemoveOwnerOpenIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) HasRemoveOwnerOpenIdList() bool {
	if o != nil && !IsNil(o.RemoveOwnerOpenIdList) {
		return true
	}

	return false
}

// SetRemoveOwnerOpenIdList gets a reference to the given []string and assigns it to the RemoveOwnerOpenIdList field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyModel) SetRemoveOwnerOpenIdList(v []string) {
	o.RemoveOwnerOpenIdList = v
}

func (o AlipayEbppInvoiceInstitutionScopeModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceInstitutionScopeModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AdapterType) {
		toSerialize["adapter_type"] = o.AdapterType
	}
	if !IsNil(o.AddOwnerIdList) {
		toSerialize["add_owner_id_list"] = o.AddOwnerIdList
	}
	if !IsNil(o.AddOwnerOpenIdList) {
		toSerialize["add_owner_open_id_list"] = o.AddOwnerOpenIdList
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.InstitutionId) {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if !IsNil(o.OwnerType) {
		toSerialize["owner_type"] = o.OwnerType
	}
	if !IsNil(o.RemoveOwnerIdList) {
		toSerialize["remove_owner_id_list"] = o.RemoveOwnerIdList
	}
	if !IsNil(o.RemoveOwnerOpenIdList) {
		toSerialize["remove_owner_open_id_list"] = o.RemoveOwnerOpenIdList
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceInstitutionScopeModifyModel struct {
	value *AlipayEbppInvoiceInstitutionScopeModifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceInstitutionScopeModifyModel) Get() *AlipayEbppInvoiceInstitutionScopeModifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceInstitutionScopeModifyModel) Set(val *AlipayEbppInvoiceInstitutionScopeModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceInstitutionScopeModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceInstitutionScopeModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceInstitutionScopeModifyModel(val *AlipayEbppInvoiceInstitutionScopeModifyModel) *NullableAlipayEbppInvoiceInstitutionScopeModifyModel {
	return &NullableAlipayEbppInvoiceInstitutionScopeModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceInstitutionScopeModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceInstitutionScopeModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
