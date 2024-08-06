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

// checks if the ArInvoiceReceiptOpenApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArInvoiceReceiptOpenApiResponse{}

// ArInvoiceReceiptOpenApiResponse struct for ArInvoiceReceiptOpenApiResponse
type ArInvoiceReceiptOpenApiResponse struct {
	// 合约号
	ArrangementNo *string `json:"arrangement_no,omitempty"`
	// 可开票单据主键ID
	Id *string `json:"id,omitempty"`
	// 机构ID
	InstId *string `json:"inst_id,omitempty"`
	// 开票时间 格式：yyyymm
	InvDt *string `json:"inv_dt,omitempty"`
	// 开票模式  01：实收开票， 02：应收开票， 03：累计实收开票
	InvMode *string `json:"inv_mode,omitempty"`
	InvoiceAmt *MultiCurrencyMoneyOpenApi `json:"invoice_amt,omitempty"`
	InvoicedAmt *MultiCurrencyMoneyOpenApi `json:"invoiced_amt,omitempty"`
	// 结算ip_id
	IpId *string `json:"ip_id,omitempty"`
	// 结算对象ip_role_id
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// 最后修改人
	LastModer *string `json:"last_moder,omitempty"`
	LinkInvoiceAmt *MultiCurrencyMoneyOpenApi `json:"link_invoice_amt,omitempty"`
	// 外部单据号，对应开票记录的月账单号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 可开票单据来源，01：主站，02：芝麻，03：金融云，04：微贷
	OutBizType *string `json:"out_biz_type,omitempty"`
	// 付款方式，1：资金，5：走量
	PayWay *string `json:"pay_way,omitempty"`
	// 产品码
	ProdCode *string `json:"prod_code,omitempty"`
	// 结算类型 ，01：实时，02：预收，03：后收，04：周期性，05：按日汇总，09：延期结算
	SettleType *string `json:"settle_type,omitempty"`
	// 开票金额消耗状态：01未开票，02部分开票，03：已开票
	Status *string `json:"status,omitempty"`
	// 税率
	TaxRate *int32 `json:"tax_rate,omitempty"`
	// 税收类型01：增值税，02：营业税
	TaxType *string `json:"tax_type,omitempty"`
	// 租户ID
	TntInstId *string `json:"tnt_inst_id,omitempty"`
	// 类型，1：应收，2：返点
	Type *string `json:"type,omitempty"`
}

// NewArInvoiceReceiptOpenApiResponse instantiates a new ArInvoiceReceiptOpenApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArInvoiceReceiptOpenApiResponse() *ArInvoiceReceiptOpenApiResponse {
	this := ArInvoiceReceiptOpenApiResponse{}
	return &this
}

// NewArInvoiceReceiptOpenApiResponseWithDefaults instantiates a new ArInvoiceReceiptOpenApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArInvoiceReceiptOpenApiResponseWithDefaults() *ArInvoiceReceiptOpenApiResponse {
	this := ArInvoiceReceiptOpenApiResponse{}
	return &this
}

// GetArrangementNo returns the ArrangementNo field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetArrangementNo() string {
	if o == nil || IsNil(o.ArrangementNo) {
		var ret string
		return ret
	}
	return *o.ArrangementNo
}

// GetArrangementNoOk returns a tuple with the ArrangementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetArrangementNoOk() (*string, bool) {
	if o == nil || IsNil(o.ArrangementNo) {
		return nil, false
	}
	return o.ArrangementNo, true
}

// HasArrangementNo returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasArrangementNo() bool {
	if o != nil && !IsNil(o.ArrangementNo) {
		return true
	}

	return false
}

// SetArrangementNo gets a reference to the given string and assigns it to the ArrangementNo field.
func (o *ArInvoiceReceiptOpenApiResponse) SetArrangementNo(v string) {
	o.ArrangementNo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArInvoiceReceiptOpenApiResponse) SetId(v string) {
	o.Id = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *ArInvoiceReceiptOpenApiResponse) SetInstId(v string) {
	o.InstId = &v
}

// GetInvDt returns the InvDt field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvDt() string {
	if o == nil || IsNil(o.InvDt) {
		var ret string
		return ret
	}
	return *o.InvDt
}

// GetInvDtOk returns a tuple with the InvDt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvDtOk() (*string, bool) {
	if o == nil || IsNil(o.InvDt) {
		return nil, false
	}
	return o.InvDt, true
}

// HasInvDt returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasInvDt() bool {
	if o != nil && !IsNil(o.InvDt) {
		return true
	}

	return false
}

