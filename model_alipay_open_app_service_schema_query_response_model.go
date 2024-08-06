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

// checks if the AlipayOpenAppServiceSchemaQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAppServiceSchemaQueryResponseModel{}

// AlipayOpenAppServiceSchemaQueryResponseModel struct for AlipayOpenAppServiceSchemaQueryResponseModel
type AlipayOpenAppServiceSchemaQueryResponseModel struct {
	// 类目id
	CategoryId *string `json:"category_id,omitempty"`
	// 服务schema版本
	SchemaVersion *string `json:"schema_version,omitempty"`
	// 服务schema
	SchemaXml *string `json:"schema_xml,omitempty"`
	// 服务模版类型，默认值: DEFAULT
	TemplateType *string `json:"template_type,omitempty"`
	// 商户提报服务履约类型
	UserServiceDeliveryType *string `json:"user_service_delivery_type,omitempty"`
}

// NewAlipayOpenAppServiceSchemaQueryResponseModel instantiates a new AlipayOpenAppServiceSchemaQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAppServiceSchemaQueryResponseModel() *AlipayOpenAppServiceSchemaQueryResponseModel {
	this := AlipayOpenAppServiceSchemaQueryResponseModel{}
	return &this
}

// NewAlipayOpenAppServiceSchemaQueryResponseModelWithDefaults instantiates a new AlipayOpenAppServiceSchemaQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAppServiceSchemaQueryResponseModelWithDefaults() *AlipayOpenAppServiceSchemaQueryResponseModel {
	this := AlipayOpenAppServiceSchemaQueryResponseModel{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetSchemaVersion() string {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetSchemaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetSchemaXml returns the SchemaXml field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetSchemaXml() string {
	if o == nil || IsNil(o.SchemaXml) {
		var ret string
		return ret
	}
	return *o.SchemaXml
}

// GetSchemaXmlOk returns a tuple with the SchemaXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetSchemaXmlOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaXml) {
		return nil, false
	}
	return o.SchemaXml, true
}

// HasSchemaXml returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) HasSchemaXml() bool {
	if o != nil && !IsNil(o.SchemaXml) {
		return true
	}

	return false
}

// SetSchemaXml gets a reference to the given string and assigns it to the SchemaXml field.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) SetSchemaXml(v string) {
	o.SchemaXml = &v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetTemplateType() string {
	if o == nil || IsNil(o.TemplateType) {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetTemplateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateType) {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) HasTemplateType() bool {
	if o != nil && !IsNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) SetTemplateType(v string) {
	o.TemplateType = &v
}

// GetUserServiceDeliveryType returns the UserServiceDeliveryType field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetUserServiceDeliveryType() string {
	if o == nil || IsNil(o.UserServiceDeliveryType) {
		var ret string
		return ret
	}
	return *o.UserServiceDeliveryType
}

// GetUserServiceDeliveryTypeOk returns a tuple with the UserServiceDeliveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) GetUserServiceDeliveryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserServiceDeliveryType) {
		return nil, false
	}
	return o.UserServiceDeliveryType, true
}

// HasUserServiceDeliveryType returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) HasUserServiceDeliveryType() bool {
	if o != nil && !IsNil(o.UserServiceDeliveryType) {
		return true
	}

	return false
}

// SetUserServiceDeliveryType gets a reference to the given string and assigns it to the UserServiceDeliveryType field.
func (o *AlipayOpenAppServiceSchemaQueryResponseModel) SetUserServiceDeliveryType(v string) {
	o.UserServiceDeliveryType = &v
}

func (o AlipayOpenAppServiceSchemaQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAppServiceSchemaQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if !IsNil(o.SchemaXml) {
		toSerialize["schema_xml"] = o.SchemaXml
	}
	if !IsNil(o.TemplateType) {
		toSerialize["template_type"] = o.TemplateType
	}
	if !IsNil(o.UserServiceDeliveryType) {
		toSerialize["user_service_delivery_type"] = o.UserServiceDeliveryType
	}
	return toSerialize, nil
}

type NullableAlipayOpenAppServiceSchemaQueryResponseModel struct {
	value *AlipayOpenAppServiceSchemaQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenAppServiceSchemaQueryResponseModel) Get() *AlipayOpenAppServiceSchemaQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAppServiceSchemaQueryResponseModel) Set(val *AlipayOpenAppServiceSchemaQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppServiceSchemaQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppServiceSchemaQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppServiceSchemaQueryResponseModel(val *AlipayOpenAppServiceSchemaQueryResponseModel) *NullableAlipayOpenAppServiceSchemaQueryResponseModel {
	return &NullableAlipayOpenAppServiceSchemaQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAppServiceSchemaQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppServiceSchemaQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

