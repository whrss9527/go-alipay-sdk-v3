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

// checks if the AlipayFundEnterprisepayUnsignModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundEnterprisepayUnsignModel{}

// AlipayFundEnterprisepayUnsignModel struct for AlipayFundEnterprisepayUnsignModel
type AlipayFundEnterprisepayUnsignModel struct {
	// 企业账号
	AccountId *string `json:"account_id,omitempty"`
	// 授权协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 业务场景
	BizScene *string `json:"biz_scene,omitempty"`
	// 销售产品码
	ProductCode *string `json:"product_code,omitempty"`
}

// NewAlipayFundEnterprisepayUnsignModel instantiates a new AlipayFundEnterprisepayUnsignModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundEnterprisepayUnsignModel() *AlipayFundEnterprisepayUnsignModel {
	this := AlipayFundEnterprisepayUnsignModel{}
	return &this
}

// NewAlipayFundEnterprisepayUnsignModelWithDefaults instantiates a new AlipayFundEnterprisepayUnsignModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundEnterprisepayUnsignModelWithDefaults() *AlipayFundEnterprisepayUnsignModel {
	this := AlipayFundEnterprisepayUnsignModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayUnsignModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayUnsignModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayUnsignModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayFundEnterprisepayUnsignModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayUnsignModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayUnsignModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayUnsignModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayFundEnterprisepayUnsignModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayUnsignModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayUnsignModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayUnsignModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundEnterprisepayUnsignModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayUnsignModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayUnsignModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayUnsignModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundEnterprisepayUnsignModel) SetProductCode(v string) {
	o.ProductCode = &v
}

func (o AlipayFundEnterprisepayUnsignModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundEnterprisepayUnsignModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	return toSerialize, nil
}

type NullableAlipayFundEnterprisepayUnsignModel struct {
	value *AlipayFundEnterprisepayUnsignModel
	isSet bool
}

func (v NullableAlipayFundEnterprisepayUnsignModel) Get() *AlipayFundEnterprisepayUnsignModel {
	return v.value
}

func (v *NullableAlipayFundEnterprisepayUnsignModel) Set(val *AlipayFundEnterprisepayUnsignModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundEnterprisepayUnsignModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundEnterprisepayUnsignModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundEnterprisepayUnsignModel(val *AlipayFundEnterprisepayUnsignModel) *NullableAlipayFundEnterprisepayUnsignModel {
	return &NullableAlipayFundEnterprisepayUnsignModel{value: val, isSet: true}
}

func (v NullableAlipayFundEnterprisepayUnsignModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundEnterprisepayUnsignModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

