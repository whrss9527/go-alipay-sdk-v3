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

// checks if the VoucherAvailableGoodsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableGoodsInfo{}

// VoucherAvailableGoodsInfo struct for VoucherAvailableGoodsInfo
type VoucherAvailableGoodsInfo struct {
	// 可优惠商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，可以使用优惠券。
	AvailableGoodsSkuIds []string `json:"available_goods_sku_ids,omitempty"`
	// 不可核销商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，不可以使用优惠券。
	ExcludeGoodsSkuIds []string `json:"exclude_goods_sku_ids,omitempty"`
	// 商品描述信息。 用于券面展示，向用户介绍商品
	GoodsDescription *string `json:"goods_description,omitempty"`
	// 商品详情，会展示在特定渠道(如服务市场团购业务插件的套餐明细)。
	GoodsDetail *string `json:"goods_detail,omitempty"`
	// 商品详情图片列表，会展示在特定渠道(如服务市场团购业务插件的头图)
	GoodsDetailImages []string `json:"goods_detail_images,omitempty"`
	// 商品详情富文本描述
	GoodsDetailRichDescription *string `json:"goods_detail_rich_description,omitempty"`
	// 商品名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// 原价。说明：该字段可不填，填入商品名称goods_name则必填;
	OriginAmount *string `json:"origin_amount,omitempty"`
}

// NewVoucherAvailableGoodsInfo instantiates a new VoucherAvailableGoodsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableGoodsInfo() *VoucherAvailableGoodsInfo {
	this := VoucherAvailableGoodsInfo{}
	return &this
}

// NewVoucherAvailableGoodsInfoWithDefaults instantiates a new VoucherAvailableGoodsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableGoodsInfoWithDefaults() *VoucherAvailableGoodsInfo {
	this := VoucherAvailableGoodsInfo{}
	return &this
}

// GetAvailableGoodsSkuIds returns the AvailableGoodsSkuIds field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetAvailableGoodsSkuIds() []string {
	if o == nil || IsNil(o.AvailableGoodsSkuIds) {
		var ret []string
		return ret
	}
	return o.AvailableGoodsSkuIds
}

// GetAvailableGoodsSkuIdsOk returns a tuple with the AvailableGoodsSkuIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetAvailableGoodsSkuIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableGoodsSkuIds) {
		return nil, false
	}
	return o.AvailableGoodsSkuIds, true
}

// HasAvailableGoodsSkuIds returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasAvailableGoodsSkuIds() bool {
	if o != nil && !IsNil(o.AvailableGoodsSkuIds) {
		return true
	}

	return false
}

// SetAvailableGoodsSkuIds gets a reference to the given []string and assigns it to the AvailableGoodsSkuIds field.
func (o *VoucherAvailableGoodsInfo) SetAvailableGoodsSkuIds(v []string) {
	o.AvailableGoodsSkuIds = v
}

// GetExcludeGoodsSkuIds returns the ExcludeGoodsSkuIds field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetExcludeGoodsSkuIds() []string {
	if o == nil || IsNil(o.ExcludeGoodsSkuIds) {
		var ret []string
		return ret
	}
	return o.ExcludeGoodsSkuIds
}

// GetExcludeGoodsSkuIdsOk returns a tuple with the ExcludeGoodsSkuIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetExcludeGoodsSkuIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeGoodsSkuIds) {
		return nil, false
	}
	return o.ExcludeGoodsSkuIds, true
}

// HasExcludeGoodsSkuIds returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasExcludeGoodsSkuIds() bool {
	if o != nil && !IsNil(o.ExcludeGoodsSkuIds) {
		return true
	}

	return false
}

// SetExcludeGoodsSkuIds gets a reference to the given []string and assigns it to the ExcludeGoodsSkuIds field.
func (o *VoucherAvailableGoodsInfo) SetExcludeGoodsSkuIds(v []string) {
	o.ExcludeGoodsSkuIds = v
}

// GetGoodsDescription returns the GoodsDescription field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetGoodsDescription() string {
	if o == nil || IsNil(o.GoodsDescription) {
		var ret string
		return ret
	}
	return *o.GoodsDescription
}

// GetGoodsDescriptionOk returns a tuple with the GoodsDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetGoodsDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.GoodsDescription) {
		return nil, false
	}
	return o.GoodsDescription, true
}

// HasGoodsDescription returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasGoodsDescription() bool {
	if o != nil && !IsNil(o.GoodsDescription) {
		return true
	}

	return false
}

// SetGoodsDescription gets a reference to the given string and assigns it to the GoodsDescription field.
func (o *VoucherAvailableGoodsInfo) SetGoodsDescription(v string) {
	o.GoodsDescription = &v
}

// GetGoodsDetail returns the GoodsDetail field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetGoodsDetail() string {
	if o == nil || IsNil(o.GoodsDetail) {
		var ret string
		return ret
	}
	return *o.GoodsDetail
}

// GetGoodsDetailOk returns a tuple with the GoodsDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetGoodsDetailOk() (*string, bool) {
	if o == nil || IsNil(o.GoodsDetail) {
		return nil, false
	}
	return o.GoodsDetail, true
}

// HasGoodsDetail returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasGoodsDetail() bool {
	if o != nil && !IsNil(o.GoodsDetail) {
		return true
	}

	return false
}

// SetGoodsDetail gets a reference to the given string and assigns it to the GoodsDetail field.
func (o *VoucherAvailableGoodsInfo) SetGoodsDetail(v string) {
	o.GoodsDetail = &v
}

