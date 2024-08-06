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


// AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse struct for AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse
type AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse struct {
	AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel *AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel);
	if err == nil {
		jsonAlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingParkinglotinfoCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse struct {
	value *AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) Get() *AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) Set(val *AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse(val *AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) *NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse {
	return &NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


