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

// checks if the BillSendExtInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillSendExtInfo{}

// BillSendExtInfo struct for BillSendExtInfo
type BillSendExtInfo struct {
	// 学校外标编号
	ExtSchoolId *string `json:"ext_school_id,omitempty"`
	// 账单唤起收银台的方式： 1-手机wap;2-当面付。默认值为1（不传值或传非数字值）
	OrderPayType *string `json:"order_pay_type,omitempty"`
	// 分账金额，单位：元
	RoyaltyAmount *string `json:"royalty_amount,omitempty"`
	// 学校名称
	SchoolName *string `json:"school_name,omitempty"`
	// 分账收入户PID
	TransInPid *string `json:"trans_in_pid,omitempty"`
}

// NewBillSendExtInfo instantiates a new BillSendExtInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillSendExtInfo() *BillSendExtInfo {
	this := BillSendExtInfo{}
	return &this
}

// NewBillSendExtInfoWithDefaults instantiates a new BillSendExtInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillSendExtInfoWithDefaults() *BillSendExtInfo {
	this := BillSendExtInfo{}
	return &this
}

// GetExtSchoolId returns the ExtSchoolId field value if set, zero value otherwise.
func (o *BillSendExtInfo) GetExtSchoolId() string {
	if o == nil || IsNil(o.ExtSchoolId) {
		var ret string
		return ret
	}
	return *o.ExtSchoolId
}

// GetExtSchoolIdOk returns a tuple with the ExtSchoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillSendExtInfo) GetExtSchoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExtSchoolId) {
		return nil, false
	}
	return o.ExtSchoolId, true
}

// HasExtSchoolId returns a boolean if a field has been set.
func (o *BillSendExtInfo) HasExtSchoolId() bool {
	if o != nil && !IsNil(o.ExtSchoolId) {
		return true
	}

	return false
}

// SetExtSchoolId gets a reference to the given string and assigns it to the ExtSchoolId field.
func (o *BillSendExtInfo) SetExtSchoolId(v string) {
	o.ExtSchoolId = &v
}

// GetOrderPayType returns the OrderPayType field value if set, zero value otherwise.
func (o *BillSendExtInfo) GetOrderPayType() string {
	if o == nil || IsNil(o.OrderPayType) {
		var ret string
		return ret
	}
	return *o.OrderPayType
}

// GetOrderPayTypeOk returns a tuple with the OrderPayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillSendExtInfo) GetOrderPayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderPayType) {
		return nil, false
	}
	return o.OrderPayType, true
}

// HasOrderPayType returns a boolean if a field has been set.
func (o *BillSendExtInfo) HasOrderPayType() bool {
	if o != nil && !IsNil(o.OrderPayType) {
		return true
	}

	return false
}

// SetOrderPayType gets a reference to the given string and assigns it to the OrderPayType field.
func (o *BillSendExtInfo) SetOrderPayType(v string) {
	o.OrderPayType = &v
}

// GetRoyaltyAmount returns the RoyaltyAmount field value if set, zero value otherwise.
func (o *BillSendExtInfo) GetRoyaltyAmount() string {
	if o == nil || IsNil(o.RoyaltyAmount) {
		var ret string
		return ret
	}
	return *o.RoyaltyAmount
}

// GetRoyaltyAmountOk returns a tuple with the RoyaltyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillSendExtInfo) GetRoyaltyAmountOk() (*string, bool) {
	if o == nil || IsNil(o.RoyaltyAmount) {
		return nil, false
	}
	return o.RoyaltyAmount, true
}

// HasRoyaltyAmount returns a boolean if a field has been set.
func (o *BillSendExtInfo) HasRoyaltyAmount() bool {
	if o != nil && !IsNil(o.RoyaltyAmount) {
		return true
	}

	return false
}

// SetRoyaltyAmount gets a reference to the given string and assigns it to the RoyaltyAmount field.
func (o *BillSendExtInfo) SetRoyaltyAmount(v string) {
	o.RoyaltyAmount = &v
}

// GetSchoolName returns the SchoolName field value if set, zero value otherwise.
func (o *BillSendExtInfo) GetSchoolName() string {
	if o == nil || IsNil(o.SchoolName) {
		var ret string
		return ret
	}
	return *o.SchoolName
}

// GetSchoolNameOk returns a tuple with the SchoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillSendExtInfo) GetSchoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchoolName) {
		return nil, false
	}
	return o.SchoolName, true
}

// HasSchoolName returns a boolean if a field has been set.
func (o *BillSendExtInfo) HasSchoolName() bool {
	if o != nil && !IsNil(o.SchoolName) {
		return true
	}

	return false
}

// SetSchoolName gets a reference to the given string and assigns it to the SchoolName field.
func (o *BillSendExtInfo) SetSchoolName(v string) {
	o.SchoolName = &v
}

// GetTransInPid returns the TransInPid field value if set, zero value otherwise.
func (o *BillSendExtInfo) GetTransInPid() string {
	if o == nil || IsNil(o.TransInPid) {
		var ret string
		return ret
	}
	return *o.TransInPid
}

// GetTransInPidOk returns a tuple with the TransInPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillSendExtInfo) GetTransInPidOk() (*string, bool) {
	if o == nil || IsNil(o.TransInPid) {
		return nil, false
	}
	return o.TransInPid, true
}

// HasTransInPid returns a boolean if a field has been set.
func (o *BillSendExtInfo) HasTransInPid() bool {
	if o != nil && !IsNil(o.TransInPid) {
		return true
	}

	return false
}

// SetTransInPid gets a reference to the given string and assigns it to the TransInPid field.
func (o *BillSendExtInfo) SetTransInPid(v string) {
	o.TransInPid = &v
}

func (o BillSendExtInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillSendExtInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtSchoolId) {
		toSerialize["ext_school_id"] = o.ExtSchoolId
	}
	if !IsNil(o.OrderPayType) {
		toSerialize["order_pay_type"] = o.OrderPayType
	}
	if !IsNil(o.RoyaltyAmount) {
		toSerialize["royalty_amount"] = o.RoyaltyAmount
	}
	if !IsNil(o.SchoolName) {
		toSerialize["school_name"] = o.SchoolName
	}
	if !IsNil(o.TransInPid) {
		toSerialize["trans_in_pid"] = o.TransInPid
	}
	return toSerialize, nil
}

type NullableBillSendExtInfo struct {
	value *BillSendExtInfo
	isSet bool
}

func (v NullableBillSendExtInfo) Get() *BillSendExtInfo {
	return v.value
}

func (v *NullableBillSendExtInfo) Set(val *BillSendExtInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillSendExtInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillSendExtInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillSendExtInfo(val *BillSendExtInfo) *NullableBillSendExtInfo {
	return &NullableBillSendExtInfo{value: val, isSet: true}
}

func (v NullableBillSendExtInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillSendExtInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
