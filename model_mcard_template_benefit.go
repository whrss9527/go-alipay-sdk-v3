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

// checks if the McardTemplateBenefit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McardTemplateBenefit{}

// McardTemplateBenefit struct for McardTemplateBenefit
type McardTemplateBenefit struct {
	// 权益描述信息
	BenefitDesc []string `json:"benefit_desc,omitempty"`
	// 权益结束时间。  注：在权益开始时间和结束时间范围内的权益才会认为是有效权益进行展示。
	EndDate *string `json:"end_date,omitempty"`
	// 会员卡模板权益扩展信息：JSON格式; openUrl 说明：跳转到商户的优惠活动页面
	ExtInfo *string `json:"ext_info,omitempty"`
	// 权益开始时间
	StartDate *string `json:"start_date,omitempty"`
	// 会员卡模板ID
	TemplateId *string `json:"template_id,omitempty"`
	// 权益标题
	Title *string `json:"title,omitempty"`
}

// NewMcardTemplateBenefit instantiates a new McardTemplateBenefit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcardTemplateBenefit() *McardTemplateBenefit {
	this := McardTemplateBenefit{}
	return &this
}

// NewMcardTemplateBenefitWithDefaults instantiates a new McardTemplateBenefit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcardTemplateBenefitWithDefaults() *McardTemplateBenefit {
	this := McardTemplateBenefit{}
	return &this
}

// GetBenefitDesc returns the BenefitDesc field value if set, zero value otherwise.
func (o *McardTemplateBenefit) GetBenefitDesc() []string {
	if o == nil || IsNil(o.BenefitDesc) {
		var ret []string
		return ret
	}
	return o.BenefitDesc
}

// GetBenefitDescOk returns a tuple with the BenefitDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardTemplateBenefit) GetBenefitDescOk() ([]string, bool) {
	if o == nil || IsNil(o.BenefitDesc) {
		return nil, false
	}
	return o.BenefitDesc, true
}

// HasBenefitDesc returns a boolean if a field has been set.
func (o *McardTemplateBenefit) HasBenefitDesc() bool {
	if o != nil && !IsNil(o.BenefitDesc) {
		return true
	}

	return false
}

// SetBenefitDesc gets a reference to the given []string and assigns it to the BenefitDesc field.
func (o *McardTemplateBenefit) SetBenefitDesc(v []string) {
	o.BenefitDesc = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *McardTemplateBenefit) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardTemplateBenefit) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *McardTemplateBenefit) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *McardTemplateBenefit) SetEndDate(v string) {
	o.EndDate = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *McardTemplateBenefit) GetExtInfo() string {
	if o == nil || IsNil(o.ExtInfo) {
		var ret string
		return ret
	}
	return *o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardTemplateBenefit) GetExtInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *McardTemplateBenefit) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given string and assigns it to the ExtInfo field.
func (o *McardTemplateBenefit) SetExtInfo(v string) {
	o.ExtInfo = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *McardTemplateBenefit) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardTemplateBenefit) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *McardTemplateBenefit) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *McardTemplateBenefit) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *McardTemplateBenefit) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardTemplateBenefit) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *McardTemplateBenefit) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *McardTemplateBenefit) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *McardTemplateBenefit) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardTemplateBenefit) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *McardTemplateBenefit) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *McardTemplateBenefit) SetTitle(v string) {
	o.Title = &v
}

func (o McardTemplateBenefit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McardTemplateBenefit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BenefitDesc) {
		toSerialize["benefit_desc"] = o.BenefitDesc
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableMcardTemplateBenefit struct {
	value *McardTemplateBenefit
	isSet bool
}

func (v NullableMcardTemplateBenefit) Get() *McardTemplateBenefit {
	return v.value
}

func (v *NullableMcardTemplateBenefit) Set(val *McardTemplateBenefit) {
	v.value = val
	v.isSet = true
}

func (v NullableMcardTemplateBenefit) IsSet() bool {
	return v.isSet
}

func (v *NullableMcardTemplateBenefit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcardTemplateBenefit(val *McardTemplateBenefit) *NullableMcardTemplateBenefit {
	return &NullableMcardTemplateBenefit{value: val, isSet: true}
}

func (v NullableMcardTemplateBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcardTemplateBenefit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


