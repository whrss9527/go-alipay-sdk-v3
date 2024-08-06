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

// checks if the PromotionTaskRelation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionTaskRelation{}

// PromotionTaskRelation struct for PromotionTaskRelation
type PromotionTaskRelation struct {
	// 申请推广时填写的申请理由
	ApplyReason *string `json:"apply_reason,omitempty"`
	// 服务商品ID
	CommodityId *string `json:"commodity_id,omitempty"`
	// 服务商品名称
	CommodityName *string `json:"commodity_name,omitempty"`
	// 推广关系创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 推广服务商联系人名称
	PromoterContactName *string `json:"promoter_contact_name,omitempty"`
	// 推广服务商联系电话
	PromoterContactPhone *string `json:"promoter_contact_phone,omitempty"`
	// 推广服务商的名称
	PromoterName *string `json:"promoter_name,omitempty"`
	// 推广服务商的pid
	PromoterPid *string `json:"promoter_pid,omitempty"`
	// 推广任务id
	PromotionId *string `json:"promotion_id,omitempty"`
	// 推广任务的名称
	PromotionName *string `json:"promotion_name,omitempty"`
}

// NewPromotionTaskRelation instantiates a new PromotionTaskRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionTaskRelation() *PromotionTaskRelation {
	this := PromotionTaskRelation{}
	return &this
}

// NewPromotionTaskRelationWithDefaults instantiates a new PromotionTaskRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionTaskRelationWithDefaults() *PromotionTaskRelation {
	this := PromotionTaskRelation{}
	return &this
}

// GetApplyReason returns the ApplyReason field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetApplyReason() string {
	if o == nil || IsNil(o.ApplyReason) {
		var ret string
		return ret
	}
	return *o.ApplyReason
}

// GetApplyReasonOk returns a tuple with the ApplyReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetApplyReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyReason) {
		return nil, false
	}
	return o.ApplyReason, true
}

// HasApplyReason returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasApplyReason() bool {
	if o != nil && !IsNil(o.ApplyReason) {
		return true
	}

	return false
}

// SetApplyReason gets a reference to the given string and assigns it to the ApplyReason field.
func (o *PromotionTaskRelation) SetApplyReason(v string) {
	o.ApplyReason = &v
}

// GetCommodityId returns the CommodityId field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetCommodityId() string {
	if o == nil || IsNil(o.CommodityId) {
		var ret string
		return ret
	}
	return *o.CommodityId
}

// GetCommodityIdOk returns a tuple with the CommodityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetCommodityIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommodityId) {
		return nil, false
	}
	return o.CommodityId, true
}

// HasCommodityId returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasCommodityId() bool {
	if o != nil && !IsNil(o.CommodityId) {
		return true
	}

	return false
}

// SetCommodityId gets a reference to the given string and assigns it to the CommodityId field.
func (o *PromotionTaskRelation) SetCommodityId(v string) {
	o.CommodityId = &v
}

// GetCommodityName returns the CommodityName field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetCommodityName() string {
	if o == nil || IsNil(o.CommodityName) {
		var ret string
		return ret
	}
	return *o.CommodityName
}

// GetCommodityNameOk returns a tuple with the CommodityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetCommodityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommodityName) {
		return nil, false
	}
	return o.CommodityName, true
}

// HasCommodityName returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasCommodityName() bool {
	if o != nil && !IsNil(o.CommodityName) {
		return true
	}

	return false
}

// SetCommodityName gets a reference to the given string and assigns it to the CommodityName field.
func (o *PromotionTaskRelation) SetCommodityName(v string) {
	o.CommodityName = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *PromotionTaskRelation) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetPromoterContactName returns the PromoterContactName field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetPromoterContactName() string {
	if o == nil || IsNil(o.PromoterContactName) {
		var ret string
		return ret
	}
	return *o.PromoterContactName
}

// GetPromoterContactNameOk returns a tuple with the PromoterContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetPromoterContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.PromoterContactName) {
		return nil, false
	}
	return o.PromoterContactName, true
}

// HasPromoterContactName returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasPromoterContactName() bool {
	if o != nil && !IsNil(o.PromoterContactName) {
		return true
	}

	return false
}

// SetPromoterContactName gets a reference to the given string and assigns it to the PromoterContactName field.
func (o *PromotionTaskRelation) SetPromoterContactName(v string) {
	o.PromoterContactName = &v
}

// GetPromoterContactPhone returns the PromoterContactPhone field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetPromoterContactPhone() string {
	if o == nil || IsNil(o.PromoterContactPhone) {
		var ret string
		return ret
	}
	return *o.PromoterContactPhone
}

// GetPromoterContactPhoneOk returns a tuple with the PromoterContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetPromoterContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.PromoterContactPhone) {
		return nil, false
	}
	return o.PromoterContactPhone, true
}

// HasPromoterContactPhone returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasPromoterContactPhone() bool {
	if o != nil && !IsNil(o.PromoterContactPhone) {
		return true
	}

	return false
}

