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

// checks if the SignFieldBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignFieldBean{}

// SignFieldBean struct for SignFieldBean
type SignFieldBean struct {
	// 签署类型：  platform-平台自动签：无需指定签署人。创建流程后，系统将自动盖上商户的默认企业印章。  person-个人签署：需要指定签署人个人信息。创建流程后，需通过签署插件完成签署。 org-企业签署：需要指定签署企业信息与经办人个人信息。创建流程后，需经办人代企业完成签署。（企业暂不支持插件签署，可通过获取签署地址接口获取企业签署地址）
	SignFieldType *string     `json:"sign_field_type,omitempty"`
	Signer        *SignerBean `json:"signer,omitempty"`
	// 签署区key值
	StructKey *string `json:"struct_key,omitempty"`
}

// NewSignFieldBean instantiates a new SignFieldBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignFieldBean() *SignFieldBean {
	this := SignFieldBean{}
	return &this
}

// NewSignFieldBeanWithDefaults instantiates a new SignFieldBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignFieldBeanWithDefaults() *SignFieldBean {
	this := SignFieldBean{}
	return &this
}

// GetSignFieldType returns the SignFieldType field value if set, zero value otherwise.
func (o *SignFieldBean) GetSignFieldType() string {
	if o == nil || IsNil(o.SignFieldType) {
		var ret string
		return ret
	}
	return *o.SignFieldType
}

// GetSignFieldTypeOk returns a tuple with the SignFieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignFieldBean) GetSignFieldTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SignFieldType) {
		return nil, false
	}
	return o.SignFieldType, true
}

// HasSignFieldType returns a boolean if a field has been set.
func (o *SignFieldBean) HasSignFieldType() bool {
	if o != nil && !IsNil(o.SignFieldType) {
		return true
	}

	return false
}

// SetSignFieldType gets a reference to the given string and assigns it to the SignFieldType field.
func (o *SignFieldBean) SetSignFieldType(v string) {
	o.SignFieldType = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *SignFieldBean) GetSigner() SignerBean {
	if o == nil || IsNil(o.Signer) {
		var ret SignerBean
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignFieldBean) GetSignerOk() (*SignerBean, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *SignFieldBean) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given SignerBean and assigns it to the Signer field.
func (o *SignFieldBean) SetSigner(v SignerBean) {
	o.Signer = &v
}

// GetStructKey returns the StructKey field value if set, zero value otherwise.
func (o *SignFieldBean) GetStructKey() string {
	if o == nil || IsNil(o.StructKey) {
		var ret string
		return ret
	}
	return *o.StructKey
}

// GetStructKeyOk returns a tuple with the StructKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignFieldBean) GetStructKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StructKey) {
		return nil, false
	}
	return o.StructKey, true
}

// HasStructKey returns a boolean if a field has been set.
func (o *SignFieldBean) HasStructKey() bool {
	if o != nil && !IsNil(o.StructKey) {
		return true
	}

	return false
}

// SetStructKey gets a reference to the given string and assigns it to the StructKey field.
func (o *SignFieldBean) SetStructKey(v string) {
	o.StructKey = &v
}

func (o SignFieldBean) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignFieldBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignFieldType) {
		toSerialize["sign_field_type"] = o.SignFieldType
	}
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.StructKey) {
		toSerialize["struct_key"] = o.StructKey
	}
	return toSerialize, nil
}

type NullableSignFieldBean struct {
	value *SignFieldBean
	isSet bool
}

func (v NullableSignFieldBean) Get() *SignFieldBean {
	return v.value
}

func (v *NullableSignFieldBean) Set(val *SignFieldBean) {
	v.value = val
	v.isSet = true
}

func (v NullableSignFieldBean) IsSet() bool {
	return v.isSet
}

func (v *NullableSignFieldBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignFieldBean(val *SignFieldBean) *NullableSignFieldBean {
	return &NullableSignFieldBean{value: val, isSet: true}
}

func (v NullableSignFieldBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignFieldBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
