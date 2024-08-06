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

// checks if the MiniAppPluginReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiniAppPluginReference{}

// MiniAppPluginReference struct for MiniAppPluginReference
type MiniAppPluginReference struct {
	// 是否懒加载
	Lazy *bool `json:"lazy,omitempty"`
	// 插件id，即插件的唯一标识，形如小程序的appId
	PluginId *string `json:"plugin_id,omitempty"`
	// 是否懒加载
	PluginLazy *bool `json:"plugin_lazy,omitempty"`
	// 小程序引用的插件版本，可以指定插件版本，如：1.0.11；也可以填*，表示引用最新的插件版本，此时，当插件发布新版本后，小程序内的插件会自动更新到最新的版本。
	PluginVersion *string `json:"plugin_version,omitempty"`
}

// NewMiniAppPluginReference instantiates a new MiniAppPluginReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniAppPluginReference() *MiniAppPluginReference {
	this := MiniAppPluginReference{}
	return &this
}

// NewMiniAppPluginReferenceWithDefaults instantiates a new MiniAppPluginReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniAppPluginReferenceWithDefaults() *MiniAppPluginReference {
	this := MiniAppPluginReference{}
	return &this
}

// GetLazy returns the Lazy field value if set, zero value otherwise.
func (o *MiniAppPluginReference) GetLazy() bool {
	if o == nil || IsNil(o.Lazy) {
		var ret bool
		return ret
	}
	return *o.Lazy
}

// GetLazyOk returns a tuple with the Lazy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppPluginReference) GetLazyOk() (*bool, bool) {
	if o == nil || IsNil(o.Lazy) {
		return nil, false
	}
	return o.Lazy, true
}

// HasLazy returns a boolean if a field has been set.
func (o *MiniAppPluginReference) HasLazy() bool {
	if o != nil && !IsNil(o.Lazy) {
		return true
	}

	return false
}

// SetLazy gets a reference to the given bool and assigns it to the Lazy field.
func (o *MiniAppPluginReference) SetLazy(v bool) {
	o.Lazy = &v
}

// GetPluginId returns the PluginId field value if set, zero value otherwise.
func (o *MiniAppPluginReference) GetPluginId() string {
	if o == nil || IsNil(o.PluginId) {
		var ret string
		return ret
	}
	return *o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppPluginReference) GetPluginIdOk() (*string, bool) {
	if o == nil || IsNil(o.PluginId) {
		return nil, false
	}
	return o.PluginId, true
}

// HasPluginId returns a boolean if a field has been set.
func (o *MiniAppPluginReference) HasPluginId() bool {
	if o != nil && !IsNil(o.PluginId) {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given string and assigns it to the PluginId field.
func (o *MiniAppPluginReference) SetPluginId(v string) {
	o.PluginId = &v
}

// GetPluginLazy returns the PluginLazy field value if set, zero value otherwise.
func (o *MiniAppPluginReference) GetPluginLazy() bool {
	if o == nil || IsNil(o.PluginLazy) {
		var ret bool
		return ret
	}
	return *o.PluginLazy
}

// GetPluginLazyOk returns a tuple with the PluginLazy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppPluginReference) GetPluginLazyOk() (*bool, bool) {
	if o == nil || IsNil(o.PluginLazy) {
		return nil, false
	}
	return o.PluginLazy, true
}

// HasPluginLazy returns a boolean if a field has been set.
func (o *MiniAppPluginReference) HasPluginLazy() bool {
	if o != nil && !IsNil(o.PluginLazy) {
		return true
	}

	return false
}

// SetPluginLazy gets a reference to the given bool and assigns it to the PluginLazy field.
func (o *MiniAppPluginReference) SetPluginLazy(v bool) {
	o.PluginLazy = &v
}

// GetPluginVersion returns the PluginVersion field value if set, zero value otherwise.
func (o *MiniAppPluginReference) GetPluginVersion() string {
	if o == nil || IsNil(o.PluginVersion) {
		var ret string
		return ret
	}
	return *o.PluginVersion
}

// GetPluginVersionOk returns a tuple with the PluginVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppPluginReference) GetPluginVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PluginVersion) {
		return nil, false
	}
	return o.PluginVersion, true
}

// HasPluginVersion returns a boolean if a field has been set.
func (o *MiniAppPluginReference) HasPluginVersion() bool {
	if o != nil && !IsNil(o.PluginVersion) {
		return true
	}

	return false
}

// SetPluginVersion gets a reference to the given string and assigns it to the PluginVersion field.
func (o *MiniAppPluginReference) SetPluginVersion(v string) {
	o.PluginVersion = &v
}

func (o MiniAppPluginReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniAppPluginReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lazy) {
		toSerialize["lazy"] = o.Lazy
	}
	if !IsNil(o.PluginId) {
		toSerialize["plugin_id"] = o.PluginId
	}
	if !IsNil(o.PluginLazy) {
		toSerialize["plugin_lazy"] = o.PluginLazy
	}
	if !IsNil(o.PluginVersion) {
		toSerialize["plugin_version"] = o.PluginVersion
	}
	return toSerialize, nil
}

type NullableMiniAppPluginReference struct {
	value *MiniAppPluginReference
	isSet bool
}

func (v NullableMiniAppPluginReference) Get() *MiniAppPluginReference {
	return v.value
}

func (v *NullableMiniAppPluginReference) Set(val *MiniAppPluginReference) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniAppPluginReference) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniAppPluginReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniAppPluginReference(val *MiniAppPluginReference) *NullableMiniAppPluginReference {
	return &NullableMiniAppPluginReference{value: val, isSet: true}
}

func (v NullableMiniAppPluginReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniAppPluginReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


