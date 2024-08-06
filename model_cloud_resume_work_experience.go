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

// checks if the CloudResumeWorkExperience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudResumeWorkExperience{}

// CloudResumeWorkExperience struct for CloudResumeWorkExperience
type CloudResumeWorkExperience struct {
	// 公司名称
	CompanyName *string `json:"company_name,omitempty"`
	// 职业id
	JobId *string `json:"job_id,omitempty"`
	// 职业
	JobName *string `json:"job_name,omitempty"`
	// 用户在求职意向页面自主录入
	PositionName *string `json:"position_name,omitempty"`
	// 行业id
	ProfessionId *string `json:"profession_id,omitempty"`
	// 行业名称
	ProfessionName *string `json:"profession_name,omitempty"`
	// 工作描述
	WorkDesc *string `json:"work_desc,omitempty"`
	// 工作结束时间
	WorkEndTime *int32 `json:"work_end_time,omitempty"`
	// 工作开始日期
	WorkStartTime *int32 `json:"work_start_time,omitempty"`
}

// NewCloudResumeWorkExperience instantiates a new CloudResumeWorkExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudResumeWorkExperience() *CloudResumeWorkExperience {
	this := CloudResumeWorkExperience{}
	return &this
}

// NewCloudResumeWorkExperienceWithDefaults instantiates a new CloudResumeWorkExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudResumeWorkExperienceWithDefaults() *CloudResumeWorkExperience {
	this := CloudResumeWorkExperience{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *CloudResumeWorkExperience) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *CloudResumeWorkExperience) SetJobId(v string) {
	o.JobId = &v
}

// GetJobName returns the JobName field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetJobName() string {
	if o == nil || IsNil(o.JobName) {
		var ret string
		return ret
	}
	return *o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetJobNameOk() (*string, bool) {
	if o == nil || IsNil(o.JobName) {
		return nil, false
	}
	return o.JobName, true
}

// HasJobName returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasJobName() bool {
	if o != nil && !IsNil(o.JobName) {
		return true
	}

	return false
}

// SetJobName gets a reference to the given string and assigns it to the JobName field.
func (o *CloudResumeWorkExperience) SetJobName(v string) {
	o.JobName = &v
}

// GetPositionName returns the PositionName field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetPositionName() string {
	if o == nil || IsNil(o.PositionName) {
		var ret string
		return ret
	}
	return *o.PositionName
}

// GetPositionNameOk returns a tuple with the PositionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetPositionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PositionName) {
		return nil, false
	}
	return o.PositionName, true
}

// HasPositionName returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasPositionName() bool {
	if o != nil && !IsNil(o.PositionName) {
		return true
	}

	return false
}

// SetPositionName gets a reference to the given string and assigns it to the PositionName field.
func (o *CloudResumeWorkExperience) SetPositionName(v string) {
	o.PositionName = &v
}

// GetProfessionId returns the ProfessionId field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetProfessionId() string {
	if o == nil || IsNil(o.ProfessionId) {
		var ret string
		return ret
	}
	return *o.ProfessionId
}

// GetProfessionIdOk returns a tuple with the ProfessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetProfessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfessionId) {
		return nil, false
	}
	return o.ProfessionId, true
}

// HasProfessionId returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasProfessionId() bool {
	if o != nil && !IsNil(o.ProfessionId) {
		return true
	}

	return false
}

// SetProfessionId gets a reference to the given string and assigns it to the ProfessionId field.
func (o *CloudResumeWorkExperience) SetProfessionId(v string) {
	o.ProfessionId = &v
}

// GetProfessionName returns the ProfessionName field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetProfessionName() string {
	if o == nil || IsNil(o.ProfessionName) {
		var ret string
		return ret
	}
	return *o.ProfessionName
}

// GetProfessionNameOk returns a tuple with the ProfessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetProfessionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProfessionName) {
		return nil, false
	}
	return o.ProfessionName, true
}

// HasProfessionName returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasProfessionName() bool {
	if o != nil && !IsNil(o.ProfessionName) {
		return true
	}

	return false
}

