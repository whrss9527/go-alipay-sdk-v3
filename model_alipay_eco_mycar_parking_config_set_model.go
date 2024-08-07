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

// checks if the AlipayEcoMycarParkingConfigSetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoMycarParkingConfigSetModel{}

// AlipayEcoMycarParkingConfigSetModel struct for AlipayEcoMycarParkingConfigSetModel
type AlipayEcoMycarParkingConfigSetModel struct {
	// 签约支付宝账号
	AccountNo *string `json:"account_no,omitempty"`
	// 接口信息列表，停车业务需要配置的接口列表，该值为JSON数据格式的LIST对象，现阶段只需要配置一个页面接口即可 。每次请将所有的接口配置信息都传入，未传的接口信息将会被置空。
	InterfaceInfoList []InterfaceInfoList `json:"interface_info_list,omitempty"`
	// 商户在停车平台首页露出的LOGO。 注意： *该图片为PNG格式，内容为BASE64的字符串，建议图片尺寸27px*27px，大小不要超过60K。 *若为空则停车平台首页将不露出商户LOGO。
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// 商户简称，由开发者提供
	MerchantName *string `json:"merchant_name,omitempty"`
	// 商户客服电话
	MerchantServicePhone *string `json:"merchant_service_phone,omitempty"`
}

// NewAlipayEcoMycarParkingConfigSetModel instantiates a new AlipayEcoMycarParkingConfigSetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoMycarParkingConfigSetModel() *AlipayEcoMycarParkingConfigSetModel {
	this := AlipayEcoMycarParkingConfigSetModel{}
	return &this
}

// NewAlipayEcoMycarParkingConfigSetModelWithDefaults instantiates a new AlipayEcoMycarParkingConfigSetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoMycarParkingConfigSetModelWithDefaults() *AlipayEcoMycarParkingConfigSetModel {
	this := AlipayEcoMycarParkingConfigSetModel{}
	return &this
}

// GetAccountNo returns the AccountNo field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigSetModel) GetAccountNo() string {
	if o == nil || IsNil(o.AccountNo) {
		var ret string
		return ret
	}
	return *o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) GetAccountNoOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNo) {
		return nil, false
	}
	return o.AccountNo, true
}

// HasAccountNo returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) HasAccountNo() bool {
	if o != nil && !IsNil(o.AccountNo) {
		return true
	}

	return false
}

// SetAccountNo gets a reference to the given string and assigns it to the AccountNo field.
func (o *AlipayEcoMycarParkingConfigSetModel) SetAccountNo(v string) {
	o.AccountNo = &v
}

// GetInterfaceInfoList returns the InterfaceInfoList field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigSetModel) GetInterfaceInfoList() []InterfaceInfoList {
	if o == nil || IsNil(o.InterfaceInfoList) {
		var ret []InterfaceInfoList
		return ret
	}
	return o.InterfaceInfoList
}

// GetInterfaceInfoListOk returns a tuple with the InterfaceInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) GetInterfaceInfoListOk() ([]InterfaceInfoList, bool) {
	if o == nil || IsNil(o.InterfaceInfoList) {
		return nil, false
	}
	return o.InterfaceInfoList, true
}

// HasInterfaceInfoList returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) HasInterfaceInfoList() bool {
	if o != nil && !IsNil(o.InterfaceInfoList) {
		return true
	}

	return false
}

// SetInterfaceInfoList gets a reference to the given []InterfaceInfoList and assigns it to the InterfaceInfoList field.
func (o *AlipayEcoMycarParkingConfigSetModel) SetInterfaceInfoList(v []InterfaceInfoList) {
	o.InterfaceInfoList = v
}

// GetMerchantLogo returns the MerchantLogo field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigSetModel) GetMerchantLogo() string {
	if o == nil || IsNil(o.MerchantLogo) {
		var ret string
		return ret
	}
	return *o.MerchantLogo
}

// GetMerchantLogoOk returns a tuple with the MerchantLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) GetMerchantLogoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantLogo) {
		return nil, false
	}
	return o.MerchantLogo, true
}

// HasMerchantLogo returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) HasMerchantLogo() bool {
	if o != nil && !IsNil(o.MerchantLogo) {
		return true
	}

	return false
}

// SetMerchantLogo gets a reference to the given string and assigns it to the MerchantLogo field.
func (o *AlipayEcoMycarParkingConfigSetModel) SetMerchantLogo(v string) {
	o.MerchantLogo = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigSetModel) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *AlipayEcoMycarParkingConfigSetModel) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetMerchantServicePhone returns the MerchantServicePhone field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingConfigSetModel) GetMerchantServicePhone() string {
	if o == nil || IsNil(o.MerchantServicePhone) {
		var ret string
		return ret
	}
	return *o.MerchantServicePhone
}

// GetMerchantServicePhoneOk returns a tuple with the MerchantServicePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) GetMerchantServicePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantServicePhone) {
		return nil, false
	}
	return o.MerchantServicePhone, true
}

// HasMerchantServicePhone returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingConfigSetModel) HasMerchantServicePhone() bool {
	if o != nil && !IsNil(o.MerchantServicePhone) {
		return true
	}

	return false
}

// SetMerchantServicePhone gets a reference to the given string and assigns it to the MerchantServicePhone field.
func (o *AlipayEcoMycarParkingConfigSetModel) SetMerchantServicePhone(v string) {
	o.MerchantServicePhone = &v
}

func (o AlipayEcoMycarParkingConfigSetModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoMycarParkingConfigSetModel) ToMap() (map[string]interface{}, error) {
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

type NullableAlipayEcoMycarParkingConfigSetModel struct {
	value *AlipayEcoMycarParkingConfigSetModel
	isSet bool
}

func (v NullableAlipayEcoMycarParkingConfigSetModel) Get() *AlipayEcoMycarParkingConfigSetModel {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingConfigSetModel) Set(val *AlipayEcoMycarParkingConfigSetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingConfigSetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingConfigSetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingConfigSetModel(val *AlipayEcoMycarParkingConfigSetModel) *NullableAlipayEcoMycarParkingConfigSetModel {
	return &NullableAlipayEcoMycarParkingConfigSetModel{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingConfigSetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingConfigSetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
