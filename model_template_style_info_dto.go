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

// checks if the TemplateStyleInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateStyleInfoDTO{}

// TemplateStyleInfoDTO struct for TemplateStyleInfoDTO
type TemplateStyleInfoDTO struct {
	// 背景图片Id，通过接口（alipay.offline.material.image.upload）上传图片    图片说明：2M以内，格式：bmp、png、jpeg、jpg、gif；  尺寸不小于1020*643px；  图片不得有圆角，不得拉伸变形
	BackgroundId *string `json:"background_id,omitempty"`
	// banner图片id。 通过接口（alipay.offline.material.image.upload）上传图片。
	BannerImgId *string `json:"banner_img_id,omitempty"`
	// banner跳转地址。
	BannerUrl *string `json:"banner_url,omitempty"`
	// 字体颜色（非背景色），只影响卡详情中部信息区域字体颜色
	BgColor *string `json:"bg_color,omitempty"`
	// 品牌商名称
	BrandName *string `json:"brand_name,omitempty"`
	// 钱包端显示名称
	CardShowName *string `json:"card_show_name,omitempty"`
	// 注意：此字段已废弃。
	Color *string `json:"color,omitempty"`
	// 如果为空则默认为list。
	ColumnInfoLayout *string `json:"column_info_layout,omitempty"`
	// 特色信息，用于领卡预览
	FeatureDescriptions []string `json:"feature_descriptions,omitempty"`
	// 设置是否在卡面展示（个人头像）图片信息，默认不展示；  当前仅用于身份验证信息类型的个人头像图片；  图片id随创建卡/更新卡时传入；  详见会员卡产品文档。
	FrontImageEnable *bool `json:"front_image_enable,omitempty"`
	// 设置是否在卡面展示文案信息，默认不展示；  文案信息分行展示，最多展示4行文案，每行文案分为label和value两部分；  文案实际内容随创建卡/更新卡时传入；  详见会员卡产品说明文档。
	FrontTextListEnable *bool `json:"front_text_list_enable,omitempty"`
	// logo的图片ID，通过接口（alipay.offline.material.image.upload）上传图片    图片说明：1M以内，格式bmp、png、jpeg、jpg、gif；  尺寸不小于500*500px的正方形；  请优先使用商家LOGO；
	LogoId *string `json:"logo_id,omitempty"`
	// 标语
	Slogan *string `json:"slogan,omitempty"`
	// 标语图片id， 通过接口（alipay.offline.material.image.upload）上传图片
	SloganImgId *string `json:"slogan_img_id,omitempty"`
}

// NewTemplateStyleInfoDTO instantiates a new TemplateStyleInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateStyleInfoDTO() *TemplateStyleInfoDTO {
	this := TemplateStyleInfoDTO{}
	return &this
}

// NewTemplateStyleInfoDTOWithDefaults instantiates a new TemplateStyleInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateStyleInfoDTOWithDefaults() *TemplateStyleInfoDTO {
	this := TemplateStyleInfoDTO{}
	return &this
}

// GetBackgroundId returns the BackgroundId field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetBackgroundId() string {
	if o == nil || IsNil(o.BackgroundId) {
		var ret string
		return ret
	}
	return *o.BackgroundId
}

// GetBackgroundIdOk returns a tuple with the BackgroundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetBackgroundIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundId) {
		return nil, false
	}
	return o.BackgroundId, true
}

// HasBackgroundId returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasBackgroundId() bool {
	if o != nil && !IsNil(o.BackgroundId) {
		return true
	}

	return false
}

// SetBackgroundId gets a reference to the given string and assigns it to the BackgroundId field.
func (o *TemplateStyleInfoDTO) SetBackgroundId(v string) {
	o.BackgroundId = &v
}

// GetBannerImgId returns the BannerImgId field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetBannerImgId() string {
	if o == nil || IsNil(o.BannerImgId) {
		var ret string
		return ret
	}
	return *o.BannerImgId
}

// GetBannerImgIdOk returns a tuple with the BannerImgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetBannerImgIdOk() (*string, bool) {
	if o == nil || IsNil(o.BannerImgId) {
		return nil, false
	}
	return o.BannerImgId, true
}

// HasBannerImgId returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasBannerImgId() bool {
	if o != nil && !IsNil(o.BannerImgId) {
		return true
	}

	return false
}

// SetBannerImgId gets a reference to the given string and assigns it to the BannerImgId field.
func (o *TemplateStyleInfoDTO) SetBannerImgId(v string) {
	o.BannerImgId = &v
}

// GetBannerUrl returns the BannerUrl field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetBannerUrl() string {
	if o == nil || IsNil(o.BannerUrl) {
		var ret string
		return ret
	}
	return *o.BannerUrl
}

// GetBannerUrlOk returns a tuple with the BannerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetBannerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BannerUrl) {
		return nil, false
	}
	return o.BannerUrl, true
}

