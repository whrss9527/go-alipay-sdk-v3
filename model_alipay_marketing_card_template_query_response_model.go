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

// checks if the AlipayMarketingCardTemplateQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardTemplateQueryResponseModel{}

// AlipayMarketingCardTemplateQueryResponseModel struct for AlipayMarketingCardTemplateQueryResponseModel
type AlipayMarketingCardTemplateQueryResponseModel struct {
	// 接入版本，接入的是基础版本一定会有值
	AccessVersion *string `json:"access_version,omitempty"`
	// 业务卡号前缀，由商户指定  支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀
	BizNoPrefix *string `json:"biz_no_prefix,omitempty"`
	// 业务卡号后缀的长度  支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀
	BizNoSuffixLen *string `json:"biz_no_suffix_len,omitempty"`
	// \"卡行动点配置； 行动点，即用户可点击跳转的区块，类似按钮控件的交互； 单张卡最多4个行动点。\"
	CardActionList []TemplateActionInfoDTO `json:"card_action_list,omitempty"`
	// 卡等级配置
	CardLevelConfs []TemplateCardLevelConfDTO `json:"card_level_confs,omitempty"`
	// 卡特定标签，只供特定业务使用，通常接入无需关注
	CardSpecTag *string `json:"card_spec_tag,omitempty"`
	// 会员卡类型：  OUT_MEMBER_CARD：外部权益卡
	CardType *string `json:"card_type,omitempty"`
	// 栏位信息（卡包详情页面展现的栏位）
	ColumnInfoList []TemplateColumnInfoDTO `json:"column_info_list,omitempty"`
	// \"字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE\"
	FieldRuleList []TemplateFieldRuleDTO `json:"field_rule_list,omitempty"`
	MdcodeNotifyConf *TemplateMdcodeNotifyConfDTO `json:"mdcode_notify_conf,omitempty"`
	OpenCardConf *TemplateOpenCardConfDTO `json:"open_card_conf,omitempty"`
	PaidOuterCardConf *PaidOuterCardTemplateConfDTO `json:"paid_outer_card_conf,omitempty"`
	// 卡模板投放渠道
	PubChannels []PubChannelDTO `json:"pub_channels,omitempty"`
	// 服务标签列表
	ServiceLabelList []string `json:"service_label_list,omitempty"`
	// 支付宝门店id
	ShopIds []string `json:"shop_ids,omitempty"`
	// spi应用id，设置的实现spi.alipay.user.opencard.get接口的app_id
	SpiAppId *string `json:"spi_app_id,omitempty"`
	// \"权益信息， 1、在卡包的卡详情页面会自动添加权益栏位，展现会员卡特权， 2、如果添加门店渠道，则可在门店页展现会员卡的权益\"
	TemplateBenefitInfo []TemplateBenefitInfoDTO `json:"template_benefit_info,omitempty"`
	TemplateFormConfig *TemplateFormConfig `json:"template_form_config,omitempty"`
	TemplateStyleInfo *TemplateStyleInfoDTO `json:"template_style_info,omitempty"`
}

// NewAlipayMarketingCardTemplateQueryResponseModel instantiates a new AlipayMarketingCardTemplateQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardTemplateQueryResponseModel() *AlipayMarketingCardTemplateQueryResponseModel {
	this := AlipayMarketingCardTemplateQueryResponseModel{}
	return &this
}

// NewAlipayMarketingCardTemplateQueryResponseModelWithDefaults instantiates a new AlipayMarketingCardTemplateQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardTemplateQueryResponseModelWithDefaults() *AlipayMarketingCardTemplateQueryResponseModel {
	this := AlipayMarketingCardTemplateQueryResponseModel{}
	return &this
}

// GetAccessVersion returns the AccessVersion field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetAccessVersion() string {
	if o == nil || IsNil(o.AccessVersion) {
		var ret string
		return ret
	}
	return *o.AccessVersion
}

// GetAccessVersionOk returns a tuple with the AccessVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetAccessVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AccessVersion) {
		return nil, false
	}
	return o.AccessVersion, true
}

// HasAccessVersion returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasAccessVersion() bool {
	if o != nil && !IsNil(o.AccessVersion) {
		return true
	}

	return false
}

// SetAccessVersion gets a reference to the given string and assigns it to the AccessVersion field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetAccessVersion(v string) {
	o.AccessVersion = &v
}

// GetBizNoPrefix returns the BizNoPrefix field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetBizNoPrefix() string {
	if o == nil || IsNil(o.BizNoPrefix) {
		var ret string
		return ret
	}
	return *o.BizNoPrefix
}

