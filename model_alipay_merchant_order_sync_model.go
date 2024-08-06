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

// checks if the AlipayMerchantOrderSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantOrderSyncModel{}

// AlipayMerchantOrderSyncModel struct for AlipayMerchantOrderSyncModel
type AlipayMerchantOrderSyncModel struct {
	// 订单总金额：某笔交易订单优惠前的总金额，单位为【元】</br> <a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*UuuWRpmekegAAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例一</a> <br><a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*kBkOTZpqP40AAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例二</a> <br><a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*ZpkjTJQlFVAAAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例三</a>
	Amount *string `json:"amount,omitempty"`
	// 买家userId
	BuyerId *string `json:"buyer_id,omitempty"`
	BuyerInfo *UserInfomation `json:"buyer_info,omitempty"`
	// 用户open_id
	BuyerOpenId *string `json:"buyer_open_id,omitempty"`
	// 标准服务类目
	CategoryId *string `json:"category_id,omitempty"`
	// 商户总计优惠金额：代表商户侧给予用户的总计优惠金额 （不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】。</br> <a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*UuuWRpmekegAAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例一</a> <br><a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*kBkOTZpqP40AAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例二</a> <br><a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*ZpkjTJQlFVAAAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例三</a>
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// 订单优惠信息
	DiscountInfoList []DiscountInfoData `json:"discount_info_list,omitempty"`
	// 扩展信息，请参见 <a href=\"https://opendocs.alipay.com/mini/04zsxt?pathHash=654d4f10\">小程序订单中心模板</a>
	ExtInfo []OrderExtInfo `json:"ext_info,omitempty"`
	// 商品信息列表
	ItemOrderList []ItemOrderInfo `json:"item_order_list,omitempty"`
	// 行程信息
	JourneyOrderList []OrderJourneyInfo `json:"journey_order_list,omitempty"`
	// 物流信息列表
	LogisticsInfoList []OrderLogisticsInformationRequest `json:"logistics_info_list,omitempty"`
	// 本字段已废弃，无需填写！(本字段已废弃 不再使用)
	// Deprecated
	OrderAuthCode *string `json:"order_auth_code,omitempty"`
	// 订单创建时间
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	// 订单修改时间
	OrderModifiedTime *string `json:"order_modified_time,omitempty"`
	// 订单支付时间
	OrderPayTime *string `json:"order_pay_time,omitempty"`
	// 订单类型
	OrderType *string `json:"order_type,omitempty"`
	// 外部订单号 out_biz_no唯一对应一笔订单，相同的订单需传入相同的out_biz_no
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 交易对应的签约商户userId(注意：该字段自2020-02-16日起，可以不传入)
	// Deprecated
	PartnerId *string `json:"partner_id,omitempty"`
	// 用户应付金额 ：用户最终结算时需要支付金额（不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】</br> <a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*UuuWRpmekegAAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例一</a> <br><a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*kBkOTZpqP40AAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例二</a> <br><a href=\"https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*ZpkjTJQlFVAAAAAAAAAAAAAAAQAAAQ/original\" target=\"_blank\">实际案例三</a>
	PayAmount *string `json:"pay_amount,omitempty"`
	// 支付超时时间，超过时间支付宝自行关闭订单
	PayTimeoutExpress *string `json:"pay_timeout_express,omitempty"`
	// 商户订单同步记录id(仅部分存量接入和行业模板需要外，其他情况可以不传入)
	// Deprecated
	RecordId *string `json:"record_id,omitempty"`
	// 卖家userId(注意：该字段自2020-02-16日起，可以不传入)
	// Deprecated
	SellerId *string `json:"seller_id,omitempty"`
	// 是否需要小程序订单代理发送模版消息。不传默认不发送
	SendMsg *string `json:"send_msg,omitempty"`
	// 服务code：传入小程序后台提报的服务id，将订单与服务关联，有利于提高服务曝光机会；入参服务id的类目须与订单类型相符，若不相符将会报错；如订单类型为“外卖”，则入参的服务ID所对应的服务类目也必须得是”外卖“；service_code 通过 <a href=\"https://opendocs.alipay.com/mini/e1c835a1_alipay.open.app.service.apply\" target=\"_blank\">alipay.open.app.service.apply</a> ，(服务提报申请)接口创建服务后获取。
	ServiceCode *string `json:"service_code,omitempty"`
	ShopInfo *OrderShopInfo `json:"shop_info,omitempty"`
	// 用于区分用户下单的订单来源
	SourceApp *string `json:"source_app,omitempty"`
	// 同步内容
	SyncContent *string `json:"sync_content,omitempty"`
	TicketInfo *TicketInfo `json:"ticket_info,omitempty"`
	// 凭证信息
	TicketOrderList []TicketOrderInfo `json:"ticket_order_list,omitempty"`
	// 订单所对应的支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
	// 交易号类型
	TradeType *string `json:"trade_type,omitempty"`
}

