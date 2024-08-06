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

// checks if the FulfillmentDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentDetail{}

// FulfillmentDetail struct for FulfillmentDetail
type FulfillmentDetail struct {
	// 履约金额
	FulfillmentAmount *string `json:"fulfillment_amount,omitempty"`
	// 履约支付时间
	GmtPayment *string `json:"gmt_payment,omitempty"`
	// 商户发起履约请求时，传入的out_request_no，标识一次请求的唯一id
	OutRequestNo *string `json:"out_request_no,omitempty"`
}

// NewFulfillmentDetail instantiates a new FulfillmentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentDetail() *FulfillmentDetail {
	this := FulfillmentDetail{}
	return &this
}

// NewFulfillmentDetailWithDefaults instantiates a new FulfillmentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentDetailWithDefaults() *FulfillmentDetail {
	this := FulfillmentDetail{}
	return &this
}

// GetFulfillmentAmount returns the FulfillmentAmount field value if set, zero value otherwise.
func (o *FulfillmentDetail) GetFulfillmentAmount() string {
	if o == nil || IsNil(o.FulfillmentAmount) {
		var ret string
		return ret
	}
	return *o.FulfillmentAmount
}

// GetFulfillmentAmountOk returns a tuple with the FulfillmentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDetail) GetFulfillmentAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentAmount) {
		return nil, false
	}
	return o.FulfillmentAmount, true
}

// HasFulfillmentAmount returns a boolean if a field has been set.
func (o *FulfillmentDetail) HasFulfillmentAmount() bool {
	if o != nil && !IsNil(o.FulfillmentAmount) {
		return true
	}

	return false
}

// SetFulfillmentAmount gets a reference to the given string and assigns it to the FulfillmentAmount field.
func (o *FulfillmentDetail) SetFulfillmentAmount(v string) {
	o.FulfillmentAmount = &v
}

// GetGmtPayment returns the GmtPayment field value if set, zero value otherwise.
func (o *FulfillmentDetail) GetGmtPayment() string {
	if o == nil || IsNil(o.GmtPayment) {
		var ret string
		return ret
	}
	return *o.GmtPayment
}

// GetGmtPaymentOk returns a tuple with the GmtPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDetail) GetGmtPaymentOk() (*string, bool) {
	if o == nil || IsNil(o.GmtPayment) {
		return nil, false
	}
	return o.GmtPayment, true
}

// HasGmtPayment returns a boolean if a field has been set.
func (o *FulfillmentDetail) HasGmtPayment() bool {
	if o != nil && !IsNil(o.GmtPayment) {
		return true
	}

	return false
}

// SetGmtPayment gets a reference to the given string and assigns it to the GmtPayment field.
func (o *FulfillmentDetail) SetGmtPayment(v string) {
	o.GmtPayment = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *FulfillmentDetail) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDetail) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *FulfillmentDetail) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *FulfillmentDetail) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

func (o FulfillmentDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulfillmentDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FulfillmentAmount) {
		toSerialize["fulfillment_amount"] = o.FulfillmentAmount
	}
	if !IsNil(o.GmtPayment) {
		toSerialize["gmt_payment"] = o.GmtPayment
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	return toSerialize, nil
}

type NullableFulfillmentDetail struct {
	value *FulfillmentDetail
	isSet bool
}

func (v NullableFulfillmentDetail) Get() *FulfillmentDetail {
	return v.value
}

func (v *NullableFulfillmentDetail) Set(val *FulfillmentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentDetail(val *FulfillmentDetail) *NullableFulfillmentDetail {
	return &NullableFulfillmentDetail{value: val, isSet: true}
}

func (v NullableFulfillmentDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


