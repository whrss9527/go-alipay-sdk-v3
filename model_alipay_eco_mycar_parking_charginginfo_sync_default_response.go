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


// AlipayEcoMycarParkingCharginginfoSyncDefaultResponse struct for AlipayEcoMycarParkingCharginginfoSyncDefaultResponse
type AlipayEcoMycarParkingCharginginfoSyncDefaultResponse struct {
	AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel *AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingCharginginfoSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel);
	if err == nil {
		jsonAlipayEcoMycarParkingCharginginfoSyncErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingCharginginfoSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingCharginginfoSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingCharginginfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingCharginginfoSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse struct {
	value *AlipayEcoMycarParkingCharginginfoSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse) Get() *AlipayEcoMycarParkingCharginginfoSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse) Set(val *AlipayEcoMycarParkingCharginginfoSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse(val *AlipayEcoMycarParkingCharginginfoSyncDefaultResponse) *NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse {
	return &NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingCharginginfoSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


