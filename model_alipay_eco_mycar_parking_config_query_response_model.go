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

// checks if the AlipayEcoMycarParkingConfigQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoMycarParkingConfigQueryResponseModel{}

// AlipayEcoMycarParkingConfigQueryResponseModel struct for AlipayEcoMycarParkingConfigQueryResponseModel
type AlipayEcoMycarParkingConfigQueryResponseModel struct {
	// 签约支付宝账号，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)传入的参数值
	AccountNo         *string            `json:"account_no,omitempty"`
	InterfaceInfoList *InterfaceInfoList `json:"interface_info_list,omitempty"`
	// 商户在停车平台首页露出的LOGO链接地址，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)调用自动生成该链接
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// 商户简称，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)传入的参数值
	MerchantName *string `json:"merchant_name,omitempty"`
	// 商户客服电话，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)传入的参数值
	MerchantServicePhone *string `json:"merchant_service_phone,omitempty"`
}

// NewAlipayEcoMycarParkingConfigQueryResponseModel instantiates a new AlipayEcoMycarParkingConfigQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoMycarParkingConfigQueryResponseModel() *AlipayEcoMycarParkingConfigQueryResponseModel {
	this := AlipayEcoMycarParkingConfigQueryResponseModel{}
	return &this
}

// NewAlipayEcoMycarParkingConfigQueryResponseModelWithDefaults instantiates a new AlipayEcoMycarParkingConfigQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoMycarParkingConfigQueryResponseModelWithDefaults() *AlipayEcoMycarParkingConfigQueryResponseModel {
	this := AlipayEcoMycarParkingConfigQueryResponseModel{}
	return &this
}

// GetAccountNo returns the AccountNo field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetAccountNo() string {
	if o == nil || IsNil(o.AccountNo) {
		var ret string
		return ret
	}
	return *o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetAccountNoOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNo) {
		return nil, false
	}
	return o.AccountNo, true
}

// HasAccountNo returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) HasAccountNo() bool {
	if o != nil && !IsNil(o.AccountNo) {
		return true
	}

	return false
}

// SetAccountNo gets a reference to the given string and assigns it to the AccountNo field.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) SetAccountNo(v string) {
	o.AccountNo = &v
}

// GetInterfaceInfoList returns the InterfaceInfoList field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetInterfaceInfoList() InterfaceInfoList {
	if o == nil || IsNil(o.InterfaceInfoList) {
		var ret InterfaceInfoList
		return ret
	}
	return *o.InterfaceInfoList
}

// GetInterfaceInfoListOk returns a tuple with the InterfaceInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetInterfaceInfoListOk() (*InterfaceInfoList, bool) {
	if o == nil || IsNil(o.InterfaceInfoList) {
		return nil, false
	}
	return o.InterfaceInfoList, true
}

// HasInterfaceInfoList returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) HasInterfaceInfoList() bool {
	if o != nil && !IsNil(o.InterfaceInfoList) {
		return true
	}

	return false
}

// SetInterfaceInfoList gets a reference to the given InterfaceInfoList and assigns it to the InterfaceInfoList field.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) SetInterfaceInfoList(v InterfaceInfoList) {
	o.InterfaceInfoList = &v
}

// GetMerchantLogo returns the MerchantLogo field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetMerchantLogo() string {
	if o == nil || IsNil(o.MerchantLogo) {
		var ret string
		return ret
	}
	return *o.MerchantLogo
}

// GetMerchantLogoOk returns a tuple with the MerchantLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetMerchantLogoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantLogo) {
		return nil, false
	}
	return o.MerchantLogo, true
}

// HasMerchantLogo returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) HasMerchantLogo() bool {
	if o != nil && !IsNil(o.MerchantLogo) {
		return true
	}

	return false
}

// SetMerchantLogo gets a reference to the given string and assigns it to the MerchantLogo field.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) SetMerchantLogo(v string) {
	o.MerchantLogo = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetMerchantServicePhone returns the MerchantServicePhone field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetMerchantServicePhone() string {
	if o == nil || IsNil(o.MerchantServicePhone) {
		var ret string
		return ret
	}
	return *o.MerchantServicePhone
}

// GetMerchantServicePhoneOk returns a tuple with the MerchantServicePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) GetMerchantServicePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantServicePhone) {
		return nil, false
	}
	return o.MerchantServicePhone, true
}

// HasMerchantServicePhone returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) HasMerchantServicePhone() bool {
	if o != nil && !IsNil(o.MerchantServicePhone) {
		return true
	}

	return false
}

// SetMerchantServicePhone gets a reference to the given string and assigns it to the MerchantServicePhone field.
func (o *AlipayEcoMycarParkingConfigQueryResponseModel) SetMerchantServicePhone(v string) {
	o.MerchantServicePhone = &v
}

func (o AlipayEcoMycarParkingConfigQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoMycarParkingConfigQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNo) {
		toSerialize["account_no"] = o.AccountNo
	}
	if !IsNil(o.InterfaceInfoList) {
		toSerialize["interface_info_list"] = o.InterfaceInfoList
	}
	if !IsNil(o.MerchantLogo) {
		toSerialize["merchant_logo"] = o.MerchantLogo
	}
	if !IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if !IsNil(o.MerchantServicePhone) {
		toSerialize["merchant_service_phone"] = o.MerchantServicePhone
	}
	return toSerialize, nil
}

type NullableAlipayEcoMycarParkingConfigQueryResponseModel struct {
	value *AlipayEcoMycarParkingConfigQueryResponseModel
	isSet bool
}

func (v NullableAlipayEcoMycarParkingConfigQueryResponseModel) Get() *AlipayEcoMycarParkingConfigQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingConfigQueryResponseModel) Set(val *AlipayEcoMycarParkingConfigQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingConfigQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingConfigQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingConfigQueryResponseModel(val *AlipayEcoMycarParkingConfigQueryResponseModel) *NullableAlipayEcoMycarParkingConfigQueryResponseModel {
	return &NullableAlipayEcoMycarParkingConfigQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingConfigQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingConfigQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