// GetBizNoPrefixOk returns a tuple with the BizNoPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetBizNoPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.BizNoPrefix) {
		return nil, false
	}
	return o.BizNoPrefix, true
}

// HasBizNoPrefix returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasBizNoPrefix() bool {
	if o != nil && !IsNil(o.BizNoPrefix) {
		return true
	}

	return false
}

// SetBizNoPrefix gets a reference to the given string and assigns it to the BizNoPrefix field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetBizNoPrefix(v string) {
	o.BizNoPrefix = &v
}

// GetBizNoSuffixLen returns the BizNoSuffixLen field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetBizNoSuffixLen() string {
	if o == nil || IsNil(o.BizNoSuffixLen) {
		var ret string
		return ret
	}
	return *o.BizNoSuffixLen
}

// GetBizNoSuffixLenOk returns a tuple with the BizNoSuffixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetBizNoSuffixLenOk() (*string, bool) {
	if o == nil || IsNil(o.BizNoSuffixLen) {
		return nil, false
	}
	return o.BizNoSuffixLen, true
}

// HasBizNoSuffixLen returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasBizNoSuffixLen() bool {
	if o != nil && !IsNil(o.BizNoSuffixLen) {
		return true
	}

	return false
}

// SetBizNoSuffixLen gets a reference to the given string and assigns it to the BizNoSuffixLen field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetBizNoSuffixLen(v string) {
	o.BizNoSuffixLen = &v
}

// GetCardActionList returns the CardActionList field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardActionList() []TemplateActionInfoDTO {
	if o == nil || IsNil(o.CardActionList) {
		var ret []TemplateActionInfoDTO
		return ret
	}
	return o.CardActionList
}

// GetCardActionListOk returns a tuple with the CardActionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardActionListOk() ([]TemplateActionInfoDTO, bool) {
	if o == nil || IsNil(o.CardActionList) {
		return nil, false
	}
	return o.CardActionList, true
}

// HasCardActionList returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasCardActionList() bool {
	if o != nil && !IsNil(o.CardActionList) {
		return true
	}

	return false
}

// SetCardActionList gets a reference to the given []TemplateActionInfoDTO and assigns it to the CardActionList field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetCardActionList(v []TemplateActionInfoDTO) {
	o.CardActionList = v
}

// GetCardLevelConfs returns the CardLevelConfs field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardLevelConfs() []TemplateCardLevelConfDTO {
	if o == nil || IsNil(o.CardLevelConfs) {
		var ret []TemplateCardLevelConfDTO
		return ret
	}
	return o.CardLevelConfs
}

// GetCardLevelConfsOk returns a tuple with the CardLevelConfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardLevelConfsOk() ([]TemplateCardLevelConfDTO, bool) {
	if o == nil || IsNil(o.CardLevelConfs) {
		return nil, false
	}
	return o.CardLevelConfs, true
}

// HasCardLevelConfs returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasCardLevelConfs() bool {
	if o != nil && !IsNil(o.CardLevelConfs) {
		return true
	}

	return false
}

// SetCardLevelConfs gets a reference to the given []TemplateCardLevelConfDTO and assigns it to the CardLevelConfs field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetCardLevelConfs(v []TemplateCardLevelConfDTO) {
	o.CardLevelConfs = v
}

// GetCardSpecTag returns the CardSpecTag field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardSpecTag() string {
	if o == nil || IsNil(o.CardSpecTag) {
		var ret string
		return ret
	}
	return *o.CardSpecTag
}

// GetCardSpecTagOk returns a tuple with the CardSpecTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardSpecTagOk() (*string, bool) {
	if o == nil || IsNil(o.CardSpecTag) {
		return nil, false
	}
	return o.CardSpecTag, true
}

// HasCardSpecTag returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasCardSpecTag() bool {
	if o != nil && !IsNil(o.CardSpecTag) {
		return true
	}

	return false
}

// SetCardSpecTag gets a reference to the given string and assigns it to the CardSpecTag field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetCardSpecTag(v string) {
	o.CardSpecTag = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetCardType(v string) {
	o.CardType = &v
}

// GetColumnInfoList returns the ColumnInfoList field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetColumnInfoList() []TemplateColumnInfoDTO {
	if o == nil || IsNil(o.ColumnInfoList) {
		var ret []TemplateColumnInfoDTO
		return ret
	}
	return o.ColumnInfoList
}

// GetColumnInfoListOk returns a tuple with the ColumnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetColumnInfoListOk() ([]TemplateColumnInfoDTO, bool) {
	if o == nil || IsNil(o.ColumnInfoList) {
		return nil, false
	}
	return o.ColumnInfoList, true
}

