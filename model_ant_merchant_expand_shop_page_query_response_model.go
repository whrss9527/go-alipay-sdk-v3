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

// checks if the AntMerchantExpandShopPageQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandShopPageQueryResponseModel{}

// AntMerchantExpandShopPageQueryResponseModel struct for AntMerchantExpandShopPageQueryResponseModel
type AntMerchantExpandShopPageQueryResponseModel struct {
	// 门店地址库 ID，按照一定的平台规则识别出的线下真实存在、真实经营的蚂蚁门店地址库 ID，将作用于服务商的返佣激励、商品/券等权益的公域分发。如平台未返回alipay_poiid，请在确认门店信息真实有效后，稍后再进行查询
	AlipayPoiid *string `json:"alipay_poiid,omitempty"`
	// 门店详情
	ShopInfos []ShopQueryOpenApiVO `json:"shop_infos,omitempty"`
	// 总页数
	TotalPages *int32 `json:"total_pages,omitempty"`
}

// NewAntMerchantExpandShopPageQueryResponseModel instantiates a new AntMerchantExpandShopPageQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandShopPageQueryResponseModel() *AntMerchantExpandShopPageQueryResponseModel {
	this := AntMerchantExpandShopPageQueryResponseModel{}
	return &this
}

// NewAntMerchantExpandShopPageQueryResponseModelWithDefaults instantiates a new AntMerchantExpandShopPageQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandShopPageQueryResponseModelWithDefaults() *AntMerchantExpandShopPageQueryResponseModel {
	this := AntMerchantExpandShopPageQueryResponseModel{}
	return &this
}

// GetAlipayPoiid returns the AlipayPoiid field value if set, zero value otherwise.
func (o *AntMerchantExpandShopPageQueryResponseModel) GetAlipayPoiid() string {
	if o == nil || IsNil(o.AlipayPoiid) {
		var ret string
		return ret
	}
	return *o.AlipayPoiid
}

// GetAlipayPoiidOk returns a tuple with the AlipayPoiid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopPageQueryResponseModel) GetAlipayPoiidOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayPoiid) {
		return nil, false
	}
	return o.AlipayPoiid, true
}

// HasAlipayPoiid returns a boolean if a field has been set.
func (o *AntMerchantExpandShopPageQueryResponseModel) HasAlipayPoiid() bool {
	if o != nil && !IsNil(o.AlipayPoiid) {
		return true
	}

	return false
}

// SetAlipayPoiid gets a reference to the given string and assigns it to the AlipayPoiid field.
func (o *AntMerchantExpandShopPageQueryResponseModel) SetAlipayPoiid(v string) {
	o.AlipayPoiid = &v
}

// GetShopInfos returns the ShopInfos field value if set, zero value otherwise.
func (o *AntMerchantExpandShopPageQueryResponseModel) GetShopInfos() []ShopQueryOpenApiVO {
	if o == nil || IsNil(o.ShopInfos) {
		var ret []ShopQueryOpenApiVO
		return ret
	}
	return o.ShopInfos
}

// GetShopInfosOk returns a tuple with the ShopInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopPageQueryResponseModel) GetShopInfosOk() ([]ShopQueryOpenApiVO, bool) {
	if o == nil || IsNil(o.ShopInfos) {
		return nil, false
	}
	return o.ShopInfos, true
}

// HasShopInfos returns a boolean if a field has been set.
func (o *AntMerchantExpandShopPageQueryResponseModel) HasShopInfos() bool {
	if o != nil && !IsNil(o.ShopInfos) {
		return true
	}

	return false
}

// SetShopInfos gets a reference to the given []ShopQueryOpenApiVO and assigns it to the ShopInfos field.
func (o *AntMerchantExpandShopPageQueryResponseModel) SetShopInfos(v []ShopQueryOpenApiVO) {
	o.ShopInfos = v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *AntMerchantExpandShopPageQueryResponseModel) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandShopPageQueryResponseModel) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *AntMerchantExpandShopPageQueryResponseModel) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *AntMerchantExpandShopPageQueryResponseModel) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o AntMerchantExpandShopPageQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandShopPageQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlipayPoiid) {
		toSerialize["alipay_poiid"] = o.AlipayPoiid
	}
	if !IsNil(o.ShopInfos) {
		toSerialize["shop_infos"] = o.ShopInfos
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandShopPageQueryResponseModel struct {
	value *AntMerchantExpandShopPageQueryResponseModel
	isSet bool
}

func (v NullableAntMerchantExpandShopPageQueryResponseModel) Get() *AntMerchantExpandShopPageQueryResponseModel {
	return v.value
}

func (v *NullableAntMerchantExpandShopPageQueryResponseModel) Set(val *AntMerchantExpandShopPageQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopPageQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopPageQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopPageQueryResponseModel(val *AntMerchantExpandShopPageQueryResponseModel) *NullableAntMerchantExpandShopPageQueryResponseModel {
	return &NullableAntMerchantExpandShopPageQueryResponseModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopPageQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopPageQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


