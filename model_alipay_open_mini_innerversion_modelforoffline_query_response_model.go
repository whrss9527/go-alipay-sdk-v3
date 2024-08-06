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

// checks if the AlipayOpenMiniInnerversionModelforofflineQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionModelforofflineQueryResponseModel{}

// AlipayOpenMiniInnerversionModelforofflineQueryResponseModel struct for AlipayOpenMiniInnerversionModelforofflineQueryResponseModel
type AlipayOpenMiniInnerversionModelforofflineQueryResponseModel struct {
	// 线上版本包数据
	ModelJson *string `json:"model_json,omitempty"`
	// 同步ID，同步失败时便于去线上追溯
	SyncId *string `json:"sync_id,omitempty"`
}

// NewAlipayOpenMiniInnerversionModelforofflineQueryResponseModel instantiates a new AlipayOpenMiniInnerversionModelforofflineQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionModelforofflineQueryResponseModel() *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel {
	this := AlipayOpenMiniInnerversionModelforofflineQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniInnerversionModelforofflineQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerversionModelforofflineQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionModelforofflineQueryResponseModelWithDefaults() *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel {
	this := AlipayOpenMiniInnerversionModelforofflineQueryResponseModel{}
	return &this
}

// GetModelJson returns the ModelJson field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) GetModelJson() string {
	if o == nil || IsNil(o.ModelJson) {
		var ret string
		return ret
	}
	return *o.ModelJson
}

// GetModelJsonOk returns a tuple with the ModelJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) GetModelJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ModelJson) {
		return nil, false
	}
	return o.ModelJson, true
}

// HasModelJson returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) HasModelJson() bool {
	if o != nil && !IsNil(o.ModelJson) {
		return true
	}

	return false
}

// SetModelJson gets a reference to the given string and assigns it to the ModelJson field.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) SetModelJson(v string) {
	o.ModelJson = &v
}

// GetSyncId returns the SyncId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) GetSyncId() string {
	if o == nil || IsNil(o.SyncId) {
		var ret string
		return ret
	}
	return *o.SyncId
}

// GetSyncIdOk returns a tuple with the SyncId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) GetSyncIdOk() (*string, bool) {
	if o == nil || IsNil(o.SyncId) {
		return nil, false
	}
	return o.SyncId, true
}

// HasSyncId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) HasSyncId() bool {
	if o != nil && !IsNil(o.SyncId) {
		return true
	}

	return false
}

// SetSyncId gets a reference to the given string and assigns it to the SyncId field.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) SetSyncId(v string) {
	o.SyncId = &v
}

func (o AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModelJson) {
		toSerialize["model_json"] = o.ModelJson
	}
	if !IsNil(o.SyncId) {
		toSerialize["sync_id"] = o.SyncId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel struct {
	value *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel) Get() *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel) Set(val *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel(val *AlipayOpenMiniInnerversionModelforofflineQueryResponseModel) *NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel {
	return &NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionModelforofflineQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


