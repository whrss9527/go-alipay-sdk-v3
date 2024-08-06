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

// checks if the AlipayCommerceEducateSchoolcardOrderSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEducateSchoolcardOrderSyncModel{}

// AlipayCommerceEducateSchoolcardOrderSyncModel struct for AlipayCommerceEducateSchoolcardOrderSyncModel
type AlipayCommerceEducateSchoolcardOrderSyncModel struct {
	// 实际金额（总支付金额），单位为【元】
	ActualAmount *string `json:"actual_amount,omitempty"`
	// 小程序appid
	AppletAppId *string `json:"applet_app_id,omitempty"`
	// 校园卡余额
	CardBalance *string `json:"card_balance,omitempty"`
	// 128
	CardNo *string `json:"card_no,omitempty"`
	// 校园卡类型
	CardType *string `json:"card_type,omitempty"`
	// 该笔订单真实的创建时间，需精确到毫秒。
	CreateTime *string `json:"create_time,omitempty"`
	// 优惠金额
	DiscountAmount *string     `json:"discount_amount,omitempty"`
	GoodsOrders    *GoodsOrder `json:"goods_orders,omitempty"`
	// 商家名称，不传默认展示学校名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// 订单修改时间
	ModifiedTime *string `json:"modified_time,omitempty"`
	// 用户open_id
	OpenId *string `json:"open_id,omitempty"`
	// 订单金额
	OrderAmount *string `json:"order_amount,omitempty"`
	// 订单详情链接
	OrderDetailUrl *string `json:"order_detail_url,omitempty"`
	// 订单状态
	OrderStatus *string `json:"order_status,omitempty"`
	// 外部业务号，由商家自定义，128个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 支付地点
	PayAddress *string `json:"pay_address,omitempty"`
	// 付款方式，不传默认展示学校名称+校园卡+（卡号后四位）
	PayMode *string `json:"pay_mode,omitempty"`
	// 系统商编号。该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	RakeBackPid *string `json:"rake_back_pid,omitempty"`
	// 学校内标，录入学校接口返回的参数
	SchoolId *string `json:"school_id,omitempty"`
	// 学校收款账号
	SchoolPid *string `json:"school_pid,omitempty"`
	// 服务提供者名称
	ServiceProviderName *string `json:"service_provider_name,omitempty"`
	// 业务类型
	Type *string `json:"type,omitempty"`
	// 买家支付宝用户ID。 2088开头的16位纯数字，小程序场景下获取用户ID请参考：用户授权; 其它场景下获取用户ID请参考：网页授权获取用户信息;
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayCommerceEducateSchoolcardOrderSyncModel instantiates a new AlipayCommerceEducateSchoolcardOrderSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEducateSchoolcardOrderSyncModel() *AlipayCommerceEducateSchoolcardOrderSyncModel {
	this := AlipayCommerceEducateSchoolcardOrderSyncModel{}
	return &this
}

// NewAlipayCommerceEducateSchoolcardOrderSyncModelWithDefaults instantiates a new AlipayCommerceEducateSchoolcardOrderSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEducateSchoolcardOrderSyncModelWithDefaults() *AlipayCommerceEducateSchoolcardOrderSyncModel {
	this := AlipayCommerceEducateSchoolcardOrderSyncModel{}
	return &this
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetActualAmount() string {
	if o == nil || IsNil(o.ActualAmount) {
		var ret string
		return ret
	}
	return *o.ActualAmount
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetActualAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ActualAmount) {
		return nil, false
	}
	return o.ActualAmount, true
}

// HasActualAmount returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasActualAmount() bool {
	if o != nil && !IsNil(o.ActualAmount) {
		return true
	}

	return false
}

// SetActualAmount gets a reference to the given string and assigns it to the ActualAmount field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetActualAmount(v string) {
	o.ActualAmount = &v
}

// GetAppletAppId returns the AppletAppId field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetAppletAppId() string {
	if o == nil || IsNil(o.AppletAppId) {
		var ret string
		return ret
	}
	return *o.AppletAppId
}

// GetAppletAppIdOk returns a tuple with the AppletAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetAppletAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppletAppId) {
		return nil, false
	}
	return o.AppletAppId, true
}

// HasAppletAppId returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasAppletAppId() bool {
	if o != nil && !IsNil(o.AppletAppId) {
		return true
	}

	return false
}

// SetAppletAppId gets a reference to the given string and assigns it to the AppletAppId field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetAppletAppId(v string) {
	o.AppletAppId = &v
}

// GetCardBalance returns the CardBalance field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCardBalance() string {
	if o == nil || IsNil(o.CardBalance) {
		var ret string
		return ret
	}
	return *o.CardBalance
}

// GetCardBalanceOk returns a tuple with the CardBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCardBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.CardBalance) {
		return nil, false
	}
	return o.CardBalance, true
}

// HasCardBalance returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasCardBalance() bool {
	if o != nil && !IsNil(o.CardBalance) {
		return true
	}

	return false
}

