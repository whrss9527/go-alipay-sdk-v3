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

// checks if the SearchBoxActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBoxActivity{}

// SearchBoxActivity struct for SearchBoxActivity
type SearchBoxActivity struct {
	// 活动id
	BoxActivityId *string `json:"box_activity_id,omitempty"`
	// 搜索直达id
	BoxId *string `json:"box_id,omitempty"`
	// 活动结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 审核失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// IMAGE-图片/VIDEO-视频
	MaterialType *string `json:"material_type,omitempty"`
	// 当material_type=\"IMAGE\"时，为图片url；当material_type=\"VIDEO\"时，为视频url
	MaterialUrl *string `json:"material_url,omitempty"`
	// 运行状态，INITIAL-初始，ONLINE-已上架，EXPIRE-已失效，OFFLINE-已下架
	RunStatus *string `json:"run_status,omitempty"`
	// 活动开始时间
	StartTime *string `json:"start_time,omitempty"`
	// 配置状态，INITIAL-初始/AUDIT-审核中/CANCEL-已取消/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架/EXPIRE-已失效
	Status *string `json:"status,omitempty"`
	// 跳转应用ID
	TargetAppid *string `json:"target_appid,omitempty"`
	// 目标小程序名称
	TargetAppname *string `json:"target_appname,omitempty"`
	// 投放目标区域
	TargetRegions []DeliveryTargetRegion `json:"target_regions,omitempty"`
	// 活动标题
	Title *string `json:"title,omitempty"`
	VideoInfo *SearchBoxActivityVideoInfo `json:"video_info,omitempty"`
}

// NewSearchBoxActivity instantiates a new SearchBoxActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBoxActivity() *SearchBoxActivity {
	this := SearchBoxActivity{}
	return &this
}

// NewSearchBoxActivityWithDefaults instantiates a new SearchBoxActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBoxActivityWithDefaults() *SearchBoxActivity {
	this := SearchBoxActivity{}
	return &this
}

// GetBoxActivityId returns the BoxActivityId field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetBoxActivityId() string {
	if o == nil || IsNil(o.BoxActivityId) {
		var ret string
		return ret
	}
	return *o.BoxActivityId
}

// GetBoxActivityIdOk returns a tuple with the BoxActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetBoxActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxActivityId) {
		return nil, false
	}
	return o.BoxActivityId, true
}

// HasBoxActivityId returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasBoxActivityId() bool {
	if o != nil && !IsNil(o.BoxActivityId) {
		return true
	}

	return false
}

// SetBoxActivityId gets a reference to the given string and assigns it to the BoxActivityId field.
func (o *SearchBoxActivity) SetBoxActivityId(v string) {
	o.BoxActivityId = &v
}

// GetBoxId returns the BoxId field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetBoxId() string {
	if o == nil || IsNil(o.BoxId) {
		var ret string
		return ret
	}
	return *o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetBoxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxId) {
		return nil, false
	}
	return o.BoxId, true
}

// HasBoxId returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasBoxId() bool {
	if o != nil && !IsNil(o.BoxId) {
		return true
	}

	return false
}

// SetBoxId gets a reference to the given string and assigns it to the BoxId field.
func (o *SearchBoxActivity) SetBoxId(v string) {
	o.BoxId = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *SearchBoxActivity) SetEndTime(v string) {
	o.EndTime = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *SearchBoxActivity) SetFailReason(v string) {
	o.FailReason = &v
}

// GetMaterialType returns the MaterialType field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetMaterialType() string {
	if o == nil || IsNil(o.MaterialType) {
		var ret string
		return ret
	}
	return *o.MaterialType
}

// GetMaterialTypeOk returns a tuple with the MaterialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetMaterialTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MaterialType) {
		return nil, false
	}
	return o.MaterialType, true
}

// HasMaterialType returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasMaterialType() bool {
	if o != nil && !IsNil(o.MaterialType) {
		return true
	}

	return false
}

// SetMaterialType gets a reference to the given string and assigns it to the MaterialType field.
func (o *SearchBoxActivity) SetMaterialType(v string) {
	o.MaterialType = &v
}

// GetMaterialUrl returns the MaterialUrl field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetMaterialUrl() string {
	if o == nil || IsNil(o.MaterialUrl) {
		var ret string
		return ret
	}
	return *o.MaterialUrl
}

// GetMaterialUrlOk returns a tuple with the MaterialUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetMaterialUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MaterialUrl) {
		return nil, false
	}
	return o.MaterialUrl, true
}

// HasMaterialUrl returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasMaterialUrl() bool {
	if o != nil && !IsNil(o.MaterialUrl) {
		return true
	}

	return false
}

// SetMaterialUrl gets a reference to the given string and assigns it to the MaterialUrl field.
func (o *SearchBoxActivity) SetMaterialUrl(v string) {
	o.MaterialUrl = &v
}

// GetRunStatus returns the RunStatus field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetRunStatus() string {
	if o == nil || IsNil(o.RunStatus) {
		var ret string
		return ret
	}
	return *o.RunStatus
}

// GetRunStatusOk returns a tuple with the RunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetRunStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RunStatus) {
		return nil, false
	}
	return o.RunStatus, true
}

// HasRunStatus returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasRunStatus() bool {
	if o != nil && !IsNil(o.RunStatus) {
		return true
	}

	return false
}