// SetInvDt gets a reference to the given string and assigns it to the InvDt field.
func (o *ArInvoiceReceiptOpenApiResponse) SetInvDt(v string) {
	o.InvDt = &v
}

// GetInvMode returns the InvMode field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvMode() string {
	if o == nil || IsNil(o.InvMode) {
		var ret string
		return ret
	}
	return *o.InvMode
}

// GetInvModeOk returns a tuple with the InvMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvModeOk() (*string, bool) {
	if o == nil || IsNil(o.InvMode) {
		return nil, false
	}
	return o.InvMode, true
}

// HasInvMode returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasInvMode() bool {
	if o != nil && !IsNil(o.InvMode) {
		return true
	}

	return false
}

// SetInvMode gets a reference to the given string and assigns it to the InvMode field.
func (o *ArInvoiceReceiptOpenApiResponse) SetInvMode(v string) {
	o.InvMode = &v
}

// GetInvoiceAmt returns the InvoiceAmt field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvoiceAmt() MultiCurrencyMoneyOpenApi {
	if o == nil || IsNil(o.InvoiceAmt) {
		var ret MultiCurrencyMoneyOpenApi
		return ret
	}
	return *o.InvoiceAmt
}

// GetInvoiceAmtOk returns a tuple with the InvoiceAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvoiceAmtOk() (*MultiCurrencyMoneyOpenApi, bool) {
	if o == nil || IsNil(o.InvoiceAmt) {
		return nil, false
	}
	return o.InvoiceAmt, true
}

// HasInvoiceAmt returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasInvoiceAmt() bool {
	if o != nil && !IsNil(o.InvoiceAmt) {
		return true
	}

	return false
}

// SetInvoiceAmt gets a reference to the given MultiCurrencyMoneyOpenApi and assigns it to the InvoiceAmt field.
func (o *ArInvoiceReceiptOpenApiResponse) SetInvoiceAmt(v MultiCurrencyMoneyOpenApi) {
	o.InvoiceAmt = &v
}

// GetInvoicedAmt returns the InvoicedAmt field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvoicedAmt() MultiCurrencyMoneyOpenApi {
	if o == nil || IsNil(o.InvoicedAmt) {
		var ret MultiCurrencyMoneyOpenApi
		return ret
	}
	return *o.InvoicedAmt
}

// GetInvoicedAmtOk returns a tuple with the InvoicedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetInvoicedAmtOk() (*MultiCurrencyMoneyOpenApi, bool) {
	if o == nil || IsNil(o.InvoicedAmt) {
		return nil, false
	}
	return o.InvoicedAmt, true
}

// HasInvoicedAmt returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasInvoicedAmt() bool {
	if o != nil && !IsNil(o.InvoicedAmt) {
		return true
	}

	return false
}

// SetInvoicedAmt gets a reference to the given MultiCurrencyMoneyOpenApi and assigns it to the InvoicedAmt field.
func (o *ArInvoiceReceiptOpenApiResponse) SetInvoicedAmt(v MultiCurrencyMoneyOpenApi) {
	o.InvoicedAmt = &v
}

// GetIpId returns the IpId field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetIpId() string {
	if o == nil || IsNil(o.IpId) {
		var ret string
		return ret
	}
	return *o.IpId
}

// GetIpIdOk returns a tuple with the IpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetIpIdOk() (*string, bool) {
	if o == nil || IsNil(o.IpId) {
		return nil, false
	}
	return o.IpId, true
}

// HasIpId returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasIpId() bool {
	if o != nil && !IsNil(o.IpId) {
		return true
	}

	return false
}

// SetIpId gets a reference to the given string and assigns it to the IpId field.
func (o *ArInvoiceReceiptOpenApiResponse) SetIpId(v string) {
	o.IpId = &v
}

// GetIpRoleId returns the IpRoleId field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetIpRoleId() string {
	if o == nil || IsNil(o.IpRoleId) {
		var ret string
		return ret
	}
	return *o.IpRoleId
}

// GetIpRoleIdOk returns a tuple with the IpRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetIpRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IpRoleId) {
		return nil, false
	}
	return o.IpRoleId, true
}

// HasIpRoleId returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasIpRoleId() bool {
	if o != nil && !IsNil(o.IpRoleId) {
		return true
	}

	return false
}

// SetIpRoleId gets a reference to the given string and assigns it to the IpRoleId field.
func (o *ArInvoiceReceiptOpenApiResponse) SetIpRoleId(v string) {
	o.IpRoleId = &v
}

