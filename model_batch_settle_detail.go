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

// checks if the BatchSettleDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchSettleDetail{}

// BatchSettleDetail struct for BatchSettleDetail
type BatchSettleDetail struct {
	// 结算金额，单位为元
	Amount *string `json:"amount,omitempty"`
	// 结算币种
	Currency *string `json:"currency,omitempty"`
	// 错误编码，SETTLE_ACCOUNT_ERROR：结算账户信息有误；BANK_DISHONOR：银行提现退票；UNKNOWN_ERROR：未知错误
	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述
	ErrorDesc *string `json:"error_desc,omitempty"`
	// 结算账户id。当结算账户id类型是cardSerialNo时，本参数为用户在支付宝绑定的卡编号；当结算账户id类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号；当结算账户id类型是loginName时，本参数为用户的支付宝登录号
	SettleAccountId *string `json:"settle_account_id,omitempty"`
	// 结算账户id类型。  当settle_account_type 为bankCard时，本参数为cardSerialNo，表示结算账户id是银行卡编号;  当settle_account_type 为alipayBalance时，本参数为userId或者loginName，其中userId表示结算账户id是支付宝唯一用户号，loginName表示结算账户id是支付宝登录号
	SettleAccountIdType *string `json:"settle_account_id_type,omitempty"`
	// 结算账户的OpenId，本参数是用户在该应用（AppId）下的唯一用户标识。
	SettleAccountOpenId *string `json:"settle_account_open_id,omitempty"`
	// 结算账户类型。 bankCard: 结算账户为银行卡； alipayBalance: 结算账户为支付宝余额户
	SettleAccountType *string `json:"settle_account_type,omitempty"`
	// 结算主体账号。 当结算主体类型为SecondMerchant，本参数为二级商户的SecondMerchantID
	SettleEntityId *string `json:"settle_entity_id,omitempty"`
	// 结算主体类型。 SecondMerchant：结算主体为二级商户; Store：结算主体为门店；
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
	// ACCEPT_SUCCESS：受理成功； SUCCESS：结算成功； FAIL：结算失败；FAIL_RETRY：失败重试。
	Status *string `json:"status,omitempty"`
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
}

// NewBatchSettleDetail instantiates a new BatchSettleDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchSettleDetail() *BatchSettleDetail {
	this := BatchSettleDetail{}
	return &this
}

// NewBatchSettleDetailWithDefaults instantiates a new BatchSettleDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchSettleDetailWithDefaults() *BatchSettleDetail {
	this := BatchSettleDetail{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *BatchSettleDetail) SetAmount(v string) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BatchSettleDetail) SetCurrency(v string) {
	o.Currency = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *BatchSettleDetail) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorDesc returns the ErrorDesc field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetErrorDesc() string {
	if o == nil || IsNil(o.ErrorDesc) {
		var ret string
		return ret
	}
	return *o.ErrorDesc
}

// GetErrorDescOk returns a tuple with the ErrorDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetErrorDescOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDesc) {
		return nil, false
	}
	return o.ErrorDesc, true
}

// HasErrorDesc returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasErrorDesc() bool {
	if o != nil && !IsNil(o.ErrorDesc) {
		return true
	}

	return false
}

// SetErrorDesc gets a reference to the given string and assigns it to the ErrorDesc field.
func (o *BatchSettleDetail) SetErrorDesc(v string) {
	o.ErrorDesc = &v
}

// GetSettleAccountId returns the SettleAccountId field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetSettleAccountId() string {
	if o == nil || IsNil(o.SettleAccountId) {
		var ret string
		return ret
	}
	return *o.SettleAccountId
}

// GetSettleAccountIdOk returns a tuple with the SettleAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetSettleAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SettleAccountId) {
		return nil, false
	}
	return o.SettleAccountId, true
}

// HasSettleAccountId returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasSettleAccountId() bool {
	if o != nil && !IsNil(o.SettleAccountId) {
		return true
	}

	return false
}

// SetSettleAccountId gets a reference to the given string and assigns it to the SettleAccountId field.
func (o *BatchSettleDetail) SetSettleAccountId(v string) {
	o.SettleAccountId = &v
}

// GetSettleAccountIdType returns the SettleAccountIdType field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetSettleAccountIdType() string {
	if o == nil || IsNil(o.SettleAccountIdType) {
		var ret string
		return ret
	}
	return *o.SettleAccountIdType
}

// GetSettleAccountIdTypeOk returns a tuple with the SettleAccountIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetSettleAccountIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleAccountIdType) {
		return nil, false
	}
	return o.SettleAccountIdType, true
}

// HasSettleAccountIdType returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasSettleAccountIdType() bool {
	if o != nil && !IsNil(o.SettleAccountIdType) {
		return true
	}

	return false
}

// SetSettleAccountIdType gets a reference to the given string and assigns it to the SettleAccountIdType field.
func (o *BatchSettleDetail) SetSettleAccountIdType(v string) {
	o.SettleAccountIdType = &v
}

// GetSettleAccountOpenId returns the SettleAccountOpenId field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetSettleAccountOpenId() string {
	if o == nil || IsNil(o.SettleAccountOpenId) {
		var ret string
		return ret
	}
	return *o.SettleAccountOpenId
}

// GetSettleAccountOpenIdOk returns a tuple with the SettleAccountOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetSettleAccountOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.SettleAccountOpenId) {
		return nil, false
	}
	return o.SettleAccountOpenId, true
}

