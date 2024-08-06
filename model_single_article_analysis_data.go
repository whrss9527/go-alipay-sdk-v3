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

// checks if the SingleArticleAnalysisData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleArticleAnalysisData{}

// SingleArticleAnalysisData struct for SingleArticleAnalysisData
type SingleArticleAnalysisData struct {
	// 人均阅读时长
	AvgReadTime *string `json:"avg_read_time,omitempty"`
	// 文章发布日期
	Date *string `json:"date,omitempty"`
	// 送达人数
	DeliverUserCnt *int32 `json:"deliver_user_cnt,omitempty"`
	// 曝光人数
	ExposeUserCnt *int32 `json:"expose_user_cnt,omitempty"`
	// 点赞数
	PraiseUserCnt *int32 `json:"praise_user_cnt,omitempty"`
	// 阅读人数
	ReadUserCnt *int32 `json:"read_user_cnt,omitempty"`
	// 评论数
	ReplyUserCnt *int32 `json:"reply_user_cnt,omitempty"`
	// 分享人数
	ShareUserCnt *int32 `json:"share_user_cnt,omitempty"`
	// 文章标题
	Title *string `json:"title,omitempty"`
}

// NewSingleArticleAnalysisData instantiates a new SingleArticleAnalysisData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleArticleAnalysisData() *SingleArticleAnalysisData {
	this := SingleArticleAnalysisData{}
	return &this
}

// NewSingleArticleAnalysisDataWithDefaults instantiates a new SingleArticleAnalysisData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleArticleAnalysisDataWithDefaults() *SingleArticleAnalysisData {
	this := SingleArticleAnalysisData{}
	return &this
}

// GetAvgReadTime returns the AvgReadTime field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetAvgReadTime() string {
	if o == nil || IsNil(o.AvgReadTime) {
		var ret string
		return ret
	}
	return *o.AvgReadTime
}

// GetAvgReadTimeOk returns a tuple with the AvgReadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetAvgReadTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AvgReadTime) {
		return nil, false
	}
	return o.AvgReadTime, true
}

// HasAvgReadTime returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasAvgReadTime() bool {
	if o != nil && !IsNil(o.AvgReadTime) {
		return true
	}

	return false
}

// SetAvgReadTime gets a reference to the given string and assigns it to the AvgReadTime field.
func (o *SingleArticleAnalysisData) SetAvgReadTime(v string) {
	o.AvgReadTime = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *SingleArticleAnalysisData) SetDate(v string) {
	o.Date = &v
}

// GetDeliverUserCnt returns the DeliverUserCnt field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetDeliverUserCnt() int32 {
	if o == nil || IsNil(o.DeliverUserCnt) {
		var ret int32
		return ret
	}
	return *o.DeliverUserCnt
}

// GetDeliverUserCntOk returns a tuple with the DeliverUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetDeliverUserCntOk() (*int32, bool) {
	if o == nil || IsNil(o.DeliverUserCnt) {
		return nil, false
	}
	return o.DeliverUserCnt, true
}

// HasDeliverUserCnt returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasDeliverUserCnt() bool {
	if o != nil && !IsNil(o.DeliverUserCnt) {
		return true
	}

	return false
}

// SetDeliverUserCnt gets a reference to the given int32 and assigns it to the DeliverUserCnt field.
func (o *SingleArticleAnalysisData) SetDeliverUserCnt(v int32) {
	o.DeliverUserCnt = &v
}

// GetExposeUserCnt returns the ExposeUserCnt field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetExposeUserCnt() int32 {
	if o == nil || IsNil(o.ExposeUserCnt) {
		var ret int32
		return ret
	}
	return *o.ExposeUserCnt
}

// GetExposeUserCntOk returns a tuple with the ExposeUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetExposeUserCntOk() (*int32, bool) {
	if o == nil || IsNil(o.ExposeUserCnt) {
		return nil, false
	}
	return o.ExposeUserCnt, true
}

// HasExposeUserCnt returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasExposeUserCnt() bool {
	if o != nil && !IsNil(o.ExposeUserCnt) {
		return true
	}

	return false
}

// SetExposeUserCnt gets a reference to the given int32 and assigns it to the ExposeUserCnt field.
func (o *SingleArticleAnalysisData) SetExposeUserCnt(v int32) {
	o.ExposeUserCnt = &v
}

// GetPraiseUserCnt returns the PraiseUserCnt field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetPraiseUserCnt() int32 {
	if o == nil || IsNil(o.PraiseUserCnt) {
		var ret int32
		return ret
	}
	return *o.PraiseUserCnt
}

// GetPraiseUserCntOk returns a tuple with the PraiseUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetPraiseUserCntOk() (*int32, bool) {
	if o == nil || IsNil(o.PraiseUserCnt) {
		return nil, false
	}
	return o.PraiseUserCnt, true
}