// GetGoodsDetailImages returns the GoodsDetailImages field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetGoodsDetailImages() []string {
	if o == nil || IsNil(o.GoodsDetailImages) {
		var ret []string
		return ret
	}
	return o.GoodsDetailImages
}

// GetGoodsDetailImagesOk returns a tuple with the GoodsDetailImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetGoodsDetailImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.GoodsDetailImages) {
		return nil, false
	}
	return o.GoodsDetailImages, true
}

// HasGoodsDetailImages returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasGoodsDetailImages() bool {
	if o != nil && !IsNil(o.GoodsDetailImages) {
		return true
	}

	return false
}

// SetGoodsDetailImages gets a reference to the given []string and assigns it to the GoodsDetailImages field.
func (o *VoucherAvailableGoodsInfo) SetGoodsDetailImages(v []string) {
	o.GoodsDetailImages = v
}

// GetGoodsDetailRichDescription returns the GoodsDetailRichDescription field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetGoodsDetailRichDescription() string {
	if o == nil || IsNil(o.GoodsDetailRichDescription) {
		var ret string
		return ret
	}
	return *o.GoodsDetailRichDescription
}

// GetGoodsDetailRichDescriptionOk returns a tuple with the GoodsDetailRichDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetGoodsDetailRichDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.GoodsDetailRichDescription) {
		return nil, false
	}
	return o.GoodsDetailRichDescription, true
}

// HasGoodsDetailRichDescription returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasGoodsDetailRichDescription() bool {
	if o != nil && !IsNil(o.GoodsDetailRichDescription) {
		return true
	}

	return false
}

// SetGoodsDetailRichDescription gets a reference to the given string and assigns it to the GoodsDetailRichDescription field.
func (o *VoucherAvailableGoodsInfo) SetGoodsDetailRichDescription(v string) {
	o.GoodsDetailRichDescription = &v
}

// GetGoodsName returns the GoodsName field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetGoodsName() string {
	if o == nil || IsNil(o.GoodsName) {
		var ret string
		return ret
	}
	return *o.GoodsName
}

// GetGoodsNameOk returns a tuple with the GoodsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetGoodsNameOk() (*string, bool) {
	if o == nil || IsNil(o.GoodsName) {
		return nil, false
	}
	return o.GoodsName, true
}

// HasGoodsName returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasGoodsName() bool {
	if o != nil && !IsNil(o.GoodsName) {
		return true
	}

	return false
}

// SetGoodsName gets a reference to the given string and assigns it to the GoodsName field.
func (o *VoucherAvailableGoodsInfo) SetGoodsName(v string) {
	o.GoodsName = &v
}

// GetOriginAmount returns the OriginAmount field value if set, zero value otherwise.
func (o *VoucherAvailableGoodsInfo) GetOriginAmount() string {
	if o == nil || IsNil(o.OriginAmount) {
		var ret string
		return ret
	}
	return *o.OriginAmount
}

// GetOriginAmountOk returns a tuple with the OriginAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGoodsInfo) GetOriginAmountOk() (*string, bool) {
	if o == nil || IsNil(o.OriginAmount) {
		return nil, false
	}
	return o.OriginAmount, true
}

// HasOriginAmount returns a boolean if a field has been set.
func (o *VoucherAvailableGoodsInfo) HasOriginAmount() bool {
	if o != nil && !IsNil(o.OriginAmount) {
		return true
	}

	return false
}

// SetOriginAmount gets a reference to the given string and assigns it to the OriginAmount field.
func (o *VoucherAvailableGoodsInfo) SetOriginAmount(v string) {
	o.OriginAmount = &v
}

func (o VoucherAvailableGoodsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableGoodsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableGoodsSkuIds) {
		toSerialize["available_goods_sku_ids"] = o.AvailableGoodsSkuIds
	}
	if !IsNil(o.ExcludeGoodsSkuIds) {
		toSerialize["exclude_goods_sku_ids"] = o.ExcludeGoodsSkuIds
	}
	if !IsNil(o.GoodsDescription) {
		toSerialize["goods_description"] = o.GoodsDescription
	}
	if !IsNil(o.GoodsDetail) {
		toSerialize["goods_detail"] = o.GoodsDetail
	}
	if !IsNil(o.GoodsDetailImages) {
		toSerialize["goods_detail_images"] = o.GoodsDetailImages
	}
	if !IsNil(o.GoodsDetailRichDescription) {
		toSerialize["goods_detail_rich_description"] = o.GoodsDetailRichDescription
	}
	if !IsNil(o.GoodsName) {
		toSerialize["goods_name"] = o.GoodsName
	}
	if !IsNil(o.OriginAmount) {
		toSerialize["origin_amount"] = o.OriginAmount
	}
	return toSerialize, nil
}

type NullableVoucherAvailableGoodsInfo struct {
	value *VoucherAvailableGoodsInfo
	isSet bool
}

func (v NullableVoucherAvailableGoodsInfo) Get() *VoucherAvailableGoodsInfo {
	return v.value
}

func (v *NullableVoucherAvailableGoodsInfo) Set(val *VoucherAvailableGoodsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableGoodsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableGoodsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableGoodsInfo(val *VoucherAvailableGoodsInfo) *NullableVoucherAvailableGoodsInfo {
	return &NullableVoucherAvailableGoodsInfo{value: val, isSet: true}
}

func (v NullableVoucherAvailableGoodsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableGoodsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

