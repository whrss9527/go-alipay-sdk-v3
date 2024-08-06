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

// checks if the AlipayOpenPublicInfoModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicInfoModifyModel{}

// AlipayOpenPublicInfoModifyModel struct for AlipayOpenPublicInfoModifyModel
type AlipayOpenPublicInfoModifyModel struct {
	// 生活号名称，2-20个字之间。注意： * 不得含有违反法律法规和公序良俗的相关信息； * 不得侵害他人名誉权、知识产权、商业秘密等合法权利； * 不得以太过广泛的、或产品、行业词组来命名，如：女装、皮革批发； * 不得以实名认证的媒体资质账号创建服务窗，或媒体相关名称命名服务窗，如：XX电视台、XX杂志等。
	AppName *string `json:"app_name,omitempty"`
	// 授权运营书图片地址，企业商户若为被经营方授权，需上传加盖公章的扫描件。支持 .jpg、 .jpeg、 .png 格式，需小于1M。使用 <a href=\"https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload\">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。
	AuthPic *string `json:"auth_pic,omitempty"`
	// 背景图片地址，建议尺寸 1600 x 1000px，支持.jpg .jpeg .png格式，小于1M。使用 <a href=\"https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload\">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。
	BackgroundUrl *string `json:"background_url,omitempty"`
	// 简介
	Introduction *string `json:"introduction,omitempty"`
	// 营业执照地址，建议尺寸 320*320 px，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 <a href=\"https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload\">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。
	LicenseUrl *string `json:"license_url,omitempty"`
	// 生活号头像地址，建议尺寸 320*320 px，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 <a href=\"https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload\">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。
	LogoUrl *string `json:"logo_url,omitempty"`
	// 通知地址
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 欢迎语
	PublicGreeting *string `json:"public_greeting,omitempty"`
	// 门店照片地址，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 <a href=\"https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload\">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。
	ShopPics []string `json:"shop_pics,omitempty"`
}

// NewAlipayOpenPublicInfoModifyModel instantiates a new AlipayOpenPublicInfoModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicInfoModifyModel() *AlipayOpenPublicInfoModifyModel {
	this := AlipayOpenPublicInfoModifyModel{}
	return &this
}

// NewAlipayOpenPublicInfoModifyModelWithDefaults instantiates a new AlipayOpenPublicInfoModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicInfoModifyModelWithDefaults() *AlipayOpenPublicInfoModifyModel {
	this := AlipayOpenPublicInfoModifyModel{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AlipayOpenPublicInfoModifyModel) SetAppName(v string) {
	o.AppName = &v
}

// GetAuthPic returns the AuthPic field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetAuthPic() string {
	if o == nil || IsNil(o.AuthPic) {
		var ret string
		return ret
	}
	return *o.AuthPic
}

// GetAuthPicOk returns a tuple with the AuthPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetAuthPicOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPic) {
		return nil, false
	}
	return o.AuthPic, true
}

// HasAuthPic returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasAuthPic() bool {
	if o != nil && !IsNil(o.AuthPic) {
		return true
	}

	return false
}

// SetAuthPic gets a reference to the given string and assigns it to the AuthPic field.
func (o *AlipayOpenPublicInfoModifyModel) SetAuthPic(v string) {
	o.AuthPic = &v
}

// GetBackgroundUrl returns the BackgroundUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetBackgroundUrl() string {
	if o == nil || IsNil(o.BackgroundUrl) {
		var ret string
		return ret
	}
	return *o.BackgroundUrl
}

// GetBackgroundUrlOk returns a tuple with the BackgroundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetBackgroundUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundUrl) {
		return nil, false
	}
	return o.BackgroundUrl, true
}

// HasBackgroundUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasBackgroundUrl() bool {
	if o != nil && !IsNil(o.BackgroundUrl) {
		return true
	}

	return false
}

// SetBackgroundUrl gets a reference to the given string and assigns it to the BackgroundUrl field.
func (o *AlipayOpenPublicInfoModifyModel) SetBackgroundUrl(v string) {
	o.BackgroundUrl = &v
}

// GetIntroduction returns the Introduction field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetIntroduction() string {
	if o == nil || IsNil(o.Introduction) {
		var ret string
		return ret
	}
	return *o.Introduction
}

// GetIntroductionOk returns a tuple with the Introduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetIntroductionOk() (*string, bool) {
	if o == nil || IsNil(o.Introduction) {
		return nil, false
	}
	return o.Introduction, true
}

