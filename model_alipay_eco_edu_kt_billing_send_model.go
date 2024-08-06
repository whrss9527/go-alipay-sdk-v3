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

// checks if the AlipayEcoEduKtBillingSendModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoEduKtBillingSendModel{}

// AlipayEcoEduKtBillingSendModel struct for AlipayEcoEduKtBillingSendModel
type AlipayEcoEduKtBillingSendModel struct {
	// 总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]，  如果是非多选项，是要和缴费项的总和相同，多选模式不做验证
	Amount *string `json:"amount,omitempty"`
	// 缴费账单名称
	ChargeBillTitle *string `json:"charge_bill_title,omitempty"`
	// 缴费详情：输入json格式字符串。Json定义：key填写缴费项名称，value填写缴费项金额，金额保留2位小数（单位：元）
	ChargeItem []ChargeItems `json:"charge_item,omitempty"`
	// 缴费项模式：空或\"N\"，表示缴费项不可选，  \"M\"表示缴费项为可选 ，支持单选和多选。
	ChargeType *string `json:"charge_type,omitempty"`
	// 孩子名字
	ChildName *string `json:"child_name,omitempty"`
	// 孩子所在班级
	ClassIn *string `json:"class_in,omitempty"`
	// 截止日期是否生效，与gmt_end发布配合使用,N为gmt_end不生效，用户过期后仍可以缴费；Y为gmt_end生效，用户过期后，不能再缴费。
	EndEnable *string          `json:"end_enable,omitempty"`
	ExtInfo   *BillSendExtInfo `json:"ext_info,omitempty"`
	// 缴费截止时间，格式\"yyyy-MM-dd HH:mm:ss\"，日期要大于当前时间。请注意，过期时间不宜设置过短。
	GmtEnd *string `json:"gmt_end,omitempty"`
	// 孩子所在年级
	Grade *string `json:"grade,omitempty"`
	// ISV端的缴费账单编号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// Isv支付宝pid, 支付宝签约后，返回给ISV编号
	PartnerId *string `json:"partner_id,omitempty"`
	// 学校编码，录入学校接口返回的参数
	SchoolNo *string `json:"school_no,omitempty"`
	// 学校支付宝pid,直付通填写smid
	SchoolPid *string `json:"school_pid,omitempty"`
	// 学生的学号，只支持字母和数字类型，一般以教育局学号为准，作为学生的唯一标识。此字段与student_identify、家长user_mobile至少选一个
	StudentCode *string `json:"student_code,omitempty"`
	// 学生的身份证号，如果ISV有学生身份证号，则同步身份证号作为学生唯一标识。此字段与student_code、家长user_mobile至少选一个。  大陆身份证必须是18位 ， 其它地区或国家的身份证开头需要加\"IC\"开头区分并且不超过18位，但查询账单的时候不要带\"IC\"
	StudentIdentify *string `json:"student_identify,omitempty"`
	// 孩子的家长信息，最多一次输入10个家长，此字段做为识别家长的孩子用，与student_identify、student_code至少选一个
	Users []UserDetails `json:"users,omitempty"`
}

// NewAlipayEcoEduKtBillingSendModel instantiates a new AlipayEcoEduKtBillingSendModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoEduKtBillingSendModel() *AlipayEcoEduKtBillingSendModel {
	this := AlipayEcoEduKtBillingSendModel{}
	return &this
}

// NewAlipayEcoEduKtBillingSendModelWithDefaults instantiates a new AlipayEcoEduKtBillingSendModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoEduKtBillingSendModelWithDefaults() *AlipayEcoEduKtBillingSendModel {
	this := AlipayEcoEduKtBillingSendModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AlipayEcoEduKtBillingSendModel) SetAmount(v string) {
	o.Amount = &v
}

// GetChargeBillTitle returns the ChargeBillTitle field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetChargeBillTitle() string {
	if o == nil || IsNil(o.ChargeBillTitle) {
		var ret string
		return ret
	}
	return *o.ChargeBillTitle
}

// GetChargeBillTitleOk returns a tuple with the ChargeBillTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetChargeBillTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeBillTitle) {
		return nil, false
	}
	return o.ChargeBillTitle, true
}

