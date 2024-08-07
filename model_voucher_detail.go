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

// checks if the VoucherDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherDetail{}

// VoucherDetail struct for VoucherDetail
type VoucherDetail struct {
	// 优惠券面额，它应该会等于商家出资加上其他出资方出资
	Amount *string `json:"amount,omitempty"`
	// 券id
	Id *string `json:"id,omitempty"`
	// 优惠券备注信息
	Memo *string `json:"memo,omitempty"`
	// 商家出资（特指发起交易的商家出资金额）
	MerchantContribute *string `json:"merchant_contribute,omitempty"`
	// 券名称
	Name *string `json:"name,omitempty"`
	// 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	OtherContribute *string `json:"other_contribute,omitempty"`
	// 优惠券的其他出资方明细
	OtherContributeDetail []ContributeDetail `json:"other_contribute_detail,omitempty"`
	// 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时平台优惠的金额
	PurchaseAntContribute *string `json:"purchase_ant_contribute,omitempty"`
	// 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时用户实际付款的金额
	PurchaseBuyerContribute *string `json:"purchase_buyer_contribute,omitempty"`
	// 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时商户优惠的金额
	PurchaseMerchantContribute *string `json:"purchase_merchant_contribute,omitempty"`
	// 券模板id
	TemplateId *string `json:"template_id,omitempty"`
	// 券类型
	Type *string `json:"type,omitempty"`
}

// NewVoucherDetail instantiates a new VoucherDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherDetail() *VoucherDetail {
	this := VoucherDetail{}
	return &this
}

// NewVoucherDetailWithDefaults instantiates a new VoucherDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherDetailWithDefaults() *VoucherDetail {
	this := VoucherDetail{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *VoucherDetail) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *VoucherDetail) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *VoucherDetail) SetAmount(v string) {
	o.Amount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VoucherDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VoucherDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VoucherDetail) SetId(v string) {
	o.Id = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *VoucherDetail) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *VoucherDetail) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *VoucherDetail) SetMemo(v string) {
	o.Memo = &v
}

// GetMerchantContribute returns the MerchantContribute field value if set, zero value otherwise.
func (o *VoucherDetail) GetMerchantContribute() string {
	if o == nil || IsNil(o.MerchantContribute) {
		var ret string
		return ret
	}
	return *o.MerchantContribute
}

// GetMerchantContributeOk returns a tuple with the MerchantContribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetMerchantContributeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantContribute) {
		return nil, false
	}
	return o.MerchantContribute, true
}

// HasMerchantContribute returns a boolean if a field has been set.
func (o *VoucherDetail) HasMerchantContribute() bool {
	if o != nil && !IsNil(o.MerchantContribute) {
		return true
	}

	return false
}

// SetMerchantContribute gets a reference to the given string and assigns it to the MerchantContribute field.
func (o *VoucherDetail) SetMerchantContribute(v string) {
	o.MerchantContribute = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VoucherDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VoucherDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VoucherDetail) SetName(v string) {
	o.Name = &v
}

// GetOtherContribute returns the OtherContribute field value if set, zero value otherwise.
func (o *VoucherDetail) GetOtherContribute() string {
	if o == nil || IsNil(o.OtherContribute) {
		var ret string
		return ret
	}
	return *o.OtherContribute
}

// GetOtherContributeOk returns a tuple with the OtherContribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetOtherContributeOk() (*string, bool) {
	if o == nil || IsNil(o.OtherContribute) {
		return nil, false
	}
	return o.OtherContribute, true
}

// HasOtherContribute returns a boolean if a field has been set.
func (o *VoucherDetail) HasOtherContribute() bool {
	if o != nil && !IsNil(o.OtherContribute) {
		return true
	}

	return false
}

// SetOtherContribute gets a reference to the given string and assigns it to the OtherContribute field.
func (o *VoucherDetail) SetOtherContribute(v string) {
	o.OtherContribute = &v
}

// GetOtherContributeDetail returns the OtherContributeDetail field value if set, zero value otherwise.
func (o *VoucherDetail) GetOtherContributeDetail() []ContributeDetail {
	if o == nil || IsNil(o.OtherContributeDetail) {
		var ret []ContributeDetail
		return ret
	}
	return o.OtherContributeDetail
}

// GetOtherContributeDetailOk returns a tuple with the OtherContributeDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetOtherContributeDetailOk() ([]ContributeDetail, bool) {
	if o == nil || IsNil(o.OtherContributeDetail) {
		return nil, false
	}
	return o.OtherContributeDetail, true
}

// HasOtherContributeDetail returns a boolean if a field has been set.
func (o *VoucherDetail) HasOtherContributeDetail() bool {
	if o != nil && !IsNil(o.OtherContributeDetail) {
		return true
	}

	return false
}

