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

// checks if the SearchOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchOrderRequest{}

// SearchOrderRequest struct for SearchOrderRequest
type SearchOrderRequest struct {
	// 申请单id
	ApplyId *string `json:"apply_id,omitempty"`
	// 申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达
	ApplyType *string `json:"apply_type,omitempty"`
	// 品牌的模板id ONE_WITH_TWO：一拖二、DEFAULT：一拖四
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// 服务code
	ServiceCode *string `json:"service_code,omitempty"`
	// 服务的类型默认使用小程序 SP_MINI_APP 小程序 SP_PUBLIC_APP 生活号
	SpecCode *string `json:"spec_code,omitempty"`
}

// NewSearchOrderRequest instantiates a new SearchOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchOrderRequest() *SearchOrderRequest {
	this := SearchOrderRequest{}
	return &this
}

// NewSearchOrderRequestWithDefaults instantiates a new SearchOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchOrderRequestWithDefaults() *SearchOrderRequest {
	this := SearchOrderRequest{}
	return &this
}

// GetApplyId returns the ApplyId field value if set, zero value otherwise.
func (o *SearchOrderRequest) GetApplyId() string {
	if o == nil || IsNil(o.ApplyId) {
		var ret string
		return ret
	}
	return *o.ApplyId
}

// GetApplyIdOk returns a tuple with the ApplyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderRequest) GetApplyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyId) {
		return nil, false
	}
	return o.ApplyId, true
}

// HasApplyId returns a boolean if a field has been set.
func (o *SearchOrderRequest) HasApplyId() bool {
	if o != nil && !IsNil(o.ApplyId) {
		return true
	}

	return false
}

// SetApplyId gets a reference to the given string and assigns it to the ApplyId field.
func (o *SearchOrderRequest) SetApplyId(v string) {
	o.ApplyId = &v
}

// GetApplyType returns the ApplyType field value if set, zero value otherwise.
func (o *SearchOrderRequest) GetApplyType() string {
	if o == nil || IsNil(o.ApplyType) {
		var ret string
		return ret
	}
	return *o.ApplyType
}

// GetApplyTypeOk returns a tuple with the ApplyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderRequest) GetApplyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyType) {
		return nil, false
	}
	return o.ApplyType, true
}

// HasApplyType returns a boolean if a field has been set.
func (o *SearchOrderRequest) HasApplyType() bool {
	if o != nil && !IsNil(o.ApplyType) {
		return true
	}

	return false
}

// SetApplyType gets a reference to the given string and assigns it to the ApplyType field.
func (o *SearchOrderRequest) SetApplyType(v string) {
	o.ApplyType = &v
}

// GetBrandTemplateId returns the BrandTemplateId field value if set, zero value otherwise.
func (o *SearchOrderRequest) GetBrandTemplateId() string {
	if o == nil || IsNil(o.BrandTemplateId) {
		var ret string
		return ret
	}
	return *o.BrandTemplateId
}

// GetBrandTemplateIdOk returns a tuple with the BrandTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderRequest) GetBrandTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandTemplateId) {
		return nil, false
	}
	return o.BrandTemplateId, true
}

// HasBrandTemplateId returns a boolean if a field has been set.
func (o *SearchOrderRequest) HasBrandTemplateId() bool {
	if o != nil && !IsNil(o.BrandTemplateId) {
		return true
	}

	return false
}

// SetBrandTemplateId gets a reference to the given string and assigns it to the BrandTemplateId field.
func (o *SearchOrderRequest) SetBrandTemplateId(v string) {
	o.BrandTemplateId = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *SearchOrderRequest) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderRequest) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *SearchOrderRequest) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *SearchOrderRequest) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetSpecCode returns the SpecCode field value if set, zero value otherwise.
func (o *SearchOrderRequest) GetSpecCode() string {
	if o == nil || IsNil(o.SpecCode) {
		var ret string
		return ret
	}
	return *o.SpecCode
}

// GetSpecCodeOk returns a tuple with the SpecCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderRequest) GetSpecCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SpecCode) {
		return nil, false
	}
	return o.SpecCode, true
}

// HasSpecCode returns a boolean if a field has been set.
func (o *SearchOrderRequest) HasSpecCode() bool {
	if o != nil && !IsNil(o.SpecCode) {
		return true
	}

	return false
}

// SetSpecCode gets a reference to the given string and assigns it to the SpecCode field.
func (o *SearchOrderRequest) SetSpecCode(v string) {
	o.SpecCode = &v
}

func (o SearchOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyId) {
		toSerialize["apply_id"] = o.ApplyId
	}
	if !IsNil(o.ApplyType) {
		toSerialize["apply_type"] = o.ApplyType
	}
	if !IsNil(o.BrandTemplateId) {
		toSerialize["brand_template_id"] = o.BrandTemplateId
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["service_code"] = o.ServiceCode
	}
	if !IsNil(o.SpecCode) {
		toSerialize["spec_code"] = o.SpecCode
	}
	return toSerialize, nil
}

type NullableSearchOrderRequest struct {
	value *SearchOrderRequest
	isSet bool
}

func (v NullableSearchOrderRequest) Get() *SearchOrderRequest {
	return v.value
}

func (v *NullableSearchOrderRequest) Set(val *SearchOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchOrderRequest(val *SearchOrderRequest) *NullableSearchOrderRequest {
	return &NullableSearchOrderRequest{value: val, isSet: true}
}

func (v NullableSearchOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


