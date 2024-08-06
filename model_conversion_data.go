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

// checks if the ConversionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConversionData{}

// ConversionData struct for ConversionData
type ConversionData struct {
	// 广告id
	AdId *string `json:"ad_id,omitempty"`
	// 转化事件属性信息，用于转化事件类型相关属性规则上传。 可支持上传属性的转化事件类型及属性规则请参考该文档：<a href='https://adpub.alipay.com/adrlark/ivdktpyh511x9r6i'>转化事件类型属性规则</a>
	AttributeList []ConversionProperty `json:"attribute_list,omitempty"`
	Attributes *ConversionProperty `json:"attributes,omitempty"`
	// 转化流水号：由用户自定义，用于幂等
	BizNo *string `json:"biz_no,omitempty"`
	// 转化回调扩展信息
	CallbackExtInfo *string `json:"callback_ext_info,omitempty"`
	// 来自监测链接配置的__CID__宏参发生用户点击替换的值
	Cid *string `json:"cid,omitempty"`
	// 转化金额，单位分
	ConversionAmount *string `json:"conversion_amount,omitempty"`
	// 当source=COMMON_CONVERSION_ID时使用
	ConversionId *string `json:"conversion_id,omitempty"`
	// 转化时间，UTC 时间戳，单位：秒
	ConversionTime *int32 `json:"conversion_time,omitempty"`
	// 转化事件类型数字
	ConversionType *string `json:"conversion_type,omitempty"`
	// 创意ID
	CreativeId *string `json:"creative_id,omitempty"`
	// 当source=XLIGHT或者DATASET时使用：  数据集id(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)
	// Deprecated
	DataId *string `json:"data_id,omitempty"`
	// 当source=OTHER时使用:  数据类型： KR_MEMBER - 客如云入会 KR_TRADE - 客如云交易 TB_LIVE -  淘宝直播(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)
	// Deprecated
	DataSrcType *string `json:"data_src_type,omitempty"`
	// 单元ID
	GroupId *string `json:"group_id,omitempty"`
	// XLIGHT - 灯火归因 MERCHANT -商家自主归因 该字段若为空则默认为XLIGHT
	JoinChannel *string `json:"join_channel,omitempty"`
	// 计划ID
	PlanId *string `json:"plan_id,omitempty"`
	// 商户在灯火pb端的id, 可代替principal_tag
	PrincipalId *string `json:"principal_id,omitempty"`
	// 商家标志，可代替principal_id
	PrincipalTag *string `json:"principal_tag,omitempty"`
	// 转化归因字段列表
	PropertyList []ConversionProperty `json:"property_list,omitempty"`
	// 来源： COMMON_TARGET-通用转化事件类型适用 CALLBACK-APP推广类转化事件类型适用 OTHER-其它
	Source *string `json:"source,omitempty"`
	// 当source=OTHER时使用： 主体id，例如品牌id(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)
	// Deprecated
	TargetId *string `json:"target_id,omitempty"`
	// 当source=OTHER时使用：  主体类型： BRAND - 品牌 STORE - 店铺 LIVE - 直播 等等(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)
	// Deprecated
	TargetType *string `json:"target_type,omitempty"`
	// 转化用户（发生真实转化的用户）唯一标识。 当uuid_type=PID时，传2088UID（例：208801217938xxxx）
	Uuid *string `json:"uuid,omitempty"`
	// 支付宝用户在应用维度下的唯一标识
	UuidOpenId *string `json:"uuid_open_id,omitempty"`
	// 通用转化事件类型数据回传可使用：PID，表示转化用户唯一标志类型。 自建站转化事件类型可使用：PID_ENCRYPT，自建站转化事件类型编码:445~450。 app推广类转化事件类型可使用： OAID_MD5：oaid md5值 IDFA_MD5：idfa md5值 IP+UA:ip ua拼接值
	UuidType *string `json:"uuid_type,omitempty"`
}

// NewConversionData instantiates a new ConversionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversionData() *ConversionData {
	this := ConversionData{}
	return &this
}

// NewConversionDataWithDefaults instantiates a new ConversionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversionDataWithDefaults() *ConversionData {
	this := ConversionData{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *ConversionData) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *ConversionData) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *ConversionData) SetAdId(v string) {
	o.AdId = &v
}

