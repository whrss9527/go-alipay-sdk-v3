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

// checks if the AlipayEbppInvoiceDetailOutputQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceDetailOutputQueryResponseModel{}

// AlipayEbppInvoiceDetailOutputQueryResponseModel struct for AlipayEbppInvoiceDetailOutputQueryResponseModel
type AlipayEbppInvoiceDetailOutputQueryResponseModel struct {
	// 防伪码，发票校验码
	AntiFakeCode *string `json:"anti_fake_code,omitempty"`
	// 不含税金额（废弃），单位元
	ExTaxAmount *string `json:"ex_tax_amount,omitempty"`
	// 发票报销状态  取值范围：  WAIT_EXPENSE－未报销  EXPENSE_PROCESSING－报销中  EXPENSE_FINISHED－已报销
	ExpenseStatus *string `json:"expense_status,omitempty"`
	// 该发票可能存在异常，请核实后使用  true:无异常  false:存在异常
	HasRisk *bool `json:"has_risk,omitempty"`
	// 是否已上传发票pdf文件  false－未上传  true－已上传
	HasUploadPdf *string `json:"has_upload_pdf,omitempty"`
	// 发票金额（价税合计金额），单位元
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// 发票代码 （全电票为空）
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// 开票日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// 发票缩略图地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// 发票内容项
	InvoiceItemContentList []InvoiceItemContent `json:"invoice_item_content_list,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// 发票状态  SUCCEED－成功  EXPIRED－已失效  TRANSFERRED－已转交
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// 发票类型  值有两种情况：blue－蓝票 red－红票
	InvoiceType *string `json:"invoice_type,omitempty"`
	// 销售方地址
	PayeeAddress *string `json:"payee_address,omitempty"`
	// 销售方开户银行账号
	PayeeBankAccount *string `json:"payee_bank_account,omitempty"`
	// 销售方开户银行名称
	PayeeBankName *string `json:"payee_bank_name,omitempty"`
	// 销售方电话
	PayeePhone *string `json:"payee_phone,omitempty"`
	// 销售方名称
	PayeeRegisterName *string `json:"payee_register_name,omitempty"`
	// 销售方税号
	PayeeTaxNo *string `json:"payee_tax_no,omitempty"`
	// 购买方地址
	PayerAddress *string `json:"payer_address,omitempty"`
	// 购买方开户银行账号
	PayerBankAccount *string `json:"payer_bank_account,omitempty"`
	// 购买方开户银行名称
	PayerBankName *string `json:"payer_bank_name,omitempty"`
	// 购买方名称
	PayerName *string `json:"payer_name,omitempty"`
	// 购买方电话
	PayerPhone *string `json:"payer_phone,omitempty"`
	// 购买方税号
	PayerTaxNo *string `json:"payer_tax_no,omitempty"`
	// 合计税额（废弃）
	SumTaxAmount *string `json:"sum_tax_amount,omitempty"`
	// 票种 PLAIN：增值税电子普通发票 SPECIAL：增值税专用发票 ALL_ELECTRONIC_GENERAL：电子发票（普通发票） ALL_ELECTRONIC_SPECIAL：电子发票（增值税专用发票） PLAIN_INVOICE:增值税普通发票 PAPER_INVOICE:增值税普通发票（卷式） SALSE_INVOICE:机动车销售统一发票
	TaxType *string `json:"tax_type,omitempty"`
}

// NewAlipayEbppInvoiceDetailOutputQueryResponseModel instantiates a new AlipayEbppInvoiceDetailOutputQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceDetailOutputQueryResponseModel() *AlipayEbppInvoiceDetailOutputQueryResponseModel {
	this := AlipayEbppInvoiceDetailOutputQueryResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceDetailOutputQueryResponseModelWithDefaults instantiates a new AlipayEbppInvoiceDetailOutputQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceDetailOutputQueryResponseModelWithDefaults() *AlipayEbppInvoiceDetailOutputQueryResponseModel {
	this := AlipayEbppInvoiceDetailOutputQueryResponseModel{}
	return &this
}

// GetAntiFakeCode returns the AntiFakeCode field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetAntiFakeCode() string {
	if o == nil || IsNil(o.AntiFakeCode) {
		var ret string
		return ret
	}
	return *o.AntiFakeCode
}

// GetAntiFakeCodeOk returns a tuple with the AntiFakeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetAntiFakeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AntiFakeCode) {
		return nil, false
	}
	return o.AntiFakeCode, true
}

// HasAntiFakeCode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasAntiFakeCode() bool {
	if o != nil && !IsNil(o.AntiFakeCode) {
		return true
	}

	return false
}

// SetAntiFakeCode gets a reference to the given string and assigns it to the AntiFakeCode field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetAntiFakeCode(v string) {
	o.AntiFakeCode = &v
}

// GetExTaxAmount returns the ExTaxAmount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetExTaxAmount() string {
	if o == nil || IsNil(o.ExTaxAmount) {
		var ret string
		return ret
	}
	return *o.ExTaxAmount
}

// GetExTaxAmountOk returns a tuple with the ExTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetExTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ExTaxAmount) {
		return nil, false
	}
	return o.ExTaxAmount, true
}

// HasExTaxAmount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasExTaxAmount() bool {
	if o != nil && !IsNil(o.ExTaxAmount) {
		return true
	}

	return false
}

// SetExTaxAmount gets a reference to the given string and assigns it to the ExTaxAmount field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetExTaxAmount(v string) {
	o.ExTaxAmount = &v
}

// GetExpenseStatus returns the ExpenseStatus field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetExpenseStatus() string {
	if o == nil || IsNil(o.ExpenseStatus) {
		var ret string
		return ret
	}
	return *o.ExpenseStatus
}

// GetExpenseStatusOk returns a tuple with the ExpenseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetExpenseStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseStatus) {
		return nil, false
	}
	return o.ExpenseStatus, true
}

// HasExpenseStatus returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasExpenseStatus() bool {
	if o != nil && !IsNil(o.ExpenseStatus) {
		return true
	}

	return false
}

// SetExpenseStatus gets a reference to the given string and assigns it to the ExpenseStatus field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetExpenseStatus(v string) {
	o.ExpenseStatus = &v
}

// GetHasRisk returns the HasRisk field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetHasRisk() bool {
	if o == nil || IsNil(o.HasRisk) {
		var ret bool
		return ret
	}
	return *o.HasRisk
}

// GetHasRiskOk returns a tuple with the HasRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetHasRiskOk() (*bool, bool) {
	if o == nil || IsNil(o.HasRisk) {
		return nil, false
	}
	return o.HasRisk, true
}

// HasHasRisk returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasHasRisk() bool {
	if o != nil && !IsNil(o.HasRisk) {
		return true
	}

	return false
}

// SetHasRisk gets a reference to the given bool and assigns it to the HasRisk field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetHasRisk(v bool) {
	o.HasRisk = &v
}

// GetHasUploadPdf returns the HasUploadPdf field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetHasUploadPdf() string {
	if o == nil || IsNil(o.HasUploadPdf) {
		var ret string
		return ret
	}
	return *o.HasUploadPdf
}

// GetHasUploadPdfOk returns a tuple with the HasUploadPdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetHasUploadPdfOk() (*string, bool) {
	if o == nil || IsNil(o.HasUploadPdf) {
		return nil, false
	}
	return o.HasUploadPdf, true
}

// HasHasUploadPdf returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasHasUploadPdf() bool {
	if o != nil && !IsNil(o.HasUploadPdf) {
		return true
	}

	return false
}

// SetHasUploadPdf gets a reference to the given string and assigns it to the HasUploadPdf field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetHasUploadPdf(v string) {
	o.HasUploadPdf = &v
}

// GetInvoiceAmount returns the InvoiceAmount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceAmount() string {
	if o == nil || IsNil(o.InvoiceAmount) {
		var ret string
		return ret
	}
	return *o.InvoiceAmount
}

// GetInvoiceAmountOk returns a tuple with the InvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceAmountOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceAmount) {
		return nil, false
	}
	return o.InvoiceAmount, true
}

// HasInvoiceAmount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceAmount() bool {
	if o != nil && !IsNil(o.InvoiceAmount) {
		return true
	}

	return false
}

// SetInvoiceAmount gets a reference to the given string and assigns it to the InvoiceAmount field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceAmount(v string) {
	o.InvoiceAmount = &v
}

// GetInvoiceCode returns the InvoiceCode field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceCode() string {
	if o == nil || IsNil(o.InvoiceCode) {
		var ret string
		return ret
	}
	return *o.InvoiceCode
}

// GetInvoiceCodeOk returns a tuple with the InvoiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCode) {
		return nil, false
	}
	return o.InvoiceCode, true
}

// HasInvoiceCode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceCode() bool {
	if o != nil && !IsNil(o.InvoiceCode) {
		return true
	}

	return false
}

// SetInvoiceCode gets a reference to the given string and assigns it to the InvoiceCode field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceCode(v string) {
	o.InvoiceCode = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetInvoiceImgUrl returns the InvoiceImgUrl field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceImgUrl() string {
	if o == nil || IsNil(o.InvoiceImgUrl) {
		var ret string
		return ret
	}
	return *o.InvoiceImgUrl
}

// GetInvoiceImgUrlOk returns a tuple with the InvoiceImgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceImgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceImgUrl) {
		return nil, false
	}
	return o.InvoiceImgUrl, true
}

// HasInvoiceImgUrl returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceImgUrl() bool {
	if o != nil && !IsNil(o.InvoiceImgUrl) {
		return true
	}

	return false
}

// SetInvoiceImgUrl gets a reference to the given string and assigns it to the InvoiceImgUrl field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceImgUrl(v string) {
	o.InvoiceImgUrl = &v
}

// GetInvoiceItemContentList returns the InvoiceItemContentList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceItemContentList() []InvoiceItemContent {
	if o == nil || IsNil(o.InvoiceItemContentList) {
		var ret []InvoiceItemContent
		return ret
	}
	return o.InvoiceItemContentList
}

// GetInvoiceItemContentListOk returns a tuple with the InvoiceItemContentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceItemContentListOk() ([]InvoiceItemContent, bool) {
	if o == nil || IsNil(o.InvoiceItemContentList) {
		return nil, false
	}
	return o.InvoiceItemContentList, true
}

// HasInvoiceItemContentList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceItemContentList() bool {
	if o != nil && !IsNil(o.InvoiceItemContentList) {
		return true
	}

	return false
}

// SetInvoiceItemContentList gets a reference to the given []InvoiceItemContent and assigns it to the InvoiceItemContentList field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceItemContentList(v []InvoiceItemContent) {
	o.InvoiceItemContentList = v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceNo() string {
	if o == nil || IsNil(o.InvoiceNo) {
		var ret string
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceNoOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNo) {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceNo() bool {
	if o != nil && !IsNil(o.InvoiceNo) {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given string and assigns it to the InvoiceNo field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceNo(v string) {
	o.InvoiceNo = &v
}

// GetInvoiceStatus returns the InvoiceStatus field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceStatus() string {
	if o == nil || IsNil(o.InvoiceStatus) {
		var ret string
		return ret
	}
	return *o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceStatus) {
		return nil, false
	}
	return o.InvoiceStatus, true
}

// HasInvoiceStatus returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceStatus() bool {
	if o != nil && !IsNil(o.InvoiceStatus) {
		return true
	}

	return false
}

// SetInvoiceStatus gets a reference to the given string and assigns it to the InvoiceStatus field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceStatus(v string) {
	o.InvoiceStatus = &v
}

// GetInvoiceType returns the InvoiceType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceType() string {
	if o == nil || IsNil(o.InvoiceType) {
		var ret string
		return ret
	}
	return *o.InvoiceType
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetInvoiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceType) {
		return nil, false
	}
	return o.InvoiceType, true
}

// HasInvoiceType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasInvoiceType() bool {
	if o != nil && !IsNil(o.InvoiceType) {
		return true
	}

	return false
}

// SetInvoiceType gets a reference to the given string and assigns it to the InvoiceType field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetInvoiceType(v string) {
	o.InvoiceType = &v
}

// GetPayeeAddress returns the PayeeAddress field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeAddress() string {
	if o == nil || IsNil(o.PayeeAddress) {
		var ret string
		return ret
	}
	return *o.PayeeAddress
}

// GetPayeeAddressOk returns a tuple with the PayeeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeAddress) {
		return nil, false
	}
	return o.PayeeAddress, true
}

// HasPayeeAddress returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayeeAddress() bool {
	if o != nil && !IsNil(o.PayeeAddress) {
		return true
	}

	return false
}

// SetPayeeAddress gets a reference to the given string and assigns it to the PayeeAddress field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayeeAddress(v string) {
	o.PayeeAddress = &v
}

// GetPayeeBankAccount returns the PayeeBankAccount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeBankAccount() string {
	if o == nil || IsNil(o.PayeeBankAccount) {
		var ret string
		return ret
	}
	return *o.PayeeBankAccount
}

// GetPayeeBankAccountOk returns a tuple with the PayeeBankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeBankAccount) {
		return nil, false
	}
	return o.PayeeBankAccount, true
}

// HasPayeeBankAccount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayeeBankAccount() bool {
	if o != nil && !IsNil(o.PayeeBankAccount) {
		return true
	}

	return false
}

// SetPayeeBankAccount gets a reference to the given string and assigns it to the PayeeBankAccount field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayeeBankAccount(v string) {
	o.PayeeBankAccount = &v
}

// GetPayeeBankName returns the PayeeBankName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeBankName() string {
	if o == nil || IsNil(o.PayeeBankName) {
		var ret string
		return ret
	}
	return *o.PayeeBankName
}

// GetPayeeBankNameOk returns a tuple with the PayeeBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeBankName) {
		return nil, false
	}
	return o.PayeeBankName, true
}

// HasPayeeBankName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayeeBankName() bool {
	if o != nil && !IsNil(o.PayeeBankName) {
		return true
	}

	return false
}

// SetPayeeBankName gets a reference to the given string and assigns it to the PayeeBankName field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayeeBankName(v string) {
	o.PayeeBankName = &v
}

// GetPayeePhone returns the PayeePhone field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeePhone() string {
	if o == nil || IsNil(o.PayeePhone) {
		var ret string
		return ret
	}
	return *o.PayeePhone
}

// GetPayeePhoneOk returns a tuple with the PayeePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.PayeePhone) {
		return nil, false
	}
	return o.PayeePhone, true
}

// HasPayeePhone returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayeePhone() bool {
	if o != nil && !IsNil(o.PayeePhone) {
		return true
	}

	return false
}

// SetPayeePhone gets a reference to the given string and assigns it to the PayeePhone field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayeePhone(v string) {
	o.PayeePhone = &v
}

// GetPayeeRegisterName returns the PayeeRegisterName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeRegisterName() string {
	if o == nil || IsNil(o.PayeeRegisterName) {
		var ret string
		return ret
	}
	return *o.PayeeRegisterName
}

// GetPayeeRegisterNameOk returns a tuple with the PayeeRegisterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeRegisterNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeRegisterName) {
		return nil, false
	}
	return o.PayeeRegisterName, true
}

// HasPayeeRegisterName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayeeRegisterName() bool {
	if o != nil && !IsNil(o.PayeeRegisterName) {
		return true
	}

	return false
}

// SetPayeeRegisterName gets a reference to the given string and assigns it to the PayeeRegisterName field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayeeRegisterName(v string) {
	o.PayeeRegisterName = &v
}

// GetPayeeTaxNo returns the PayeeTaxNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeTaxNo() string {
	if o == nil || IsNil(o.PayeeTaxNo) {
		var ret string
		return ret
	}
	return *o.PayeeTaxNo
}

// GetPayeeTaxNoOk returns a tuple with the PayeeTaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayeeTaxNoOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeTaxNo) {
		return nil, false
	}
	return o.PayeeTaxNo, true
}

// HasPayeeTaxNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayeeTaxNo() bool {
	if o != nil && !IsNil(o.PayeeTaxNo) {
		return true
	}

	return false
}

// SetPayeeTaxNo gets a reference to the given string and assigns it to the PayeeTaxNo field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayeeTaxNo(v string) {
	o.PayeeTaxNo = &v
}

// GetPayerAddress returns the PayerAddress field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerAddress() string {
	if o == nil || IsNil(o.PayerAddress) {
		var ret string
		return ret
	}
	return *o.PayerAddress
}

// GetPayerAddressOk returns a tuple with the PayerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PayerAddress) {
		return nil, false
	}
	return o.PayerAddress, true
}

// HasPayerAddress returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayerAddress() bool {
	if o != nil && !IsNil(o.PayerAddress) {
		return true
	}

	return false
}

// SetPayerAddress gets a reference to the given string and assigns it to the PayerAddress field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayerAddress(v string) {
	o.PayerAddress = &v
}

// GetPayerBankAccount returns the PayerBankAccount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerBankAccount() string {
	if o == nil || IsNil(o.PayerBankAccount) {
		var ret string
		return ret
	}
	return *o.PayerBankAccount
}

// GetPayerBankAccountOk returns a tuple with the PayerBankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.PayerBankAccount) {
		return nil, false
	}
	return o.PayerBankAccount, true
}

// HasPayerBankAccount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayerBankAccount() bool {
	if o != nil && !IsNil(o.PayerBankAccount) {
		return true
	}

	return false
}

// SetPayerBankAccount gets a reference to the given string and assigns it to the PayerBankAccount field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayerBankAccount(v string) {
	o.PayerBankAccount = &v
}

// GetPayerBankName returns the PayerBankName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerBankName() string {
	if o == nil || IsNil(o.PayerBankName) {
		var ret string
		return ret
	}
	return *o.PayerBankName
}

// GetPayerBankNameOk returns a tuple with the PayerBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayerBankName) {
		return nil, false
	}
	return o.PayerBankName, true
}

// HasPayerBankName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayerBankName() bool {
	if o != nil && !IsNil(o.PayerBankName) {
		return true
	}

	return false
}

// SetPayerBankName gets a reference to the given string and assigns it to the PayerBankName field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayerBankName(v string) {
	o.PayerBankName = &v
}

// GetPayerName returns the PayerName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerName() string {
	if o == nil || IsNil(o.PayerName) {
		var ret string
		return ret
	}
	return *o.PayerName
}

// GetPayerNameOk returns a tuple with the PayerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayerName) {
		return nil, false
	}
	return o.PayerName, true
}

// HasPayerName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayerName() bool {
	if o != nil && !IsNil(o.PayerName) {
		return true
	}

	return false
}

// SetPayerName gets a reference to the given string and assigns it to the PayerName field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayerName(v string) {
	o.PayerName = &v
}

// GetPayerPhone returns the PayerPhone field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerPhone() string {
	if o == nil || IsNil(o.PayerPhone) {
		var ret string
		return ret
	}
	return *o.PayerPhone
}

// GetPayerPhoneOk returns a tuple with the PayerPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.PayerPhone) {
		return nil, false
	}
	return o.PayerPhone, true
}

// HasPayerPhone returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayerPhone() bool {
	if o != nil && !IsNil(o.PayerPhone) {
		return true
	}

	return false
}

// SetPayerPhone gets a reference to the given string and assigns it to the PayerPhone field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayerPhone(v string) {
	o.PayerPhone = &v
}

// GetPayerTaxNo returns the PayerTaxNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerTaxNo() string {
	if o == nil || IsNil(o.PayerTaxNo) {
		var ret string
		return ret
	}
	return *o.PayerTaxNo
}

// GetPayerTaxNoOk returns a tuple with the PayerTaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetPayerTaxNoOk() (*string, bool) {
	if o == nil || IsNil(o.PayerTaxNo) {
		return nil, false
	}
	return o.PayerTaxNo, true
}

// HasPayerTaxNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasPayerTaxNo() bool {
	if o != nil && !IsNil(o.PayerTaxNo) {
		return true
	}

	return false
}

// SetPayerTaxNo gets a reference to the given string and assigns it to the PayerTaxNo field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetPayerTaxNo(v string) {
	o.PayerTaxNo = &v
}

// GetSumTaxAmount returns the SumTaxAmount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetSumTaxAmount() string {
	if o == nil || IsNil(o.SumTaxAmount) {
		var ret string
		return ret
	}
	return *o.SumTaxAmount
}

// GetSumTaxAmountOk returns a tuple with the SumTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetSumTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.SumTaxAmount) {
		return nil, false
	}
	return o.SumTaxAmount, true
}

// HasSumTaxAmount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasSumTaxAmount() bool {
	if o != nil && !IsNil(o.SumTaxAmount) {
		return true
	}

	return false
}

// SetSumTaxAmount gets a reference to the given string and assigns it to the SumTaxAmount field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetSumTaxAmount(v string) {
	o.SumTaxAmount = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *AlipayEbppInvoiceDetailOutputQueryResponseModel) SetTaxType(v string) {
	o.TaxType = &v
}

func (o AlipayEbppInvoiceDetailOutputQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceDetailOutputQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AntiFakeCode) {
		toSerialize["anti_fake_code"] = o.AntiFakeCode
	}
	if !IsNil(o.ExTaxAmount) {
		toSerialize["ex_tax_amount"] = o.ExTaxAmount
	}
	if !IsNil(o.ExpenseStatus) {
		toSerialize["expense_status"] = o.ExpenseStatus
	}
	if !IsNil(o.HasRisk) {
		toSerialize["has_risk"] = o.HasRisk
	}
	if !IsNil(o.HasUploadPdf) {
		toSerialize["has_upload_pdf"] = o.HasUploadPdf
	}
	if !IsNil(o.InvoiceAmount) {
		toSerialize["invoice_amount"] = o.InvoiceAmount
	}
	if !IsNil(o.InvoiceCode) {
		toSerialize["invoice_code"] = o.InvoiceCode
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoice_date"] = o.InvoiceDate
	}
	if !IsNil(o.InvoiceImgUrl) {
		toSerialize["invoice_img_url"] = o.InvoiceImgUrl
	}
	if !IsNil(o.InvoiceItemContentList) {
		toSerialize["invoice_item_content_list"] = o.InvoiceItemContentList
	}
	if !IsNil(o.InvoiceNo) {
		toSerialize["invoice_no"] = o.InvoiceNo
	}
	if !IsNil(o.InvoiceStatus) {
		toSerialize["invoice_status"] = o.InvoiceStatus
	}
	if !IsNil(o.InvoiceType) {
		toSerialize["invoice_type"] = o.InvoiceType
	}
	if !IsNil(o.PayeeAddress) {
		toSerialize["payee_address"] = o.PayeeAddress
	}
	if !IsNil(o.PayeeBankAccount) {
		toSerialize["payee_bank_account"] = o.PayeeBankAccount
	}
	if !IsNil(o.PayeeBankName) {
		toSerialize["payee_bank_name"] = o.PayeeBankName
	}
	if !IsNil(o.PayeePhone) {
		toSerialize["payee_phone"] = o.PayeePhone
	}
	if !IsNil(o.PayeeRegisterName) {
		toSerialize["payee_register_name"] = o.PayeeRegisterName
	}
	if !IsNil(o.PayeeTaxNo) {
		toSerialize["payee_tax_no"] = o.PayeeTaxNo
	}
	if !IsNil(o.PayerAddress) {
		toSerialize["payer_address"] = o.PayerAddress
	}
	if !IsNil(o.PayerBankAccount) {
		toSerialize["payer_bank_account"] = o.PayerBankAccount
	}
	if !IsNil(o.PayerBankName) {
		toSerialize["payer_bank_name"] = o.PayerBankName
	}
	if !IsNil(o.PayerName) {
		toSerialize["payer_name"] = o.PayerName
	}
	if !IsNil(o.PayerPhone) {
		toSerialize["payer_phone"] = o.PayerPhone
	}
	if !IsNil(o.PayerTaxNo) {
		toSerialize["payer_tax_no"] = o.PayerTaxNo
	}
	if !IsNil(o.SumTaxAmount) {
		toSerialize["sum_tax_amount"] = o.SumTaxAmount
	}
	if !IsNil(o.TaxType) {
		toSerialize["tax_type"] = o.TaxType
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceDetailOutputQueryResponseModel struct {
	value *AlipayEbppInvoiceDetailOutputQueryResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryResponseModel) Get() *AlipayEbppInvoiceDetailOutputQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryResponseModel) Set(val *AlipayEbppInvoiceDetailOutputQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceDetailOutputQueryResponseModel(val *AlipayEbppInvoiceDetailOutputQueryResponseModel) *NullableAlipayEbppInvoiceDetailOutputQueryResponseModel {
	return &NullableAlipayEbppInvoiceDetailOutputQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


