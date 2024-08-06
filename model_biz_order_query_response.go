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

// checks if the BizOrderQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BizOrderQueryResponse{}

// BizOrderQueryResponse struct for BizOrderQueryResponse
type BizOrderQueryResponse struct {
	// 操作动作。  CREATE_SHOP-创建门店，  MODIFY_SHOP-修改门店，  CREATE_ITEM-创建商品，  MODIFY_ITEM-修改商品，  EFFECTIVE_ITEM-上架商品，  INVALID_ITEM-下架商品，  RESUME_ITEM-暂停售卖商品，  PAUSE_ITEM-恢复售卖商品
	Action *string `json:"action,omitempty"`
	// 操作模式：NORMAL-普通开店
	ActionMode *string `json:"action_mode,omitempty"`
	// 支付宝流水ID
	ApplyId *string `json:"apply_id,omitempty"`
	// 流水上下文信息，JSON格式。根据action不同对应的结构也不同，JSON字段与含义可参考各个接口的请求参数。
	BizContextInfo *string `json:"biz_context_info,omitempty"`
	// 业务主体ID。根据biz_type不同可能对应shop_id或item_id。  特别注意对于门店创建，当流水status=SUCCESS时，此字段才为shop_id，其他状态时为0或空。
	BizId *string `json:"biz_id,omitempty"`
	// 业务类型：SHOP-店铺，ITEM-商品
	BizType *string `json:"biz_type,omitempty"`
	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 操作用户的支付账号id
	OpId *string `json:"op_id,omitempty"`
	// 注意：此字段并非外部商户请求时传入的request_id，暂时代表支付宝内部字段，请勿用。
	RequestId *string `json:"request_id,omitempty"`
	// 流水处理结果码  <a href=\"https://doc.open.alipay.com/doc2/detail.htm?spm=a219a.7629140.0.0.lL9hGI&treeId=78&articleId=103834&docType=1#s2\">点此查看</a>
	ResultCode *string `json:"result_code,omitempty"`
	// 流水处理结果描述
	ResultDesc *string `json:"result_desc,omitempty"`
	// 流水状态：INIT-初始，PROCESS-处理中，SUCCESS-成功，FAIL-失败。
	Status *string `json:"status,omitempty"`
	// 流水子状态：WAIT_CERTIFY-等待认证，LICENSE_AUDITING-证照审核中，RISK_AUDITING-风控审核中，WAIT_SIGN-等待签约，FINISH-终结。
	SubStatus *string `json:"sub_status,omitempty"`
	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

// NewBizOrderQueryResponse instantiates a new BizOrderQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBizOrderQueryResponse() *BizOrderQueryResponse {
	this := BizOrderQueryResponse{}
	return &this
}

// NewBizOrderQueryResponseWithDefaults instantiates a new BizOrderQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBizOrderQueryResponseWithDefaults() *BizOrderQueryResponse {
	this := BizOrderQueryResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BizOrderQueryResponse) SetAction(v string) {
	o.Action = &v
}

// GetActionMode returns the ActionMode field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetActionMode() string {
	if o == nil || IsNil(o.ActionMode) {
		var ret string
		return ret
	}
	return *o.ActionMode
}

// GetActionModeOk returns a tuple with the ActionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetActionModeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionMode) {
		return nil, false
	}
	return o.ActionMode, true
}

// HasActionMode returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasActionMode() bool {
	if o != nil && !IsNil(o.ActionMode) {
		return true
	}

	return false
}

// SetActionMode gets a reference to the given string and assigns it to the ActionMode field.
func (o *BizOrderQueryResponse) SetActionMode(v string) {
	o.ActionMode = &v
}

// GetApplyId returns the ApplyId field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetApplyId() string {
	if o == nil || IsNil(o.ApplyId) {
		var ret string
		return ret
	}
	return *o.ApplyId
}

// GetApplyIdOk returns a tuple with the ApplyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetApplyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyId) {
		return nil, false
	}
	return o.ApplyId, true
}

// HasApplyId returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasApplyId() bool {
	if o != nil && !IsNil(o.ApplyId) {
		return true
	}

	return false
}

// SetApplyId gets a reference to the given string and assigns it to the ApplyId field.
func (o *BizOrderQueryResponse) SetApplyId(v string) {
	o.ApplyId = &v
}

// GetBizContextInfo returns the BizContextInfo field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetBizContextInfo() string {
	if o == nil || IsNil(o.BizContextInfo) {
		var ret string
		return ret
	}
	return *o.BizContextInfo
}

