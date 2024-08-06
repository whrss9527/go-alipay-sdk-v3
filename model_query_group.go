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

// checks if the QueryGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryGroup{}

// QueryGroup struct for QueryGroup
type QueryGroup struct {
	// 人群中包含人数
	Count *int32 `json:"count,omitempty"`
	// 分组id
	Id *string `json:"id,omitempty"`
	// 分组中的圈人规则
	LabelRule []QueryComplexLabelRule `json:"label_rule,omitempty"`
	// 用户分组名称
	Name *string `json:"name,omitempty"`
}

// NewQueryGroup instantiates a new QueryGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryGroup() *QueryGroup {
	this := QueryGroup{}
	return &this
}

// NewQueryGroupWithDefaults instantiates a new QueryGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryGroupWithDefaults() *QueryGroup {
	this := QueryGroup{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueryGroup) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryGroup) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueryGroup) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *QueryGroup) SetCount(v int32) {
	o.Count = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QueryGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QueryGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QueryGroup) SetId(v string) {
	o.Id = &v
}

// GetLabelRule returns the LabelRule field value if set, zero value otherwise.
func (o *QueryGroup) GetLabelRule() []QueryComplexLabelRule {
	if o == nil || IsNil(o.LabelRule) {
		var ret []QueryComplexLabelRule
		return ret
	}
	return o.LabelRule
}

// GetLabelRuleOk returns a tuple with the LabelRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryGroup) GetLabelRuleOk() ([]QueryComplexLabelRule, bool) {
	if o == nil || IsNil(o.LabelRule) {
		return nil, false
	}
	return o.LabelRule, true
}

// HasLabelRule returns a boolean if a field has been set.
func (o *QueryGroup) HasLabelRule() bool {
	if o != nil && !IsNil(o.LabelRule) {
		return true
	}

	return false
}

// SetLabelRule gets a reference to the given []QueryComplexLabelRule and assigns it to the LabelRule field.
func (o *QueryGroup) SetLabelRule(v []QueryComplexLabelRule) {
	o.LabelRule = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueryGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryGroup) SetName(v string) {
	o.Name = &v
}

func (o QueryGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LabelRule) {
		toSerialize["label_rule"] = o.LabelRule
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableQueryGroup struct {
	value *QueryGroup
	isSet bool
}

func (v NullableQueryGroup) Get() *QueryGroup {
	return v.value
}

func (v *NullableQueryGroup) Set(val *QueryGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryGroup(val *QueryGroup) *NullableQueryGroup {
	return &NullableQueryGroup{value: val, isSet: true}
}

func (v NullableQueryGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


