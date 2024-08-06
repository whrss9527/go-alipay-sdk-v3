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

// checks if the AlipayEcoEduKtStudentModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoEduKtStudentModifyModel{}

// AlipayEcoEduKtStudentModifyModel struct for AlipayEcoEduKtStudentModifyModel
type AlipayEcoEduKtStudentModifyModel struct {
	// 修改后的学生姓名  本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	ChildName *string `json:"child_name,omitempty"`
	// 已经签约教育缴费的isv的支付宝PID
	IsvPid *string `json:"isv_pid,omitempty"`
	// 学校编号，调用alipay.eco.edu.kt.schoolinfo.modify接口录入学校信息时，接口返回的编号
	SchoolNo *string `json:"school_no,omitempty"`
	// 学校用来签约支付宝教育缴费的支付宝PID
	SchoolPid *string `json:"school_pid,omitempty"`
	// 区分ISV操作，“D”表示删除，“U”表示更新，区分大小写。  如果为U，则学生名字，学号，身份证至少填写一项
	Status *string `json:"status,omitempty"`
	// 修改后的学号  本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	StudentCode *string `json:"student_code,omitempty"`
	// 修改后的身份证号码  本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	StudentIdentify *string `json:"student_identify,omitempty"`
	// 支付宝-中小学-教育缴费生成的学生唯一编号
	StudentNo *string `json:"student_no,omitempty"`
	// 孩子的家长信息，最多一次输入20个家长。如果输入的家长信息不存在，则该改学生增加家长信息  本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	Users []UserDetails `json:"users,omitempty"`
}

// NewAlipayEcoEduKtStudentModifyModel instantiates a new AlipayEcoEduKtStudentModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoEduKtStudentModifyModel() *AlipayEcoEduKtStudentModifyModel {
	this := AlipayEcoEduKtStudentModifyModel{}
	return &this
}

// NewAlipayEcoEduKtStudentModifyModelWithDefaults instantiates a new AlipayEcoEduKtStudentModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoEduKtStudentModifyModelWithDefaults() *AlipayEcoEduKtStudentModifyModel {
	this := AlipayEcoEduKtStudentModifyModel{}
	return &this
}

// GetChildName returns the ChildName field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetChildName() string {
	if o == nil || IsNil(o.ChildName) {
		var ret string
		return ret
	}
	return *o.ChildName
}

// GetChildNameOk returns a tuple with the ChildName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetChildNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChildName) {
		return nil, false
	}
	return o.ChildName, true
}

// HasChildName returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasChildName() bool {
	if o != nil && !IsNil(o.ChildName) {
		return true
	}

	return false
}

// SetChildName gets a reference to the given string and assigns it to the ChildName field.
func (o *AlipayEcoEduKtStudentModifyModel) SetChildName(v string) {
	o.ChildName = &v
}

// GetIsvPid returns the IsvPid field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetIsvPid() string {
	if o == nil || IsNil(o.IsvPid) {
		var ret string
		return ret
	}
	return *o.IsvPid
}

// GetIsvPidOk returns a tuple with the IsvPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetIsvPidOk() (*string, bool) {
	if o == nil || IsNil(o.IsvPid) {
		return nil, false
	}
	return o.IsvPid, true
}

// HasIsvPid returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasIsvPid() bool {
	if o != nil && !IsNil(o.IsvPid) {
		return true
	}

	return false
}

// SetIsvPid gets a reference to the given string and assigns it to the IsvPid field.
func (o *AlipayEcoEduKtStudentModifyModel) SetIsvPid(v string) {
	o.IsvPid = &v
}

// GetSchoolNo returns the SchoolNo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetSchoolNo() string {
	if o == nil || IsNil(o.SchoolNo) {
		var ret string
		return ret
	}
	return *o.SchoolNo
}

// GetSchoolNoOk returns a tuple with the SchoolNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetSchoolNoOk() (*string, bool) {
	if o == nil || IsNil(o.SchoolNo) {
		return nil, false
	}
	return o.SchoolNo, true
}

// HasSchoolNo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasSchoolNo() bool {
	if o != nil && !IsNil(o.SchoolNo) {
		return true
	}

	return false
}

// SetSchoolNo gets a reference to the given string and assigns it to the SchoolNo field.
func (o *AlipayEcoEduKtStudentModifyModel) SetSchoolNo(v string) {
	o.SchoolNo = &v
}

// GetSchoolPid returns the SchoolPid field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetSchoolPid() string {
	if o == nil || IsNil(o.SchoolPid) {
		var ret string
		return ret
	}
	return *o.SchoolPid
}

// GetSchoolPidOk returns a tuple with the SchoolPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetSchoolPidOk() (*string, bool) {
	if o == nil || IsNil(o.SchoolPid) {
		return nil, false
	}
	return o.SchoolPid, true
}

