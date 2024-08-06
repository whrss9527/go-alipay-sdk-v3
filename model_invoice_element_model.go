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

// checks if the InvoiceElementModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceElementModel{}

// InvoiceElementModel struct for InvoiceElementModel
type InvoiceElementModel struct {
	// 发票报销状态  取值范围：  WAIT_EXPENSE－未报销  EXPENSE_PROCESSING－报销中  EXPENSE_FINISHED－已报销
	ExpenseStatus *string `json:"expense_status,omitempty"`
	// 扩展字段
	ExtendFields *string `json:"extend_fields,omitempty"`
	// 防伪校验码
	FakeCode *string `json:"fake_code,omitempty"`
	// 发票是否有pdf文件
	HasPdfFile *bool `json:"has_pdf_file,omitempty"`
	// 该发票可能存在异常，请核实后使用  true:无异常  false:存在异常
	HasRisk *bool `json:"has_risk,omitempty"`
	// 发票金额，含税，单位元
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// 发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// 开票日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// 发票pdf文件转换后jpg预览地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// 发票类型 可选值 PLAIN：增值税电子普通发票 SPECIAL：增值税专用发票 ALL_ELECTRONIC_GENERAL： \"电子发票（普通发票） ALL_ELECTRONIC_SPECIAL： \"电子发票（专用发票） PLAIN_INVOICE:增值税普通发票 PAPER_INVOICE:增值税普通发票（卷式） SALSE_INVOICE:机动车销售统一发票 财政电子票据：FINANCIAL_ELECTRONIC_BILL
	InvoiceKind *string `json:"invoice_kind,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// 发票状态　  取值范围  SUCCEED－正常蓝票  EXPIRED－已失效
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// 服务商联系方式
	IsvContact *string `json:"isv_contact,omitempty"`
	// 服务商名称
	IsvName *string `json:"isv_name,omitempty"`
	// logo地址
	LogoUrl *string `json:"logo_url,omitempty"`
	// 商户全称
	MName *string `json:"m_name,omitempty"`
	// 发票金额，不含税，单位元
	OutTaxAmount *string `json:"out_tax_amount,omitempty"`
	// 销方名称
	PayeeName *string `json:"payee_name,omitempty"`
	// 销方税号
	PayeeTaxNo *string `json:"payee_tax_no,omitempty"`
	// 购方名称
	PayerName *string `json:"payer_name,omitempty"`
	// 购方税号
	PayerTaxNo *string `json:"payer_tax_no,omitempty"`
	// PDF的下载链接
	PdfUrl *string `json:"pdf_url,omitempty"`
	// 表示发票来源，由发票回传方带入。例如：bz_gd，bz_ele，bz_tmall等
	Source *string `json:"source,omitempty"`
	// 该发票对应的交易
	TradeList []EinvTrade `json:"trade_list,omitempty"`
	// 交易匹配结果 match-匹配到 noMatched-未匹配到 notMatch-未做匹配
	TradeMatchResult *string `json:"trade_match_result,omitempty"`
}

// NewInvoiceElementModel instantiates a new InvoiceElementModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceElementModel() *InvoiceElementModel {
	this := InvoiceElementModel{}
	return &this
}

// NewInvoiceElementModelWithDefaults instantiates a new InvoiceElementModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceElementModelWithDefaults() *InvoiceElementModel {
	this := InvoiceElementModel{}
	return &this
}

// GetExpenseStatus returns the ExpenseStatus field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetExpenseStatus() string {
	if o == nil || IsNil(o.ExpenseStatus) {
		var ret string
		return ret
	}
	return *o.ExpenseStatus
}

// GetExpenseStatusOk returns a tuple with the ExpenseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetExpenseStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseStatus) {
		return nil, false
	}
	return o.ExpenseStatus, true
}

// HasExpenseStatus returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasExpenseStatus() bool {
	if o != nil && !IsNil(o.ExpenseStatus) {
		return true
	}

	return false
}

// SetExpenseStatus gets a reference to the given string and assigns it to the ExpenseStatus field.
func (o *InvoiceElementModel) SetExpenseStatus(v string) {
	o.ExpenseStatus = &v
}

// GetExtendFields returns the ExtendFields field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetExtendFields() string {
	if o == nil || IsNil(o.ExtendFields) {
		var ret string
		return ret
	}
	return *o.ExtendFields
}

// GetExtendFieldsOk returns a tuple with the ExtendFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetExtendFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendFields) {
		return nil, false
	}
	return o.ExtendFields, true
}

// HasExtendFields returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasExtendFields() bool {
	if o != nil && !IsNil(o.ExtendFields) {
		return true
	}

	return false
}

// SetExtendFields gets a reference to the given string and assigns it to the ExtendFields field.
func (o *InvoiceElementModel) SetExtendFields(v string) {
	o.ExtendFields = &v
}

// GetFakeCode returns the FakeCode field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetFakeCode() string {
	if o == nil || IsNil(o.FakeCode) {
		var ret string
		return ret
	}
	return *o.FakeCode
}

// GetFakeCodeOk returns a tuple with the FakeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetFakeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FakeCode) {
		return nil, false
	}
	return o.FakeCode, true
}

// HasFakeCode returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasFakeCode() bool {
	if o != nil && !IsNil(o.FakeCode) {
		return true
	}

	return false
}

// SetFakeCode gets a reference to the given string and assigns it to the FakeCode field.
func (o *InvoiceElementModel) SetFakeCode(v string) {
	o.FakeCode = &v
}

// GetHasPdfFile returns the HasPdfFile field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetHasPdfFile() bool {
	if o == nil || IsNil(o.HasPdfFile) {
		var ret bool
		return ret
	}
	return *o.HasPdfFile
}

// GetHasPdfFileOk returns a tuple with the HasPdfFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetHasPdfFileOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPdfFile) {
		return nil, false
	}
	return o.HasPdfFile, true
}

// HasHasPdfFile returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasHasPdfFile() bool {
	if o != nil && !IsNil(o.HasPdfFile) {
		return true
	}

	return false
}

// SetHasPdfFile gets a reference to the given bool and assigns it to the HasPdfFile field.
func (o *InvoiceElementModel) SetHasPdfFile(v bool) {
	o.HasPdfFile = &v
}

// GetHasRisk returns the HasRisk field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetHasRisk() bool {
	if o == nil || IsNil(o.HasRisk) {
		var ret bool
		return ret
	}
	return *o.HasRisk
}

// GetHasRiskOk returns a tuple with the HasRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetHasRiskOk() (*bool, bool) {
	if o == nil || IsNil(o.HasRisk) {
		return nil, false
	}
	return o.HasRisk, true
}

// HasHasRisk returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasHasRisk() bool {
	if o != nil && !IsNil(o.HasRisk) {
		return true
	}

	return false
}

// SetHasRisk gets a reference to the given bool and assigns it to the HasRisk field.
func (o *InvoiceElementModel) SetHasRisk(v bool) {
	o.HasRisk = &v
}

// GetInvoiceAmount returns the InvoiceAmount field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetInvoiceAmount() string {
	if o == nil || IsNil(o.InvoiceAmount) {
		var ret string
		return ret
	}
	return *o.InvoiceAmount
}

// GetInvoiceAmountOk returns a tuple with the InvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetInvoiceAmountOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceAmount) {
		return nil, false
	}
	return o.InvoiceAmount, true
}

// HasInvoiceAmount returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasInvoiceAmount() bool {
	if o != nil && !IsNil(o.InvoiceAmount) {
		return true
	}

	return false
}

// SetInvoiceAmount gets a reference to the given string and assigns it to the InvoiceAmount field.
func (o *InvoiceElementModel) SetInvoiceAmount(v string) {
	o.InvoiceAmount = &v
}

// GetInvoiceCode returns the InvoiceCode field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetInvoiceCode() string {
	if o == nil || IsNil(o.InvoiceCode) {
		var ret string
		return ret
	}
	return *o.InvoiceCode
}

// GetInvoiceCodeOk returns a tuple with the InvoiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetInvoiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCode) {
		return nil, false
	}
	return o.InvoiceCode, true
}

// HasInvoiceCode returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasInvoiceCode() bool {
	if o != nil && !IsNil(o.InvoiceCode) {
		return true
	}

	return false
}

// SetInvoiceCode gets a reference to the given string and assigns it to the InvoiceCode field.
func (o *InvoiceElementModel) SetInvoiceCode(v string) {
	o.InvoiceCode = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *InvoiceElementModel) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetInvoiceImgUrl returns the InvoiceImgUrl field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetInvoiceImgUrl() string {
	if o == nil || IsNil(o.InvoiceImgUrl) {
		var ret string
		return ret
	}
	return *o.InvoiceImgUrl
}

// GetInvoiceImgUrlOk returns a tuple with the InvoiceImgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetInvoiceImgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceImgUrl) {
		return nil, false
	}
	return o.InvoiceImgUrl, true
}

// HasInvoiceImgUrl returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasInvoiceImgUrl() bool {
	if o != nil && !IsNil(o.InvoiceImgUrl) {
		return true
	}

	return false
}

// SetInvoiceImgUrl gets a reference to the given string and assigns it to the InvoiceImgUrl field.
func (o *InvoiceElementModel) SetInvoiceImgUrl(v string) {
	o.InvoiceImgUrl = &v
}

// GetInvoiceKind returns the InvoiceKind field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetInvoiceKind() string {
	if o == nil || IsNil(o.InvoiceKind) {
		var ret string
		return ret
	}
	return *o.InvoiceKind
}

// GetInvoiceKindOk returns a tuple with the InvoiceKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetInvoiceKindOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceKind) {
		return nil, false
	}
	return o.InvoiceKind, true
}

// HasInvoiceKind returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasInvoiceKind() bool {
	if o != nil && !IsNil(o.InvoiceKind) {
		return true
	}

	return false
}

// SetInvoiceKind gets a reference to the given string and assigns it to the InvoiceKind field.
func (o *InvoiceElementModel) SetInvoiceKind(v string) {
	o.InvoiceKind = &v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetInvoiceNo() string {
	if o == nil || IsNil(o.InvoiceNo) {
		var ret string
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetInvoiceNoOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNo) {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasInvoiceNo() bool {
	if o != nil && !IsNil(o.InvoiceNo) {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given string and assigns it to the InvoiceNo field.
func (o *InvoiceElementModel) SetInvoiceNo(v string) {
	o.InvoiceNo = &v
}

// GetInvoiceStatus returns the InvoiceStatus field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetInvoiceStatus() string {
	if o == nil || IsNil(o.InvoiceStatus) {
		var ret string
		return ret
	}
	return *o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetInvoiceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceStatus) {
		return nil, false
	}
	return o.InvoiceStatus, true
}

// HasInvoiceStatus returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasInvoiceStatus() bool {
	if o != nil && !IsNil(o.InvoiceStatus) {
		return true
	}

	return false
}

// SetInvoiceStatus gets a reference to the given string and assigns it to the InvoiceStatus field.
func (o *InvoiceElementModel) SetInvoiceStatus(v string) {
	o.InvoiceStatus = &v
}

// GetIsvContact returns the IsvContact field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetIsvContact() string {
	if o == nil || IsNil(o.IsvContact) {
		var ret string
		return ret
	}
	return *o.IsvContact
}

// GetIsvContactOk returns a tuple with the IsvContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetIsvContactOk() (*string, bool) {
	if o == nil || IsNil(o.IsvContact) {
		return nil, false
	}
	return o.IsvContact, true
}

// HasIsvContact returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasIsvContact() bool {
	if o != nil && !IsNil(o.IsvContact) {
		return true
	}

	return false
}

// SetIsvContact gets a reference to the given string and assigns it to the IsvContact field.
func (o *InvoiceElementModel) SetIsvContact(v string) {
	o.IsvContact = &v
}

// GetIsvName returns the IsvName field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetIsvName() string {
	if o == nil || IsNil(o.IsvName) {
		var ret string
		return ret
	}
	return *o.IsvName
}

// GetIsvNameOk returns a tuple with the IsvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetIsvNameOk() (*string, bool) {
	if o == nil || IsNil(o.IsvName) {
		return nil, false
	}
	return o.IsvName, true
}

// HasIsvName returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasIsvName() bool {
	if o != nil && !IsNil(o.IsvName) {
		return true
	}

	return false
}

// SetIsvName gets a reference to the given string and assigns it to the IsvName field.
func (o *InvoiceElementModel) SetIsvName(v string) {
	o.IsvName = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *InvoiceElementModel) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetMName returns the MName field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetMName() string {
	if o == nil || IsNil(o.MName) {
		var ret string
		return ret
	}
	return *o.MName
}

// GetMNameOk returns a tuple with the MName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetMNameOk() (*string, bool) {
	if o == nil || IsNil(o.MName) {
		return nil, false
	}
	return o.MName, true
}

// HasMName returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasMName() bool {
	if o != nil && !IsNil(o.MName) {
		return true
	}

	return false
}

// SetMName gets a reference to the given string and assigns it to the MName field.
func (o *InvoiceElementModel) SetMName(v string) {
	o.MName = &v
}

// GetOutTaxAmount returns the OutTaxAmount field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetOutTaxAmount() string {
	if o == nil || IsNil(o.OutTaxAmount) {
		var ret string
		return ret
	}
	return *o.OutTaxAmount
}

// GetOutTaxAmountOk returns a tuple with the OutTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetOutTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.OutTaxAmount) {
		return nil, false
	}
	return o.OutTaxAmount, true
}

// HasOutTaxAmount returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasOutTaxAmount() bool {
	if o != nil && !IsNil(o.OutTaxAmount) {
		return true
	}

	return false
}

// SetOutTaxAmount gets a reference to the given string and assigns it to the OutTaxAmount field.
func (o *InvoiceElementModel) SetOutTaxAmount(v string) {
	o.OutTaxAmount = &v
}

// GetPayeeName returns the PayeeName field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetPayeeName() string {
	if o == nil || IsNil(o.PayeeName) {
		var ret string
		return ret
	}
	return *o.PayeeName
}

// GetPayeeNameOk returns a tuple with the PayeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetPayeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeName) {
		return nil, false
	}
	return o.PayeeName, true
}

// HasPayeeName returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasPayeeName() bool {
	if o != nil && !IsNil(o.PayeeName) {
		return true
	}

	return false
}

// SetPayeeName gets a reference to the given string and assigns it to the PayeeName field.
func (o *InvoiceElementModel) SetPayeeName(v string) {
	o.PayeeName = &v
}

// GetPayeeTaxNo returns the PayeeTaxNo field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetPayeeTaxNo() string {
	if o == nil || IsNil(o.PayeeTaxNo) {
		var ret string
		return ret
	}
	return *o.PayeeTaxNo
}

// GetPayeeTaxNoOk returns a tuple with the PayeeTaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetPayeeTaxNoOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeTaxNo) {
		return nil, false
	}
	return o.PayeeTaxNo, true
}

// HasPayeeTaxNo returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasPayeeTaxNo() bool {
	if o != nil && !IsNil(o.PayeeTaxNo) {
		return true
	}

	return false
}

// SetPayeeTaxNo gets a reference to the given string and assigns it to the PayeeTaxNo field.
func (o *InvoiceElementModel) SetPayeeTaxNo(v string) {
	o.PayeeTaxNo = &v
}

// GetPayerName returns the PayerName field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetPayerName() string {
	if o == nil || IsNil(o.PayerName) {
		var ret string
		return ret
	}
	return *o.PayerName
}

// GetPayerNameOk returns a tuple with the PayerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetPayerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayerName) {
		return nil, false
	}
	return o.PayerName, true
}

// HasPayerName returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasPayerName() bool {
	if o != nil && !IsNil(o.PayerName) {
		return true
	}

	return false
}

// SetPayerName gets a reference to the given string and assigns it to the PayerName field.
func (o *InvoiceElementModel) SetPayerName(v string) {
	o.PayerName = &v
}

// GetPayerTaxNo returns the PayerTaxNo field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetPayerTaxNo() string {
	if o == nil || IsNil(o.PayerTaxNo) {
		var ret string
		return ret
	}
	return *o.PayerTaxNo
}

// GetPayerTaxNoOk returns a tuple with the PayerTaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetPayerTaxNoOk() (*string, bool) {
	if o == nil || IsNil(o.PayerTaxNo) {
		return nil, false
	}
	return o.PayerTaxNo, true
}

// HasPayerTaxNo returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasPayerTaxNo() bool {
	if o != nil && !IsNil(o.PayerTaxNo) {
		return true
	}

	return false
}

// SetPayerTaxNo gets a reference to the given string and assigns it to the PayerTaxNo field.
func (o *InvoiceElementModel) SetPayerTaxNo(v string) {
	o.PayerTaxNo = &v
}

// GetPdfUrl returns the PdfUrl field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetPdfUrl() string {
	if o == nil || IsNil(o.PdfUrl) {
		var ret string
		return ret
	}
	return *o.PdfUrl
}

// GetPdfUrlOk returns a tuple with the PdfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetPdfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PdfUrl) {
		return nil, false
	}
	return o.PdfUrl, true
}

// HasPdfUrl returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasPdfUrl() bool {
	if o != nil && !IsNil(o.PdfUrl) {
		return true
	}

	return false
}

// SetPdfUrl gets a reference to the given string and assigns it to the PdfUrl field.
func (o *InvoiceElementModel) SetPdfUrl(v string) {
	o.PdfUrl = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *InvoiceElementModel) SetSource(v string) {
	o.Source = &v
}

// GetTradeList returns the TradeList field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetTradeList() []EinvTrade {
	if o == nil || IsNil(o.TradeList) {
		var ret []EinvTrade
		return ret
	}
	return o.TradeList
}

// GetTradeListOk returns a tuple with the TradeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetTradeListOk() ([]EinvTrade, bool) {
	if o == nil || IsNil(o.TradeList) {
		return nil, false
	}
	return o.TradeList, true
}

// HasTradeList returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasTradeList() bool {
	if o != nil && !IsNil(o.TradeList) {
		return true
	}

	return false
}

// SetTradeList gets a reference to the given []EinvTrade and assigns it to the TradeList field.
func (o *InvoiceElementModel) SetTradeList(v []EinvTrade) {
	o.TradeList = v
}

// GetTradeMatchResult returns the TradeMatchResult field value if set, zero value otherwise.
func (o *InvoiceElementModel) GetTradeMatchResult() string {
	if o == nil || IsNil(o.TradeMatchResult) {
		var ret string
		return ret
	}
	return *o.TradeMatchResult
}

// GetTradeMatchResultOk returns a tuple with the TradeMatchResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceElementModel) GetTradeMatchResultOk() (*string, bool) {
	if o == nil || IsNil(o.TradeMatchResult) {
		return nil, false
	}
	return o.TradeMatchResult, true
}

// HasTradeMatchResult returns a boolean if a field has been set.
func (o *InvoiceElementModel) HasTradeMatchResult() bool {
	if o != nil && !IsNil(o.TradeMatchResult) {
		return true
	}

	return false
}

// SetTradeMatchResult gets a reference to the given string and assigns it to the TradeMatchResult field.
func (o *InvoiceElementModel) SetTradeMatchResult(v string) {
	o.TradeMatchResult = &v
}

func (o InvoiceElementModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceElementModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpenseStatus) {
		toSerialize["expense_status"] = o.ExpenseStatus
	}
	if !IsNil(o.ExtendFields) {
		toSerialize["extend_fields"] = o.ExtendFields
	}
	if !IsNil(o.FakeCode) {
		toSerialize["fake_code"] = o.FakeCode
	}
	if !IsNil(o.HasPdfFile) {
		toSerialize["has_pdf_file"] = o.HasPdfFile
	}
	if !IsNil(o.HasRisk) {
		toSerialize["has_risk"] = o.HasRisk
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
	if !IsNil(o.InvoiceKind) {
		toSerialize["invoice_kind"] = o.InvoiceKind
	}
	if !IsNil(o.InvoiceNo) {
		toSerialize["invoice_no"] = o.InvoiceNo
	}
	if !IsNil(o.InvoiceStatus) {
		toSerialize["invoice_status"] = o.InvoiceStatus
	}
	if !IsNil(o.IsvContact) {
		toSerialize["isv_contact"] = o.IsvContact
	}
	if !IsNil(o.IsvName) {
		toSerialize["isv_name"] = o.IsvName
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.MName) {
		toSerialize["m_name"] = o.MName
	}
	if !IsNil(o.OutTaxAmount) {
		toSerialize["out_tax_amount"] = o.OutTaxAmount
	}
	if !IsNil(o.PayeeName) {
		toSerialize["payee_name"] = o.PayeeName
	}
	if !IsNil(o.PayeeTaxNo) {
		toSerialize["payee_tax_no"] = o.PayeeTaxNo
	}
	if !IsNil(o.PayerName) {
		toSerialize["payer_name"] = o.PayerName
	}
	if !IsNil(o.PayerTaxNo) {
		toSerialize["payer_tax_no"] = o.PayerTaxNo
	}
	if !IsNil(o.PdfUrl) {
		toSerialize["pdf_url"] = o.PdfUrl
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.TradeList) {
		toSerialize["trade_list"] = o.TradeList
	}
	if !IsNil(o.TradeMatchResult) {
		toSerialize["trade_match_result"] = o.TradeMatchResult
	}
	return toSerialize, nil
}

type NullableInvoiceElementModel struct {
	value *InvoiceElementModel
	isSet bool
}

func (v NullableInvoiceElementModel) Get() *InvoiceElementModel {
	return v.value
}

func (v *NullableInvoiceElementModel) Set(val *InvoiceElementModel) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceElementModel) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceElementModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceElementModel(val *InvoiceElementModel) *NullableInvoiceElementModel {
	return &NullableInvoiceElementModel{value: val, isSet: true}
}

func (v NullableInvoiceElementModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceElementModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
