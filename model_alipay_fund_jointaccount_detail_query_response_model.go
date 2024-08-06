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

// checks if the AlipayFundJointaccountDetailQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundJointaccountDetailQueryResponseModel{}

// AlipayFundJointaccountDetailQueryResponseModel struct for AlipayFundJointaccountDetailQueryResponseModel
type AlipayFundJointaccountDetailQueryResponseModel struct {
	// 合花群ID（支付宝侧生成）
	AccountId *string `json:"account_id,omitempty"`
	// 账户名称（支付宝侧生成）
	AccountName *string `json:"account_name,omitempty"`
	// 额度模型
	AccountQuota []JointAccountQuotaRespDTO `json:"account_quota,omitempty"`
	// 账户状态<br> -NORMAL：正常<br> -RELEASING：注销中<br> -RELEASED：已注销<br> - FREEZE：冻结
	AccountStatus *string `json:"account_status,omitempty"`
	// 授权协议号（支付宝侧生成）
	AgreementNo *string `json:"agreement_no,omitempty"`
	AuthorizedRule *AuthorizedRuleDTO `json:"authorized_rule,omitempty"`
	// 当前可用金额（单位为元，必须大于0且最多小数点后两位）
	AvailableBalance *string `json:"available_balance,omitempty"`
	// 业务场景码
	BizScene *string `json:"biz_scene,omitempty"`
	// （创建人）支付宝侧用户唯一标识
	CreatorId *string `json:"creator_id,omitempty"`
	// （创建人）支付宝侧用户唯一标识
	CreatorOpenId *string `json:"creator_open_id,omitempty"`
	// （创建人）商户侧用户唯一标识<br> 补充说明：<br> -如果签约时，发起人标识传递的是商户侧用户唯一标识，则该字段会返回<br> -如果签约时，发起人标识传递的是支付宝侧用户唯一标识，则该字段为空
	CreatorOutId *string `json:"creator_out_id,omitempty"`
	// 当前冻结金额（单位为元，必须大于0且最多小数点后两位）
	FreezeBalance *string `json:"freeze_balance,omitempty"`
	// 签约时邀请的成员列表（快照）
	InviteResultList []InviteResultDTO `json:"invite_result_list,omitempty"`
	// 已加入合花群的成员列表
	MemberList []JointAccountMemberInfoRespDTO `json:"member_list,omitempty"`
	// 销售产品码
	ProductCode *string `json:"product_code,omitempty"`
	// 生息状态：</br> - MAKING    ：开启中</br> - MADE.       ：已开启</br> - CLEARING：关闭中</br> - NONE.       ：未生息
	ProfitStatus *string `json:"profit_status,omitempty"`
	// 昨日收益
	YesterdayProfit *string `json:"yesterday_profit,omitempty"`
}

// NewAlipayFundJointaccountDetailQueryResponseModel instantiates a new AlipayFundJointaccountDetailQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundJointaccountDetailQueryResponseModel() *AlipayFundJointaccountDetailQueryResponseModel {
	this := AlipayFundJointaccountDetailQueryResponseModel{}
	return &this
}

// NewAlipayFundJointaccountDetailQueryResponseModelWithDefaults instantiates a new AlipayFundJointaccountDetailQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundJointaccountDetailQueryResponseModelWithDefaults() *AlipayFundJointaccountDetailQueryResponseModel {
	this := AlipayFundJointaccountDetailQueryResponseModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountQuota returns the AccountQuota field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountQuota() []JointAccountQuotaRespDTO {
	if o == nil || IsNil(o.AccountQuota) {
		var ret []JointAccountQuotaRespDTO
		return ret
	}
	return o.AccountQuota
}

// GetAccountQuotaOk returns a tuple with the AccountQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountQuotaOk() ([]JointAccountQuotaRespDTO, bool) {
	if o == nil || IsNil(o.AccountQuota) {
		return nil, false
	}
	return o.AccountQuota, true
}

// HasAccountQuota returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasAccountQuota() bool {
	if o != nil && !IsNil(o.AccountQuota) {
		return true
	}

	return false
}

