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

// checks if the AlipayFundEnterprisepayGroupAddModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundEnterprisepayGroupAddModel{}

// AlipayFundEnterprisepayGroupAddModel struct for AlipayFundEnterprisepayGroupAddModel
type AlipayFundEnterprisepayGroupAddModel struct {
	// 企业签约账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 平台和企业的三方授权协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 场景码，联系支付宝分配
	BizScene *string `json:"biz_scene,omitempty"`
	FundExtInfo *FundExtInfo `json:"fund_ext_info,omitempty"`
	// 群组名称
	Name *string `json:"name,omitempty"`
	// 外部业务号，外部群组号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 产品码，默认值 ENTERPRISE_PAY
	ProductCode *string `json:"product_code,omitempty"`
}

// NewAlipayFundEnterprisepayGroupAddModel instantiates a new AlipayFundEnterprisepayGroupAddModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundEnterprisepayGroupAddModel() *AlipayFundEnterprisepayGroupAddModel {
	this := AlipayFundEnterprisepayGroupAddModel{}
	return &this
}

// NewAlipayFundEnterprisepayGroupAddModelWithDefaults instantiates a new AlipayFundEnterprisepayGroupAddModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundEnterprisepayGroupAddModelWithDefaults() *AlipayFundEnterprisepayGroupAddModel {
	this := AlipayFundEnterprisepayGroupAddModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupAddModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayFundEnterprisepayGroupAddModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupAddModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayFundEnterprisepayGroupAddModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupAddModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundEnterprisepayGroupAddModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetFundExtInfo returns the FundExtInfo field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupAddModel) GetFundExtInfo() FundExtInfo {
	if o == nil || IsNil(o.FundExtInfo) {
		var ret FundExtInfo
		return ret
	}
	return *o.FundExtInfo
}

// GetFundExtInfoOk returns a tuple with the FundExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) GetFundExtInfoOk() (*FundExtInfo, bool) {
	if o == nil || IsNil(o.FundExtInfo) {
		return nil, false
	}
	return o.FundExtInfo, true
}

// HasFundExtInfo returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) HasFundExtInfo() bool {
	if o != nil && !IsNil(o.FundExtInfo) {
		return true
	}

	return false
}

// SetFundExtInfo gets a reference to the given FundExtInfo and assigns it to the FundExtInfo field.
func (o *AlipayFundEnterprisepayGroupAddModel) SetFundExtInfo(v FundExtInfo) {
	o.FundExtInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupAddModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayFundEnterprisepayGroupAddModel) SetName(v string) {
	o.Name = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupAddModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayFundEnterprisepayGroupAddModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupAddModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupAddModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundEnterprisepayGroupAddModel) SetProductCode(v string) {
	o.ProductCode = &v
}

func (o AlipayFundEnterprisepayGroupAddModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundEnterprisepayGroupAddModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FundExtInfo) {
		toSerialize["fund_ext_info"] = o.FundExtInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	return toSerialize, nil
}

type NullableAlipayFundEnterprisepayGroupAddModel struct {
	value *AlipayFundEnterprisepayGroupAddModel
	isSet bool
}

func (v NullableAlipayFundEnterprisepayGroupAddModel) Get() *AlipayFundEnterprisepayGroupAddModel {
	return v.value
}

func (v *NullableAlipayFundEnterprisepayGroupAddModel) Set(val *AlipayFundEnterprisepayGroupAddModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundEnterprisepayGroupAddModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundEnterprisepayGroupAddModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundEnterprisepayGroupAddModel(val *AlipayFundEnterprisepayGroupAddModel) *NullableAlipayFundEnterprisepayGroupAddModel {
	return &NullableAlipayFundEnterprisepayGroupAddModel{value: val, isSet: true}
}

func (v NullableAlipayFundEnterprisepayGroupAddModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundEnterprisepayGroupAddModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


