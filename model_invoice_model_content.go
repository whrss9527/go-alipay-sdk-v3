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

// checks if the InvoiceModelContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceModelContent{}

// InvoiceModelContent struct for InvoiceModelContent
type InvoiceModelContent struct {
	// 支付宝端的申请id。如果在开票过程中，是通过支付宝提交的申请到机构端，支付宝会带上开票申请在支付宝生成的申请id，机构在回传发票的时候只需要回传这个申请id，不用获取用户的uid，支付宝可以根据申请id将发票归集到对应的用户名下
	ApplyId *string `json:"apply_id,omitempty"`
	// key=value，每组键值对以回车分割
	ExtendFields *string `json:"extend_fields,omitempty"`
	// 下载的发票文件类型，可选值： pdf（发票原文件） ofd（发票原文件） jpg（发票原文件缩略图）
	FileDownloadType *string `json:"file_download_type,omitempty"`
	// 发票原文件下载地址 1.当tax_type=PLAIN、ALL_ELECTRONIC_GENERAL或ALL_ELECTRONIC_SPECIAL时， file_download_url必传 且file_download_type取值范围为pdf或ofd； 2.当tax_type=SPECIAL时， file_download_url必传 file_download_type可以传入pdf，ofd，jpg 3.当其他票种时，file_download_url可以不传
	FileDownloadUrl *string `json:"file_download_url,omitempty"`
	// 财政电子票据子类型，当tax_type=FINANCIAL_ELECTRONIC_BILL时要求必填 可选值如下： 01:非税收入通用票据  02:非税收入专用票据 03:非税收入一般缴款书 04:资金往来结算票据 05:公益事业捐赠票据 06:医疗收费票据 07:社会团体会费票据 08:社会保险基金票据 09:工会经费收入票据 99:其他财政票据
	FinancialElectronicType *string `json:"financial_electronic_type,omitempty"`
	// 发票金额，大于0且精确到小数点两位，以元为单位  需要传入税价合计金额
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// 发票代码，为国税局生成的唯一值，全电票时为空，其他情况不可为空
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// 发票内容项
	InvoiceContent []InvoiceItemContent `json:"invoice_content,omitempty"`
	// 发票日期，用户填写，目前精确到日
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// 发票防伪码
	InvoiceFakeCode *string `json:"invoice_fake_code,omitempty"`
	// 原始发票PDF/OFD文件流
	InvoiceFileData *string `json:"invoice_file_data,omitempty"`
	// 发票原始文件jpg文件地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// 发票号码，国税局生成的唯一号码，不可为空串；  使用时请注意，invoice_no+invoice_code唯一，不能重复
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// 发票开具操作人
	InvoiceOperator *string `json:"invoice_operator,omitempty"`
	InvoiceTitle *InvoiceTitleModel `json:"invoice_title,omitempty"`
	// 发票类型，按照可选值只传入英文部分，该字段严格要求大小写  可选值:  blue（蓝票）  red（红票）
	InvoiceType *string `json:"invoice_type,omitempty"`
	// 支付宝用户id，支付宝端的申请id存在的时候也不需要传，其他情况下，当同步的是蓝票时，必传，红票时不需传。
	OpenId *string `json:"open_id,omitempty"`
	// 仅用于同步红票，原始蓝票发票代码，同步红票时必传（全电票时为空）
	OriginalBlueInvoiceCode *string `json:"original_blue_invoice_code,omitempty"`
	// 仅用于同步红票，原始蓝票发票号码，同步红票时必传
	OriginalBlueInvoiceNo *string `json:"original_blue_invoice_no,omitempty"`
	// 商户交易流水号，不可为空串;  传入红票时请注意，此字段的值要和蓝票保持一致
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 商户唯一开票申请业务流水号，同一个isv下不能重复
	OutInvoiceId *string `json:"out_invoice_id,omitempty"`
	// 开票单位地址
	RegisterAddress *string `json:"register_address,omitempty"`
	// 开票单位开户行账号
	RegisterBankAccount *string `json:"register_bank_account,omitempty"`
	// 开票单位开户行名称
	RegisterBankName *string `json:"register_bank_name,omitempty"`
	// 开票单位
	RegisterName *string `json:"register_name,omitempty"`
	// 纳税人识别号，不可为空串
	RegisterNo *string `json:"register_no,omitempty"`
	// 开票人电话，支持座机和手机两种格式
	RegisterPhoneNo *string `json:"register_phone_no,omitempty"`
	// 价税合计
	SumAmount *string `json:"sum_amount,omitempty"`
	// 税额
	TaxAmount *string `json:"tax_amount,omitempty"`
	// 税种 可选值： PLAIN：增值税电子普通发票 ALL_ELECTRONIC_GENERAL：电子发票（普通发票） ALL_ELECTRONIC_SPECIAL：电子发票（增值税专用发票） SPECIAL：增值税专用发票 PLAIN_INVOICE:增值税普通发票 FINANCIAL_ELECTRONIC_BILL:财政电子票据
	TaxType *string `json:"tax_type,omitempty"`
	// 支付宝用户id，支付宝端的申请id存在的时候也不需要传，其他情况下，当同步的是蓝票时，必传，红票时不需传。
	UserId *string `json:"user_id,omitempty"`
}

