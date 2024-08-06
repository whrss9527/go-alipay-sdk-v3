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

// checks if the ZhimaCreditPayafteruseCreditagreementTransferModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCreditPayafteruseCreditagreementTransferModel{}

// ZhimaCreditPayafteruseCreditagreementTransferModel struct for ZhimaCreditPayafteruseCreditagreementTransferModel
type ZhimaCreditPayafteruseCreditagreementTransferModel struct {
	// 芝麻外部类目
	CategoryId *string `json:"category_id,omitempty"`
	// 用户的代扣协议号
	DeductAgreementNo *string `json:"deduct_agreement_no,omitempty"`
	// 业务扩展参数，Json格式
	ExtendParams *string `json:"extend_params,omitempty"`
	// 商户外部协议号，不同的支付宝用户需要传递不同的外部单号
	OutAgreementNo *string `json:"out_agreement_no,omitempty"`
	// 商户签约的芝麻产品的产品码
	ProductCode *string `json:"product_code,omitempty"`
	// 芝麻信用服务ID
	ZmServiceId *string `json:"zm_service_id,omitempty"`
}

// NewZhimaCreditPayafteruseCreditagreementTransferModel instantiates a new ZhimaCreditPayafteruseCreditagreementTransferModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCreditPayafteruseCreditagreementTransferModel() *ZhimaCreditPayafteruseCreditagreementTransferModel {
	this := ZhimaCreditPayafteruseCreditagreementTransferModel{}
	return &this
}

// NewZhimaCreditPayafteruseCreditagreementTransferModelWithDefaults instantiates a new ZhimaCreditPayafteruseCreditagreementTransferModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCreditPayafteruseCreditagreementTransferModelWithDefaults() *ZhimaCreditPayafteruseCreditagreementTransferModel {
	this := ZhimaCreditPayafteruseCreditagreementTransferModel{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetDeductAgreementNo returns the DeductAgreementNo field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetDeductAgreementNo() string {
	if o == nil || IsNil(o.DeductAgreementNo) {
		var ret string
		return ret
	}
	return *o.DeductAgreementNo
}

// GetDeductAgreementNoOk returns a tuple with the DeductAgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetDeductAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.DeductAgreementNo) {
		return nil, false
	}
	return o.DeductAgreementNo, true
}

// HasDeductAgreementNo returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) HasDeductAgreementNo() bool {
	if o != nil && !IsNil(o.DeductAgreementNo) {
		return true
	}

	return false
}

// SetDeductAgreementNo gets a reference to the given string and assigns it to the DeductAgreementNo field.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) SetDeductAgreementNo(v string) {
	o.DeductAgreementNo = &v
}

// GetExtendParams returns the ExtendParams field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetExtendParams() string {
	if o == nil || IsNil(o.ExtendParams) {
		var ret string
		return ret
	}
	return *o.ExtendParams
}

// GetExtendParamsOk returns a tuple with the ExtendParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetExtendParamsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendParams) {
		return nil, false
	}
	return o.ExtendParams, true
}

// HasExtendParams returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) HasExtendParams() bool {
	if o != nil && !IsNil(o.ExtendParams) {
		return true
	}

	return false
}

// SetExtendParams gets a reference to the given string and assigns it to the ExtendParams field.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) SetExtendParams(v string) {
	o.ExtendParams = &v
}

// GetOutAgreementNo returns the OutAgreementNo field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetOutAgreementNo() string {
	if o == nil || IsNil(o.OutAgreementNo) {
		var ret string
		return ret
	}
	return *o.OutAgreementNo
}

// GetOutAgreementNoOk returns a tuple with the OutAgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetOutAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutAgreementNo) {
		return nil, false
	}
	return o.OutAgreementNo, true
}

// HasOutAgreementNo returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) HasOutAgreementNo() bool {
	if o != nil && !IsNil(o.OutAgreementNo) {
		return true
	}

	return false
}

// SetOutAgreementNo gets a reference to the given string and assigns it to the OutAgreementNo field.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) SetOutAgreementNo(v string) {
	o.OutAgreementNo = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetZmServiceId returns the ZmServiceId field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetZmServiceId() string {
	if o == nil || IsNil(o.ZmServiceId) {
		var ret string
		return ret
	}
	return *o.ZmServiceId
}

// GetZmServiceIdOk returns a tuple with the ZmServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) GetZmServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ZmServiceId) {
		return nil, false
	}
	return o.ZmServiceId, true
}

// HasZmServiceId returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) HasZmServiceId() bool {
	if o != nil && !IsNil(o.ZmServiceId) {
		return true
	}

	return false
}

// SetZmServiceId gets a reference to the given string and assigns it to the ZmServiceId field.
func (o *ZhimaCreditPayafteruseCreditagreementTransferModel) SetZmServiceId(v string) {
	o.ZmServiceId = &v
}

func (o ZhimaCreditPayafteruseCreditagreementTransferModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCreditPayafteruseCreditagreementTransferModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.DeductAgreementNo) {
		toSerialize["deduct_agreement_no"] = o.DeductAgreementNo
	}
	if !IsNil(o.ExtendParams) {
		toSerialize["extend_params"] = o.ExtendParams
	}
	if !IsNil(o.OutAgreementNo) {
		toSerialize["out_agreement_no"] = o.OutAgreementNo
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.ZmServiceId) {
		toSerialize["zm_service_id"] = o.ZmServiceId
	}
	return toSerialize, nil
}

type NullableZhimaCreditPayafteruseCreditagreementTransferModel struct {
	value *ZhimaCreditPayafteruseCreditagreementTransferModel
	isSet bool
}

func (v NullableZhimaCreditPayafteruseCreditagreementTransferModel) Get() *ZhimaCreditPayafteruseCreditagreementTransferModel {
	return v.value
}

func (v *NullableZhimaCreditPayafteruseCreditagreementTransferModel) Set(val *ZhimaCreditPayafteruseCreditagreementTransferModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPayafteruseCreditagreementTransferModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPayafteruseCreditagreementTransferModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPayafteruseCreditagreementTransferModel(val *ZhimaCreditPayafteruseCreditagreementTransferModel) *NullableZhimaCreditPayafteruseCreditagreementTransferModel {
	return &NullableZhimaCreditPayafteruseCreditagreementTransferModel{value: val, isSet: true}
}

func (v NullableZhimaCreditPayafteruseCreditagreementTransferModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPayafteruseCreditagreementTransferModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