// HasColumnInfoList returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasColumnInfoList() bool {
	if o != nil && !IsNil(o.ColumnInfoList) {
		return true
	}

	return false
}

// SetColumnInfoList gets a reference to the given []TemplateColumnInfoDTO and assigns it to the ColumnInfoList field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetColumnInfoList(v []TemplateColumnInfoDTO) {
	o.ColumnInfoList = v
}

// GetFieldRuleList returns the FieldRuleList field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetFieldRuleList() []TemplateFieldRuleDTO {
	if o == nil || IsNil(o.FieldRuleList) {
		var ret []TemplateFieldRuleDTO
		return ret
	}
	return o.FieldRuleList
}

// GetFieldRuleListOk returns a tuple with the FieldRuleList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetFieldRuleListOk() ([]TemplateFieldRuleDTO, bool) {
	if o == nil || IsNil(o.FieldRuleList) {
		return nil, false
	}
	return o.FieldRuleList, true
}

// HasFieldRuleList returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasFieldRuleList() bool {
	if o != nil && !IsNil(o.FieldRuleList) {
		return true
	}

	return false
}

// SetFieldRuleList gets a reference to the given []TemplateFieldRuleDTO and assigns it to the FieldRuleList field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetFieldRuleList(v []TemplateFieldRuleDTO) {
	o.FieldRuleList = v
}

// GetMdcodeNotifyConf returns the MdcodeNotifyConf field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetMdcodeNotifyConf() TemplateMdcodeNotifyConfDTO {
	if o == nil || IsNil(o.MdcodeNotifyConf) {
		var ret TemplateMdcodeNotifyConfDTO
		return ret
	}
	return *o.MdcodeNotifyConf
}

// GetMdcodeNotifyConfOk returns a tuple with the MdcodeNotifyConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetMdcodeNotifyConfOk() (*TemplateMdcodeNotifyConfDTO, bool) {
	if o == nil || IsNil(o.MdcodeNotifyConf) {
		return nil, false
	}
	return o.MdcodeNotifyConf, true
}

// HasMdcodeNotifyConf returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasMdcodeNotifyConf() bool {
	if o != nil && !IsNil(o.MdcodeNotifyConf) {
		return true
	}

	return false
}

// SetMdcodeNotifyConf gets a reference to the given TemplateMdcodeNotifyConfDTO and assigns it to the MdcodeNotifyConf field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetMdcodeNotifyConf(v TemplateMdcodeNotifyConfDTO) {
	o.MdcodeNotifyConf = &v
}

// GetOpenCardConf returns the OpenCardConf field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetOpenCardConf() TemplateOpenCardConfDTO {
	if o == nil || IsNil(o.OpenCardConf) {
		var ret TemplateOpenCardConfDTO
		return ret
	}
	return *o.OpenCardConf
}

// GetOpenCardConfOk returns a tuple with the OpenCardConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetOpenCardConfOk() (*TemplateOpenCardConfDTO, bool) {
	if o == nil || IsNil(o.OpenCardConf) {
		return nil, false
	}
	return o.OpenCardConf, true
}

// HasOpenCardConf returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasOpenCardConf() bool {
	if o != nil && !IsNil(o.OpenCardConf) {
		return true
	}

	return false
}

// SetOpenCardConf gets a reference to the given TemplateOpenCardConfDTO and assigns it to the OpenCardConf field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetOpenCardConf(v TemplateOpenCardConfDTO) {
	o.OpenCardConf = &v
}

// GetPaidOuterCardConf returns the PaidOuterCardConf field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetPaidOuterCardConf() PaidOuterCardTemplateConfDTO {
	if o == nil || IsNil(o.PaidOuterCardConf) {
		var ret PaidOuterCardTemplateConfDTO
		return ret
	}
	return *o.PaidOuterCardConf
}

// GetPaidOuterCardConfOk returns a tuple with the PaidOuterCardConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetPaidOuterCardConfOk() (*PaidOuterCardTemplateConfDTO, bool) {
	if o == nil || IsNil(o.PaidOuterCardConf) {
		return nil, false
	}
	return o.PaidOuterCardConf, true
}

// HasPaidOuterCardConf returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasPaidOuterCardConf() bool {
	if o != nil && !IsNil(o.PaidOuterCardConf) {
		return true
	}

	return false
}

// SetPaidOuterCardConf gets a reference to the given PaidOuterCardTemplateConfDTO and assigns it to the PaidOuterCardConf field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetPaidOuterCardConf(v PaidOuterCardTemplateConfDTO) {
	o.PaidOuterCardConf = &v
}

