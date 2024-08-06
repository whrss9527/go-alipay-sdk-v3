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

// checks if the AlipayFundTransOrderQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundTransOrderQueryResponseModel{}

// AlipayFundTransOrderQueryResponseModel struct for AlipayFundTransOrderQueryResponseModel
type AlipayFundTransOrderQueryResponseModel struct {
	// 预计到账时间，转账到银行卡专用，格式为yyyy-MM-dd HH:mm:ss，转账受理失败不返回。     注意：  此参数为预计时间，可能与实际到账时间有较大误差，不能作为实际到账时间使用，仅供参考用途。
	ArrivalTimeEnd *string `json:"arrival_time_end,omitempty"`
	// 查询失败时，本参数为错误代 码。   查询成功不返回。 对于退票订单，不返回该参数。
	ErrorCode *string `json:"error_code,omitempty"`
	// 查询到的订单状态为FAIL失败或REFUND退票时，返回具体的原因。
	FailReason *string `json:"fail_reason,omitempty"`
	// 预计收费金额（元），转账到银行卡专用，数字格式，精确到小数点后2位，转账失败或转账受理失败不返回。
	OrderFee *string `json:"order_fee,omitempty"`
	// 支付宝转账单据号，查询失败不返回。
	OrderId *string `json:"order_id,omitempty"`
	// 发起转账来源方定义的转账单据号。   该参数的赋值均以查询结果中 的 out_biz_no 为准。   如果查询失败，不返回该参数。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 支付时间，格式为yyyy-MM-dd HH:mm:ss，转账失败不返回。
	PayDate *string `json:"pay_date,omitempty"`
	// 转账单据状态。   SUCCESS：成功（配合\"单笔转账到银行账户接口\"产品使用时, 同一笔单据多次查询有可能从成功变成退票状态）；   FAIL：失败（具体失败原因请参见error_code以及fail_reason返回值）；   INIT：等待处理；   DEALING：处理中；   REFUND：退票（仅配合\"单笔转账到银行账户接口\"产品使用时会涉及, 具体退票原因请参见fail_reason返回值）；   UNKNOWN：状态未知。
	Status *string `json:"status,omitempty"`
}

// NewAlipayFundTransOrderQueryResponseModel instantiates a new AlipayFundTransOrderQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundTransOrderQueryResponseModel() *AlipayFundTransOrderQueryResponseModel {
	this := AlipayFundTransOrderQueryResponseModel{}
	return &this
}

// NewAlipayFundTransOrderQueryResponseModelWithDefaults instantiates a new AlipayFundTransOrderQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundTransOrderQueryResponseModelWithDefaults() *AlipayFundTransOrderQueryResponseModel {
	this := AlipayFundTransOrderQueryResponseModel{}
	return &this
}

// GetArrivalTimeEnd returns the ArrivalTimeEnd field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetArrivalTimeEnd() string {
	if o == nil || IsNil(o.ArrivalTimeEnd) {
		var ret string
		return ret
	}
	return *o.ArrivalTimeEnd
}

// GetArrivalTimeEndOk returns a tuple with the ArrivalTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetArrivalTimeEndOk() (*string, bool) {
	if o == nil || IsNil(o.ArrivalTimeEnd) {
		return nil, false
	}
	return o.ArrivalTimeEnd, true
}

// HasArrivalTimeEnd returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasArrivalTimeEnd() bool {
	if o != nil && !IsNil(o.ArrivalTimeEnd) {
		return true
	}

	return false
}

// SetArrivalTimeEnd gets a reference to the given string and assigns it to the ArrivalTimeEnd field.
func (o *AlipayFundTransOrderQueryResponseModel) SetArrivalTimeEnd(v string) {
	o.ArrivalTimeEnd = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *AlipayFundTransOrderQueryResponseModel) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *AlipayFundTransOrderQueryResponseModel) SetFailReason(v string) {
	o.FailReason = &v
}

// GetOrderFee returns the OrderFee field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetOrderFee() string {
	if o == nil || IsNil(o.OrderFee) {
		var ret string
		return ret
	}
	return *o.OrderFee
}

// GetOrderFeeOk returns a tuple with the OrderFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetOrderFeeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderFee) {
		return nil, false
	}
	return o.OrderFee, true
}

// HasOrderFee returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasOrderFee() bool {
	if o != nil && !IsNil(o.OrderFee) {
		return true
	}

	return false
}

// SetOrderFee gets a reference to the given string and assigns it to the OrderFee field.
func (o *AlipayFundTransOrderQueryResponseModel) SetOrderFee(v string) {
	o.OrderFee = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AlipayFundTransOrderQueryResponseModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayFundTransOrderQueryResponseModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPayDate returns the PayDate field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetPayDate() string {
	if o == nil || IsNil(o.PayDate) {
		var ret string
		return ret
	}
	return *o.PayDate
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetPayDateOk() (*string, bool) {
	if o == nil || IsNil(o.PayDate) {
		return nil, false
	}
	return o.PayDate, true
}

// HasPayDate returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasPayDate() bool {
	if o != nil && !IsNil(o.PayDate) {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given string and assigns it to the PayDate field.
func (o *AlipayFundTransOrderQueryResponseModel) SetPayDate(v string) {
	o.PayDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayFundTransOrderQueryResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransOrderQueryResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayFundTransOrderQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayFundTransOrderQueryResponseModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayFundTransOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundTransOrderQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArrivalTimeEnd) {
		toSerialize["arrival_time_end"] = o.ArrivalTimeEnd
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.FailReason) {
		toSerialize["fail_reason"] = o.FailReason
	}
	if !IsNil(o.OrderFee) {
		toSerialize["order_fee"] = o.OrderFee
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PayDate) {
		toSerialize["pay_date"] = o.PayDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayFundTransOrderQueryResponseModel struct {
	value *AlipayFundTransOrderQueryResponseModel
	isSet bool
}

func (v NullableAlipayFundTransOrderQueryResponseModel) Get() *AlipayFundTransOrderQueryResponseModel {
	return v.value
}

func (v *NullableAlipayFundTransOrderQueryResponseModel) Set(val *AlipayFundTransOrderQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundTransOrderQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundTransOrderQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundTransOrderQueryResponseModel(val *AlipayFundTransOrderQueryResponseModel) *NullableAlipayFundTransOrderQueryResponseModel {
	return &NullableAlipayFundTransOrderQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundTransOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundTransOrderQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
