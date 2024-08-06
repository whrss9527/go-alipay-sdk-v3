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

// checks if the EnterpriseOpenRuleRelationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterpriseOpenRuleRelationInfo{}

// EnterpriseOpenRuleRelationInfo struct for EnterpriseOpenRuleRelationInfo
type EnterpriseOpenRuleRelationInfo struct {
	// 企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// 所有者ID（企业情况下即为企业ID）。
	OwnerId *string `json:"owner_id,omitempty"`
	// 费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
}

// NewEnterpriseOpenRuleRelationInfo instantiates a new EnterpriseOpenRuleRelationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseOpenRuleRelationInfo() *EnterpriseOpenRuleRelationInfo {
	this := EnterpriseOpenRuleRelationInfo{}
	return &this
}

// NewEnterpriseOpenRuleRelationInfoWithDefaults instantiates a new EnterpriseOpenRuleRelationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseOpenRuleRelationInfoWithDefaults() *EnterpriseOpenRuleRelationInfo {
	this := EnterpriseOpenRuleRelationInfo{}
	return &this
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRelationInfo) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRelationInfo) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRelationInfo) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *EnterpriseOpenRuleRelationInfo) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRelationInfo) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRelationInfo) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRelationInfo) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *EnterpriseOpenRuleRelationInfo) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRelationInfo) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRelationInfo) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRelationInfo) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *EnterpriseOpenRuleRelationInfo) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRelationInfo) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRelationInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRelationInfo) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *EnterpriseOpenRuleRelationInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetStandardId returns the StandardId field value if set, zero value otherwise.
func (o *EnterpriseOpenRuleRelationInfo) GetStandardId() string {
	if o == nil || IsNil(o.StandardId) {
		var ret string
		return ret
	}
	return *o.StandardId
}

// GetStandardIdOk returns a tuple with the StandardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseOpenRuleRelationInfo) GetStandardIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardId) {
		return nil, false
	}
	return o.StandardId, true
}

// HasStandardId returns a boolean if a field has been set.
func (o *EnterpriseOpenRuleRelationInfo) HasStandardId() bool {
	if o != nil && !IsNil(o.StandardId) {
		return true
	}

	return false
}

// SetStandardId gets a reference to the given string and assigns it to the StandardId field.
func (o *EnterpriseOpenRuleRelationInfo) SetStandardId(v string) {
	o.StandardId = &v
}

func (o EnterpriseOpenRuleRelationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseOpenRuleRelationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmt_create"] = o.GmtCreate
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.StandardId) {
		toSerialize["standard_id"] = o.StandardId
	}
	return toSerialize, nil
}

type NullableEnterpriseOpenRuleRelationInfo struct {
	value *EnterpriseOpenRuleRelationInfo
	isSet bool
}

func (v NullableEnterpriseOpenRuleRelationInfo) Get() *EnterpriseOpenRuleRelationInfo {
	return v.value
}

func (v *NullableEnterpriseOpenRuleRelationInfo) Set(val *EnterpriseOpenRuleRelationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseOpenRuleRelationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseOpenRuleRelationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseOpenRuleRelationInfo(val *EnterpriseOpenRuleRelationInfo) *NullableEnterpriseOpenRuleRelationInfo {
	return &NullableEnterpriseOpenRuleRelationInfo{value: val, isSet: true}
}

func (v NullableEnterpriseOpenRuleRelationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseOpenRuleRelationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
