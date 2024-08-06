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

// checks if the DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel{}

// DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel struct for DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel
type DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel struct {
	// 人脸核身具体类型目前仅支持：DATA_DIGITAL_BIZ_CODE_FACE_VERIFICATION
	BizCode *string `json:"biz_code,omitempty"`
	// 真实姓名
	CertName *string `json:"cert_name,omitempty"`
	// 证件号
	CertNo *string `json:"cert_no,omitempty"`
	// 证件类型，当前枚举支持： IDENTITY_CARD：身份证 RESIDENCE_HK_MC：港澳居民居住证 RESIDENCE_TAIWAN：台湾居民居住证
	CertType *string `json:"cert_type,omitempty"`
	// 人脸保存策略，非必填。具体取值为：reserve(保存活体人脸)/never(不保存活体人脸)，不传默认为reserve
	FaceReserveStrategy *string `json:"face_reserve_strategy,omitempty"`
	// 认证类型，固定值为：CERT_INFO
	IdentityType *string `json:"identity_type,omitempty"`
	// 国家地区代码，当证件类型为外国人永久居留证时必填。
	Nation *string `json:"nation,omitempty"`
	// 商户请求的唯一标识，商户要保证其唯一性，值为64位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	OuterOrderNo *string `json:"outer_order_no,omitempty"`
	// 手机号码
	PhoneNo *string `json:"phone_no,omitempty"`
}

// NewDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel instantiates a new DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel() *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel {
	this := DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel{}
	return &this
}

// NewDatadigitalFincloudGeneralsaasFaceVerificationInitializeModelWithDefaults instantiates a new DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatadigitalFincloudGeneralsaasFaceVerificationInitializeModelWithDefaults() *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel {
	this := DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel{}
	return &this
}

// GetBizCode returns the BizCode field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetBizCode() string {
	if o == nil || IsNil(o.BizCode) {
		var ret string
		return ret
	}
	return *o.BizCode
}

// GetBizCodeOk returns a tuple with the BizCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetBizCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BizCode) {
		return nil, false
	}
	return o.BizCode, true
}

// HasBizCode returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasBizCode() bool {
	if o != nil && !IsNil(o.BizCode) {
		return true
	}

	return false
}

// SetBizCode gets a reference to the given string and assigns it to the BizCode field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetBizCode(v string) {
	o.BizCode = &v
}

// GetCertName returns the CertName field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetCertName() string {
	if o == nil || IsNil(o.CertName) {
		var ret string
		return ret
	}
	return *o.CertName
}

// GetCertNameOk returns a tuple with the CertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetCertNameOk() (*string, bool) {
	if o == nil || IsNil(o.CertName) {
		return nil, false
	}
	return o.CertName, true
}

// HasCertName returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasCertName() bool {
	if o != nil && !IsNil(o.CertName) {
		return true
	}

	return false
}

// SetCertName gets a reference to the given string and assigns it to the CertName field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetCertName(v string) {
	o.CertName = &v
}

// GetCertNo returns the CertNo field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetCertNo() string {
	if o == nil || IsNil(o.CertNo) {
		var ret string
		return ret
	}
	return *o.CertNo
}

// GetCertNoOk returns a tuple with the CertNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetCertNoOk() (*string, bool) {
	if o == nil || IsNil(o.CertNo) {
		return nil, false
	}
	return o.CertNo, true
}

// HasCertNo returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasCertNo() bool {
	if o != nil && !IsNil(o.CertNo) {
		return true
	}

	return false
}

