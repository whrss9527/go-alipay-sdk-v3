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

// checks if the AlipayMarketingCardTemplateModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardTemplateModifyModel{}

// AlipayMarketingCardTemplateModifyModel struct for AlipayMarketingCardTemplateModifyModel
type AlipayMarketingCardTemplateModifyModel struct {
	// 接口接入版本
	AccessVersion *string `json:"access_version,omitempty"`
	// 业务卡号前缀，由商户指定  支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀
	BizNoPrefix *string `json:"biz_no_prefix,omitempty"`
	// 业务卡号后缀长度，与biz_no_prefix配合，扣除系统预留2位，剩下对应seq长度。在生成卡号时，若seq位数不足前置补0，若seq位数超出则以实际为准。举例：设为10，其中8位用于生成seq，可覆盖1亿用户，即使seq超过1亿，生成卡号也不报错，但总长度不得超过32位。建议按需设置合适的值，以获得长度一致的业务卡号，建议长度20，性能更好。单位：/位
	BizNoSuffixLen *string `json:"biz_no_suffix_len,omitempty"`
	// \"卡行动点配置； 行动点，即用户可点击跳转的区块，类似按钮控件的交互； 单张卡最多定制4个行动点。\"
	CardActionList []TemplateActionInfoDTO `json:"card_action_list,omitempty"`
	// 卡级别配置
	CardLevelConf []TemplateCardLevelConfDTO `json:"card_level_conf,omitempty"`
	// 卡特定标签，只供特定业务使用，通常接入无需关注
	CardSpecTag *string `json:"card_spec_tag,omitempty"`
	// 栏位信息（卡包详情页面展现的栏位）
	ColumnInfoList []TemplateColumnInfoDTO `json:"column_info_list,omitempty"`
	// \"字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE 注意：商家会员卡场景不支持修改该内容。 \"
	FieldRuleList     []TemplateFieldRuleDTO        `json:"field_rule_list,omitempty"`
	MdcodeNotifyConf  *TemplateMdcodeNotifyConfDTO  `json:"mdcode_notify_conf,omitempty"`
	OpenCardConf      *TemplateOpenCardConfDTO      `json:"open_card_conf,omitempty"`
	PaidOuterCardConf *PaidOuterCardTemplateConfDTO `json:"paid_outer_card_conf,omitempty"`
	// 卡模板投放渠道
	PubChannels []PubChannelDTO `json:"pub_channels,omitempty"`
	// 请求 ID，商户自定义且需确保每次请求唯一。
	RequestId *string `json:"request_id,omitempty"`
	// 会员卡上架门店 id（支付宝门店id），即发放会员卡的商家门店id。不传则保持创建模板时信息。 注意：不支持删除已有shop_id，仅支持新增。
	ShopIds []string `json:"shop_ids,omitempty"`
	// spi应用id，为实现spi.alipay.user.opencard.get接口的app_id。若是第三方代理模式，可以设置成服务商的 APPID 或者是商家自己的 APPID； 若不是第三方代理模式，只能设置商家自己的 APPID。
	SpiAppId *string `json:"spi_app_id,omitempty"`
	// \"权益信息， 1、在卡包的卡详情页面会自动添加权益栏位，展现会员卡特权， 2、如果添加门店渠道，则可在门店页展现会员卡的权益\"
	TemplateBenefitInfo []TemplateBenefitInfoDTO `json:"template_benefit_info,omitempty"`
	TemplateFormConfig  *TemplateFormConfig      `json:"template_form_config,omitempty"`
	// 会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。
	TemplateId        *string               `json:"template_id,omitempty"`
	TemplateStyleInfo *TemplateStyleInfoDTO `json:"template_style_info,omitempty"`
	// 卡包详情页面中展现出的卡码（可用于扫码核销）  (1) 静态码 qrcode: 二维码，扫码得商户开卡传入的external_card_no barcode: 条形码，扫码得商户开卡传入的external_card_no  (2) 动态码-支付宝生成码值(动态码会在2分钟左右后过期) dqrcode: 动态二维码，扫码得到的码值可配合会员卡查询接口使用 dbarcode: 动态条形码，扫码得到的码值可配合会员卡查询接口使用  (3) 动态码-商家自主生成码值（码值、时效性都由商户控制） mdqrcode: 商户动态二维码，扫码得商户自主传入的码值 mdbarcode: 商户动态条码，扫码得商户自主传入的码值。 如需使用商户动态码，请联系支付宝技术支持获取相关文档。
	WriteOffType *string `json:"write_off_type,omitempty"`
}

