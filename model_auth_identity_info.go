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

// checks if the AuthIdentityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthIdentityInfo{}

// AuthIdentityInfo struct for AuthIdentityInfo
type AuthIdentityInfo struct {
	CertificateInfo *IndirectCertificateInfo `json:"certificate_info,omitempty"`
	// 主体类型为企业/个体户/政府机关/事业单位/社会组织时，必填； 证照类型：营业执照(BUSINESS_CERT)/登记证书(REGISTER_CERT) 主体为政府机关/事业单位/社会组织时，填登记证书； 主体类型为企业/个体户时，填营业执照；
	CertificateType *string `json:"certificate_type,omitempty"`
	// 单位证明函照片（使用图片上传接口）主体类型为政府机关/事业单位时，单位证明函照片必填
	EmployerLetterImg *string `json:"employer_letter_img,omitempty"`
	FinancialOrgInfo *IndirectFinancialOrgInfo `json:"financial_org_info,omitempty"`
	// 主体类型，枚举定义：企业(ENTERPRISE)、个体工商户(IND_BIZ)、事业单位(INST)、党政机关(GOV)、社会组织(ORG)、小微商户(MSE)
	IdentityType *string `json:"identity_type,omitempty"`
	// 是否金融机构
	IsFinancialOrg *bool `json:"is_financial_org,omitempty"`
	// 经营许可证列表，填写特殊行业的经营许可证信息，一个主体最多5个行业
	QualificationInfoList []IndirectQualificationInfo `json:"qualification_info_list,omitempty"`
	SupportCredentials *IndirectSupportCredentials `json:"support_credentials,omitempty"`
}

// NewAuthIdentityInfo instantiates a new AuthIdentityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthIdentityInfo() *AuthIdentityInfo {
	this := AuthIdentityInfo{}
	return &this
}

// NewAuthIdentityInfoWithDefaults instantiates a new AuthIdentityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthIdentityInfoWithDefaults() *AuthIdentityInfo {
	this := AuthIdentityInfo{}
	return &this
}

// GetCertificateInfo returns the CertificateInfo field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetCertificateInfo() IndirectCertificateInfo {
	if o == nil || IsNil(o.CertificateInfo) {
		var ret IndirectCertificateInfo
		return ret
	}
	return *o.CertificateInfo
}

// GetCertificateInfoOk returns a tuple with the CertificateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetCertificateInfoOk() (*IndirectCertificateInfo, bool) {
	if o == nil || IsNil(o.CertificateInfo) {
		return nil, false
	}
	return o.CertificateInfo, true
}

// HasCertificateInfo returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasCertificateInfo() bool {
	if o != nil && !IsNil(o.CertificateInfo) {
		return true
	}

	return false
}

// SetCertificateInfo gets a reference to the given IndirectCertificateInfo and assigns it to the CertificateInfo field.
func (o *AuthIdentityInfo) SetCertificateInfo(v IndirectCertificateInfo) {
	o.CertificateInfo = &v
}

// GetCertificateType returns the CertificateType field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetCertificateType() string {
	if o == nil || IsNil(o.CertificateType) {
		var ret string
		return ret
	}
	return *o.CertificateType
}

// GetCertificateTypeOk returns a tuple with the CertificateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetCertificateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateType) {
		return nil, false
	}
	return o.CertificateType, true
}

// HasCertificateType returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasCertificateType() bool {
	if o != nil && !IsNil(o.CertificateType) {
		return true
	}

	return false
}

// SetCertificateType gets a reference to the given string and assigns it to the CertificateType field.
func (o *AuthIdentityInfo) SetCertificateType(v string) {
	o.CertificateType = &v
}

// GetEmployerLetterImg returns the EmployerLetterImg field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetEmployerLetterImg() string {
	if o == nil || IsNil(o.EmployerLetterImg) {
		var ret string
		return ret
	}
	return *o.EmployerLetterImg
}

// GetEmployerLetterImgOk returns a tuple with the EmployerLetterImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetEmployerLetterImgOk() (*string, bool) {
	if o == nil || IsNil(o.EmployerLetterImg) {
		return nil, false
	}
	return o.EmployerLetterImg, true
}

// HasEmployerLetterImg returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasEmployerLetterImg() bool {
	if o != nil && !IsNil(o.EmployerLetterImg) {
		return true
	}

	return false
}

// SetEmployerLetterImg gets a reference to the given string and assigns it to the EmployerLetterImg field.
func (o *AuthIdentityInfo) SetEmployerLetterImg(v string) {
	o.EmployerLetterImg = &v
}

// GetFinancialOrgInfo returns the FinancialOrgInfo field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetFinancialOrgInfo() IndirectFinancialOrgInfo {
	if o == nil || IsNil(o.FinancialOrgInfo) {
		var ret IndirectFinancialOrgInfo
		return ret
	}
	return *o.FinancialOrgInfo
}

// GetFinancialOrgInfoOk returns a tuple with the FinancialOrgInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetFinancialOrgInfoOk() (*IndirectFinancialOrgInfo, bool) {
	if o == nil || IsNil(o.FinancialOrgInfo) {
		return nil, false
	}
	return o.FinancialOrgInfo, true
}