// NewInvoiceModelContent instantiates a new InvoiceModelContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceModelContent() *InvoiceModelContent {
	this := InvoiceModelContent{}
	return &this
}

// NewInvoiceModelContentWithDefaults instantiates a new InvoiceModelContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceModelContentWithDefaults() *InvoiceModelContent {
	this := InvoiceModelContent{}
	return &this
}

// GetApplyId returns the ApplyId field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetApplyId() string {
	if o == nil || IsNil(o.ApplyId) {
		var ret string
		return ret
	}
	return *o.ApplyId
}

// GetApplyIdOk returns a tuple with the ApplyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetApplyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyId) {
		return nil, false
	}
	return o.ApplyId, true
}

// HasApplyId returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasApplyId() bool {
	if o != nil && !IsNil(o.ApplyId) {
		return true
	}

	return false
}

// SetApplyId gets a reference to the given string and assigns it to the ApplyId field.
func (o *InvoiceModelContent) SetApplyId(v string) {
	o.ApplyId = &v
}

// GetExtendFields returns the ExtendFields field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetExtendFields() string {
	if o == nil || IsNil(o.ExtendFields) {
		var ret string
		return ret
	}
	return *o.ExtendFields
}

// GetExtendFieldsOk returns a tuple with the ExtendFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetExtendFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendFields) {
		return nil, false
	}
	return o.ExtendFields, true
}

// HasExtendFields returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasExtendFields() bool {
	if o != nil && !IsNil(o.ExtendFields) {
		return true
	}

	return false
}

// SetExtendFields gets a reference to the given string and assigns it to the ExtendFields field.
func (o *InvoiceModelContent) SetExtendFields(v string) {
	o.ExtendFields = &v
}

// GetFileDownloadType returns the FileDownloadType field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetFileDownloadType() string {
	if o == nil || IsNil(o.FileDownloadType) {
		var ret string
		return ret
	}
	return *o.FileDownloadType
}

// GetFileDownloadTypeOk returns a tuple with the FileDownloadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetFileDownloadTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileDownloadType) {
		return nil, false
	}
	return o.FileDownloadType, true
}

// HasFileDownloadType returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasFileDownloadType() bool {
	if o != nil && !IsNil(o.FileDownloadType) {
		return true
	}

	return false
}

// SetFileDownloadType gets a reference to the given string and assigns it to the FileDownloadType field.
func (o *InvoiceModelContent) SetFileDownloadType(v string) {
	o.FileDownloadType = &v
}

// GetFileDownloadUrl returns the FileDownloadUrl field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetFileDownloadUrl() string {
	if o == nil || IsNil(o.FileDownloadUrl) {
		var ret string
		return ret
	}
	return *o.FileDownloadUrl
}

// GetFileDownloadUrlOk returns a tuple with the FileDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetFileDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileDownloadUrl) {
		return nil, false
	}
	return o.FileDownloadUrl, true
}

// HasFileDownloadUrl returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasFileDownloadUrl() bool {
	if o != nil && !IsNil(o.FileDownloadUrl) {
		return true
	}

	return false
}

// SetFileDownloadUrl gets a reference to the given string and assigns it to the FileDownloadUrl field.
func (o *InvoiceModelContent) SetFileDownloadUrl(v string) {
	o.FileDownloadUrl = &v
}

