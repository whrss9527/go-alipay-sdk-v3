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

// checks if the IsvMerchantSalesDetailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsvMerchantSalesDetailRequest{}

// IsvMerchantSalesDetailRequest struct for IsvMerchantSalesDetailRequest
type IsvMerchantSalesDetailRequest struct {
	// 配券数
	CouponsQuantity *string `json:"coupons_quantity,omitempty"`
	// 设备详情
	DeviceDetail *string `json:"device_detail,omitempty"`
	// 商户pid
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// 小程序appid，若推广的商品不为小程序，则不传此参数
	MiniAppid *string `json:"mini_appid,omitempty"`
	// 作业地
	OperationPlace *string `json:"operation_place,omitempty"`
	// 外部业务号，传isv系统生成的账单号，需要保证唯一
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 推广服务商(S2)pid
	PromotorPid *string `json:"promotor_pid,omitempty"`
	// 销售金额，这里填写的是整数，单位为分，比如1元，那么输入100
	SalesAmount *string `json:"sales_amount,omitempty"`
	// 销售笔数
	SalesQuantity *string `json:"sales_quantity,omitempty"`
	// 推广服务商(S2)子账号pid
	SubPromotorPid *string `json:"sub_promotor_pid,omitempty"`
	// 核销金额，这里填写的是整数，单位为分，比如1元，那么输入100
	WriteOffAmount *string `json:"write_off_amount,omitempty"`
	// 核销数
	WriteOffQuantity *string `json:"write_off_quantity,omitempty"`
}

// NewIsvMerchantSalesDetailRequest instantiates a new IsvMerchantSalesDetailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsvMerchantSalesDetailRequest() *IsvMerchantSalesDetailRequest {
	this := IsvMerchantSalesDetailRequest{}
	return &this
}

// NewIsvMerchantSalesDetailRequestWithDefaults instantiates a new IsvMerchantSalesDetailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsvMerchantSalesDetailRequestWithDefaults() *IsvMerchantSalesDetailRequest {
	this := IsvMerchantSalesDetailRequest{}
	return &this
}

// GetCouponsQuantity returns the CouponsQuantity field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetCouponsQuantity() string {
	if o == nil || IsNil(o.CouponsQuantity) {
		var ret string
		return ret
	}
	return *o.CouponsQuantity
}

// GetCouponsQuantityOk returns a tuple with the CouponsQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetCouponsQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.CouponsQuantity) {
		return nil, false
	}
	return o.CouponsQuantity, true
}

// HasCouponsQuantity returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasCouponsQuantity() bool {
	if o != nil && !IsNil(o.CouponsQuantity) {
		return true
	}

	return false
}

// SetCouponsQuantity gets a reference to the given string and assigns it to the CouponsQuantity field.
func (o *IsvMerchantSalesDetailRequest) SetCouponsQuantity(v string) {
	o.CouponsQuantity = &v
}

// GetDeviceDetail returns the DeviceDetail field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetDeviceDetail() string {
	if o == nil || IsNil(o.DeviceDetail) {
		var ret string
		return ret
	}
	return *o.DeviceDetail
}

// GetDeviceDetailOk returns a tuple with the DeviceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetDeviceDetailOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceDetail) {
		return nil, false
	}
	return o.DeviceDetail, true
}

// HasDeviceDetail returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasDeviceDetail() bool {
	if o != nil && !IsNil(o.DeviceDetail) {
		return true
	}

	return false
}

// SetDeviceDetail gets a reference to the given string and assigns it to the DeviceDetail field.
func (o *IsvMerchantSalesDetailRequest) SetDeviceDetail(v string) {
	o.DeviceDetail = &v
}

// GetMerchantPid returns the MerchantPid field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetMerchantPid() string {
	if o == nil || IsNil(o.MerchantPid) {
		var ret string
		return ret
	}
	return *o.MerchantPid
}

// GetMerchantPidOk returns a tuple with the MerchantPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetMerchantPidOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantPid) {
		return nil, false
	}
	return o.MerchantPid, true
}

// HasMerchantPid returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasMerchantPid() bool {
	if o != nil && !IsNil(o.MerchantPid) {
		return true
	}

	return false
}

// SetMerchantPid gets a reference to the given string and assigns it to the MerchantPid field.
func (o *IsvMerchantSalesDetailRequest) SetMerchantPid(v string) {
	o.MerchantPid = &v
}

