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

// checks if the ExpenseInvoiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpenseInvoiceInfo{}

// ExpenseInvoiceInfo struct for ExpenseInvoiceInfo
type ExpenseInvoiceInfo struct {
	// 员工ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// 员工ID
	EmployeeOpenId    *string            `json:"employee_open_id,omitempty"`
	InvoiceOutputInfo *InvoiceOutputInfo `json:"invoice_output_info,omitempty"`
	OcrNormalScanInfo *OcrNormalScanInfo `json:"ocr_normal_scan_info,omitempty"`
	OcrPlaneScanInfo  *OcrPlaneScanInfo  `json:"ocr_plane_scan_info,omitempty"`
	OcrTaxiScanInfo   *OcrTaxiScanInfo   `json:"ocr_taxi_scan_info,omitempty"`
	OcrTrainScanInfo  *OcrTrainScanInfo  `json:"ocr_train_scan_info,omitempty"`
	VoucherFileInfo   *VoucherFileInfo   `json:"voucher_file_info,omitempty"`
	// 凭证ID
	VoucherId *string `json:"voucher_id,omitempty"`
}

// NewExpenseInvoiceInfo instantiates a new ExpenseInvoiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseInvoiceInfo() *ExpenseInvoiceInfo {
	this := ExpenseInvoiceInfo{}
	return &this
}

// NewExpenseInvoiceInfoWithDefaults instantiates a new ExpenseInvoiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseInvoiceInfoWithDefaults() *ExpenseInvoiceInfo {
	this := ExpenseInvoiceInfo{}
	return &this
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetEmployeeId() string {
	if o == nil || IsNil(o.EmployeeId) {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetEmployeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeId) {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasEmployeeId() bool {
	if o != nil && !IsNil(o.EmployeeId) {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *ExpenseInvoiceInfo) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

// GetEmployeeOpenId returns the EmployeeOpenId field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetEmployeeOpenId() string {
	if o == nil || IsNil(o.EmployeeOpenId) {
		var ret string
		return ret
	}
	return *o.EmployeeOpenId
}

// GetEmployeeOpenIdOk returns a tuple with the EmployeeOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetEmployeeOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeOpenId) {
		return nil, false
	}
	return o.EmployeeOpenId, true
}

// HasEmployeeOpenId returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasEmployeeOpenId() bool {
	if o != nil && !IsNil(o.EmployeeOpenId) {
		return true
	}

	return false
}

// SetEmployeeOpenId gets a reference to the given string and assigns it to the EmployeeOpenId field.
func (o *ExpenseInvoiceInfo) SetEmployeeOpenId(v string) {
	o.EmployeeOpenId = &v
}

// GetInvoiceOutputInfo returns the InvoiceOutputInfo field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetInvoiceOutputInfo() InvoiceOutputInfo {
	if o == nil || IsNil(o.InvoiceOutputInfo) {
		var ret InvoiceOutputInfo
		return ret
	}
	return *o.InvoiceOutputInfo
}

// GetInvoiceOutputInfoOk returns a tuple with the InvoiceOutputInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetInvoiceOutputInfoOk() (*InvoiceOutputInfo, bool) {
	if o == nil || IsNil(o.InvoiceOutputInfo) {
		return nil, false
	}
	return o.InvoiceOutputInfo, true
}

// HasInvoiceOutputInfo returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasInvoiceOutputInfo() bool {
	if o != nil && !IsNil(o.InvoiceOutputInfo) {
		return true
	}

	return false
}

// SetInvoiceOutputInfo gets a reference to the given InvoiceOutputInfo and assigns it to the InvoiceOutputInfo field.
func (o *ExpenseInvoiceInfo) SetInvoiceOutputInfo(v InvoiceOutputInfo) {
	o.InvoiceOutputInfo = &v
}

// GetOcrNormalScanInfo returns the OcrNormalScanInfo field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetOcrNormalScanInfo() OcrNormalScanInfo {
	if o == nil || IsNil(o.OcrNormalScanInfo) {
		var ret OcrNormalScanInfo
		return ret
	}
	return *o.OcrNormalScanInfo
}

