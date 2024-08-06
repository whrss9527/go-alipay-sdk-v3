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

// checks if the AlipayCommerceTransportParkingExitinfoSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceTransportParkingExitinfoSyncModel{}

// AlipayCommerceTransportParkingExitinfoSyncModel struct for AlipayCommerceTransportParkingExitinfoSyncModel
type AlipayCommerceTransportParkingExitinfoSyncModel struct {
	// 车牌是否加密，true为加密，false为不加密，默认为false
	IsEncryptPlateNo *bool `json:"is_encrypt_plate_no,omitempty"`
	// 蚂蚁会员统一ID对应的归属应用appid
	OpenAppid *string `json:"open_appid,omitempty"`
	// 蚂蚁会员统一ID
	OpenId *string `json:"open_id,omitempty"`
	// 外部停车流水号(用于串通进场与出场信息)
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// 车辆出场的时间，格式\"YYYY-MM-DD HH:mm:ss\"，24小时制，请保证服务器时间准确，出场时间不应晚于当前网络时间，也不当早于入场时间。
	OutTime *string `json:"out_time,omitempty"`
	// 车牌颜色，车牌颜色，枚举支持： *BLUE：蓝。 *GREEN：绿。 *YELLOW：黄。 *WHITE：白。 *BLACK：黑。 *LIMEGREEN：黄绿色。
	PlateColor *string `json:"plate_color,omitempty"`
	// 车牌号（支持加密格式）
	PlateNo *string `json:"plate_no,omitempty"`
	// 停车服务页面地址。 1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见：https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F 2、若服务商没有服务链接，可传输支付宝停车官方小程序的服务链接：alipays://platformapi/startapp?appId=2021001102642986&page=pages%2Fparking-fee%2Findex 3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex
	ServiceUrl *string `json:"service_url,omitempty"`
}

// NewAlipayCommerceTransportParkingExitinfoSyncModel instantiates a new AlipayCommerceTransportParkingExitinfoSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceTransportParkingExitinfoSyncModel() *AlipayCommerceTransportParkingExitinfoSyncModel {
	this := AlipayCommerceTransportParkingExitinfoSyncModel{}
	return &this
}

// NewAlipayCommerceTransportParkingExitinfoSyncModelWithDefaults instantiates a new AlipayCommerceTransportParkingExitinfoSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceTransportParkingExitinfoSyncModelWithDefaults() *AlipayCommerceTransportParkingExitinfoSyncModel {
	this := AlipayCommerceTransportParkingExitinfoSyncModel{}
	return &this
}

// GetIsEncryptPlateNo returns the IsEncryptPlateNo field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetIsEncryptPlateNo() bool {
	if o == nil || IsNil(o.IsEncryptPlateNo) {
		var ret bool
		return ret
	}
	return *o.IsEncryptPlateNo
}

// GetIsEncryptPlateNoOk returns a tuple with the IsEncryptPlateNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetIsEncryptPlateNoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEncryptPlateNo) {
		return nil, false
	}
	return o.IsEncryptPlateNo, true
}

// HasIsEncryptPlateNo returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasIsEncryptPlateNo() bool {
	if o != nil && !IsNil(o.IsEncryptPlateNo) {
		return true
	}

	return false
}

// SetIsEncryptPlateNo gets a reference to the given bool and assigns it to the IsEncryptPlateNo field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetIsEncryptPlateNo(v bool) {
	o.IsEncryptPlateNo = &v
}

// GetOpenAppid returns the OpenAppid field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOpenAppid() string {
	if o == nil || IsNil(o.OpenAppid) {
		var ret string
		return ret
	}
	return *o.OpenAppid
}

// GetOpenAppidOk returns a tuple with the OpenAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOpenAppidOk() (*string, bool) {
	if o == nil || IsNil(o.OpenAppid) {
		return nil, false
	}
	return o.OpenAppid, true
}

// HasOpenAppid returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasOpenAppid() bool {
	if o != nil && !IsNil(o.OpenAppid) {
		return true
	}

	return false
}

// SetOpenAppid gets a reference to the given string and assigns it to the OpenAppid field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetOpenAppid(v string) {
	o.OpenAppid = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOutSerialNo returns the OutSerialNo field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOutSerialNo() string {
	if o == nil || IsNil(o.OutSerialNo) {
		var ret string
		return ret
	}
	return *o.OutSerialNo
}

// GetOutSerialNoOk returns a tuple with the OutSerialNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOutSerialNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutSerialNo) {
		return nil, false
	}
	return o.OutSerialNo, true
}