// GetAttributeList returns the AttributeList field value if set, zero value otherwise.
func (o *ConversionData) GetAttributeList() []ConversionProperty {
	if o == nil || IsNil(o.AttributeList) {
		var ret []ConversionProperty
		return ret
	}
	return o.AttributeList
}

// GetAttributeListOk returns a tuple with the AttributeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetAttributeListOk() ([]ConversionProperty, bool) {
	if o == nil || IsNil(o.AttributeList) {
		return nil, false
	}
	return o.AttributeList, true
}

// HasAttributeList returns a boolean if a field has been set.
func (o *ConversionData) HasAttributeList() bool {
	if o != nil && !IsNil(o.AttributeList) {
		return true
	}

	return false
}

// SetAttributeList gets a reference to the given []ConversionProperty and assigns it to the AttributeList field.
func (o *ConversionData) SetAttributeList(v []ConversionProperty) {
	o.AttributeList = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ConversionData) GetAttributes() ConversionProperty {
	if o == nil || IsNil(o.Attributes) {
		var ret ConversionProperty
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetAttributesOk() (*ConversionProperty, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ConversionData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ConversionProperty and assigns it to the Attributes field.
func (o *ConversionData) SetAttributes(v ConversionProperty) {
	o.Attributes = &v
}

// GetBizNo returns the BizNo field value if set, zero value otherwise.
func (o *ConversionData) GetBizNo() string {
	if o == nil || IsNil(o.BizNo) {
		var ret string
		return ret
	}
	return *o.BizNo
}

// GetBizNoOk returns a tuple with the BizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.BizNo) {
		return nil, false
	}
	return o.BizNo, true
}

// HasBizNo returns a boolean if a field has been set.
func (o *ConversionData) HasBizNo() bool {
	if o != nil && !IsNil(o.BizNo) {
		return true
	}

	return false
}

// SetBizNo gets a reference to the given string and assigns it to the BizNo field.
func (o *ConversionData) SetBizNo(v string) {
	o.BizNo = &v
}

// GetCallbackExtInfo returns the CallbackExtInfo field value if set, zero value otherwise.
func (o *ConversionData) GetCallbackExtInfo() string {
	if o == nil || IsNil(o.CallbackExtInfo) {
		var ret string
		return ret
	}
	return *o.CallbackExtInfo
}

// GetCallbackExtInfoOk returns a tuple with the CallbackExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetCallbackExtInfoOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackExtInfo) {
		return nil, false
	}
	return o.CallbackExtInfo, true
}

// HasCallbackExtInfo returns a boolean if a field has been set.
func (o *ConversionData) HasCallbackExtInfo() bool {
	if o != nil && !IsNil(o.CallbackExtInfo) {
		return true
	}

	return false
}

// SetCallbackExtInfo gets a reference to the given string and assigns it to the CallbackExtInfo field.
func (o *ConversionData) SetCallbackExtInfo(v string) {
	o.CallbackExtInfo = &v
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *ConversionData) GetCid() string {
	if o == nil || IsNil(o.Cid) {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetCidOk() (*string, bool) {
	if o == nil || IsNil(o.Cid) {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *ConversionData) HasCid() bool {
	if o != nil && !IsNil(o.Cid) {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *ConversionData) SetCid(v string) {
	o.Cid = &v
}

// GetConversionAmount returns the ConversionAmount field value if set, zero value otherwise.
func (o *ConversionData) GetConversionAmount() string {
	if o == nil || IsNil(o.ConversionAmount) {
		var ret string
		return ret
	}
	return *o.ConversionAmount
}

// GetConversionAmountOk returns a tuple with the ConversionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetConversionAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ConversionAmount) {
		return nil, false
	}
	return o.ConversionAmount, true
}

// HasConversionAmount returns a boolean if a field has been set.
func (o *ConversionData) HasConversionAmount() bool {
	if o != nil && !IsNil(o.ConversionAmount) {
		return true
	}

	return false
}

// SetConversionAmount gets a reference to the given string and assigns it to the ConversionAmount field.
func (o *ConversionData) SetConversionAmount(v string) {
	o.ConversionAmount = &v
}

// GetConversionId returns the ConversionId field value if set, zero value otherwise.
func (o *ConversionData) GetConversionId() string {
	if o == nil || IsNil(o.ConversionId) {
		var ret string
		return ret
	}
	return *o.ConversionId
}

// GetConversionIdOk returns a tuple with the ConversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetConversionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversionId) {
		return nil, false
	}
	return o.ConversionId, true
}

