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

// checks if the AlipayFundJointaccountListQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundJointaccountListQueryResponseModel{}

// AlipayFundJointaccountListQueryResponseModel struct for AlipayFundJointaccountListQueryResponseModel
type AlipayFundJointaccountListQueryResponseModel struct {
	// 共同账户列表
	AccountList []JointAccountDTO `json:"account_list,omitempty"`
	// 场景
	BizScene *string `json:"biz_scene,omitempty"`
	// 产品码
	ProductCode *string `json:"product_code,omitempty"`
}

// NewAlipayFundJointaccountListQueryResponseModel instantiates a new AlipayFundJointaccountListQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundJointaccountListQueryResponseModel() *AlipayFundJointaccountListQueryResponseModel {
	this := AlipayFundJointaccountListQueryResponseModel{}
	return &this
}

// NewAlipayFundJointaccountListQueryResponseModelWithDefaults instantiates a new AlipayFundJointaccountListQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundJointaccountListQueryResponseModelWithDefaults() *AlipayFundJointaccountListQueryResponseModel {
	this := AlipayFundJointaccountListQueryResponseModel{}
	return &this
}

// GetAccountList returns the AccountList field value if set, zero value otherwise.
func (o *AlipayFundJointaccountListQueryResponseModel) GetAccountList() []JointAccountDTO {
	if o == nil || IsNil(o.AccountList) {
		var ret []JointAccountDTO
		return ret
	}
	return o.AccountList
}

// GetAccountListOk returns a tuple with the AccountList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountListQueryResponseModel) GetAccountListOk() ([]JointAccountDTO, bool) {
	if o == nil || IsNil(o.AccountList) {
		return nil, false
	}
	return o.AccountList, true
}

// HasAccountList returns a boolean if a field has been set.
func (o *AlipayFundJointaccountListQueryResponseModel) HasAccountList() bool {
	if o != nil && !IsNil(o.AccountList) {
		return true
	}

	return false
}

// SetAccountList gets a reference to the given []JointAccountDTO and assigns it to the AccountList field.
func (o *AlipayFundJointaccountListQueryResponseModel) SetAccountList(v []JointAccountDTO) {
	o.AccountList = v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundJointaccountListQueryResponseModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountListQueryResponseModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundJointaccountListQueryResponseModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundJointaccountListQueryResponseModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundJointaccountListQueryResponseModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountListQueryResponseModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundJointaccountListQueryResponseModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundJointaccountListQueryResponseModel) SetProductCode(v string) {
	o.ProductCode = &v
}

func (o AlipayFundJointaccountListQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundJointaccountListQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountList) {
		toSerialize["account_list"] = o.AccountList
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	return toSerialize, nil
}

type NullableAlipayFundJointaccountListQueryResponseModel struct {
	value *AlipayFundJointaccountListQueryResponseModel
	isSet bool
}

func (v NullableAlipayFundJointaccountListQueryResponseModel) Get() *AlipayFundJointaccountListQueryResponseModel {
	return v.value
}

func (v *NullableAlipayFundJointaccountListQueryResponseModel) Set(val *AlipayFundJointaccountListQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountListQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountListQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountListQueryResponseModel(val *AlipayFundJointaccountListQueryResponseModel) *NullableAlipayFundJointaccountListQueryResponseModel {
	return &NullableAlipayFundJointaccountListQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountListQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountListQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

