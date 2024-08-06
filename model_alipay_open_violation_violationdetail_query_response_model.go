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

// checks if the AlipayOpenViolationViolationdetailQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenViolationViolationdetailQueryResponseModel{}

// AlipayOpenViolationViolationdetailQueryResponseModel struct for AlipayOpenViolationViolationdetailQueryResponseModel
type AlipayOpenViolationViolationdetailQueryResponseModel struct {
	// 申诉截止日期。 即商户可进行申诉的截止日期，若超过该时间未申诉或申诉不通过，平台不提供申诉机会，维持原处罚结果；若该时间内完成申诉且申诉通过，平台将撤销对商户的处罚
	AppealDeadLine *string `json:"appeal_dead_line,omitempty"`
	// 商家申诉记录
	AppealReplyRecords []ReplyRecord `json:"appeal_reply_records,omitempty"`
	// 商家是否可以申诉
	CanAppeal *bool `json:"can_appeal,omitempty"`
	// 商家是否可以整改
	CanRectify *bool `json:"can_rectify,omitempty"`
	// 处罚动作及有效期
	PunishAction []string `json:"punish_action,omitempty"`
	// 截止整改时间。即商户可进行整改的截止日期，若超过该时间未整改或整改不通过，平台将对商户进行处罚，若该时间内完成整改且整改通过，平台对商户不处罚
	RectifyDeadLine *string `json:"rectify_dead_line,omitempty"`
	// 商家整改记录
	RectifyReplyRecords []ReplyRecord `json:"rectify_reply_records,omitempty"`
	// 违规工单状态枚举
	Status *string `json:"status,omitempty"`
	// 剩余申诉次数。 即在申诉截止日期内，平台允许商家申诉的次数，当商户提交申诉但申诉未通过时，剩余申诉次数减1，当整改次数为0，不管是否在整改截止日期内，平台将对商户进行处罚
	SurplusAppealCnt *string `json:"surplus_appeal_cnt,omitempty"`
	// 剩余整改次数。即在整改截止日期内，平台允许商家整改的次数，当商户提交整改但整改未通过时，剩余整改次数减1，当整改次数为0，不管是否在整改截止日期内，平台将对商户进行处罚
	SurplusRectifyCnt *string `json:"surplus_rectify_cnt,omitempty"`
	// 违规对象ID
	TargetId *string `json:"target_id,omitempty"`
	// 违规对象名称
	TargetName *string `json:"target_name,omitempty"`
	// 违规对象类型 小程序ID:APPID 生活号ID:PUBLICID
	TargetType *string `json:"target_type,omitempty"`
	// 平台判定商家违规凭证
	ViolationEvidence []AuditEvidenceInfo `json:"violation_evidence,omitempty"`
	// 违规原因
	ViolationReason *string `json:"violation_reason,omitempty"`
	// 支付宝侧生成的违规记录唯一标识
	ViolationRecordId *string `json:"violation_record_id,omitempty"`
	// 违规时间，格式为 yyyy-MM-dd HH:mm:ss
	ViolationTime *string `json:"violation_time,omitempty"`
	// 即平台依据平台规范/规则，判定商户的违规类型
	ViolationType *string `json:"violation_type,omitempty"`
}

// NewAlipayOpenViolationViolationdetailQueryResponseModel instantiates a new AlipayOpenViolationViolationdetailQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenViolationViolationdetailQueryResponseModel() *AlipayOpenViolationViolationdetailQueryResponseModel {
	this := AlipayOpenViolationViolationdetailQueryResponseModel{}
	return &this
}

// NewAlipayOpenViolationViolationdetailQueryResponseModelWithDefaults instantiates a new AlipayOpenViolationViolationdetailQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenViolationViolationdetailQueryResponseModelWithDefaults() *AlipayOpenViolationViolationdetailQueryResponseModel {
	this := AlipayOpenViolationViolationdetailQueryResponseModel{}
	return &this
}

// GetAppealDeadLine returns the AppealDeadLine field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetAppealDeadLine() string {
	if o == nil || IsNil(o.AppealDeadLine) {
		var ret string
		return ret
	}
	return *o.AppealDeadLine
}

// GetAppealDeadLineOk returns a tuple with the AppealDeadLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetAppealDeadLineOk() (*string, bool) {
	if o == nil || IsNil(o.AppealDeadLine) {
		return nil, false
	}
	return o.AppealDeadLine, true
}