// NewAlipayMerchantOrderSyncModel instantiates a new AlipayMerchantOrderSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantOrderSyncModel() *AlipayMerchantOrderSyncModel {
	this := AlipayMerchantOrderSyncModel{}
	return &this
}

// NewAlipayMerchantOrderSyncModelWithDefaults instantiates a new AlipayMerchantOrderSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantOrderSyncModelWithDefaults() *AlipayMerchantOrderSyncModel {
	this := AlipayMerchantOrderSyncModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AlipayMerchantOrderSyncModel) SetAmount(v string) {
	o.Amount = &v
}

// GetBuyerId returns the BuyerId field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetBuyerId() string {
	if o == nil || IsNil(o.BuyerId) {
		var ret string
		return ret
	}
	return *o.BuyerId
}

// GetBuyerIdOk returns a tuple with the BuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetBuyerIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerId) {
		return nil, false
	}
	return o.BuyerId, true
}

// HasBuyerId returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasBuyerId() bool {
	if o != nil && !IsNil(o.BuyerId) {
		return true
	}

	return false
}

// SetBuyerId gets a reference to the given string and assigns it to the BuyerId field.
func (o *AlipayMerchantOrderSyncModel) SetBuyerId(v string) {
	o.BuyerId = &v
}

// GetBuyerInfo returns the BuyerInfo field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetBuyerInfo() UserInfomation {
	if o == nil || IsNil(o.BuyerInfo) {
		var ret UserInfomation
		return ret
	}
	return *o.BuyerInfo
}

// GetBuyerInfoOk returns a tuple with the BuyerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetBuyerInfoOk() (*UserInfomation, bool) {
	if o == nil || IsNil(o.BuyerInfo) {
		return nil, false
	}
	return o.BuyerInfo, true
}

// HasBuyerInfo returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasBuyerInfo() bool {
	if o != nil && !IsNil(o.BuyerInfo) {
		return true
	}

	return false
}

// SetBuyerInfo gets a reference to the given UserInfomation and assigns it to the BuyerInfo field.
func (o *AlipayMerchantOrderSyncModel) SetBuyerInfo(v UserInfomation) {
	o.BuyerInfo = &v
}

// GetBuyerOpenId returns the BuyerOpenId field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetBuyerOpenId() string {
	if o == nil || IsNil(o.BuyerOpenId) {
		var ret string
		return ret
	}
	return *o.BuyerOpenId
}

// GetBuyerOpenIdOk returns a tuple with the BuyerOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetBuyerOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerOpenId) {
		return nil, false
	}
	return o.BuyerOpenId, true
}

// HasBuyerOpenId returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasBuyerOpenId() bool {
	if o != nil && !IsNil(o.BuyerOpenId) {
		return true
	}

	return false
}

// SetBuyerOpenId gets a reference to the given string and assigns it to the BuyerOpenId field.
func (o *AlipayMerchantOrderSyncModel) SetBuyerOpenId(v string) {
	o.BuyerOpenId = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *AlipayMerchantOrderSyncModel) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetDiscountAmount() string {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret string
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetDiscountAmountOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given string and assigns it to the DiscountAmount field.
func (o *AlipayMerchantOrderSyncModel) SetDiscountAmount(v string) {
	o.DiscountAmount = &v
}

// GetDiscountInfoList returns the DiscountInfoList field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetDiscountInfoList() []DiscountInfoData {
	if o == nil || IsNil(o.DiscountInfoList) {
		var ret []DiscountInfoData
		return ret
	}
	return o.DiscountInfoList
}

// GetDiscountInfoListOk returns a tuple with the DiscountInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetDiscountInfoListOk() ([]DiscountInfoData, bool) {
	if o == nil || IsNil(o.DiscountInfoList) {
		return nil, false
	}
	return o.DiscountInfoList, true
}

