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

// checks if the ZhimaMerchantSubsidiariesApplyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaMerchantSubsidiariesApplyModel{}

// ZhimaMerchantSubsidiariesApplyModel struct for ZhimaMerchantSubsidiariesApplyModel
type ZhimaMerchantSubsidiariesApplyModel struct {
	// 商户pid
	Pid *string `json:"pid,omitempty"`
	// 信用服务id
	ServiceId *string `json:"service_id,omitempty"`
	// 二级商户进件id
	Smid *string `json:"smid,omitempty"`
	// 子商户联系电话
	SubMerchantContactNumber *string `json:"sub_merchant_contact_number,omitempty"`
	// 子商户跳转链接
	SubMerchantJumpLink *string `json:"sub_merchant_jump_link,omitempty"`
	// 子商户logo地址
	SubMerchantLogoUrl *string `json:"sub_merchant_logo_url,omitempty"`
	// 子商户名
	SubMerchantName *string `json:"sub_merchant_name,omitempty"`
	// 子商户id
	SubPid *string `json:"sub_pid,omitempty"`
}

// NewZhimaMerchantSubsidiariesApplyModel instantiates a new ZhimaMerchantSubsidiariesApplyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaMerchantSubsidiariesApplyModel() *ZhimaMerchantSubsidiariesApplyModel {
	this := ZhimaMerchantSubsidiariesApplyModel{}
	return &this
}

// NewZhimaMerchantSubsidiariesApplyModelWithDefaults instantiates a new ZhimaMerchantSubsidiariesApplyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaMerchantSubsidiariesApplyModelWithDefaults() *ZhimaMerchantSubsidiariesApplyModel {
	this := ZhimaMerchantSubsidiariesApplyModel{}
	return &this
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetPid(v string) {
	o.Pid = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetSmid returns the Smid field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSmid() string {
	if o == nil || IsNil(o.Smid) {
		var ret string
		return ret
	}
	return *o.Smid
}

// GetSmidOk returns a tuple with the Smid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSmidOk() (*string, bool) {
	if o == nil || IsNil(o.Smid) {
		return nil, false
	}
	return o.Smid, true
}

// HasSmid returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasSmid() bool {
	if o != nil && !IsNil(o.Smid) {
		return true
	}

	return false
}

// SetSmid gets a reference to the given string and assigns it to the Smid field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetSmid(v string) {
	o.Smid = &v
}

// GetSubMerchantContactNumber returns the SubMerchantContactNumber field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantContactNumber() string {
	if o == nil || IsNil(o.SubMerchantContactNumber) {
		var ret string
		return ret
	}
	return *o.SubMerchantContactNumber
}

// GetSubMerchantContactNumberOk returns a tuple with the SubMerchantContactNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantContactNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantContactNumber) {
		return nil, false
	}
	return o.SubMerchantContactNumber, true
}

// HasSubMerchantContactNumber returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasSubMerchantContactNumber() bool {
	if o != nil && !IsNil(o.SubMerchantContactNumber) {
		return true
	}

	return false
}

// SetSubMerchantContactNumber gets a reference to the given string and assigns it to the SubMerchantContactNumber field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetSubMerchantContactNumber(v string) {
	o.SubMerchantContactNumber = &v
}

// GetSubMerchantJumpLink returns the SubMerchantJumpLink field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantJumpLink() string {
	if o == nil || IsNil(o.SubMerchantJumpLink) {
		var ret string
		return ret
	}
	return *o.SubMerchantJumpLink
}

// GetSubMerchantJumpLinkOk returns a tuple with the SubMerchantJumpLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantJumpLinkOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantJumpLink) {
		return nil, false
	}
	return o.SubMerchantJumpLink, true
}

// HasSubMerchantJumpLink returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasSubMerchantJumpLink() bool {
	if o != nil && !IsNil(o.SubMerchantJumpLink) {
		return true
	}

	return false
}

// SetSubMerchantJumpLink gets a reference to the given string and assigns it to the SubMerchantJumpLink field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetSubMerchantJumpLink(v string) {
	o.SubMerchantJumpLink = &v
}

