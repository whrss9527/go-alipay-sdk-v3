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

// checks if the ZolozAuthenticationCustomerFacemanageDeleteModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZolozAuthenticationCustomerFacemanageDeleteModel{}

// ZolozAuthenticationCustomerFacemanageDeleteModel struct for ZolozAuthenticationCustomerFacemanageDeleteModel
type ZolozAuthenticationCustomerFacemanageDeleteModel struct {
	// 地域编码
	Areacode *string `json:"areacode,omitempty"`
	// 人脸产品能力
	BizType *string `json:"biz_type,omitempty"`
	// 业务量
	Bizscale *string `json:"bizscale,omitempty"`
	// 商户品牌
	Brandcode *string `json:"brandcode,omitempty"`
	// 商户机具唯一编码，关键参数
	Devicenum *string `json:"devicenum,omitempty"`
	// 拓展参数
	Extinfo *string `json:"extinfo,omitempty"`
	// 入库类型 IDCARD:身份证 ALIPAY_USER:支付宝用户id, ALIPAY_TEL:手机号入库 CUSTOMER:自定义
	Facetype *string `json:"facetype,omitempty"`
	// 入库用户信息
	Faceval *string `json:"faceval,omitempty"`
	// 分组
	Group *string `json:"group,omitempty"`
	// 门店编码
	Storecode *string `json:"storecode,omitempty"`
	// 有效期天数，如7天、30天、365天
	Validtimes *string `json:"validtimes,omitempty"`
}

// NewZolozAuthenticationCustomerFacemanageDeleteModel instantiates a new ZolozAuthenticationCustomerFacemanageDeleteModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZolozAuthenticationCustomerFacemanageDeleteModel() *ZolozAuthenticationCustomerFacemanageDeleteModel {
	this := ZolozAuthenticationCustomerFacemanageDeleteModel{}
	return &this
}

// NewZolozAuthenticationCustomerFacemanageDeleteModelWithDefaults instantiates a new ZolozAuthenticationCustomerFacemanageDeleteModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZolozAuthenticationCustomerFacemanageDeleteModelWithDefaults() *ZolozAuthenticationCustomerFacemanageDeleteModel {
	this := ZolozAuthenticationCustomerFacemanageDeleteModel{}
	return &this
}

// GetAreacode returns the Areacode field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetAreacode() string {
	if o == nil || IsNil(o.Areacode) {
		var ret string
		return ret
	}
	return *o.Areacode
}

// GetAreacodeOk returns a tuple with the Areacode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetAreacodeOk() (*string, bool) {
	if o == nil || IsNil(o.Areacode) {
		return nil, false
	}
	return o.Areacode, true
}

// HasAreacode returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasAreacode() bool {
	if o != nil && !IsNil(o.Areacode) {
		return true
	}

	return false
}

// SetAreacode gets a reference to the given string and assigns it to the Areacode field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetAreacode(v string) {
	o.Areacode = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetBizType(v string) {
	o.BizType = &v
}

// GetBizscale returns the Bizscale field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetBizscale() string {
	if o == nil || IsNil(o.Bizscale) {
		var ret string
		return ret
	}
	return *o.Bizscale
}

// GetBizscaleOk returns a tuple with the Bizscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetBizscaleOk() (*string, bool) {
	if o == nil || IsNil(o.Bizscale) {
		return nil, false
	}
	return o.Bizscale, true
}

// HasBizscale returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasBizscale() bool {
	if o != nil && !IsNil(o.Bizscale) {
		return true
	}

	return false
}

// SetBizscale gets a reference to the given string and assigns it to the Bizscale field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetBizscale(v string) {
	o.Bizscale = &v
}

// GetBrandcode returns the Brandcode field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetBrandcode() string {
	if o == nil || IsNil(o.Brandcode) {
		var ret string
		return ret
	}
	return *o.Brandcode
}

// GetBrandcodeOk returns a tuple with the Brandcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetBrandcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Brandcode) {
		return nil, false
	}
	return o.Brandcode, true
}

// HasBrandcode returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasBrandcode() bool {
	if o != nil && !IsNil(o.Brandcode) {
		return true
	}

	return false
}

// SetBrandcode gets a reference to the given string and assigns it to the Brandcode field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetBrandcode(v string) {
	o.Brandcode = &v
}

// GetDevicenum returns the Devicenum field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetDevicenum() string {
	if o == nil || IsNil(o.Devicenum) {
		var ret string
		return ret
	}
	return *o.Devicenum
}

// GetDevicenumOk returns a tuple with the Devicenum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetDevicenumOk() (*string, bool) {
	if o == nil || IsNil(o.Devicenum) {
		return nil, false
	}
	return o.Devicenum, true
}

// HasDevicenum returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasDevicenum() bool {
	if o != nil && !IsNil(o.Devicenum) {
		return true
	}

	return false
}

// SetDevicenum gets a reference to the given string and assigns it to the Devicenum field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetDevicenum(v string) {
	o.Devicenum = &v
}

// GetExtinfo returns the Extinfo field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetExtinfo() string {
	if o == nil || IsNil(o.Extinfo) {
		var ret string
		return ret
	}
	return *o.Extinfo
}

// GetExtinfoOk returns a tuple with the Extinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetExtinfoOk() (*string, bool) {
	if o == nil || IsNil(o.Extinfo) {
		return nil, false
	}
	return o.Extinfo, true
}