// GetOcrNormalScanInfoOk returns a tuple with the OcrNormalScanInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetOcrNormalScanInfoOk() (*OcrNormalScanInfo, bool) {
	if o == nil || IsNil(o.OcrNormalScanInfo) {
		return nil, false
	}
	return o.OcrNormalScanInfo, true
}

// HasOcrNormalScanInfo returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasOcrNormalScanInfo() bool {
	if o != nil && !IsNil(o.OcrNormalScanInfo) {
		return true
	}

	return false
}

// SetOcrNormalScanInfo gets a reference to the given OcrNormalScanInfo and assigns it to the OcrNormalScanInfo field.
func (o *ExpenseInvoiceInfo) SetOcrNormalScanInfo(v OcrNormalScanInfo) {
	o.OcrNormalScanInfo = &v
}

// GetOcrPlaneScanInfo returns the OcrPlaneScanInfo field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetOcrPlaneScanInfo() OcrPlaneScanInfo {
	if o == nil || IsNil(o.OcrPlaneScanInfo) {
		var ret OcrPlaneScanInfo
		return ret
	}
	return *o.OcrPlaneScanInfo
}

// GetOcrPlaneScanInfoOk returns a tuple with the OcrPlaneScanInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetOcrPlaneScanInfoOk() (*OcrPlaneScanInfo, bool) {
	if o == nil || IsNil(o.OcrPlaneScanInfo) {
		return nil, false
	}
	return o.OcrPlaneScanInfo, true
}

// HasOcrPlaneScanInfo returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasOcrPlaneScanInfo() bool {
	if o != nil && !IsNil(o.OcrPlaneScanInfo) {
		return true
	}

	return false
}

// SetOcrPlaneScanInfo gets a reference to the given OcrPlaneScanInfo and assigns it to the OcrPlaneScanInfo field.
func (o *ExpenseInvoiceInfo) SetOcrPlaneScanInfo(v OcrPlaneScanInfo) {
	o.OcrPlaneScanInfo = &v
}

// GetOcrTaxiScanInfo returns the OcrTaxiScanInfo field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetOcrTaxiScanInfo() OcrTaxiScanInfo {
	if o == nil || IsNil(o.OcrTaxiScanInfo) {
		var ret OcrTaxiScanInfo
		return ret
	}
	return *o.OcrTaxiScanInfo
}

// GetOcrTaxiScanInfoOk returns a tuple with the OcrTaxiScanInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetOcrTaxiScanInfoOk() (*OcrTaxiScanInfo, bool) {
	if o == nil || IsNil(o.OcrTaxiScanInfo) {
		return nil, false
	}
	return o.OcrTaxiScanInfo, true
}

// HasOcrTaxiScanInfo returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasOcrTaxiScanInfo() bool {
	if o != nil && !IsNil(o.OcrTaxiScanInfo) {
		return true
	}

	return false
}

// SetOcrTaxiScanInfo gets a reference to the given OcrTaxiScanInfo and assigns it to the OcrTaxiScanInfo field.
func (o *ExpenseInvoiceInfo) SetOcrTaxiScanInfo(v OcrTaxiScanInfo) {
	o.OcrTaxiScanInfo = &v
}

// GetOcrTrainScanInfo returns the OcrTrainScanInfo field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetOcrTrainScanInfo() OcrTrainScanInfo {
	if o == nil || IsNil(o.OcrTrainScanInfo) {
		var ret OcrTrainScanInfo
		return ret
	}
	return *o.OcrTrainScanInfo
}

// GetOcrTrainScanInfoOk returns a tuple with the OcrTrainScanInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetOcrTrainScanInfoOk() (*OcrTrainScanInfo, bool) {
	if o == nil || IsNil(o.OcrTrainScanInfo) {
		return nil, false
	}
	return o.OcrTrainScanInfo, true
}

// HasOcrTrainScanInfo returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasOcrTrainScanInfo() bool {
	if o != nil && !IsNil(o.OcrTrainScanInfo) {
		return true
	}

	return false
}