// NewAlipayMarketingCardTemplateModifyModel instantiates a new AlipayMarketingCardTemplateModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardTemplateModifyModel() *AlipayMarketingCardTemplateModifyModel {
	this := AlipayMarketingCardTemplateModifyModel{}
	return &this
}

// NewAlipayMarketingCardTemplateModifyModelWithDefaults instantiates a new AlipayMarketingCardTemplateModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardTemplateModifyModelWithDefaults() *AlipayMarketingCardTemplateModifyModel {
	this := AlipayMarketingCardTemplateModifyModel{}
	return &this
}

// GetAccessVersion returns the AccessVersion field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetAccessVersion() string {
	if o == nil || IsNil(o.AccessVersion) {
		var ret string
		return ret
	}
	return *o.AccessVersion
}

// GetAccessVersionOk returns a tuple with the AccessVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetAccessVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AccessVersion) {
		return nil, false
	}
	return o.AccessVersion, true
}

// HasAccessVersion returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasAccessVersion() bool {
	if o != nil && !IsNil(o.AccessVersion) {
		return true
	}

	return false
}

// SetAccessVersion gets a reference to the given string and assigns it to the AccessVersion field.
func (o *AlipayMarketingCardTemplateModifyModel) SetAccessVersion(v string) {
	o.AccessVersion = &v
}

// GetBizNoPrefix returns the BizNoPrefix field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetBizNoPrefix() string {
	if o == nil || IsNil(o.BizNoPrefix) {
		var ret string
		return ret
	}
	return *o.BizNoPrefix
}

// GetBizNoPrefixOk returns a tuple with the BizNoPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetBizNoPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.BizNoPrefix) {
		return nil, false
	}
	return o.BizNoPrefix, true
}

// HasBizNoPrefix returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasBizNoPrefix() bool {
	if o != nil && !IsNil(o.BizNoPrefix) {
		return true
	}

	return false
}

// SetBizNoPrefix gets a reference to the given string and assigns it to the BizNoPrefix field.
func (o *AlipayMarketingCardTemplateModifyModel) SetBizNoPrefix(v string) {
	o.BizNoPrefix = &v
}

// GetBizNoSuffixLen returns the BizNoSuffixLen field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetBizNoSuffixLen() string {
	if o == nil || IsNil(o.BizNoSuffixLen) {
		var ret string
		return ret
	}
	return *o.BizNoSuffixLen
}

// GetBizNoSuffixLenOk returns a tuple with the BizNoSuffixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetBizNoSuffixLenOk() (*string, bool) {
	if o == nil || IsNil(o.BizNoSuffixLen) {
		return nil, false
	}
	return o.BizNoSuffixLen, true
}

// HasBizNoSuffixLen returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasBizNoSuffixLen() bool {
	if o != nil && !IsNil(o.BizNoSuffixLen) {
		return true
	}

	return false
}

// SetBizNoSuffixLen gets a reference to the given string and assigns it to the BizNoSuffixLen field.
func (o *AlipayMarketingCardTemplateModifyModel) SetBizNoSuffixLen(v string) {
	o.BizNoSuffixLen = &v
}

// GetCardActionList returns the CardActionList field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetCardActionList() []TemplateActionInfoDTO {
	if o == nil || IsNil(o.CardActionList) {
		var ret []TemplateActionInfoDTO
		return ret
	}
	return o.CardActionList
}

