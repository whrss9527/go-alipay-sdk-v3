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

// checks if the CloudResumePositionIntention type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudResumePositionIntention{}

// CloudResumePositionIntention struct for CloudResumePositionIntention
type CloudResumePositionIntention struct {
	// 意向城市Code。具体地区编码参见https://lbs.amap.com/api/webservice/download 里面城市编码表
	IntentionCity *string `json:"intention_city,omitempty"`
	// 职业id(这个字段在目前版本不对外暴露)
	// Deprecated
	JobId *string `json:"job_id,omitempty"`
	// 职业名称
	JobName *string `json:"job_name,omitempty"`
	// 行业id(这个字段在目前版本不对外暴露)
	// Deprecated
	ProfessionId *string `json:"profession_id,omitempty"`
	// 行业名称
	ProfessionName *string `json:"profession_name,omitempty"`
	// 最大工资
	SalaryMax *string `json:"salary_max,omitempty"`
	// 最小薪资
	SalaryMin *string `json:"salary_min,omitempty"`
	// 工资单位，/月 /日
	SalaryUnit *string `json:"salary_unit,omitempty"`
	// 工作属性
	WorkProperty *string `json:"work_property,omitempty"`
}

// NewCloudResumePositionIntention instantiates a new CloudResumePositionIntention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudResumePositionIntention() *CloudResumePositionIntention {
	this := CloudResumePositionIntention{}
	return &this
}

// NewCloudResumePositionIntentionWithDefaults instantiates a new CloudResumePositionIntention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudResumePositionIntentionWithDefaults() *CloudResumePositionIntention {
	this := CloudResumePositionIntention{}
	return &this
}

// GetIntentionCity returns the IntentionCity field value if set, zero value otherwise.
func (o *CloudResumePositionIntention) GetIntentionCity() string {
	if o == nil || IsNil(o.IntentionCity) {
		var ret string
		return ret
	}
	return *o.IntentionCity
}

// GetIntentionCityOk returns a tuple with the IntentionCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumePositionIntention) GetIntentionCityOk() (*string, bool) {
	if o == nil || IsNil(o.IntentionCity) {
		return nil, false
	}
	return o.IntentionCity, true
}

// HasIntentionCity returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasIntentionCity() bool {
	if o != nil && !IsNil(o.IntentionCity) {
		return true
	}

	return false
}

// SetIntentionCity gets a reference to the given string and assigns it to the IntentionCity field.
func (o *CloudResumePositionIntention) SetIntentionCity(v string) {
	o.IntentionCity = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
// Deprecated
func (o *CloudResumePositionIntention) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CloudResumePositionIntention) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
// Deprecated
func (o *CloudResumePositionIntention) SetJobId(v string) {
	o.JobId = &v
}

// GetJobName returns the JobName field value if set, zero value otherwise.
func (o *CloudResumePositionIntention) GetJobName() string {
	if o == nil || IsNil(o.JobName) {
		var ret string
		return ret
	}
	return *o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumePositionIntention) GetJobNameOk() (*string, bool) {
	if o == nil || IsNil(o.JobName) {
		return nil, false
	}
	return o.JobName, true
}

// HasJobName returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasJobName() bool {
	if o != nil && !IsNil(o.JobName) {
		return true
	}

	return false
}

// SetJobName gets a reference to the given string and assigns it to the JobName field.
func (o *CloudResumePositionIntention) SetJobName(v string) {
	o.JobName = &v
}

// GetProfessionId returns the ProfessionId field value if set, zero value otherwise.
// Deprecated
func (o *CloudResumePositionIntention) GetProfessionId() string {
	if o == nil || IsNil(o.ProfessionId) {
		var ret string
		return ret
	}
	return *o.ProfessionId
}

// GetProfessionIdOk returns a tuple with the ProfessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CloudResumePositionIntention) GetProfessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfessionId) {
		return nil, false
	}
	return o.ProfessionId, true
}

// HasProfessionId returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasProfessionId() bool {
	if o != nil && !IsNil(o.ProfessionId) {
		return true
	}

	return false
}

// SetProfessionId gets a reference to the given string and assigns it to the ProfessionId field.
// Deprecated
func (o *CloudResumePositionIntention) SetProfessionId(v string) {
	o.ProfessionId = &v
}

// GetProfessionName returns the ProfessionName field value if set, zero value otherwise.
func (o *CloudResumePositionIntention) GetProfessionName() string {
	if o == nil || IsNil(o.ProfessionName) {
		var ret string
		return ret
	}
	return *o.ProfessionName
}

// GetProfessionNameOk returns a tuple with the ProfessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumePositionIntention) GetProfessionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProfessionName) {
		return nil, false
	}
	return o.ProfessionName, true
}

// HasProfessionName returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasProfessionName() bool {
	if o != nil && !IsNil(o.ProfessionName) {
		return true
	}

	return false
}

// SetProfessionName gets a reference to the given string and assigns it to the ProfessionName field.
func (o *CloudResumePositionIntention) SetProfessionName(v string) {
	o.ProfessionName = &v
}