// HasExtinfo returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasExtinfo() bool {
	if o != nil && !IsNil(o.Extinfo) {
		return true
	}

	return false
}

// SetExtinfo gets a reference to the given string and assigns it to the Extinfo field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetExtinfo(v string) {
	o.Extinfo = &v
}

// GetFacetype returns the Facetype field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetFacetype() string {
	if o == nil || IsNil(o.Facetype) {
		var ret string
		return ret
	}
	return *o.Facetype
}

// GetFacetypeOk returns a tuple with the Facetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetFacetypeOk() (*string, bool) {
	if o == nil || IsNil(o.Facetype) {
		return nil, false
	}
	return o.Facetype, true
}

// HasFacetype returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasFacetype() bool {
	if o != nil && !IsNil(o.Facetype) {
		return true
	}

	return false
}

// SetFacetype gets a reference to the given string and assigns it to the Facetype field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetFacetype(v string) {
	o.Facetype = &v
}

// GetFaceval returns the Faceval field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetFaceval() string {
	if o == nil || IsNil(o.Faceval) {
		var ret string
		return ret
	}
	return *o.Faceval
}

// GetFacevalOk returns a tuple with the Faceval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetFacevalOk() (*string, bool) {
	if o == nil || IsNil(o.Faceval) {
		return nil, false
	}
	return o.Faceval, true
}

// HasFaceval returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasFaceval() bool {
	if o != nil && !IsNil(o.Faceval) {
		return true
	}

	return false
}

// SetFaceval gets a reference to the given string and assigns it to the Faceval field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetFaceval(v string) {
	o.Faceval = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetGroup(v string) {
	o.Group = &v
}

// GetStorecode returns the Storecode field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetStorecode() string {
	if o == nil || IsNil(o.Storecode) {
		var ret string
		return ret
	}
	return *o.Storecode
}

// GetStorecodeOk returns a tuple with the Storecode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetStorecodeOk() (*string, bool) {
	if o == nil || IsNil(o.Storecode) {
		return nil, false
	}
	return o.Storecode, true
}

// HasStorecode returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasStorecode() bool {
	if o != nil && !IsNil(o.Storecode) {
		return true
	}

	return false
}

// SetStorecode gets a reference to the given string and assigns it to the Storecode field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetStorecode(v string) {
	o.Storecode = &v
}

// GetValidtimes returns the Validtimes field value if set, zero value otherwise.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetValidtimes() string {
	if o == nil || IsNil(o.Validtimes) {
		var ret string
		return ret
	}
	return *o.Validtimes
}

// GetValidtimesOk returns a tuple with the Validtimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) GetValidtimesOk() (*string, bool) {
	if o == nil || IsNil(o.Validtimes) {
		return nil, false
	}
	return o.Validtimes, true
}

// HasValidtimes returns a boolean if a field has been set.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) HasValidtimes() bool {
	if o != nil && !IsNil(o.Validtimes) {
		return true
	}

	return false
}

// SetValidtimes gets a reference to the given string and assigns it to the Validtimes field.
func (o *ZolozAuthenticationCustomerFacemanageDeleteModel) SetValidtimes(v string) {
	o.Validtimes = &v
}

func (o ZolozAuthenticationCustomerFacemanageDeleteModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZolozAuthenticationCustomerFacemanageDeleteModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Areacode) {
		toSerialize["areacode"] = o.Areacode
	}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.Bizscale) {
		toSerialize["bizscale"] = o.Bizscale
	}
	if !IsNil(o.Brandcode) {
		toSerialize["brandcode"] = o.Brandcode
	}
	if !IsNil(o.Devicenum) {
		toSerialize["devicenum"] = o.Devicenum
	}
	if !IsNil(o.Extinfo) {
		toSerialize["extinfo"] = o.Extinfo
	}
	if !IsNil(o.Facetype) {
		toSerialize["facetype"] = o.Facetype
	}
	if !IsNil(o.Faceval) {
		toSerialize["faceval"] = o.Faceval
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Storecode) {
		toSerialize["storecode"] = o.Storecode
	}
	if !IsNil(o.Validtimes) {
		toSerialize["validtimes"] = o.Validtimes
	}
	return toSerialize, nil
}

type NullableZolozAuthenticationCustomerFacemanageDeleteModel struct {
	value *ZolozAuthenticationCustomerFacemanageDeleteModel
	isSet bool
}

func (v NullableZolozAuthenticationCustomerFacemanageDeleteModel) Get() *ZolozAuthenticationCustomerFacemanageDeleteModel {
	return v.value
}

func (v *NullableZolozAuthenticationCustomerFacemanageDeleteModel) Set(val *ZolozAuthenticationCustomerFacemanageDeleteModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZolozAuthenticationCustomerFacemanageDeleteModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZolozAuthenticationCustomerFacemanageDeleteModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZolozAuthenticationCustomerFacemanageDeleteModel(val *ZolozAuthenticationCustomerFacemanageDeleteModel) *NullableZolozAuthenticationCustomerFacemanageDeleteModel {
	return &NullableZolozAuthenticationCustomerFacemanageDeleteModel{value: val, isSet: true}
}

func (v NullableZolozAuthenticationCustomerFacemanageDeleteModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZolozAuthenticationCustomerFacemanageDeleteModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
