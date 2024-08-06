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

// checks if the CumulateDataDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CumulateDataDetail{}

// CumulateDataDetail struct for CumulateDataDetail
type CumulateDataDetail struct {
	// 数据回传的动作类型，正向或逆向。
	ActionType *string `json:"action_type,omitempty"`
	// 回传数据发生的实际时间
	BizTime *string `json:"biz_time,omitempty"`
	// 回传数据类型。
	DataType *string `json:"data_type,omitempty"`
	// 优惠金额
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// 优惠描述
	DiscountDesc *string `json:"discount_desc,omitempty"`
	// 数据回传时传入的外部业务号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 数据回传时逆向单据记录的对应正向单据的外部业务号。
	ReferOutBizNo *string `json:"refer_out_biz_no,omitempty"`
	// 回传数据子类型。
	SubDataType *string `json:"sub_data_type,omitempty"`
	// 任务金额
	TaskAmount *string `json:"task_amount,omitempty"`
	// 任务描述
	TaskDesc *string `json:"task_desc,omitempty"`
	// 任务次数
	TaskTimes *int32 `json:"task_times,omitempty"`
}

// NewCumulateDataDetail instantiates a new CumulateDataDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCumulateDataDetail() *CumulateDataDetail {
	this := CumulateDataDetail{}
	return &this
}

// NewCumulateDataDetailWithDefaults instantiates a new CumulateDataDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCumulateDataDetailWithDefaults() *CumulateDataDetail {
	this := CumulateDataDetail{}
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetActionType() string {
	if o == nil || IsNil(o.ActionType) {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetActionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *CumulateDataDetail) SetActionType(v string) {
	o.ActionType = &v
}

// GetBizTime returns the BizTime field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetBizTime() string {
	if o == nil || IsNil(o.BizTime) {
		var ret string
		return ret
	}
	return *o.BizTime
}

// GetBizTimeOk returns a tuple with the BizTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetBizTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BizTime) {
		return nil, false
	}
	return o.BizTime, true
}

// HasBizTime returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasBizTime() bool {
	if o != nil && !IsNil(o.BizTime) {
		return true
	}

	return false
}

// SetBizTime gets a reference to the given string and assigns it to the BizTime field.
func (o *CumulateDataDetail) SetBizTime(v string) {
	o.BizTime = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *CumulateDataDetail) SetDataType(v string) {
	o.DataType = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetDiscountAmount() string {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret string
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetDiscountAmountOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given string and assigns it to the DiscountAmount field.
func (o *CumulateDataDetail) SetDiscountAmount(v string) {
	o.DiscountAmount = &v
}

// GetDiscountDesc returns the DiscountDesc field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetDiscountDesc() string {
	if o == nil || IsNil(o.DiscountDesc) {
		var ret string
		return ret
	}
	return *o.DiscountDesc
}

// GetDiscountDescOk returns a tuple with the DiscountDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetDiscountDescOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountDesc) {
		return nil, false
	}
	return o.DiscountDesc, true
}

// HasDiscountDesc returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasDiscountDesc() bool {
	if o != nil && !IsNil(o.DiscountDesc) {
		return true
	}

	return false
}

// SetDiscountDesc gets a reference to the given string and assigns it to the DiscountDesc field.
func (o *CumulateDataDetail) SetDiscountDesc(v string) {
	o.DiscountDesc = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *CumulateDataDetail) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetReferOutBizNo returns the ReferOutBizNo field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetReferOutBizNo() string {
	if o == nil || IsNil(o.ReferOutBizNo) {
		var ret string
		return ret
	}
	return *o.ReferOutBizNo
}

// GetReferOutBizNoOk returns a tuple with the ReferOutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetReferOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.ReferOutBizNo) {
		return nil, false
	}
	return o.ReferOutBizNo, true
}

// HasReferOutBizNo returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasReferOutBizNo() bool {
	if o != nil && !IsNil(o.ReferOutBizNo) {
		return true
	}

	return false
}

// SetReferOutBizNo gets a reference to the given string and assigns it to the ReferOutBizNo field.
func (o *CumulateDataDetail) SetReferOutBizNo(v string) {
	o.ReferOutBizNo = &v
}

// GetSubDataType returns the SubDataType field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetSubDataType() string {
	if o == nil || IsNil(o.SubDataType) {
		var ret string
		return ret
	}
	return *o.SubDataType
}

