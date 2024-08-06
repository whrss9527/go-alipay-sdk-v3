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

// checks if the TemplateColumnInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateColumnInfoDTO{}

// TemplateColumnInfoDTO struct for TemplateColumnInfoDTO
type TemplateColumnInfoDTO struct {
	// 标准栏位：行为由支付宝统一定，同时已经分配标准Code  BALANCE：会员卡余额  POINT：积分  LEVEL：等级  TELEPHONE：联系方式  自定义栏位：行为由商户定义，自定义Code码（只要无重复）
	Code *string `json:"code,omitempty"`
	// 若template_style_info.column_info_layout 的值为grid，此项为宫格项所属分组标题。可空。如果需要展示该项，还需支付宝内部进行特殊配置。
	GroupTitle *string `json:"group_title,omitempty"`
	// 当template_style_info.column_info_layout 的值为grid时，此参数必填。此项为宫格项的展示icon。通过接口（alipay.offline.material.image.upload）上传图片。
	IconId *string `json:"icon_id,omitempty"`
	MoreInfo *MoreInfoDTO `json:"more_info,omitempty"`
	// 1、openNative：打开二级页面，展现 more中descs  2、openWeb：打开URL  3、staticinfo：静态信息  注意：  不填则默认staticinfo；  标准code尽量使用staticinfo，例如TELEPHONE商家电话栏位就只支持staticinfo；
	OperateType *string `json:"operate_type,omitempty"`
	// 只有当template_style_info.column_info_layout 的值为grid时，此参数有效。此项为宫格项标签，最多只会展示一个标签。
	Tag *string `json:"tag,omitempty"`
	// 栏目标题
	Title *string `json:"title,omitempty"`
	// 卡包详情页面，卡栏位右边展现的值    TELEPHONE栏位的商家联系电话号码由此value字段传入
	Value *string `json:"value,omitempty"`
}

// NewTemplateColumnInfoDTO instantiates a new TemplateColumnInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateColumnInfoDTO() *TemplateColumnInfoDTO {
	this := TemplateColumnInfoDTO{}
	return &this
}

// NewTemplateColumnInfoDTOWithDefaults instantiates a new TemplateColumnInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateColumnInfoDTOWithDefaults() *TemplateColumnInfoDTO {
	this := TemplateColumnInfoDTO{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TemplateColumnInfoDTO) SetCode(v string) {
	o.Code = &v
}

// GetGroupTitle returns the GroupTitle field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetGroupTitle() string {
	if o == nil || IsNil(o.GroupTitle) {
		var ret string
		return ret
	}
	return *o.GroupTitle
}

// GetGroupTitleOk returns a tuple with the GroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetGroupTitleOk() (*string, bool) {
	if o == nil || IsNil(o.GroupTitle) {
		return nil, false
	}
	return o.GroupTitle, true
}

// HasGroupTitle returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasGroupTitle() bool {
	if o != nil && !IsNil(o.GroupTitle) {
		return true
	}

	return false
}

// SetGroupTitle gets a reference to the given string and assigns it to the GroupTitle field.
func (o *TemplateColumnInfoDTO) SetGroupTitle(v string) {
	o.GroupTitle = &v
}

// GetIconId returns the IconId field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetIconId() string {
	if o == nil || IsNil(o.IconId) {
		var ret string
		return ret
	}
	return *o.IconId
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetIconIdOk() (*string, bool) {
	if o == nil || IsNil(o.IconId) {
		return nil, false
	}
	return o.IconId, true
}

// HasIconId returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasIconId() bool {
	if o != nil && !IsNil(o.IconId) {
		return true
	}

	return false
}

// SetIconId gets a reference to the given string and assigns it to the IconId field.
func (o *TemplateColumnInfoDTO) SetIconId(v string) {
	o.IconId = &v
}

// GetMoreInfo returns the MoreInfo field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetMoreInfo() MoreInfoDTO {
	if o == nil || IsNil(o.MoreInfo) {
		var ret MoreInfoDTO
		return ret
	}
	return *o.MoreInfo
}

// GetMoreInfoOk returns a tuple with the MoreInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetMoreInfoOk() (*MoreInfoDTO, bool) {
	if o == nil || IsNil(o.MoreInfo) {
		return nil, false
	}
	return o.MoreInfo, true
}

// HasMoreInfo returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasMoreInfo() bool {
	if o != nil && !IsNil(o.MoreInfo) {
		return true
	}

	return false
}

// SetMoreInfo gets a reference to the given MoreInfoDTO and assigns it to the MoreInfo field.
func (o *TemplateColumnInfoDTO) SetMoreInfo(v MoreInfoDTO) {
	o.MoreInfo = &v
}

// GetOperateType returns the OperateType field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetOperateType() string {
	if o == nil || IsNil(o.OperateType) {
		var ret string
		return ret
	}
	return *o.OperateType
}

// GetOperateTypeOk returns a tuple with the OperateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetOperateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OperateType) {
		return nil, false
	}
	return o.OperateType, true
}

// HasOperateType returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasOperateType() bool {
	if o != nil && !IsNil(o.OperateType) {
		return true
	}

	return false
}

// SetOperateType gets a reference to the given string and assigns it to the OperateType field.
func (o *TemplateColumnInfoDTO) SetOperateType(v string) {
	o.OperateType = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *TemplateColumnInfoDTO) SetTag(v string) {
	o.Tag = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TemplateColumnInfoDTO) SetTitle(v string) {
	o.Title = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TemplateColumnInfoDTO) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateColumnInfoDTO) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TemplateColumnInfoDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TemplateColumnInfoDTO) SetValue(v string) {
	o.Value = &v
}

func (o TemplateColumnInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateColumnInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.GroupTitle) {
		toSerialize["group_title"] = o.GroupTitle
	}
	if !IsNil(o.IconId) {
		toSerialize["icon_id"] = o.IconId
	}
	if !IsNil(o.MoreInfo) {
		toSerialize["more_info"] = o.MoreInfo
	}
	if !IsNil(o.OperateType) {
		toSerialize["operate_type"] = o.OperateType
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTemplateColumnInfoDTO struct {
	value *TemplateColumnInfoDTO
	isSet bool
}

func (v NullableTemplateColumnInfoDTO) Get() *TemplateColumnInfoDTO {
	return v.value
}

func (v *NullableTemplateColumnInfoDTO) Set(val *TemplateColumnInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateColumnInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateColumnInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateColumnInfoDTO(val *TemplateColumnInfoDTO) *NullableTemplateColumnInfoDTO {
	return &NullableTemplateColumnInfoDTO{value: val, isSet: true}
}

func (v NullableTemplateColumnInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateColumnInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


