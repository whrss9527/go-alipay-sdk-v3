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

// checks if the JointAccountBillDetailDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JointAccountBillDetailDTO{}

// JointAccountBillDetailDTO struct for JointAccountBillDetailDTO
type JointAccountBillDetailDTO struct {
	// 共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 消费金额
	Amount *string `json:"amount,omitempty"`
	// 账单业务号
	BillNo *string `json:"bill_no,omitempty"`
	// 业务时间
	BizDate *string `json:"biz_date,omitempty"`
	// 订单号
	BizNo *string `json:"biz_no,omitempty"`
	// 1-退款，2-支付
	InOut *string `json:"in_out,omitempty"`
	// 共同账户消费的成员openid
	OpenId *string `json:"open_id,omitempty"`
	// 外部交易单号，正向支付为外部交易单号，逆向退款是为外部退款单号
	OutTradeNo     *string          `json:"out_trade_no,omitempty"`
	PayerAssetInfo *UserAssetInfoVO `json:"payer_asset_info,omitempty"`
	// 间连商户显示二级商户全名，直连、直付通则显示一级商户全名。企业商户该字段不脱敏，非企业商户该字段会脱敏
	SellerFullName *string `json:"seller_full_name,omitempty"`
	// 收款方登录号信息。间连商户显示二级商户的登录号信息；直连、直付通显示一级商户的登录号信息；该字段脱敏
	SellerLogonId *string `json:"seller_logon_id,omitempty"`
	// 账单标题
	Title *string `json:"title,omitempty"`
	// 共同账户消费的成员ID
	UserId *string `json:"user_id,omitempty"`
}

// NewJointAccountBillDetailDTO instantiates a new JointAccountBillDetailDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJointAccountBillDetailDTO() *JointAccountBillDetailDTO {
	this := JointAccountBillDetailDTO{}
	return &this
}

// NewJointAccountBillDetailDTOWithDefaults instantiates a new JointAccountBillDetailDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJointAccountBillDetailDTOWithDefaults() *JointAccountBillDetailDTO {
	this := JointAccountBillDetailDTO{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *JointAccountBillDetailDTO) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *JointAccountBillDetailDTO) SetAmount(v string) {
	o.Amount = &v
}

// GetBillNo returns the BillNo field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetBillNo() string {
	if o == nil || IsNil(o.BillNo) {
		var ret string
		return ret
	}
	return *o.BillNo
}

// GetBillNoOk returns a tuple with the BillNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetBillNoOk() (*string, bool) {
	if o == nil || IsNil(o.BillNo) {
		return nil, false
	}
	return o.BillNo, true
}

// HasBillNo returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasBillNo() bool {
	if o != nil && !IsNil(o.BillNo) {
		return true
	}

	return false
}

// SetBillNo gets a reference to the given string and assigns it to the BillNo field.
func (o *JointAccountBillDetailDTO) SetBillNo(v string) {
	o.BillNo = &v
}

// GetBizDate returns the BizDate field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetBizDate() string {
	if o == nil || IsNil(o.BizDate) {
		var ret string
		return ret
	}
	return *o.BizDate
}

// GetBizDateOk returns a tuple with the BizDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetBizDateOk() (*string, bool) {
	if o == nil || IsNil(o.BizDate) {
		return nil, false
	}
	return o.BizDate, true
}

// HasBizDate returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasBizDate() bool {
	if o != nil && !IsNil(o.BizDate) {
		return true
	}

	return false
}

// SetBizDate gets a reference to the given string and assigns it to the BizDate field.
func (o *JointAccountBillDetailDTO) SetBizDate(v string) {
	o.BizDate = &v
}

// GetBizNo returns the BizNo field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetBizNo() string {
	if o == nil || IsNil(o.BizNo) {
		var ret string
		return ret
	}
	return *o.BizNo
}

// GetBizNoOk returns a tuple with the BizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.BizNo) {
		return nil, false
	}
	return o.BizNo, true
}

// HasBizNo returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasBizNo() bool {
	if o != nil && !IsNil(o.BizNo) {
		return true
	}

	return false
}

// SetBizNo gets a reference to the given string and assigns it to the BizNo field.
func (o *JointAccountBillDetailDTO) SetBizNo(v string) {
	o.BizNo = &v
}

// GetInOut returns the InOut field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetInOut() string {
	if o == nil || IsNil(o.InOut) {
		var ret string
		return ret
	}
	return *o.InOut
}

// GetInOutOk returns a tuple with the InOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetInOutOk() (*string, bool) {
	if o == nil || IsNil(o.InOut) {
		return nil, false
	}
	return o.InOut, true
}

// HasInOut returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasInOut() bool {
	if o != nil && !IsNil(o.InOut) {
		return true
	}

	return false
}

// SetInOut gets a reference to the given string and assigns it to the InOut field.
func (o *JointAccountBillDetailDTO) SetInOut(v string) {
	o.InOut = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *JointAccountBillDetailDTO) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOutTradeNo returns the OutTradeNo field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetOutTradeNo() string {
	if o == nil || IsNil(o.OutTradeNo) {
		var ret string
		return ret
	}
	return *o.OutTradeNo
}