// SetCardBalance gets a reference to the given string and assigns it to the CardBalance field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetCardBalance(v string) {
	o.CardBalance = &v
}

// GetCardNo returns the CardNo field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCardNo() string {
	if o == nil || IsNil(o.CardNo) {
		var ret string
		return ret
	}
	return *o.CardNo
}

// GetCardNoOk returns a tuple with the CardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.CardNo) {
		return nil, false
	}
	return o.CardNo, true
}

// HasCardNo returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasCardNo() bool {
	if o != nil && !IsNil(o.CardNo) {
		return true
	}

	return false
}

// SetCardNo gets a reference to the given string and assigns it to the CardNo field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetCardNo(v string) {
	o.CardNo = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetCardType(v string) {
	o.CardType = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetDiscountAmount() string {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret string
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetDiscountAmountOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given string and assigns it to the DiscountAmount field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetDiscountAmount(v string) {
	o.DiscountAmount = &v
}

// GetGoodsOrders returns the GoodsOrders field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetGoodsOrders() GoodsOrder {
	if o == nil || IsNil(o.GoodsOrders) {
		var ret GoodsOrder
		return ret
	}
	return *o.GoodsOrders
}

// GetGoodsOrdersOk returns a tuple with the GoodsOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetGoodsOrdersOk() (*GoodsOrder, bool) {
	if o == nil || IsNil(o.GoodsOrders) {
		return nil, false
	}
	return o.GoodsOrders, true
}

// HasGoodsOrders returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasGoodsOrders() bool {
	if o != nil && !IsNil(o.GoodsOrders) {
		return true
	}

	return false
}

// SetGoodsOrders gets a reference to the given GoodsOrder and assigns it to the GoodsOrders field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetGoodsOrders(v GoodsOrder) {
	o.GoodsOrders = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetModifiedTime() string {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret string
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetModifiedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given string and assigns it to the ModifiedTime field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetModifiedTime(v string) {
	o.ModifiedTime = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOrderAmount returns the OrderAmount field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOrderAmount() string {
	if o == nil || IsNil(o.OrderAmount) {
		var ret string
		return ret
	}
	return *o.OrderAmount
}

// GetOrderAmountOk returns a tuple with the OrderAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOrderAmountOk() (*string, bool) {
	if o == nil || IsNil(o.OrderAmount) {
		return nil, false
	}
	return o.OrderAmount, true
}

// HasOrderAmount returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasOrderAmount() bool {
	if o != nil && !IsNil(o.OrderAmount) {
		return true
	}

	return false
}

// SetOrderAmount gets a reference to the given string and assigns it to the OrderAmount field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetOrderAmount(v string) {
	o.OrderAmount = &v
}

// GetOrderDetailUrl returns the OrderDetailUrl field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOrderDetailUrl() string {
	if o == nil || IsNil(o.OrderDetailUrl) {
		var ret string
		return ret
	}
	return *o.OrderDetailUrl
}

// GetOrderDetailUrlOk returns a tuple with the OrderDetailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOrderDetailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OrderDetailUrl) {
		return nil, false
	}
	return o.OrderDetailUrl, true
}

// HasOrderDetailUrl returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasOrderDetailUrl() bool {
	if o != nil && !IsNil(o.OrderDetailUrl) {
		return true
	}

	return false
}

// SetOrderDetailUrl gets a reference to the given string and assigns it to the OrderDetailUrl field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetOrderDetailUrl(v string) {
	o.OrderDetailUrl = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPayAddress returns the PayAddress field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetPayAddress() string {
	if o == nil || IsNil(o.PayAddress) {
		var ret string
		return ret
	}
	return *o.PayAddress
}

// GetPayAddressOk returns a tuple with the PayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetPayAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PayAddress) {
		return nil, false
	}
	return o.PayAddress, true
}

// HasPayAddress returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasPayAddress() bool {
	if o != nil && !IsNil(o.PayAddress) {
		return true
	}

	return false
}

// SetPayAddress gets a reference to the given string and assigns it to the PayAddress field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetPayAddress(v string) {
	o.PayAddress = &v
}

// GetPayMode returns the PayMode field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetPayMode() string {
	if o == nil || IsNil(o.PayMode) {
		var ret string
		return ret
	}
	return *o.PayMode
}

// GetPayModeOk returns a tuple with the PayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetPayModeOk() (*string, bool) {
	if o == nil || IsNil(o.PayMode) {
		return nil, false
	}
	return o.PayMode, true
}

// HasPayMode returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasPayMode() bool {
	if o != nil && !IsNil(o.PayMode) {
		return true
	}

	return false
}

// SetPayMode gets a reference to the given string and assigns it to the PayMode field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetPayMode(v string) {
	o.PayMode = &v
}

// GetRakeBackPid returns the RakeBackPid field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetRakeBackPid() string {
	if o == nil || IsNil(o.RakeBackPid) {
		var ret string
		return ret
	}
	return *o.RakeBackPid
}

