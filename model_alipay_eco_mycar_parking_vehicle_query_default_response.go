/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// AlipayEcoMycarParkingVehicleQueryDefaultResponse struct for AlipayEcoMycarParkingVehicleQueryDefaultResponse
type AlipayEcoMycarParkingVehicleQueryDefaultResponse struct {
	AlipayEcoMycarParkingVehicleQueryErrorResponseModel *AlipayEcoMycarParkingVehicleQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingVehicleQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingVehicleQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingVehicleQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEcoMycarParkingVehicleQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingVehicleQueryErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingVehicleQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingVehicleQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingVehicleQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingVehicleQueryErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType);
	if err == nil {
		jsonCommonErrorType, _ := json.Marshal(dst.CommonErrorType)
		if string(jsonCommonErrorType) == "{}" { // empty struct
			dst.CommonErrorType = nil
		} else {
			return nil // data stored in dst.CommonErrorType, return on the first match
		}
	} else {
		dst.CommonErrorType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingVehicleQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingVehicleQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingVehicleQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingVehicleQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse struct {
	value *AlipayEcoMycarParkingVehicleQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse) Get() *AlipayEcoMycarParkingVehicleQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse) Set(val *AlipayEcoMycarParkingVehicleQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingVehicleQueryDefaultResponse(val *AlipayEcoMycarParkingVehicleQueryDefaultResponse) *NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse {
	return &NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingVehicleQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