// SetProfessionName gets a reference to the given string and assigns it to the ProfessionName field.
func (o *CloudResumeWorkExperience) SetProfessionName(v string) {
	o.ProfessionName = &v
}

// GetWorkDesc returns the WorkDesc field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetWorkDesc() string {
	if o == nil || IsNil(o.WorkDesc) {
		var ret string
		return ret
	}
	return *o.WorkDesc
}

// GetWorkDescOk returns a tuple with the WorkDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetWorkDescOk() (*string, bool) {
	if o == nil || IsNil(o.WorkDesc) {
		return nil, false
	}
	return o.WorkDesc, true
}

// HasWorkDesc returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasWorkDesc() bool {
	if o != nil && !IsNil(o.WorkDesc) {
		return true
	}

	return false
}

// SetWorkDesc gets a reference to the given string and assigns it to the WorkDesc field.
func (o *CloudResumeWorkExperience) SetWorkDesc(v string) {
	o.WorkDesc = &v
}

// GetWorkEndTime returns the WorkEndTime field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetWorkEndTime() int32 {
	if o == nil || IsNil(o.WorkEndTime) {
		var ret int32
		return ret
	}
	return *o.WorkEndTime
}

// GetWorkEndTimeOk returns a tuple with the WorkEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetWorkEndTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WorkEndTime) {
		return nil, false
	}
	return o.WorkEndTime, true
}

// HasWorkEndTime returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasWorkEndTime() bool {
	if o != nil && !IsNil(o.WorkEndTime) {
		return true
	}

	return false
}

// SetWorkEndTime gets a reference to the given int32 and assigns it to the WorkEndTime field.
func (o *CloudResumeWorkExperience) SetWorkEndTime(v int32) {
	o.WorkEndTime = &v
}

// GetWorkStartTime returns the WorkStartTime field value if set, zero value otherwise.
func (o *CloudResumeWorkExperience) GetWorkStartTime() int32 {
	if o == nil || IsNil(o.WorkStartTime) {
		var ret int32
		return ret
	}
	return *o.WorkStartTime
}

// GetWorkStartTimeOk returns a tuple with the WorkStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeWorkExperience) GetWorkStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WorkStartTime) {
		return nil, false
	}
	return o.WorkStartTime, true
}

// HasWorkStartTime returns a boolean if a field has been set.
func (o *CloudResumeWorkExperience) HasWorkStartTime() bool {
	if o != nil && !IsNil(o.WorkStartTime) {
		return true
	}

	return false
}

// SetWorkStartTime gets a reference to the given int32 and assigns it to the WorkStartTime field.
func (o *CloudResumeWorkExperience) SetWorkStartTime(v int32) {
	o.WorkStartTime = &v
}

func (o CloudResumeWorkExperience) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudResumeWorkExperience) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobName) {
		toSerialize["job_name"] = o.JobName
	}
	if !IsNil(o.PositionName) {
		toSerialize["position_name"] = o.PositionName
	}
	if !IsNil(o.ProfessionId) {
		toSerialize["profession_id"] = o.ProfessionId
	}
	if !IsNil(o.ProfessionName) {
		toSerialize["profession_name"] = o.ProfessionName
	}
	if !IsNil(o.WorkDesc) {
		toSerialize["work_desc"] = o.WorkDesc
	}
	if !IsNil(o.WorkEndTime) {
		toSerialize["work_end_time"] = o.WorkEndTime
	}
	if !IsNil(o.WorkStartTime) {
		toSerialize["work_start_time"] = o.WorkStartTime
	}
	return toSerialize, nil
}

type NullableCloudResumeWorkExperience struct {
	value *CloudResumeWorkExperience
	isSet bool
}

func (v NullableCloudResumeWorkExperience) Get() *CloudResumeWorkExperience {
	return v.value
}

func (v *NullableCloudResumeWorkExperience) Set(val *CloudResumeWorkExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudResumeWorkExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudResumeWorkExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudResumeWorkExperience(val *CloudResumeWorkExperience) *NullableCloudResumeWorkExperience {
	return &NullableCloudResumeWorkExperience{value: val, isSet: true}
}

func (v NullableCloudResumeWorkExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudResumeWorkExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


