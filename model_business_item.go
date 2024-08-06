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

// checks if the BusinessItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessItem{}

// BusinessItem struct for BusinessItem
type BusinessItem struct {
	// 业务归属appid
	BusinessAppid *string `json:"business_appid,omitempty"`
	// 业务收款pid，business_type为AGREEMENT_PAY时，此参数必传
	BusinessPayeeId *string `json:"business_payee_id,omitempty"`
	// 业务归属pid
	BusinessPid *string `json:"business_pid,omitempty"`
	// 业务类型ONLINE_PAY(在线缴费) 、AGREEMENT_PAY(无感停车) DEVICE_ONLINE_PAY(车机在线缴费)
	BusinessType *string `json:"business_type,omitempty"`
}

// NewBusinessItem instantiates a new BusinessItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessItem() *BusinessItem {
	this := BusinessItem{}
	return &this
}

// NewBusinessItemWithDefaults instantiates a new BusinessItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessItemWithDefaults() *BusinessItem {
	this := BusinessItem{}
	return &this
}

// GetBusinessAppid returns the BusinessAppid field value if set, zero value otherwise.
func (o *BusinessItem) GetBusinessAppid() string {
	if o == nil || IsNil(o.BusinessAppid) {
		var ret string
		return ret
	}
	return *o.BusinessAppid
}

// GetBusinessAppidOk returns a tuple with the BusinessAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessItem) GetBusinessAppidOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessAppid) {
		return nil, false
	}
	return o.BusinessAppid, true
}

// HasBusinessAppid returns a boolean if a field has been set.
func (o *BusinessItem) HasBusinessAppid() bool {
	if o != nil && !IsNil(o.BusinessAppid) {
		return true
	}

	return false
}

// SetBusinessAppid gets a reference to the given string and assigns it to the BusinessAppid field.
func (o *BusinessItem) SetBusinessAppid(v string) {
	o.BusinessAppid = &v
}

// GetBusinessPayeeId returns the BusinessPayeeId field value if set, zero value otherwise.
func (o *BusinessItem) GetBusinessPayeeId() string {
	if o == nil || IsNil(o.BusinessPayeeId) {
		var ret string
		return ret
	}
	return *o.BusinessPayeeId
}

// GetBusinessPayeeIdOk returns a tuple with the BusinessPayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessItem) GetBusinessPayeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessPayeeId) {
		return nil, false
	}
	return o.BusinessPayeeId, true
}

// HasBusinessPayeeId returns a boolean if a field has been set.
func (o *BusinessItem) HasBusinessPayeeId() bool {
	if o != nil && !IsNil(o.BusinessPayeeId) {
		return true
	}

	return false
}

// SetBusinessPayeeId gets a reference to the given string and assigns it to the BusinessPayeeId field.
func (o *BusinessItem) SetBusinessPayeeId(v string) {
	o.BusinessPayeeId = &v
}

// GetBusinessPid returns the BusinessPid field value if set, zero value otherwise.
func (o *BusinessItem) GetBusinessPid() string {
	if o == nil || IsNil(o.BusinessPid) {
		var ret string
		return ret
	}
	return *o.BusinessPid
}

// GetBusinessPidOk returns a tuple with the BusinessPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessItem) GetBusinessPidOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessPid) {
		return nil, false
	}
	return o.BusinessPid, true
}

// HasBusinessPid returns a boolean if a field has been set.
func (o *BusinessItem) HasBusinessPid() bool {
	if o != nil && !IsNil(o.BusinessPid) {
		return true
	}

	return false
}

// SetBusinessPid gets a reference to the given string and assigns it to the BusinessPid field.
func (o *BusinessItem) SetBusinessPid(v string) {
	o.BusinessPid = &v
}

// GetBusinessType returns the BusinessType field value if set, zero value otherwise.
func (o *BusinessItem) GetBusinessType() string {
	if o == nil || IsNil(o.BusinessType) {
		var ret string
		return ret
	}
	return *o.BusinessType
}

// GetBusinessTypeOk returns a tuple with the BusinessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessItem) GetBusinessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessType) {
		return nil, false
	}
	return o.BusinessType, true
}

// HasBusinessType returns a boolean if a field has been set.
func (o *BusinessItem) HasBusinessType() bool {
	if o != nil && !IsNil(o.BusinessType) {
		return true
	}

	return false
}

// SetBusinessType gets a reference to the given string and assigns it to the BusinessType field.
func (o *BusinessItem) SetBusinessType(v string) {
	o.BusinessType = &v
}

func (o BusinessItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessAppid) {
		toSerialize["business_appid"] = o.BusinessAppid
	}
	if !IsNil(o.BusinessPayeeId) {
		toSerialize["business_payee_id"] = o.BusinessPayeeId
	}
	if !IsNil(o.BusinessPid) {
		toSerialize["business_pid"] = o.BusinessPid
	}
	if !IsNil(o.BusinessType) {
		toSerialize["business_type"] = o.BusinessType
	}
	return toSerialize, nil
}

type NullableBusinessItem struct {
	value *BusinessItem
	isSet bool
}

func (v NullableBusinessItem) Get() *BusinessItem {
	return v.value
}

func (v *NullableBusinessItem) Set(val *BusinessItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessItem(val *BusinessItem) *NullableBusinessItem {
	return &NullableBusinessItem{value: val, isSet: true}
}

func (v NullableBusinessItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


