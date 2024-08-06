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

// checks if the AlipayOpenAgentMobilepaySignModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAgentMobilepaySignModel{}

// AlipayOpenAgentMobilepaySignModel struct for AlipayOpenAgentMobilepaySignModel
type AlipayOpenAgentMobilepaySignModel struct {
	// 应用在哪些市场上架，枚举值为：苹果,应用宝,华为,360,小米,豌豆荚,其他
	AppMarket []string `json:"app_market,omitempty"`
	// 商户的APP应用名称
	AppName *string `json:"app_name,omitempty"`
	// 应用上架状态，枚举值为 已上线，未上线
	AppStatus *string `json:"app_status,omitempty"`
	// 可以登录该应用的测试账户
	AppTestAccount *string `json:"app_test_account,omitempty"`
	// 可以登录此应用的账户的密码。对应app_test_account的登录密码
	AppTestAccountPassword *string `json:"app_test_account_password,omitempty"`
	// 应用类型，枚举值为：IOS，安卓
	AppType []string `json:"app_type,omitempty"`
	// 代商户操作事务编号，通过alipay.open.isv.agent.create接口进行创建。
	BatchNo *string `json:"batch_no,omitempty"`
	// 营业执照法人手机号码
	BusinessLicenseMobile *string `json:"business_license_mobile,omitempty"`
	// 营业执照号码。
	BusinessLicenseNo *string `json:"business_license_no,omitempty"`
	// 营业期限
	DateLimitation *string `json:"date_limitation,omitempty"`
	// 应用下载链接
	DownloadLink *string `json:"download_link,omitempty"`
	// 营业期限是否长期有效
	LongTerm *bool `json:"long_term,omitempty"`
	// 所属MCCCode，详情可参考 <a href=\"https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE\">商家经营类目</a> 中的“经营类目编码”
	MccCode *string `json:"mcc_code,omitempty"`
	// 传参：APP，代表设备类型是APP
	MobileType *string `json:"mobile_type,omitempty"`
	// 传参：mobile，代表交易场景是移动设备
	TradeScene *string `json:"trade_scene,omitempty"`
}

// NewAlipayOpenAgentMobilepaySignModel instantiates a new AlipayOpenAgentMobilepaySignModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAgentMobilepaySignModel() *AlipayOpenAgentMobilepaySignModel {
	this := AlipayOpenAgentMobilepaySignModel{}
	return &this
}

// NewAlipayOpenAgentMobilepaySignModelWithDefaults instantiates a new AlipayOpenAgentMobilepaySignModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAgentMobilepaySignModelWithDefaults() *AlipayOpenAgentMobilepaySignModel {
	this := AlipayOpenAgentMobilepaySignModel{}
	return &this
}

// GetAppMarket returns the AppMarket field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppMarket() []string {
	if o == nil || IsNil(o.AppMarket) {
		var ret []string
		return ret
	}
	return o.AppMarket
}

// GetAppMarketOk returns a tuple with the AppMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppMarketOk() ([]string, bool) {
	if o == nil || IsNil(o.AppMarket) {
		return nil, false
	}
	return o.AppMarket, true
}

// HasAppMarket returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasAppMarket() bool {
	if o != nil && !IsNil(o.AppMarket) {
		return true
	}

	return false
}

// SetAppMarket gets a reference to the given []string and assigns it to the AppMarket field.
func (o *AlipayOpenAgentMobilepaySignModel) SetAppMarket(v []string) {
	o.AppMarket = v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AlipayOpenAgentMobilepaySignModel) SetAppName(v string) {
	o.AppName = &v
}

// GetAppStatus returns the AppStatus field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppStatus() string {
	if o == nil || IsNil(o.AppStatus) {
		var ret string
		return ret
	}
	return *o.AppStatus
}

// GetAppStatusOk returns a tuple with the AppStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AppStatus) {
		return nil, false
	}
	return o.AppStatus, true
}

// HasAppStatus returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasAppStatus() bool {
	if o != nil && !IsNil(o.AppStatus) {
		return true
	}

	return false
}