// HasAppealDeadLine returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasAppealDeadLine() bool {
	if o != nil && !IsNil(o.AppealDeadLine) {
		return true
	}

	return false
}

// SetAppealDeadLine gets a reference to the given string and assigns it to the AppealDeadLine field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetAppealDeadLine(v string) {
	o.AppealDeadLine = &v
}

// GetAppealReplyRecords returns the AppealReplyRecords field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetAppealReplyRecords() []ReplyRecord {
	if o == nil || IsNil(o.AppealReplyRecords) {
		var ret []ReplyRecord
		return ret
	}
	return o.AppealReplyRecords
}

// GetAppealReplyRecordsOk returns a tuple with the AppealReplyRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetAppealReplyRecordsOk() ([]ReplyRecord, bool) {
	if o == nil || IsNil(o.AppealReplyRecords) {
		return nil, false
	}
	return o.AppealReplyRecords, true
}

// HasAppealReplyRecords returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasAppealReplyRecords() bool {
	if o != nil && !IsNil(o.AppealReplyRecords) {
		return true
	}

	return false
}

// SetAppealReplyRecords gets a reference to the given []ReplyRecord and assigns it to the AppealReplyRecords field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetAppealReplyRecords(v []ReplyRecord) {
	o.AppealReplyRecords = v
}

// GetCanAppeal returns the CanAppeal field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetCanAppeal() bool {
	if o == nil || IsNil(o.CanAppeal) {
		var ret bool
		return ret
	}
	return *o.CanAppeal
}

// GetCanAppealOk returns a tuple with the CanAppeal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetCanAppealOk() (*bool, bool) {
	if o == nil || IsNil(o.CanAppeal) {
		return nil, false
	}
	return o.CanAppeal, true
}

// HasCanAppeal returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasCanAppeal() bool {
	if o != nil && !IsNil(o.CanAppeal) {
		return true
	}

	return false
}

// SetCanAppeal gets a reference to the given bool and assigns it to the CanAppeal field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetCanAppeal(v bool) {
	o.CanAppeal = &v
}

// GetCanRectify returns the CanRectify field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetCanRectify() bool {
	if o == nil || IsNil(o.CanRectify) {
		var ret bool
		return ret
	}
	return *o.CanRectify
}

// GetCanRectifyOk returns a tuple with the CanRectify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetCanRectifyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRectify) {
		return nil, false
	}
	return o.CanRectify, true
}

// HasCanRectify returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasCanRectify() bool {
	if o != nil && !IsNil(o.CanRectify) {
		return true
	}

	return false
}

// SetCanRectify gets a reference to the given bool and assigns it to the CanRectify field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetCanRectify(v bool) {
	o.CanRectify = &v
}

// GetPunishAction returns the PunishAction field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetPunishAction() []string {
	if o == nil || IsNil(o.PunishAction) {
		var ret []string
		return ret
	}
	return o.PunishAction
}

// GetPunishActionOk returns a tuple with the PunishAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetPunishActionOk() ([]string, bool) {
	if o == nil || IsNil(o.PunishAction) {
		return nil, false
	}
	return o.PunishAction, true
}

// HasPunishAction returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasPunishAction() bool {
	if o != nil && !IsNil(o.PunishAction) {
		return true
	}

	return false
}

// SetPunishAction gets a reference to the given []string and assigns it to the PunishAction field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetPunishAction(v []string) {
	o.PunishAction = v
}

// GetRectifyDeadLine returns the RectifyDeadLine field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetRectifyDeadLine() string {
	if o == nil || IsNil(o.RectifyDeadLine) {
		var ret string
		return ret
	}
	return *o.RectifyDeadLine
}

// GetRectifyDeadLineOk returns a tuple with the RectifyDeadLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetRectifyDeadLineOk() (*string, bool) {
	if o == nil || IsNil(o.RectifyDeadLine) {
		return nil, false
	}
	return o.RectifyDeadLine, true
}

// HasRectifyDeadLine returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasRectifyDeadLine() bool {
	if o != nil && !IsNil(o.RectifyDeadLine) {
		return true
	}

	return false
}

// SetRectifyDeadLine gets a reference to the given string and assigns it to the RectifyDeadLine field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetRectifyDeadLine(v string) {
	o.RectifyDeadLine = &v
}

// GetRectifyReplyRecords returns the RectifyReplyRecords field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetRectifyReplyRecords() []ReplyRecord {
	if o == nil || IsNil(o.RectifyReplyRecords) {
		var ret []ReplyRecord
		return ret
	}
	return o.RectifyReplyRecords
}