// SetRunStatus gets a reference to the given string and assigns it to the RunStatus field.
func (o *SearchBoxActivity) SetRunStatus(v string) {
	o.RunStatus = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *SearchBoxActivity) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SearchBoxActivity) SetStatus(v string) {
	o.Status = &v
}

// GetTargetAppid returns the TargetAppid field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetTargetAppid() string {
	if o == nil || IsNil(o.TargetAppid) {
		var ret string
		return ret
	}
	return *o.TargetAppid
}

// GetTargetAppidOk returns a tuple with the TargetAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetTargetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAppid) {
		return nil, false
	}
	return o.TargetAppid, true
}

// HasTargetAppid returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasTargetAppid() bool {
	if o != nil && !IsNil(o.TargetAppid) {
		return true
	}

	return false
}

// SetTargetAppid gets a reference to the given string and assigns it to the TargetAppid field.
func (o *SearchBoxActivity) SetTargetAppid(v string) {
	o.TargetAppid = &v
}

// GetTargetAppname returns the TargetAppname field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetTargetAppname() string {
	if o == nil || IsNil(o.TargetAppname) {
		var ret string
		return ret
	}
	return *o.TargetAppname
}

// GetTargetAppnameOk returns a tuple with the TargetAppname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetTargetAppnameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAppname) {
		return nil, false
	}
	return o.TargetAppname, true
}

// HasTargetAppname returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasTargetAppname() bool {
	if o != nil && !IsNil(o.TargetAppname) {
		return true
	}

	return false
}

// SetTargetAppname gets a reference to the given string and assigns it to the TargetAppname field.
func (o *SearchBoxActivity) SetTargetAppname(v string) {
	o.TargetAppname = &v
}

// GetTargetRegions returns the TargetRegions field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetTargetRegions() []DeliveryTargetRegion {
	if o == nil || IsNil(o.TargetRegions) {
		var ret []DeliveryTargetRegion
		return ret
	}
	return o.TargetRegions
}

// GetTargetRegionsOk returns a tuple with the TargetRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetTargetRegionsOk() ([]DeliveryTargetRegion, bool) {
	if o == nil || IsNil(o.TargetRegions) {
		return nil, false
	}
	return o.TargetRegions, true
}

// HasTargetRegions returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasTargetRegions() bool {
	if o != nil && !IsNil(o.TargetRegions) {
		return true
	}

	return false
}

// SetTargetRegions gets a reference to the given []DeliveryTargetRegion and assigns it to the TargetRegions field.
func (o *SearchBoxActivity) SetTargetRegions(v []DeliveryTargetRegion) {
	o.TargetRegions = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SearchBoxActivity) SetTitle(v string) {
	o.Title = &v
}

// GetVideoInfo returns the VideoInfo field value if set, zero value otherwise.
func (o *SearchBoxActivity) GetVideoInfo() SearchBoxActivityVideoInfo {
	if o == nil || IsNil(o.VideoInfo) {
		var ret SearchBoxActivityVideoInfo
		return ret
	}
	return *o.VideoInfo
}

// GetVideoInfoOk returns a tuple with the VideoInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxActivity) GetVideoInfoOk() (*SearchBoxActivityVideoInfo, bool) {
	if o == nil || IsNil(o.VideoInfo) {
		return nil, false
	}
	return o.VideoInfo, true
}

// HasVideoInfo returns a boolean if a field has been set.
func (o *SearchBoxActivity) HasVideoInfo() bool {
	if o != nil && !IsNil(o.VideoInfo) {
		return true
	}

	return false
}

// SetVideoInfo gets a reference to the given SearchBoxActivityVideoInfo and assigns it to the VideoInfo field.
func (o *SearchBoxActivity) SetVideoInfo(v SearchBoxActivityVideoInfo) {
	o.VideoInfo = &v
}

func (o SearchBoxActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBoxActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxActivityId) {
		toSerialize["box_activity_id"] = o.BoxActivityId
	}
	if !IsNil(o.BoxId) {
		toSerialize["box_id"] = o.BoxId
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.FailReason) {
		toSerialize["fail_reason"] = o.FailReason
	}
	if !IsNil(o.MaterialType) {
		toSerialize["material_type"] = o.MaterialType
	}
	if !IsNil(o.MaterialUrl) {
		toSerialize["material_url"] = o.MaterialUrl
	}
	if !IsNil(o.RunStatus) {
		toSerialize["run_status"] = o.RunStatus
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TargetAppid) {
		toSerialize["target_appid"] = o.TargetAppid
	}
	if !IsNil(o.TargetAppname) {
		toSerialize["target_appname"] = o.TargetAppname
	}
	if !IsNil(o.TargetRegions) {
		toSerialize["target_regions"] = o.TargetRegions
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.VideoInfo) {
		toSerialize["video_info"] = o.VideoInfo
	}
	return toSerialize, nil
}

type NullableSearchBoxActivity struct {
	value *SearchBoxActivity
	isSet bool
}

func (v NullableSearchBoxActivity) Get() *SearchBoxActivity {
	return v.value
}

func (v *NullableSearchBoxActivity) Set(val *SearchBoxActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBoxActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBoxActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBoxActivity(val *SearchBoxActivity) *NullableSearchBoxActivity {
	return &NullableSearchBoxActivity{value: val, isSet: true}
}

func (v NullableSearchBoxActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBoxActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


