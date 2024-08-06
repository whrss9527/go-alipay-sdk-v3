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

// checks if the AlipayOpenMiniTipsDeliveryQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniTipsDeliveryQueryResponseModel{}

// AlipayOpenMiniTipsDeliveryQueryResponseModel struct for AlipayOpenMiniTipsDeliveryQueryResponseModel
type AlipayOpenMiniTipsDeliveryQueryResponseModel struct {
	// 收藏引导文案内容，不得超过14个字
	DeliveryContent *string `json:"delivery_content,omitempty"`
	// 收藏引导投放活动ID
	DeliveryId *string `json:"delivery_id,omitempty"`
	// 收藏引导配置投放活动的名称
	DeliveryName *string `json:"delivery_name,omitempty"`
	// 活动结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 文案审核驳回理由，仅AUDIT_REJECTED生效
	FailReason *string `json:"fail_reason,omitempty"`
	// 匹配类型
	MatchType *string `json:"match_type,omitempty"`
	// 目标页面地址，当匹配类型为TARGETURL时，投放文案仅在当前页面生效
	MatchUrl *string `json:"match_url,omitempty"`
	// 活动开始时间
	StartTime *string `json:"start_time,omitempty"`
	// 活动状态
	Status *string `json:"status,omitempty"`
}

// NewAlipayOpenMiniTipsDeliveryQueryResponseModel instantiates a new AlipayOpenMiniTipsDeliveryQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniTipsDeliveryQueryResponseModel() *AlipayOpenMiniTipsDeliveryQueryResponseModel {
	this := AlipayOpenMiniTipsDeliveryQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniTipsDeliveryQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniTipsDeliveryQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniTipsDeliveryQueryResponseModelWithDefaults() *AlipayOpenMiniTipsDeliveryQueryResponseModel {
	this := AlipayOpenMiniTipsDeliveryQueryResponseModel{}
	return &this
}

// GetDeliveryContent returns the DeliveryContent field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetDeliveryContent() string {
	if o == nil || IsNil(o.DeliveryContent) {
		var ret string
		return ret
	}
	return *o.DeliveryContent
}

// GetDeliveryContentOk returns a tuple with the DeliveryContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetDeliveryContentOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryContent) {
		return nil, false
	}
	return o.DeliveryContent, true
}

// HasDeliveryContent returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasDeliveryContent() bool {
	if o != nil && !IsNil(o.DeliveryContent) {
		return true
	}

	return false
}

// SetDeliveryContent gets a reference to the given string and assigns it to the DeliveryContent field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetDeliveryContent(v string) {
	o.DeliveryContent = &v
}

// GetDeliveryId returns the DeliveryId field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetDeliveryId() string {
	if o == nil || IsNil(o.DeliveryId) {
		var ret string
		return ret
	}
	return *o.DeliveryId
}

// GetDeliveryIdOk returns a tuple with the DeliveryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetDeliveryIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryId) {
		return nil, false
	}
	return o.DeliveryId, true
}

// HasDeliveryId returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasDeliveryId() bool {
	if o != nil && !IsNil(o.DeliveryId) {
		return true
	}

	return false
}

// SetDeliveryId gets a reference to the given string and assigns it to the DeliveryId field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetDeliveryId(v string) {
	o.DeliveryId = &v
}

// GetDeliveryName returns the DeliveryName field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetDeliveryName() string {
	if o == nil || IsNil(o.DeliveryName) {
		var ret string
		return ret
	}
	return *o.DeliveryName
}

// GetDeliveryNameOk returns a tuple with the DeliveryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetDeliveryNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryName) {
		return nil, false
	}
	return o.DeliveryName, true
}

// HasDeliveryName returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasDeliveryName() bool {
	if o != nil && !IsNil(o.DeliveryName) {
		return true
	}

	return false
}

// SetDeliveryName gets a reference to the given string and assigns it to the DeliveryName field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetDeliveryName(v string) {
	o.DeliveryName = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetEndTime(v string) {
	o.EndTime = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetFailReason(v string) {
	o.FailReason = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetMatchType(v string) {
	o.MatchType = &v
}

// GetMatchUrl returns the MatchUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetMatchUrl() string {
	if o == nil || IsNil(o.MatchUrl) {
		var ret string
		return ret
	}
	return *o.MatchUrl
}

// GetMatchUrlOk returns a tuple with the MatchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetMatchUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MatchUrl) {
		return nil, false
	}
	return o.MatchUrl, true
}

// HasMatchUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasMatchUrl() bool {
	if o != nil && !IsNil(o.MatchUrl) {
		return true
	}

	return false
}

// SetMatchUrl gets a reference to the given string and assigns it to the MatchUrl field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetMatchUrl(v string) {
	o.MatchUrl = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayOpenMiniTipsDeliveryQueryResponseModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayOpenMiniTipsDeliveryQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniTipsDeliveryQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryContent) {
		toSerialize["delivery_content"] = o.DeliveryContent
	}
	if !IsNil(o.DeliveryId) {
		toSerialize["delivery_id"] = o.DeliveryId
	}
	if !IsNil(o.DeliveryName) {
		toSerialize["delivery_name"] = o.DeliveryName
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.FailReason) {
		toSerialize["fail_reason"] = o.FailReason
	}
	if !IsNil(o.MatchType) {
		toSerialize["match_type"] = o.MatchType
	}
	if !IsNil(o.MatchUrl) {
		toSerialize["match_url"] = o.MatchUrl
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniTipsDeliveryQueryResponseModel struct {
	value *AlipayOpenMiniTipsDeliveryQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniTipsDeliveryQueryResponseModel) Get() *AlipayOpenMiniTipsDeliveryQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniTipsDeliveryQueryResponseModel) Set(val *AlipayOpenMiniTipsDeliveryQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTipsDeliveryQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTipsDeliveryQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTipsDeliveryQueryResponseModel(val *AlipayOpenMiniTipsDeliveryQueryResponseModel) *NullableAlipayOpenMiniTipsDeliveryQueryResponseModel {
	return &NullableAlipayOpenMiniTipsDeliveryQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTipsDeliveryQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTipsDeliveryQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