// HasChargeBillTitle returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasChargeBillTitle() bool {
	if o != nil && !IsNil(o.ChargeBillTitle) {
		return true
	}

	return false
}

// SetChargeBillTitle gets a reference to the given string and assigns it to the ChargeBillTitle field.
func (o *AlipayEcoEduKtBillingSendModel) SetChargeBillTitle(v string) {
	o.ChargeBillTitle = &v
}

// GetChargeItem returns the ChargeItem field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetChargeItem() []ChargeItems {
	if o == nil || IsNil(o.ChargeItem) {
		var ret []ChargeItems
		return ret
	}
	return o.ChargeItem
}

// GetChargeItemOk returns a tuple with the ChargeItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetChargeItemOk() ([]ChargeItems, bool) {
	if o == nil || IsNil(o.ChargeItem) {
		return nil, false
	}
	return o.ChargeItem, true
}

// HasChargeItem returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasChargeItem() bool {
	if o != nil && !IsNil(o.ChargeItem) {
		return true
	}

	return false
}

// SetChargeItem gets a reference to the given []ChargeItems and assigns it to the ChargeItem field.
func (o *AlipayEcoEduKtBillingSendModel) SetChargeItem(v []ChargeItems) {
	o.ChargeItem = v
}

// GetChargeType returns the ChargeType field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetChargeType() string {
	if o == nil || IsNil(o.ChargeType) {
		var ret string
		return ret
	}
	return *o.ChargeType
}

// GetChargeTypeOk returns a tuple with the ChargeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetChargeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeType) {
		return nil, false
	}
	return o.ChargeType, true
}

// HasChargeType returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasChargeType() bool {
	if o != nil && !IsNil(o.ChargeType) {
		return true
	}

	return false
}

// SetChargeType gets a reference to the given string and assigns it to the ChargeType field.
func (o *AlipayEcoEduKtBillingSendModel) SetChargeType(v string) {
	o.ChargeType = &v
}

// GetChildName returns the ChildName field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetChildName() string {
	if o == nil || IsNil(o.ChildName) {
		var ret string
		return ret
	}
	return *o.ChildName
}

// GetChildNameOk returns a tuple with the ChildName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetChildNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChildName) {
		return nil, false
	}
	return o.ChildName, true
}

// HasChildName returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasChildName() bool {
	if o != nil && !IsNil(o.ChildName) {
		return true
	}

	return false
}

// SetChildName gets a reference to the given string and assigns it to the ChildName field.
func (o *AlipayEcoEduKtBillingSendModel) SetChildName(v string) {
	o.ChildName = &v
}

// GetClassIn returns the ClassIn field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetClassIn() string {
	if o == nil || IsNil(o.ClassIn) {
		var ret string
		return ret
	}
	return *o.ClassIn
}

// GetClassInOk returns a tuple with the ClassIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetClassInOk() (*string, bool) {
	if o == nil || IsNil(o.ClassIn) {
		return nil, false
	}
	return o.ClassIn, true
}

// HasClassIn returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasClassIn() bool {
	if o != nil && !IsNil(o.ClassIn) {
		return true
	}

	return false
}

// SetClassIn gets a reference to the given string and assigns it to the ClassIn field.
func (o *AlipayEcoEduKtBillingSendModel) SetClassIn(v string) {
	o.ClassIn = &v
}

// GetEndEnable returns the EndEnable field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetEndEnable() string {
	if o == nil || IsNil(o.EndEnable) {
		var ret string
		return ret
	}
	return *o.EndEnable
}

// GetEndEnableOk returns a tuple with the EndEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetEndEnableOk() (*string, bool) {
	if o == nil || IsNil(o.EndEnable) {
		return nil, false
	}
	return o.EndEnable, true
}

// HasEndEnable returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasEndEnable() bool {
	if o != nil && !IsNil(o.EndEnable) {
		return true
	}

	return false
}

