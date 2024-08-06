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

// checks if the AlipayBossFncInvoicereceiptBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayBossFncInvoicereceiptBatchqueryResponseModel{}

// AlipayBossFncInvoicereceiptBatchqueryResponseModel struct for AlipayBossFncInvoicereceiptBatchqueryResponseModel
type AlipayBossFncInvoicereceiptBatchqueryResponseModel struct {
	// 返回结果对象：可开票单据
	ResultSet []ArInvoiceReceiptOpenApiResponse `json:"result_set,omitempty"`
	TotalInvAmt *MultiCurrencyMoneyOpenApi `json:"total_inv_amt,omitempty"`
	TotalInvedAmt *MultiCurrencyMoneyOpenApi `json:"total_inved_amt,omitempty"`
	TotalLinkInvoiceAmt *MultiCurrencyMoneyOpenApi `json:"total_link_invoice_amt,omitempty"`
}

// NewAlipayBossFncInvoicereceiptBatchqueryResponseModel instantiates a new AlipayBossFncInvoicereceiptBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayBossFncInvoicereceiptBatchqueryResponseModel() *AlipayBossFncInvoicereceiptBatchqueryResponseModel {
	this := AlipayBossFncInvoicereceiptBatchqueryResponseModel{}
	return &this
}

// NewAlipayBossFncInvoicereceiptBatchqueryResponseModelWithDefaults instantiates a new AlipayBossFncInvoicereceiptBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayBossFncInvoicereceiptBatchqueryResponseModelWithDefaults() *AlipayBossFncInvoicereceiptBatchqueryResponseModel {
	this := AlipayBossFncInvoicereceiptBatchqueryResponseModel{}
	return &this
}

// GetResultSet returns the ResultSet field value if set, zero value otherwise.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetResultSet() []ArInvoiceReceiptOpenApiResponse {
	if o == nil || IsNil(o.ResultSet) {
		var ret []ArInvoiceReceiptOpenApiResponse
		return ret
	}
	return o.ResultSet
}

// GetResultSetOk returns a tuple with the ResultSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetResultSetOk() ([]ArInvoiceReceiptOpenApiResponse, bool) {
	if o == nil || IsNil(o.ResultSet) {
		return nil, false
	}
	return o.ResultSet, true
}

// HasResultSet returns a boolean if a field has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) HasResultSet() bool {
	if o != nil && !IsNil(o.ResultSet) {
		return true
	}

	return false
}

// SetResultSet gets a reference to the given []ArInvoiceReceiptOpenApiResponse and assigns it to the ResultSet field.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) SetResultSet(v []ArInvoiceReceiptOpenApiResponse) {
	o.ResultSet = v
}

// GetTotalInvAmt returns the TotalInvAmt field value if set, zero value otherwise.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetTotalInvAmt() MultiCurrencyMoneyOpenApi {
	if o == nil || IsNil(o.TotalInvAmt) {
		var ret MultiCurrencyMoneyOpenApi
		return ret
	}
	return *o.TotalInvAmt
}

// GetTotalInvAmtOk returns a tuple with the TotalInvAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetTotalInvAmtOk() (*MultiCurrencyMoneyOpenApi, bool) {
	if o == nil || IsNil(o.TotalInvAmt) {
		return nil, false
	}
	return o.TotalInvAmt, true
}

// HasTotalInvAmt returns a boolean if a field has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) HasTotalInvAmt() bool {
	if o != nil && !IsNil(o.TotalInvAmt) {
		return true
	}

	return false
}

// SetTotalInvAmt gets a reference to the given MultiCurrencyMoneyOpenApi and assigns it to the TotalInvAmt field.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) SetTotalInvAmt(v MultiCurrencyMoneyOpenApi) {
	o.TotalInvAmt = &v
}

// GetTotalInvedAmt returns the TotalInvedAmt field value if set, zero value otherwise.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetTotalInvedAmt() MultiCurrencyMoneyOpenApi {
	if o == nil || IsNil(o.TotalInvedAmt) {
		var ret MultiCurrencyMoneyOpenApi
		return ret
	}
	return *o.TotalInvedAmt
}

// GetTotalInvedAmtOk returns a tuple with the TotalInvedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetTotalInvedAmtOk() (*MultiCurrencyMoneyOpenApi, bool) {
	if o == nil || IsNil(o.TotalInvedAmt) {
		return nil, false
	}
	return o.TotalInvedAmt, true
}

// HasTotalInvedAmt returns a boolean if a field has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) HasTotalInvedAmt() bool {
	if o != nil && !IsNil(o.TotalInvedAmt) {
		return true
	}

	return false
}

// SetTotalInvedAmt gets a reference to the given MultiCurrencyMoneyOpenApi and assigns it to the TotalInvedAmt field.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) SetTotalInvedAmt(v MultiCurrencyMoneyOpenApi) {
	o.TotalInvedAmt = &v
}

// GetTotalLinkInvoiceAmt returns the TotalLinkInvoiceAmt field value if set, zero value otherwise.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetTotalLinkInvoiceAmt() MultiCurrencyMoneyOpenApi {
	if o == nil || IsNil(o.TotalLinkInvoiceAmt) {
		var ret MultiCurrencyMoneyOpenApi
		return ret
	}
	return *o.TotalLinkInvoiceAmt
}

// GetTotalLinkInvoiceAmtOk returns a tuple with the TotalLinkInvoiceAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) GetTotalLinkInvoiceAmtOk() (*MultiCurrencyMoneyOpenApi, bool) {
	if o == nil || IsNil(o.TotalLinkInvoiceAmt) {
		return nil, false
	}
	return o.TotalLinkInvoiceAmt, true
}

// HasTotalLinkInvoiceAmt returns a boolean if a field has been set.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) HasTotalLinkInvoiceAmt() bool {
	if o != nil && !IsNil(o.TotalLinkInvoiceAmt) {
		return true
	}

	return false
}

// SetTotalLinkInvoiceAmt gets a reference to the given MultiCurrencyMoneyOpenApi and assigns it to the TotalLinkInvoiceAmt field.
func (o *AlipayBossFncInvoicereceiptBatchqueryResponseModel) SetTotalLinkInvoiceAmt(v MultiCurrencyMoneyOpenApi) {
	o.TotalLinkInvoiceAmt = &v
}

func (o AlipayBossFncInvoicereceiptBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayBossFncInvoicereceiptBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultSet) {
		toSerialize["result_set"] = o.ResultSet
	}
	if !IsNil(o.TotalInvAmt) {
		toSerialize["total_inv_amt"] = o.TotalInvAmt
	}
	if !IsNil(o.TotalInvedAmt) {
		toSerialize["total_inved_amt"] = o.TotalInvedAmt
	}
	if !IsNil(o.TotalLinkInvoiceAmt) {
		toSerialize["total_link_invoice_amt"] = o.TotalLinkInvoiceAmt
	}
	return toSerialize, nil
}

type NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel struct {
	value *AlipayBossFncInvoicereceiptBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel) Get() *AlipayBossFncInvoicereceiptBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel) Set(val *AlipayBossFncInvoicereceiptBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayBossFncInvoicereceiptBatchqueryResponseModel(val *AlipayBossFncInvoicereceiptBatchqueryResponseModel) *NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel {
	return &NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayBossFncInvoicereceiptBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

