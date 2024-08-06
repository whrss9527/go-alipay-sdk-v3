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

// checks if the DatadigitalFincloudGeneralsaasFaceCheckInitializeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatadigitalFincloudGeneralsaasFaceCheckInitializeModel{}

// DatadigitalFincloudGeneralsaasFaceCheckInitializeModel struct for DatadigitalFincloudGeneralsaasFaceCheckInitializeModel
type DatadigitalFincloudGeneralsaasFaceCheckInitializeModel struct {
	// DATA_DIGITAL_BIZ_CODE_FACE_CHECK_LIVE，代表活体检测。
	BizCode *string `json:"biz_code,omitempty"`
	// 客户业务单据号。
	OuterOrderNo *string `json:"outer_order_no,omitempty"`
}

// NewDatadigitalFincloudGeneralsaasFaceCheckInitializeModel instantiates a new DatadigitalFincloudGeneralsaasFaceCheckInitializeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadigitalFincloudGeneralsaasFaceCheckInitializeModel() *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel {
	this := DatadigitalFincloudGeneralsaasFaceCheckInitializeModel{}
	return &this
}

// NewDatadigitalFincloudGeneralsaasFaceCheckInitializeModelWithDefaults instantiates a new DatadigitalFincloudGeneralsaasFaceCheckInitializeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadigitalFincloudGeneralsaasFaceCheckInitializeModelWithDefaults() *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel {
	this := DatadigitalFincloudGeneralsaasFaceCheckInitializeModel{}
	return &this
}

// GetBizCode returns the BizCode field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) GetBizCode() string {
	if o == nil || IsNil(o.BizCode) {
		var ret string
		return ret
	}
	return *o.BizCode
}

// GetBizCodeOk returns a tuple with the BizCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) GetBizCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BizCode) {
		return nil, false
	}
	return o.BizCode, true
}

// HasBizCode returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) HasBizCode() bool {
	if o != nil && !IsNil(o.BizCode) {
		return true
	}

	return false
}

// SetBizCode gets a reference to the given string and assigns it to the BizCode field.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) SetBizCode(v string) {
	o.BizCode = &v
}

// GetOuterOrderNo returns the OuterOrderNo field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) GetOuterOrderNo() string {
	if o == nil || IsNil(o.OuterOrderNo) {
		var ret string
		return ret
	}
	return *o.OuterOrderNo
}

// GetOuterOrderNoOk returns a tuple with the OuterOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) GetOuterOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OuterOrderNo) {
		return nil, false
	}
	return o.OuterOrderNo, true
}

// HasOuterOrderNo returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) HasOuterOrderNo() bool {
	if o != nil && !IsNil(o.OuterOrderNo) {
		return true
	}

	return false
}

// SetOuterOrderNo gets a reference to the given string and assigns it to the OuterOrderNo field.
func (o *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) SetOuterOrderNo(v string) {
	o.OuterOrderNo = &v
}

func (o DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizCode) {
		toSerialize["biz_code"] = o.BizCode
	}
	if !IsNil(o.OuterOrderNo) {
		toSerialize["outer_order_no"] = o.OuterOrderNo
	}
	return toSerialize, nil
}

type NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel struct {
	value *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel) Get() *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel) Set(val *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel(val *DatadigitalFincloudGeneralsaasFaceCheckInitializeModel) *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel {
	return &NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckInitializeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