// GetOutTradeNoOk returns a tuple with the OutTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetOutTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutTradeNo) {
		return nil, false
	}
	return o.OutTradeNo, true
}

// HasOutTradeNo returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasOutTradeNo() bool {
	if o != nil && !IsNil(o.OutTradeNo) {
		return true
	}

	return false
}

// SetOutTradeNo gets a reference to the given string and assigns it to the OutTradeNo field.
func (o *JointAccountBillDetailDTO) SetOutTradeNo(v string) {
	o.OutTradeNo = &v
}

// GetPayerAssetInfo returns the PayerAssetInfo field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetPayerAssetInfo() UserAssetInfoVO {
	if o == nil || IsNil(o.PayerAssetInfo) {
		var ret UserAssetInfoVO
		return ret
	}
	return *o.PayerAssetInfo
}

// GetPayerAssetInfoOk returns a tuple with the PayerAssetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetPayerAssetInfoOk() (*UserAssetInfoVO, bool) {
	if o == nil || IsNil(o.PayerAssetInfo) {
		return nil, false
	}
	return o.PayerAssetInfo, true
}

// HasPayerAssetInfo returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasPayerAssetInfo() bool {
	if o != nil && !IsNil(o.PayerAssetInfo) {
		return true
	}

	return false
}

// SetPayerAssetInfo gets a reference to the given UserAssetInfoVO and assigns it to the PayerAssetInfo field.
func (o *JointAccountBillDetailDTO) SetPayerAssetInfo(v UserAssetInfoVO) {
	o.PayerAssetInfo = &v
}

// GetSellerFullName returns the SellerFullName field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetSellerFullName() string {
	if o == nil || IsNil(o.SellerFullName) {
		var ret string
		return ret
	}
	return *o.SellerFullName
}

// GetSellerFullNameOk returns a tuple with the SellerFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetSellerFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.SellerFullName) {
		return nil, false
	}
	return o.SellerFullName, true
}

// HasSellerFullName returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasSellerFullName() bool {
	if o != nil && !IsNil(o.SellerFullName) {
		return true
	}

	return false
}

// SetSellerFullName gets a reference to the given string and assigns it to the SellerFullName field.
func (o *JointAccountBillDetailDTO) SetSellerFullName(v string) {
	o.SellerFullName = &v
}

// GetSellerLogonId returns the SellerLogonId field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetSellerLogonId() string {
	if o == nil || IsNil(o.SellerLogonId) {
		var ret string
		return ret
	}
	return *o.SellerLogonId
}

// GetSellerLogonIdOk returns a tuple with the SellerLogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetSellerLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellerLogonId) {
		return nil, false
	}
	return o.SellerLogonId, true
}

// HasSellerLogonId returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasSellerLogonId() bool {
	if o != nil && !IsNil(o.SellerLogonId) {
		return true
	}

	return false
}

// SetSellerLogonId gets a reference to the given string and assigns it to the SellerLogonId field.
func (o *JointAccountBillDetailDTO) SetSellerLogonId(v string) {
	o.SellerLogonId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *JointAccountBillDetailDTO) SetTitle(v string) {
	o.Title = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *JointAccountBillDetailDTO) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountBillDetailDTO) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *JointAccountBillDetailDTO) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *JointAccountBillDetailDTO) SetUserId(v string) {
	o.UserId = &v
}

func (o JointAccountBillDetailDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JointAccountBillDetailDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BillNo) {
		toSerialize["bill_no"] = o.BillNo
	}
	if !IsNil(o.BizDate) {
		toSerialize["biz_date"] = o.BizDate
	}
	if !IsNil(o.BizNo) {
		toSerialize["biz_no"] = o.BizNo
	}
	if !IsNil(o.InOut) {
		toSerialize["in_out"] = o.InOut
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OutTradeNo) {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if !IsNil(o.PayerAssetInfo) {
		toSerialize["payer_asset_info"] = o.PayerAssetInfo
	}
	if !IsNil(o.SellerFullName) {
		toSerialize["seller_full_name"] = o.SellerFullName
	}
	if !IsNil(o.SellerLogonId) {
		toSerialize["seller_logon_id"] = o.SellerLogonId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableJointAccountBillDetailDTO struct {
	value *JointAccountBillDetailDTO
	isSet bool
}

func (v NullableJointAccountBillDetailDTO) Get() *JointAccountBillDetailDTO {
	return v.value
}

func (v *NullableJointAccountBillDetailDTO) Set(val *JointAccountBillDetailDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableJointAccountBillDetailDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableJointAccountBillDetailDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJointAccountBillDetailDTO(val *JointAccountBillDetailDTO) *NullableJointAccountBillDetailDTO {
	return &NullableJointAccountBillDetailDTO{value: val, isSet: true}
}

func (v NullableJointAccountBillDetailDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJointAccountBillDetailDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