// GetPubChannels returns the PubChannels field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetPubChannels() []PubChannelDTO {
	if o == nil || IsNil(o.PubChannels) {
		var ret []PubChannelDTO
		return ret
	}
	return o.PubChannels
}

// GetPubChannelsOk returns a tuple with the PubChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetPubChannelsOk() ([]PubChannelDTO, bool) {
	if o == nil || IsNil(o.PubChannels) {
		return nil, false
	}
	return o.PubChannels, true
}

// HasPubChannels returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasPubChannels() bool {
	if o != nil && !IsNil(o.PubChannels) {
		return true
	}

	return false
}

// SetPubChannels gets a reference to the given []PubChannelDTO and assigns it to the PubChannels field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetPubChannels(v []PubChannelDTO) {
	o.PubChannels = v
}

// GetServiceLabelList returns the ServiceLabelList field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetServiceLabelList() []string {
	if o == nil || IsNil(o.ServiceLabelList) {
		var ret []string
		return ret
	}
	return o.ServiceLabelList
}

// GetServiceLabelListOk returns a tuple with the ServiceLabelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetServiceLabelListOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceLabelList) {
		return nil, false
	}
	return o.ServiceLabelList, true
}

// HasServiceLabelList returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasServiceLabelList() bool {
	if o != nil && !IsNil(o.ServiceLabelList) {
		return true
	}

	return false
}

// SetServiceLabelList gets a reference to the given []string and assigns it to the ServiceLabelList field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetServiceLabelList(v []string) {
	o.ServiceLabelList = v
}

// GetShopIds returns the ShopIds field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetShopIds() []string {
	if o == nil || IsNil(o.ShopIds) {
		var ret []string
		return ret
	}
	return o.ShopIds
}

// GetShopIdsOk returns a tuple with the ShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShopIds) {
		return nil, false
	}
	return o.ShopIds, true
}

// HasShopIds returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasShopIds() bool {
	if o != nil && !IsNil(o.ShopIds) {
		return true
	}

	return false
}

// SetShopIds gets a reference to the given []string and assigns it to the ShopIds field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetShopIds(v []string) {
	o.ShopIds = v
}

// GetSpiAppId returns the SpiAppId field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetSpiAppId() string {
	if o == nil || IsNil(o.SpiAppId) {
		var ret string
		return ret
	}
	return *o.SpiAppId
}

// GetSpiAppIdOk returns a tuple with the SpiAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetSpiAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpiAppId) {
		return nil, false
	}
	return o.SpiAppId, true
}

// HasSpiAppId returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasSpiAppId() bool {
	if o != nil && !IsNil(o.SpiAppId) {
		return true
	}

	return false
}

// SetSpiAppId gets a reference to the given string and assigns it to the SpiAppId field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetSpiAppId(v string) {
	o.SpiAppId = &v
}

// GetTemplateBenefitInfo returns the TemplateBenefitInfo field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetTemplateBenefitInfo() []TemplateBenefitInfoDTO {
	if o == nil || IsNil(o.TemplateBenefitInfo) {
		var ret []TemplateBenefitInfoDTO
		return ret
	}
	return o.TemplateBenefitInfo
}

// GetTemplateBenefitInfoOk returns a tuple with the TemplateBenefitInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetTemplateBenefitInfoOk() ([]TemplateBenefitInfoDTO, bool) {
	if o == nil || IsNil(o.TemplateBenefitInfo) {
		return nil, false
	}
	return o.TemplateBenefitInfo, true
}

// HasTemplateBenefitInfo returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasTemplateBenefitInfo() bool {
	if o != nil && !IsNil(o.TemplateBenefitInfo) {
		return true
	}

	return false
}

// SetTemplateBenefitInfo gets a reference to the given []TemplateBenefitInfoDTO and assigns it to the TemplateBenefitInfo field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetTemplateBenefitInfo(v []TemplateBenefitInfoDTO) {
	o.TemplateBenefitInfo = v
}

// GetTemplateFormConfig returns the TemplateFormConfig field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetTemplateFormConfig() TemplateFormConfig {
	if o == nil || IsNil(o.TemplateFormConfig) {
		var ret TemplateFormConfig
		return ret
	}
	return *o.TemplateFormConfig
}

// GetTemplateFormConfigOk returns a tuple with the TemplateFormConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetTemplateFormConfigOk() (*TemplateFormConfig, bool) {
	if o == nil || IsNil(o.TemplateFormConfig) {
		return nil, false
	}
	return o.TemplateFormConfig, true
}