// HasSchoolPid returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasSchoolPid() bool {
	if o != nil && !IsNil(o.SchoolPid) {
		return true
	}

	return false
}

// SetSchoolPid gets a reference to the given string and assigns it to the SchoolPid field.
func (o *AlipayEcoEduKtStudentModifyModel) SetSchoolPid(v string) {
	o.SchoolPid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayEcoEduKtStudentModifyModel) SetStatus(v string) {
	o.Status = &v
}

// GetStudentCode returns the StudentCode field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetStudentCode() string {
	if o == nil || IsNil(o.StudentCode) {
		var ret string
		return ret
	}
	return *o.StudentCode
}

// GetStudentCodeOk returns a tuple with the StudentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetStudentCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StudentCode) {
		return nil, false
	}
	return o.StudentCode, true
}

// HasStudentCode returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasStudentCode() bool {
	if o != nil && !IsNil(o.StudentCode) {
		return true
	}

	return false
}

// SetStudentCode gets a reference to the given string and assigns it to the StudentCode field.
func (o *AlipayEcoEduKtStudentModifyModel) SetStudentCode(v string) {
	o.StudentCode = &v
}

// GetStudentIdentify returns the StudentIdentify field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetStudentIdentify() string {
	if o == nil || IsNil(o.StudentIdentify) {
		var ret string
		return ret
	}
	return *o.StudentIdentify
}

// GetStudentIdentifyOk returns a tuple with the StudentIdentify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetStudentIdentifyOk() (*string, bool) {
	if o == nil || IsNil(o.StudentIdentify) {
		return nil, false
	}
	return o.StudentIdentify, true
}

// HasStudentIdentify returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasStudentIdentify() bool {
	if o != nil && !IsNil(o.StudentIdentify) {
		return true
	}

	return false
}

// SetStudentIdentify gets a reference to the given string and assigns it to the StudentIdentify field.
func (o *AlipayEcoEduKtStudentModifyModel) SetStudentIdentify(v string) {
	o.StudentIdentify = &v
}

// GetStudentNo returns the StudentNo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetStudentNo() string {
	if o == nil || IsNil(o.StudentNo) {
		var ret string
		return ret
	}
	return *o.StudentNo
}

// GetStudentNoOk returns a tuple with the StudentNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetStudentNoOk() (*string, bool) {
	if o == nil || IsNil(o.StudentNo) {
		return nil, false
	}
	return o.StudentNo, true
}

// HasStudentNo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasStudentNo() bool {
	if o != nil && !IsNil(o.StudentNo) {
		return true
	}

	return false
}

// SetStudentNo gets a reference to the given string and assigns it to the StudentNo field.
func (o *AlipayEcoEduKtStudentModifyModel) SetStudentNo(v string) {
	o.StudentNo = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AlipayEcoEduKtStudentModifyModel) GetUsers() []UserDetails {
	if o == nil || IsNil(o.Users) {
		var ret []UserDetails
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtStudentModifyModel) GetUsersOk() ([]UserDetails, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AlipayEcoEduKtStudentModifyModel) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserDetails and assigns it to the Users field.
func (o *AlipayEcoEduKtStudentModifyModel) SetUsers(v []UserDetails) {
	o.Users = v
}

func (o AlipayEcoEduKtStudentModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoEduKtStudentModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChildName) {
		toSerialize["child_name"] = o.ChildName
	}
	if !IsNil(o.IsvPid) {
		toSerialize["isv_pid"] = o.IsvPid
	}
	if !IsNil(o.SchoolNo) {
		toSerialize["school_no"] = o.SchoolNo
	}
	if !IsNil(o.SchoolPid) {
		toSerialize["school_pid"] = o.SchoolPid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StudentCode) {
		toSerialize["student_code"] = o.StudentCode
	}
	if !IsNil(o.StudentIdentify) {
		toSerialize["student_identify"] = o.StudentIdentify
	}
	if !IsNil(o.StudentNo) {
		toSerialize["student_no"] = o.StudentNo
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableAlipayEcoEduKtStudentModifyModel struct {
	value *AlipayEcoEduKtStudentModifyModel
	isSet bool
}

func (v NullableAlipayEcoEduKtStudentModifyModel) Get() *AlipayEcoEduKtStudentModifyModel {
	return v.value
}

func (v *NullableAlipayEcoEduKtStudentModifyModel) Set(val *AlipayEcoEduKtStudentModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoEduKtStudentModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoEduKtStudentModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoEduKtStudentModifyModel(val *AlipayEcoEduKtStudentModifyModel) *NullableAlipayEcoEduKtStudentModifyModel {
	return &NullableAlipayEcoEduKtStudentModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEcoEduKtStudentModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoEduKtStudentModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

