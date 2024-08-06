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

// checks if the AlipayOpenPublicMessageContentModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicMessageContentModifyModel{}

// AlipayOpenPublicMessageContentModifyModel struct for AlipayOpenPublicMessageContentModifyModel
type AlipayOpenPublicMessageContentModifyModel struct {
	// 活动利益点，图文类型ctype为activity类型时才需要传，最多10个字符
	Benefit *string `json:"benefit,omitempty"`
	// 消息正文（支持富文本）
	Content *string `json:"content,omitempty"`
	// 内容id
	ContentId *string `json:"content_id,omitempty"`
	// 是否允许评论 T:允许 F:不允许，默认不允许
	CouldComment *string `json:"could_comment,omitempty"`
	// 封面图url, 尺寸为996*450，最大不超过3M，支持格式:.jpg、.png ，请先调用<ahref=\"https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload\"> 图片上传接口</a>获得图片url。
	Cover *string `json:"cover,omitempty"`
	// 图文类型  activity: 活动图文，不填默认普通图文
	Ctype *string `json:"ctype,omitempty"`
	// 关键词列表，英文逗号分隔，最多不超过5个
	ExtTags *string `json:"ext_tags,omitempty"`
	// 可预览支付宝账号列表，需要预览时才填写， 英文逗号分隔，最多不超过10个
	LoginIds *string `json:"login_ids,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
}

// NewAlipayOpenPublicMessageContentModifyModel instantiates a new AlipayOpenPublicMessageContentModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicMessageContentModifyModel() *AlipayOpenPublicMessageContentModifyModel {
	this := AlipayOpenPublicMessageContentModifyModel{}
	return &this
}

// NewAlipayOpenPublicMessageContentModifyModelWithDefaults instantiates a new AlipayOpenPublicMessageContentModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicMessageContentModifyModelWithDefaults() *AlipayOpenPublicMessageContentModifyModel {
	this := AlipayOpenPublicMessageContentModifyModel{}
	return &this
}

// GetBenefit returns the Benefit field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetBenefit() string {
	if o == nil || IsNil(o.Benefit) {
		var ret string
		return ret
	}
	return *o.Benefit
}

// GetBenefitOk returns a tuple with the Benefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetBenefitOk() (*string, bool) {
	if o == nil || IsNil(o.Benefit) {
		return nil, false
	}
	return o.Benefit, true
}

// HasBenefit returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasBenefit() bool {
	if o != nil && !IsNil(o.Benefit) {
		return true
	}

	return false
}

// SetBenefit gets a reference to the given string and assigns it to the Benefit field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetBenefit(v string) {
	o.Benefit = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetContent(v string) {
	o.Content = &v
}

// GetContentId returns the ContentId field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetContentId() string {
	if o == nil || IsNil(o.ContentId) {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetContentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContentId) {
		return nil, false
	}
	return o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasContentId() bool {
	if o != nil && !IsNil(o.ContentId) {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetContentId(v string) {
	o.ContentId = &v
}

// GetCouldComment returns the CouldComment field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetCouldComment() string {
	if o == nil || IsNil(o.CouldComment) {
		var ret string
		return ret
	}
	return *o.CouldComment
}

// GetCouldCommentOk returns a tuple with the CouldComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetCouldCommentOk() (*string, bool) {
	if o == nil || IsNil(o.CouldComment) {
		return nil, false
	}
	return o.CouldComment, true
}

// HasCouldComment returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasCouldComment() bool {
	if o != nil && !IsNil(o.CouldComment) {
		return true
	}

	return false
}

// SetCouldComment gets a reference to the given string and assigns it to the CouldComment field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetCouldComment(v string) {
	o.CouldComment = &v
}

// GetCover returns the Cover field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetCover() string {
	if o == nil || IsNil(o.Cover) {
		var ret string
		return ret
	}
	return *o.Cover
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetCoverOk() (*string, bool) {
	if o == nil || IsNil(o.Cover) {
		return nil, false
	}
	return o.Cover, true
}

// HasCover returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasCover() bool {
	if o != nil && !IsNil(o.Cover) {
		return true
	}

	return false
}

// SetCover gets a reference to the given string and assigns it to the Cover field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetCover(v string) {
	o.Cover = &v
}

// GetCtype returns the Ctype field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetCtype() string {
	if o == nil || IsNil(o.Ctype) {
		var ret string
		return ret
	}
	return *o.Ctype
}

// GetCtypeOk returns a tuple with the Ctype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetCtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Ctype) {
		return nil, false
	}
	return o.Ctype, true
}

// HasCtype returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasCtype() bool {
	if o != nil && !IsNil(o.Ctype) {
		return true
	}

	return false
}

// SetCtype gets a reference to the given string and assigns it to the Ctype field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetCtype(v string) {
	o.Ctype = &v
}

// GetExtTags returns the ExtTags field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetExtTags() string {
	if o == nil || IsNil(o.ExtTags) {
		var ret string
		return ret
	}
	return *o.ExtTags
}

// GetExtTagsOk returns a tuple with the ExtTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetExtTagsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtTags) {
		return nil, false
	}
	return o.ExtTags, true
}

// HasExtTags returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasExtTags() bool {
	if o != nil && !IsNil(o.ExtTags) {
		return true
	}

	return false
}

// SetExtTags gets a reference to the given string and assigns it to the ExtTags field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetExtTags(v string) {
	o.ExtTags = &v
}

// GetLoginIds returns the LoginIds field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetLoginIds() string {
	if o == nil || IsNil(o.LoginIds) {
		var ret string
		return ret
	}
	return *o.LoginIds
}

// GetLoginIdsOk returns a tuple with the LoginIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetLoginIdsOk() (*string, bool) {
	if o == nil || IsNil(o.LoginIds) {
		return nil, false
	}
	return o.LoginIds, true
}

// HasLoginIds returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasLoginIds() bool {
	if o != nil && !IsNil(o.LoginIds) {
		return true
	}

	return false
}

// SetLoginIds gets a reference to the given string and assigns it to the LoginIds field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetLoginIds(v string) {
	o.LoginIds = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentModifyModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentModifyModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayOpenPublicMessageContentModifyModel) SetTitle(v string) {
	o.Title = &v
}

func (o AlipayOpenPublicMessageContentModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicMessageContentModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Benefit) {
		toSerialize["benefit"] = o.Benefit
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.ContentId) {
		toSerialize["content_id"] = o.ContentId
	}
	if !IsNil(o.CouldComment) {
		toSerialize["could_comment"] = o.CouldComment
	}
	if !IsNil(o.Cover) {
		toSerialize["cover"] = o.Cover
	}
	if !IsNil(o.Ctype) {
		toSerialize["ctype"] = o.Ctype
	}
	if !IsNil(o.ExtTags) {
		toSerialize["ext_tags"] = o.ExtTags
	}
	if !IsNil(o.LoginIds) {
		toSerialize["login_ids"] = o.LoginIds
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicMessageContentModifyModel struct {
	value *AlipayOpenPublicMessageContentModifyModel
	isSet bool
}

func (v NullableAlipayOpenPublicMessageContentModifyModel) Get() *AlipayOpenPublicMessageContentModifyModel {
	return v.value
}

func (v *NullableAlipayOpenPublicMessageContentModifyModel) Set(val *AlipayOpenPublicMessageContentModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMessageContentModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMessageContentModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMessageContentModifyModel(val *AlipayOpenPublicMessageContentModifyModel) *NullableAlipayOpenPublicMessageContentModifyModel {
	return &NullableAlipayOpenPublicMessageContentModifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMessageContentModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMessageContentModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