// HasTemplateFormConfig returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasTemplateFormConfig() bool {
	if o != nil && !IsNil(o.TemplateFormConfig) {
		return true
	}

	return false
}

// SetTemplateFormConfig gets a reference to the given TemplateFormConfig and assigns it to the TemplateFormConfig field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetTemplateFormConfig(v TemplateFormConfig) {
	o.TemplateFormConfig = &v
}

// GetTemplateStyleInfo returns the TemplateStyleInfo field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetTemplateStyleInfo() TemplateStyleInfoDTO {
	if o == nil || IsNil(o.TemplateStyleInfo) {
		var ret TemplateStyleInfoDTO
		return ret
	}
	return *o.TemplateStyleInfo
}

// GetTemplateStyleInfoOk returns a tuple with the TemplateStyleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) GetTemplateStyleInfoOk() (*TemplateStyleInfoDTO, bool) {
	if o == nil || IsNil(o.TemplateStyleInfo) {
		return nil, false
	}
	return o.TemplateStyleInfo, true
}

// HasTemplateStyleInfo returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateQueryResponseModel) HasTemplateStyleInfo() bool {
	if o != nil && !IsNil(o.TemplateStyleInfo) {
		return true
	}

	return false
}

// SetTemplateStyleInfo gets a reference to the given TemplateStyleInfoDTO and assigns it to the TemplateStyleInfo field.
func (o *AlipayMarketingCardTemplateQueryResponseModel) SetTemplateStyleInfo(v TemplateStyleInfoDTO) {
	o.TemplateStyleInfo = &v
}

func (o AlipayMarketingCardTemplateQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardTemplateQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessVersion) {
		toSerialize["access_version"] = o.AccessVersion
	}
	if !IsNil(o.BizNoPrefix) {
		toSerialize["biz_no_prefix"] = o.BizNoPrefix
	}
	if !IsNil(o.BizNoSuffixLen) {
		toSerialize["biz_no_suffix_len"] = o.BizNoSuffixLen
	}
	if !IsNil(o.CardActionList) {
		toSerialize["card_action_list"] = o.CardActionList
	}
	if !IsNil(o.CardLevelConfs) {
		toSerialize["card_level_confs"] = o.CardLevelConfs
	}
	if !IsNil(o.CardSpecTag) {
		toSerialize["card_spec_tag"] = o.CardSpecTag
	}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.ColumnInfoList) {
		toSerialize["column_info_list"] = o.ColumnInfoList
	}
	if !IsNil(o.FieldRuleList) {
		toSerialize["field_rule_list"] = o.FieldRuleList
	}
	if !IsNil(o.MdcodeNotifyConf) {
		toSerialize["mdcode_notify_conf"] = o.MdcodeNotifyConf
	}
	if !IsNil(o.OpenCardConf) {
		toSerialize["open_card_conf"] = o.OpenCardConf
	}
	if !IsNil(o.PaidOuterCardConf) {
		toSerialize["paid_outer_card_conf"] = o.PaidOuterCardConf
	}
	if !IsNil(o.PubChannels) {
		toSerialize["pub_channels"] = o.PubChannels
	}
	if !IsNil(o.ServiceLabelList) {
		toSerialize["service_label_list"] = o.ServiceLabelList
	}
	if !IsNil(o.ShopIds) {
		toSerialize["shop_ids"] = o.ShopIds
	}
	if !IsNil(o.SpiAppId) {
		toSerialize["spi_app_id"] = o.SpiAppId
	}
	if !IsNil(o.TemplateBenefitInfo) {
		toSerialize["template_benefit_info"] = o.TemplateBenefitInfo
	}
	if !IsNil(o.TemplateFormConfig) {
		toSerialize["template_form_config"] = o.TemplateFormConfig
	}
	if !IsNil(o.TemplateStyleInfo) {
		toSerialize["template_style_info"] = o.TemplateStyleInfo
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardTemplateQueryResponseModel struct {
	value *AlipayMarketingCardTemplateQueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCardTemplateQueryResponseModel) Get() *AlipayMarketingCardTemplateQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCardTemplateQueryResponseModel) Set(val *AlipayMarketingCardTemplateQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardTemplateQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardTemplateQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardTemplateQueryResponseModel(val *AlipayMarketingCardTemplateQueryResponseModel) *NullableAlipayMarketingCardTemplateQueryResponseModel {
	return &NullableAlipayMarketingCardTemplateQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardTemplateQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardTemplateQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