// GetMiniAppid returns the MiniAppid field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetMiniAppid() string {
	if o == nil || IsNil(o.MiniAppid) {
		var ret string
		return ret
	}
	return *o.MiniAppid
}

// GetMiniAppidOk returns a tuple with the MiniAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetMiniAppidOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppid) {
		return nil, false
	}
	return o.MiniAppid, true
}

// HasMiniAppid returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasMiniAppid() bool {
	if o != nil && !IsNil(o.MiniAppid) {
		return true
	}

	return false
}

// SetMiniAppid gets a reference to the given string and assigns it to the MiniAppid field.
func (o *IsvMerchantSalesDetailRequest) SetMiniAppid(v string) {
	o.MiniAppid = &v
}

// GetOperationPlace returns the OperationPlace field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetOperationPlace() string {
	if o == nil || IsNil(o.OperationPlace) {
		var ret string
		return ret
	}
	return *o.OperationPlace
}

// GetOperationPlaceOk returns a tuple with the OperationPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetOperationPlaceOk() (*string, bool) {
	if o == nil || IsNil(o.OperationPlace) {
		return nil, false
	}
	return o.OperationPlace, true
}

// HasOperationPlace returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasOperationPlace() bool {
	if o != nil && !IsNil(o.OperationPlace) {
		return true
	}

	return false
}

// SetOperationPlace gets a reference to the given string and assigns it to the OperationPlace field.
func (o *IsvMerchantSalesDetailRequest) SetOperationPlace(v string) {
	o.OperationPlace = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *IsvMerchantSalesDetailRequest) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPromotorPid returns the PromotorPid field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetPromotorPid() string {
	if o == nil || IsNil(o.PromotorPid) {
		var ret string
		return ret
	}
	return *o.PromotorPid
}

// GetPromotorPidOk returns a tuple with the PromotorPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetPromotorPidOk() (*string, bool) {
	if o == nil || IsNil(o.PromotorPid) {
		return nil, false
	}
	return o.PromotorPid, true
}

// HasPromotorPid returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasPromotorPid() bool {
	if o != nil && !IsNil(o.PromotorPid) {
		return true
	}

	return false
}

// SetPromotorPid gets a reference to the given string and assigns it to the PromotorPid field.
func (o *IsvMerchantSalesDetailRequest) SetPromotorPid(v string) {
	o.PromotorPid = &v
}

// GetSalesAmount returns the SalesAmount field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetSalesAmount() string {
	if o == nil || IsNil(o.SalesAmount) {
		var ret string
		return ret
	}
	return *o.SalesAmount
}

// GetSalesAmountOk returns a tuple with the SalesAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetSalesAmountOk() (*string, bool) {
	if o == nil || IsNil(o.SalesAmount) {
		return nil, false
	}
	return o.SalesAmount, true
}

// HasSalesAmount returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasSalesAmount() bool {
	if o != nil && !IsNil(o.SalesAmount) {
		return true
	}

	return false
}

// SetSalesAmount gets a reference to the given string and assigns it to the SalesAmount field.
func (o *IsvMerchantSalesDetailRequest) SetSalesAmount(v string) {
	o.SalesAmount = &v
}

// GetSalesQuantity returns the SalesQuantity field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetSalesQuantity() string {
	if o == nil || IsNil(o.SalesQuantity) {
		var ret string
		return ret
	}
	return *o.SalesQuantity
}

// GetSalesQuantityOk returns a tuple with the SalesQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetSalesQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.SalesQuantity) {
		return nil, false
	}
	return o.SalesQuantity, true
}

// HasSalesQuantity returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasSalesQuantity() bool {
	if o != nil && !IsNil(o.SalesQuantity) {
		return true
	}

	return false
}

// SetSalesQuantity gets a reference to the given string and assigns it to the SalesQuantity field.
func (o *IsvMerchantSalesDetailRequest) SetSalesQuantity(v string) {
	o.SalesQuantity = &v
}

// GetSubPromotorPid returns the SubPromotorPid field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetSubPromotorPid() string {
	if o == nil || IsNil(o.SubPromotorPid) {
		var ret string
		return ret
	}
	return *o.SubPromotorPid
}

// GetSubPromotorPidOk returns a tuple with the SubPromotorPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetSubPromotorPidOk() (*string, bool) {
	if o == nil || IsNil(o.SubPromotorPid) {
		return nil, false
	}
	return o.SubPromotorPid, true
}