// SetEndEnable gets a reference to the given string and assigns it to the EndEnable field.
func (o *AlipayEcoEduKtBillingSendModel) SetEndEnable(v string) {
	o.EndEnable = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetExtInfo() BillSendExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret BillSendExtInfo
		return ret
	}
	return *o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetExtInfoOk() (*BillSendExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given BillSendExtInfo and assigns it to the ExtInfo field.
func (o *AlipayEcoEduKtBillingSendModel) SetExtInfo(v BillSendExtInfo) {
	o.ExtInfo = &v
}

// GetGmtEnd returns the GmtEnd field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetGmtEnd() string {
	if o == nil || IsNil(o.GmtEnd) {
		var ret string
		return ret
	}
	return *o.GmtEnd
}

// GetGmtEndOk returns a tuple with the GmtEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetGmtEndOk() (*string, bool) {
	if o == nil || IsNil(o.GmtEnd) {
		return nil, false
	}
	return o.GmtEnd, true
}

// HasGmtEnd returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasGmtEnd() bool {
	if o != nil && !IsNil(o.GmtEnd) {
		return true
	}

	return false
}

// SetGmtEnd gets a reference to the given string and assigns it to the GmtEnd field.
func (o *AlipayEcoEduKtBillingSendModel) SetGmtEnd(v string) {
	o.GmtEnd = &v
}

// GetGrade returns the Grade field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetGrade() string {
	if o == nil || IsNil(o.Grade) {
		var ret string
		return ret
	}
	return *o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetGradeOk() (*string, bool) {
	if o == nil || IsNil(o.Grade) {
		return nil, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given string and assigns it to the Grade field.
func (o *AlipayEcoEduKtBillingSendModel) SetGrade(v string) {
	o.Grade = &v
}

// GetOutTradeNo returns the OutTradeNo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetOutTradeNo() string {
	if o == nil || IsNil(o.OutTradeNo) {
		var ret string
		return ret
	}
	return *o.OutTradeNo
}

// GetOutTradeNoOk returns a tuple with the OutTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetOutTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutTradeNo) {
		return nil, false
	}
	return o.OutTradeNo, true
}

// HasOutTradeNo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasOutTradeNo() bool {
	if o != nil && !IsNil(o.OutTradeNo) {
		return true
	}

	return false
}

// SetOutTradeNo gets a reference to the given string and assigns it to the OutTradeNo field.
func (o *AlipayEcoEduKtBillingSendModel) SetOutTradeNo(v string) {
	o.OutTradeNo = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *AlipayEcoEduKtBillingSendModel) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetSchoolNo returns the SchoolNo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetSchoolNo() string {
	if o == nil || IsNil(o.SchoolNo) {
		var ret string
		return ret
	}
	return *o.SchoolNo
}

// GetSchoolNoOk returns a tuple with the SchoolNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetSchoolNoOk() (*string, bool) {
	if o == nil || IsNil(o.SchoolNo) {
		return nil, false
	}
	return o.SchoolNo, true
}

// HasSchoolNo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasSchoolNo() bool {
	if o != nil && !IsNil(o.SchoolNo) {
		return true
	}

	return false
}

// SetSchoolNo gets a reference to the given string and assigns it to the SchoolNo field.
func (o *AlipayEcoEduKtBillingSendModel) SetSchoolNo(v string) {
	o.SchoolNo = &v
}

// GetSchoolPid returns the SchoolPid field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetSchoolPid() string {
	if o == nil || IsNil(o.SchoolPid) {
		var ret string
		return ret
	}
	return *o.SchoolPid
}

// GetSchoolPidOk returns a tuple with the SchoolPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetSchoolPidOk() (*string, bool) {
	if o == nil || IsNil(o.SchoolPid) {
		return nil, false
	}
	return o.SchoolPid, true
}

// HasSchoolPid returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasSchoolPid() bool {
	if o != nil && !IsNil(o.SchoolPid) {
		return true
	}

	return false
}

// SetSchoolPid gets a reference to the given string and assigns it to the SchoolPid field.
func (o *AlipayEcoEduKtBillingSendModel) SetSchoolPid(v string) {
	o.SchoolPid = &v
}