// SetPromoterContactPhone gets a reference to the given string and assigns it to the PromoterContactPhone field.
func (o *PromotionTaskRelation) SetPromoterContactPhone(v string) {
	o.PromoterContactPhone = &v
}

// GetPromoterName returns the PromoterName field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetPromoterName() string {
	if o == nil || IsNil(o.PromoterName) {
		var ret string
		return ret
	}
	return *o.PromoterName
}

// GetPromoterNameOk returns a tuple with the PromoterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetPromoterNameOk() (*string, bool) {
	if o == nil || IsNil(o.PromoterName) {
		return nil, false
	}
	return o.PromoterName, true
}

// HasPromoterName returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasPromoterName() bool {
	if o != nil && !IsNil(o.PromoterName) {
		return true
	}

	return false
}

// SetPromoterName gets a reference to the given string and assigns it to the PromoterName field.
func (o *PromotionTaskRelation) SetPromoterName(v string) {
	o.PromoterName = &v
}

// GetPromoterPid returns the PromoterPid field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetPromoterPid() string {
	if o == nil || IsNil(o.PromoterPid) {
		var ret string
		return ret
	}
	return *o.PromoterPid
}

// GetPromoterPidOk returns a tuple with the PromoterPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetPromoterPidOk() (*string, bool) {
	if o == nil || IsNil(o.PromoterPid) {
		return nil, false
	}
	return o.PromoterPid, true
}

// HasPromoterPid returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasPromoterPid() bool {
	if o != nil && !IsNil(o.PromoterPid) {
		return true
	}

	return false
}

// SetPromoterPid gets a reference to the given string and assigns it to the PromoterPid field.
func (o *PromotionTaskRelation) SetPromoterPid(v string) {
	o.PromoterPid = &v
}

// GetPromotionId returns the PromotionId field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetPromotionId() string {
	if o == nil || IsNil(o.PromotionId) {
		var ret string
		return ret
	}
	return *o.PromotionId
}

// GetPromotionIdOk returns a tuple with the PromotionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetPromotionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionId) {
		return nil, false
	}
	return o.PromotionId, true
}

// HasPromotionId returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasPromotionId() bool {
	if o != nil && !IsNil(o.PromotionId) {
		return true
	}

	return false
}

// SetPromotionId gets a reference to the given string and assigns it to the PromotionId field.
func (o *PromotionTaskRelation) SetPromotionId(v string) {
	o.PromotionId = &v
}

// GetPromotionName returns the PromotionName field value if set, zero value otherwise.
func (o *PromotionTaskRelation) GetPromotionName() string {
	if o == nil || IsNil(o.PromotionName) {
		var ret string
		return ret
	}
	return *o.PromotionName
}

// GetPromotionNameOk returns a tuple with the PromotionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionTaskRelation) GetPromotionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionName) {
		return nil, false
	}
	return o.PromotionName, true
}

// HasPromotionName returns a boolean if a field has been set.
func (o *PromotionTaskRelation) HasPromotionName() bool {
	if o != nil && !IsNil(o.PromotionName) {
		return true
	}

	return false
}

// SetPromotionName gets a reference to the given string and assigns it to the PromotionName field.
func (o *PromotionTaskRelation) SetPromotionName(v string) {
	o.PromotionName = &v
}

func (o PromotionTaskRelation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionTaskRelation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyReason) {
		toSerialize["apply_reason"] = o.ApplyReason
	}
	if !IsNil(o.CommodityId) {
		toSerialize["commodity_id"] = o.CommodityId
	}
	if !IsNil(o.CommodityName) {
		toSerialize["commodity_name"] = o.CommodityName
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.PromoterContactName) {
		toSerialize["promoter_contact_name"] = o.PromoterContactName
	}
	if !IsNil(o.PromoterContactPhone) {
		toSerialize["promoter_contact_phone"] = o.PromoterContactPhone
	}
	if !IsNil(o.PromoterName) {
		toSerialize["promoter_name"] = o.PromoterName
	}
	if !IsNil(o.PromoterPid) {
		toSerialize["promoter_pid"] = o.PromoterPid
	}
	if !IsNil(o.PromotionId) {
		toSerialize["promotion_id"] = o.PromotionId
	}
	if !IsNil(o.PromotionName) {
		toSerialize["promotion_name"] = o.PromotionName
	}
	return toSerialize, nil
}

type NullablePromotionTaskRelation struct {
	value *PromotionTaskRelation
	isSet bool
}

func (v NullablePromotionTaskRelation) Get() *PromotionTaskRelation {
	return v.value
}

func (v *NullablePromotionTaskRelation) Set(val *PromotionTaskRelation) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionTaskRelation) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionTaskRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionTaskRelation(val *PromotionTaskRelation) *NullablePromotionTaskRelation {
	return &NullablePromotionTaskRelation{value: val, isSet: true}
}

func (v NullablePromotionTaskRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionTaskRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
