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

// checks if the AlipayOpenMiniInnerappPluginservicePublishModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerappPluginservicePublishModel{}

// AlipayOpenMiniInnerappPluginservicePublishModel struct for AlipayOpenMiniInnerappPluginservicePublishModel
type AlipayOpenMiniInnerappPluginservicePublishModel struct {
	// 功能类型，目前支持的有：1039支付, 1040会员, 1041基础, 1056资金, 1058信用, 1111口碑, 330120安全, 360101营销
	AbilityTypeList []string `json:"ability_type_list,omitempty"`
	// 服务发布logo
	AppLogo *string `json:"app_logo,omitempty"`
	// 业务来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 服务描述
	Description *string `json:"description,omitempty"`
	// 移动端详情，用于能力中心展示
	DetailForClient *string `json:"detail_for_client,omitempty"`
	// pc端详情，用于能力中心展示
	DetailForPc *string `json:"detail_for_pc,omitempty"`
	// 插件id
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 可订购人群，1003个人, 1004企业，-1无限制
	SellCrowd *string `json:"sell_crowd,omitempty"`
	// 发布后是否展示，01展示（默认）、02隐藏
	ShowType *string `json:"show_type,omitempty"`
	// 服务标题
	Title *string `json:"title,omitempty"`
	// pc端url地址，不需要则为空
	VisitUrlForPc *string `json:"visit_url_for_pc,omitempty"`
}

// NewAlipayOpenMiniInnerappPluginservicePublishModel instantiates a new AlipayOpenMiniInnerappPluginservicePublishModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerappPluginservicePublishModel() *AlipayOpenMiniInnerappPluginservicePublishModel {
	this := AlipayOpenMiniInnerappPluginservicePublishModel{}
	return &this
}

// NewAlipayOpenMiniInnerappPluginservicePublishModelWithDefaults instantiates a new AlipayOpenMiniInnerappPluginservicePublishModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerappPluginservicePublishModelWithDefaults() *AlipayOpenMiniInnerappPluginservicePublishModel {
	this := AlipayOpenMiniInnerappPluginservicePublishModel{}
	return &this
}

// GetAbilityTypeList returns the AbilityTypeList field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetAbilityTypeList() []string {
	if o == nil || IsNil(o.AbilityTypeList) {
		var ret []string
		return ret
	}
	return o.AbilityTypeList
}

// GetAbilityTypeListOk returns a tuple with the AbilityTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetAbilityTypeListOk() ([]string, bool) {
	if o == nil || IsNil(o.AbilityTypeList) {
		return nil, false
	}
	return o.AbilityTypeList, true
}

// HasAbilityTypeList returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasAbilityTypeList() bool {
	if o != nil && !IsNil(o.AbilityTypeList) {
		return true
	}

	return false
}

// SetAbilityTypeList gets a reference to the given []string and assigns it to the AbilityTypeList field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetAbilityTypeList(v []string) {
	o.AbilityTypeList = v
}

// GetAppLogo returns the AppLogo field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetAppLogo() string {
	if o == nil || IsNil(o.AppLogo) {
		var ret string
		return ret
	}
	return *o.AppLogo
}

// GetAppLogoOk returns a tuple with the AppLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetAppLogoOk() (*string, bool) {
	if o == nil || IsNil(o.AppLogo) {
		return nil, false
	}
	return o.AppLogo, true
}

// HasAppLogo returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasAppLogo() bool {
	if o != nil && !IsNil(o.AppLogo) {
		return true
	}

	return false
}

// SetAppLogo gets a reference to the given string and assigns it to the AppLogo field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetAppLogo(v string) {
	o.AppLogo = &v
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetDescription(v string) {
	o.Description = &v
}

// GetDetailForClient returns the DetailForClient field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetDetailForClient() string {
	if o == nil || IsNil(o.DetailForClient) {
		var ret string
		return ret
	}
	return *o.DetailForClient
}

// GetDetailForClientOk returns a tuple with the DetailForClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetDetailForClientOk() (*string, bool) {
	if o == nil || IsNil(o.DetailForClient) {
		return nil, false
	}
	return o.DetailForClient, true
}

// HasDetailForClient returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasDetailForClient() bool {
	if o != nil && !IsNil(o.DetailForClient) {
		return true
	}

	return false
}

// SetDetailForClient gets a reference to the given string and assigns it to the DetailForClient field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetDetailForClient(v string) {
	o.DetailForClient = &v
}