// GetLastModer returns the LastModer field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetLastModer() string {
	if o == nil || IsNil(o.LastModer) {
		var ret string
		return ret
	}
	return *o.LastModer
}

// GetLastModerOk returns a tuple with the LastModer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetLastModerOk() (*string, bool) {
	if o == nil || IsNil(o.LastModer) {
		return nil, false
	}
	return o.LastModer, true
}

// HasLastModer returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasLastModer() bool {
	if o != nil && !IsNil(o.LastModer) {
		return true
	}

	return false
}

// SetLastModer gets a reference to the given string and assigns it to the LastModer field.
func (o *ArInvoiceReceiptOpenApiResponse) SetLastModer(v string) {
	o.LastModer = &v
}

// GetLinkInvoiceAmt returns the LinkInvoiceAmt field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetLinkInvoiceAmt() MultiCurrencyMoneyOpenApi {
	if o == nil || IsNil(o.LinkInvoiceAmt) {
		var ret MultiCurrencyMoneyOpenApi
		return ret
	}
	return *o.LinkInvoiceAmt
}

// GetLinkInvoiceAmtOk returns a tuple with the LinkInvoiceAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetLinkInvoiceAmtOk() (*MultiCurrencyMoneyOpenApi, bool) {
	if o == nil || IsNil(o.LinkInvoiceAmt) {
		return nil, false
	}
	return o.LinkInvoiceAmt, true
}

// HasLinkInvoiceAmt returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasLinkInvoiceAmt() bool {
	if o != nil && !IsNil(o.LinkInvoiceAmt) {
		return true
	}

	return false
}

// SetLinkInvoiceAmt gets a reference to the given MultiCurrencyMoneyOpenApi and assigns it to the LinkInvoiceAmt field.
func (o *ArInvoiceReceiptOpenApiResponse) SetLinkInvoiceAmt(v MultiCurrencyMoneyOpenApi) {
	o.LinkInvoiceAmt = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *ArInvoiceReceiptOpenApiResponse) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetOutBizType returns the OutBizType field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetOutBizType() string {
	if o == nil || IsNil(o.OutBizType) {
		var ret string
		return ret
	}
	return *o.OutBizType
}

// GetOutBizTypeOk returns a tuple with the OutBizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetOutBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizType) {
		return nil, false
	}
	return o.OutBizType, true
}

// HasOutBizType returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasOutBizType() bool {
	if o != nil && !IsNil(o.OutBizType) {
		return true
	}

	return false
}

// SetOutBizType gets a reference to the given string and assigns it to the OutBizType field.
func (o *ArInvoiceReceiptOpenApiResponse) SetOutBizType(v string) {
	o.OutBizType = &v
}

// GetPayWay returns the PayWay field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetPayWay() string {
	if o == nil || IsNil(o.PayWay) {
		var ret string
		return ret
	}
	return *o.PayWay
}

// GetPayWayOk returns a tuple with the PayWay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetPayWayOk() (*string, bool) {
	if o == nil || IsNil(o.PayWay) {
		return nil, false
	}
	return o.PayWay, true
}

// HasPayWay returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasPayWay() bool {
	if o != nil && !IsNil(o.PayWay) {
		return true
	}

	return false
}

// SetPayWay gets a reference to the given string and assigns it to the PayWay field.
func (o *ArInvoiceReceiptOpenApiResponse) SetPayWay(v string) {
	o.PayWay = &v
}

// GetProdCode returns the ProdCode field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetProdCode() string {
	if o == nil || IsNil(o.ProdCode) {
		var ret string
		return ret
	}
	return *o.ProdCode
}

// GetProdCodeOk returns a tuple with the ProdCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetProdCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProdCode) {
		return nil, false
	}
	return o.ProdCode, true
}

// HasProdCode returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasProdCode() bool {
	if o != nil && !IsNil(o.ProdCode) {
		return true
	}

	return false
}

// SetProdCode gets a reference to the given string and assigns it to the ProdCode field.
func (o *ArInvoiceReceiptOpenApiResponse) SetProdCode(v string) {
	o.ProdCode = &v
}

// GetSettleType returns the SettleType field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetSettleType() string {
	if o == nil || IsNil(o.SettleType) {
		var ret string
		return ret
	}
	return *o.SettleType
}

// GetSettleTypeOk returns a tuple with the SettleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetSettleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleType) {
		return nil, false
	}
	return o.SettleType, true
}

// HasSettleType returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasSettleType() bool {
	if o != nil && !IsNil(o.SettleType) {
		return true
	}

	return false
}

