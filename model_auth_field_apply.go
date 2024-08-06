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

// checks if the AuthFieldApply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthFieldApply{}

// AuthFieldApply struct for AuthFieldApply
type AuthFieldApply struct {
	// 接口英文名称，通过alipay.open.app.api.query接口查询获得。
	ApiName *string `json:"api_name,omitempty"`
	// 贵公司是否有自己的客服渠道，能及时响应和处理舆情数量是多少？
	CustomerAnswer *string `json:"customer_answer,omitempty"`
	// 字段英文名称，通过alipay.open.app.api.query接口查询获得。
	FieldName *string `json:"field_name,omitempty"`
	// 备注
	Memo *string `json:"memo,omitempty"`
	// 功能code，通过alipay.open.app.api.query接口查询获得
	PackageCode *string `json:"package_code,omitempty"`
	// 接入后一年内预计接口秒级调用量峰值是多少？（最高峰值：1000QPS）
	QpsAnswer *string `json:"qps_answer,omitempty"`
	// 场景code，alipay.open.app.api.scene.query 接口查询获得
	SceneCode *string `json:"scene_code,omitempty"`
	// 当为使用使用模板的小程序申请时可传入所使用的小程序模板id
	TinyAppTemplateId *string `json:"tiny_app_template_id,omitempty"`
}

// NewAuthFieldApply instantiates a new AuthFieldApply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthFieldApply() *AuthFieldApply {
	this := AuthFieldApply{}
	return &this
}

// NewAuthFieldApplyWithDefaults instantiates a new AuthFieldApply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthFieldApplyWithDefaults() *AuthFieldApply {
	this := AuthFieldApply{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *AuthFieldApply) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *AuthFieldApply) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *AuthFieldApply) SetApiName(v string) {
	o.ApiName = &v
}

// GetCustomerAnswer returns the CustomerAnswer field value if set, zero value otherwise.
func (o *AuthFieldApply) GetCustomerAnswer() string {
	if o == nil || IsNil(o.CustomerAnswer) {
		var ret string
		return ret
	}
	return *o.CustomerAnswer
}

// GetCustomerAnswerOk returns a tuple with the CustomerAnswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetCustomerAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerAnswer) {
		return nil, false
	}
	return o.CustomerAnswer, true
}

// HasCustomerAnswer returns a boolean if a field has been set.
func (o *AuthFieldApply) HasCustomerAnswer() bool {
	if o != nil && !IsNil(o.CustomerAnswer) {
		return true
	}

	return false
}

// SetCustomerAnswer gets a reference to the given string and assigns it to the CustomerAnswer field.
func (o *AuthFieldApply) SetCustomerAnswer(v string) {
	o.CustomerAnswer = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *AuthFieldApply) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *AuthFieldApply) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *AuthFieldApply) SetFieldName(v string) {
	o.FieldName = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *AuthFieldApply) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *AuthFieldApply) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *AuthFieldApply) SetMemo(v string) {
	o.Memo = &v
}

// GetPackageCode returns the PackageCode field value if set, zero value otherwise.
func (o *AuthFieldApply) GetPackageCode() string {
	if o == nil || IsNil(o.PackageCode) {
		var ret string
		return ret
	}
	return *o.PackageCode
}

// GetPackageCodeOk returns a tuple with the PackageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetPackageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageCode) {
		return nil, false
	}
	return o.PackageCode, true
}

// HasPackageCode returns a boolean if a field has been set.
func (o *AuthFieldApply) HasPackageCode() bool {
	if o != nil && !IsNil(o.PackageCode) {
		return true
	}

	return false
}

// SetPackageCode gets a reference to the given string and assigns it to the PackageCode field.
func (o *AuthFieldApply) SetPackageCode(v string) {
	o.PackageCode = &v
}

// GetQpsAnswer returns the QpsAnswer field value if set, zero value otherwise.
func (o *AuthFieldApply) GetQpsAnswer() string {
	if o == nil || IsNil(o.QpsAnswer) {
		var ret string
		return ret
	}
	return *o.QpsAnswer
}

// GetQpsAnswerOk returns a tuple with the QpsAnswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetQpsAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.QpsAnswer) {
		return nil, false
	}
	return o.QpsAnswer, true
}

// HasQpsAnswer returns a boolean if a field has been set.
func (o *AuthFieldApply) HasQpsAnswer() bool {
	if o != nil && !IsNil(o.QpsAnswer) {
		return true
	}

	return false
}

// SetQpsAnswer gets a reference to the given string and assigns it to the QpsAnswer field.
func (o *AuthFieldApply) SetQpsAnswer(v string) {
	o.QpsAnswer = &v
}

// GetSceneCode returns the SceneCode field value if set, zero value otherwise.
func (o *AuthFieldApply) GetSceneCode() string {
	if o == nil || IsNil(o.SceneCode) {
		var ret string
		return ret
	}
	return *o.SceneCode
}

// GetSceneCodeOk returns a tuple with the SceneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetSceneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneCode) {
		return nil, false
	}
	return o.SceneCode, true
}

// HasSceneCode returns a boolean if a field has been set.
func (o *AuthFieldApply) HasSceneCode() bool {
	if o != nil && !IsNil(o.SceneCode) {
		return true
	}

	return false
}

// SetSceneCode gets a reference to the given string and assigns it to the SceneCode field.
func (o *AuthFieldApply) SetSceneCode(v string) {
	o.SceneCode = &v
}

// GetTinyAppTemplateId returns the TinyAppTemplateId field value if set, zero value otherwise.
func (o *AuthFieldApply) GetTinyAppTemplateId() string {
	if o == nil || IsNil(o.TinyAppTemplateId) {
		var ret string
		return ret
	}
	return *o.TinyAppTemplateId
}

// GetTinyAppTemplateIdOk returns a tuple with the TinyAppTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFieldApply) GetTinyAppTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TinyAppTemplateId) {
		return nil, false
	}
	return o.TinyAppTemplateId, true
}

// HasTinyAppTemplateId returns a boolean if a field has been set.
func (o *AuthFieldApply) HasTinyAppTemplateId() bool {
	if o != nil && !IsNil(o.TinyAppTemplateId) {
		return true
	}

	return false
}

// SetTinyAppTemplateId gets a reference to the given string and assigns it to the TinyAppTemplateId field.
func (o *AuthFieldApply) SetTinyAppTemplateId(v string) {
	o.TinyAppTemplateId = &v
}

func (o AuthFieldApply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthFieldApply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiName) {
		toSerialize["api_name"] = o.ApiName
	}
	if !IsNil(o.CustomerAnswer) {
		toSerialize["customer_answer"] = o.CustomerAnswer
	}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.PackageCode) {
		toSerialize["package_code"] = o.PackageCode
	}
	if !IsNil(o.QpsAnswer) {
		toSerialize["qps_answer"] = o.QpsAnswer
	}
	if !IsNil(o.SceneCode) {
		toSerialize["scene_code"] = o.SceneCode
	}
	if !IsNil(o.TinyAppTemplateId) {
		toSerialize["tiny_app_template_id"] = o.TinyAppTemplateId
	}
	return toSerialize, nil
}

type NullableAuthFieldApply struct {
	value *AuthFieldApply
	isSet bool
}

func (v NullableAuthFieldApply) Get() *AuthFieldApply {
	return v.value
}

func (v *NullableAuthFieldApply) Set(val *AuthFieldApply) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthFieldApply) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthFieldApply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthFieldApply(val *AuthFieldApply) *NullableAuthFieldApply {
	return &NullableAuthFieldApply{value: val, isSet: true}
}

func (v NullableAuthFieldApply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthFieldApply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