// HasDiscountInfoList returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasDiscountInfoList() bool {
	if o != nil && !IsNil(o.DiscountInfoList) {
		return true
	}

	return false
}

// SetDiscountInfoList gets a reference to the given []DiscountInfoData and assigns it to the DiscountInfoList field.
func (o *AlipayMerchantOrderSyncModel) SetDiscountInfoList(v []DiscountInfoData) {
	o.DiscountInfoList = v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetExtInfo() []OrderExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret []OrderExtInfo
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetExtInfoOk() ([]OrderExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given []OrderExtInfo and assigns it to the ExtInfo field.
func (o *AlipayMerchantOrderSyncModel) SetExtInfo(v []OrderExtInfo) {
	o.ExtInfo = v
}

// GetItemOrderList returns the ItemOrderList field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetItemOrderList() []ItemOrderInfo {
	if o == nil || IsNil(o.ItemOrderList) {
		var ret []ItemOrderInfo
		return ret
	}
	return o.ItemOrderList
}

// GetItemOrderListOk returns a tuple with the ItemOrderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetItemOrderListOk() ([]ItemOrderInfo, bool) {
	if o == nil || IsNil(o.ItemOrderList) {
		return nil, false
	}
	return o.ItemOrderList, true
}

// HasItemOrderList returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasItemOrderList() bool {
	if o != nil && !IsNil(o.ItemOrderList) {
		return true
	}

	return false
}

// SetItemOrderList gets a reference to the given []ItemOrderInfo and assigns it to the ItemOrderList field.
func (o *AlipayMerchantOrderSyncModel) SetItemOrderList(v []ItemOrderInfo) {
	o.ItemOrderList = v
}

// GetJourneyOrderList returns the JourneyOrderList field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetJourneyOrderList() []OrderJourneyInfo {
	if o == nil || IsNil(o.JourneyOrderList) {
		var ret []OrderJourneyInfo
		return ret
	}
	return o.JourneyOrderList
}

// GetJourneyOrderListOk returns a tuple with the JourneyOrderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetJourneyOrderListOk() ([]OrderJourneyInfo, bool) {
	if o == nil || IsNil(o.JourneyOrderList) {
		return nil, false
	}
	return o.JourneyOrderList, true
}

// HasJourneyOrderList returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasJourneyOrderList() bool {
	if o != nil && !IsNil(o.JourneyOrderList) {
		return true
	}

	return false
}

// SetJourneyOrderList gets a reference to the given []OrderJourneyInfo and assigns it to the JourneyOrderList field.
func (o *AlipayMerchantOrderSyncModel) SetJourneyOrderList(v []OrderJourneyInfo) {
	o.JourneyOrderList = v
}

// GetLogisticsInfoList returns the LogisticsInfoList field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetLogisticsInfoList() []OrderLogisticsInformationRequest {
	if o == nil || IsNil(o.LogisticsInfoList) {
		var ret []OrderLogisticsInformationRequest
		return ret
	}
	return o.LogisticsInfoList
}

// GetLogisticsInfoListOk returns a tuple with the LogisticsInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetLogisticsInfoListOk() ([]OrderLogisticsInformationRequest, bool) {
	if o == nil || IsNil(o.LogisticsInfoList) {
		return nil, false
	}
	return o.LogisticsInfoList, true
}

// HasLogisticsInfoList returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasLogisticsInfoList() bool {
	if o != nil && !IsNil(o.LogisticsInfoList) {
		return true
	}

	return false
}

// SetLogisticsInfoList gets a reference to the given []OrderLogisticsInformationRequest and assigns it to the LogisticsInfoList field.
func (o *AlipayMerchantOrderSyncModel) SetLogisticsInfoList(v []OrderLogisticsInformationRequest) {
	o.LogisticsInfoList = v
}

// GetOrderAuthCode returns the OrderAuthCode field value if set, zero value otherwise.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetOrderAuthCode() string {
	if o == nil || IsNil(o.OrderAuthCode) {
		var ret string
		return ret
	}
	return *o.OrderAuthCode
}

// GetOrderAuthCodeOk returns a tuple with the OrderAuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetOrderAuthCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderAuthCode) {
		return nil, false
	}
	return o.OrderAuthCode, true
}

// HasOrderAuthCode returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasOrderAuthCode() bool {
	if o != nil && !IsNil(o.OrderAuthCode) {
		return true
	}

	return false
}

