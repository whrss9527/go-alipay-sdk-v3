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

// checks if the AlipayEbppPdeductSignCancelModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppPdeductSignCancelModel{}

// AlipayEbppPdeductSignCancelModel struct for AlipayEbppPdeductSignCancelModel
type AlipayEbppPdeductSignCancelModel struct {
	// 此值只是供代扣中心做最后的渠道统计用，并不做值是什么的强校验，只要不为空就可以
	AgentChannel *string `json:"agent_channel,omitempty"`
	// 标识发起方的ID，从服务窗发起则为appId的值，appId即开放平台分配给接入ISV的id，此处也可以随便真其它值，只要能标识唯一即可
	AgentCode *string `json:"agent_code,omitempty"`
	// 支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// 用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// 需要用户首先处于登陆态，然后访问https://ebppprod.alipay.com/deduct/enterMobileicPayPassword.htm?cb=自己指定的回跳连接地址,访问页面后会进到独立密码校验页，用户输入密码校验成功后，会生成token回调到指定的回跳地址，如果设置cb=www.baidu.com则最后回调的内容是www.baidu.com?token=312314ADFDSFAS,然后签约时直接取得地址后token的值即可
	PayPasswordToken *string `json:"pay_password_token,omitempty"`
	// 用户ID
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayEbppPdeductSignCancelModel instantiates a new AlipayEbppPdeductSignCancelModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppPdeductSignCancelModel() *AlipayEbppPdeductSignCancelModel {
	this := AlipayEbppPdeductSignCancelModel{}
	return &this
}

// NewAlipayEbppPdeductSignCancelModelWithDefaults instantiates a new AlipayEbppPdeductSignCancelModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppPdeductSignCancelModelWithDefaults() *AlipayEbppPdeductSignCancelModel {
	this := AlipayEbppPdeductSignCancelModel{}
	return &this
}

// GetAgentChannel returns the AgentChannel field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelModel) GetAgentChannel() string {
	if o == nil || IsNil(o.AgentChannel) {
		var ret string
		return ret
	}
	return *o.AgentChannel
}

// GetAgentChannelOk returns a tuple with the AgentChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelModel) GetAgentChannelOk() (*string, bool) {
	if o == nil || IsNil(o.AgentChannel) {
		return nil, false
	}
	return o.AgentChannel, true
}

// HasAgentChannel returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelModel) HasAgentChannel() bool {
	if o != nil && !IsNil(o.AgentChannel) {
		return true
	}

	return false
}

// SetAgentChannel gets a reference to the given string and assigns it to the AgentChannel field.
func (o *AlipayEbppPdeductSignCancelModel) SetAgentChannel(v string) {
	o.AgentChannel = &v
}

// GetAgentCode returns the AgentCode field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelModel) GetAgentCode() string {
	if o == nil || IsNil(o.AgentCode) {
		var ret string
		return ret
	}
	return *o.AgentCode
}

// GetAgentCodeOk returns a tuple with the AgentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelModel) GetAgentCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AgentCode) {
		return nil, false
	}
	return o.AgentCode, true
}

// HasAgentCode returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelModel) HasAgentCode() bool {
	if o != nil && !IsNil(o.AgentCode) {
		return true
	}

	return false
}

// SetAgentCode gets a reference to the given string and assigns it to the AgentCode field.
func (o *AlipayEbppPdeductSignCancelModel) SetAgentCode(v string) {
	o.AgentCode = &v
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayEbppPdeductSignCancelModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayEbppPdeductSignCancelModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetPayPasswordToken returns the PayPasswordToken field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelModel) GetPayPasswordToken() string {
	if o == nil || IsNil(o.PayPasswordToken) {
		var ret string
		return ret
	}
	return *o.PayPasswordToken
}

// GetPayPasswordTokenOk returns a tuple with the PayPasswordToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelModel) GetPayPasswordTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PayPasswordToken) {
		return nil, false
	}
	return o.PayPasswordToken, true
}

// HasPayPasswordToken returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelModel) HasPayPasswordToken() bool {
	if o != nil && !IsNil(o.PayPasswordToken) {
		return true
	}

	return false
}

// SetPayPasswordToken gets a reference to the given string and assigns it to the PayPasswordToken field.
func (o *AlipayEbppPdeductSignCancelModel) SetPayPasswordToken(v string) {
	o.PayPasswordToken = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayEbppPdeductSignCancelModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayEbppPdeductSignCancelModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppPdeductSignCancelModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentChannel) {
		toSerialize["agent_channel"] = o.AgentChannel
	}
	if !IsNil(o.AgentCode) {
		toSerialize["agent_code"] = o.AgentCode
	}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.PayPasswordToken) {
		toSerialize["pay_password_token"] = o.PayPasswordToken
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayEbppPdeductSignCancelModel struct {
	value *AlipayEbppPdeductSignCancelModel
	isSet bool
}

func (v NullableAlipayEbppPdeductSignCancelModel) Get() *AlipayEbppPdeductSignCancelModel {
	return v.value
}

func (v *NullableAlipayEbppPdeductSignCancelModel) Set(val *AlipayEbppPdeductSignCancelModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppPdeductSignCancelModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppPdeductSignCancelModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppPdeductSignCancelModel(val *AlipayEbppPdeductSignCancelModel) *NullableAlipayEbppPdeductSignCancelModel {
	return &NullableAlipayEbppPdeductSignCancelModel{value: val, isSet: true}
}

func (v NullableAlipayEbppPdeductSignCancelModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppPdeductSignCancelModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