// HasConversionId returns a boolean if a field has been set.
func (o *ConversionData) HasConversionId() bool {
	if o != nil && !IsNil(o.ConversionId) {
		return true
	}

	return false
}

// SetConversionId gets a reference to the given string and assigns it to the ConversionId field.
func (o *ConversionData) SetConversionId(v string) {
	o.ConversionId = &v
}

// GetConversionTime returns the ConversionTime field value if set, zero value otherwise.
func (o *ConversionData) GetConversionTime() int32 {
	if o == nil || IsNil(o.ConversionTime) {
		var ret int32
		return ret
	}
	return *o.ConversionTime
}

// GetConversionTimeOk returns a tuple with the ConversionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetConversionTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ConversionTime) {
		return nil, false
	}
	return o.ConversionTime, true
}

// HasConversionTime returns a boolean if a field has been set.
func (o *ConversionData) HasConversionTime() bool {
	if o != nil && !IsNil(o.ConversionTime) {
		return true
	}

	return false
}

// SetConversionTime gets a reference to the given int32 and assigns it to the ConversionTime field.
func (o *ConversionData) SetConversionTime(v int32) {
	o.ConversionTime = &v
}

// GetConversionType returns the ConversionType field value if set, zero value otherwise.
func (o *ConversionData) GetConversionType() string {
	if o == nil || IsNil(o.ConversionType) {
		var ret string
		return ret
	}
	return *o.ConversionType
}

// GetConversionTypeOk returns a tuple with the ConversionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetConversionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConversionType) {
		return nil, false
	}
	return o.ConversionType, true
}

// HasConversionType returns a boolean if a field has been set.
func (o *ConversionData) HasConversionType() bool {
	if o != nil && !IsNil(o.ConversionType) {
		return true
	}

	return false
}

// SetConversionType gets a reference to the given string and assigns it to the ConversionType field.
func (o *ConversionData) SetConversionType(v string) {
	o.ConversionType = &v
}

// GetCreativeId returns the CreativeId field value if set, zero value otherwise.
func (o *ConversionData) GetCreativeId() string {
	if o == nil || IsNil(o.CreativeId) {
		var ret string
		return ret
	}
	return *o.CreativeId
}

// GetCreativeIdOk returns a tuple with the CreativeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetCreativeIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreativeId) {
		return nil, false
	}
	return o.CreativeId, true
}

// HasCreativeId returns a boolean if a field has been set.
func (o *ConversionData) HasCreativeId() bool {
	if o != nil && !IsNil(o.CreativeId) {
		return true
	}

	return false
}

// SetCreativeId gets a reference to the given string and assigns it to the CreativeId field.
func (o *ConversionData) SetCreativeId(v string) {
	o.CreativeId = &v
}

// GetDataId returns the DataId field value if set, zero value otherwise.
// Deprecated
func (o *ConversionData) GetDataId() string {
	if o == nil || IsNil(o.DataId) {
		var ret string
		return ret
	}
	return *o.DataId
}

// GetDataIdOk returns a tuple with the DataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ConversionData) GetDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataId) {
		return nil, false
	}
	return o.DataId, true
}

// HasDataId returns a boolean if a field has been set.
func (o *ConversionData) HasDataId() bool {
	if o != nil && !IsNil(o.DataId) {
		return true
	}

	return false
}

// SetDataId gets a reference to the given string and assigns it to the DataId field.
// Deprecated
func (o *ConversionData) SetDataId(v string) {
	o.DataId = &v
}

// GetDataSrcType returns the DataSrcType field value if set, zero value otherwise.
// Deprecated
func (o *ConversionData) GetDataSrcType() string {
	if o == nil || IsNil(o.DataSrcType) {
		var ret string
		return ret
	}
	return *o.DataSrcType
}

