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

// checks if the AntMerchantExpandItemCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandItemCreateResponseModel{}

// AntMerchantExpandItemCreateResponseModel struct for AntMerchantExpandItemCreateResponseModel
type AntMerchantExpandItemCreateResponseModel struct {
	// 商品id
	ItemId *string `json:"item_id,omitempty"`
}

// NewAntMerchantExpandItemCreateResponseModel instantiates a new AntMerchantExpandItemCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandItemCreateResponseModel() *AntMerchantExpandItemCreateResponseModel {
	this := AntMerchantExpandItemCreateResponseModel{}
	return &this
}

// NewAntMerchantExpandItemCreateResponseModelWithDefaults instantiates a new AntMerchantExpandItemCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandItemCreateResponseModelWithDefaults() *AntMerchantExpandItemCreateResponseModel {
	this := AntMerchantExpandItemCreateResponseModel{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *AntMerchantExpandItemCreateResponseModel) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandItemCreateResponseModel) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *AntMerchantExpandItemCreateResponseModel) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *AntMerchantExpandItemCreateResponseModel) SetItemId(v string) {
	o.ItemId = &v
}

func (o AntMerchantExpandItemCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandItemCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandItemCreateResponseModel struct {
	value *AntMerchantExpandItemCreateResponseModel
	isSet bool
}

func (v NullableAntMerchantExpandItemCreateResponseModel) Get() *AntMerchantExpandItemCreateResponseModel {
	return v.value
}

func (v *NullableAntMerchantExpandItemCreateResponseModel) Set(val *AntMerchantExpandItemCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandItemCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandItemCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandItemCreateResponseModel(val *AntMerchantExpandItemCreateResponseModel) *NullableAntMerchantExpandItemCreateResponseModel {
	return &NullableAntMerchantExpandItemCreateResponseModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandItemCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandItemCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