// SetOrderAuthCode gets a reference to the given string and assigns it to the OrderAuthCode field.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) SetOrderAuthCode(v string) {
	o.OrderAuthCode = &v
}

// GetOrderCreateTime returns the OrderCreateTime field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetOrderCreateTime() string {
	if o == nil || IsNil(o.OrderCreateTime) {
		var ret string
		return ret
	}
	return *o.OrderCreateTime
}

// GetOrderCreateTimeOk returns a tuple with the OrderCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetOrderCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderCreateTime) {
		return nil, false
	}
	return o.OrderCreateTime, true
}

// HasOrderCreateTime returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasOrderCreateTime() bool {
	if o != nil && !IsNil(o.OrderCreateTime) {
		return true
	}

	return false
}

// SetOrderCreateTime gets a reference to the given string and assigns it to the OrderCreateTime field.
func (o *AlipayMerchantOrderSyncModel) SetOrderCreateTime(v string) {
	o.OrderCreateTime = &v
}

// GetOrderModifiedTime returns the OrderModifiedTime field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetOrderModifiedTime() string {
	if o == nil || IsNil(o.OrderModifiedTime) {
		var ret string
		return ret
	}
	return *o.OrderModifiedTime
}

// GetOrderModifiedTimeOk returns a tuple with the OrderModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetOrderModifiedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderModifiedTime) {
		return nil, false
	}
	return o.OrderModifiedTime, true
}

// HasOrderModifiedTime returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasOrderModifiedTime() bool {
	if o != nil && !IsNil(o.OrderModifiedTime) {
		return true
	}

	return false
}

// SetOrderModifiedTime gets a reference to the given string and assigns it to the OrderModifiedTime field.
func (o *AlipayMerchantOrderSyncModel) SetOrderModifiedTime(v string) {
	o.OrderModifiedTime = &v
}

// GetOrderPayTime returns the OrderPayTime field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetOrderPayTime() string {
	if o == nil || IsNil(o.OrderPayTime) {
		var ret string
		return ret
	}
	return *o.OrderPayTime
}

// GetOrderPayTimeOk returns a tuple with the OrderPayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetOrderPayTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderPayTime) {
		return nil, false
	}
	return o.OrderPayTime, true
}

// HasOrderPayTime returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasOrderPayTime() bool {
	if o != nil && !IsNil(o.OrderPayTime) {
		return true
	}

	return false
}

// SetOrderPayTime gets a reference to the given string and assigns it to the OrderPayTime field.
func (o *AlipayMerchantOrderSyncModel) SetOrderPayTime(v string) {
	o.OrderPayTime = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *AlipayMerchantOrderSyncModel) SetOrderType(v string) {
	o.OrderType = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMerchantOrderSyncModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *AlipayMerchantOrderSyncModel) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetPayTimeoutExpress returns the PayTimeoutExpress field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetPayTimeoutExpress() string {
	if o == nil || IsNil(o.PayTimeoutExpress) {
		var ret string
		return ret
	}
	return *o.PayTimeoutExpress
}

// GetPayTimeoutExpressOk returns a tuple with the PayTimeoutExpress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetPayTimeoutExpressOk() (*string, bool) {
	if o == nil || IsNil(o.PayTimeoutExpress) {
		return nil, false
	}
	return o.PayTimeoutExpress, true
}

// HasPayTimeoutExpress returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasPayTimeoutExpress() bool {
	if o != nil && !IsNil(o.PayTimeoutExpress) {
		return true
	}

	return false
}

// SetPayTimeoutExpress gets a reference to the given string and assigns it to the PayTimeoutExpress field.
func (o *AlipayMerchantOrderSyncModel) SetPayTimeoutExpress(v string) {
	o.PayTimeoutExpress = &v
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetRecordId() string {
	if o == nil || IsNil(o.RecordId) {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordId) {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasRecordId() bool {
	if o != nil && !IsNil(o.RecordId) {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) SetRecordId(v string) {
	o.RecordId = &v
}

// GetSellerId returns the SellerId field value if set, zero value otherwise.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetSellerId() string {
	if o == nil || IsNil(o.SellerId) {
		var ret string
		return ret
	}
	return *o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) GetSellerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellerId) {
		return nil, false
	}
	return o.SellerId, true
}

// HasSellerId returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasSellerId() bool {
	if o != nil && !IsNil(o.SellerId) {
		return true
	}

	return false
}