// GetDataSrcTypeOk returns a tuple with the DataSrcType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ConversionData) GetDataSrcTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataSrcType) {
		return nil, false
	}
	return o.DataSrcType, true
}

// HasDataSrcType returns a boolean if a field has been set.
func (o *ConversionData) HasDataSrcType() bool {
	if o != nil && !IsNil(o.DataSrcType) {
		return true
	}

	return false
}

// SetDataSrcType gets a reference to the given string and assigns it to the DataSrcType field.
// Deprecated
func (o *ConversionData) SetDataSrcType(v string) {
	o.DataSrcType = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ConversionData) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ConversionData) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ConversionData) SetGroupId(v string) {
	o.GroupId = &v
}

// GetJoinChannel returns the JoinChannel field value if set, zero value otherwise.
func (o *ConversionData) GetJoinChannel() string {
	if o == nil || IsNil(o.JoinChannel) {
		var ret string
		return ret
	}
	return *o.JoinChannel
}

// GetJoinChannelOk returns a tuple with the JoinChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetJoinChannelOk() (*string, bool) {
	if o == nil || IsNil(o.JoinChannel) {
		return nil, false
	}
	return o.JoinChannel, true
}

// HasJoinChannel returns a boolean if a field has been set.
func (o *ConversionData) HasJoinChannel() bool {
	if o != nil && !IsNil(o.JoinChannel) {
		return true
	}

	return false
}

// SetJoinChannel gets a reference to the given string and assigns it to the JoinChannel field.
func (o *ConversionData) SetJoinChannel(v string) {
	o.JoinChannel = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *ConversionData) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *ConversionData) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *ConversionData) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise.
func (o *ConversionData) GetPrincipalId() string {
	if o == nil || IsNil(o.PrincipalId) {
		var ret string
		return ret
	}
	return *o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetPrincipalIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalId) {
		return nil, false
	}
	return o.PrincipalId, true
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *ConversionData) HasPrincipalId() bool {
	if o != nil && !IsNil(o.PrincipalId) {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given string and assigns it to the PrincipalId field.
func (o *ConversionData) SetPrincipalId(v string) {
	o.PrincipalId = &v
}

// GetPrincipalTag returns the PrincipalTag field value if set, zero value otherwise.
func (o *ConversionData) GetPrincipalTag() string {
	if o == nil || IsNil(o.PrincipalTag) {
		var ret string
		return ret
	}
	return *o.PrincipalTag
}

// GetPrincipalTagOk returns a tuple with the PrincipalTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetPrincipalTagOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalTag) {
		return nil, false
	}
	return o.PrincipalTag, true
}

// HasPrincipalTag returns a boolean if a field has been set.
func (o *ConversionData) HasPrincipalTag() bool {
	if o != nil && !IsNil(o.PrincipalTag) {
		return true
	}

	return false
}

// SetPrincipalTag gets a reference to the given string and assigns it to the PrincipalTag field.
func (o *ConversionData) SetPrincipalTag(v string) {
	o.PrincipalTag = &v
}

// GetPropertyList returns the PropertyList field value if set, zero value otherwise.
func (o *ConversionData) GetPropertyList() []ConversionProperty {
	if o == nil || IsNil(o.PropertyList) {
		var ret []ConversionProperty
		return ret
	}
	return o.PropertyList
}

// GetPropertyListOk returns a tuple with the PropertyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetPropertyListOk() ([]ConversionProperty, bool) {
	if o == nil || IsNil(o.PropertyList) {
		return nil, false
	}
	return o.PropertyList, true
}

// HasPropertyList returns a boolean if a field has been set.
func (o *ConversionData) HasPropertyList() bool {
	if o != nil && !IsNil(o.PropertyList) {
		return true
	}

	return false
}

// SetPropertyList gets a reference to the given []ConversionProperty and assigns it to the PropertyList field.
func (o *ConversionData) SetPropertyList(v []ConversionProperty) {
	o.PropertyList = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConversionData) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConversionData) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ConversionData) SetSource(v string) {
	o.Source = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
// Deprecated
func (o *ConversionData) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ConversionData) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *ConversionData) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
// Deprecated
func (o *ConversionData) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
// Deprecated
func (o *ConversionData) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ConversionData) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *ConversionData) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
// Deprecated
func (o *ConversionData) SetTargetType(v string) {
	o.TargetType = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ConversionData) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ConversionData) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ConversionData) SetUuid(v string) {
	o.Uuid = &v
}