// GetFinancialElectronicType returns the FinancialElectronicType field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetFinancialElectronicType() string {
	if o == nil || IsNil(o.FinancialElectronicType) {
		var ret string
		return ret
	}
	return *o.FinancialElectronicType
}

// GetFinancialElectronicTypeOk returns a tuple with the FinancialElectronicType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetFinancialElectronicTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FinancialElectronicType) {
		return nil, false
	}
	return o.FinancialElectronicType, true
}

// HasFinancialElectronicType returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasFinancialElectronicType() bool {
	if o != nil && !IsNil(o.FinancialElectronicType) {
		return true
	}

	return false
}

// SetFinancialElectronicType gets a reference to the given string and assigns it to the FinancialElectronicType field.
func (o *InvoiceModelContent) SetFinancialElectronicType(v string) {
	o.FinancialElectronicType = &v
}

// GetInvoiceAmount returns the InvoiceAmount field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceAmount() string {
	if o == nil || IsNil(o.InvoiceAmount) {
		var ret string
		return ret
	}
	return *o.InvoiceAmount
}

// GetInvoiceAmountOk returns a tuple with the InvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceAmountOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceAmount) {
		return nil, false
	}
	return o.InvoiceAmount, true
}

// HasInvoiceAmount returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceAmount() bool {
	if o != nil && !IsNil(o.InvoiceAmount) {
		return true
	}

	return false
}

// SetInvoiceAmount gets a reference to the given string and assigns it to the InvoiceAmount field.
func (o *InvoiceModelContent) SetInvoiceAmount(v string) {
	o.InvoiceAmount = &v
}

// GetInvoiceCode returns the InvoiceCode field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceCode() string {
	if o == nil || IsNil(o.InvoiceCode) {
		var ret string
		return ret
	}
	return *o.InvoiceCode
}

// GetInvoiceCodeOk returns a tuple with the InvoiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCode) {
		return nil, false
	}
	return o.InvoiceCode, true
}

// HasInvoiceCode returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceCode() bool {
	if o != nil && !IsNil(o.InvoiceCode) {
		return true
	}

	return false
}

// SetInvoiceCode gets a reference to the given string and assigns it to the InvoiceCode field.
func (o *InvoiceModelContent) SetInvoiceCode(v string) {
	o.InvoiceCode = &v
}

// GetInvoiceContent returns the InvoiceContent field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceContent() []InvoiceItemContent {
	if o == nil || IsNil(o.InvoiceContent) {
		var ret []InvoiceItemContent
		return ret
	}
	return o.InvoiceContent
}

// GetInvoiceContentOk returns a tuple with the InvoiceContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceContentOk() ([]InvoiceItemContent, bool) {
	if o == nil || IsNil(o.InvoiceContent) {
		return nil, false
	}
	return o.InvoiceContent, true
}

// HasInvoiceContent returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceContent() bool {
	if o != nil && !IsNil(o.InvoiceContent) {
		return true
	}

	return false
}

// SetInvoiceContent gets a reference to the given []InvoiceItemContent and assigns it to the InvoiceContent field.
func (o *InvoiceModelContent) SetInvoiceContent(v []InvoiceItemContent) {
	o.InvoiceContent = v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *InvoiceModelContent) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetInvoiceFakeCode returns the InvoiceFakeCode field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceFakeCode() string {
	if o == nil || IsNil(o.InvoiceFakeCode) {
		var ret string
		return ret
	}
	return *o.InvoiceFakeCode
}

// GetInvoiceFakeCodeOk returns a tuple with the InvoiceFakeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceFakeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceFakeCode) {
		return nil, false
	}
	return o.InvoiceFakeCode, true
}

// HasInvoiceFakeCode returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceFakeCode() bool {
	if o != nil && !IsNil(o.InvoiceFakeCode) {
		return true
	}

	return false
}

// SetInvoiceFakeCode gets a reference to the given string and assigns it to the InvoiceFakeCode field.
func (o *InvoiceModelContent) SetInvoiceFakeCode(v string) {
	o.InvoiceFakeCode = &v
}

