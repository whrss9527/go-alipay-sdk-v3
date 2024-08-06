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

// checks if the AlipayOpenAgentOfflinepaymentSignModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAgentOfflinepaymentSignModel{}

// AlipayOpenAgentOfflinepaymentSignModel struct for AlipayOpenAgentOfflinepaymentSignModel
type AlipayOpenAgentOfflinepaymentSignModel struct {
	// 代商户操作事务编号，通过alipay.open.agent.create接口进行创建。
	BatchNo *string `json:"batch_no,omitempty"`
	// 被邀请授权的营业执照法人手机号码，上传非同人营业执照时必填
	BusinessLicenseMobile *string `json:"business_license_mobile,omitempty"`
	// 营业执照号码。若填写，请与以下营业执照图片、期限、一起提供。
	BusinessLicenseNo *string `json:"business_license_no,omitempty"`
	// 营业期限
	DateLimitation *string `json:"date_limitation,omitempty"`
	DeliveryAddress *SignAddressInfo `json:"delivery_address,omitempty"`
	// 营业期限是否长期有效
	LongTerm *bool `json:"long_term,omitempty"`
	// 所属MCCCode，可参考 <a href=\"https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE\">商家经营类目</a> 中的“经营类目编码”
	MccCode *string `json:"mcc_code,omitempty"`
	// 服务费率（%），0.38~3之间，精确到0.01
	Rate *string `json:"rate,omitempty"`
	ShopAddress *SignAddressInfo `json:"shop_address,omitempty"`
	// 店铺名称
	ShopName *string `json:"shop_name,omitempty"`
}

// NewAlipayOpenAgentOfflinepaymentSignModel instantiates a new AlipayOpenAgentOfflinepaymentSignModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAgentOfflinepaymentSignModel() *AlipayOpenAgentOfflinepaymentSignModel {
	this := AlipayOpenAgentOfflinepaymentSignModel{}
	return &this
}

// NewAlipayOpenAgentOfflinepaymentSignModelWithDefaults instantiates a new AlipayOpenAgentOfflinepaymentSignModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAgentOfflinepaymentSignModelWithDefaults() *AlipayOpenAgentOfflinepaymentSignModel {
	this := AlipayOpenAgentOfflinepaymentSignModel{}
	return &this
}

// GetBatchNo returns the BatchNo field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetBatchNo() string {
	if o == nil || IsNil(o.BatchNo) {
		var ret string
		return ret
	}
	return *o.BatchNo
}

// GetBatchNoOk returns a tuple with the BatchNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetBatchNoOk() (*string, bool) {
	if o == nil || IsNil(o.BatchNo) {
		return nil, false
	}
	return o.BatchNo, true
}

// HasBatchNo returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasBatchNo() bool {
	if o != nil && !IsNil(o.BatchNo) {
		return true
	}

	return false
}

// SetBatchNo gets a reference to the given string and assigns it to the BatchNo field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetBatchNo(v string) {
	o.BatchNo = &v
}

// GetBusinessLicenseMobile returns the BusinessLicenseMobile field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetBusinessLicenseMobile() string {
	if o == nil || IsNil(o.BusinessLicenseMobile) {
		var ret string
		return ret
	}
	return *o.BusinessLicenseMobile
}

// GetBusinessLicenseMobileOk returns a tuple with the BusinessLicenseMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetBusinessLicenseMobileOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLicenseMobile) {
		return nil, false
	}
	return o.BusinessLicenseMobile, true
}

// HasBusinessLicenseMobile returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasBusinessLicenseMobile() bool {
	if o != nil && !IsNil(o.BusinessLicenseMobile) {
		return true
	}

	return false
}

// SetBusinessLicenseMobile gets a reference to the given string and assigns it to the BusinessLicenseMobile field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetBusinessLicenseMobile(v string) {
	o.BusinessLicenseMobile = &v
}

// GetBusinessLicenseNo returns the BusinessLicenseNo field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetBusinessLicenseNo() string {
	if o == nil || IsNil(o.BusinessLicenseNo) {
		var ret string
		return ret
	}
	return *o.BusinessLicenseNo
}

// GetBusinessLicenseNoOk returns a tuple with the BusinessLicenseNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetBusinessLicenseNoOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLicenseNo) {
		return nil, false
	}
	return o.BusinessLicenseNo, true
}

// HasBusinessLicenseNo returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasBusinessLicenseNo() bool {
	if o != nil && !IsNil(o.BusinessLicenseNo) {
		return true
	}

	return false
}

// SetBusinessLicenseNo gets a reference to the given string and assigns it to the BusinessLicenseNo field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetBusinessLicenseNo(v string) {
	o.BusinessLicenseNo = &v
}

// GetDateLimitation returns the DateLimitation field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetDateLimitation() string {
	if o == nil || IsNil(o.DateLimitation) {
		var ret string
		return ret
	}
	return *o.DateLimitation
}

// GetDateLimitationOk returns a tuple with the DateLimitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetDateLimitationOk() (*string, bool) {
	if o == nil || IsNil(o.DateLimitation) {
		return nil, false
	}
	return o.DateLimitation, true
}

// HasDateLimitation returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasDateLimitation() bool {
	if o != nil && !IsNil(o.DateLimitation) {
		return true
	}

	return false
}

