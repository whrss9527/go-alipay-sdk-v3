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

// checks if the OpenCertifyMerchantConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenCertifyMerchantConfig{}

// OpenCertifyMerchantConfig struct for OpenCertifyMerchantConfig
type OpenCertifyMerchantConfig struct {
	// 用于开放认证授权
	AuthScope *string `json:"auth_scope,omitempty"`
	// 用于指定授权类型，与auth_scope配合使用
	AuthType *string `json:"auth_type,omitempty"`
	// 不传默认为reserve
	FaceReserveStrategy *string `json:"face_reserve_strategy,omitempty"`
	// 若有特殊人脸等级采集要求，可指定等级
	FacialPictureLevel *string `json:"facial_picture_level,omitempty"`
	// 用于授权二级商户操作
	LinkedMerchantAppId *string `json:"linked_merchant_app_id,omitempty"`
	// 用于指定展示的商户logo
	LinkedMerchantLogoUrl *string `json:"linked_merchant_logo_url,omitempty"`
	// 用于指定展示的商户名称
	LinkedMerchantName *string `json:"linked_merchant_name,omitempty"`
	// 在拥有该权限前提下，用于商户控制是否透出活体人脸
	OutPutFacialPicture *bool `json:"out_put_facial_picture,omitempty"`
	// 认证成功后需要跳转的地址，一般为商户业务页面；若无跳转地址可填空字符\"\";
	ReturnUrl *string `json:"return_url,omitempty"`
}

// NewOpenCertifyMerchantConfig instantiates a new OpenCertifyMerchantConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenCertifyMerchantConfig() *OpenCertifyMerchantConfig {
	this := OpenCertifyMerchantConfig{}
	return &this
}

// NewOpenCertifyMerchantConfigWithDefaults instantiates a new OpenCertifyMerchantConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenCertifyMerchantConfigWithDefaults() *OpenCertifyMerchantConfig {
	this := OpenCertifyMerchantConfig{}
	return &this
}

// GetAuthScope returns the AuthScope field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetAuthScope() string {
	if o == nil || IsNil(o.AuthScope) {
		var ret string
		return ret
	}
	return *o.AuthScope
}

// GetAuthScopeOk returns a tuple with the AuthScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetAuthScopeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthScope) {
		return nil, false
	}
	return o.AuthScope, true
}

// HasAuthScope returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasAuthScope() bool {
	if o != nil && !IsNil(o.AuthScope) {
		return true
	}

	return false
}

// SetAuthScope gets a reference to the given string and assigns it to the AuthScope field.
func (o *OpenCertifyMerchantConfig) SetAuthScope(v string) {
	o.AuthScope = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *OpenCertifyMerchantConfig) SetAuthType(v string) {
	o.AuthType = &v
}

// GetFaceReserveStrategy returns the FaceReserveStrategy field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetFaceReserveStrategy() string {
	if o == nil || IsNil(o.FaceReserveStrategy) {
		var ret string
		return ret
	}
	return *o.FaceReserveStrategy
}

// GetFaceReserveStrategyOk returns a tuple with the FaceReserveStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetFaceReserveStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.FaceReserveStrategy) {
		return nil, false
	}
	return o.FaceReserveStrategy, true
}

// HasFaceReserveStrategy returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasFaceReserveStrategy() bool {
	if o != nil && !IsNil(o.FaceReserveStrategy) {
		return true
	}

	return false
}

// SetFaceReserveStrategy gets a reference to the given string and assigns it to the FaceReserveStrategy field.
func (o *OpenCertifyMerchantConfig) SetFaceReserveStrategy(v string) {
	o.FaceReserveStrategy = &v
}

// GetFacialPictureLevel returns the FacialPictureLevel field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetFacialPictureLevel() string {
	if o == nil || IsNil(o.FacialPictureLevel) {
		var ret string
		return ret
	}
	return *o.FacialPictureLevel
}

// GetFacialPictureLevelOk returns a tuple with the FacialPictureLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetFacialPictureLevelOk() (*string, bool) {
	if o == nil || IsNil(o.FacialPictureLevel) {
		return nil, false
	}
	return o.FacialPictureLevel, true
}

// HasFacialPictureLevel returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasFacialPictureLevel() bool {
	if o != nil && !IsNil(o.FacialPictureLevel) {
		return true
	}

	return false
}

// SetFacialPictureLevel gets a reference to the given string and assigns it to the FacialPictureLevel field.
func (o *OpenCertifyMerchantConfig) SetFacialPictureLevel(v string) {
	o.FacialPictureLevel = &v
}

// GetLinkedMerchantAppId returns the LinkedMerchantAppId field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetLinkedMerchantAppId() string {
	if o == nil || IsNil(o.LinkedMerchantAppId) {
		var ret string
		return ret
	}
	return *o.LinkedMerchantAppId
}

// GetLinkedMerchantAppIdOk returns a tuple with the LinkedMerchantAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetLinkedMerchantAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedMerchantAppId) {
		return nil, false
	}
	return o.LinkedMerchantAppId, true
}

// HasLinkedMerchantAppId returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasLinkedMerchantAppId() bool {
	if o != nil && !IsNil(o.LinkedMerchantAppId) {
		return true
	}

	return false
}