// GetBizContextInfoOk returns a tuple with the BizContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetBizContextInfoOk() (*string, bool) {
	if o == nil || IsNil(o.BizContextInfo) {
		return nil, false
	}
	return o.BizContextInfo, true
}

// HasBizContextInfo returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasBizContextInfo() bool {
	if o != nil && !IsNil(o.BizContextInfo) {
		return true
	}

	return false
}

// SetBizContextInfo gets a reference to the given string and assigns it to the BizContextInfo field.
func (o *BizOrderQueryResponse) SetBizContextInfo(v string) {
	o.BizContextInfo = &v
}

// GetBizId returns the BizId field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetBizId() string {
	if o == nil || IsNil(o.BizId) {
		var ret string
		return ret
	}
	return *o.BizId
}

// GetBizIdOk returns a tuple with the BizId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetBizIdOk() (*string, bool) {
	if o == nil || IsNil(o.BizId) {
		return nil, false
	}
	return o.BizId, true
}

// HasBizId returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasBizId() bool {
	if o != nil && !IsNil(o.BizId) {
		return true
	}

	return false
}

// SetBizId gets a reference to the given string and assigns it to the BizId field.
func (o *BizOrderQueryResponse) SetBizId(v string) {
	o.BizId = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *BizOrderQueryResponse) SetBizType(v string) {
	o.BizType = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *BizOrderQueryResponse) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetOpId returns the OpId field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetOpId() string {
	if o == nil || IsNil(o.OpId) {
		var ret string
		return ret
	}
	return *o.OpId
}

// GetOpIdOk returns a tuple with the OpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetOpIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpId) {
		return nil, false
	}
	return o.OpId, true
}

// HasOpId returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasOpId() bool {
	if o != nil && !IsNil(o.OpId) {
		return true
	}

	return false
}

// SetOpId gets a reference to the given string and assigns it to the OpId field.
func (o *BizOrderQueryResponse) SetOpId(v string) {
	o.OpId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *BizOrderQueryResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *BizOrderQueryResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultDesc returns the ResultDesc field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetResultDesc() string {
	if o == nil || IsNil(o.ResultDesc) {
		var ret string
		return ret
	}
	return *o.ResultDesc
}

// GetResultDescOk returns a tuple with the ResultDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetResultDescOk() (*string, bool) {
	if o == nil || IsNil(o.ResultDesc) {
		return nil, false
	}
	return o.ResultDesc, true
}

// HasResultDesc returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasResultDesc() bool {
	if o != nil && !IsNil(o.ResultDesc) {
		return true
	}

	return false
}

// SetResultDesc gets a reference to the given string and assigns it to the ResultDesc field.
func (o *BizOrderQueryResponse) SetResultDesc(v string) {
	o.ResultDesc = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BizOrderQueryResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetSubStatus() string {
	if o == nil || IsNil(o.SubStatus) {
		var ret string
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetSubStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given string and assigns it to the SubStatus field.
func (o *BizOrderQueryResponse) SetSubStatus(v string) {
	o.SubStatus = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *BizOrderQueryResponse) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizOrderQueryResponse) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *BizOrderQueryResponse) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *BizOrderQueryResponse) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

func (o BizOrderQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BizOrderQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ActionMode) {
		toSerialize["action_mode"] = o.ActionMode
	}
	if !IsNil(o.ApplyId) {
		toSerialize["apply_id"] = o.ApplyId
	}
	if !IsNil(o.BizContextInfo) {
		toSerialize["biz_context_info"] = o.BizContextInfo
	}
	if !IsNil(o.BizId) {
		toSerialize["biz_id"] = o.BizId
	}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.OpId) {
		toSerialize["op_id"] = o.OpId
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.ResultCode) {
		toSerialize["result_code"] = o.ResultCode
	}
	if !IsNil(o.ResultDesc) {
		toSerialize["result_desc"] = o.ResultDesc
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubStatus) {
		toSerialize["sub_status"] = o.SubStatus
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableBizOrderQueryResponse struct {
	value *BizOrderQueryResponse
	isSet bool
}

func (v NullableBizOrderQueryResponse) Get() *BizOrderQueryResponse {
	return v.value
}

func (v *NullableBizOrderQueryResponse) Set(val *BizOrderQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBizOrderQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBizOrderQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBizOrderQueryResponse(val *BizOrderQueryResponse) *NullableBizOrderQueryResponse {
	return &NullableBizOrderQueryResponse{value: val, isSet: true}
}

func (v NullableBizOrderQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBizOrderQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


