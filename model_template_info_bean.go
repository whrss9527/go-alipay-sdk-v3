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

// checks if the TemplateInfoBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateInfoBean{}

// TemplateInfoBean struct for TemplateInfoBean
type TemplateInfoBean struct {
	// 模板填充项，根据组件key值传入填写内容
	FillContents []FillContent `json:"fill_contents,omitempty"`
	// 签署文件名称
	Name *string `json:"name,omitempty"`
	// 签署区，根据签署区key值传入对应的签署人信息
	Signfields []SignFieldBean `json:"signfields,omitempty"`
	// 模板id，通过创建合同模板获取
	TemplateId *string `json:"template_id,omitempty"`
}

// NewTemplateInfoBean instantiates a new TemplateInfoBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateInfoBean() *TemplateInfoBean {
	this := TemplateInfoBean{}
	return &this
}

// NewTemplateInfoBeanWithDefaults instantiates a new TemplateInfoBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateInfoBeanWithDefaults() *TemplateInfoBean {
	this := TemplateInfoBean{}
	return &this
}

// GetFillContents returns the FillContents field value if set, zero value otherwise.
func (o *TemplateInfoBean) GetFillContents() []FillContent {
	if o == nil || IsNil(o.FillContents) {
		var ret []FillContent
		return ret
	}
	return o.FillContents
}

// GetFillContentsOk returns a tuple with the FillContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInfoBean) GetFillContentsOk() ([]FillContent, bool) {
	if o == nil || IsNil(o.FillContents) {
		return nil, false
	}
	return o.FillContents, true
}

// HasFillContents returns a boolean if a field has been set.
func (o *TemplateInfoBean) HasFillContents() bool {
	if o != nil && !IsNil(o.FillContents) {
		return true
	}

	return false
}

// SetFillContents gets a reference to the given []FillContent and assigns it to the FillContents field.
func (o *TemplateInfoBean) SetFillContents(v []FillContent) {
	o.FillContents = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateInfoBean) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInfoBean) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateInfoBean) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateInfoBean) SetName(v string) {
	o.Name = &v
}

// GetSignfields returns the Signfields field value if set, zero value otherwise.
func (o *TemplateInfoBean) GetSignfields() []SignFieldBean {
	if o == nil || IsNil(o.Signfields) {
		var ret []SignFieldBean
		return ret
	}
	return o.Signfields
}

// GetSignfieldsOk returns a tuple with the Signfields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInfoBean) GetSignfieldsOk() ([]SignFieldBean, bool) {
	if o == nil || IsNil(o.Signfields) {
		return nil, false
	}
	return o.Signfields, true
}

// HasSignfields returns a boolean if a field has been set.
func (o *TemplateInfoBean) HasSignfields() bool {
	if o != nil && !IsNil(o.Signfields) {
		return true
	}

	return false
}

// SetSignfields gets a reference to the given []SignFieldBean and assigns it to the Signfields field.
func (o *TemplateInfoBean) SetSignfields(v []SignFieldBean) {
	o.Signfields = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *TemplateInfoBean) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInfoBean) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *TemplateInfoBean) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *TemplateInfoBean) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o TemplateInfoBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateInfoBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FillContents) {
		toSerialize["fill_contents"] = o.FillContents
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Signfields) {
		toSerialize["signfields"] = o.Signfields
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableTemplateInfoBean struct {
	value *TemplateInfoBean
	isSet bool
}

func (v NullableTemplateInfoBean) Get() *TemplateInfoBean {
	return v.value
}

func (v *NullableTemplateInfoBean) Set(val *TemplateInfoBean) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateInfoBean) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateInfoBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateInfoBean(val *TemplateInfoBean) *NullableTemplateInfoBean {
	return &NullableTemplateInfoBean{value: val, isSet: true}
}

func (v NullableTemplateInfoBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateInfoBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