// GetInvoiceFileData returns the InvoiceFileData field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceFileData() string {
	if o == nil || IsNil(o.InvoiceFileData) {
		var ret string
		return ret
	}
	return *o.InvoiceFileData
}

// GetInvoiceFileDataOk returns a tuple with the InvoiceFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceFileDataOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceFileData) {
		return nil, false
	}
	return o.InvoiceFileData, true
}

// HasInvoiceFileData returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceFileData() bool {
	if o != nil && !IsNil(o.InvoiceFileData) {
		return true
	}

	return false
}

// SetInvoiceFileData gets a reference to the given string and assigns it to the InvoiceFileData field.
func (o *InvoiceModelContent) SetInvoiceFileData(v string) {
	o.InvoiceFileData = &v
}

// GetInvoiceImgUrl returns the InvoiceImgUrl field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceImgUrl() string {
	if o == nil || IsNil(o.InvoiceImgUrl) {
		var ret string
		return ret
	}
	return *o.InvoiceImgUrl
}

// GetInvoiceImgUrlOk returns a tuple with the InvoiceImgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceImgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceImgUrl) {
		return nil, false
	}
	return o.InvoiceImgUrl, true
}

// HasInvoiceImgUrl returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceImgUrl() bool {
	if o != nil && !IsNil(o.InvoiceImgUrl) {
		return true
	}

	return false
}

// SetInvoiceImgUrl gets a reference to the given string and assigns it to the InvoiceImgUrl field.
func (o *InvoiceModelContent) SetInvoiceImgUrl(v string) {
	o.InvoiceImgUrl = &v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceNo() string {
	if o == nil || IsNil(o.InvoiceNo) {
		var ret string
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceNoOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNo) {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceNo() bool {
	if o != nil && !IsNil(o.InvoiceNo) {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given string and assigns it to the InvoiceNo field.
func (o *InvoiceModelContent) SetInvoiceNo(v string) {
	o.InvoiceNo = &v
}

// GetInvoiceOperator returns the InvoiceOperator field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceOperator() string {
	if o == nil || IsNil(o.InvoiceOperator) {
		var ret string
		return ret
	}
	return *o.InvoiceOperator
}

// GetInvoiceOperatorOk returns a tuple with the InvoiceOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceOperator) {
		return nil, false
	}
	return o.InvoiceOperator, true
}

// HasInvoiceOperator returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceOperator() bool {
	if o != nil && !IsNil(o.InvoiceOperator) {
		return true
	}

	return false
}

// SetInvoiceOperator gets a reference to the given string and assigns it to the InvoiceOperator field.
func (o *InvoiceModelContent) SetInvoiceOperator(v string) {
	o.InvoiceOperator = &v
}

// GetInvoiceTitle returns the InvoiceTitle field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceTitle() InvoiceTitleModel {
	if o == nil || IsNil(o.InvoiceTitle) {
		var ret InvoiceTitleModel
		return ret
	}
	return *o.InvoiceTitle
}

// GetInvoiceTitleOk returns a tuple with the InvoiceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceTitleOk() (*InvoiceTitleModel, bool) {
	if o == nil || IsNil(o.InvoiceTitle) {
		return nil, false
	}
	return o.InvoiceTitle, true
}

// HasInvoiceTitle returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceTitle() bool {
	if o != nil && !IsNil(o.InvoiceTitle) {
		return true
	}

	return false
}

// SetInvoiceTitle gets a reference to the given InvoiceTitleModel and assigns it to the InvoiceTitle field.
func (o *InvoiceModelContent) SetInvoiceTitle(v InvoiceTitleModel) {
	o.InvoiceTitle = &v
}

// GetInvoiceType returns the InvoiceType field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetInvoiceType() string {
	if o == nil || IsNil(o.InvoiceType) {
		var ret string
		return ret
	}
	return *o.InvoiceType
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetInvoiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceType) {
		return nil, false
	}
	return o.InvoiceType, true
}

// HasInvoiceType returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasInvoiceType() bool {
	if o != nil && !IsNil(o.InvoiceType) {
		return true
	}

	return false
}

