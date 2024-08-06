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

// checks if the AlipayFundUsertradeBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundUsertradeBatchqueryModel{}

// AlipayFundUsertradeBatchqueryModel struct for AlipayFundUsertradeBatchqueryModel
type AlipayFundUsertradeBatchqueryModel struct {
	// 查询交易数据场景码
	BizScene *string `json:"biz_scene,omitempty"`
	// 拉取的结算时间，格式yyyy-MM-dd HH:mm:ss，拉取逻辑包含该时刻。
	EndTime *string `json:"end_time,omitempty"`
	// 分页页码，从1开始，必须大于0
	PageIndex *string `json:"page_index,omitempty"`
	// 分页大小，必须大于0，最大设置100
	PageSize *string `json:"page_size,omitempty"`
	// 查询交易数据产品码
	ProductCode *string `json:"product_code,omitempty"`
	// 拉取的起始时间，格式yyyy-MM-dd HH:mm:ss，时间必须早于拉取的截止时间，并且，接口仅限查询用户30天内交易数据，拉取逻辑包含该时刻。
	StartTime *string `json:"start_time,omitempty"`
}

// NewAlipayFundUsertradeBatchqueryModel instantiates a new AlipayFundUsertradeBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundUsertradeBatchqueryModel() *AlipayFundUsertradeBatchqueryModel {
	this := AlipayFundUsertradeBatchqueryModel{}
	return &this
}

// NewAlipayFundUsertradeBatchqueryModelWithDefaults instantiates a new AlipayFundUsertradeBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundUsertradeBatchqueryModelWithDefaults() *AlipayFundUsertradeBatchqueryModel {
	this := AlipayFundUsertradeBatchqueryModel{}
	return &this
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundUsertradeBatchqueryModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundUsertradeBatchqueryModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundUsertradeBatchqueryModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundUsertradeBatchqueryModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AlipayFundUsertradeBatchqueryModel) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundUsertradeBatchqueryModel) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AlipayFundUsertradeBatchqueryModel) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *AlipayFundUsertradeBatchqueryModel) SetEndTime(v string) {
	o.EndTime = &v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise.
func (o *AlipayFundUsertradeBatchqueryModel) GetPageIndex() string {
	if o == nil || IsNil(o.PageIndex) {
		var ret string
		return ret
	}
	return *o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundUsertradeBatchqueryModel) GetPageIndexOk() (*string, bool) {
	if o == nil || IsNil(o.PageIndex) {
		return nil, false
	}
	return o.PageIndex, true
}

// HasPageIndex returns a boolean if a field has been set.
func (o *AlipayFundUsertradeBatchqueryModel) HasPageIndex() bool {
	if o != nil && !IsNil(o.PageIndex) {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given string and assigns it to the PageIndex field.
func (o *AlipayFundUsertradeBatchqueryModel) SetPageIndex(v string) {
	o.PageIndex = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayFundUsertradeBatchqueryModel) GetPageSize() string {
	if o == nil || IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundUsertradeBatchqueryModel) GetPageSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayFundUsertradeBatchqueryModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *AlipayFundUsertradeBatchqueryModel) SetPageSize(v string) {
	o.PageSize = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundUsertradeBatchqueryModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundUsertradeBatchqueryModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundUsertradeBatchqueryModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundUsertradeBatchqueryModel) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AlipayFundUsertradeBatchqueryModel) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundUsertradeBatchqueryModel) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AlipayFundUsertradeBatchqueryModel) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *AlipayFundUsertradeBatchqueryModel) SetStartTime(v string) {
	o.StartTime = &v
}

func (o AlipayFundUsertradeBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundUsertradeBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.PageIndex) {
		toSerialize["page_index"] = o.PageIndex
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableAlipayFundUsertradeBatchqueryModel struct {
	value *AlipayFundUsertradeBatchqueryModel
	isSet bool
}

func (v NullableAlipayFundUsertradeBatchqueryModel) Get() *AlipayFundUsertradeBatchqueryModel {
	return v.value
}

func (v *NullableAlipayFundUsertradeBatchqueryModel) Set(val *AlipayFundUsertradeBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundUsertradeBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundUsertradeBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundUsertradeBatchqueryModel(val *AlipayFundUsertradeBatchqueryModel) *NullableAlipayFundUsertradeBatchqueryModel {
	return &NullableAlipayFundUsertradeBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayFundUsertradeBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundUsertradeBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


