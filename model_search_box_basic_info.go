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

// checks if the SearchBoxBasicInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBoxBasicInfo{}

// SearchBoxBasicInfo struct for SearchBoxBasicInfo
type SearchBoxBasicInfo struct {
	// 搜索直达配置id
	BoxId *string `json:"box_id,omitempty"`
	// 品牌id
	BrandId *string `json:"brand_id,omitempty"`
	// 搜索直达名称
	Name *string `json:"name,omitempty"`
	// 搜索直达配置状态，INITIAL-初始/ONLINE-已上架/EXPIRE-已失效/OFFLINE-已下架
	Status *string `json:"status,omitempty"`
	// 小程序id
	TargetAppid *string `json:"target_appid,omitempty"`
}

// NewSearchBoxBasicInfo instantiates a new SearchBoxBasicInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBoxBasicInfo() *SearchBoxBasicInfo {
	this := SearchBoxBasicInfo{}
	return &this
}

// NewSearchBoxBasicInfoWithDefaults instantiates a new SearchBoxBasicInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBoxBasicInfoWithDefaults() *SearchBoxBasicInfo {
	this := SearchBoxBasicInfo{}
	return &this
}

// GetBoxId returns the BoxId field value if set, zero value otherwise.
func (o *SearchBoxBasicInfo) GetBoxId() string {
	if o == nil || IsNil(o.BoxId) {
		var ret string
		return ret
	}
	return *o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxBasicInfo) GetBoxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxId) {
		return nil, false
	}
	return o.BoxId, true
}

// HasBoxId returns a boolean if a field has been set.
func (o *SearchBoxBasicInfo) HasBoxId() bool {
	if o != nil && !IsNil(o.BoxId) {
		return true
	}

	return false
}

// SetBoxId gets a reference to the given string and assigns it to the BoxId field.
func (o *SearchBoxBasicInfo) SetBoxId(v string) {
	o.BoxId = &v
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *SearchBoxBasicInfo) GetBrandId() string {
	if o == nil || IsNil(o.BrandId) {
		var ret string
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxBasicInfo) GetBrandIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandId) {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *SearchBoxBasicInfo) HasBrandId() bool {
	if o != nil && !IsNil(o.BrandId) {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given string and assigns it to the BrandId field.
func (o *SearchBoxBasicInfo) SetBrandId(v string) {
	o.BrandId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchBoxBasicInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxBasicInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchBoxBasicInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchBoxBasicInfo) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SearchBoxBasicInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxBasicInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SearchBoxBasicInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SearchBoxBasicInfo) SetStatus(v string) {
	o.Status = &v
}

// GetTargetAppid returns the TargetAppid field value if set, zero value otherwise.
func (o *SearchBoxBasicInfo) GetTargetAppid() string {
	if o == nil || IsNil(o.TargetAppid) {
		var ret string
		return ret
	}
	return *o.TargetAppid
}

// GetTargetAppidOk returns a tuple with the TargetAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxBasicInfo) GetTargetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAppid) {
		return nil, false
	}
	return o.TargetAppid, true
}

// HasTargetAppid returns a boolean if a field has been set.
func (o *SearchBoxBasicInfo) HasTargetAppid() bool {
	if o != nil && !IsNil(o.TargetAppid) {
		return true
	}

	return false
}

// SetTargetAppid gets a reference to the given string and assigns it to the TargetAppid field.
func (o *SearchBoxBasicInfo) SetTargetAppid(v string) {
	o.TargetAppid = &v
}

func (o SearchBoxBasicInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBoxBasicInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxId) {
		toSerialize["box_id"] = o.BoxId
	}
	if !IsNil(o.BrandId) {
		toSerialize["brand_id"] = o.BrandId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TargetAppid) {
		toSerialize["target_appid"] = o.TargetAppid
	}
	return toSerialize, nil
}

type NullableSearchBoxBasicInfo struct {
	value *SearchBoxBasicInfo
	isSet bool
}

func (v NullableSearchBoxBasicInfo) Get() *SearchBoxBasicInfo {
	return v.value
}

func (v *NullableSearchBoxBasicInfo) Set(val *SearchBoxBasicInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBoxBasicInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBoxBasicInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBoxBasicInfo(val *SearchBoxBasicInfo) *NullableSearchBoxBasicInfo {
	return &NullableSearchBoxBasicInfo{value: val, isSet: true}
}

func (v NullableSearchBoxBasicInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBoxBasicInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