// SetInvoiceType gets a reference to the given string and assigns it to the InvoiceType field.
func (o *InvoiceModelContent) SetInvoiceType(v string) {
	o.InvoiceType = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *InvoiceModelContent) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOriginalBlueInvoiceCode returns the OriginalBlueInvoiceCode field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetOriginalBlueInvoiceCode() string {
	if o == nil || IsNil(o.OriginalBlueInvoiceCode) {
		var ret string
		return ret
	}
	return *o.OriginalBlueInvoiceCode
}

// GetOriginalBlueInvoiceCodeOk returns a tuple with the OriginalBlueInvoiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetOriginalBlueInvoiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalBlueInvoiceCode) {
		return nil, false
	}
	return o.OriginalBlueInvoiceCode, true
}

// HasOriginalBlueInvoiceCode returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasOriginalBlueInvoiceCode() bool {
	if o != nil && !IsNil(o.OriginalBlueInvoiceCode) {
		return true
	}

	return false
}

// SetOriginalBlueInvoiceCode gets a reference to the given string and assigns it to the OriginalBlueInvoiceCode field.
func (o *InvoiceModelContent) SetOriginalBlueInvoiceCode(v string) {
	o.OriginalBlueInvoiceCode = &v
}

// GetOriginalBlueInvoiceNo returns the OriginalBlueInvoiceNo field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetOriginalBlueInvoiceNo() string {
	if o == nil || IsNil(o.OriginalBlueInvoiceNo) {
		var ret string
		return ret
	}
	return *o.OriginalBlueInvoiceNo
}

// GetOriginalBlueInvoiceNoOk returns a tuple with the OriginalBlueInvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetOriginalBlueInvoiceNoOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalBlueInvoiceNo) {
		return nil, false
	}
	return o.OriginalBlueInvoiceNo, true
}

// HasOriginalBlueInvoiceNo returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasOriginalBlueInvoiceNo() bool {
	if o != nil && !IsNil(o.OriginalBlueInvoiceNo) {
		return true
	}

	return false
}

// SetOriginalBlueInvoiceNo gets a reference to the given string and assigns it to the OriginalBlueInvoiceNo field.
func (o *InvoiceModelContent) SetOriginalBlueInvoiceNo(v string) {
	o.OriginalBlueInvoiceNo = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *InvoiceModelContent) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetOutInvoiceId returns the OutInvoiceId field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetOutInvoiceId() string {
	if o == nil || IsNil(o.OutInvoiceId) {
		var ret string
		return ret
	}
	return *o.OutInvoiceId
}

// GetOutInvoiceIdOk returns a tuple with the OutInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetOutInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutInvoiceId) {
		return nil, false
	}
	return o.OutInvoiceId, true
}

// HasOutInvoiceId returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasOutInvoiceId() bool {
	if o != nil && !IsNil(o.OutInvoiceId) {
		return true
	}

	return false
}

// SetOutInvoiceId gets a reference to the given string and assigns it to the OutInvoiceId field.
func (o *InvoiceModelContent) SetOutInvoiceId(v string) {
	o.OutInvoiceId = &v
}

// GetRegisterAddress returns the RegisterAddress field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetRegisterAddress() string {
	if o == nil || IsNil(o.RegisterAddress) {
		var ret string
		return ret
	}
	return *o.RegisterAddress
}

// GetRegisterAddressOk returns a tuple with the RegisterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetRegisterAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RegisterAddress) {
		return nil, false
	}
	return o.RegisterAddress, true
}

// HasRegisterAddress returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasRegisterAddress() bool {
	if o != nil && !IsNil(o.RegisterAddress) {
		return true
	}

	return false
}

// SetRegisterAddress gets a reference to the given string and assigns it to the RegisterAddress field.
func (o *InvoiceModelContent) SetRegisterAddress(v string) {
	o.RegisterAddress = &v
}

// GetRegisterBankAccount returns the RegisterBankAccount field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetRegisterBankAccount() string {
	if o == nil || IsNil(o.RegisterBankAccount) {
		var ret string
		return ret
	}
	return *o.RegisterBankAccount
}

// GetRegisterBankAccountOk returns a tuple with the RegisterBankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetRegisterBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.RegisterBankAccount) {
		return nil, false
	}
	return o.RegisterBankAccount, true
}

// HasRegisterBankAccount returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasRegisterBankAccount() bool {
	if o != nil && !IsNil(o.RegisterBankAccount) {
		return true
	}

	return false
}