// SetOtherContributeDetail gets a reference to the given []ContributeDetail and assigns it to the OtherContributeDetail field.
func (o *VoucherDetail) SetOtherContributeDetail(v []ContributeDetail) {
	o.OtherContributeDetail = v
}

// GetPurchaseAntContribute returns the PurchaseAntContribute field value if set, zero value otherwise.
func (o *VoucherDetail) GetPurchaseAntContribute() string {
	if o == nil || IsNil(o.PurchaseAntContribute) {
		var ret string
		return ret
	}
	return *o.PurchaseAntContribute
}

// GetPurchaseAntContributeOk returns a tuple with the PurchaseAntContribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetPurchaseAntContributeOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseAntContribute) {
		return nil, false
	}
	return o.PurchaseAntContribute, true
}

// HasPurchaseAntContribute returns a boolean if a field has been set.
func (o *VoucherDetail) HasPurchaseAntContribute() bool {
	if o != nil && !IsNil(o.PurchaseAntContribute) {
		return true
	}

	return false
}

// SetPurchaseAntContribute gets a reference to the given string and assigns it to the PurchaseAntContribute field.
func (o *VoucherDetail) SetPurchaseAntContribute(v string) {
	o.PurchaseAntContribute = &v
}

// GetPurchaseBuyerContribute returns the PurchaseBuyerContribute field value if set, zero value otherwise.
func (o *VoucherDetail) GetPurchaseBuyerContribute() string {
	if o == nil || IsNil(o.PurchaseBuyerContribute) {
		var ret string
		return ret
	}
	return *o.PurchaseBuyerContribute
}

// GetPurchaseBuyerContributeOk returns a tuple with the PurchaseBuyerContribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetPurchaseBuyerContributeOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseBuyerContribute) {
		return nil, false
	}
	return o.PurchaseBuyerContribute, true
}

// HasPurchaseBuyerContribute returns a boolean if a field has been set.
func (o *VoucherDetail) HasPurchaseBuyerContribute() bool {
	if o != nil && !IsNil(o.PurchaseBuyerContribute) {
		return true
	}

	return false
}

// SetPurchaseBuyerContribute gets a reference to the given string and assigns it to the PurchaseBuyerContribute field.
func (o *VoucherDetail) SetPurchaseBuyerContribute(v string) {
	o.PurchaseBuyerContribute = &v
}

// GetPurchaseMerchantContribute returns the PurchaseMerchantContribute field value if set, zero value otherwise.
func (o *VoucherDetail) GetPurchaseMerchantContribute() string {
	if o == nil || IsNil(o.PurchaseMerchantContribute) {
		var ret string
		return ret
	}
	return *o.PurchaseMerchantContribute
}

// GetPurchaseMerchantContributeOk returns a tuple with the PurchaseMerchantContribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetPurchaseMerchantContributeOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseMerchantContribute) {
		return nil, false
	}
	return o.PurchaseMerchantContribute, true
}

// HasPurchaseMerchantContribute returns a boolean if a field has been set.
func (o *VoucherDetail) HasPurchaseMerchantContribute() bool {
	if o != nil && !IsNil(o.PurchaseMerchantContribute) {
		return true
	}

	return false
}

// SetPurchaseMerchantContribute gets a reference to the given string and assigns it to the PurchaseMerchantContribute field.
func (o *VoucherDetail) SetPurchaseMerchantContribute(v string) {
	o.PurchaseMerchantContribute = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *VoucherDetail) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *VoucherDetail) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *VoucherDetail) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VoucherDetail) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherDetail) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VoucherDetail) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VoucherDetail) SetType(v string) {
	o.Type = &v
}

func (o VoucherDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.MerchantContribute) {
		toSerialize["merchant_contribute"] = o.MerchantContribute
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OtherContribute) {
		toSerialize["other_contribute"] = o.OtherContribute
	}
	if !IsNil(o.OtherContributeDetail) {
		toSerialize["other_contribute_detail"] = o.OtherContributeDetail
	}
	if !IsNil(o.PurchaseAntContribute) {
		toSerialize["purchase_ant_contribute"] = o.PurchaseAntContribute
	}
	if !IsNil(o.PurchaseBuyerContribute) {
		toSerialize["purchase_buyer_contribute"] = o.PurchaseBuyerContribute
	}
	if !IsNil(o.PurchaseMerchantContribute) {
		toSerialize["purchase_merchant_contribute"] = o.PurchaseMerchantContribute
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableVoucherDetail struct {
	value *VoucherDetail
	isSet bool
}

func (v NullableVoucherDetail) Get() *VoucherDetail {
	return v.value
}

func (v *NullableVoucherDetail) Set(val *VoucherDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherDetail(val *VoucherDetail) *NullableVoucherDetail {
	return &NullableVoucherDetail{value: val, isSet: true}
}

func (v NullableVoucherDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