// SetAppStatus gets a reference to the given string and assigns it to the AppStatus field.
func (o *AlipayOpenAgentMobilepaySignModel) SetAppStatus(v string) {
	o.AppStatus = &v
}

// GetAppTestAccount returns the AppTestAccount field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppTestAccount() string {
	if o == nil || IsNil(o.AppTestAccount) {
		var ret string
		return ret
	}
	return *o.AppTestAccount
}

// GetAppTestAccountOk returns a tuple with the AppTestAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppTestAccountOk() (*string, bool) {
	if o == nil || IsNil(o.AppTestAccount) {
		return nil, false
	}
	return o.AppTestAccount, true
}

// HasAppTestAccount returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasAppTestAccount() bool {
	if o != nil && !IsNil(o.AppTestAccount) {
		return true
	}

	return false
}

// SetAppTestAccount gets a reference to the given string and assigns it to the AppTestAccount field.
func (o *AlipayOpenAgentMobilepaySignModel) SetAppTestAccount(v string) {
	o.AppTestAccount = &v
}

// GetAppTestAccountPassword returns the AppTestAccountPassword field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppTestAccountPassword() string {
	if o == nil || IsNil(o.AppTestAccountPassword) {
		var ret string
		return ret
	}
	return *o.AppTestAccountPassword
}

// GetAppTestAccountPasswordOk returns a tuple with the AppTestAccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppTestAccountPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AppTestAccountPassword) {
		return nil, false
	}
	return o.AppTestAccountPassword, true
}

// HasAppTestAccountPassword returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasAppTestAccountPassword() bool {
	if o != nil && !IsNil(o.AppTestAccountPassword) {
		return true
	}

	return false
}

// SetAppTestAccountPassword gets a reference to the given string and assigns it to the AppTestAccountPassword field.
func (o *AlipayOpenAgentMobilepaySignModel) SetAppTestAccountPassword(v string) {
	o.AppTestAccountPassword = &v
}

// GetAppType returns the AppType field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppType() []string {
	if o == nil || IsNil(o.AppType) {
		var ret []string
		return ret
	}
	return o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetAppTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.AppType) {
		return nil, false
	}
	return o.AppType, true
}

// HasAppType returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasAppType() bool {
	if o != nil && !IsNil(o.AppType) {
		return true
	}

	return false
}

// SetAppType gets a reference to the given []string and assigns it to the AppType field.
func (o *AlipayOpenAgentMobilepaySignModel) SetAppType(v []string) {
	o.AppType = v
}

// GetBatchNo returns the BatchNo field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetBatchNo() string {
	if o == nil || IsNil(o.BatchNo) {
		var ret string
		return ret
	}
	return *o.BatchNo
}

// GetBatchNoOk returns a tuple with the BatchNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetBatchNoOk() (*string, bool) {
	if o == nil || IsNil(o.BatchNo) {
		return nil, false
	}
	return o.BatchNo, true
}

// HasBatchNo returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasBatchNo() bool {
	if o != nil && !IsNil(o.BatchNo) {
		return true
	}

	return false
}

// SetBatchNo gets a reference to the given string and assigns it to the BatchNo field.
func (o *AlipayOpenAgentMobilepaySignModel) SetBatchNo(v string) {
	o.BatchNo = &v
}

// GetBusinessLicenseMobile returns the BusinessLicenseMobile field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetBusinessLicenseMobile() string {
	if o == nil || IsNil(o.BusinessLicenseMobile) {
		var ret string
		return ret
	}
	return *o.BusinessLicenseMobile
}

// GetBusinessLicenseMobileOk returns a tuple with the BusinessLicenseMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetBusinessLicenseMobileOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLicenseMobile) {
		return nil, false
	}
	return o.BusinessLicenseMobile, true
}

// HasBusinessLicenseMobile returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasBusinessLicenseMobile() bool {
	if o != nil && !IsNil(o.BusinessLicenseMobile) {
		return true
	}

	return false
}

