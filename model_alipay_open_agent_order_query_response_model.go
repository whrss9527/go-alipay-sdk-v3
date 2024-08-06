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

// checks if the AlipayOpenAgentOrderQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAgentOrderQueryResponseModel{}

// AlipayOpenAgentOrderQueryResponseModel struct for AlipayOpenAgentOrderQueryResponseModel
type AlipayOpenAgentOrderQueryResponseModel struct {
	// 代理创建的应用ID，如果有代理商户创建应用，商户确认成功后，才返回应用ID，否则不返回。
	AgentAppId *string `json:"agent_app_id,omitempty"`
	// 只有申请单状态在MERCHANT_CONFIRM状态下，才会返回商户确认签约链接
	ConfirmUrl *string `json:"confirm_url,omitempty"`
	// 商户pid
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// 签约单号
	OrderNo *string `json:"order_no,omitempty"`
	// 支付宝商户入驻申请单状态，申请单状态包括：  MERCHANT_INFO_HOLD=暂存，提交事务出现业务校验异常时，会暂存申请单信息，可以调用业务接口修正参数，并重新提交  MERCHANT_AUDITING=审核中，申请信息正在人工审核中  MERCHANT_CONFIRM=待商户确认，申请信息审核通过，等待联系人确认签约或授权  MERCHANT_CONFIRM_SUCCESS=商户确认成功，商户同意签约或授权  MERCHANT_CONFIRM_TIME_OUT=商户超时未确认，如果商户受到确认信息15天内未确认，则需要重新提交申请信息  MERCHANT_APPLY_ORDER_CANCELED=审核失败或商户拒绝，申请信息审核被驳回，或者商户选择拒绝签约或授权
	OrderStatus *string `json:"order_status,omitempty"`
	// 申请单中每个产品的签约状态
	ProductAgentStatusInfos []ProductAgentStatusInfo `json:"product_agent_status_infos,omitempty"`
	// 审核失败的拒绝原因，只有审核失败才会返回该值
	RejectReason *string `json:"reject_reason,omitempty"`
	// 受限信息
	RestrictInfos []SignRestrictInfo `json:"restrict_infos,omitempty"`
}

// NewAlipayOpenAgentOrderQueryResponseModel instantiates a new AlipayOpenAgentOrderQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAgentOrderQueryResponseModel() *AlipayOpenAgentOrderQueryResponseModel {
	this := AlipayOpenAgentOrderQueryResponseModel{}
	return &this
}

// NewAlipayOpenAgentOrderQueryResponseModelWithDefaults instantiates a new AlipayOpenAgentOrderQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAgentOrderQueryResponseModelWithDefaults() *AlipayOpenAgentOrderQueryResponseModel {
	this := AlipayOpenAgentOrderQueryResponseModel{}
	return &this
}

// GetAgentAppId returns the AgentAppId field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetAgentAppId() string {
	if o == nil || IsNil(o.AgentAppId) {
		var ret string
		return ret
	}
	return *o.AgentAppId
}

// GetAgentAppIdOk returns a tuple with the AgentAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetAgentAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgentAppId) {
		return nil, false
	}
	return o.AgentAppId, true
}

// HasAgentAppId returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasAgentAppId() bool {
	if o != nil && !IsNil(o.AgentAppId) {
		return true
	}

	return false
}

// SetAgentAppId gets a reference to the given string and assigns it to the AgentAppId field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetAgentAppId(v string) {
	o.AgentAppId = &v
}

// GetConfirmUrl returns the ConfirmUrl field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetConfirmUrl() string {
	if o == nil || IsNil(o.ConfirmUrl) {
		var ret string
		return ret
	}
	return *o.ConfirmUrl
}

// GetConfirmUrlOk returns a tuple with the ConfirmUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetConfirmUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmUrl) {
		return nil, false
	}
	return o.ConfirmUrl, true
}

// HasConfirmUrl returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasConfirmUrl() bool {
	if o != nil && !IsNil(o.ConfirmUrl) {
		return true
	}

	return false
}

// SetConfirmUrl gets a reference to the given string and assigns it to the ConfirmUrl field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetConfirmUrl(v string) {
	o.ConfirmUrl = &v
}

// GetMerchantPid returns the MerchantPid field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetMerchantPid() string {
	if o == nil || IsNil(o.MerchantPid) {
		var ret string
		return ret
	}
	return *o.MerchantPid
}

// GetMerchantPidOk returns a tuple with the MerchantPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetMerchantPidOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantPid) {
		return nil, false
	}
	return o.MerchantPid, true
}