// GetCardActionListOk returns a tuple with the CardActionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetCardActionListOk() ([]TemplateActionInfoDTO, bool) {
	if o == nil || IsNil(o.CardActionList) {
		return nil, false
	}
	return o.CardActionList, true
}

// HasCardActionList returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasCardActionList() bool {
	if o != nil && !IsNil(o.CardActionList) {
		return true
	}

	return false
}

// SetCardActionList gets a reference to the given []TemplateActionInfoDTO and assigns it to the CardActionList field.
func (o *AlipayMarketingCardTemplateModifyModel) SetCardActionList(v []TemplateActionInfoDTO) {
	o.CardActionList = v
}

// GetCardLevelConf returns the CardLevelConf field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetCardLevelConf() []TemplateCardLevelConfDTO {
	if o == nil || IsNil(o.CardLevelConf) {
		var ret []TemplateCardLevelConfDTO
		return ret
	}
	return o.CardLevelConf
}

// GetCardLevelConfOk returns a tuple with the CardLevelConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetCardLevelConfOk() ([]TemplateCardLevelConfDTO, bool) {
	if o == nil || IsNil(o.CardLevelConf) {
		return nil, false
	}
	return o.CardLevelConf, true
}

// HasCardLevelConf returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasCardLevelConf() bool {
	if o != nil && !IsNil(o.CardLevelConf) {
		return true
	}

	return false
}

// SetCardLevelConf gets a reference to the given []TemplateCardLevelConfDTO and assigns it to the CardLevelConf field.
func (o *AlipayMarketingCardTemplateModifyModel) SetCardLevelConf(v []TemplateCardLevelConfDTO) {
	o.CardLevelConf = v
}

// GetCardSpecTag returns the CardSpecTag field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetCardSpecTag() string {
	if o == nil || IsNil(o.CardSpecTag) {
		var ret string
		return ret
	}
	return *o.CardSpecTag
}

// GetCardSpecTagOk returns a tuple with the CardSpecTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetCardSpecTagOk() (*string, bool) {
	if o == nil || IsNil(o.CardSpecTag) {
		return nil, false
	}
	return o.CardSpecTag, true
}

// HasCardSpecTag returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasCardSpecTag() bool {
	if o != nil && !IsNil(o.CardSpecTag) {
		return true
	}

	return false
}

// SetCardSpecTag gets a reference to the given string and assigns it to the CardSpecTag field.
func (o *AlipayMarketingCardTemplateModifyModel) SetCardSpecTag(v string) {
	o.CardSpecTag = &v
}

// GetColumnInfoList returns the ColumnInfoList field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetColumnInfoList() []TemplateColumnInfoDTO {
	if o == nil || IsNil(o.ColumnInfoList) {
		var ret []TemplateColumnInfoDTO
		return ret
	}
	return o.ColumnInfoList
}

// GetColumnInfoListOk returns a tuple with the ColumnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetColumnInfoListOk() ([]TemplateColumnInfoDTO, bool) {
	if o == nil || IsNil(o.ColumnInfoList) {
		return nil, false
	}
	return o.ColumnInfoList, true
}

// HasColumnInfoList returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasColumnInfoList() bool {
	if o != nil && !IsNil(o.ColumnInfoList) {
		return true
	}

	return false
}

// SetColumnInfoList gets a reference to the given []TemplateColumnInfoDTO and assigns it to the ColumnInfoList field.
func (o *AlipayMarketingCardTemplateModifyModel) SetColumnInfoList(v []TemplateColumnInfoDTO) {
	o.ColumnInfoList = v
}

// GetFieldRuleList returns the FieldRuleList field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetFieldRuleList() []TemplateFieldRuleDTO {
	if o == nil || IsNil(o.FieldRuleList) {
		var ret []TemplateFieldRuleDTO
		return ret
	}
	return o.FieldRuleList
}

// GetFieldRuleListOk returns a tuple with the FieldRuleList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetFieldRuleListOk() ([]TemplateFieldRuleDTO, bool) {
	if o == nil || IsNil(o.FieldRuleList) {
		return nil, false
	}
	return o.FieldRuleList, true
}