// SetDateLimitation gets a reference to the given string and assigns it to the DateLimitation field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetDateLimitation(v string) {
	o.DateLimitation = &v
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetDeliveryAddress() SignAddressInfo {
	if o == nil || IsNil(o.DeliveryAddress) {
		var ret SignAddressInfo
		return ret
	}
	return *o.DeliveryAddress
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetDeliveryAddressOk() (*SignAddressInfo, bool) {
	if o == nil || IsNil(o.DeliveryAddress) {
		return nil, false
	}
	return o.DeliveryAddress, true
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasDeliveryAddress() bool {
	if o != nil && !IsNil(o.DeliveryAddress) {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given SignAddressInfo and assigns it to the DeliveryAddress field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetDeliveryAddress(v SignAddressInfo) {
	o.DeliveryAddress = &v
}

// GetLongTerm returns the LongTerm field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetLongTerm() bool {
	if o == nil || IsNil(o.LongTerm) {
		var ret bool
		return ret
	}
	return *o.LongTerm
}

// GetLongTermOk returns a tuple with the LongTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetLongTermOk() (*bool, bool) {
	if o == nil || IsNil(o.LongTerm) {
		return nil, false
	}
	return o.LongTerm, true
}

// HasLongTerm returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasLongTerm() bool {
	if o != nil && !IsNil(o.LongTerm) {
		return true
	}

	return false
}

// SetLongTerm gets a reference to the given bool and assigns it to the LongTerm field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetLongTerm(v bool) {
	o.LongTerm = &v
}

// GetMccCode returns the MccCode field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetMccCode() string {
	if o == nil || IsNil(o.MccCode) {
		var ret string
		return ret
	}
	return *o.MccCode
}

// GetMccCodeOk returns a tuple with the MccCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetMccCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MccCode) {
		return nil, false
	}
	return o.MccCode, true
}

// HasMccCode returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasMccCode() bool {
	if o != nil && !IsNil(o.MccCode) {
		return true
	}

	return false
}

// SetMccCode gets a reference to the given string and assigns it to the MccCode field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetMccCode(v string) {
	o.MccCode = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetRate(v string) {
	o.Rate = &v
}

// GetShopAddress returns the ShopAddress field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetShopAddress() SignAddressInfo {
	if o == nil || IsNil(o.ShopAddress) {
		var ret SignAddressInfo
		return ret
	}
	return *o.ShopAddress
}

// GetShopAddressOk returns a tuple with the ShopAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetShopAddressOk() (*SignAddressInfo, bool) {
	if o == nil || IsNil(o.ShopAddress) {
		return nil, false
	}
	return o.ShopAddress, true
}

// HasShopAddress returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasShopAddress() bool {
	if o != nil && !IsNil(o.ShopAddress) {
		return true
	}

	return false
}

// SetShopAddress gets a reference to the given SignAddressInfo and assigns it to the ShopAddress field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetShopAddress(v SignAddressInfo) {
	o.ShopAddress = &v
}

// GetShopName returns the ShopName field value if set, zero value otherwise.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetShopName() string {
	if o == nil || IsNil(o.ShopName) {
		var ret string
		return ret
	}
	return *o.ShopName
}

// GetShopNameOk returns a tuple with the ShopName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) GetShopNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShopName) {
		return nil, false
	}
	return o.ShopName, true
}

// HasShopName returns a boolean if a field has been set.
func (o *AlipayOpenAgentOfflinepaymentSignModel) HasShopName() bool {
	if o != nil && !IsNil(o.ShopName) {
		return true
	}

	return false
}

// SetShopName gets a reference to the given string and assigns it to the ShopName field.
func (o *AlipayOpenAgentOfflinepaymentSignModel) SetShopName(v string) {
	o.ShopName = &v
}

func (o AlipayOpenAgentOfflinepaymentSignModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAgentOfflinepaymentSignModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatchNo) {
		toSerialize["batch_no"] = o.BatchNo
	}
	if !IsNil(o.BusinessLicenseMobile) {
		toSerialize["business_license_mobile"] = o.BusinessLicenseMobile
	}
	if !IsNil(o.BusinessLicenseNo) {
		toSerialize["business_license_no"] = o.BusinessLicenseNo
	}
	if !IsNil(o.DateLimitation) {
		toSerialize["date_limitation"] = o.DateLimitation
	}
	if !IsNil(o.DeliveryAddress) {
		toSerialize["delivery_address"] = o.DeliveryAddress
	}
	if !IsNil(o.LongTerm) {
		toSerialize["long_term"] = o.LongTerm
	}
	if !IsNil(o.MccCode) {
		toSerialize["mcc_code"] = o.MccCode
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.ShopAddress) {
		toSerialize["shop_address"] = o.ShopAddress
	}
	if !IsNil(o.ShopName) {
		toSerialize["shop_name"] = o.ShopName
	}
	return toSerialize, nil
}

type NullableAlipayOpenAgentOfflinepaymentSignModel struct {
	value *AlipayOpenAgentOfflinepaymentSignModel
	isSet bool
}

func (v NullableAlipayOpenAgentOfflinepaymentSignModel) Get() *AlipayOpenAgentOfflinepaymentSignModel {
	return v.value
}

func (v *NullableAlipayOpenAgentOfflinepaymentSignModel) Set(val *AlipayOpenAgentOfflinepaymentSignModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentOfflinepaymentSignModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentOfflinepaymentSignModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentOfflinepaymentSignModel(val *AlipayOpenAgentOfflinepaymentSignModel) *NullableAlipayOpenAgentOfflinepaymentSignModel {
	return &NullableAlipayOpenAgentOfflinepaymentSignModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentOfflinepaymentSignModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentOfflinepaymentSignModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

