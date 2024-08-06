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

// checks if the VoucherDisplayPatternInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherDisplayPatternInfo{}

// VoucherDisplayPatternInfo struct for VoucherDisplayPatternInfo
type VoucherDisplayPatternInfo struct {
	// 商家logo 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_BRAND_LOGO，返回的resource_id即为该参数的值。
	BrandLogo *string `json:"brand_logo,omitempty"`
	// 商家品牌logo链接
	BrandLogoUrl *string `json:"brand_logo_url,omitempty"`
	// 商户品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// 客服电话。
	CustomerServiceMobile *string `json:"customer_service_mobile,omitempty"`
	// 客服链接。
	CustomerServiceUrl *string `json:"customer_service_url,omitempty"`
	// 券详细使用说明。
	VoucherDescription *string `json:"voucher_description,omitempty"`
	// 券详情链接
	VoucherDetailImageUrls []string `json:"voucher_detail_image_urls,omitempty"`
	// 券详细图列表，会展示在用户支付宝卡包券详情页 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE  ,接口返回的resource_id即为该参数的值 上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB。
	VoucherDetailImages []string `json:"voucher_detail_images,omitempty"`
	// 券详情页封面图，会展示在用户支付宝卡包券详情页。 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值。
	VoucherImage *string `json:"voucher_image,omitempty"`
	// 券封面链接
	VoucherImageUrl *string `json:"voucher_image_url,omitempty"`
	// 券名称。
	VoucherName *string `json:"voucher_name,omitempty"`
}

// NewVoucherDisplayPatternInfo instantiates a new VoucherDisplayPatternInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherDisplayPatternInfo() *VoucherDisplayPatternInfo {
	this := VoucherDisplayPatternInfo{}
	return &this
}

// NewVoucherDisplayPatternInfoWithDefaults instantiates a new VoucherDisplayPatternInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherDisplayPatternInfoWithDefaults() *VoucherDisplayPatternInfo {
	this := VoucherDisplayPatternInfo{}
	return &this
}

// GetBrandLogo returns the BrandLogo field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetBrandLogo() string {
	if o == nil || IsNil(o.BrandLogo) {
		var ret string
		return ret
	}
	return *o.BrandLogo
}

// GetBrandLogoOk returns a tuple with the BrandLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetBrandLogoOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogo) {
		return nil, false
	}
	return o.BrandLogo, true
}

// HasBrandLogo returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasBrandLogo() bool {
	if o != nil && !IsNil(o.BrandLogo) {
		return true
	}

	return false
}

// SetBrandLogo gets a reference to the given string and assigns it to the BrandLogo field.
func (o *VoucherDisplayPatternInfo) SetBrandLogo(v string) {
	o.BrandLogo = &v
}

// GetBrandLogoUrl returns the BrandLogoUrl field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetBrandLogoUrl() string {
	if o == nil || IsNil(o.BrandLogoUrl) {
		var ret string
		return ret
	}
	return *o.BrandLogoUrl
}

// GetBrandLogoUrlOk returns a tuple with the BrandLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetBrandLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoUrl) {
		return nil, false
	}
	return o.BrandLogoUrl, true
}

// HasBrandLogoUrl returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasBrandLogoUrl() bool {
	if o != nil && !IsNil(o.BrandLogoUrl) {
		return true
	}

	return false
}

// SetBrandLogoUrl gets a reference to the given string and assigns it to the BrandLogoUrl field.
func (o *VoucherDisplayPatternInfo) SetBrandLogoUrl(v string) {
	o.BrandLogoUrl = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *VoucherDisplayPatternInfo) SetBrandName(v string) {
	o.BrandName = &v
}

// GetCustomerServiceMobile returns the CustomerServiceMobile field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetCustomerServiceMobile() string {
	if o == nil || IsNil(o.CustomerServiceMobile) {
		var ret string
		return ret
	}
	return *o.CustomerServiceMobile
}

// GetCustomerServiceMobileOk returns a tuple with the CustomerServiceMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetCustomerServiceMobileOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerServiceMobile) {
		return nil, false
	}
	return o.CustomerServiceMobile, true
}

// HasCustomerServiceMobile returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasCustomerServiceMobile() bool {
	if o != nil && !IsNil(o.CustomerServiceMobile) {
		return true
	}

	return false
}