// GetRectifyReplyRecordsOk returns a tuple with the RectifyReplyRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetRectifyReplyRecordsOk() ([]ReplyRecord, bool) {
	if o == nil || IsNil(o.RectifyReplyRecords) {
		return nil, false
	}
	return o.RectifyReplyRecords, true
}

// HasRectifyReplyRecords returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasRectifyReplyRecords() bool {
	if o != nil && !IsNil(o.RectifyReplyRecords) {
		return true
	}

	return false
}

// SetRectifyReplyRecords gets a reference to the given []ReplyRecord and assigns it to the RectifyReplyRecords field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetRectifyReplyRecords(v []ReplyRecord) {
	o.RectifyReplyRecords = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetSurplusAppealCnt returns the SurplusAppealCnt field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetSurplusAppealCnt() string {
	if o == nil || IsNil(o.SurplusAppealCnt) {
		var ret string
		return ret
	}
	return *o.SurplusAppealCnt
}

// GetSurplusAppealCntOk returns a tuple with the SurplusAppealCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetSurplusAppealCntOk() (*string, bool) {
	if o == nil || IsNil(o.SurplusAppealCnt) {
		return nil, false
	}
	return o.SurplusAppealCnt, true
}

// HasSurplusAppealCnt returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasSurplusAppealCnt() bool {
	if o != nil && !IsNil(o.SurplusAppealCnt) {
		return true
	}

	return false
}

// SetSurplusAppealCnt gets a reference to the given string and assigns it to the SurplusAppealCnt field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetSurplusAppealCnt(v string) {
	o.SurplusAppealCnt = &v
}

// GetSurplusRectifyCnt returns the SurplusRectifyCnt field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetSurplusRectifyCnt() string {
	if o == nil || IsNil(o.SurplusRectifyCnt) {
		var ret string
		return ret
	}
	return *o.SurplusRectifyCnt
}

// GetSurplusRectifyCntOk returns a tuple with the SurplusRectifyCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetSurplusRectifyCntOk() (*string, bool) {
	if o == nil || IsNil(o.SurplusRectifyCnt) {
		return nil, false
	}
	return o.SurplusRectifyCnt, true
}

// HasSurplusRectifyCnt returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasSurplusRectifyCnt() bool {
	if o != nil && !IsNil(o.SurplusRectifyCnt) {
		return true
	}

	return false
}

// SetSurplusRectifyCnt gets a reference to the given string and assigns it to the SurplusRectifyCnt field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetSurplusRectifyCnt(v string) {
	o.SurplusRectifyCnt = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetTargetType(v string) {
	o.TargetType = &v
}

// GetViolationEvidence returns the ViolationEvidence field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationEvidence() []AuditEvidenceInfo {
	if o == nil || IsNil(o.ViolationEvidence) {
		var ret []AuditEvidenceInfo
		return ret
	}
	return o.ViolationEvidence
}

// GetViolationEvidenceOk returns a tuple with the ViolationEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationEvidenceOk() ([]AuditEvidenceInfo, bool) {
	if o == nil || IsNil(o.ViolationEvidence) {
		return nil, false
	}
	return o.ViolationEvidence, true
}

// HasViolationEvidence returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasViolationEvidence() bool {
	if o != nil && !IsNil(o.ViolationEvidence) {
		return true
	}

	return false
}

// SetViolationEvidence gets a reference to the given []AuditEvidenceInfo and assigns it to the ViolationEvidence field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetViolationEvidence(v []AuditEvidenceInfo) {
	o.ViolationEvidence = v
}

// GetViolationReason returns the ViolationReason field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationReason() string {
	if o == nil || IsNil(o.ViolationReason) {
		var ret string
		return ret
	}
	return *o.ViolationReason
}

// GetViolationReasonOk returns a tuple with the ViolationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationReason) {
		return nil, false
	}
	return o.ViolationReason, true
}

// HasViolationReason returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasViolationReason() bool {
	if o != nil && !IsNil(o.ViolationReason) {
		return true
	}

	return false
}

// SetViolationReason gets a reference to the given string and assigns it to the ViolationReason field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetViolationReason(v string) {
	o.ViolationReason = &v
}

// GetViolationRecordId returns the ViolationRecordId field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationRecordId() string {
	if o == nil || IsNil(o.ViolationRecordId) {
		var ret string
		return ret
	}
	return *o.ViolationRecordId
}

