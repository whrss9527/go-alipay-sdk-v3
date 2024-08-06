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

// checks if the AlipayOpenAppServiceApplyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAppServiceApplyModel{}

// AlipayOpenAppServiceApplyModel struct for AlipayOpenAppServiceApplyModel
type AlipayOpenAppServiceApplyModel struct {
	// 行业类目id,获取请参考<a href=\"https://opendocs.alipay.com/mini/03ci0w?pathHash=ed3c875c\">各个行业场景服务接入资料</a>
	CategoryId *string `json:"category_id,omitempty"`
	// 外部业务编号,平台会基于appId+out_biz_no做幂等控制，如果出现幂等，会返回幂等的service_code
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 服务schema版本
	SchemaVersion *string `json:"schema_version,omitempty"`
	// 服务编码, 首次提报时设置空值，支付宝侧返回ServiceCode；当传入serviceCode，则对已提报服务做编辑操作。
	ServiceCode *string `json:"service_code,omitempty"`
	// 服务xml,格式请参考<a href=\"https://opendocs.alipay.com/mini/03cj40?pathHash=a5de4bd9\">Schema 规则介绍</a>
	ServiceXml *string `json:"service_xml,omitempty"`
	// 默认值:DEFAULT。
	TemplateType *string `json:"template_type,omitempty"`
	// 商户提报服务履约类型
	UserServiceDeliveryType *string `json:"user_service_delivery_type,omitempty"`
}

// NewAlipayOpenAppServiceApplyModel instantiates a new AlipayOpenAppServiceApplyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAppServiceApplyModel() *AlipayOpenAppServiceApplyModel {
	this := AlipayOpenAppServiceApplyModel{}
	return &this
}

// NewAlipayOpenAppServiceApplyModelWithDefaults instantiates a new AlipayOpenAppServiceApplyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAppServiceApplyModelWithDefaults() *AlipayOpenAppServiceApplyModel {
	this := AlipayOpenAppServiceApplyModel{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceApplyModel) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceApplyModel) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceApplyModel) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *AlipayOpenAppServiceApplyModel) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceApplyModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceApplyModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceApplyModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayOpenAppServiceApplyModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceApplyModel) GetSchemaVersion() string {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceApplyModel) GetSchemaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceApplyModel) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *AlipayOpenAppServiceApplyModel) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceApplyModel) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceApplyModel) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceApplyModel) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *AlipayOpenAppServiceApplyModel) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetServiceXml returns the ServiceXml field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceApplyModel) GetServiceXml() string {
	if o == nil || IsNil(o.ServiceXml) {
		var ret string
		return ret
	}
	return *o.ServiceXml
}

// GetServiceXmlOk returns a tuple with the ServiceXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceApplyModel) GetServiceXmlOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceXml) {
		return nil, false
	}
	return o.ServiceXml, true
}

// HasServiceXml returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceApplyModel) HasServiceXml() bool {
	if o != nil && !IsNil(o.ServiceXml) {
		return true
	}

	return false
}

// SetServiceXml gets a reference to the given string and assigns it to the ServiceXml field.
func (o *AlipayOpenAppServiceApplyModel) SetServiceXml(v string) {
	o.ServiceXml = &v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceApplyModel) GetTemplateType() string {
	if o == nil || IsNil(o.TemplateType) {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceApplyModel) GetTemplateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateType) {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceApplyModel) HasTemplateType() bool {
	if o != nil && !IsNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *AlipayOpenAppServiceApplyModel) SetTemplateType(v string) {
	o.TemplateType = &v
}

// GetUserServiceDeliveryType returns the UserServiceDeliveryType field value if set, zero value otherwise.
func (o *AlipayOpenAppServiceApplyModel) GetUserServiceDeliveryType() string {
	if o == nil || IsNil(o.UserServiceDeliveryType) {
		var ret string
		return ret
	}
	return *o.UserServiceDeliveryType
}

// GetUserServiceDeliveryTypeOk returns a tuple with the UserServiceDeliveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppServiceApplyModel) GetUserServiceDeliveryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserServiceDeliveryType) {
		return nil, false
	}
	return o.UserServiceDeliveryType, true
}

// HasUserServiceDeliveryType returns a boolean if a field has been set.
func (o *AlipayOpenAppServiceApplyModel) HasUserServiceDeliveryType() bool {
	if o != nil && !IsNil(o.UserServiceDeliveryType) {
		return true
	}

	return false
}

// SetUserServiceDeliveryType gets a reference to the given string and assigns it to the UserServiceDeliveryType field.
func (o *AlipayOpenAppServiceApplyModel) SetUserServiceDeliveryType(v string) {
	o.UserServiceDeliveryType = &v
}

func (o AlipayOpenAppServiceApplyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAppServiceApplyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["service_code"] = o.ServiceCode
	}
	if !IsNil(o.ServiceXml) {
		toSerialize["service_xml"] = o.ServiceXml
	}
	if !IsNil(o.TemplateType) {
		toSerialize["template_type"] = o.TemplateType
	}
	if !IsNil(o.UserServiceDeliveryType) {
		toSerialize["user_service_delivery_type"] = o.UserServiceDeliveryType
	}
	return toSerialize, nil
}

type NullableAlipayOpenAppServiceApplyModel struct {
	value *AlipayOpenAppServiceApplyModel
	isSet bool
}

func (v NullableAlipayOpenAppServiceApplyModel) Get() *AlipayOpenAppServiceApplyModel {
	return v.value
}

func (v *NullableAlipayOpenAppServiceApplyModel) Set(val *AlipayOpenAppServiceApplyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppServiceApplyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppServiceApplyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppServiceApplyModel(val *AlipayOpenAppServiceApplyModel) *NullableAlipayOpenAppServiceApplyModel {
	return &NullableAlipayOpenAppServiceApplyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAppServiceApplyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppServiceApplyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


