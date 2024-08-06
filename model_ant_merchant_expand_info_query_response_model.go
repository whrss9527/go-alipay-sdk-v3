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

// checks if the AntMerchantExpandInfoQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandInfoQueryResponseModel{}

// AntMerchantExpandInfoQueryResponseModel struct for AntMerchantExpandInfoQueryResponseModel
type AntMerchantExpandInfoQueryResponseModel struct {
	MerchantQueryResult *MerchantQueryResult `json:"merchant_query_result,omitempty"`
}

// NewAntMerchantExpandInfoQueryResponseModel instantiates a new AntMerchantExpandInfoQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandInfoQueryResponseModel() *AntMerchantExpandInfoQueryResponseModel {
	this := AntMerchantExpandInfoQueryResponseModel{}
	return &this
}

// NewAntMerchantExpandInfoQueryResponseModelWithDefaults instantiates a new AntMerchantExpandInfoQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandInfoQueryResponseModelWithDefaults() *AntMerchantExpandInfoQueryResponseModel {
	this := AntMerchantExpandInfoQueryResponseModel{}
	return &this
}

// GetMerchantQueryResult returns the MerchantQueryResult field value if set, zero value otherwise.
func (o *AntMerchantExpandInfoQueryResponseModel) GetMerchantQueryResult() MerchantQueryResult {
	if o == nil || IsNil(o.MerchantQueryResult) {
		var ret MerchantQueryResult
		return ret
	}
	return *o.MerchantQueryResult
}

// GetMerchantQueryResultOk returns a tuple with the MerchantQueryResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandInfoQueryResponseModel) GetMerchantQueryResultOk() (*MerchantQueryResult, bool) {
	if o == nil || IsNil(o.MerchantQueryResult) {
		return nil, false
	}
	return o.MerchantQueryResult, true
}

// HasMerchantQueryResult returns a boolean if a field has been set.
func (o *AntMerchantExpandInfoQueryResponseModel) HasMerchantQueryResult() bool {
	if o != nil && !IsNil(o.MerchantQueryResult) {
		return true
	}

	return false
}

// SetMerchantQueryResult gets a reference to the given MerchantQueryResult and assigns it to the MerchantQueryResult field.
func (o *AntMerchantExpandInfoQueryResponseModel) SetMerchantQueryResult(v MerchantQueryResult) {
	o.MerchantQueryResult = &v
}

func (o AntMerchantExpandInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandInfoQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantQueryResult) {
		toSerialize["merchant_query_result"] = o.MerchantQueryResult
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandInfoQueryResponseModel struct {
	value *AntMerchantExpandInfoQueryResponseModel
	isSet bool
}

func (v NullableAntMerchantExpandInfoQueryResponseModel) Get() *AntMerchantExpandInfoQueryResponseModel {
	return v.value
}

func (v *NullableAntMerchantExpandInfoQueryResponseModel) Set(val *AntMerchantExpandInfoQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandInfoQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandInfoQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandInfoQueryResponseModel(val *AntMerchantExpandInfoQueryResponseModel) *NullableAntMerchantExpandInfoQueryResponseModel {
	return &NullableAntMerchantExpandInfoQueryResponseModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandInfoQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