// HasMerchantPid returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasMerchantPid() bool {
	if o != nil && !IsNil(o.MerchantPid) {
		return true
	}

	return false
}

// SetMerchantPid gets a reference to the given string and assigns it to the MerchantPid field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetMerchantPid(v string) {
	o.MerchantPid = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetProductAgentStatusInfos returns the ProductAgentStatusInfos field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetProductAgentStatusInfos() []ProductAgentStatusInfo {
	if o == nil || IsNil(o.ProductAgentStatusInfos) {
		var ret []ProductAgentStatusInfo
		return ret
	}
	return o.ProductAgentStatusInfos
}

// GetProductAgentStatusInfosOk returns a tuple with the ProductAgentStatusInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetProductAgentStatusInfosOk() ([]ProductAgentStatusInfo, bool) {
	if o == nil || IsNil(o.ProductAgentStatusInfos) {
		return nil, false
	}
	return o.ProductAgentStatusInfos, true
}

// HasProductAgentStatusInfos returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasProductAgentStatusInfos() bool {
	if o != nil && !IsNil(o.ProductAgentStatusInfos) {
		return true
	}

	return false
}

// SetProductAgentStatusInfos gets a reference to the given []ProductAgentStatusInfo and assigns it to the ProductAgentStatusInfos field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetProductAgentStatusInfos(v []ProductAgentStatusInfo) {
	o.ProductAgentStatusInfos = v
}

// GetRejectReason returns the RejectReason field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetRejectReason() string {
	if o == nil || IsNil(o.RejectReason) {
		var ret string
		return ret
	}
	return *o.RejectReason
}

// GetRejectReasonOk returns a tuple with the RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetRejectReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectReason) {
		return nil, false
	}
	return o.RejectReason, true
}

// HasRejectReason returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasRejectReason() bool {
	if o != nil && !IsNil(o.RejectReason) {
		return true
	}

	return false
}

// SetRejectReason gets a reference to the given string and assigns it to the RejectReason field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetRejectReason(v string) {
	o.RejectReason = &v
}

// GetRestrictInfos returns the RestrictInfos field value if set, zero value otherwise.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetRestrictInfos() []SignRestrictInfo {
	if o == nil || IsNil(o.RestrictInfos) {
		var ret []SignRestrictInfo
		return ret
	}
	return o.RestrictInfos
}

// GetRestrictInfosOk returns a tuple with the RestrictInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) GetRestrictInfosOk() ([]SignRestrictInfo, bool) {
	if o == nil || IsNil(o.RestrictInfos) {
		return nil, false
	}
	return o.RestrictInfos, true
}

// HasRestrictInfos returns a boolean if a field has been set.
func (o *AlipayOpenAgentOrderQueryResponseModel) HasRestrictInfos() bool {
	if o != nil && !IsNil(o.RestrictInfos) {
		return true
	}

	return false
}

// SetRestrictInfos gets a reference to the given []SignRestrictInfo and assigns it to the RestrictInfos field.
func (o *AlipayOpenAgentOrderQueryResponseModel) SetRestrictInfos(v []SignRestrictInfo) {
	o.RestrictInfos = v
}

func (o AlipayOpenAgentOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAgentOrderQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentAppId) {
		toSerialize["agent_app_id"] = o.AgentAppId
	}
	if !IsNil(o.ConfirmUrl) {
		toSerialize["confirm_url"] = o.ConfirmUrl
	}
	if !IsNil(o.MerchantPid) {
		toSerialize["merchant_pid"] = o.MerchantPid
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.ProductAgentStatusInfos) {
		toSerialize["product_agent_status_infos"] = o.ProductAgentStatusInfos
	}
	if !IsNil(o.RejectReason) {
		toSerialize["reject_reason"] = o.RejectReason
	}
	if !IsNil(o.RestrictInfos) {
		toSerialize["restrict_infos"] = o.RestrictInfos
	}
	return toSerialize, nil
}

type NullableAlipayOpenAgentOrderQueryResponseModel struct {
	value *AlipayOpenAgentOrderQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenAgentOrderQueryResponseModel) Get() *AlipayOpenAgentOrderQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAgentOrderQueryResponseModel) Set(val *AlipayOpenAgentOrderQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentOrderQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentOrderQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentOrderQueryResponseModel(val *AlipayOpenAgentOrderQueryResponseModel) *NullableAlipayOpenAgentOrderQueryResponseModel {
	return &NullableAlipayOpenAgentOrderQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentOrderQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