// GetStudentCode returns the StudentCode field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetStudentCode() string {
	if o == nil || IsNil(o.StudentCode) {
		var ret string
		return ret
	}
	return *o.StudentCode
}

// GetStudentCodeOk returns a tuple with the StudentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetStudentCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StudentCode) {
		return nil, false
	}
	return o.StudentCode, true
}

// HasStudentCode returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasStudentCode() bool {
	if o != nil && !IsNil(o.StudentCode) {
		return true
	}

	return false
}

// SetStudentCode gets a reference to the given string and assigns it to the StudentCode field.
func (o *AlipayEcoEduKtBillingSendModel) SetStudentCode(v string) {
	o.StudentCode = &v
}

// GetStudentIdentify returns the StudentIdentify field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetStudentIdentify() string {
	if o == nil || IsNil(o.StudentIdentify) {
		var ret string
		return ret
	}
	return *o.StudentIdentify
}

// GetStudentIdentifyOk returns a tuple with the StudentIdentify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetStudentIdentifyOk() (*string, bool) {
	if o == nil || IsNil(o.StudentIdentify) {
		return nil, false
	}
	return o.StudentIdentify, true
}

// HasStudentIdentify returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasStudentIdentify() bool {
	if o != nil && !IsNil(o.StudentIdentify) {
		return true
	}

	return false
}

// SetStudentIdentify gets a reference to the given string and assigns it to the StudentIdentify field.
func (o *AlipayEcoEduKtBillingSendModel) SetStudentIdentify(v string) {
	o.StudentIdentify = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingSendModel) GetUsers() []UserDetails {
	if o == nil || IsNil(o.Users) {
		var ret []UserDetails
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingSendModel) GetUsersOk() ([]UserDetails, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingSendModel) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserDetails and assigns it to the Users field.
func (o *AlipayEcoEduKtBillingSendModel) SetUsers(v []UserDetails) {
	o.Users = v
}

func (o AlipayEcoEduKtBillingSendModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoEduKtBillingSendModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ChargeBillTitle) {
		toSerialize["charge_bill_title"] = o.ChargeBillTitle
	}
	if !IsNil(o.ChargeItem) {
		toSerialize["charge_item"] = o.ChargeItem
	}
	if !IsNil(o.ChargeType) {
		toSerialize["charge_type"] = o.ChargeType
	}
	if !IsNil(o.ChildName) {
		toSerialize["child_name"] = o.ChildName
	}
	if !IsNil(o.ClassIn) {
		toSerialize["class_in"] = o.ClassIn
	}
	if !IsNil(o.EndEnable) {
		toSerialize["end_enable"] = o.EndEnable
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.GmtEnd) {
		toSerialize["gmt_end"] = o.GmtEnd
	}
	if !IsNil(o.Grade) {
		toSerialize["grade"] = o.Grade
	}
	if !IsNil(o.OutTradeNo) {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.SchoolNo) {
		toSerialize["school_no"] = o.SchoolNo
	}
	if !IsNil(o.SchoolPid) {
		toSerialize["school_pid"] = o.SchoolPid
	}
	if !IsNil(o.StudentCode) {
		toSerialize["student_code"] = o.StudentCode
	}
	if !IsNil(o.StudentIdentify) {
		toSerialize["student_identify"] = o.StudentIdentify
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableAlipayEcoEduKtBillingSendModel struct {
	value *AlipayEcoEduKtBillingSendModel
	isSet bool
}

func (v NullableAlipayEcoEduKtBillingSendModel) Get() *AlipayEcoEduKtBillingSendModel {
	return v.value
}

func (v *NullableAlipayEcoEduKtBillingSendModel) Set(val *AlipayEcoEduKtBillingSendModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoEduKtBillingSendModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoEduKtBillingSendModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoEduKtBillingSendModel(val *AlipayEcoEduKtBillingSendModel) *NullableAlipayEcoEduKtBillingSendModel {
	return &NullableAlipayEcoEduKtBillingSendModel{value: val, isSet: true}
}

func (v NullableAlipayEcoEduKtBillingSendModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoEduKtBillingSendModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
