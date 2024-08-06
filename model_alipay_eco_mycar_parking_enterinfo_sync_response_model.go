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

// checks if the AlipayEcoMycarParkingEnterinfoSyncResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoMycarParkingEnterinfoSyncResponseModel{}

// AlipayEcoMycarParkingEnterinfoSyncResponseModel struct for AlipayEcoMycarParkingEnterinfoSyncResponseModel
type AlipayEcoMycarParkingEnterinfoSyncResponseModel struct {
	// 用户签约的代扣场景字符集，多个英文逗号分割，当车场为ORC识别车牌的场景，返回值包含PLATE_PAY时，表示代扣协议可用。 当车场为ETC设备识别车牌，返回值包含ETC_PAY时表示协议可用。当用户未签约任何代扣场景时返回NO_AGREEMENT_SCENE
	AgreementScene *string `json:"agreement_scene,omitempty"`
	// 该字段已废弃,请使用agreement_scene字段
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// 当前车辆在支付宝侧的信息，是否可以使用无感支付，可选返回项为： NORMAL：正常缴费用户 NON-SENSE-PAY：可无感支付用户
	CarStatus *string `json:"car_status,omitempty"`
	// 当前时间戳(查询签约状态和是否支持无感支付仅代表当前时间点查询结果，不作为最后代扣结果，应在发起代扣前再次查询)
	CurrentTime *string `json:"current_time,omitempty"`
	// 支付宝业务流水号，用于记录车辆从驶入到驶出的全流程
	SerialNo *string `json:"serial_no,omitempty"`
}

// NewAlipayEcoMycarParkingEnterinfoSyncResponseModel instantiates a new AlipayEcoMycarParkingEnterinfoSyncResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoMycarParkingEnterinfoSyncResponseModel() *AlipayEcoMycarParkingEnterinfoSyncResponseModel {
	this := AlipayEcoMycarParkingEnterinfoSyncResponseModel{}
	return &this
}

// NewAlipayEcoMycarParkingEnterinfoSyncResponseModelWithDefaults instantiates a new AlipayEcoMycarParkingEnterinfoSyncResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoMycarParkingEnterinfoSyncResponseModelWithDefaults() *AlipayEcoMycarParkingEnterinfoSyncResponseModel {
	this := AlipayEcoMycarParkingEnterinfoSyncResponseModel{}
	return &this
}

// GetAgreementScene returns the AgreementScene field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetAgreementScene() string {
	if o == nil || IsNil(o.AgreementScene) {
		var ret string
		return ret
	}
	return *o.AgreementScene
}

// GetAgreementSceneOk returns a tuple with the AgreementScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetAgreementSceneOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementScene) {
		return nil, false
	}
	return o.AgreementScene, true
}

// HasAgreementScene returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) HasAgreementScene() bool {
	if o != nil && !IsNil(o.AgreementScene) {
		return true
	}

	return false
}

// SetAgreementScene gets a reference to the given string and assigns it to the AgreementScene field.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) SetAgreementScene(v string) {
	o.AgreementScene = &v
}

// GetAgreementStatus returns the AgreementStatus field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetAgreementStatus() string {
	if o == nil || IsNil(o.AgreementStatus) {
		var ret string
		return ret
	}
	return *o.AgreementStatus
}

// GetAgreementStatusOk returns a tuple with the AgreementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetAgreementStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementStatus) {
		return nil, false
	}
	return o.AgreementStatus, true
}

// HasAgreementStatus returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) HasAgreementStatus() bool {
	if o != nil && !IsNil(o.AgreementStatus) {
		return true
	}

	return false
}

// SetAgreementStatus gets a reference to the given string and assigns it to the AgreementStatus field.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) SetAgreementStatus(v string) {
	o.AgreementStatus = &v
}

// GetCarStatus returns the CarStatus field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetCarStatus() string {
	if o == nil || IsNil(o.CarStatus) {
		var ret string
		return ret
	}
	return *o.CarStatus
}

// GetCarStatusOk returns a tuple with the CarStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetCarStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CarStatus) {
		return nil, false
	}
	return o.CarStatus, true
}

// HasCarStatus returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) HasCarStatus() bool {
	if o != nil && !IsNil(o.CarStatus) {
		return true
	}

	return false
}

// SetCarStatus gets a reference to the given string and assigns it to the CarStatus field.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) SetCarStatus(v string) {
	o.CarStatus = &v
}

// GetCurrentTime returns the CurrentTime field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetCurrentTime() string {
	if o == nil || IsNil(o.CurrentTime) {
		var ret string
		return ret
	}
	return *o.CurrentTime
}

// GetCurrentTimeOk returns a tuple with the CurrentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetCurrentTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentTime) {
		return nil, false
	}
	return o.CurrentTime, true
}

// HasCurrentTime returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) HasCurrentTime() bool {
	if o != nil && !IsNil(o.CurrentTime) {
		return true
	}

	return false
}

// SetCurrentTime gets a reference to the given string and assigns it to the CurrentTime field.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) SetCurrentTime(v string) {
	o.CurrentTime = &v
}

// GetSerialNo returns the SerialNo field value if set, zero value otherwise.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetSerialNo() string {
	if o == nil || IsNil(o.SerialNo) {
		var ret string
		return ret
	}
	return *o.SerialNo
}

// GetSerialNoOk returns a tuple with the SerialNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) GetSerialNoOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNo) {
		return nil, false
	}
	return o.SerialNo, true
}

// HasSerialNo returns a boolean if a field has been set.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) HasSerialNo() bool {
	if o != nil && !IsNil(o.SerialNo) {
		return true
	}

	return false
}

// SetSerialNo gets a reference to the given string and assigns it to the SerialNo field.
func (o *AlipayEcoMycarParkingEnterinfoSyncResponseModel) SetSerialNo(v string) {
	o.SerialNo = &v
}

func (o AlipayEcoMycarParkingEnterinfoSyncResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoMycarParkingEnterinfoSyncResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementScene) {
		toSerialize["agreement_scene"] = o.AgreementScene
	}
	if !IsNil(o.AgreementStatus) {
		toSerialize["agreement_status"] = o.AgreementStatus
	}
	if !IsNil(o.CarStatus) {
		toSerialize["car_status"] = o.CarStatus
	}
	if !IsNil(o.CurrentTime) {
		toSerialize["current_time"] = o.CurrentTime
	}
	if !IsNil(o.SerialNo) {
		toSerialize["serial_no"] = o.SerialNo
	}
	return toSerialize, nil
}

type NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel struct {
	value *AlipayEcoMycarParkingEnterinfoSyncResponseModel
	isSet bool
}

func (v NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel) Get() *AlipayEcoMycarParkingEnterinfoSyncResponseModel {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel) Set(val *AlipayEcoMycarParkingEnterinfoSyncResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingEnterinfoSyncResponseModel(val *AlipayEcoMycarParkingEnterinfoSyncResponseModel) *NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel {
	return &NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingEnterinfoSyncResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
