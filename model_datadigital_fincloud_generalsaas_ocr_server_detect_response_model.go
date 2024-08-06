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

// checks if the DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel{}

// DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel struct for DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel
type DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel struct {
	// OCR识别单据号，计费依据，请保留以便排查问题。
	CertifyId *string `json:"certify_id,omitempty"`
	// OCR识别结果
	OcrData *string `json:"ocr_data,omitempty"`
}

// NewDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel instantiates a new DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel() *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel {
	this := DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel{}
	return &this
}

// NewDatadigitalFincloudGeneralsaasOcrServerDetectResponseModelWithDefaults instantiates a new DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadigitalFincloudGeneralsaasOcrServerDetectResponseModelWithDefaults() *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel {
	this := DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel{}
	return &this
}

// GetCertifyId returns the CertifyId field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) GetCertifyId() string {
	if o == nil || IsNil(o.CertifyId) {
		var ret string
		return ret
	}
	return *o.CertifyId
}

// GetCertifyIdOk returns a tuple with the CertifyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) GetCertifyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CertifyId) {
		return nil, false
	}
	return o.CertifyId, true
}

// HasCertifyId returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) HasCertifyId() bool {
	if o != nil && !IsNil(o.CertifyId) {
		return true
	}

	return false
}

// SetCertifyId gets a reference to the given string and assigns it to the CertifyId field.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) SetCertifyId(v string) {
	o.CertifyId = &v
}

// GetOcrData returns the OcrData field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) GetOcrData() string {
	if o == nil || IsNil(o.OcrData) {
		var ret string
		return ret
	}
	return *o.OcrData
}

// GetOcrDataOk returns a tuple with the OcrData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) GetOcrDataOk() (*string, bool) {
	if o == nil || IsNil(o.OcrData) {
		return nil, false
	}
	return o.OcrData, true
}

// HasOcrData returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) HasOcrData() bool {
	if o != nil && !IsNil(o.OcrData) {
		return true
	}

	return false
}

// SetOcrData gets a reference to the given string and assigns it to the OcrData field.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) SetOcrData(v string) {
	o.OcrData = &v
}

func (o DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertifyId) {
		toSerialize["certify_id"] = o.CertifyId
	}
	if !IsNil(o.OcrData) {
		toSerialize["ocr_data"] = o.OcrData
	}
	return toSerialize, nil
}

type NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel struct {
	value *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) Get() *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) Set(val *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel(val *DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) *NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel {
	return &NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