// HasPraiseUserCnt returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasPraiseUserCnt() bool {
	if o != nil && !IsNil(o.PraiseUserCnt) {
		return true
	}

	return false
}

// SetPraiseUserCnt gets a reference to the given int32 and assigns it to the PraiseUserCnt field.
func (o *SingleArticleAnalysisData) SetPraiseUserCnt(v int32) {
	o.PraiseUserCnt = &v
}

// GetReadUserCnt returns the ReadUserCnt field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetReadUserCnt() int32 {
	if o == nil || IsNil(o.ReadUserCnt) {
		var ret int32
		return ret
	}
	return *o.ReadUserCnt
}

// GetReadUserCntOk returns a tuple with the ReadUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetReadUserCntOk() (*int32, bool) {
	if o == nil || IsNil(o.ReadUserCnt) {
		return nil, false
	}
	return o.ReadUserCnt, true
}

// HasReadUserCnt returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasReadUserCnt() bool {
	if o != nil && !IsNil(o.ReadUserCnt) {
		return true
	}

	return false
}

// SetReadUserCnt gets a reference to the given int32 and assigns it to the ReadUserCnt field.
func (o *SingleArticleAnalysisData) SetReadUserCnt(v int32) {
	o.ReadUserCnt = &v
}

// GetReplyUserCnt returns the ReplyUserCnt field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetReplyUserCnt() int32 {
	if o == nil || IsNil(o.ReplyUserCnt) {
		var ret int32
		return ret
	}
	return *o.ReplyUserCnt
}

// GetReplyUserCntOk returns a tuple with the ReplyUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetReplyUserCntOk() (*int32, bool) {
	if o == nil || IsNil(o.ReplyUserCnt) {
		return nil, false
	}
	return o.ReplyUserCnt, true
}

// HasReplyUserCnt returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasReplyUserCnt() bool {
	if o != nil && !IsNil(o.ReplyUserCnt) {
		return true
	}

	return false
}

// SetReplyUserCnt gets a reference to the given int32 and assigns it to the ReplyUserCnt field.
func (o *SingleArticleAnalysisData) SetReplyUserCnt(v int32) {
	o.ReplyUserCnt = &v
}

// GetShareUserCnt returns the ShareUserCnt field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetShareUserCnt() int32 {
	if o == nil || IsNil(o.ShareUserCnt) {
		var ret int32
		return ret
	}
	return *o.ShareUserCnt
}

// GetShareUserCntOk returns a tuple with the ShareUserCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetShareUserCntOk() (*int32, bool) {
	if o == nil || IsNil(o.ShareUserCnt) {
		return nil, false
	}
	return o.ShareUserCnt, true
}

// HasShareUserCnt returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasShareUserCnt() bool {
	if o != nil && !IsNil(o.ShareUserCnt) {
		return true
	}

	return false
}

// SetShareUserCnt gets a reference to the given int32 and assigns it to the ShareUserCnt field.
func (o *SingleArticleAnalysisData) SetShareUserCnt(v int32) {
	o.ShareUserCnt = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SingleArticleAnalysisData) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleArticleAnalysisData) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SingleArticleAnalysisData) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SingleArticleAnalysisData) SetTitle(v string) {
	o.Title = &v
}

func (o SingleArticleAnalysisData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleArticleAnalysisData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgReadTime) {
		toSerialize["avg_read_time"] = o.AvgReadTime
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.DeliverUserCnt) {
		toSerialize["deliver_user_cnt"] = o.DeliverUserCnt
	}
	if !IsNil(o.ExposeUserCnt) {
		toSerialize["expose_user_cnt"] = o.ExposeUserCnt
	}
	if !IsNil(o.PraiseUserCnt) {
		toSerialize["praise_user_cnt"] = o.PraiseUserCnt
	}
	if !IsNil(o.ReadUserCnt) {
		toSerialize["read_user_cnt"] = o.ReadUserCnt
	}
	if !IsNil(o.ReplyUserCnt) {
		toSerialize["reply_user_cnt"] = o.ReplyUserCnt
	}
	if !IsNil(o.ShareUserCnt) {
		toSerialize["share_user_cnt"] = o.ShareUserCnt
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableSingleArticleAnalysisData struct {
	value *SingleArticleAnalysisData
	isSet bool
}

func (v NullableSingleArticleAnalysisData) Get() *SingleArticleAnalysisData {
	return v.value
}

func (v *NullableSingleArticleAnalysisData) Set(val *SingleArticleAnalysisData) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleArticleAnalysisData) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleArticleAnalysisData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleArticleAnalysisData(val *SingleArticleAnalysisData) *NullableSingleArticleAnalysisData {
	return &NullableSingleArticleAnalysisData{value: val, isSet: true}
}

func (v NullableSingleArticleAnalysisData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleArticleAnalysisData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