// SetAccountQuota gets a reference to the given []JointAccountQuotaRespDTO and assigns it to the AccountQuota field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetAccountQuota(v []JointAccountQuotaRespDTO) {
	o.AccountQuota = v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountStatus() string {
	if o == nil || IsNil(o.AccountStatus) {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAccountStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasAccountStatus() bool {
	if o != nil && !IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetAuthorizedRule returns the AuthorizedRule field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAuthorizedRule() AuthorizedRuleDTO {
	if o == nil || IsNil(o.AuthorizedRule) {
		var ret AuthorizedRuleDTO
		return ret
	}
	return *o.AuthorizedRule
}

// GetAuthorizedRuleOk returns a tuple with the AuthorizedRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAuthorizedRuleOk() (*AuthorizedRuleDTO, bool) {
	if o == nil || IsNil(o.AuthorizedRule) {
		return nil, false
	}
	return o.AuthorizedRule, true
}

// HasAuthorizedRule returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasAuthorizedRule() bool {
	if o != nil && !IsNil(o.AuthorizedRule) {
		return true
	}

	return false
}

// SetAuthorizedRule gets a reference to the given AuthorizedRuleDTO and assigns it to the AuthorizedRule field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetAuthorizedRule(v AuthorizedRuleDTO) {
	o.AuthorizedRule = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAvailableBalance() string {
	if o == nil || IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasAvailableBalance() bool {
	if o != nil && !IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetCreatorOpenId returns the CreatorOpenId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetCreatorOpenId() string {
	if o == nil || IsNil(o.CreatorOpenId) {
		var ret string
		return ret
	}
	return *o.CreatorOpenId
}

// GetCreatorOpenIdOk returns a tuple with the CreatorOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetCreatorOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorOpenId) {
		return nil, false
	}
	return o.CreatorOpenId, true
}

// HasCreatorOpenId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasCreatorOpenId() bool {
	if o != nil && !IsNil(o.CreatorOpenId) {
		return true
	}

	return false
}

// SetCreatorOpenId gets a reference to the given string and assigns it to the CreatorOpenId field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetCreatorOpenId(v string) {
	o.CreatorOpenId = &v
}

// GetCreatorOutId returns the CreatorOutId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetCreatorOutId() string {
	if o == nil || IsNil(o.CreatorOutId) {
		var ret string
		return ret
	}
	return *o.CreatorOutId
}

// GetCreatorOutIdOk returns a tuple with the CreatorOutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetCreatorOutIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorOutId) {
		return nil, false
	}
	return o.CreatorOutId, true
}

// HasCreatorOutId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasCreatorOutId() bool {
	if o != nil && !IsNil(o.CreatorOutId) {
		return true
	}

	return false
}

// SetCreatorOutId gets a reference to the given string and assigns it to the CreatorOutId field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetCreatorOutId(v string) {
	o.CreatorOutId = &v
}

// GetFreezeBalance returns the FreezeBalance field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetFreezeBalance() string {
	if o == nil || IsNil(o.FreezeBalance) {
		var ret string
		return ret
	}
	return *o.FreezeBalance
}

// GetFreezeBalanceOk returns a tuple with the FreezeBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetFreezeBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.FreezeBalance) {
		return nil, false
	}
	return o.FreezeBalance, true
}

// HasFreezeBalance returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasFreezeBalance() bool {
	if o != nil && !IsNil(o.FreezeBalance) {
		return true
	}

	return false
}

// SetFreezeBalance gets a reference to the given string and assigns it to the FreezeBalance field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetFreezeBalance(v string) {
	o.FreezeBalance = &v
}

// GetInviteResultList returns the InviteResultList field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetInviteResultList() []InviteResultDTO {
	if o == nil || IsNil(o.InviteResultList) {
		var ret []InviteResultDTO
		return ret
	}
	return o.InviteResultList
}

// GetInviteResultListOk returns a tuple with the InviteResultList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetInviteResultListOk() ([]InviteResultDTO, bool) {
	if o == nil || IsNil(o.InviteResultList) {
		return nil, false
	}
	return o.InviteResultList, true
}

// HasInviteResultList returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasInviteResultList() bool {
	if o != nil && !IsNil(o.InviteResultList) {
		return true
	}

	return false
}

// SetInviteResultList gets a reference to the given []InviteResultDTO and assigns it to the InviteResultList field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetInviteResultList(v []InviteResultDTO) {
	o.InviteResultList = v
}

// GetMemberList returns the MemberList field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetMemberList() []JointAccountMemberInfoRespDTO {
	if o == nil || IsNil(o.MemberList) {
		var ret []JointAccountMemberInfoRespDTO
		return ret
	}
	return o.MemberList
}