// HasOutSerialNo returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasOutSerialNo() bool {
	if o != nil && !IsNil(o.OutSerialNo) {
		return true
	}

	return false
}

// SetOutSerialNo gets a reference to the given string and assigns it to the OutSerialNo field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetOutSerialNo(v string) {
	o.OutSerialNo = &v
}

// GetOutTime returns the OutTime field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOutTime() string {
	if o == nil || IsNil(o.OutTime) {
		var ret string
		return ret
	}
	return *o.OutTime
}

// GetOutTimeOk returns a tuple with the OutTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetOutTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OutTime) {
		return nil, false
	}
	return o.OutTime, true
}

// HasOutTime returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasOutTime() bool {
	if o != nil && !IsNil(o.OutTime) {
		return true
	}

	return false
}

// SetOutTime gets a reference to the given string and assigns it to the OutTime field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetOutTime(v string) {
	o.OutTime = &v
}

// GetPlateColor returns the PlateColor field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetPlateColor() string {
	if o == nil || IsNil(o.PlateColor) {
		var ret string
		return ret
	}
	return *o.PlateColor
}

// GetPlateColorOk returns a tuple with the PlateColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetPlateColorOk() (*string, bool) {
	if o == nil || IsNil(o.PlateColor) {
		return nil, false
	}
	return o.PlateColor, true
}

// HasPlateColor returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasPlateColor() bool {
	if o != nil && !IsNil(o.PlateColor) {
		return true
	}

	return false
}

// SetPlateColor gets a reference to the given string and assigns it to the PlateColor field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetPlateColor(v string) {
	o.PlateColor = &v
}

// GetPlateNo returns the PlateNo field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetPlateNo() string {
	if o == nil || IsNil(o.PlateNo) {
		var ret string
		return ret
	}
	return *o.PlateNo
}

// GetPlateNoOk returns a tuple with the PlateNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetPlateNoOk() (*string, bool) {
	if o == nil || IsNil(o.PlateNo) {
		return nil, false
	}
	return o.PlateNo, true
}

// HasPlateNo returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasPlateNo() bool {
	if o != nil && !IsNil(o.PlateNo) {
		return true
	}

	return false
}

// SetPlateNo gets a reference to the given string and assigns it to the PlateNo field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetPlateNo(v string) {
	o.PlateNo = &v
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetServiceUrl() string {
	if o == nil || IsNil(o.ServiceUrl) {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) GetServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUrl) {
		return nil, false
	}
	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) HasServiceUrl() bool {
	if o != nil && !IsNil(o.ServiceUrl) {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *AlipayCommerceTransportParkingExitinfoSyncModel) SetServiceUrl(v string) {
	o.ServiceUrl = &v
}

func (o AlipayCommerceTransportParkingExitinfoSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceTransportParkingExitinfoSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsEncryptPlateNo) {
		toSerialize["is_encrypt_plate_no"] = o.IsEncryptPlateNo
	}
	if !IsNil(o.OpenAppid) {
		toSerialize["open_appid"] = o.OpenAppid
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OutSerialNo) {
		toSerialize["out_serial_no"] = o.OutSerialNo
	}
	if !IsNil(o.OutTime) {
		toSerialize["out_time"] = o.OutTime
	}
	if !IsNil(o.PlateColor) {
		toSerialize["plate_color"] = o.PlateColor
	}
	if !IsNil(o.PlateNo) {
		toSerialize["plate_no"] = o.PlateNo
	}
	if !IsNil(o.ServiceUrl) {
		toSerialize["service_url"] = o.ServiceUrl
	}
	return toSerialize, nil
}

type NullableAlipayCommerceTransportParkingExitinfoSyncModel struct {
	value *AlipayCommerceTransportParkingExitinfoSyncModel
	isSet bool
}

func (v NullableAlipayCommerceTransportParkingExitinfoSyncModel) Get() *AlipayCommerceTransportParkingExitinfoSyncModel {
	return v.value
}

func (v *NullableAlipayCommerceTransportParkingExitinfoSyncModel) Set(val *AlipayCommerceTransportParkingExitinfoSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceTransportParkingExitinfoSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceTransportParkingExitinfoSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceTransportParkingExitinfoSyncModel(val *AlipayCommerceTransportParkingExitinfoSyncModel) *NullableAlipayCommerceTransportParkingExitinfoSyncModel {
	return &NullableAlipayCommerceTransportParkingExitinfoSyncModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceTransportParkingExitinfoSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceTransportParkingExitinfoSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

