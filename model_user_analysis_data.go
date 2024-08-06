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

// checks if the UserAnalysisData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAnalysisData{}

// UserAnalysisData struct for UserAnalysisData
type UserAnalysisData struct {
	// 取消关注人数
	CancelUserCnt *string `json:"cancel_user_cnt,omitempty"`
	// 累积关注人数
	CumulateUserCnt *string `json:"cumulate_user_cnt,omitempty"`
	// 日期
	Date *string `json:"date,omitempty"`
	// 净增关注人数
	GrowUserCnt *string `json:"grow_user_cnt,omitempty"`
	// 新关注人数
	NewUserCnt *string `json:"new_user_cnt,omitempty"`
}

// NewUserAnalysisData instantiates a new UserAnalysisData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAnalysisData() *UserAnalysisData {
	this := UserAnalysisData{}
	return &this
}

// NewUserAnalysisDataWithDefaults instantiates a new UserAnalysisData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAnalysisDataWithDefaults() *UserAnalysisData {
	this := UserAnalysisData{}
	return &this
}

// GetCancelUserCnt returns the CancelUserCnt field value if set, zero value otherwise.
func (o *UserAnalysisData) GetCancelUserCnt() string {
	if o == nil || IsNil(o.CancelUserCnt) {
		var ret string
		return ret
	}
	return *o.CancelUserCnt
}

// GetCancelUserCntOk returns a tuple with the CancelUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAnalysisData) GetCancelUserCntOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUserCnt) {
		return nil, false
	}
	return o.CancelUserCnt, true
}

// HasCancelUserCnt returns a boolean if a field has been set.
func (o *UserAnalysisData) HasCancelUserCnt() bool {
	if o != nil && !IsNil(o.CancelUserCnt) {
		return true
	}

	return false
}

// SetCancelUserCnt gets a reference to the given string and assigns it to the CancelUserCnt field.
func (o *UserAnalysisData) SetCancelUserCnt(v string) {
	o.CancelUserCnt = &v
}

// GetCumulateUserCnt returns the CumulateUserCnt field value if set, zero value otherwise.
func (o *UserAnalysisData) GetCumulateUserCnt() string {
	if o == nil || IsNil(o.CumulateUserCnt) {
		var ret string
		return ret
	}
	return *o.CumulateUserCnt
}

// GetCumulateUserCntOk returns a tuple with the CumulateUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAnalysisData) GetCumulateUserCntOk() (*string, bool) {
	if o == nil || IsNil(o.CumulateUserCnt) {
		return nil, false
	}
	return o.CumulateUserCnt, true
}

// HasCumulateUserCnt returns a boolean if a field has been set.
func (o *UserAnalysisData) HasCumulateUserCnt() bool {
	if o != nil && !IsNil(o.CumulateUserCnt) {
		return true
	}

	return false
}

// SetCumulateUserCnt gets a reference to the given string and assigns it to the CumulateUserCnt field.
func (o *UserAnalysisData) SetCumulateUserCnt(v string) {
	o.CumulateUserCnt = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *UserAnalysisData) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAnalysisData) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *UserAnalysisData) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *UserAnalysisData) SetDate(v string) {
	o.Date = &v
}

// GetGrowUserCnt returns the GrowUserCnt field value if set, zero value otherwise.
func (o *UserAnalysisData) GetGrowUserCnt() string {
	if o == nil || IsNil(o.GrowUserCnt) {
		var ret string
		return ret
	}
	return *o.GrowUserCnt
}

// GetGrowUserCntOk returns a tuple with the GrowUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAnalysisData) GetGrowUserCntOk() (*string, bool) {
	if o == nil || IsNil(o.GrowUserCnt) {
		return nil, false
	}
	return o.GrowUserCnt, true
}

// HasGrowUserCnt returns a boolean if a field has been set.
func (o *UserAnalysisData) HasGrowUserCnt() bool {
	if o != nil && !IsNil(o.GrowUserCnt) {
		return true
	}

	return false
}

// SetGrowUserCnt gets a reference to the given string and assigns it to the GrowUserCnt field.
func (o *UserAnalysisData) SetGrowUserCnt(v string) {
	o.GrowUserCnt = &v
}

// GetNewUserCnt returns the NewUserCnt field value if set, zero value otherwise.
func (o *UserAnalysisData) GetNewUserCnt() string {
	if o == nil || IsNil(o.NewUserCnt) {
		var ret string
		return ret
	}
	return *o.NewUserCnt
}

// GetNewUserCntOk returns a tuple with the NewUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAnalysisData) GetNewUserCntOk() (*string, bool) {
	if o == nil || IsNil(o.NewUserCnt) {
		return nil, false
	}
	return o.NewUserCnt, true
}

// HasNewUserCnt returns a boolean if a field has been set.
func (o *UserAnalysisData) HasNewUserCnt() bool {
	if o != nil && !IsNil(o.NewUserCnt) {
		return true
	}

	return false
}

// SetNewUserCnt gets a reference to the given string and assigns it to the NewUserCnt field.
func (o *UserAnalysisData) SetNewUserCnt(v string) {
	o.NewUserCnt = &v
}

func (o UserAnalysisData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAnalysisData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancelUserCnt) {
		toSerialize["cancel_user_cnt"] = o.CancelUserCnt
	}
	if !IsNil(o.CumulateUserCnt) {
		toSerialize["cumulate_user_cnt"] = o.CumulateUserCnt
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.GrowUserCnt) {
		toSerialize["grow_user_cnt"] = o.GrowUserCnt
	}
	if !IsNil(o.NewUserCnt) {
		toSerialize["new_user_cnt"] = o.NewUserCnt
	}
	return toSerialize, nil
}

type NullableUserAnalysisData struct {
	value *UserAnalysisData
	isSet bool
}

func (v NullableUserAnalysisData) Get() *UserAnalysisData {
	return v.value
}

func (v *NullableUserAnalysisData) Set(val *UserAnalysisData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAnalysisData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAnalysisData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAnalysisData(val *UserAnalysisData) *NullableUserAnalysisData {
	return &NullableUserAnalysisData{value: val, isSet: true}
}

func (v NullableUserAnalysisData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAnalysisData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