// SetSettleType gets a reference to the given string and assigns it to the SettleType field.
func (o *ArInvoiceReceiptOpenApiResponse) SetSettleType(v string) {
	o.SettleType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ArInvoiceReceiptOpenApiResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetTaxRate() int32 {
	if o == nil || IsNil(o.TaxRate) {
		var ret int32
		return ret
	}
	return *o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetTaxRateOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxRate) {
		return nil, false
	}
	return o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasTaxRate() bool {
	if o != nil && !IsNil(o.TaxRate) {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given int32 and assigns it to the TaxRate field.
func (o *ArInvoiceReceiptOpenApiResponse) SetTaxRate(v int32) {
	o.TaxRate = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *ArInvoiceReceiptOpenApiResponse) SetTaxType(v string) {
	o.TaxType = &v
}

// GetTntInstId returns the TntInstId field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetTntInstId() string {
	if o == nil || IsNil(o.TntInstId) {
		var ret string
		return ret
	}
	return *o.TntInstId
}

// GetTntInstIdOk returns a tuple with the TntInstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetTntInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.TntInstId) {
		return nil, false
	}
	return o.TntInstId, true
}

// HasTntInstId returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasTntInstId() bool {
	if o != nil && !IsNil(o.TntInstId) {
		return true
	}

	return false
}

// SetTntInstId gets a reference to the given string and assigns it to the TntInstId field.
func (o *ArInvoiceReceiptOpenApiResponse) SetTntInstId(v string) {
	o.TntInstId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ArInvoiceReceiptOpenApiResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArInvoiceReceiptOpenApiResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ArInvoiceReceiptOpenApiResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ArInvoiceReceiptOpenApiResponse) SetType(v string) {
	o.Type = &v
}

func (o ArInvoiceReceiptOpenApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArInvoiceReceiptOpenApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArrangementNo) {
		toSerialize["arrangement_no"] = o.ArrangementNo
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InstId) {
		toSerialize["inst_id"] = o.InstId
	}
	if !IsNil(o.InvDt) {
		toSerialize["inv_dt"] = o.InvDt
	}
	if !IsNil(o.InvMode) {
		toSerialize["inv_mode"] = o.InvMode
	}
	if !IsNil(o.InvoiceAmt) {
		toSerialize["invoice_amt"] = o.InvoiceAmt
	}
	if !IsNil(o.InvoicedAmt) {
		toSerialize["invoiced_amt"] = o.InvoicedAmt
	}
	if !IsNil(o.IpId) {
		toSerialize["ip_id"] = o.IpId
	}
	if !IsNil(o.IpRoleId) {
		toSerialize["ip_role_id"] = o.IpRoleId
	}
	if !IsNil(o.LastModer) {
		toSerialize["last_moder"] = o.LastModer
	}
	if !IsNil(o.LinkInvoiceAmt) {
		toSerialize["link_invoice_amt"] = o.LinkInvoiceAmt
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.OutBizType) {
		toSerialize["out_biz_type"] = o.OutBizType
	}
	if !IsNil(o.PayWay) {
		toSerialize["pay_way"] = o.PayWay
	}
	if !IsNil(o.ProdCode) {
		toSerialize["prod_code"] = o.ProdCode
	}
	if !IsNil(o.SettleType) {
		toSerialize["settle_type"] = o.SettleType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TaxRate) {
		toSerialize["tax_rate"] = o.TaxRate
	}
	if !IsNil(o.TaxType) {
		toSerialize["tax_type"] = o.TaxType
	}
	if !IsNil(o.TntInstId) {
		toSerialize["tnt_inst_id"] = o.TntInstId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableArInvoiceReceiptOpenApiResponse struct {
	value *ArInvoiceReceiptOpenApiResponse
	isSet bool
}

func (v NullableArInvoiceReceiptOpenApiResponse) Get() *ArInvoiceReceiptOpenApiResponse {
	return v.value
}

func (v *NullableArInvoiceReceiptOpenApiResponse) Set(val *ArInvoiceReceiptOpenApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArInvoiceReceiptOpenApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArInvoiceReceiptOpenApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArInvoiceReceiptOpenApiResponse(val *ArInvoiceReceiptOpenApiResponse) *NullableArInvoiceReceiptOpenApiResponse {
	return &NullableArInvoiceReceiptOpenApiResponse{value: val, isSet: true}
}

func (v NullableArInvoiceReceiptOpenApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArInvoiceReceiptOpenApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