// GetViolationRecordIdOk returns a tuple with the ViolationRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationRecordId) {
		return nil, false
	}
	return o.ViolationRecordId, true
}

// HasViolationRecordId returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasViolationRecordId() bool {
	if o != nil && !IsNil(o.ViolationRecordId) {
		return true
	}

	return false
}

// SetViolationRecordId gets a reference to the given string and assigns it to the ViolationRecordId field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetViolationRecordId(v string) {
	o.ViolationRecordId = &v
}

// GetViolationTime returns the ViolationTime field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationTime() string {
	if o == nil || IsNil(o.ViolationTime) {
		var ret string
		return ret
	}
	return *o.ViolationTime
}

// GetViolationTimeOk returns a tuple with the ViolationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationTime) {
		return nil, false
	}
	return o.ViolationTime, true
}

// HasViolationTime returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasViolationTime() bool {
	if o != nil && !IsNil(o.ViolationTime) {
		return true
	}

	return false
}

// SetViolationTime gets a reference to the given string and assigns it to the ViolationTime field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetViolationTime(v string) {
	o.ViolationTime = &v
}

// GetViolationType returns the ViolationType field value if set, zero value otherwise.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationType() string {
	if o == nil || IsNil(o.ViolationType) {
		var ret string
		return ret
	}
	return *o.ViolationType
}

// GetViolationTypeOk returns a tuple with the ViolationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) GetViolationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationType) {
		return nil, false
	}
	return o.ViolationType, true
}

// HasViolationType returns a boolean if a field has been set.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) HasViolationType() bool {
	if o != nil && !IsNil(o.ViolationType) {
		return true
	}

	return false
}

// SetViolationType gets a reference to the given string and assigns it to the ViolationType field.
func (o *AlipayOpenViolationViolationdetailQueryResponseModel) SetViolationType(v string) {
	o.ViolationType = &v
}

func (o AlipayOpenViolationViolationdetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenViolationViolationdetailQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppealDeadLine) {
		toSerialize["appeal_dead_line"] = o.AppealDeadLine
	}
	if !IsNil(o.AppealReplyRecords) {
		toSerialize["appeal_reply_records"] = o.AppealReplyRecords
	}
	if !IsNil(o.CanAppeal) {
		toSerialize["can_appeal"] = o.CanAppeal
	}
	if !IsNil(o.CanRectify) {
		toSerialize["can_rectify"] = o.CanRectify
	}
	if !IsNil(o.PunishAction) {
		toSerialize["punish_action"] = o.PunishAction
	}
	if !IsNil(o.RectifyDeadLine) {
		toSerialize["rectify_dead_line"] = o.RectifyDeadLine
	}
	if !IsNil(o.RectifyReplyRecords) {
		toSerialize["rectify_reply_records"] = o.RectifyReplyRecords
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SurplusAppealCnt) {
		toSerialize["surplus_appeal_cnt"] = o.SurplusAppealCnt
	}
	if !IsNil(o.SurplusRectifyCnt) {
		toSerialize["surplus_rectify_cnt"] = o.SurplusRectifyCnt
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetName) {
		toSerialize["target_name"] = o.TargetName
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	if !IsNil(o.ViolationEvidence) {
		toSerialize["violation_evidence"] = o.ViolationEvidence
	}
	if !IsNil(o.ViolationReason) {
		toSerialize["violation_reason"] = o.ViolationReason
	}
	if !IsNil(o.ViolationRecordId) {
		toSerialize["violation_record_id"] = o.ViolationRecordId
	}
	if !IsNil(o.ViolationTime) {
		toSerialize["violation_time"] = o.ViolationTime
	}
	if !IsNil(o.ViolationType) {
		toSerialize["violation_type"] = o.ViolationType
	}
	return toSerialize, nil
}

type NullableAlipayOpenViolationViolationdetailQueryResponseModel struct {
	value *AlipayOpenViolationViolationdetailQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenViolationViolationdetailQueryResponseModel) Get() *AlipayOpenViolationViolationdetailQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenViolationViolationdetailQueryResponseModel) Set(val *AlipayOpenViolationViolationdetailQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenViolationViolationdetailQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenViolationViolationdetailQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenViolationViolationdetailQueryResponseModel(val *AlipayOpenViolationViolationdetailQueryResponseModel) *NullableAlipayOpenViolationViolationdetailQueryResponseModel {
	return &NullableAlipayOpenViolationViolationdetailQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenViolationViolationdetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenViolationViolationdetailQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