// SetCustomerServiceMobile gets a reference to the given string and assigns it to the CustomerServiceMobile field.
func (o *VoucherDisplayPatternInfo) SetCustomerServiceMobile(v string) {
	o.CustomerServiceMobile = &v
}

// GetCustomerServiceUrl returns the CustomerServiceUrl field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetCustomerServiceUrl() string {
	if o == nil || IsNil(o.CustomerServiceUrl) {
		var ret string
		return ret
	}
	return *o.CustomerServiceUrl
}

// GetCustomerServiceUrlOk returns a tuple with the CustomerServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetCustomerServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerServiceUrl) {
		return nil, false
	}
	return o.CustomerServiceUrl, true
}

// HasCustomerServiceUrl returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasCustomerServiceUrl() bool {
	if o != nil && !IsNil(o.CustomerServiceUrl) {
		return true
	}

	return false
}

// SetCustomerServiceUrl gets a reference to the given string and assigns it to the CustomerServiceUrl field.
func (o *VoucherDisplayPatternInfo) SetCustomerServiceUrl(v string) {
	o.CustomerServiceUrl = &v
}

// GetVoucherDescription returns the VoucherDescription field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetVoucherDescription() string {
	if o == nil || IsNil(o.VoucherDescription) {
		var ret string
		return ret
	}
	return *o.VoucherDescription
}

// GetVoucherDescriptionOk returns a tuple with the VoucherDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetVoucherDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherDescription) {
		return nil, false
	}
	return o.VoucherDescription, true
}

// HasVoucherDescription returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasVoucherDescription() bool {
	if o != nil && !IsNil(o.VoucherDescription) {
		return true
	}

	return false
}

// SetVoucherDescription gets a reference to the given string and assigns it to the VoucherDescription field.
func (o *VoucherDisplayPatternInfo) SetVoucherDescription(v string) {
	o.VoucherDescription = &v
}

// GetVoucherDetailImageUrls returns the VoucherDetailImageUrls field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetVoucherDetailImageUrls() []string {
	if o == nil || IsNil(o.VoucherDetailImageUrls) {
		var ret []string
		return ret
	}
	return o.VoucherDetailImageUrls
}

// GetVoucherDetailImageUrlsOk returns a tuple with the VoucherDetailImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetVoucherDetailImageUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.VoucherDetailImageUrls) {
		return nil, false
	}
	return o.VoucherDetailImageUrls, true
}

// HasVoucherDetailImageUrls returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasVoucherDetailImageUrls() bool {
	if o != nil && !IsNil(o.VoucherDetailImageUrls) {
		return true
	}

	return false
}

// SetVoucherDetailImageUrls gets a reference to the given []string and assigns it to the VoucherDetailImageUrls field.
func (o *VoucherDisplayPatternInfo) SetVoucherDetailImageUrls(v []string) {
	o.VoucherDetailImageUrls = v
}

// GetVoucherDetailImages returns the VoucherDetailImages field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetVoucherDetailImages() []string {
	if o == nil || IsNil(o.VoucherDetailImages) {
		var ret []string
		return ret
	}
	return o.VoucherDetailImages
}

// GetVoucherDetailImagesOk returns a tuple with the VoucherDetailImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetVoucherDetailImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.VoucherDetailImages) {
		return nil, false
	}
	return o.VoucherDetailImages, true
}

// HasVoucherDetailImages returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasVoucherDetailImages() bool {
	if o != nil && !IsNil(o.VoucherDetailImages) {
		return true
	}

	return false
}

// SetVoucherDetailImages gets a reference to the given []string and assigns it to the VoucherDetailImages field.
func (o *VoucherDisplayPatternInfo) SetVoucherDetailImages(v []string) {
	o.VoucherDetailImages = v
}

// GetVoucherImage returns the VoucherImage field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetVoucherImage() string {
	if o == nil || IsNil(o.VoucherImage) {
		var ret string
		return ret
	}
	return *o.VoucherImage
}

// GetVoucherImageOk returns a tuple with the VoucherImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetVoucherImageOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherImage) {
		return nil, false
	}
	return o.VoucherImage, true
}