// HasFinancialOrgInfo returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasFinancialOrgInfo() bool {
	if o != nil && !IsNil(o.FinancialOrgInfo) {
		return true
	}

	return false
}

// SetFinancialOrgInfo gets a reference to the given IndirectFinancialOrgInfo and assigns it to the FinancialOrgInfo field.
func (o *AuthIdentityInfo) SetFinancialOrgInfo(v IndirectFinancialOrgInfo) {
	o.FinancialOrgInfo = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *AuthIdentityInfo) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetIsFinancialOrg returns the IsFinancialOrg field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetIsFinancialOrg() bool {
	if o == nil || IsNil(o.IsFinancialOrg) {
		var ret bool
		return ret
	}
	return *o.IsFinancialOrg
}

// GetIsFinancialOrgOk returns a tuple with the IsFinancialOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetIsFinancialOrgOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFinancialOrg) {
		return nil, false
	}
	return o.IsFinancialOrg, true
}

// HasIsFinancialOrg returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasIsFinancialOrg() bool {
	if o != nil && !IsNil(o.IsFinancialOrg) {
		return true
	}

	return false
}

// SetIsFinancialOrg gets a reference to the given bool and assigns it to the IsFinancialOrg field.
func (o *AuthIdentityInfo) SetIsFinancialOrg(v bool) {
	o.IsFinancialOrg = &v
}

// GetQualificationInfoList returns the QualificationInfoList field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetQualificationInfoList() []IndirectQualificationInfo {
	if o == nil || IsNil(o.QualificationInfoList) {
		var ret []IndirectQualificationInfo
		return ret
	}
	return o.QualificationInfoList
}

// GetQualificationInfoListOk returns a tuple with the QualificationInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetQualificationInfoListOk() ([]IndirectQualificationInfo, bool) {
	if o == nil || IsNil(o.QualificationInfoList) {
		return nil, false
	}
	return o.QualificationInfoList, true
}

// HasQualificationInfoList returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasQualificationInfoList() bool {
	if o != nil && !IsNil(o.QualificationInfoList) {
		return true
	}

	return false
}

// SetQualificationInfoList gets a reference to the given []IndirectQualificationInfo and assigns it to the QualificationInfoList field.
func (o *AuthIdentityInfo) SetQualificationInfoList(v []IndirectQualificationInfo) {
	o.QualificationInfoList = v
}

// GetSupportCredentials returns the SupportCredentials field value if set, zero value otherwise.
func (o *AuthIdentityInfo) GetSupportCredentials() IndirectSupportCredentials {
	if o == nil || IsNil(o.SupportCredentials) {
		var ret IndirectSupportCredentials
		return ret
	}
	return *o.SupportCredentials
}

// GetSupportCredentialsOk returns a tuple with the SupportCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthIdentityInfo) GetSupportCredentialsOk() (*IndirectSupportCredentials, bool) {
	if o == nil || IsNil(o.SupportCredentials) {
		return nil, false
	}
	return o.SupportCredentials, true
}

// HasSupportCredentials returns a boolean if a field has been set.
func (o *AuthIdentityInfo) HasSupportCredentials() bool {
	if o != nil && !IsNil(o.SupportCredentials) {
		return true
	}

	return false
}

// SetSupportCredentials gets a reference to the given IndirectSupportCredentials and assigns it to the SupportCredentials field.
func (o *AuthIdentityInfo) SetSupportCredentials(v IndirectSupportCredentials) {
	o.SupportCredentials = &v
}

func (o AuthIdentityInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthIdentityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificateInfo) {
		toSerialize["certificate_info"] = o.CertificateInfo
	}
	if !IsNil(o.CertificateType) {
		toSerialize["certificate_type"] = o.CertificateType
	}
	if !IsNil(o.EmployerLetterImg) {
		toSerialize["employer_letter_img"] = o.EmployerLetterImg
	}
	if !IsNil(o.FinancialOrgInfo) {
		toSerialize["financial_org_info"] = o.FinancialOrgInfo
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.IsFinancialOrg) {
		toSerialize["is_financial_org"] = o.IsFinancialOrg
	}
	if !IsNil(o.QualificationInfoList) {
		toSerialize["qualification_info_list"] = o.QualificationInfoList
	}
	if !IsNil(o.SupportCredentials) {
		toSerialize["support_credentials"] = o.SupportCredentials
	}
	return toSerialize, nil
}

type NullableAuthIdentityInfo struct {
	value *AuthIdentityInfo
	isSet bool
}

func (v NullableAuthIdentityInfo) Get() *AuthIdentityInfo {
	return v.value
}

func (v *NullableAuthIdentityInfo) Set(val *AuthIdentityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthIdentityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthIdentityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthIdentityInfo(val *AuthIdentityInfo) *NullableAuthIdentityInfo {
	return &NullableAuthIdentityInfo{value: val, isSet: true}
}

func (v NullableAuthIdentityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthIdentityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