// SetSellerId gets a reference to the given string and assigns it to the SellerId field.
// Deprecated
func (o *AlipayMerchantOrderSyncModel) SetSellerId(v string) {
	o.SellerId = &v
}

// GetSendMsg returns the SendMsg field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetSendMsg() string {
	if o == nil || IsNil(o.SendMsg) {
		var ret string
		return ret
	}
	return *o.SendMsg
}

// GetSendMsgOk returns a tuple with the SendMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetSendMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SendMsg) {
		return nil, false
	}
	return o.SendMsg, true
}

// HasSendMsg returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasSendMsg() bool {
	if o != nil && !IsNil(o.SendMsg) {
		return true
	}

	return false
}

// SetSendMsg gets a reference to the given string and assigns it to the SendMsg field.
func (o *AlipayMerchantOrderSyncModel) SetSendMsg(v string) {
	o.SendMsg = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *AlipayMerchantOrderSyncModel) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetShopInfo returns the ShopInfo field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetShopInfo() OrderShopInfo {
	if o == nil || IsNil(o.ShopInfo) {
		var ret OrderShopInfo
		return ret
	}
	return *o.ShopInfo
}

// GetShopInfoOk returns a tuple with the ShopInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetShopInfoOk() (*OrderShopInfo, bool) {
	if o == nil || IsNil(o.ShopInfo) {
		return nil, false
	}
	return o.ShopInfo, true
}

// HasShopInfo returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasShopInfo() bool {
	if o != nil && !IsNil(o.ShopInfo) {
		return true
	}

	return false
}

// SetShopInfo gets a reference to the given OrderShopInfo and assigns it to the ShopInfo field.
func (o *AlipayMerchantOrderSyncModel) SetShopInfo(v OrderShopInfo) {
	o.ShopInfo = &v
}

// GetSourceApp returns the SourceApp field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetSourceApp() string {
	if o == nil || IsNil(o.SourceApp) {
		var ret string
		return ret
	}
	return *o.SourceApp
}

// GetSourceAppOk returns a tuple with the SourceApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetSourceAppOk() (*string, bool) {
	if o == nil || IsNil(o.SourceApp) {
		return nil, false
	}
	return o.SourceApp, true
}

// HasSourceApp returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasSourceApp() bool {
	if o != nil && !IsNil(o.SourceApp) {
		return true
	}

	return false
}

// SetSourceApp gets a reference to the given string and assigns it to the SourceApp field.
func (o *AlipayMerchantOrderSyncModel) SetSourceApp(v string) {
	o.SourceApp = &v
}

// GetSyncContent returns the SyncContent field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetSyncContent() string {
	if o == nil || IsNil(o.SyncContent) {
		var ret string
		return ret
	}
	return *o.SyncContent
}

// GetSyncContentOk returns a tuple with the SyncContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetSyncContentOk() (*string, bool) {
	if o == nil || IsNil(o.SyncContent) {
		return nil, false
	}
	return o.SyncContent, true
}

// HasSyncContent returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasSyncContent() bool {
	if o != nil && !IsNil(o.SyncContent) {
		return true
	}

	return false
}

// SetSyncContent gets a reference to the given string and assigns it to the SyncContent field.
func (o *AlipayMerchantOrderSyncModel) SetSyncContent(v string) {
	o.SyncContent = &v
}

// GetTicketInfo returns the TicketInfo field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetTicketInfo() TicketInfo {
	if o == nil || IsNil(o.TicketInfo) {
		var ret TicketInfo
		return ret
	}
	return *o.TicketInfo
}

// GetTicketInfoOk returns a tuple with the TicketInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetTicketInfoOk() (*TicketInfo, bool) {
	if o == nil || IsNil(o.TicketInfo) {
		return nil, false
	}
	return o.TicketInfo, true
}

// HasTicketInfo returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasTicketInfo() bool {
	if o != nil && !IsNil(o.TicketInfo) {
		return true
	}

	return false
}

// SetTicketInfo gets a reference to the given TicketInfo and assigns it to the TicketInfo field.
func (o *AlipayMerchantOrderSyncModel) SetTicketInfo(v TicketInfo) {
	o.TicketInfo = &v
}

// GetTicketOrderList returns the TicketOrderList field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetTicketOrderList() []TicketOrderInfo {
	if o == nil || IsNil(o.TicketOrderList) {
		var ret []TicketOrderInfo
		return ret
	}
	return o.TicketOrderList
}

