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

// checks if the SearchBaseOrderCreateApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBaseOrderCreateApiRequest{}

// SearchBaseOrderCreateApiRequest struct for SearchBaseOrderCreateApiRequest
type SearchBaseOrderCreateApiRequest struct {
	// 搜索直达类型 BASE：基础信息，只支持基础信息工单提报
	AccessType *string `json:"access_type,omitempty"`
	// 小程序Id
	Appid *string `json:"appid,omitempty"`
	BaseItems *SearchBaseItems `json:"base_items,omitempty"`
	// 服务描述
	Descprise *string `json:"descprise,omitempty"`
	// 是否为草稿态
	IsDraft *bool `json:"is_draft,omitempty"`
	// 申请单id，仅仅驳回或修改是传入
	OrderId *string `json:"order_id,omitempty"`
	// 服务的类型 SP_MINI_APP 小程序 SP_PUBLIC_APP 生活号
	SpecCode *string `json:"spec_code,omitempty"`
}

// NewSearchBaseOrderCreateApiRequest instantiates a new SearchBaseOrderCreateApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBaseOrderCreateApiRequest() *SearchBaseOrderCreateApiRequest {
	this := SearchBaseOrderCreateApiRequest{}
	return &this
}

// NewSearchBaseOrderCreateApiRequestWithDefaults instantiates a new SearchBaseOrderCreateApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBaseOrderCreateApiRequestWithDefaults() *SearchBaseOrderCreateApiRequest {
	this := SearchBaseOrderCreateApiRequest{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SearchBaseOrderCreateApiRequest) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBaseOrderCreateApiRequest) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SearchBaseOrderCreateApiRequest) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *SearchBaseOrderCreateApiRequest) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAppid returns the Appid field value if set, zero value otherwise.
func (o *SearchBaseOrderCreateApiRequest) GetAppid() string {
	if o == nil || IsNil(o.Appid) {
		var ret string
		return ret
	}
	return *o.Appid
}

// GetAppidOk returns a tuple with the Appid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBaseOrderCreateApiRequest) GetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.Appid) {
		return nil, false
	}
	return o.Appid, true
}

// HasAppid returns a boolean if a field has been set.
func (o *SearchBaseOrderCreateApiRequest) HasAppid() bool {
	if o != nil && !IsNil(o.Appid) {
		return true
	}

	return false
}

// SetAppid gets a reference to the given string and assigns it to the Appid field.
func (o *SearchBaseOrderCreateApiRequest) SetAppid(v string) {
	o.Appid = &v
}

// GetBaseItems returns the BaseItems field value if set, zero value otherwise.
func (o *SearchBaseOrderCreateApiRequest) GetBaseItems() SearchBaseItems {
	if o == nil || IsNil(o.BaseItems) {
		var ret SearchBaseItems
		return ret
	}
	return *o.BaseItems
}

// GetBaseItemsOk returns a tuple with the BaseItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBaseOrderCreateApiRequest) GetBaseItemsOk() (*SearchBaseItems, bool) {
	if o == nil || IsNil(o.BaseItems) {
		return nil, false
	}
	return o.BaseItems, true
}

// HasBaseItems returns a boolean if a field has been set.
func (o *SearchBaseOrderCreateApiRequest) HasBaseItems() bool {
	if o != nil && !IsNil(o.BaseItems) {
		return true
	}

	return false
}

// SetBaseItems gets a reference to the given SearchBaseItems and assigns it to the BaseItems field.
func (o *SearchBaseOrderCreateApiRequest) SetBaseItems(v SearchBaseItems) {
	o.BaseItems = &v
}

// GetDescprise returns the Descprise field value if set, zero value otherwise.
func (o *SearchBaseOrderCreateApiRequest) GetDescprise() string {
	if o == nil || IsNil(o.Descprise) {
		var ret string
		return ret
	}
	return *o.Descprise
}

// GetDescpriseOk returns a tuple with the Descprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBaseOrderCreateApiRequest) GetDescpriseOk() (*string, bool) {
	if o == nil || IsNil(o.Descprise) {
		return nil, false
	}
	return o.Descprise, true
}