// HasIntroduction returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasIntroduction() bool {
	if o != nil && !IsNil(o.Introduction) {
		return true
	}

	return false
}

// SetIntroduction gets a reference to the given string and assigns it to the Introduction field.
func (o *AlipayOpenPublicInfoModifyModel) SetIntroduction(v string) {
	o.Introduction = &v
}

// GetLicenseUrl returns the LicenseUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetLicenseUrl() string {
	if o == nil || IsNil(o.LicenseUrl) {
		var ret string
		return ret
	}
	return *o.LicenseUrl
}

// GetLicenseUrlOk returns a tuple with the LicenseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetLicenseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseUrl) {
		return nil, false
	}
	return o.LicenseUrl, true
}

// HasLicenseUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasLicenseUrl() bool {
	if o != nil && !IsNil(o.LicenseUrl) {
		return true
	}

	return false
}

// SetLicenseUrl gets a reference to the given string and assigns it to the LicenseUrl field.
func (o *AlipayOpenPublicInfoModifyModel) SetLicenseUrl(v string) {
	o.LicenseUrl = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *AlipayOpenPublicInfoModifyModel) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetNotifyUrl() string {
	if o == nil || IsNil(o.NotifyUrl) {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetNotifyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyUrl) {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasNotifyUrl() bool {
	if o != nil && !IsNil(o.NotifyUrl) {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *AlipayOpenPublicInfoModifyModel) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

// GetPublicGreeting returns the PublicGreeting field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetPublicGreeting() string {
	if o == nil || IsNil(o.PublicGreeting) {
		var ret string
		return ret
	}
	return *o.PublicGreeting
}

// GetPublicGreetingOk returns a tuple with the PublicGreeting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetPublicGreetingOk() (*string, bool) {
	if o == nil || IsNil(o.PublicGreeting) {
		return nil, false
	}
	return o.PublicGreeting, true
}

// HasPublicGreeting returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasPublicGreeting() bool {
	if o != nil && !IsNil(o.PublicGreeting) {
		return true
	}

	return false
}

// SetPublicGreeting gets a reference to the given string and assigns it to the PublicGreeting field.
func (o *AlipayOpenPublicInfoModifyModel) SetPublicGreeting(v string) {
	o.PublicGreeting = &v
}

// GetShopPics returns the ShopPics field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoModifyModel) GetShopPics() []string {
	if o == nil || IsNil(o.ShopPics) {
		var ret []string
		return ret
	}
	return o.ShopPics
}

// GetShopPicsOk returns a tuple with the ShopPics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoModifyModel) GetShopPicsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShopPics) {
		return nil, false
	}
	return o.ShopPics, true
}

// HasShopPics returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoModifyModel) HasShopPics() bool {
	if o != nil && !IsNil(o.ShopPics) {
		return true
	}

	return false
}

// SetShopPics gets a reference to the given []string and assigns it to the ShopPics field.
func (o *AlipayOpenPublicInfoModifyModel) SetShopPics(v []string) {
	o.ShopPics = v
}

func (o AlipayOpenPublicInfoModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicInfoModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AuthPic) {
		toSerialize["auth_pic"] = o.AuthPic
	}
	if !IsNil(o.BackgroundUrl) {
		toSerialize["background_url"] = o.BackgroundUrl
	}
	if !IsNil(o.Introduction) {
		toSerialize["introduction"] = o.Introduction
	}
	if !IsNil(o.LicenseUrl) {
		toSerialize["license_url"] = o.LicenseUrl
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if !IsNil(o.PublicGreeting) {
		toSerialize["public_greeting"] = o.PublicGreeting
	}
	if !IsNil(o.ShopPics) {
		toSerialize["shop_pics"] = o.ShopPics
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicInfoModifyModel struct {
	value *AlipayOpenPublicInfoModifyModel
	isSet bool
}

func (v NullableAlipayOpenPublicInfoModifyModel) Get() *AlipayOpenPublicInfoModifyModel {
	return v.value
}

func (v *NullableAlipayOpenPublicInfoModifyModel) Set(val *AlipayOpenPublicInfoModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicInfoModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicInfoModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicInfoModifyModel(val *AlipayOpenPublicInfoModifyModel) *NullableAlipayOpenPublicInfoModifyModel {
	return &NullableAlipayOpenPublicInfoModifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicInfoModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicInfoModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