// HasBannerUrl returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasBannerUrl() bool {
	if o != nil && !IsNil(o.BannerUrl) {
		return true
	}

	return false
}

// SetBannerUrl gets a reference to the given string and assigns it to the BannerUrl field.
func (o *TemplateStyleInfoDTO) SetBannerUrl(v string) {
	o.BannerUrl = &v
}

// GetBgColor returns the BgColor field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetBgColor() string {
	if o == nil || IsNil(o.BgColor) {
		var ret string
		return ret
	}
	return *o.BgColor
}

// GetBgColorOk returns a tuple with the BgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetBgColorOk() (*string, bool) {
	if o == nil || IsNil(o.BgColor) {
		return nil, false
	}
	return o.BgColor, true
}

// HasBgColor returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasBgColor() bool {
	if o != nil && !IsNil(o.BgColor) {
		return true
	}

	return false
}

// SetBgColor gets a reference to the given string and assigns it to the BgColor field.
func (o *TemplateStyleInfoDTO) SetBgColor(v string) {
	o.BgColor = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *TemplateStyleInfoDTO) SetBrandName(v string) {
	o.BrandName = &v
}

// GetCardShowName returns the CardShowName field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetCardShowName() string {
	if o == nil || IsNil(o.CardShowName) {
		var ret string
		return ret
	}
	return *o.CardShowName
}

// GetCardShowNameOk returns a tuple with the CardShowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetCardShowNameOk() (*string, bool) {
	if o == nil || IsNil(o.CardShowName) {
		return nil, false
	}
	return o.CardShowName, true
}

// HasCardShowName returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasCardShowName() bool {
	if o != nil && !IsNil(o.CardShowName) {
		return true
	}

	return false
}

// SetCardShowName gets a reference to the given string and assigns it to the CardShowName field.
func (o *TemplateStyleInfoDTO) SetCardShowName(v string) {
	o.CardShowName = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TemplateStyleInfoDTO) SetColor(v string) {
	o.Color = &v
}

// GetColumnInfoLayout returns the ColumnInfoLayout field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetColumnInfoLayout() string {
	if o == nil || IsNil(o.ColumnInfoLayout) {
		var ret string
		return ret
	}
	return *o.ColumnInfoLayout
}

// GetColumnInfoLayoutOk returns a tuple with the ColumnInfoLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetColumnInfoLayoutOk() (*string, bool) {
	if o == nil || IsNil(o.ColumnInfoLayout) {
		return nil, false
	}
	return o.ColumnInfoLayout, true
}

// HasColumnInfoLayout returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasColumnInfoLayout() bool {
	if o != nil && !IsNil(o.ColumnInfoLayout) {
		return true
	}

	return false
}

// SetColumnInfoLayout gets a reference to the given string and assigns it to the ColumnInfoLayout field.
func (o *TemplateStyleInfoDTO) SetColumnInfoLayout(v string) {
	o.ColumnInfoLayout = &v
}

// GetFeatureDescriptions returns the FeatureDescriptions field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetFeatureDescriptions() []string {
	if o == nil || IsNil(o.FeatureDescriptions) {
		var ret []string
		return ret
	}
	return o.FeatureDescriptions
}

// GetFeatureDescriptionsOk returns a tuple with the FeatureDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetFeatureDescriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.FeatureDescriptions) {
		return nil, false
	}
	return o.FeatureDescriptions, true
}

// HasFeatureDescriptions returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasFeatureDescriptions() bool {
	if o != nil && !IsNil(o.FeatureDescriptions) {
		return true
	}

	return false
}

// SetFeatureDescriptions gets a reference to the given []string and assigns it to the FeatureDescriptions field.
func (o *TemplateStyleInfoDTO) SetFeatureDescriptions(v []string) {
	o.FeatureDescriptions = v
}

// GetFrontImageEnable returns the FrontImageEnable field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetFrontImageEnable() bool {
	if o == nil || IsNil(o.FrontImageEnable) {
		var ret bool
		return ret
	}
	return *o.FrontImageEnable
}

// GetFrontImageEnableOk returns a tuple with the FrontImageEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetFrontImageEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.FrontImageEnable) {
		return nil, false
	}
	return o.FrontImageEnable, true
}

// HasFrontImageEnable returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasFrontImageEnable() bool {
	if o != nil && !IsNil(o.FrontImageEnable) {
		return true
	}

	return false
}

// SetFrontImageEnable gets a reference to the given bool and assigns it to the FrontImageEnable field.
func (o *TemplateStyleInfoDTO) SetFrontImageEnable(v bool) {
	o.FrontImageEnable = &v
}

// GetFrontTextListEnable returns the FrontTextListEnable field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetFrontTextListEnable() bool {
	if o == nil || IsNil(o.FrontTextListEnable) {
		var ret bool
		return ret
	}
	return *o.FrontTextListEnable
}

// GetFrontTextListEnableOk returns a tuple with the FrontTextListEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetFrontTextListEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.FrontTextListEnable) {
		return nil, false
	}
	return o.FrontTextListEnable, true
}

