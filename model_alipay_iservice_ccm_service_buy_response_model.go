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

// checks if the AlipayIserviceCcmServiceBuyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmServiceBuyResponseModel{}

// AlipayIserviceCcmServiceBuyResponseModel struct for AlipayIserviceCcmServiceBuyResponseModel
type AlipayIserviceCcmServiceBuyResponseModel struct {
	// 服务实例id
	ServiceInstanceId *int32 `json:"service_instance_id,omitempty"`
	// 商户在CCM的租户id
	TenantId *string `json:"tenant_id,omitempty"`
}

// NewAlipayIserviceCcmServiceBuyResponseModel instantiates a new AlipayIserviceCcmServiceBuyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmServiceBuyResponseModel() *AlipayIserviceCcmServiceBuyResponseModel {
	this := AlipayIserviceCcmServiceBuyResponseModel{}
	return &this
}

// NewAlipayIserviceCcmServiceBuyResponseModelWithDefaults instantiates a new AlipayIserviceCcmServiceBuyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmServiceBuyResponseModelWithDefaults() *AlipayIserviceCcmServiceBuyResponseModel {
	this := AlipayIserviceCcmServiceBuyResponseModel{}
	return &this
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmServiceBuyResponseModel) GetServiceInstanceId() int32 {
	if o == nil || IsNil(o.ServiceInstanceId) {
		var ret int32
		return ret
	}
	return *o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmServiceBuyResponseModel) GetServiceInstanceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceInstanceId) {
		return nil, false
	}
	return o.ServiceInstanceId, true
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmServiceBuyResponseModel) HasServiceInstanceId() bool {
	if o != nil && !IsNil(o.ServiceInstanceId) {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given int32 and assigns it to the ServiceInstanceId field.
func (o *AlipayIserviceCcmServiceBuyResponseModel) SetServiceInstanceId(v int32) {
	o.ServiceInstanceId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmServiceBuyResponseModel) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmServiceBuyResponseModel) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmServiceBuyResponseModel) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AlipayIserviceCcmServiceBuyResponseModel) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AlipayIserviceCcmServiceBuyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmServiceBuyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceInstanceId) {
		toSerialize["service_instance_id"] = o.ServiceInstanceId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmServiceBuyResponseModel struct {
	value *AlipayIserviceCcmServiceBuyResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmServiceBuyResponseModel) Get() *AlipayIserviceCcmServiceBuyResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmServiceBuyResponseModel) Set(val *AlipayIserviceCcmServiceBuyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmServiceBuyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmServiceBuyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmServiceBuyResponseModel(val *AlipayIserviceCcmServiceBuyResponseModel) *NullableAlipayIserviceCcmServiceBuyResponseModel {
	return &NullableAlipayIserviceCcmServiceBuyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmServiceBuyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmServiceBuyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
