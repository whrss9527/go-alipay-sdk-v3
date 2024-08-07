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

// checks if the OcrPlaneScanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OcrPlaneScanInfo{}

// OcrPlaneScanInfo struct for OcrPlaneScanInfo
type OcrPlaneScanInfo struct {
	// 目的地
	Destination *string `json:"destination,omitempty"`
	// 航班号
	FlightNo *string `json:"flight_no,omitempty"`
	// 乘机日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// 出发地
	Origin *string `json:"origin,omitempty"`
	// 乘客
	Passenger *string `json:"passenger,omitempty"`
	// 金额（元）
	Price *string `json:"price,omitempty"`
	// 明细事由
	Remark *string `json:"remark,omitempty"`
	// 飞机舱位
	SeatClass *string `json:"seat_class,omitempty"`
}

// NewOcrPlaneScanInfo instantiates a new OcrPlaneScanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOcrPlaneScanInfo() *OcrPlaneScanInfo {
	this := OcrPlaneScanInfo{}
	return &this
}

// NewOcrPlaneScanInfoWithDefaults instantiates a new OcrPlaneScanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOcrPlaneScanInfoWithDefaults() *OcrPlaneScanInfo {
	this := OcrPlaneScanInfo{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *OcrPlaneScanInfo) SetDestination(v string) {
	o.Destination = &v
}

// GetFlightNo returns the FlightNo field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetFlightNo() string {
	if o == nil || IsNil(o.FlightNo) {
		var ret string
		return ret
	}
	return *o.FlightNo
}

// GetFlightNoOk returns a tuple with the FlightNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetFlightNoOk() (*string, bool) {
	if o == nil || IsNil(o.FlightNo) {
		return nil, false
	}
	return o.FlightNo, true
}

// HasFlightNo returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasFlightNo() bool {
	if o != nil && !IsNil(o.FlightNo) {
		return true
	}

	return false
}

// SetFlightNo gets a reference to the given string and assigns it to the FlightNo field.
func (o *OcrPlaneScanInfo) SetFlightNo(v string) {
	o.FlightNo = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *OcrPlaneScanInfo) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *OcrPlaneScanInfo) SetOrigin(v string) {
	o.Origin = &v
}

// GetPassenger returns the Passenger field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetPassenger() string {
	if o == nil || IsNil(o.Passenger) {
		var ret string
		return ret
	}
	return *o.Passenger
}

// GetPassengerOk returns a tuple with the Passenger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetPassengerOk() (*string, bool) {
	if o == nil || IsNil(o.Passenger) {
		return nil, false
	}
	return o.Passenger, true
}

// HasPassenger returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasPassenger() bool {
	if o != nil && !IsNil(o.Passenger) {
		return true
	}

	return false
}

// SetPassenger gets a reference to the given string and assigns it to the Passenger field.
func (o *OcrPlaneScanInfo) SetPassenger(v string) {
	o.Passenger = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OcrPlaneScanInfo) SetPrice(v string) {
	o.Price = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *OcrPlaneScanInfo) SetRemark(v string) {
	o.Remark = &v
}

// GetSeatClass returns the SeatClass field value if set, zero value otherwise.
func (o *OcrPlaneScanInfo) GetSeatClass() string {
	if o == nil || IsNil(o.SeatClass) {
		var ret string
		return ret
	}
	return *o.SeatClass
}

// GetSeatClassOk returns a tuple with the SeatClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrPlaneScanInfo) GetSeatClassOk() (*string, bool) {
	if o == nil || IsNil(o.SeatClass) {
		return nil, false
	}
	return o.SeatClass, true
}

// HasSeatClass returns a boolean if a field has been set.
func (o *OcrPlaneScanInfo) HasSeatClass() bool {
	if o != nil && !IsNil(o.SeatClass) {
		return true
	}

	return false
}

// SetSeatClass gets a reference to the given string and assigns it to the SeatClass field.
func (o *OcrPlaneScanInfo) SetSeatClass(v string) {
	o.SeatClass = &v
}

func (o OcrPlaneScanInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OcrPlaneScanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.FlightNo) {
		toSerialize["flight_no"] = o.FlightNo
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoice_date"] = o.InvoiceDate
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Passenger) {
		toSerialize["passenger"] = o.Passenger
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.SeatClass) {
		toSerialize["seat_class"] = o.SeatClass
	}
	return toSerialize, nil
}

type NullableOcrPlaneScanInfo struct {
	value *OcrPlaneScanInfo
	isSet bool
}

func (v NullableOcrPlaneScanInfo) Get() *OcrPlaneScanInfo {
	return v.value
}

func (v *NullableOcrPlaneScanInfo) Set(val *OcrPlaneScanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOcrPlaneScanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOcrPlaneScanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOcrPlaneScanInfo(val *OcrPlaneScanInfo) *NullableOcrPlaneScanInfo {
	return &NullableOcrPlaneScanInfo{value: val, isSet: true}
}

func (v NullableOcrPlaneScanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOcrPlaneScanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