// GetRakeBackPidOk returns a tuple with the RakeBackPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetRakeBackPidOk() (*string, bool) {
	if o == nil || IsNil(o.RakeBackPid) {
		return nil, false
	}
	return o.RakeBackPid, true
}

// HasRakeBackPid returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasRakeBackPid() bool {
	if o != nil && !IsNil(o.RakeBackPid) {
		return true
	}

	return false
}

// SetRakeBackPid gets a reference to the given string and assigns it to the RakeBackPid field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetRakeBackPid(v string) {
	o.RakeBackPid = &v
}

// GetSchoolId returns the SchoolId field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetSchoolId() string {
	if o == nil || IsNil(o.SchoolId) {
		var ret string
		return ret
	}
	return *o.SchoolId
}

// GetSchoolIdOk returns a tuple with the SchoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetSchoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.SchoolId) {
		return nil, false
	}
	return o.SchoolId, true
}

// HasSchoolId returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasSchoolId() bool {
	if o != nil && !IsNil(o.SchoolId) {
		return true
	}

	return false
}

// SetSchoolId gets a reference to the given string and assigns it to the SchoolId field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetSchoolId(v string) {
	o.SchoolId = &v
}

// GetSchoolPid returns the SchoolPid field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetSchoolPid() string {
	if o == nil || IsNil(o.SchoolPid) {
		var ret string
		return ret
	}
	return *o.SchoolPid
}

// GetSchoolPidOk returns a tuple with the SchoolPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetSchoolPidOk() (*string, bool) {
	if o == nil || IsNil(o.SchoolPid) {
		return nil, false
	}
	return o.SchoolPid, true
}

// HasSchoolPid returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasSchoolPid() bool {
	if o != nil && !IsNil(o.SchoolPid) {
		return true
	}

	return false
}

// SetSchoolPid gets a reference to the given string and assigns it to the SchoolPid field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetSchoolPid(v string) {
	o.SchoolPid = &v
}

// GetServiceProviderName returns the ServiceProviderName field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetServiceProviderName() string {
	if o == nil || IsNil(o.ServiceProviderName) {
		var ret string
		return ret
	}
	return *o.ServiceProviderName
}

// GetServiceProviderNameOk returns a tuple with the ServiceProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetServiceProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceProviderName) {
		return nil, false
	}
	return o.ServiceProviderName, true
}

// HasServiceProviderName returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasServiceProviderName() bool {
	if o != nil && !IsNil(o.ServiceProviderName) {
		return true
	}

	return false
}

// SetServiceProviderName gets a reference to the given string and assigns it to the ServiceProviderName field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetServiceProviderName(v string) {
	o.ServiceProviderName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetType(v string) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayCommerceEducateSchoolcardOrderSyncModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayCommerceEducateSchoolcardOrderSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEducateSchoolcardOrderSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActualAmount) {
		toSerialize["actual_amount"] = o.ActualAmount
	}
	if !IsNil(o.AppletAppId) {
		toSerialize["applet_app_id"] = o.AppletAppId
	}
	if !IsNil(o.CardBalance) {
		toSerialize["card_balance"] = o.CardBalance
	}
	if !IsNil(o.CardNo) {
		toSerialize["card_no"] = o.CardNo
	}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.GoodsOrders) {
		toSerialize["goods_orders"] = o.GoodsOrders
	}
	if !IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OrderAmount) {
		toSerialize["order_amount"] = o.OrderAmount
	}
	if !IsNil(o.OrderDetailUrl) {
		toSerialize["order_detail_url"] = o.OrderDetailUrl
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PayAddress) {
		toSerialize["pay_address"] = o.PayAddress
	}
	if !IsNil(o.PayMode) {
		toSerialize["pay_mode"] = o.PayMode
	}
	if !IsNil(o.RakeBackPid) {
		toSerialize["rake_back_pid"] = o.RakeBackPid
	}
	if !IsNil(o.SchoolId) {
		toSerialize["school_id"] = o.SchoolId
	}
	if !IsNil(o.SchoolPid) {
		toSerialize["school_pid"] = o.SchoolPid
	}
	if !IsNil(o.ServiceProviderName) {
		toSerialize["service_provider_name"] = o.ServiceProviderName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEducateSchoolcardOrderSyncModel struct {
	value *AlipayCommerceEducateSchoolcardOrderSyncModel
	isSet bool
}

func (v NullableAlipayCommerceEducateSchoolcardOrderSyncModel) Get() *AlipayCommerceEducateSchoolcardOrderSyncModel {
	return v.value
}

func (v *NullableAlipayCommerceEducateSchoolcardOrderSyncModel) Set(val *AlipayCommerceEducateSchoolcardOrderSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEducateSchoolcardOrderSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEducateSchoolcardOrderSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEducateSchoolcardOrderSyncModel(val *AlipayCommerceEducateSchoolcardOrderSyncModel) *NullableAlipayCommerceEducateSchoolcardOrderSyncModel {
	return &NullableAlipayCommerceEducateSchoolcardOrderSyncModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEducateSchoolcardOrderSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEducateSchoolcardOrderSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