// GetSubMerchantLogoUrl returns the SubMerchantLogoUrl field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantLogoUrl() string {
	if o == nil || IsNil(o.SubMerchantLogoUrl) {
		var ret string
		return ret
	}
	return *o.SubMerchantLogoUrl
}

// GetSubMerchantLogoUrlOk returns a tuple with the SubMerchantLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantLogoUrl) {
		return nil, false
	}
	return o.SubMerchantLogoUrl, true
}

// HasSubMerchantLogoUrl returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasSubMerchantLogoUrl() bool {
	if o != nil && !IsNil(o.SubMerchantLogoUrl) {
		return true
	}

	return false
}

// SetSubMerchantLogoUrl gets a reference to the given string and assigns it to the SubMerchantLogoUrl field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetSubMerchantLogoUrl(v string) {
	o.SubMerchantLogoUrl = &v
}

// GetSubMerchantName returns the SubMerchantName field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantName() string {
	if o == nil || IsNil(o.SubMerchantName) {
		var ret string
		return ret
	}
	return *o.SubMerchantName
}

// GetSubMerchantNameOk returns a tuple with the SubMerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubMerchantName) {
		return nil, false
	}
	return o.SubMerchantName, true
}

// HasSubMerchantName returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasSubMerchantName() bool {
	if o != nil && !IsNil(o.SubMerchantName) {
		return true
	}

	return false
}

// SetSubMerchantName gets a reference to the given string and assigns it to the SubMerchantName field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetSubMerchantName(v string) {
	o.SubMerchantName = &v
}

// GetSubPid returns the SubPid field value if set, zero value otherwise.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubPid() string {
	if o == nil || IsNil(o.SubPid) {
		var ret string
		return ret
	}
	return *o.SubPid
}

// GetSubPidOk returns a tuple with the SubPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) GetSubPidOk() (*string, bool) {
	if o == nil || IsNil(o.SubPid) {
		return nil, false
	}
	return o.SubPid, true
}

// HasSubPid returns a boolean if a field has been set.
func (o *ZhimaMerchantSubsidiariesApplyModel) HasSubPid() bool {
	if o != nil && !IsNil(o.SubPid) {
		return true
	}

	return false
}

// SetSubPid gets a reference to the given string and assigns it to the SubPid field.
func (o *ZhimaMerchantSubsidiariesApplyModel) SetSubPid(v string) {
	o.SubPid = &v
}

func (o ZhimaMerchantSubsidiariesApplyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaMerchantSubsidiariesApplyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !IsNil(o.Smid) {
		toSerialize["smid"] = o.Smid
	}
	if !IsNil(o.SubMerchantContactNumber) {
		toSerialize["sub_merchant_contact_number"] = o.SubMerchantContactNumber
	}
	if !IsNil(o.SubMerchantJumpLink) {
		toSerialize["sub_merchant_jump_link"] = o.SubMerchantJumpLink
	}
	if !IsNil(o.SubMerchantLogoUrl) {
		toSerialize["sub_merchant_logo_url"] = o.SubMerchantLogoUrl
	}
	if !IsNil(o.SubMerchantName) {
		toSerialize["sub_merchant_name"] = o.SubMerchantName
	}
	if !IsNil(o.SubPid) {
		toSerialize["sub_pid"] = o.SubPid
	}
	return toSerialize, nil
}

type NullableZhimaMerchantSubsidiariesApplyModel struct {
	value *ZhimaMerchantSubsidiariesApplyModel
	isSet bool
}

func (v NullableZhimaMerchantSubsidiariesApplyModel) Get() *ZhimaMerchantSubsidiariesApplyModel {
	return v.value
}

func (v *NullableZhimaMerchantSubsidiariesApplyModel) Set(val *ZhimaMerchantSubsidiariesApplyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaMerchantSubsidiariesApplyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaMerchantSubsidiariesApplyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaMerchantSubsidiariesApplyModel(val *ZhimaMerchantSubsidiariesApplyModel) *NullableZhimaMerchantSubsidiariesApplyModel {
	return &NullableZhimaMerchantSubsidiariesApplyModel{value: val, isSet: true}
}

func (v NullableZhimaMerchantSubsidiariesApplyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaMerchantSubsidiariesApplyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