// GetSalaryMax returns the SalaryMax field value if set, zero value otherwise.
func (o *CloudResumePositionIntention) GetSalaryMax() string {
	if o == nil || IsNil(o.SalaryMax) {
		var ret string
		return ret
	}
	return *o.SalaryMax
}

// GetSalaryMaxOk returns a tuple with the SalaryMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumePositionIntention) GetSalaryMaxOk() (*string, bool) {
	if o == nil || IsNil(o.SalaryMax) {
		return nil, false
	}
	return o.SalaryMax, true
}

// HasSalaryMax returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasSalaryMax() bool {
	if o != nil && !IsNil(o.SalaryMax) {
		return true
	}

	return false
}

// SetSalaryMax gets a reference to the given string and assigns it to the SalaryMax field.
func (o *CloudResumePositionIntention) SetSalaryMax(v string) {
	o.SalaryMax = &v
}

// GetSalaryMin returns the SalaryMin field value if set, zero value otherwise.
func (o *CloudResumePositionIntention) GetSalaryMin() string {
	if o == nil || IsNil(o.SalaryMin) {
		var ret string
		return ret
	}
	return *o.SalaryMin
}

// GetSalaryMinOk returns a tuple with the SalaryMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumePositionIntention) GetSalaryMinOk() (*string, bool) {
	if o == nil || IsNil(o.SalaryMin) {
		return nil, false
	}
	return o.SalaryMin, true
}

// HasSalaryMin returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasSalaryMin() bool {
	if o != nil && !IsNil(o.SalaryMin) {
		return true
	}

	return false
}

// SetSalaryMin gets a reference to the given string and assigns it to the SalaryMin field.
func (o *CloudResumePositionIntention) SetSalaryMin(v string) {
	o.SalaryMin = &v
}

// GetSalaryUnit returns the SalaryUnit field value if set, zero value otherwise.
func (o *CloudResumePositionIntention) GetSalaryUnit() string {
	if o == nil || IsNil(o.SalaryUnit) {
		var ret string
		return ret
	}
	return *o.SalaryUnit
}

// GetSalaryUnitOk returns a tuple with the SalaryUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumePositionIntention) GetSalaryUnitOk() (*string, bool) {
	if o == nil || IsNil(o.SalaryUnit) {
		return nil, false
	}
	return o.SalaryUnit, true
}

// HasSalaryUnit returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasSalaryUnit() bool {
	if o != nil && !IsNil(o.SalaryUnit) {
		return true
	}

	return false
}

// SetSalaryUnit gets a reference to the given string and assigns it to the SalaryUnit field.
func (o *CloudResumePositionIntention) SetSalaryUnit(v string) {
	o.SalaryUnit = &v
}

// GetWorkProperty returns the WorkProperty field value if set, zero value otherwise.
func (o *CloudResumePositionIntention) GetWorkProperty() string {
	if o == nil || IsNil(o.WorkProperty) {
		var ret string
		return ret
	}
	return *o.WorkProperty
}

// GetWorkPropertyOk returns a tuple with the WorkProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumePositionIntention) GetWorkPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.WorkProperty) {
		return nil, false
	}
	return o.WorkProperty, true
}

// HasWorkProperty returns a boolean if a field has been set.
func (o *CloudResumePositionIntention) HasWorkProperty() bool {
	if o != nil && !IsNil(o.WorkProperty) {
		return true
	}

	return false
}

// SetWorkProperty gets a reference to the given string and assigns it to the WorkProperty field.
func (o *CloudResumePositionIntention) SetWorkProperty(v string) {
	o.WorkProperty = &v
}

func (o CloudResumePositionIntention) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudResumePositionIntention) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntentionCity) {
		toSerialize["intention_city"] = o.IntentionCity
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobName) {
		toSerialize["job_name"] = o.JobName
	}
	if !IsNil(o.ProfessionId) {
		toSerialize["profession_id"] = o.ProfessionId
	}
	if !IsNil(o.ProfessionName) {
		toSerialize["profession_name"] = o.ProfessionName
	}
	if !IsNil(o.SalaryMax) {
		toSerialize["salary_max"] = o.SalaryMax
	}
	if !IsNil(o.SalaryMin) {
		toSerialize["salary_min"] = o.SalaryMin
	}
	if !IsNil(o.SalaryUnit) {
		toSerialize["salary_unit"] = o.SalaryUnit
	}
	if !IsNil(o.WorkProperty) {
		toSerialize["work_property"] = o.WorkProperty
	}
	return toSerialize, nil
}

type NullableCloudResumePositionIntention struct {
	value *CloudResumePositionIntention
	isSet bool
}

func (v NullableCloudResumePositionIntention) Get() *CloudResumePositionIntention {
	return v.value
}

func (v *NullableCloudResumePositionIntention) Set(val *CloudResumePositionIntention) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudResumePositionIntention) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudResumePositionIntention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudResumePositionIntention(val *CloudResumePositionIntention) *NullableCloudResumePositionIntention {
	return &NullableCloudResumePositionIntention{value: val, isSet: true}
}

func (v NullableCloudResumePositionIntention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudResumePositionIntention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