// GetMemberListOk returns a tuple with the MemberList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetMemberListOk() ([]JointAccountMemberInfoRespDTO, bool) {
	if o == nil || IsNil(o.MemberList) {
		return nil, false
	}
	return o.MemberList, true
}

// HasMemberList returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasMemberList() bool {
	if o != nil && !IsNil(o.MemberList) {
		return true
	}

	return false
}

// SetMemberList gets a reference to the given []JointAccountMemberInfoRespDTO and assigns it to the MemberList field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetMemberList(v []JointAccountMemberInfoRespDTO) {
	o.MemberList = v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProfitStatus returns the ProfitStatus field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetProfitStatus() string {
	if o == nil || IsNil(o.ProfitStatus) {
		var ret string
		return ret
	}
	return *o.ProfitStatus
}

// GetProfitStatusOk returns a tuple with the ProfitStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetProfitStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitStatus) {
		return nil, false
	}
	return o.ProfitStatus, true
}

// HasProfitStatus returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasProfitStatus() bool {
	if o != nil && !IsNil(o.ProfitStatus) {
		return true
	}

	return false
}

// SetProfitStatus gets a reference to the given string and assigns it to the ProfitStatus field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetProfitStatus(v string) {
	o.ProfitStatus = &v
}

// GetYesterdayProfit returns the YesterdayProfit field value if set, zero value otherwise.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetYesterdayProfit() string {
	if o == nil || IsNil(o.YesterdayProfit) {
		var ret string
		return ret
	}
	return *o.YesterdayProfit
}

// GetYesterdayProfitOk returns a tuple with the YesterdayProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) GetYesterdayProfitOk() (*string, bool) {
	if o == nil || IsNil(o.YesterdayProfit) {
		return nil, false
	}
	return o.YesterdayProfit, true
}

// HasYesterdayProfit returns a boolean if a field has been set.
func (o *AlipayFundJointaccountDetailQueryResponseModel) HasYesterdayProfit() bool {
	if o != nil && !IsNil(o.YesterdayProfit) {
		return true
	}

	return false
}

// SetYesterdayProfit gets a reference to the given string and assigns it to the YesterdayProfit field.
func (o *AlipayFundJointaccountDetailQueryResponseModel) SetYesterdayProfit(v string) {
	o.YesterdayProfit = &v
}

func (o AlipayFundJointaccountDetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundJointaccountDetailQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.AccountQuota) {
		toSerialize["account_quota"] = o.AccountQuota
	}
	if !IsNil(o.AccountStatus) {
		toSerialize["account_status"] = o.AccountStatus
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.AuthorizedRule) {
		toSerialize["authorized_rule"] = o.AuthorizedRule
	}
	if !IsNil(o.AvailableBalance) {
		toSerialize["available_balance"] = o.AvailableBalance
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if !IsNil(o.CreatorOpenId) {
		toSerialize["creator_open_id"] = o.CreatorOpenId
	}
	if !IsNil(o.CreatorOutId) {
		toSerialize["creator_out_id"] = o.CreatorOutId
	}
	if !IsNil(o.FreezeBalance) {
		toSerialize["freeze_balance"] = o.FreezeBalance
	}
	if !IsNil(o.InviteResultList) {
		toSerialize["invite_result_list"] = o.InviteResultList
	}
	if !IsNil(o.MemberList) {
		toSerialize["member_list"] = o.MemberList
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.ProfitStatus) {
		toSerialize["profit_status"] = o.ProfitStatus
	}
	if !IsNil(o.YesterdayProfit) {
		toSerialize["yesterday_profit"] = o.YesterdayProfit
	}
	return toSerialize, nil
}

type NullableAlipayFundJointaccountDetailQueryResponseModel struct {
	value *AlipayFundJointaccountDetailQueryResponseModel
	isSet bool
}

func (v NullableAlipayFundJointaccountDetailQueryResponseModel) Get() *AlipayFundJointaccountDetailQueryResponseModel {
	return v.value
}

func (v *NullableAlipayFundJointaccountDetailQueryResponseModel) Set(val *AlipayFundJointaccountDetailQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountDetailQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountDetailQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountDetailQueryResponseModel(val *AlipayFundJointaccountDetailQueryResponseModel) *NullableAlipayFundJointaccountDetailQueryResponseModel {
	return &NullableAlipayFundJointaccountDetailQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountDetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountDetailQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


