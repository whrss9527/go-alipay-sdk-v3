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

// checks if the UserInvoiceInfoOpenApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvoiceInfoOpenApiResponse{}

// UserInvoiceInfoOpenApiResponse struct for UserInvoiceInfoOpenApiResponse
type UserInvoiceInfoOpenApiResponse struct {
	// (AR开票使用)是否接受电子票 true:是，false:否
	AcceptElectronic *bool `json:"accept_electronic,omitempty"`
	// 公司注册地址
	Address *string `json:"address,omitempty"`
	// (AR开票使用)是否自动申请开票 true:是，false:否
	Auto *bool `json:"auto,omitempty"`
	// 银行账户
	BankAccount *string `json:"bank_account,omitempty"`
	// 开户行
	BankName *string `json:"bank_name,omitempty"`
	// 营业执照地址
	BusinessLicenceUrl *string `json:"business_licence_url,omitempty"`
	// 资料创建人
	Creator *string `json:"creator,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// (AR开票使用)是否暂停开票  true:暂停开票，false:可开票
	Hold *bool `json:"hold,omitempty"`
	// 开票资料ID
	Id *string `json:"id,omitempty"`
	// 商户PID,  type=01时表示OU用户,填充的是InstId机构代码,例如Z50,  type=02时表示商户MID,  type=03时表示商户PID
	IpId *string `json:"ip_id,omitempty"`
	// 最后修改人
	LastModifier *string `json:"last_modifier,omitempty"`
	// 银行开户许可证地址
	OpenAccountPermitUrl *string `json:"open_account_permit_url,omitempty"`
	// 其它资质证明地址
	OtherQualificationUrl *string `json:"other_qualification_url,omitempty"`
	// 资料状态，01：待审，02：已审(有效)， 03：已作废
	Status *string `json:"status,omitempty"`
	// 上一次任务执行时间(针对自动开票场景),格式:YYYYMMDD
	TaskDate *string `json:"task_date,omitempty"`
	// 纳税人识别号（购方税号）
	TaxNo *string `json:"tax_no,omitempty"`
	// 纳税人开票资格种类  01：一般纳税人，02：小规模纳税人
	TaxPayerQualification *string `json:"tax_payer_qualification,omitempty"`
	// 一般纳税人资格证书地址
	TaxQualificationUrl *string `json:"tax_qualification_url,omitempty"`
	// 税务登记证地址
	TaxRegCertUrl *string `json:"tax_reg_cert_url,omitempty"`
	// 纳税人资格开始时间 （时间戳）
	TaxpayerQualiValid *string `json:"taxpayer_quali_valid,omitempty"`
	// 公司注册电话
	Telephone *string `json:"telephone,omitempty"`
	// 发票抬头
	Title *string `json:"title,omitempty"`
	// 租户ID
	TntInstId *string `json:"tnt_inst_id,omitempty"`
	// 开票资料用户类型   01：OU，02：商户，03：C用户
	Type *string `json:"type,omitempty"`
	// 收件人地址列表
	UserMailInfoList []UserMailInfoVO `json:"user_mail_info_list,omitempty"`
}

// NewUserInvoiceInfoOpenApiResponse instantiates a new UserInvoiceInfoOpenApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvoiceInfoOpenApiResponse() *UserInvoiceInfoOpenApiResponse {
	this := UserInvoiceInfoOpenApiResponse{}
	return &this
}

// NewUserInvoiceInfoOpenApiResponseWithDefaults instantiates a new UserInvoiceInfoOpenApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvoiceInfoOpenApiResponseWithDefaults() *UserInvoiceInfoOpenApiResponse {
	this := UserInvoiceInfoOpenApiResponse{}
	return &this
}

// GetAcceptElectronic returns the AcceptElectronic field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetAcceptElectronic() bool {
	if o == nil || IsNil(o.AcceptElectronic) {
		var ret bool
		return ret
	}
	return *o.AcceptElectronic
}

// GetAcceptElectronicOk returns a tuple with the AcceptElectronic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetAcceptElectronicOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptElectronic) {
		return nil, false
	}
	return o.AcceptElectronic, true
}

// HasAcceptElectronic returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasAcceptElectronic() bool {
	if o != nil && !IsNil(o.AcceptElectronic) {
		return true
	}

	return false
}

// SetAcceptElectronic gets a reference to the given bool and assigns it to the AcceptElectronic field.
func (o *UserInvoiceInfoOpenApiResponse) SetAcceptElectronic(v bool) {
	o.AcceptElectronic = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UserInvoiceInfoOpenApiResponse) SetAddress(v string) {
	o.Address = &v
}

// GetAuto returns the Auto field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetAuto() bool {
	if o == nil || IsNil(o.Auto) {
		var ret bool
		return ret
	}
	return *o.Auto
}

// GetAutoOk returns a tuple with the Auto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetAutoOk() (*bool, bool) {
	if o == nil || IsNil(o.Auto) {
		return nil, false
	}
	return o.Auto, true
}

// HasAuto returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasAuto() bool {
	if o != nil && !IsNil(o.Auto) {
		return true
	}

	return false
}

// SetAuto gets a reference to the given bool and assigns it to the Auto field.
func (o *UserInvoiceInfoOpenApiResponse) SetAuto(v bool) {
	o.Auto = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetBankAccount() string {
	if o == nil || IsNil(o.BankAccount) {
		var ret string
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given string and assigns it to the BankAccount field.
func (o *UserInvoiceInfoOpenApiResponse) SetBankAccount(v string) {
	o.BankAccount = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *UserInvoiceInfoOpenApiResponse) SetBankName(v string) {
	o.BankName = &v
}

// GetBusinessLicenceUrl returns the BusinessLicenceUrl field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetBusinessLicenceUrl() string {
	if o == nil || IsNil(o.BusinessLicenceUrl) {
		var ret string
		return ret
	}
	return *o.BusinessLicenceUrl
}

// GetBusinessLicenceUrlOk returns a tuple with the BusinessLicenceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetBusinessLicenceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLicenceUrl) {
		return nil, false
	}
	return o.BusinessLicenceUrl, true
}

// HasBusinessLicenceUrl returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasBusinessLicenceUrl() bool {
	if o != nil && !IsNil(o.BusinessLicenceUrl) {
		return true
	}

	return false
}

// SetBusinessLicenceUrl gets a reference to the given string and assigns it to the BusinessLicenceUrl field.
func (o *UserInvoiceInfoOpenApiResponse) SetBusinessLicenceUrl(v string) {
	o.BusinessLicenceUrl = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetCreator() string {
	if o == nil || IsNil(o.Creator) {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetCreatorOk() (*string, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *UserInvoiceInfoOpenApiResponse) SetCreator(v string) {
	o.Creator = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *UserInvoiceInfoOpenApiResponse) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *UserInvoiceInfoOpenApiResponse) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetHold returns the Hold field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetHold() bool {
	if o == nil || IsNil(o.Hold) {
		var ret bool
		return ret
	}
	return *o.Hold
}

// GetHoldOk returns a tuple with the Hold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.Hold) {
		return nil, false
	}
	return o.Hold, true
}

// HasHold returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasHold() bool {
	if o != nil && !IsNil(o.Hold) {
		return true
	}

	return false
}

// SetHold gets a reference to the given bool and assigns it to the Hold field.
func (o *UserInvoiceInfoOpenApiResponse) SetHold(v bool) {
	o.Hold = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserInvoiceInfoOpenApiResponse) SetId(v string) {
	o.Id = &v
}

// GetIpId returns the IpId field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetIpId() string {
	if o == nil || IsNil(o.IpId) {
		var ret string
		return ret
	}
	return *o.IpId
}

// GetIpIdOk returns a tuple with the IpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetIpIdOk() (*string, bool) {
	if o == nil || IsNil(o.IpId) {
		return nil, false
	}
	return o.IpId, true
}

// HasIpId returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasIpId() bool {
	if o != nil && !IsNil(o.IpId) {
		return true
	}

	return false
}

// SetIpId gets a reference to the given string and assigns it to the IpId field.
func (o *UserInvoiceInfoOpenApiResponse) SetIpId(v string) {
	o.IpId = &v
}

// GetLastModifier returns the LastModifier field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetLastModifier() string {
	if o == nil || IsNil(o.LastModifier) {
		var ret string
		return ret
	}
	return *o.LastModifier
}

// GetLastModifierOk returns a tuple with the LastModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetLastModifierOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifier) {
		return nil, false
	}
	return o.LastModifier, true
}

// HasLastModifier returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasLastModifier() bool {
	if o != nil && !IsNil(o.LastModifier) {
		return true
	}

	return false
}

// SetLastModifier gets a reference to the given string and assigns it to the LastModifier field.
func (o *UserInvoiceInfoOpenApiResponse) SetLastModifier(v string) {
	o.LastModifier = &v
}

// GetOpenAccountPermitUrl returns the OpenAccountPermitUrl field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetOpenAccountPermitUrl() string {
	if o == nil || IsNil(o.OpenAccountPermitUrl) {
		var ret string
		return ret
	}
	return *o.OpenAccountPermitUrl
}

// GetOpenAccountPermitUrlOk returns a tuple with the OpenAccountPermitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetOpenAccountPermitUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OpenAccountPermitUrl) {
		return nil, false
	}
	return o.OpenAccountPermitUrl, true
}

// HasOpenAccountPermitUrl returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasOpenAccountPermitUrl() bool {
	if o != nil && !IsNil(o.OpenAccountPermitUrl) {
		return true
	}

	return false
}

// SetOpenAccountPermitUrl gets a reference to the given string and assigns it to the OpenAccountPermitUrl field.
func (o *UserInvoiceInfoOpenApiResponse) SetOpenAccountPermitUrl(v string) {
	o.OpenAccountPermitUrl = &v
}

// GetOtherQualificationUrl returns the OtherQualificationUrl field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetOtherQualificationUrl() string {
	if o == nil || IsNil(o.OtherQualificationUrl) {
		var ret string
		return ret
	}
	return *o.OtherQualificationUrl
}

// GetOtherQualificationUrlOk returns a tuple with the OtherQualificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetOtherQualificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OtherQualificationUrl) {
		return nil, false
	}
	return o.OtherQualificationUrl, true
}

// HasOtherQualificationUrl returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasOtherQualificationUrl() bool {
	if o != nil && !IsNil(o.OtherQualificationUrl) {
		return true
	}

	return false
}

// SetOtherQualificationUrl gets a reference to the given string and assigns it to the OtherQualificationUrl field.
func (o *UserInvoiceInfoOpenApiResponse) SetOtherQualificationUrl(v string) {
	o.OtherQualificationUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserInvoiceInfoOpenApiResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTaskDate returns the TaskDate field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTaskDate() string {
	if o == nil || IsNil(o.TaskDate) {
		var ret string
		return ret
	}
	return *o.TaskDate
}

// GetTaskDateOk returns a tuple with the TaskDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTaskDateOk() (*string, bool) {
	if o == nil || IsNil(o.TaskDate) {
		return nil, false
	}
	return o.TaskDate, true
}

// HasTaskDate returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTaskDate() bool {
	if o != nil && !IsNil(o.TaskDate) {
		return true
	}

	return false
}

// SetTaskDate gets a reference to the given string and assigns it to the TaskDate field.
func (o *UserInvoiceInfoOpenApiResponse) SetTaskDate(v string) {
	o.TaskDate = &v
}

// GetTaxNo returns the TaxNo field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxNo() string {
	if o == nil || IsNil(o.TaxNo) {
		var ret string
		return ret
	}
	return *o.TaxNo
}

// GetTaxNoOk returns a tuple with the TaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxNoOk() (*string, bool) {
	if o == nil || IsNil(o.TaxNo) {
		return nil, false
	}
	return o.TaxNo, true
}

// HasTaxNo returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTaxNo() bool {
	if o != nil && !IsNil(o.TaxNo) {
		return true
	}

	return false
}

// SetTaxNo gets a reference to the given string and assigns it to the TaxNo field.
func (o *UserInvoiceInfoOpenApiResponse) SetTaxNo(v string) {
	o.TaxNo = &v
}

// GetTaxPayerQualification returns the TaxPayerQualification field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxPayerQualification() string {
	if o == nil || IsNil(o.TaxPayerQualification) {
		var ret string
		return ret
	}
	return *o.TaxPayerQualification
}

// GetTaxPayerQualificationOk returns a tuple with the TaxPayerQualification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxPayerQualificationOk() (*string, bool) {
	if o == nil || IsNil(o.TaxPayerQualification) {
		return nil, false
	}
	return o.TaxPayerQualification, true
}

// HasTaxPayerQualification returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTaxPayerQualification() bool {
	if o != nil && !IsNil(o.TaxPayerQualification) {
		return true
	}

	return false
}

// SetTaxPayerQualification gets a reference to the given string and assigns it to the TaxPayerQualification field.
func (o *UserInvoiceInfoOpenApiResponse) SetTaxPayerQualification(v string) {
	o.TaxPayerQualification = &v
}

// GetTaxQualificationUrl returns the TaxQualificationUrl field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxQualificationUrl() string {
	if o == nil || IsNil(o.TaxQualificationUrl) {
		var ret string
		return ret
	}
	return *o.TaxQualificationUrl
}

// GetTaxQualificationUrlOk returns a tuple with the TaxQualificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxQualificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TaxQualificationUrl) {
		return nil, false
	}
	return o.TaxQualificationUrl, true
}

// HasTaxQualificationUrl returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTaxQualificationUrl() bool {
	if o != nil && !IsNil(o.TaxQualificationUrl) {
		return true
	}

	return false
}

// SetTaxQualificationUrl gets a reference to the given string and assigns it to the TaxQualificationUrl field.
func (o *UserInvoiceInfoOpenApiResponse) SetTaxQualificationUrl(v string) {
	o.TaxQualificationUrl = &v
}

// GetTaxRegCertUrl returns the TaxRegCertUrl field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxRegCertUrl() string {
	if o == nil || IsNil(o.TaxRegCertUrl) {
		var ret string
		return ret
	}
	return *o.TaxRegCertUrl
}

// GetTaxRegCertUrlOk returns a tuple with the TaxRegCertUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxRegCertUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRegCertUrl) {
		return nil, false
	}
	return o.TaxRegCertUrl, true
}

// HasTaxRegCertUrl returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTaxRegCertUrl() bool {
	if o != nil && !IsNil(o.TaxRegCertUrl) {
		return true
	}

	return false
}

// SetTaxRegCertUrl gets a reference to the given string and assigns it to the TaxRegCertUrl field.
func (o *UserInvoiceInfoOpenApiResponse) SetTaxRegCertUrl(v string) {
	o.TaxRegCertUrl = &v
}

// GetTaxpayerQualiValid returns the TaxpayerQualiValid field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxpayerQualiValid() string {
	if o == nil || IsNil(o.TaxpayerQualiValid) {
		var ret string
		return ret
	}
	return *o.TaxpayerQualiValid
}

// GetTaxpayerQualiValidOk returns a tuple with the TaxpayerQualiValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTaxpayerQualiValidOk() (*string, bool) {
	if o == nil || IsNil(o.TaxpayerQualiValid) {
		return nil, false
	}
	return o.TaxpayerQualiValid, true
}

// HasTaxpayerQualiValid returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTaxpayerQualiValid() bool {
	if o != nil && !IsNil(o.TaxpayerQualiValid) {
		return true
	}

	return false
}

// SetTaxpayerQualiValid gets a reference to the given string and assigns it to the TaxpayerQualiValid field.
func (o *UserInvoiceInfoOpenApiResponse) SetTaxpayerQualiValid(v string) {
	o.TaxpayerQualiValid = &v
}

// GetTelephone returns the Telephone field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTelephone() string {
	if o == nil || IsNil(o.Telephone) {
		var ret string
		return ret
	}
	return *o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTelephoneOk() (*string, bool) {
	if o == nil || IsNil(o.Telephone) {
		return nil, false
	}
	return o.Telephone, true
}

// HasTelephone returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTelephone() bool {
	if o != nil && !IsNil(o.Telephone) {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given string and assigns it to the Telephone field.
func (o *UserInvoiceInfoOpenApiResponse) SetTelephone(v string) {
	o.Telephone = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserInvoiceInfoOpenApiResponse) SetTitle(v string) {
	o.Title = &v
}

// GetTntInstId returns the TntInstId field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetTntInstId() string {
	if o == nil || IsNil(o.TntInstId) {
		var ret string
		return ret
	}
	return *o.TntInstId
}

// GetTntInstIdOk returns a tuple with the TntInstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTntInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.TntInstId) {
		return nil, false
	}
	return o.TntInstId, true
}

// HasTntInstId returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasTntInstId() bool {
	if o != nil && !IsNil(o.TntInstId) {
		return true
	}

	return false
}

// SetTntInstId gets a reference to the given string and assigns it to the TntInstId field.
func (o *UserInvoiceInfoOpenApiResponse) SetTntInstId(v string) {
	o.TntInstId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserInvoiceInfoOpenApiResponse) SetType(v string) {
	o.Type = &v
}

// GetUserMailInfoList returns the UserMailInfoList field value if set, zero value otherwise.
func (o *UserInvoiceInfoOpenApiResponse) GetUserMailInfoList() []UserMailInfoVO {
	if o == nil || IsNil(o.UserMailInfoList) {
		var ret []UserMailInfoVO
		return ret
	}
	return o.UserMailInfoList
}

// GetUserMailInfoListOk returns a tuple with the UserMailInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvoiceInfoOpenApiResponse) GetUserMailInfoListOk() ([]UserMailInfoVO, bool) {
	if o == nil || IsNil(o.UserMailInfoList) {
		return nil, false
	}
	return o.UserMailInfoList, true
}

// HasUserMailInfoList returns a boolean if a field has been set.
func (o *UserInvoiceInfoOpenApiResponse) HasUserMailInfoList() bool {
	if o != nil && !IsNil(o.UserMailInfoList) {
		return true
	}

	return false
}

// SetUserMailInfoList gets a reference to the given []UserMailInfoVO and assigns it to the UserMailInfoList field.
func (o *UserInvoiceInfoOpenApiResponse) SetUserMailInfoList(v []UserMailInfoVO) {
	o.UserMailInfoList = v
}

func (o UserInvoiceInfoOpenApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvoiceInfoOpenApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptElectronic) {
		toSerialize["accept_electronic"] = o.AcceptElectronic
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Auto) {
		toSerialize["auto"] = o.Auto
	}
	if !IsNil(o.BankAccount) {
		toSerialize["bank_account"] = o.BankAccount
	}
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.BusinessLicenceUrl) {
		toSerialize["business_licence_url"] = o.BusinessLicenceUrl
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmt_create"] = o.GmtCreate
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.Hold) {
		toSerialize["hold"] = o.Hold
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpId) {
		toSerialize["ip_id"] = o.IpId
	}
	if !IsNil(o.LastModifier) {
		toSerialize["last_modifier"] = o.LastModifier
	}
	if !IsNil(o.OpenAccountPermitUrl) {
		toSerialize["open_account_permit_url"] = o.OpenAccountPermitUrl
	}
	if !IsNil(o.OtherQualificationUrl) {
		toSerialize["other_qualification_url"] = o.OtherQualificationUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TaskDate) {
		toSerialize["task_date"] = o.TaskDate
	}
	if !IsNil(o.TaxNo) {
		toSerialize["tax_no"] = o.TaxNo
	}
	if !IsNil(o.TaxPayerQualification) {
		toSerialize["tax_payer_qualification"] = o.TaxPayerQualification
	}
	if !IsNil(o.TaxQualificationUrl) {
		toSerialize["tax_qualification_url"] = o.TaxQualificationUrl
	}
	if !IsNil(o.TaxRegCertUrl) {
		toSerialize["tax_reg_cert_url"] = o.TaxRegCertUrl
	}
	if !IsNil(o.TaxpayerQualiValid) {
		toSerialize["taxpayer_quali_valid"] = o.TaxpayerQualiValid
	}
	if !IsNil(o.Telephone) {
		toSerialize["telephone"] = o.Telephone
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.TntInstId) {
		toSerialize["tnt_inst_id"] = o.TntInstId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserMailInfoList) {
		toSerialize["user_mail_info_list"] = o.UserMailInfoList
	}
	return toSerialize, nil
}

type NullableUserInvoiceInfoOpenApiResponse struct {
	value *UserInvoiceInfoOpenApiResponse
	isSet bool
}

func (v NullableUserInvoiceInfoOpenApiResponse) Get() *UserInvoiceInfoOpenApiResponse {
	return v.value
}

func (v *NullableUserInvoiceInfoOpenApiResponse) Set(val *UserInvoiceInfoOpenApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvoiceInfoOpenApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvoiceInfoOpenApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvoiceInfoOpenApiResponse(val *UserInvoiceInfoOpenApiResponse) *NullableUserInvoiceInfoOpenApiResponse {
	return &NullableUserInvoiceInfoOpenApiResponse{value: val, isSet: true}
}

func (v NullableUserInvoiceInfoOpenApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvoiceInfoOpenApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