// HasFieldRuleList returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasFieldRuleList() bool {
	if o != nil && !IsNil(o.FieldRuleList) {
		return true
	}

	return false
}

// SetFieldRuleList gets a reference to the given []TemplateFieldRuleDTO and assigns it to the FieldRuleList field.
func (o *AlipayMarketingCardTemplateModifyModel) SetFieldRuleList(v []TemplateFieldRuleDTO) {
	o.FieldRuleList = v
}

// GetMdcodeNotifyConf returns the MdcodeNotifyConf field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetMdcodeNotifyConf() TemplateMdcodeNotifyConfDTO {
	if o == nil || IsNil(o.MdcodeNotifyConf) {
		var ret TemplateMdcodeNotifyConfDTO
		return ret
	}
	return *o.MdcodeNotifyConf
}

// GetMdcodeNotifyConfOk returns a tuple with the MdcodeNotifyConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetMdcodeNotifyConfOk() (*TemplateMdcodeNotifyConfDTO, bool) {
	if o == nil || IsNil(o.MdcodeNotifyConf) {
		return nil, false
	}
	return o.MdcodeNotifyConf, true
}

// HasMdcodeNotifyConf returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasMdcodeNotifyConf() bool {
	if o != nil && !IsNil(o.MdcodeNotifyConf) {
		return true
	}

	return false
}

// SetMdcodeNotifyConf gets a reference to the given TemplateMdcodeNotifyConfDTO and assigns it to the MdcodeNotifyConf field.
func (o *AlipayMarketingCardTemplateModifyModel) SetMdcodeNotifyConf(v TemplateMdcodeNotifyConfDTO) {
	o.MdcodeNotifyConf = &v
}

// GetOpenCardConf returns the OpenCardConf field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetOpenCardConf() TemplateOpenCardConfDTO {
	if o == nil || IsNil(o.OpenCardConf) {
		var ret TemplateOpenCardConfDTO
		return ret
	}
	return *o.OpenCardConf
}

// GetOpenCardConfOk returns a tuple with the OpenCardConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetOpenCardConfOk() (*TemplateOpenCardConfDTO, bool) {
	if o == nil || IsNil(o.OpenCardConf) {
		return nil, false
	}
	return o.OpenCardConf, true
}

// HasOpenCardConf returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasOpenCardConf() bool {
	if o != nil && !IsNil(o.OpenCardConf) {
		return true
	}

	return false
}

// SetOpenCardConf gets a reference to the given TemplateOpenCardConfDTO and assigns it to the OpenCardConf field.
func (o *AlipayMarketingCardTemplateModifyModel) SetOpenCardConf(v TemplateOpenCardConfDTO) {
	o.OpenCardConf = &v
}

// GetPaidOuterCardConf returns the PaidOuterCardConf field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetPaidOuterCardConf() PaidOuterCardTemplateConfDTO {
	if o == nil || IsNil(o.PaidOuterCardConf) {
		var ret PaidOuterCardTemplateConfDTO
		return ret
	}
	return *o.PaidOuterCardConf
}

// GetPaidOuterCardConfOk returns a tuple with the PaidOuterCardConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetPaidOuterCardConfOk() (*PaidOuterCardTemplateConfDTO, bool) {
	if o == nil || IsNil(o.PaidOuterCardConf) {
		return nil, false
	}
	return o.PaidOuterCardConf, true
}

// HasPaidOuterCardConf returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasPaidOuterCardConf() bool {
	if o != nil && !IsNil(o.PaidOuterCardConf) {
		return true
	}

	return false
}

// SetPaidOuterCardConf gets a reference to the given PaidOuterCardTemplateConfDTO and assigns it to the PaidOuterCardConf field.
func (o *AlipayMarketingCardTemplateModifyModel) SetPaidOuterCardConf(v PaidOuterCardTemplateConfDTO) {
	o.PaidOuterCardConf = &v
}