// SetCertNo gets a reference to the given string and assigns it to the CertNo field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetCertNo(v string) {
	o.CertNo = &v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetCertType() string {
	if o == nil || IsNil(o.CertType) {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertType) {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasCertType() bool {
	if o != nil && !IsNil(o.CertType) {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetCertType(v string) {
	o.CertType = &v
}

// GetFaceReserveStrategy returns the FaceReserveStrategy field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetFaceReserveStrategy() string {
	if o == nil || IsNil(o.FaceReserveStrategy) {
		var ret string
		return ret
	}
	return *o.FaceReserveStrategy
}

// GetFaceReserveStrategyOk returns a tuple with the FaceReserveStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetFaceReserveStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.FaceReserveStrategy) {
		return nil, false
	}
	return o.FaceReserveStrategy, true
}

// HasFaceReserveStrategy returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasFaceReserveStrategy() bool {
	if o != nil && !IsNil(o.FaceReserveStrategy) {
		return true
	}

	return false
}

// SetFaceReserveStrategy gets a reference to the given string and assigns it to the FaceReserveStrategy field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetFaceReserveStrategy(v string) {
	o.FaceReserveStrategy = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetNation returns the Nation field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetNation() string {
	if o == nil || IsNil(o.Nation) {
		var ret string
		return ret
	}
	return *o.Nation
}

// GetNationOk returns a tuple with the Nation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetNationOk() (*string, bool) {
	if o == nil || IsNil(o.Nation) {
		return nil, false
	}
	return o.Nation, true
}

// HasNation returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasNation() bool {
	if o != nil && !IsNil(o.Nation) {
		return true
	}

	return false
}

// SetNation gets a reference to the given string and assigns it to the Nation field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetNation(v string) {
	o.Nation = &v
}

// GetOuterOrderNo returns the OuterOrderNo field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetOuterOrderNo() string {
	if o == nil || IsNil(o.OuterOrderNo) {
		var ret string
		return ret
	}
	return *o.OuterOrderNo
}

// GetOuterOrderNoOk returns a tuple with the OuterOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetOuterOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OuterOrderNo) {
		return nil, false
	}
	return o.OuterOrderNo, true
}

// HasOuterOrderNo returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasOuterOrderNo() bool {
	if o != nil && !IsNil(o.OuterOrderNo) {
		return true
	}

	return false
}

// SetOuterOrderNo gets a reference to the given string and assigns it to the OuterOrderNo field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetOuterOrderNo(v string) {
	o.OuterOrderNo = &v
}

// GetPhoneNo returns the PhoneNo field value if set, zero value otherwise.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetPhoneNo() string {
	if o == nil || IsNil(o.PhoneNo) {
		var ret string
		return ret
	}
	return *o.PhoneNo
}

// GetPhoneNoOk returns a tuple with the PhoneNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) GetPhoneNoOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNo) {
		return nil, false
	}
	return o.PhoneNo, true
}

// HasPhoneNo returns a boolean if a field has been set.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) HasPhoneNo() bool {
	if o != nil && !IsNil(o.PhoneNo) {
		return true
	}

	return false
}

// SetPhoneNo gets a reference to the given string and assigns it to the PhoneNo field.
func (o *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) SetPhoneNo(v string) {
	o.PhoneNo = &v
}

func (o DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizCode) {
		toSerialize["biz_code"] = o.BizCode
	}
	if !IsNil(o.CertName) {
		toSerialize["cert_name"] = o.CertName
	}
	if !IsNil(o.CertNo) {
		toSerialize["cert_no"] = o.CertNo
	}
	if !IsNil(o.CertType) {
		toSerialize["cert_type"] = o.CertType
	}
	if !IsNil(o.FaceReserveStrategy) {
		toSerialize["face_reserve_strategy"] = o.FaceReserveStrategy
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.Nation) {
		toSerialize["nation"] = o.Nation
	}
	if !IsNil(o.OuterOrderNo) {
		toSerialize["outer_order_no"] = o.OuterOrderNo
	}
	if !IsNil(o.PhoneNo) {
		toSerialize["phone_no"] = o.PhoneNo
	}
	return toSerialize, nil
}

type NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel struct {
	value *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) Get() *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) Set(val *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel(val *DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) *NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel {
	return &NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceVerificationInitializeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