// GetTicketOrderListOk returns a tuple with the TicketOrderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetTicketOrderListOk() ([]TicketOrderInfo, bool) {
	if o == nil || IsNil(o.TicketOrderList) {
		return nil, false
	}
	return o.TicketOrderList, true
}

// HasTicketOrderList returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasTicketOrderList() bool {
	if o != nil && !IsNil(o.TicketOrderList) {
		return true
	}

	return false
}

// SetTicketOrderList gets a reference to the given []TicketOrderInfo and assigns it to the TicketOrderList field.
func (o *AlipayMerchantOrderSyncModel) SetTicketOrderList(v []TicketOrderInfo) {
	o.TicketOrderList = v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayMerchantOrderSyncModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *AlipayMerchantOrderSyncModel) GetTradeType() string {
	if o == nil || IsNil(o.TradeType) {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantOrderSyncModel) GetTradeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TradeType) {
		return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *AlipayMerchantOrderSyncModel) HasTradeType() bool {
	if o != nil && !IsNil(o.TradeType) {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *AlipayMerchantOrderSyncModel) SetTradeType(v string) {
	o.TradeType = &v
}

func (o AlipayMerchantOrderSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantOrderSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BuyerId) {
		toSerialize["buyer_id"] = o.BuyerId
	}
	if !IsNil(o.BuyerInfo) {
		toSerialize["buyer_info"] = o.BuyerInfo
	}
	if !IsNil(o.BuyerOpenId) {
		toSerialize["buyer_open_id"] = o.BuyerOpenId
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountInfoList) {
		toSerialize["discount_info_list"] = o.DiscountInfoList
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.ItemOrderList) {
		toSerialize["item_order_list"] = o.ItemOrderList
	}
	if !IsNil(o.JourneyOrderList) {
		toSerialize["journey_order_list"] = o.JourneyOrderList
	}
	if !IsNil(o.LogisticsInfoList) {
		toSerialize["logistics_info_list"] = o.LogisticsInfoList
	}
	if !IsNil(o.OrderAuthCode) {
		toSerialize["order_auth_code"] = o.OrderAuthCode
	}
	if !IsNil(o.OrderCreateTime) {
		toSerialize["order_create_time"] = o.OrderCreateTime
	}
	if !IsNil(o.OrderModifiedTime) {
		toSerialize["order_modified_time"] = o.OrderModifiedTime
	}
	if !IsNil(o.OrderPayTime) {
		toSerialize["order_pay_time"] = o.OrderPayTime
	}
	if !IsNil(o.OrderType) {
		toSerialize["order_type"] = o.OrderType
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.PayTimeoutExpress) {
		toSerialize["pay_timeout_express"] = o.PayTimeoutExpress
	}
	if !IsNil(o.RecordId) {
		toSerialize["record_id"] = o.RecordId
	}
	if !IsNil(o.SellerId) {
		toSerialize["seller_id"] = o.SellerId
	}
	if !IsNil(o.SendMsg) {
		toSerialize["send_msg"] = o.SendMsg
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["service_code"] = o.ServiceCode
	}
	if !IsNil(o.ShopInfo) {
		toSerialize["shop_info"] = o.ShopInfo
	}
	if !IsNil(o.SourceApp) {
		toSerialize["source_app"] = o.SourceApp
	}
	if !IsNil(o.SyncContent) {
		toSerialize["sync_content"] = o.SyncContent
	}
	if !IsNil(o.TicketInfo) {
		toSerialize["ticket_info"] = o.TicketInfo
	}
	if !IsNil(o.TicketOrderList) {
		toSerialize["ticket_order_list"] = o.TicketOrderList
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	if !IsNil(o.TradeType) {
		toSerialize["trade_type"] = o.TradeType
	}
	return toSerialize, nil
}

type NullableAlipayMerchantOrderSyncModel struct {
	value *AlipayMerchantOrderSyncModel
	isSet bool
}

func (v NullableAlipayMerchantOrderSyncModel) Get() *AlipayMerchantOrderSyncModel {
	return v.value
}

func (v *NullableAlipayMerchantOrderSyncModel) Set(val *AlipayMerchantOrderSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantOrderSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantOrderSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantOrderSyncModel(val *AlipayMerchantOrderSyncModel) *NullableAlipayMerchantOrderSyncModel {
	return &NullableAlipayMerchantOrderSyncModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantOrderSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantOrderSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