// SetLinkedMerchantAppId gets a reference to the given string and assigns it to the LinkedMerchantAppId field.
func (o *OpenCertifyMerchantConfig) SetLinkedMerchantAppId(v string) {
	o.LinkedMerchantAppId = &v
}

// GetLinkedMerchantLogoUrl returns the LinkedMerchantLogoUrl field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetLinkedMerchantLogoUrl() string {
	if o == nil || IsNil(o.LinkedMerchantLogoUrl) {
		var ret string
		return ret
	}
	return *o.LinkedMerchantLogoUrl
}

// GetLinkedMerchantLogoUrlOk returns a tuple with the LinkedMerchantLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetLinkedMerchantLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedMerchantLogoUrl) {
		return nil, false
	}
	return o.LinkedMerchantLogoUrl, true
}

// HasLinkedMerchantLogoUrl returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasLinkedMerchantLogoUrl() bool {
	if o != nil && !IsNil(o.LinkedMerchantLogoUrl) {
		return true
	}

	return false
}

// SetLinkedMerchantLogoUrl gets a reference to the given string and assigns it to the LinkedMerchantLogoUrl field.
func (o *OpenCertifyMerchantConfig) SetLinkedMerchantLogoUrl(v string) {
	o.LinkedMerchantLogoUrl = &v
}

// GetLinkedMerchantName returns the LinkedMerchantName field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetLinkedMerchantName() string {
	if o == nil || IsNil(o.LinkedMerchantName) {
		var ret string
		return ret
	}
	return *o.LinkedMerchantName
}

// GetLinkedMerchantNameOk returns a tuple with the LinkedMerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetLinkedMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedMerchantName) {
		return nil, false
	}
	return o.LinkedMerchantName, true
}

// HasLinkedMerchantName returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasLinkedMerchantName() bool {
	if o != nil && !IsNil(o.LinkedMerchantName) {
		return true
	}

	return false
}

// SetLinkedMerchantName gets a reference to the given string and assigns it to the LinkedMerchantName field.
func (o *OpenCertifyMerchantConfig) SetLinkedMerchantName(v string) {
	o.LinkedMerchantName = &v
}

// GetOutPutFacialPicture returns the OutPutFacialPicture field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetOutPutFacialPicture() bool {
	if o == nil || IsNil(o.OutPutFacialPicture) {
		var ret bool
		return ret
	}
	return *o.OutPutFacialPicture
}

// GetOutPutFacialPictureOk returns a tuple with the OutPutFacialPicture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetOutPutFacialPictureOk() (*bool, bool) {
	if o == nil || IsNil(o.OutPutFacialPicture) {
		return nil, false
	}
	return o.OutPutFacialPicture, true
}

// HasOutPutFacialPicture returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasOutPutFacialPicture() bool {
	if o != nil && !IsNil(o.OutPutFacialPicture) {
		return true
	}

	return false
}

// SetOutPutFacialPicture gets a reference to the given bool and assigns it to the OutPutFacialPicture field.
func (o *OpenCertifyMerchantConfig) SetOutPutFacialPicture(v bool) {
	o.OutPutFacialPicture = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *OpenCertifyMerchantConfig) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenCertifyMerchantConfig) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *OpenCertifyMerchantConfig) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *OpenCertifyMerchantConfig) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

func (o OpenCertifyMerchantConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenCertifyMerchantConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthScope) {
		toSerialize["auth_scope"] = o.AuthScope
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.FaceReserveStrategy) {
		toSerialize["face_reserve_strategy"] = o.FaceReserveStrategy
	}
	if !IsNil(o.FacialPictureLevel) {
		toSerialize["facial_picture_level"] = o.FacialPictureLevel
	}
	if !IsNil(o.LinkedMerchantAppId) {
		toSerialize["linked_merchant_app_id"] = o.LinkedMerchantAppId
	}
	if !IsNil(o.LinkedMerchantLogoUrl) {
		toSerialize["linked_merchant_logo_url"] = o.LinkedMerchantLogoUrl
	}
	if !IsNil(o.LinkedMerchantName) {
		toSerialize["linked_merchant_name"] = o.LinkedMerchantName
	}
	if !IsNil(o.OutPutFacialPicture) {
		toSerialize["out_put_facial_picture"] = o.OutPutFacialPicture
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["return_url"] = o.ReturnUrl
	}
	return toSerialize, nil
}

type NullableOpenCertifyMerchantConfig struct {
	value *OpenCertifyMerchantConfig
	isSet bool
}

func (v NullableOpenCertifyMerchantConfig) Get() *OpenCertifyMerchantConfig {
	return v.value
}

func (v *NullableOpenCertifyMerchantConfig) Set(val *OpenCertifyMerchantConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenCertifyMerchantConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenCertifyMerchantConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenCertifyMerchantConfig(val *OpenCertifyMerchantConfig) *NullableOpenCertifyMerchantConfig {
	return &NullableOpenCertifyMerchantConfig{value: val, isSet: true}
}

func (v NullableOpenCertifyMerchantConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenCertifyMerchantConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


