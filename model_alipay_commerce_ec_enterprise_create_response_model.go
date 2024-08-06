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

// checks if the AlipayCommerceEcEnterpriseCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcEnterpriseCreateResponseModel{}

// AlipayCommerceEcEnterpriseCreateResponseModel struct for AlipayCommerceEcEnterpriseCreateResponseModel
type AlipayCommerceEcEnterpriseCreateResponseModel struct {
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 封闭场景（如班车）的人脸库id，如果入参create_iot_group传false，则不会创建企业人脸库，该字段为空。
	IotGroupId *string `json:"iot_group_id,omitempty"`
	// 开放场景（如团餐）的人脸库id，如果入参create_iot_group传false，则不会创建企业人脸库，该字段为空。
	IotLogicGroupId *string `json:"iot_logic_group_id,omitempty"`
	// 资金代付签约链接
	SignUrl *string `json:"sign_url,omitempty"`
}

// NewAlipayCommerceEcEnterpriseCreateResponseModel instantiates a new AlipayCommerceEcEnterpriseCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcEnterpriseCreateResponseModel() *AlipayCommerceEcEnterpriseCreateResponseModel {
	this := AlipayCommerceEcEnterpriseCreateResponseModel{}
	return &this
}

// NewAlipayCommerceEcEnterpriseCreateResponseModelWithDefaults instantiates a new AlipayCommerceEcEnterpriseCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcEnterpriseCreateResponseModelWithDefaults() *AlipayCommerceEcEnterpriseCreateResponseModel {
	this := AlipayCommerceEcEnterpriseCreateResponseModel{}
	return &this
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetIotGroupId returns the IotGroupId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetIotGroupId() string {
	if o == nil || IsNil(o.IotGroupId) {
		var ret string
		return ret
	}
	return *o.IotGroupId
}

// GetIotGroupIdOk returns a tuple with the IotGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetIotGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.IotGroupId) {
		return nil, false
	}
	return o.IotGroupId, true
}

// HasIotGroupId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) HasIotGroupId() bool {
	if o != nil && !IsNil(o.IotGroupId) {
		return true
	}

	return false
}

// SetIotGroupId gets a reference to the given string and assigns it to the IotGroupId field.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) SetIotGroupId(v string) {
	o.IotGroupId = &v
}

// GetIotLogicGroupId returns the IotLogicGroupId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetIotLogicGroupId() string {
	if o == nil || IsNil(o.IotLogicGroupId) {
		var ret string
		return ret
	}
	return *o.IotLogicGroupId
}

// GetIotLogicGroupIdOk returns a tuple with the IotLogicGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetIotLogicGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.IotLogicGroupId) {
		return nil, false
	}
	return o.IotLogicGroupId, true
}

// HasIotLogicGroupId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) HasIotLogicGroupId() bool {
	if o != nil && !IsNil(o.IotLogicGroupId) {
		return true
	}

	return false
}

// SetIotLogicGroupId gets a reference to the given string and assigns it to the IotLogicGroupId field.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) SetIotLogicGroupId(v string) {
	o.IotLogicGroupId = &v
}

// GetSignUrl returns the SignUrl field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetSignUrl() string {
	if o == nil || IsNil(o.SignUrl) {
		var ret string
		return ret
	}
	return *o.SignUrl
}

// GetSignUrlOk returns a tuple with the SignUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) GetSignUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SignUrl) {
		return nil, false
	}
	return o.SignUrl, true
}

// HasSignUrl returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) HasSignUrl() bool {
	if o != nil && !IsNil(o.SignUrl) {
		return true
	}

	return false
}

// SetSignUrl gets a reference to the given string and assigns it to the SignUrl field.
func (o *AlipayCommerceEcEnterpriseCreateResponseModel) SetSignUrl(v string) {
	o.SignUrl = &v
}

func (o AlipayCommerceEcEnterpriseCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcEnterpriseCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.IotGroupId) {
		toSerialize["iot_group_id"] = o.IotGroupId
	}
	if !IsNil(o.IotLogicGroupId) {
		toSerialize["iot_logic_group_id"] = o.IotLogicGroupId
	}
	if !IsNil(o.SignUrl) {
		toSerialize["sign_url"] = o.SignUrl
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcEnterpriseCreateResponseModel struct {
	value *AlipayCommerceEcEnterpriseCreateResponseModel
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseCreateResponseModel) Get() *AlipayCommerceEcEnterpriseCreateResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseCreateResponseModel) Set(val *AlipayCommerceEcEnterpriseCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseCreateResponseModel(val *AlipayCommerceEcEnterpriseCreateResponseModel) *NullableAlipayCommerceEcEnterpriseCreateResponseModel {
	return &NullableAlipayCommerceEcEnterpriseCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