// HasSubPromotorPid returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasSubPromotorPid() bool {
	if o != nil && !IsNil(o.SubPromotorPid) {
		return true
	}

	return false
}

// SetSubPromotorPid gets a reference to the given string and assigns it to the SubPromotorPid field.
func (o *IsvMerchantSalesDetailRequest) SetSubPromotorPid(v string) {
	o.SubPromotorPid = &v
}

// GetWriteOffAmount returns the WriteOffAmount field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetWriteOffAmount() string {
	if o == nil || IsNil(o.WriteOffAmount) {
		var ret string
		return ret
	}
	return *o.WriteOffAmount
}

// GetWriteOffAmountOk returns a tuple with the WriteOffAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetWriteOffAmountOk() (*string, bool) {
	if o == nil || IsNil(o.WriteOffAmount) {
		return nil, false
	}
	return o.WriteOffAmount, true
}

// HasWriteOffAmount returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasWriteOffAmount() bool {
	if o != nil && !IsNil(o.WriteOffAmount) {
		return true
	}

	return false
}

// SetWriteOffAmount gets a reference to the given string and assigns it to the WriteOffAmount field.
func (o *IsvMerchantSalesDetailRequest) SetWriteOffAmount(v string) {
	o.WriteOffAmount = &v
}

// GetWriteOffQuantity returns the WriteOffQuantity field value if set, zero value otherwise.
func (o *IsvMerchantSalesDetailRequest) GetWriteOffQuantity() string {
	if o == nil || IsNil(o.WriteOffQuantity) {
		var ret string
		return ret
	}
	return *o.WriteOffQuantity
}

// GetWriteOffQuantityOk returns a tuple with the WriteOffQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvMerchantSalesDetailRequest) GetWriteOffQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.WriteOffQuantity) {
		return nil, false
	}
	return o.WriteOffQuantity, true
}

// HasWriteOffQuantity returns a boolean if a field has been set.
func (o *IsvMerchantSalesDetailRequest) HasWriteOffQuantity() bool {
	if o != nil && !IsNil(o.WriteOffQuantity) {
		return true
	}

	return false
}

// SetWriteOffQuantity gets a reference to the given string and assigns it to the WriteOffQuantity field.
func (o *IsvMerchantSalesDetailRequest) SetWriteOffQuantity(v string) {
	o.WriteOffQuantity = &v
}

func (o IsvMerchantSalesDetailRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsvMerchantSalesDetailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CouponsQuantity) {
		toSerialize["coupons_quantity"] = o.CouponsQuantity
	}
	if !IsNil(o.DeviceDetail) {
		toSerialize["device_detail"] = o.DeviceDetail
	}
	if !IsNil(o.MerchantPid) {
		toSerialize["merchant_pid"] = o.MerchantPid
	}
	if !IsNil(o.MiniAppid) {
		toSerialize["mini_appid"] = o.MiniAppid
	}
	if !IsNil(o.OperationPlace) {
		toSerialize["operation_place"] = o.OperationPlace
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PromotorPid) {
		toSerialize["promotor_pid"] = o.PromotorPid
	}
	if !IsNil(o.SalesAmount) {
		toSerialize["sales_amount"] = o.SalesAmount
	}
	if !IsNil(o.SalesQuantity) {
		toSerialize["sales_quantity"] = o.SalesQuantity
	}
	if !IsNil(o.SubPromotorPid) {
		toSerialize["sub_promotor_pid"] = o.SubPromotorPid
	}
	if !IsNil(o.WriteOffAmount) {
		toSerialize["write_off_amount"] = o.WriteOffAmount
	}
	if !IsNil(o.WriteOffQuantity) {
		toSerialize["write_off_quantity"] = o.WriteOffQuantity
	}
	return toSerialize, nil
}

type NullableIsvMerchantSalesDetailRequest struct {
	value *IsvMerchantSalesDetailRequest
	isSet bool
}

func (v NullableIsvMerchantSalesDetailRequest) Get() *IsvMerchantSalesDetailRequest {
	return v.value
}

func (v *NullableIsvMerchantSalesDetailRequest) Set(val *IsvMerchantSalesDetailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIsvMerchantSalesDetailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIsvMerchantSalesDetailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsvMerchantSalesDetailRequest(val *IsvMerchantSalesDetailRequest) *NullableIsvMerchantSalesDetailRequest {
	return &NullableIsvMerchantSalesDetailRequest{value: val, isSet: true}
}

func (v NullableIsvMerchantSalesDetailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsvMerchantSalesDetailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