// SetBusinessLicenseMobile gets a reference to the given string and assigns it to the BusinessLicenseMobile field.
func (o *AlipayOpenAgentMobilepaySignModel) SetBusinessLicenseMobile(v string) {
	o.BusinessLicenseMobile = &v
}

// GetBusinessLicenseNo returns the BusinessLicenseNo field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetBusinessLicenseNo() string {
	if o == nil || IsNil(o.BusinessLicenseNo) {
		var ret string
		return ret
	}
	return *o.BusinessLicenseNo
}

// GetBusinessLicenseNoOk returns a tuple with the BusinessLicenseNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetBusinessLicenseNoOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLicenseNo) {
		return nil, false
	}
	return o.BusinessLicenseNo, true
}

// HasBusinessLicenseNo returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasBusinessLicenseNo() bool {
	if o != nil && !IsNil(o.BusinessLicenseNo) {
		return true
	}

	return false
}

// SetBusinessLicenseNo gets a reference to the given string and assigns it to the BusinessLicenseNo field.
func (o *AlipayOpenAgentMobilepaySignModel) SetBusinessLicenseNo(v string) {
	o.BusinessLicenseNo = &v
}

// GetDateLimitation returns the DateLimitation field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetDateLimitation() string {
	if o == nil || IsNil(o.DateLimitation) {
		var ret string
		return ret
	}
	return *o.DateLimitation
}

// GetDateLimitationOk returns a tuple with the DateLimitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetDateLimitationOk() (*string, bool) {
	if o == nil || IsNil(o.DateLimitation) {
		return nil, false
	}
	return o.DateLimitation, true
}

// HasDateLimitation returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasDateLimitation() bool {
	if o != nil && !IsNil(o.DateLimitation) {
		return true
	}

	return false
}

// SetDateLimitation gets a reference to the given string and assigns it to the DateLimitation field.
func (o *AlipayOpenAgentMobilepaySignModel) SetDateLimitation(v string) {
	o.DateLimitation = &v
}

// GetDownloadLink returns the DownloadLink field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetDownloadLink() string {
	if o == nil || IsNil(o.DownloadLink) {
		var ret string
		return ret
	}
	return *o.DownloadLink
}

// GetDownloadLinkOk returns a tuple with the DownloadLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetDownloadLinkOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadLink) {
		return nil, false
	}
	return o.DownloadLink, true
}

// HasDownloadLink returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasDownloadLink() bool {
	if o != nil && !IsNil(o.DownloadLink) {
		return true
	}

	return false
}

// SetDownloadLink gets a reference to the given string and assigns it to the DownloadLink field.
func (o *AlipayOpenAgentMobilepaySignModel) SetDownloadLink(v string) {
	o.DownloadLink = &v
}

// GetLongTerm returns the LongTerm field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetLongTerm() bool {
	if o == nil || IsNil(o.LongTerm) {
		var ret bool
		return ret
	}
	return *o.LongTerm
}

// GetLongTermOk returns a tuple with the LongTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetLongTermOk() (*bool, bool) {
	if o == nil || IsNil(o.LongTerm) {
		return nil, false
	}
	return o.LongTerm, true
}

// HasLongTerm returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasLongTerm() bool {
	if o != nil && !IsNil(o.LongTerm) {
		return true
	}

	return false
}

// SetLongTerm gets a reference to the given bool and assigns it to the LongTerm field.
func (o *AlipayOpenAgentMobilepaySignModel) SetLongTerm(v bool) {
	o.LongTerm = &v
}

// GetMccCode returns the MccCode field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetMccCode() string {
	if o == nil || IsNil(o.MccCode) {
		var ret string
		return ret
	}
	return *o.MccCode
}

// GetMccCodeOk returns a tuple with the MccCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetMccCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MccCode) {
		return nil, false
	}
	return o.MccCode, true
}

// HasMccCode returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasMccCode() bool {
	if o != nil && !IsNil(o.MccCode) {
		return true
	}

	return false
}

// SetMccCode gets a reference to the given string and assigns it to the MccCode field.
func (o *AlipayOpenAgentMobilepaySignModel) SetMccCode(v string) {
	o.MccCode = &v
}