// HasDescprise returns a boolean if a field has been set.
func (o *SearchBaseOrderCreateApiRequest) HasDescprise() bool {
	if o != nil && !IsNil(o.Descprise) {
		return true
	}

	return false
}

// SetDescprise gets a reference to the given string and assigns it to the Descprise field.
func (o *SearchBaseOrderCreateApiRequest) SetDescprise(v string) {
	o.Descprise = &v
}

// GetIsDraft returns the IsDraft field value if set, zero value otherwise.
func (o *SearchBaseOrderCreateApiRequest) GetIsDraft() bool {
	if o == nil || IsNil(o.IsDraft) {
		var ret bool
		return ret
	}
	return *o.IsDraft
}

// GetIsDraftOk returns a tuple with the IsDraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBaseOrderCreateApiRequest) GetIsDraftOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDraft) {
		return nil, false
	}
	return o.IsDraft, true
}

// HasIsDraft returns a boolean if a field has been set.
func (o *SearchBaseOrderCreateApiRequest) HasIsDraft() bool {
	if o != nil && !IsNil(o.IsDraft) {
		return true
	}

	return false
}

// SetIsDraft gets a reference to the given bool and assigns it to the IsDraft field.
func (o *SearchBaseOrderCreateApiRequest) SetIsDraft(v bool) {
	o.IsDraft = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *SearchBaseOrderCreateApiRequest) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBaseOrderCreateApiRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *SearchBaseOrderCreateApiRequest) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *SearchBaseOrderCreateApiRequest) SetOrderId(v string) {
	o.OrderId = &v
}

// GetSpecCode returns the SpecCode field value if set, zero value otherwise.
func (o *SearchBaseOrderCreateApiRequest) GetSpecCode() string {
	if o == nil || IsNil(o.SpecCode) {
		var ret string
		return ret
	}
	return *o.SpecCode
}

// GetSpecCodeOk returns a tuple with the SpecCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBaseOrderCreateApiRequest) GetSpecCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SpecCode) {
		return nil, false
	}
	return o.SpecCode, true
}

// HasSpecCode returns a boolean if a field has been set.
func (o *SearchBaseOrderCreateApiRequest) HasSpecCode() bool {
	if o != nil && !IsNil(o.SpecCode) {
		return true
	}

	return false
}

// SetSpecCode gets a reference to the given string and assigns it to the SpecCode field.
func (o *SearchBaseOrderCreateApiRequest) SetSpecCode(v string) {
	o.SpecCode = &v
}

func (o SearchBaseOrderCreateApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBaseOrderCreateApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["access_type"] = o.AccessType
	}
	if !IsNil(o.Appid) {
		toSerialize["appid"] = o.Appid
	}
	if !IsNil(o.BaseItems) {
		toSerialize["base_items"] = o.BaseItems
	}
	if !IsNil(o.Descprise) {
		toSerialize["descprise"] = o.Descprise
	}
	if !IsNil(o.IsDraft) {
		toSerialize["is_draft"] = o.IsDraft
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.SpecCode) {
		toSerialize["spec_code"] = o.SpecCode
	}
	return toSerialize, nil
}

type NullableSearchBaseOrderCreateApiRequest struct {
	value *SearchBaseOrderCreateApiRequest
	isSet bool
}

func (v NullableSearchBaseOrderCreateApiRequest) Get() *SearchBaseOrderCreateApiRequest {
	return v.value
}

func (v *NullableSearchBaseOrderCreateApiRequest) Set(val *SearchBaseOrderCreateApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBaseOrderCreateApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBaseOrderCreateApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBaseOrderCreateApiRequest(val *SearchBaseOrderCreateApiRequest) *NullableSearchBaseOrderCreateApiRequest {
	return &NullableSearchBaseOrderCreateApiRequest{value: val, isSet: true}
}

func (v NullableSearchBaseOrderCreateApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBaseOrderCreateApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