// SetRegisterBankAccount gets a reference to the given string and assigns it to the RegisterBankAccount field.
func (o *InvoiceModelContent) SetRegisterBankAccount(v string) {
	o.RegisterBankAccount = &v
}

// GetRegisterBankName returns the RegisterBankName field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetRegisterBankName() string {
	if o == nil || IsNil(o.RegisterBankName) {
		var ret string
		return ret
	}
	return *o.RegisterBankName
}

// GetRegisterBankNameOk returns a tuple with the RegisterBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetRegisterBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegisterBankName) {
		return nil, false
	}
	return o.RegisterBankName, true
}

// HasRegisterBankName returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasRegisterBankName() bool {
	if o != nil && !IsNil(o.RegisterBankName) {
		return true
	}

	return false
}

// SetRegisterBankName gets a reference to the given string and assigns it to the RegisterBankName field.
func (o *InvoiceModelContent) SetRegisterBankName(v string) {
	o.RegisterBankName = &v
}

// GetRegisterName returns the RegisterName field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetRegisterName() string {
	if o == nil || IsNil(o.RegisterName) {
		var ret string
		return ret
	}
	return *o.RegisterName
}

// GetRegisterNameOk returns a tuple with the RegisterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetRegisterNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegisterName) {
		return nil, false
	}
	return o.RegisterName, true
}

// HasRegisterName returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasRegisterName() bool {
	if o != nil && !IsNil(o.RegisterName) {
		return true
	}

	return false
}

// SetRegisterName gets a reference to the given string and assigns it to the RegisterName field.
func (o *InvoiceModelContent) SetRegisterName(v string) {
	o.RegisterName = &v
}

// GetRegisterNo returns the RegisterNo field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetRegisterNo() string {
	if o == nil || IsNil(o.RegisterNo) {
		var ret string
		return ret
	}
	return *o.RegisterNo
}

// GetRegisterNoOk returns a tuple with the RegisterNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetRegisterNoOk() (*string, bool) {
	if o == nil || IsNil(o.RegisterNo) {
		return nil, false
	}
	return o.RegisterNo, true
}

// HasRegisterNo returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasRegisterNo() bool {
	if o != nil && !IsNil(o.RegisterNo) {
		return true
	}

	return false
}

// SetRegisterNo gets a reference to the given string and assigns it to the RegisterNo field.
func (o *InvoiceModelContent) SetRegisterNo(v string) {
	o.RegisterNo = &v
}

// GetRegisterPhoneNo returns the RegisterPhoneNo field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetRegisterPhoneNo() string {
	if o == nil || IsNil(o.RegisterPhoneNo) {
		var ret string
		return ret
	}
	return *o.RegisterPhoneNo
}

// GetRegisterPhoneNoOk returns a tuple with the RegisterPhoneNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetRegisterPhoneNoOk() (*string, bool) {
	if o == nil || IsNil(o.RegisterPhoneNo) {
		return nil, false
	}
	return o.RegisterPhoneNo, true
}

// HasRegisterPhoneNo returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasRegisterPhoneNo() bool {
	if o != nil && !IsNil(o.RegisterPhoneNo) {
		return true
	}

	return false
}

// SetRegisterPhoneNo gets a reference to the given string and assigns it to the RegisterPhoneNo field.
func (o *InvoiceModelContent) SetRegisterPhoneNo(v string) {
	o.RegisterPhoneNo = &v
}

// GetSumAmount returns the SumAmount field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetSumAmount() string {
	if o == nil || IsNil(o.SumAmount) {
		var ret string
		return ret
	}
	return *o.SumAmount
}

// GetSumAmountOk returns a tuple with the SumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetSumAmountOk() (*string, bool) {
	if o == nil || IsNil(o.SumAmount) {
		return nil, false
	}
	return o.SumAmount, true
}

// HasSumAmount returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasSumAmount() bool {
	if o != nil && !IsNil(o.SumAmount) {
		return true
	}

	return false
}