// HasVoucherImage returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasVoucherImage() bool {
	if o != nil && !IsNil(o.VoucherImage) {
		return true
	}

	return false
}

// SetVoucherImage gets a reference to the given string and assigns it to the VoucherImage field.
func (o *VoucherDisplayPatternInfo) SetVoucherImage(v string) {
	o.VoucherImage = &v
}

// GetVoucherImageUrl returns the VoucherImageUrl field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetVoucherImageUrl() string {
	if o == nil || IsNil(o.VoucherImageUrl) {
		var ret string
		return ret
	}
	return *o.VoucherImageUrl
}

// GetVoucherImageUrlOk returns a tuple with the VoucherImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetVoucherImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherImageUrl) {
		return nil, false
	}
	return o.VoucherImageUrl, true
}

// HasVoucherImageUrl returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasVoucherImageUrl() bool {
	if o != nil && !IsNil(o.VoucherImageUrl) {
		return true
	}

	return false
}

// SetVoucherImageUrl gets a reference to the given string and assigns it to the VoucherImageUrl field.
func (o *VoucherDisplayPatternInfo) SetVoucherImageUrl(v string) {
	o.VoucherImageUrl = &v
}

// GetVoucherName returns the VoucherName field value if set, zero value otherwise.
func (o *VoucherDisplayPatternInfo) GetVoucherName() string {
	if o == nil || IsNil(o.VoucherName) {
		var ret string
		return ret
	}
	return *o.VoucherName
}

// GetVoucherNameOk returns a tuple with the VoucherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDisplayPatternInfo) GetVoucherNameOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherName) {
		return nil, false
	}
	return o.VoucherName, true
}

// HasVoucherName returns a boolean if a field has been set.
func (o *VoucherDisplayPatternInfo) HasVoucherName() bool {
	if o != nil && !IsNil(o.VoucherName) {
		return true
	}

	return false
}

// SetVoucherName gets a reference to the given string and assigns it to the VoucherName field.
func (o *VoucherDisplayPatternInfo) SetVoucherName(v string) {
	o.VoucherName = &v
}

func (o VoucherDisplayPatternInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherDisplayPatternInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogo) {
		toSerialize["brand_logo"] = o.BrandLogo
	}
	if !IsNil(o.BrandLogoUrl) {
		toSerialize["brand_logo_url"] = o.BrandLogoUrl
	}
	if !IsNil(o.BrandName) {
		toSerialize["brand_name"] = o.BrandName
	}
	if !IsNil(o.CustomerServiceMobile) {
		toSerialize["customer_service_mobile"] = o.CustomerServiceMobile
	}
	if !IsNil(o.CustomerServiceUrl) {
		toSerialize["customer_service_url"] = o.CustomerServiceUrl
	}
	if !IsNil(o.VoucherDescription) {
		toSerialize["voucher_description"] = o.VoucherDescription
	}
	if !IsNil(o.VoucherDetailImageUrls) {
		toSerialize["voucher_detail_image_urls"] = o.VoucherDetailImageUrls
	}
	if !IsNil(o.VoucherDetailImages) {
		toSerialize["voucher_detail_images"] = o.VoucherDetailImages
	}
	if !IsNil(o.VoucherImage) {
		toSerialize["voucher_image"] = o.VoucherImage
	}
	if !IsNil(o.VoucherImageUrl) {
		toSerialize["voucher_image_url"] = o.VoucherImageUrl
	}
	if !IsNil(o.VoucherName) {
		toSerialize["voucher_name"] = o.VoucherName
	}
	return toSerialize, nil
}

type NullableVoucherDisplayPatternInfo struct {
	value *VoucherDisplayPatternInfo
	isSet bool
}

func (v NullableVoucherDisplayPatternInfo) Get() *VoucherDisplayPatternInfo {
	return v.value
}

func (v *NullableVoucherDisplayPatternInfo) Set(val *VoucherDisplayPatternInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherDisplayPatternInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherDisplayPatternInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherDisplayPatternInfo(val *VoucherDisplayPatternInfo) *NullableVoucherDisplayPatternInfo {
	return &NullableVoucherDisplayPatternInfo{value: val, isSet: true}
}

func (v NullableVoucherDisplayPatternInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherDisplayPatternInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