// HasSettleAccountOpenId returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasSettleAccountOpenId() bool {
	if o != nil && !IsNil(o.SettleAccountOpenId) {
		return true
	}

	return false
}

// SetSettleAccountOpenId gets a reference to the given string and assigns it to the SettleAccountOpenId field.
func (o *BatchSettleDetail) SetSettleAccountOpenId(v string) {
	o.SettleAccountOpenId = &v
}

// GetSettleAccountType returns the SettleAccountType field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetSettleAccountType() string {
	if o == nil || IsNil(o.SettleAccountType) {
		var ret string
		return ret
	}
	return *o.SettleAccountType
}

// GetSettleAccountTypeOk returns a tuple with the SettleAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetSettleAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleAccountType) {
		return nil, false
	}
	return o.SettleAccountType, true
}

// HasSettleAccountType returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasSettleAccountType() bool {
	if o != nil && !IsNil(o.SettleAccountType) {
		return true
	}

	return false
}

// SetSettleAccountType gets a reference to the given string and assigns it to the SettleAccountType field.
func (o *BatchSettleDetail) SetSettleAccountType(v string) {
	o.SettleAccountType = &v
}

// GetSettleEntityId returns the SettleEntityId field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetSettleEntityId() string {
	if o == nil || IsNil(o.SettleEntityId) {
		var ret string
		return ret
	}
	return *o.SettleEntityId
}

// GetSettleEntityIdOk returns a tuple with the SettleEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetSettleEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.SettleEntityId) {
		return nil, false
	}
	return o.SettleEntityId, true
}

// HasSettleEntityId returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasSettleEntityId() bool {
	if o != nil && !IsNil(o.SettleEntityId) {
		return true
	}

	return false
}

// SetSettleEntityId gets a reference to the given string and assigns it to the SettleEntityId field.
func (o *BatchSettleDetail) SetSettleEntityId(v string) {
	o.SettleEntityId = &v
}

// GetSettleEntityType returns the SettleEntityType field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetSettleEntityType() string {
	if o == nil || IsNil(o.SettleEntityType) {
		var ret string
		return ret
	}
	return *o.SettleEntityType
}

// GetSettleEntityTypeOk returns a tuple with the SettleEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetSettleEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleEntityType) {
		return nil, false
	}
	return o.SettleEntityType, true
}

// HasSettleEntityType returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasSettleEntityType() bool {
	if o != nil && !IsNil(o.SettleEntityType) {
		return true
	}

	return false
}

// SetSettleEntityType gets a reference to the given string and assigns it to the SettleEntityType field.
func (o *BatchSettleDetail) SetSettleEntityType(v string) {
	o.SettleEntityType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BatchSettleDetail) SetStatus(v string) {
	o.Status = &v
}

// GetSubMerchant returns the SubMerchant field value if set, zero value otherwise.
func (o *BatchSettleDetail) GetSubMerchant() SubMerchant {
	if o == nil || IsNil(o.SubMerchant) {
		var ret SubMerchant
		return ret
	}
	return *o.SubMerchant
}

// GetSubMerchantOk returns a tuple with the SubMerchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettleDetail) GetSubMerchantOk() (*SubMerchant, bool) {
	if o == nil || IsNil(o.SubMerchant) {
		return nil, false
	}
	return o.SubMerchant, true
}

// HasSubMerchant returns a boolean if a field has been set.
func (o *BatchSettleDetail) HasSubMerchant() bool {
	if o != nil && !IsNil(o.SubMerchant) {
		return true
	}

	return false
}

// SetSubMerchant gets a reference to the given SubMerchant and assigns it to the SubMerchant field.
func (o *BatchSettleDetail) SetSubMerchant(v SubMerchant) {
	o.SubMerchant = &v
}

func (o BatchSettleDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchSettleDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.ErrorDesc) {
		toSerialize["error_desc"] = o.ErrorDesc
	}
	if !IsNil(o.SettleAccountId) {
		toSerialize["settle_account_id"] = o.SettleAccountId
	}
	if !IsNil(o.SettleAccountIdType) {
		toSerialize["settle_account_id_type"] = o.SettleAccountIdType
	}
	if !IsNil(o.SettleAccountOpenId) {
		toSerialize["settle_account_open_id"] = o.SettleAccountOpenId
	}
	if !IsNil(o.SettleAccountType) {
		toSerialize["settle_account_type"] = o.SettleAccountType
	}
	if !IsNil(o.SettleEntityId) {
		toSerialize["settle_entity_id"] = o.SettleEntityId
	}
	if !IsNil(o.SettleEntityType) {
		toSerialize["settle_entity_type"] = o.SettleEntityType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubMerchant) {
		toSerialize["sub_merchant"] = o.SubMerchant
	}
	return toSerialize, nil
}

type NullableBatchSettleDetail struct {
	value *BatchSettleDetail
	isSet bool
}

func (v NullableBatchSettleDetail) Get() *BatchSettleDetail {
	return v.value
}

func (v *NullableBatchSettleDetail) Set(val *BatchSettleDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchSettleDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchSettleDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchSettleDetail(val *BatchSettleDetail) *NullableBatchSettleDetail {
	return &NullableBatchSettleDetail{value: val, isSet: true}
}

func (v NullableBatchSettleDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchSettleDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