// GetDetailForPc returns the DetailForPc field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetDetailForPc() string {
	if o == nil || IsNil(o.DetailForPc) {
		var ret string
		return ret
	}
	return *o.DetailForPc
}

// GetDetailForPcOk returns a tuple with the DetailForPc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetDetailForPcOk() (*string, bool) {
	if o == nil || IsNil(o.DetailForPc) {
		return nil, false
	}
	return o.DetailForPc, true
}

// HasDetailForPc returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasDetailForPc() bool {
	if o != nil && !IsNil(o.DetailForPc) {
		return true
	}

	return false
}

// SetDetailForPc gets a reference to the given string and assigns it to the DetailForPc field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetDetailForPc(v string) {
	o.DetailForPc = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetSellCrowd returns the SellCrowd field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetSellCrowd() string {
	if o == nil || IsNil(o.SellCrowd) {
		var ret string
		return ret
	}
	return *o.SellCrowd
}

// GetSellCrowdOk returns a tuple with the SellCrowd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetSellCrowdOk() (*string, bool) {
	if o == nil || IsNil(o.SellCrowd) {
		return nil, false
	}
	return o.SellCrowd, true
}

// HasSellCrowd returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasSellCrowd() bool {
	if o != nil && !IsNil(o.SellCrowd) {
		return true
	}

	return false
}

// SetSellCrowd gets a reference to the given string and assigns it to the SellCrowd field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetSellCrowd(v string) {
	o.SellCrowd = &v
}

// GetShowType returns the ShowType field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetShowType() string {
	if o == nil || IsNil(o.ShowType) {
		var ret string
		return ret
	}
	return *o.ShowType
}

// GetShowTypeOk returns a tuple with the ShowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetShowTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ShowType) {
		return nil, false
	}
	return o.ShowType, true
}

// HasShowType returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasShowType() bool {
	if o != nil && !IsNil(o.ShowType) {
		return true
	}

	return false
}

// SetShowType gets a reference to the given string and assigns it to the ShowType field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetShowType(v string) {
	o.ShowType = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetTitle(v string) {
	o.Title = &v
}

// GetVisitUrlForPc returns the VisitUrlForPc field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetVisitUrlForPc() string {
	if o == nil || IsNil(o.VisitUrlForPc) {
		var ret string
		return ret
	}
	return *o.VisitUrlForPc
}

// GetVisitUrlForPcOk returns a tuple with the VisitUrlForPc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) GetVisitUrlForPcOk() (*string, bool) {
	if o == nil || IsNil(o.VisitUrlForPc) {
		return nil, false
	}
	return o.VisitUrlForPc, true
}

// HasVisitUrlForPc returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) HasVisitUrlForPc() bool {
	if o != nil && !IsNil(o.VisitUrlForPc) {
		return true
	}

	return false
}

// SetVisitUrlForPc gets a reference to the given string and assigns it to the VisitUrlForPc field.
func (o *AlipayOpenMiniInnerappPluginservicePublishModel) SetVisitUrlForPc(v string) {
	o.VisitUrlForPc = &v
}

func (o AlipayOpenMiniInnerappPluginservicePublishModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerappPluginservicePublishModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbilityTypeList) {
		toSerialize["ability_type_list"] = o.AbilityTypeList
	}
	if !IsNil(o.AppLogo) {
		toSerialize["app_logo"] = o.AppLogo
	}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DetailForClient) {
		toSerialize["detail_for_client"] = o.DetailForClient
	}
	if !IsNil(o.DetailForPc) {
		toSerialize["detail_for_pc"] = o.DetailForPc
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.SellCrowd) {
		toSerialize["sell_crowd"] = o.SellCrowd
	}
	if !IsNil(o.ShowType) {
		toSerialize["show_type"] = o.ShowType
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.VisitUrlForPc) {
		toSerialize["visit_url_for_pc"] = o.VisitUrlForPc
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerappPluginservicePublishModel struct {
	value *AlipayOpenMiniInnerappPluginservicePublishModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappPluginservicePublishModel) Get() *AlipayOpenMiniInnerappPluginservicePublishModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappPluginservicePublishModel) Set(val *AlipayOpenMiniInnerappPluginservicePublishModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappPluginservicePublishModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappPluginservicePublishModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappPluginservicePublishModel(val *AlipayOpenMiniInnerappPluginservicePublishModel) *NullableAlipayOpenMiniInnerappPluginservicePublishModel {
	return &NullableAlipayOpenMiniInnerappPluginservicePublishModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappPluginservicePublishModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappPluginservicePublishModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
