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

// checks if the AlipayIserviceCcmRoleCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmRoleCreateModel{}

// AlipayIserviceCcmRoleCreateModel struct for AlipayIserviceCcmRoleCreateModel
type AlipayIserviceCcmRoleCreateModel struct {
	// 部门id（即租户实例ID、数据权限ID）
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// 创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// 角色描述信息
	Description *string `json:"description,omitempty"`
	// 角色关联额功能点id
	FunctionIds []string `json:"function_ids,omitempty"`
	// 角色名
	Name *string `json:"name,omitempty"`
}

// NewAlipayIserviceCcmRoleCreateModel instantiates a new AlipayIserviceCcmRoleCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmRoleCreateModel() *AlipayIserviceCcmRoleCreateModel {
	this := AlipayIserviceCcmRoleCreateModel{}
	return &this
}

// NewAlipayIserviceCcmRoleCreateModelWithDefaults instantiates a new AlipayIserviceCcmRoleCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmRoleCreateModelWithDefaults() *AlipayIserviceCcmRoleCreateModel {
	this := AlipayIserviceCcmRoleCreateModel{}
	return &this
}

// GetCcsInstanceId returns the CcsInstanceId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRoleCreateModel) GetCcsInstanceId() string {
	if o == nil || IsNil(o.CcsInstanceId) {
		var ret string
		return ret
	}
	return *o.CcsInstanceId
}

// GetCcsInstanceIdOk returns a tuple with the CcsInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRoleCreateModel) GetCcsInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CcsInstanceId) {
		return nil, false
	}
	return o.CcsInstanceId, true
}

// HasCcsInstanceId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRoleCreateModel) HasCcsInstanceId() bool {
	if o != nil && !IsNil(o.CcsInstanceId) {
		return true
	}

	return false
}

// SetCcsInstanceId gets a reference to the given string and assigns it to the CcsInstanceId field.
func (o *AlipayIserviceCcmRoleCreateModel) SetCcsInstanceId(v string) {
	o.CcsInstanceId = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRoleCreateModel) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRoleCreateModel) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRoleCreateModel) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AlipayIserviceCcmRoleCreateModel) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRoleCreateModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRoleCreateModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRoleCreateModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlipayIserviceCcmRoleCreateModel) SetDescription(v string) {
	o.Description = &v
}

// GetFunctionIds returns the FunctionIds field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRoleCreateModel) GetFunctionIds() []string {
	if o == nil || IsNil(o.FunctionIds) {
		var ret []string
		return ret
	}
	return o.FunctionIds
}

// GetFunctionIdsOk returns a tuple with the FunctionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRoleCreateModel) GetFunctionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FunctionIds) {
		return nil, false
	}
	return o.FunctionIds, true
}

// HasFunctionIds returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRoleCreateModel) HasFunctionIds() bool {
	if o != nil && !IsNil(o.FunctionIds) {
		return true
	}

	return false
}

// SetFunctionIds gets a reference to the given []string and assigns it to the FunctionIds field.
func (o *AlipayIserviceCcmRoleCreateModel) SetFunctionIds(v []string) {
	o.FunctionIds = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRoleCreateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRoleCreateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRoleCreateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayIserviceCcmRoleCreateModel) SetName(v string) {
	o.Name = &v
}

func (o AlipayIserviceCcmRoleCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmRoleCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CcsInstanceId) {
		toSerialize["ccs_instance_id"] = o.CcsInstanceId
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FunctionIds) {
		toSerialize["function_ids"] = o.FunctionIds
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmRoleCreateModel struct {
	value *AlipayIserviceCcmRoleCreateModel
	isSet bool
}

func (v NullableAlipayIserviceCcmRoleCreateModel) Get() *AlipayIserviceCcmRoleCreateModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmRoleCreateModel) Set(val *AlipayIserviceCcmRoleCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmRoleCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmRoleCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmRoleCreateModel(val *AlipayIserviceCcmRoleCreateModel) *NullableAlipayIserviceCcmRoleCreateModel {
	return &NullableAlipayIserviceCcmRoleCreateModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmRoleCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmRoleCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


