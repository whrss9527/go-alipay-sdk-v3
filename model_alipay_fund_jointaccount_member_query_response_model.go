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

// checks if the AlipayFundJointaccountMemberQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundJointaccountMemberQueryResponseModel{}

// AlipayFundJointaccountMemberQueryResponseModel struct for AlipayFundJointaccountMemberQueryResponseModel
type AlipayFundJointaccountMemberQueryResponseModel struct {
	// 共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 场景码
	BizScene *string `json:"biz_scene,omitempty"`
	// 成员列表
	MemberList []JointAccountMemberRespDTO `json:"member_list,omitempty"`
	// 页码
	PageNum *string `json:"page_num,omitempty"`
	// 页面记录数
	PageSize *string `json:"page_size,omitempty"`
	// 产品码
	ProductCode *string `json:"product_code,omitempty"`
	// 总页数
	TotalPageCount *string `json:"total_page_count,omitempty"`
}

// NewAlipayFundJointaccountMemberQueryResponseModel instantiates a new AlipayFundJointaccountMemberQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundJointaccountMemberQueryResponseModel() *AlipayFundJointaccountMemberQueryResponseModel {
	this := AlipayFundJointaccountMemberQueryResponseModel{}
	return &this
}

// NewAlipayFundJointaccountMemberQueryResponseModelWithDefaults instantiates a new AlipayFundJointaccountMemberQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundJointaccountMemberQueryResponseModelWithDefaults() *AlipayFundJointaccountMemberQueryResponseModel {
	this := AlipayFundJointaccountMemberQueryResponseModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayFundJointaccountMemberQueryResponseModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundJointaccountMemberQueryResponseModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetMemberList returns the MemberList field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetMemberList() []JointAccountMemberRespDTO {
	if o == nil || IsNil(o.MemberList) {
		var ret []JointAccountMemberRespDTO
		return ret
	}
	return o.MemberList
}

// GetMemberListOk returns a tuple with the MemberList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetMemberListOk() ([]JointAccountMemberRespDTO, bool) {
	if o == nil || IsNil(o.MemberList) {
		return nil, false
	}
	return o.MemberList, true
}

// HasMemberList returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) HasMemberList() bool {
	if o != nil && !IsNil(o.MemberList) {
		return true
	}

	return false
}

// SetMemberList gets a reference to the given []JointAccountMemberRespDTO and assigns it to the MemberList field.
func (o *AlipayFundJointaccountMemberQueryResponseModel) SetMemberList(v []JointAccountMemberRespDTO) {
	o.MemberList = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetPageNum() string {
	if o == nil || IsNil(o.PageNum) {
		var ret string
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetPageNumOk() (*string, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given string and assigns it to the PageNum field.
func (o *AlipayFundJointaccountMemberQueryResponseModel) SetPageNum(v string) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetPageSize() string {
	if o == nil || IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetPageSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *AlipayFundJointaccountMemberQueryResponseModel) SetPageSize(v string) {
	o.PageSize = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundJointaccountMemberQueryResponseModel) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetTotalPageCount returns the TotalPageCount field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetTotalPageCount() string {
	if o == nil || IsNil(o.TotalPageCount) {
		var ret string
		return ret
	}
	return *o.TotalPageCount
}

// GetTotalPageCountOk returns a tuple with the TotalPageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) GetTotalPageCountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalPageCount) {
		return nil, false
	}
	return o.TotalPageCount, true
}

// HasTotalPageCount returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberQueryResponseModel) HasTotalPageCount() bool {
	if o != nil && !IsNil(o.TotalPageCount) {
		return true
	}

	return false
}

// SetTotalPageCount gets a reference to the given string and assigns it to the TotalPageCount field.
func (o *AlipayFundJointaccountMemberQueryResponseModel) SetTotalPageCount(v string) {
	o.TotalPageCount = &v
}

func (o AlipayFundJointaccountMemberQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundJointaccountMemberQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.MemberList) {
		toSerialize["member_list"] = o.MemberList
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.TotalPageCount) {
		toSerialize["total_page_count"] = o.TotalPageCount
	}
	return toSerialize, nil
}

type NullableAlipayFundJointaccountMemberQueryResponseModel struct {
	value *AlipayFundJointaccountMemberQueryResponseModel
	isSet bool
}

func (v NullableAlipayFundJointaccountMemberQueryResponseModel) Get() *AlipayFundJointaccountMemberQueryResponseModel {
	return v.value
}

func (v *NullableAlipayFundJointaccountMemberQueryResponseModel) Set(val *AlipayFundJointaccountMemberQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountMemberQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountMemberQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountMemberQueryResponseModel(val *AlipayFundJointaccountMemberQueryResponseModel) *NullableAlipayFundJointaccountMemberQueryResponseModel {
	return &NullableAlipayFundJointaccountMemberQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountMemberQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountMemberQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


