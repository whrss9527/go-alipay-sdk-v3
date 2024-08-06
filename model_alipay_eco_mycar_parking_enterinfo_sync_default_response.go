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


// AlipayEcoMycarParkingEnterinfoSyncDefaultResponse struct for AlipayEcoMycarParkingEnterinfoSyncDefaultResponse
type AlipayEcoMycarParkingEnterinfoSyncDefaultResponse struct {
	AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel *AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingEnterinfoSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel);
	if err == nil {
		jsonAlipayEcoMycarParkingEnterinfoSyncErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingEnterinfoSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingEnterinfoSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingEnterinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingEnterinfoSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse struct {
	value *AlipayEcoMycarParkingEnterinfoSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse) Get() *AlipayEcoMycarParkingEnterinfoSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse) Set(val *AlipayEcoMycarParkingEnterinfoSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse(val *AlipayEcoMycarParkingEnterinfoSyncDefaultResponse) *NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse {
	return &NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingEnterinfoSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

