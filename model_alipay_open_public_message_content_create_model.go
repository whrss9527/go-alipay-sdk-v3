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

// checks if the AlipayOpenPublicMessageContentCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicMessageContentCreateModel{}

// AlipayOpenPublicMessageContentCreateModel struct for AlipayOpenPublicMessageContentCreateModel
type AlipayOpenPublicMessageContentCreateModel struct {
	// 活动利益点，最多10个字符。仅 ctype 为 activity 类型时才需传入。
	Benefit *string `json:"benefit,omitempty"`
	// 消息正文（支持富文本）
	Content *string `json:"content,omitempty"`
	// 是否允许评论。枚举支持： *T：允许。 *F：不允许，默认不允许。
	CouldComment *string `json:"could_comment,omitempty"`
	// 封面图url，需传入 <a href=\"https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload\">图片上传接口</a>上传图片获取的 image_url。 注意：尺寸为 996*450，最大不超过3M，支持格式：jpg、.png 。
	Cover *string `json:"cover,omitempty"`
	// 图文类型，不填默认普通图文。还支持 activity（活动图文）。
	Ctype *string `json:"ctype,omitempty"`
	// 关键词列表，英文逗号分隔，最多不超过5个
	ExtTags *string `json:"ext_tags,omitempty"`
	// 可预览支付宝账号列表，需要预览时才填写， 英文逗号分隔，最多不超过10个
	LoginIds *string `json:"login_ids,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
}

// NewAlipayOpenPublicMessageContentCreateModel instantiates a new AlipayOpenPublicMessageContentCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicMessageContentCreateModel() *AlipayOpenPublicMessageContentCreateModel {
	this := AlipayOpenPublicMessageContentCreateModel{}
	return &this
}

// NewAlipayOpenPublicMessageContentCreateModelWithDefaults instantiates a new AlipayOpenPublicMessageContentCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicMessageContentCreateModelWithDefaults() *AlipayOpenPublicMessageContentCreateModel {
	this := AlipayOpenPublicMessageContentCreateModel{}
	return &this
}

// GetBenefit returns the Benefit field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetBenefit() string {
	if o == nil || IsNil(o.Benefit) {
		var ret string
		return ret
	}
	return *o.Benefit
}

// GetBenefitOk returns a tuple with the Benefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetBenefitOk() (*string, bool) {
	if o == nil || IsNil(o.Benefit) {
		return nil, false
	}
	return o.Benefit, true
}

// HasBenefit returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasBenefit() bool {
	if o != nil && !IsNil(o.Benefit) {
		return true
	}

	return false
}

// SetBenefit gets a reference to the given string and assigns it to the Benefit field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetBenefit(v string) {
	o.Benefit = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetContent(v string) {
	o.Content = &v
}

// GetCouldComment returns the CouldComment field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetCouldComment() string {
	if o == nil || IsNil(o.CouldComment) {
		var ret string
		return ret
	}
	return *o.CouldComment
}

// GetCouldCommentOk returns a tuple with the CouldComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetCouldCommentOk() (*string, bool) {
	if o == nil || IsNil(o.CouldComment) {
		return nil, false
	}
	return o.CouldComment, true
}

// HasCouldComment returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasCouldComment() bool {
	if o != nil && !IsNil(o.CouldComment) {
		return true
	}

	return false
}

// SetCouldComment gets a reference to the given string and assigns it to the CouldComment field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetCouldComment(v string) {
	o.CouldComment = &v
}

// GetCover returns the Cover field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetCover() string {
	if o == nil || IsNil(o.Cover) {
		var ret string
		return ret
	}
	return *o.Cover
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetCoverOk() (*string, bool) {
	if o == nil || IsNil(o.Cover) {
		return nil, false
	}
	return o.Cover, true
}

// HasCover returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasCover() bool {
	if o != nil && !IsNil(o.Cover) {
		return true
	}

	return false
}

// SetCover gets a reference to the given string and assigns it to the Cover field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetCover(v string) {
	o.Cover = &v
}

// GetCtype returns the Ctype field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetCtype() string {
	if o == nil || IsNil(o.Ctype) {
		var ret string
		return ret
	}
	return *o.Ctype
}

// GetCtypeOk returns a tuple with the Ctype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetCtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Ctype) {
		return nil, false
	}
	return o.Ctype, true
}

// HasCtype returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasCtype() bool {
	if o != nil && !IsNil(o.Ctype) {
		return true
	}

	return false
}

// SetCtype gets a reference to the given string and assigns it to the Ctype field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetCtype(v string) {
	o.Ctype = &v
}

// GetExtTags returns the ExtTags field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetExtTags() string {
	if o == nil || IsNil(o.ExtTags) {
		var ret string
		return ret
	}
	return *o.ExtTags
}

// GetExtTagsOk returns a tuple with the ExtTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetExtTagsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtTags) {
		return nil, false
	}
	return o.ExtTags, true
}

// HasExtTags returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasExtTags() bool {
	if o != nil && !IsNil(o.ExtTags) {
		return true
	}

	return false
}

// SetExtTags gets a reference to the given string and assigns it to the ExtTags field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetExtTags(v string) {
	o.ExtTags = &v
}

// GetLoginIds returns the LoginIds field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetLoginIds() string {
	if o == nil || IsNil(o.LoginIds) {
		var ret string
		return ret
	}
	return *o.LoginIds
}

// GetLoginIdsOk returns a tuple with the LoginIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetLoginIdsOk() (*string, bool) {
	if o == nil || IsNil(o.LoginIds) {
		return nil, false
	}
	return o.LoginIds, true
}

// HasLoginIds returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasLoginIds() bool {
	if o != nil && !IsNil(o.LoginIds) {
		return true
	}

	return false
}

// SetLoginIds gets a reference to the given string and assigns it to the LoginIds field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetLoginIds(v string) {
	o.LoginIds = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayOpenPublicMessageContentCreateModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayOpenPublicMessageContentCreateModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayOpenPublicMessageContentCreateModel) SetTitle(v string) {
	o.Title = &v
}

func (o AlipayOpenPublicMessageContentCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicMessageContentCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Benefit) {
		toSerialize["benefit"] = o.Benefit
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
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

type NullableAlipayOpenPublicMessageContentCreateModel struct {
	value *AlipayOpenPublicMessageContentCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicMessageContentCreateModel) Get() *AlipayOpenPublicMessageContentCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicMessageContentCreateModel) Set(val *AlipayOpenPublicMessageContentCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMessageContentCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMessageContentCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMessageContentCreateModel(val *AlipayOpenPublicMessageContentCreateModel) *NullableAlipayOpenPublicMessageContentCreateModel {
	return &NullableAlipayOpenPublicMessageContentCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMessageContentCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMessageContentCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