// GetPubChannels returns the PubChannels field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetPubChannels() []PubChannelDTO {
	if o == nil || IsNil(o.PubChannels) {
		var ret []PubChannelDTO
		return ret
	}
	return o.PubChannels
}

// GetPubChannelsOk returns a tuple with the PubChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetPubChannelsOk() ([]PubChannelDTO, bool) {
	if o == nil || IsNil(o.PubChannels) {
		return nil, false
	}
	return o.PubChannels, true
}

// HasPubChannels returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasPubChannels() bool {
	if o != nil && !IsNil(o.PubChannels) {
		return true
	}

	return false
}

// SetPubChannels gets a reference to the given []PubChannelDTO and assigns it to the PubChannels field.
func (o *AlipayMarketingCardTemplateModifyModel) SetPubChannels(v []PubChannelDTO) {
	o.PubChannels = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AlipayMarketingCardTemplateModifyModel) SetRequestId(v string) {
	o.RequestId = &v
}

// GetShopIds returns the ShopIds field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetShopIds() []string {
	if o == nil || IsNil(o.ShopIds) {
		var ret []string
		return ret
	}
	return o.ShopIds
}

// GetShopIdsOk returns a tuple with the ShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShopIds) {
		return nil, false
	}
	return o.ShopIds, true
}

// HasShopIds returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasShopIds() bool {
	if o != nil && !IsNil(o.ShopIds) {
		return true
	}

	return false
}

// SetShopIds gets a reference to the given []string and assigns it to the ShopIds field.
func (o *AlipayMarketingCardTemplateModifyModel) SetShopIds(v []string) {
	o.ShopIds = v
}

// GetSpiAppId returns the SpiAppId field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetSpiAppId() string {
	if o == nil || IsNil(o.SpiAppId) {
		var ret string
		return ret
	}
	return *o.SpiAppId
}

// GetSpiAppIdOk returns a tuple with the SpiAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetSpiAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpiAppId) {
		return nil, false
	}
	return o.SpiAppId, true
}

// HasSpiAppId returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasSpiAppId() bool {
	if o != nil && !IsNil(o.SpiAppId) {
		return true
	}

	return false
}

// SetSpiAppId gets a reference to the given string and assigns it to the SpiAppId field.
func (o *AlipayMarketingCardTemplateModifyModel) SetSpiAppId(v string) {
	o.SpiAppId = &v
}

// GetTemplateBenefitInfo returns the TemplateBenefitInfo field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateBenefitInfo() []TemplateBenefitInfoDTO {
	if o == nil || IsNil(o.TemplateBenefitInfo) {
		var ret []TemplateBenefitInfoDTO
		return ret
	}
	return o.TemplateBenefitInfo
}

// GetTemplateBenefitInfoOk returns a tuple with the TemplateBenefitInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateBenefitInfoOk() ([]TemplateBenefitInfoDTO, bool) {
	if o == nil || IsNil(o.TemplateBenefitInfo) {
		return nil, false
	}
	return o.TemplateBenefitInfo, true
}

// HasTemplateBenefitInfo returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasTemplateBenefitInfo() bool {
	if o != nil && !IsNil(o.TemplateBenefitInfo) {
		return true
	}

	return false
}

// SetTemplateBenefitInfo gets a reference to the given []TemplateBenefitInfoDTO and assigns it to the TemplateBenefitInfo field.
func (o *AlipayMarketingCardTemplateModifyModel) SetTemplateBenefitInfo(v []TemplateBenefitInfoDTO) {
	o.TemplateBenefitInfo = v
}

// GetTemplateFormConfig returns the TemplateFormConfig field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateFormConfig() TemplateFormConfig {
	if o == nil || IsNil(o.TemplateFormConfig) {
		var ret TemplateFormConfig
		return ret
	}
	return *o.TemplateFormConfig
}