// SetSumAmount gets a reference to the given string and assigns it to the SumAmount field.
func (o *InvoiceModelContent) SetSumAmount(v string) {
	o.SumAmount = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetTaxAmount() string {
	if o == nil || IsNil(o.TaxAmount) {
		var ret string
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given string and assigns it to the TaxAmount field.
func (o *InvoiceModelContent) SetTaxAmount(v string) {
	o.TaxAmount = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *InvoiceModelContent) SetTaxType(v string) {
	o.TaxType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *InvoiceModelContent) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceModelContent) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InvoiceModelContent) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *InvoiceModelContent) SetUserId(v string) {
	o.UserId = &v
}

func (o InvoiceModelContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceModelContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyId) {
		toSerialize["apply_id"] = o.ApplyId
	}
	if !IsNil(o.ExtendFields) {
		toSerialize["extend_fields"] = o.ExtendFields
	}
	if !IsNil(o.FileDownloadType) {
		toSerialize["file_download_type"] = o.FileDownloadType
	}
	if !IsNil(o.FileDownloadUrl) {
		toSerialize["file_download_url"] = o.FileDownloadUrl
	}
	if !IsNil(o.FinancialElectronicType) {
		toSerialize["financial_electronic_type"] = o.FinancialElectronicType
	}
	if !IsNil(o.InvoiceAmount) {
		toSerialize["invoice_amount"] = o.InvoiceAmount
	}
	if !IsNil(o.InvoiceCode) {
		toSerialize["invoice_code"] = o.InvoiceCode
	}
	if !IsNil(o.InvoiceContent) {
		toSerialize["invoice_content"] = o.InvoiceContent
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoice_date"] = o.InvoiceDate
	}
	if !IsNil(o.InvoiceFakeCode) {
		toSerialize["invoice_fake_code"] = o.InvoiceFakeCode
	}
	if !IsNil(o.InvoiceFileData) {
		toSerialize["invoice_file_data"] = o.InvoiceFileData
	}
	if !IsNil(o.InvoiceImgUrl) {
		toSerialize["invoice_img_url"] = o.InvoiceImgUrl
	}
	if !IsNil(o.InvoiceNo) {
		toSerialize["invoice_no"] = o.InvoiceNo
	}
	if !IsNil(o.InvoiceOperator) {
		toSerialize["invoice_operator"] = o.InvoiceOperator
	}
	if !IsNil(o.InvoiceTitle) {
		toSerialize["invoice_title"] = o.InvoiceTitle
	}
	if !IsNil(o.InvoiceType) {
		toSerialize["invoice_type"] = o.InvoiceType
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OriginalBlueInvoiceCode) {
		toSerialize["original_blue_invoice_code"] = o.OriginalBlueInvoiceCode
	}
	if !IsNil(o.OriginalBlueInvoiceNo) {
		toSerialize["original_blue_invoice_no"] = o.OriginalBlueInvoiceNo
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.OutInvoiceId) {
		toSerialize["out_invoice_id"] = o.OutInvoiceId
	}
	if !IsNil(o.RegisterAddress) {
		toSerialize["register_address"] = o.RegisterAddress
	}
	if !IsNil(o.RegisterBankAccount) {
		toSerialize["register_bank_account"] = o.RegisterBankAccount
	}
	if !IsNil(o.RegisterBankName) {
		toSerialize["register_bank_name"] = o.RegisterBankName
	}
	if !IsNil(o.RegisterName) {
		toSerialize["register_name"] = o.RegisterName
	}
	if !IsNil(o.RegisterNo) {
		toSerialize["register_no"] = o.RegisterNo
	}
	if !IsNil(o.RegisterPhoneNo) {
		toSerialize["register_phone_no"] = o.RegisterPhoneNo
	}
	if !IsNil(o.SumAmount) {
		toSerialize["sum_amount"] = o.SumAmount
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["tax_amount"] = o.TaxAmount
	}
	if !IsNil(o.TaxType) {
		toSerialize["tax_type"] = o.TaxType
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableInvoiceModelContent struct {
	value *InvoiceModelContent
	isSet bool
}

func (v NullableInvoiceModelContent) Get() *InvoiceModelContent {
	return v.value
}

func (v *NullableInvoiceModelContent) Set(val *InvoiceModelContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceModelContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceModelContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceModelContent(val *InvoiceModelContent) *NullableInvoiceModelContent {
	return &NullableInvoiceModelContent{value: val, isSet: true}
}

func (v NullableInvoiceModelContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceModelContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