// GetUuidOpenId returns the UuidOpenId field value if set, zero value otherwise.
func (o *ConversionData) GetUuidOpenId() string {
	if o == nil || IsNil(o.UuidOpenId) {
		var ret string
		return ret
	}
	return *o.UuidOpenId
}

// GetUuidOpenIdOk returns a tuple with the UuidOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetUuidOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.UuidOpenId) {
		return nil, false
	}
	return o.UuidOpenId, true
}

// HasUuidOpenId returns a boolean if a field has been set.
func (o *ConversionData) HasUuidOpenId() bool {
	if o != nil && !IsNil(o.UuidOpenId) {
		return true
	}

	return false
}

// SetUuidOpenId gets a reference to the given string and assigns it to the UuidOpenId field.
func (o *ConversionData) SetUuidOpenId(v string) {
	o.UuidOpenId = &v
}

// GetUuidType returns the UuidType field value if set, zero value otherwise.
func (o *ConversionData) GetUuidType() string {
	if o == nil || IsNil(o.UuidType) {
		var ret string
		return ret
	}
	return *o.UuidType
}

// GetUuidTypeOk returns a tuple with the UuidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionData) GetUuidTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UuidType) {
		return nil, false
	}
	return o.UuidType, true
}

// HasUuidType returns a boolean if a field has been set.
func (o *ConversionData) HasUuidType() bool {
	if o != nil && !IsNil(o.UuidType) {
		return true
	}

	return false
}

// SetUuidType gets a reference to the given string and assigns it to the UuidType field.
func (o *ConversionData) SetUuidType(v string) {
	o.UuidType = &v
}

func (o ConversionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConversionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["ad_id"] = o.AdId
	}
	if !IsNil(o.AttributeList) {
		toSerialize["attribute_list"] = o.AttributeList
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.BizNo) {
		toSerialize["biz_no"] = o.BizNo
	}
	if !IsNil(o.CallbackExtInfo) {
		toSerialize["callback_ext_info"] = o.CallbackExtInfo
	}
	if !IsNil(o.Cid) {
		toSerialize["cid"] = o.Cid
	}
	if !IsNil(o.ConversionAmount) {
		toSerialize["conversion_amount"] = o.ConversionAmount
	}
	if !IsNil(o.ConversionId) {
		toSerialize["conversion_id"] = o.ConversionId
	}
	if !IsNil(o.ConversionTime) {
		toSerialize["conversion_time"] = o.ConversionTime
	}
	if !IsNil(o.ConversionType) {
		toSerialize["conversion_type"] = o.ConversionType
	}
	if !IsNil(o.CreativeId) {
		toSerialize["creative_id"] = o.CreativeId
	}
	if !IsNil(o.DataId) {
		toSerialize["data_id"] = o.DataId
	}
	if !IsNil(o.DataSrcType) {
		toSerialize["data_src_type"] = o.DataSrcType
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.JoinChannel) {
		toSerialize["join_channel"] = o.JoinChannel
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.PrincipalId) {
		toSerialize["principal_id"] = o.PrincipalId
	}
	if !IsNil(o.PrincipalTag) {
		toSerialize["principal_tag"] = o.PrincipalTag
	}
	if !IsNil(o.PropertyList) {
		toSerialize["property_list"] = o.PropertyList
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.UuidOpenId) {
		toSerialize["uuid_open_id"] = o.UuidOpenId
	}
	if !IsNil(o.UuidType) {
		toSerialize["uuid_type"] = o.UuidType
	}
	return toSerialize, nil
}

type NullableConversionData struct {
	value *ConversionData
	isSet bool
}

func (v NullableConversionData) Get() *ConversionData {
	return v.value
}

func (v *NullableConversionData) Set(val *ConversionData) {
	v.value = val
	v.isSet = true
}

func (v NullableConversionData) IsSet() bool {
	return v.isSet
}

func (v *NullableConversionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversionData(val *ConversionData) *NullableConversionData {
	return &NullableConversionData{value: val, isSet: true}
}

func (v NullableConversionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