// SetOcrTrainScanInfo gets a reference to the given OcrTrainScanInfo and assigns it to the OcrTrainScanInfo field.
func (o *ExpenseInvoiceInfo) SetOcrTrainScanInfo(v OcrTrainScanInfo) {
	o.OcrTrainScanInfo = &v
}

// GetVoucherFileInfo returns the VoucherFileInfo field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetVoucherFileInfo() VoucherFileInfo {
	if o == nil || IsNil(o.VoucherFileInfo) {
		var ret VoucherFileInfo
		return ret
	}
	return *o.VoucherFileInfo
}

// GetVoucherFileInfoOk returns a tuple with the VoucherFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetVoucherFileInfoOk() (*VoucherFileInfo, bool) {
	if o == nil || IsNil(o.VoucherFileInfo) {
		return nil, false
	}
	return o.VoucherFileInfo, true
}

// HasVoucherFileInfo returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasVoucherFileInfo() bool {
	if o != nil && !IsNil(o.VoucherFileInfo) {
		return true
	}

	return false
}

// SetVoucherFileInfo gets a reference to the given VoucherFileInfo and assigns it to the VoucherFileInfo field.
func (o *ExpenseInvoiceInfo) SetVoucherFileInfo(v VoucherFileInfo) {
	o.VoucherFileInfo = &v
}

// GetVoucherId returns the VoucherId field value if set, zero value otherwise.
func (o *ExpenseInvoiceInfo) GetVoucherId() string {
	if o == nil || IsNil(o.VoucherId) {
		var ret string
		return ret
	}
	return *o.VoucherId
}

// GetVoucherIdOk returns a tuple with the VoucherId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseInvoiceInfo) GetVoucherIdOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherId) {
		return nil, false
	}
	return o.VoucherId, true
}

// HasVoucherId returns a boolean if a field has been set.
func (o *ExpenseInvoiceInfo) HasVoucherId() bool {
	if o != nil && !IsNil(o.VoucherId) {
		return true
	}

	return false
}

// SetVoucherId gets a reference to the given string and assigns it to the VoucherId field.
func (o *ExpenseInvoiceInfo) SetVoucherId(v string) {
	o.VoucherId = &v
}

func (o ExpenseInvoiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpenseInvoiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmployeeId) {
		toSerialize["employee_id"] = o.EmployeeId
	}
	if !IsNil(o.EmployeeOpenId) {
		toSerialize["employee_open_id"] = o.EmployeeOpenId
	}
	if !IsNil(o.InvoiceOutputInfo) {
		toSerialize["invoice_output_info"] = o.InvoiceOutputInfo
	}
	if !IsNil(o.OcrNormalScanInfo) {
		toSerialize["ocr_normal_scan_info"] = o.OcrNormalScanInfo
	}
	if !IsNil(o.OcrPlaneScanInfo) {
		toSerialize["ocr_plane_scan_info"] = o.OcrPlaneScanInfo
	}
	if !IsNil(o.OcrTaxiScanInfo) {
		toSerialize["ocr_taxi_scan_info"] = o.OcrTaxiScanInfo
	}
	if !IsNil(o.OcrTrainScanInfo) {
		toSerialize["ocr_train_scan_info"] = o.OcrTrainScanInfo
	}
	if !IsNil(o.VoucherFileInfo) {
		toSerialize["voucher_file_info"] = o.VoucherFileInfo
	}
	if !IsNil(o.VoucherId) {
		toSerialize["voucher_id"] = o.VoucherId
	}
	return toSerialize, nil
}

type NullableExpenseInvoiceInfo struct {
	value *ExpenseInvoiceInfo
	isSet bool
}

func (v NullableExpenseInvoiceInfo) Get() *ExpenseInvoiceInfo {
	return v.value
}

func (v *NullableExpenseInvoiceInfo) Set(val *ExpenseInvoiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseInvoiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseInvoiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseInvoiceInfo(val *ExpenseInvoiceInfo) *NullableExpenseInvoiceInfo {
	return &NullableExpenseInvoiceInfo{value: val, isSet: true}
}

func (v NullableExpenseInvoiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseInvoiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
