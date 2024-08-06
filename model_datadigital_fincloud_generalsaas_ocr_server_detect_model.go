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

// checks if the DatadigitalFincloudGeneralsaasOcrServerDetectModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatadigitalFincloudGeneralsaasOcrServerDetectModel{}

// DatadigitalFincloudGeneralsaasOcrServerDetectModel struct for DatadigitalFincloudGeneralsaasOcrServerDetectModel
type DatadigitalFincloudGeneralsaasOcrServerDetectModel struct {
	// 可识别OCR类型。如，ID_CARD_FRONT 身份证含照片的一面；ID_CARD_BACK 身份证带国徽的一面。具体支持的类型以接入文档为准。
	OcrType *string `json:"ocr_type,omitempty"`
	// 客户业务单据号，请保证幂等性。
	OuterOrderNo *string `json:"outer_order_no,omitempty"`
}

// NewDatadigitalFincloudGeneralsaasOcrServerDetectModel instantiates a new DatadigitalFincloudGeneralsaasOcrServerDetectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadigitalFincloudGeneralsaasOcrServerDetectModel() *DatadigitalFincloudGeneralsaasOcrServerDetectModel {
	this := DatadigitalFincloudGeneralsaasOcrServerDetectModel{}
	return &this
}

// NewDatadigitalFincloudGeneralsaasOcrServerDetectModelWithDefaults instantiates a new DatadigitalFincloudGeneralsaasOcrServerDetectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadigitalFincloudGeneralsaasOcrServerDetectModelWithDefaults() *DatadigitalFincloudGeneralsaasOcrServerDetectModel {
	this := DatadigitalFincloudGeneralsaasOcrServerDetectModel{}
	return &this
}

// GetOcrType returns the OcrType field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) GetOcrType() string {
	if o == nil || IsNil(o.OcrType) {
		var ret string
		return ret
	}
	return *o.OcrType
}

// GetOcrTypeOk returns a tuple with the OcrType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) GetOcrTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OcrType) {
		return nil, false
	}
	return o.OcrType, true
}

// HasOcrType returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) HasOcrType() bool {
	if o != nil && !IsNil(o.OcrType) {
		return true
	}

	return false
}

// SetOcrType gets a reference to the given string and assigns it to the OcrType field.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) SetOcrType(v string) {
	o.OcrType = &v
}

// GetOuterOrderNo returns the OuterOrderNo field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) GetOuterOrderNo() string {
	if o == nil || IsNil(o.OuterOrderNo) {
		var ret string
		return ret
	}
	return *o.OuterOrderNo
}

// GetOuterOrderNoOk returns a tuple with the OuterOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) GetOuterOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OuterOrderNo) {
		return nil, false
	}
	return o.OuterOrderNo, true
}

// HasOuterOrderNo returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) HasOuterOrderNo() bool {
	if o != nil && !IsNil(o.OuterOrderNo) {
		return true
	}

	return false
}

// SetOuterOrderNo gets a reference to the given string and assigns it to the OuterOrderNo field.
func (o *DatadigitalFincloudGeneralsaasOcrServerDetectModel) SetOuterOrderNo(v string) {
	o.OuterOrderNo = &v
}

func (o DatadigitalFincloudGeneralsaasOcrServerDetectModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatadigitalFincloudGeneralsaasOcrServerDetectModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OcrType) {
		toSerialize["ocr_type"] = o.OcrType
	}
	if !IsNil(o.OuterOrderNo) {
		toSerialize["outer_order_no"] = o.OuterOrderNo
	}
	return toSerialize, nil
}

type NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel struct {
	value *DatadigitalFincloudGeneralsaasOcrServerDetectModel
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel) Get() *DatadigitalFincloudGeneralsaasOcrServerDetectModel {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel) Set(val *DatadigitalFincloudGeneralsaasOcrServerDetectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasOcrServerDetectModel(val *DatadigitalFincloudGeneralsaasOcrServerDetectModel) *NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel {
	return &NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