// GetSubDataTypeOk returns a tuple with the SubDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetSubDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubDataType) {
		return nil, false
	}
	return o.SubDataType, true
}

// HasSubDataType returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasSubDataType() bool {
	if o != nil && !IsNil(o.SubDataType) {
		return true
	}

	return false
}

// SetSubDataType gets a reference to the given string and assigns it to the SubDataType field.
func (o *CumulateDataDetail) SetSubDataType(v string) {
	o.SubDataType = &v
}

// GetTaskAmount returns the TaskAmount field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetTaskAmount() string {
	if o == nil || IsNil(o.TaskAmount) {
		var ret string
		return ret
	}
	return *o.TaskAmount
}

// GetTaskAmountOk returns a tuple with the TaskAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetTaskAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TaskAmount) {
		return nil, false
	}
	return o.TaskAmount, true
}

// HasTaskAmount returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasTaskAmount() bool {
	if o != nil && !IsNil(o.TaskAmount) {
		return true
	}

	return false
}

// SetTaskAmount gets a reference to the given string and assigns it to the TaskAmount field.
func (o *CumulateDataDetail) SetTaskAmount(v string) {
	o.TaskAmount = &v
}

// GetTaskDesc returns the TaskDesc field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetTaskDesc() string {
	if o == nil || IsNil(o.TaskDesc) {
		var ret string
		return ret
	}
	return *o.TaskDesc
}

// GetTaskDescOk returns a tuple with the TaskDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetTaskDescOk() (*string, bool) {
	if o == nil || IsNil(o.TaskDesc) {
		return nil, false
	}
	return o.TaskDesc, true
}

// HasTaskDesc returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasTaskDesc() bool {
	if o != nil && !IsNil(o.TaskDesc) {
		return true
	}

	return false
}

// SetTaskDesc gets a reference to the given string and assigns it to the TaskDesc field.
func (o *CumulateDataDetail) SetTaskDesc(v string) {
	o.TaskDesc = &v
}

// GetTaskTimes returns the TaskTimes field value if set, zero value otherwise.
func (o *CumulateDataDetail) GetTaskTimes() int32 {
	if o == nil || IsNil(o.TaskTimes) {
		var ret int32
		return ret
	}
	return *o.TaskTimes
}

// GetTaskTimesOk returns a tuple with the TaskTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CumulateDataDetail) GetTaskTimesOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskTimes) {
		return nil, false
	}
	return o.TaskTimes, true
}

// HasTaskTimes returns a boolean if a field has been set.
func (o *CumulateDataDetail) HasTaskTimes() bool {
	if o != nil && !IsNil(o.TaskTimes) {
		return true
	}

	return false
}

// SetTaskTimes gets a reference to the given int32 and assigns it to the TaskTimes field.
func (o *CumulateDataDetail) SetTaskTimes(v int32) {
	o.TaskTimes = &v
}

func (o CumulateDataDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CumulateDataDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionType) {
		toSerialize["action_type"] = o.ActionType
	}
	if !IsNil(o.BizTime) {
		toSerialize["biz_time"] = o.BizTime
	}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountDesc) {
		toSerialize["discount_desc"] = o.DiscountDesc
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.ReferOutBizNo) {
		toSerialize["refer_out_biz_no"] = o.ReferOutBizNo
	}
	if !IsNil(o.SubDataType) {
		toSerialize["sub_data_type"] = o.SubDataType
	}
	if !IsNil(o.TaskAmount) {
		toSerialize["task_amount"] = o.TaskAmount
	}
	if !IsNil(o.TaskDesc) {
		toSerialize["task_desc"] = o.TaskDesc
	}
	if !IsNil(o.TaskTimes) {
		toSerialize["task_times"] = o.TaskTimes
	}
	return toSerialize, nil
}

type NullableCumulateDataDetail struct {
	value *CumulateDataDetail
	isSet bool
}

func (v NullableCumulateDataDetail) Get() *CumulateDataDetail {
	return v.value
}

func (v *NullableCumulateDataDetail) Set(val *CumulateDataDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCumulateDataDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCumulateDataDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCumulateDataDetail(val *CumulateDataDetail) *NullableCumulateDataDetail {
	return &NullableCumulateDataDetail{value: val, isSet: true}
}

func (v NullableCumulateDataDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCumulateDataDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