// GetMobileType returns the MobileType field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetMobileType() string {
	if o == nil || IsNil(o.MobileType) {
		var ret string
		return ret
	}
	return *o.MobileType
}

// GetMobileTypeOk returns a tuple with the MobileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetMobileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MobileType) {
		return nil, false
	}
	return o.MobileType, true
}

// HasMobileType returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasMobileType() bool {
	if o != nil && !IsNil(o.MobileType) {
		return true
	}

	return false
}

// SetMobileType gets a reference to the given string and assigns it to the MobileType field.
func (o *AlipayOpenAgentMobilepaySignModel) SetMobileType(v string) {
	o.MobileType = &v
}

// GetTradeScene returns the TradeScene field value if set, zero value otherwise.
func (o *AlipayOpenAgentMobilepaySignModel) GetTradeScene() string {
	if o == nil || IsNil(o.TradeScene) {
		var ret string
		return ret
	}
	return *o.TradeScene
}

// GetTradeSceneOk returns a tuple with the TradeScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentMobilepaySignModel) GetTradeSceneOk() (*string, bool) {
	if o == nil || IsNil(o.TradeScene) {
		return nil, false
	}
	return o.TradeScene, true
}

// HasTradeScene returns a boolean if a field has been set.
func (o *AlipayOpenAgentMobilepaySignModel) HasTradeScene() bool {
	if o != nil && !IsNil(o.TradeScene) {
		return true
	}

	return false
}

// SetTradeScene gets a reference to the given string and assigns it to the TradeScene field.
func (o *AlipayOpenAgentMobilepaySignModel) SetTradeScene(v string) {
	o.TradeScene = &v
}

func (o AlipayOpenAgentMobilepaySignModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAgentMobilepaySignModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppMarket) {
		toSerialize["app_market"] = o.AppMarket
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AppStatus) {
		toSerialize["app_status"] = o.AppStatus
	}
	if !IsNil(o.AppTestAccount) {
		toSerialize["app_test_account"] = o.AppTestAccount
	}
	if !IsNil(o.AppTestAccountPassword) {
		toSerialize["app_test_account_password"] = o.AppTestAccountPassword
	}
	if !IsNil(o.AppType) {
		toSerialize["app_type"] = o.AppType
	}
	if !IsNil(o.BatchNo) {
		toSerialize["batch_no"] = o.BatchNo
	}
	if !IsNil(o.BusinessLicenseMobile) {
		toSerialize["business_license_mobile"] = o.BusinessLicenseMobile
	}
	if !IsNil(o.BusinessLicenseNo) {
		toSerialize["business_license_no"] = o.BusinessLicenseNo
	}
	if !IsNil(o.DateLimitation) {
		toSerialize["date_limitation"] = o.DateLimitation
	}
	if !IsNil(o.DownloadLink) {
		toSerialize["download_link"] = o.DownloadLink
	}
	if !IsNil(o.LongTerm) {
		toSerialize["long_term"] = o.LongTerm
	}
	if !IsNil(o.MccCode) {
		toSerialize["mcc_code"] = o.MccCode
	}
	if !IsNil(o.MobileType) {
		toSerialize["mobile_type"] = o.MobileType
	}
	if !IsNil(o.TradeScene) {
		toSerialize["trade_scene"] = o.TradeScene
	}
	return toSerialize, nil
}

type NullableAlipayOpenAgentMobilepaySignModel struct {
	value *AlipayOpenAgentMobilepaySignModel
	isSet bool
}

func (v NullableAlipayOpenAgentMobilepaySignModel) Get() *AlipayOpenAgentMobilepaySignModel {
	return v.value
}

func (v *NullableAlipayOpenAgentMobilepaySignModel) Set(val *AlipayOpenAgentMobilepaySignModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentMobilepaySignModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentMobilepaySignModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentMobilepaySignModel(val *AlipayOpenAgentMobilepaySignModel) *NullableAlipayOpenAgentMobilepaySignModel {
	return &NullableAlipayOpenAgentMobilepaySignModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentMobilepaySignModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentMobilepaySignModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