// HasFrontTextListEnable returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasFrontTextListEnable() bool {
	if o != nil && !IsNil(o.FrontTextListEnable) {
		return true
	}

	return false
}

// SetFrontTextListEnable gets a reference to the given bool and assigns it to the FrontTextListEnable field.
func (o *TemplateStyleInfoDTO) SetFrontTextListEnable(v bool) {
	o.FrontTextListEnable = &v
}

// GetLogoId returns the LogoId field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetLogoId() string {
	if o == nil || IsNil(o.LogoId) {
		var ret string
		return ret
	}
	return *o.LogoId
}

// GetLogoIdOk returns a tuple with the LogoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetLogoIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogoId) {
		return nil, false
	}
	return o.LogoId, true
}

// HasLogoId returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasLogoId() bool {
	if o != nil && !IsNil(o.LogoId) {
		return true
	}

	return false
}

// SetLogoId gets a reference to the given string and assigns it to the LogoId field.
func (o *TemplateStyleInfoDTO) SetLogoId(v string) {
	o.LogoId = &v
}

// GetSlogan returns the Slogan field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetSlogan() string {
	if o == nil || IsNil(o.Slogan) {
		var ret string
		return ret
	}
	return *o.Slogan
}

// GetSloganOk returns a tuple with the Slogan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetSloganOk() (*string, bool) {
	if o == nil || IsNil(o.Slogan) {
		return nil, false
	}
	return o.Slogan, true
}

// HasSlogan returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasSlogan() bool {
	if o != nil && !IsNil(o.Slogan) {
		return true
	}

	return false
}

// SetSlogan gets a reference to the given string and assigns it to the Slogan field.
func (o *TemplateStyleInfoDTO) SetSlogan(v string) {
	o.Slogan = &v
}

// GetSloganImgId returns the SloganImgId field value if set, zero value otherwise.
func (o *TemplateStyleInfoDTO) GetSloganImgId() string {
	if o == nil || IsNil(o.SloganImgId) {
		var ret string
		return ret
	}
	return *o.SloganImgId
}

// GetSloganImgIdOk returns a tuple with the SloganImgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStyleInfoDTO) GetSloganImgIdOk() (*string, bool) {
	if o == nil || IsNil(o.SloganImgId) {
		return nil, false
	}
	return o.SloganImgId, true
}

// HasSloganImgId returns a boolean if a field has been set.
func (o *TemplateStyleInfoDTO) HasSloganImgId() bool {
	if o != nil && !IsNil(o.SloganImgId) {
		return true
	}

	return false
}

// SetSloganImgId gets a reference to the given string and assigns it to the SloganImgId field.
func (o *TemplateStyleInfoDTO) SetSloganImgId(v string) {
	o.SloganImgId = &v
}

func (o TemplateStyleInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateStyleInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackgroundId) {
		toSerialize["background_id"] = o.BackgroundId
	}
	if !IsNil(o.BannerImgId) {
		toSerialize["banner_img_id"] = o.BannerImgId
	}
	if !IsNil(o.BannerUrl) {
		toSerialize["banner_url"] = o.BannerUrl
	}
	if !IsNil(o.BgColor) {
		toSerialize["bg_color"] = o.BgColor
	}
	if !IsNil(o.BrandName) {
		toSerialize["brand_name"] = o.BrandName
	}
	if !IsNil(o.CardShowName) {
		toSerialize["card_show_name"] = o.CardShowName
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.ColumnInfoLayout) {
		toSerialize["column_info_layout"] = o.ColumnInfoLayout
	}
	if !IsNil(o.FeatureDescriptions) {
		toSerialize["feature_descriptions"] = o.FeatureDescriptions
	}
	if !IsNil(o.FrontImageEnable) {
		toSerialize["front_image_enable"] = o.FrontImageEnable
	}
	if !IsNil(o.FrontTextListEnable) {
		toSerialize["front_text_list_enable"] = o.FrontTextListEnable
	}
	if !IsNil(o.LogoId) {
		toSerialize["logo_id"] = o.LogoId
	}
	if !IsNil(o.Slogan) {
		toSerialize["slogan"] = o.Slogan
	}
	if !IsNil(o.SloganImgId) {
		toSerialize["slogan_img_id"] = o.SloganImgId
	}
	return toSerialize, nil
}

type NullableTemplateStyleInfoDTO struct {
	value *TemplateStyleInfoDTO
	isSet bool
}

func (v NullableTemplateStyleInfoDTO) Get() *TemplateStyleInfoDTO {
	return v.value
}

func (v *NullableTemplateStyleInfoDTO) Set(val *TemplateStyleInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateStyleInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateStyleInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateStyleInfoDTO(val *TemplateStyleInfoDTO) *NullableTemplateStyleInfoDTO {
	return &NullableTemplateStyleInfoDTO{value: val, isSet: true}
}

func (v NullableTemplateStyleInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateStyleInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
