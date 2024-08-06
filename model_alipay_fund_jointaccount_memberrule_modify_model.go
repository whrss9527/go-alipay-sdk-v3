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

// checks if the AlipayFundJointaccountMemberruleModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundJointaccountMemberruleModifyModel{}

// AlipayFundJointaccountMemberruleModifyModel struct for AlipayFundJointaccountMemberruleModifyModel
type AlipayFundJointaccountMemberruleModifyModel struct {
	// 账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 业务场景
	BizScene   *string                 `json:"biz_scene,omitempty"`
	MemberList *JointAccountMemberList `json:"member_list,omitempty"`
	// 产品码
	ProductCode *string `json:"product_code,omitempty"`
}

// NewAlipayFundJointaccountMemberruleModifyModel instantiates a new AlipayFundJointaccountMemberruleModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundJointaccountMemberruleModifyModel() *AlipayFundJointaccountMemberruleModifyModel {
	this := AlipayFundJointaccountMemberruleModifyModel{}
	return &this
}

// NewAlipayFundJointaccountMemberruleModifyModelWithDefaults instantiates a new AlipayFundJointaccountMemberruleModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundJointaccountMemberruleModifyModelWithDefaults() *AlipayFundJointaccountMemberruleModifyModel {
	this := AlipayFundJointaccountMemberruleModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayFundJointaccountMemberruleModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayFundJointaccountMemberruleModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundJointaccountMemberruleModifyModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetMemberList returns the MemberList field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetMemberList() JointAccountMemberList {
	if o == nil || IsNil(o.MemberList) {
		var ret JointAccountMemberList
		return ret
	}
	return *o.MemberList
}

// GetMemberListOk returns a tuple with the MemberList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetMemberListOk() (*JointAccountMemberList, bool) {
	if o == nil || IsNil(o.MemberList) {
		return nil, false
	}
	return o.MemberList, true
}

// HasMemberList returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) HasMemberList() bool {
	if o != nil && !IsNil(o.MemberList) {
		return true
	}

	return false
}

// SetMemberList gets a reference to the given JointAccountMemberList and assigns it to the MemberList field.
func (o *AlipayFundJointaccountMemberruleModifyModel) SetMemberList(v JointAccountMemberList) {
	o.MemberList = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberruleModifyModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundJointaccountMemberruleModifyModel) SetProductCode(v string) {
	o.ProductCode = &v
}

func (o AlipayFundJointaccountMemberruleModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundJointaccountMemberruleModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.MemberList) {
		toSerialize["member_list"] = o.MemberList
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	return toSerialize, nil
}

type NullableAlipayFundJointaccountMemberruleModifyModel struct {
	value *AlipayFundJointaccountMemberruleModifyModel
	isSet bool
}

func (v NullableAlipayFundJointaccountMemberruleModifyModel) Get() *AlipayFundJointaccountMemberruleModifyModel {
	return v.value
}

func (v *NullableAlipayFundJointaccountMemberruleModifyModel) Set(val *AlipayFundJointaccountMemberruleModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountMemberruleModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountMemberruleModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountMemberruleModifyModel(val *AlipayFundJointaccountMemberruleModifyModel) *NullableAlipayFundJointaccountMemberruleModifyModel {
	return &NullableAlipayFundJointaccountMemberruleModifyModel{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountMemberruleModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountMemberruleModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