// GetTemplateFormConfigOk returns a tuple with the TemplateFormConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateFormConfigOk() (*TemplateFormConfig, bool) {
	if o == nil || IsNil(o.TemplateFormConfig) {
		return nil, false
	}
	return o.TemplateFormConfig, true
}

// HasTemplateFormConfig returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasTemplateFormConfig() bool {
	if o != nil && !IsNil(o.TemplateFormConfig) {
		return true
	}

	return false
}

// SetTemplateFormConfig gets a reference to the given TemplateFormConfig and assigns it to the TemplateFormConfig field.
func (o *AlipayMarketingCardTemplateModifyModel) SetTemplateFormConfig(v TemplateFormConfig) {
	o.TemplateFormConfig = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *AlipayMarketingCardTemplateModifyModel) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateStyleInfo returns the TemplateStyleInfo field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateStyleInfo() TemplateStyleInfoDTO {
	if o == nil || IsNil(o.TemplateStyleInfo) {
		var ret TemplateStyleInfoDTO
		return ret
	}
	return *o.TemplateStyleInfo
}

// GetTemplateStyleInfoOk returns a tuple with the TemplateStyleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetTemplateStyleInfoOk() (*TemplateStyleInfoDTO, bool) {
	if o == nil || IsNil(o.TemplateStyleInfo) {
		return nil, false
	}
	return o.TemplateStyleInfo, true
}

// HasTemplateStyleInfo returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasTemplateStyleInfo() bool {
	if o != nil && !IsNil(o.TemplateStyleInfo) {
		return true
	}

	return false
}

// SetTemplateStyleInfo gets a reference to the given TemplateStyleInfoDTO and assigns it to the TemplateStyleInfo field.
func (o *AlipayMarketingCardTemplateModifyModel) SetTemplateStyleInfo(v TemplateStyleInfoDTO) {
	o.TemplateStyleInfo = &v
}

// GetWriteOffType returns the WriteOffType field value if set, zero value otherwise.
func (o *AlipayMarketingCardTemplateModifyModel) GetWriteOffType() string {
	if o == nil || IsNil(o.WriteOffType) {
		var ret string
		return ret
	}
	return *o.WriteOffType
}

// GetWriteOffTypeOk returns a tuple with the WriteOffType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardTemplateModifyModel) GetWriteOffTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WriteOffType) {
		return nil, false
	}
	return o.WriteOffType, true
}

// HasWriteOffType returns a boolean if a field has been set.
func (o *AlipayMarketingCardTemplateModifyModel) HasWriteOffType() bool {
	if o != nil && !IsNil(o.WriteOffType) {
		return true
	}

	return false
}

// SetWriteOffType gets a reference to the given string and assigns it to the WriteOffType field.
func (o *AlipayMarketingCardTemplateModifyModel) SetWriteOffType(v string) {
	o.WriteOffType = &v
}

func (o AlipayMarketingCardTemplateModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardTemplateModifyModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CardLevelConf) {
		toSerialize["card_level_conf"] = o.CardLevelConf
	}
	if !IsNil(o.CardSpecTag) {
		toSerialize["card_spec_tag"] = o.CardSpecTag
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
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
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
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.TemplateStyleInfo) {
		toSerialize["template_style_info"] = o.TemplateStyleInfo
	}
	if !IsNil(o.WriteOffType) {
		toSerialize["write_off_type"] = o.WriteOffType
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardTemplateModifyModel struct {
	value *AlipayMarketingCardTemplateModifyModel
	isSet bool
}

func (v NullableAlipayMarketingCardTemplateModifyModel) Get() *AlipayMarketingCardTemplateModifyModel {
	return v.value
}

func (v *NullableAlipayMarketingCardTemplateModifyModel) Set(val *AlipayMarketingCardTemplateModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardTemplateModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardTemplateModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardTemplateModifyModel(val *AlipayMarketingCardTemplateModifyModel) *NullableAlipayMarketingCardTemplateModifyModel {
	return &NullableAlipayMarketingCardTemplateModifyModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardTemplateModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardTemplateModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
