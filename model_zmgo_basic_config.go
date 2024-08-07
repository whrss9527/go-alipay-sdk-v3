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

// checks if the ZMGOBasicConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZMGOBasicConfig{}

// ZMGOBasicConfig struct for ZMGOBasicConfig
type ZMGOBasicConfig struct {
	// 商户在芝麻GO配置的业务身份编码
	BizType *string `json:"biz_type,omitempty"`
	// 商家客服电话
	Contact *string `json:"contact,omitempty"`
	// 运营商商户支付宝ID。若非ISV代理模式，也就是商户自运营模式，此属性取值与partner_id一致。
	IsvPid *string `json:"isv_pid,omitempty"`
	// 商户LOGO
	MerchantCustomLogo *string `json:"merchant_custom_logo,omitempty"`
	// 外部业务单号，供幂等使用，需保证每次请求的值都不同
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 商户的支付宝ID，即为此商户创建芝麻GO模板
	PartnerId *string `json:"partner_id,omitempty"`
	// 芝麻GO模板名称
	TemplateName *string `json:"template_name,omitempty"`
	// 模板编号
	TemplateNo *string `json:"template_no,omitempty"`
}

// NewZMGOBasicConfig instantiates a new ZMGOBasicConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZMGOBasicConfig() *ZMGOBasicConfig {
	this := ZMGOBasicConfig{}
	return &this
}

// NewZMGOBasicConfigWithDefaults instantiates a new ZMGOBasicConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZMGOBasicConfigWithDefaults() *ZMGOBasicConfig {
	this := ZMGOBasicConfig{}
	return &this
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *ZMGOBasicConfig) SetBizType(v string) {
	o.BizType = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *ZMGOBasicConfig) SetContact(v string) {
	o.Contact = &v
}

// GetIsvPid returns the IsvPid field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetIsvPid() string {
	if o == nil || IsNil(o.IsvPid) {
		var ret string
		return ret
	}
	return *o.IsvPid
}

// GetIsvPidOk returns a tuple with the IsvPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetIsvPidOk() (*string, bool) {
	if o == nil || IsNil(o.IsvPid) {
		return nil, false
	}
	return o.IsvPid, true
}

// HasIsvPid returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasIsvPid() bool {
	if o != nil && !IsNil(o.IsvPid) {
		return true
	}

	return false
}

// SetIsvPid gets a reference to the given string and assigns it to the IsvPid field.
func (o *ZMGOBasicConfig) SetIsvPid(v string) {
	o.IsvPid = &v
}

// GetMerchantCustomLogo returns the MerchantCustomLogo field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetMerchantCustomLogo() string {
	if o == nil || IsNil(o.MerchantCustomLogo) {
		var ret string
		return ret
	}
	return *o.MerchantCustomLogo
}

// GetMerchantCustomLogoOk returns a tuple with the MerchantCustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetMerchantCustomLogoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantCustomLogo) {
		return nil, false
	}
	return o.MerchantCustomLogo, true
}

// HasMerchantCustomLogo returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasMerchantCustomLogo() bool {
	if o != nil && !IsNil(o.MerchantCustomLogo) {
		return true
	}

	return false
}

// SetMerchantCustomLogo gets a reference to the given string and assigns it to the MerchantCustomLogo field.
func (o *ZMGOBasicConfig) SetMerchantCustomLogo(v string) {
	o.MerchantCustomLogo = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *ZMGOBasicConfig) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *ZMGOBasicConfig) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *ZMGOBasicConfig) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTemplateNo returns the TemplateNo field value if set, zero value otherwise.
func (o *ZMGOBasicConfig) GetTemplateNo() string {
	if o == nil || IsNil(o.TemplateNo) {
		var ret string
		return ret
	}
	return *o.TemplateNo
}

// GetTemplateNoOk returns a tuple with the TemplateNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOBasicConfig) GetTemplateNoOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateNo) {
		return nil, false
	}
	return o.TemplateNo, true
}

// HasTemplateNo returns a boolean if a field has been set.
func (o *ZMGOBasicConfig) HasTemplateNo() bool {
	if o != nil && !IsNil(o.TemplateNo) {
		return true
	}

	return false
}

// SetTemplateNo gets a reference to the given string and assigns it to the TemplateNo field.
func (o *ZMGOBasicConfig) SetTemplateNo(v string) {
	o.TemplateNo = &v
}

func (o ZMGOBasicConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZMGOBasicConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.IsvPid) {
		toSerialize["isv_pid"] = o.IsvPid
	}
	if !IsNil(o.MerchantCustomLogo) {
		toSerialize["merchant_custom_logo"] = o.MerchantCustomLogo
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.TemplateName) {
		toSerialize["template_name"] = o.TemplateName
	}
	if !IsNil(o.TemplateNo) {
		toSerialize["template_no"] = o.TemplateNo
	}
	return toSerialize, nil
}

type NullableZMGOBasicConfig struct {
	value *ZMGOBasicConfig
	isSet bool
}

func (v NullableZMGOBasicConfig) Get() *ZMGOBasicConfig {
	return v.value
}

func (v *NullableZMGOBasicConfig) Set(val *ZMGOBasicConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableZMGOBasicConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableZMGOBasicConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZMGOBasicConfig(val *ZMGOBasicConfig) *NullableZMGOBasicConfig {
	return &NullableZMGOBasicConfig{value: val, isSet: true}
}

func (v NullableZMGOBasicConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZMGOBasicConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
