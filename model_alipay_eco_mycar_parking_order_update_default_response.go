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


// AlipayEcoMycarParkingOrderUpdateDefaultResponse struct for AlipayEcoMycarParkingOrderUpdateDefaultResponse
type AlipayEcoMycarParkingOrderUpdateDefaultResponse struct {
	AlipayEcoMycarParkingOrderUpdateErrorResponseModel *AlipayEcoMycarParkingOrderUpdateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingOrderUpdateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingOrderUpdateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingOrderUpdateErrorResponseModel);
	if err == nil {
		jsonAlipayEcoMycarParkingOrderUpdateErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingOrderUpdateErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingOrderUpdateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingOrderUpdateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingOrderUpdateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingOrderUpdateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingOrderUpdateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingOrderUpdateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingOrderUpdateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingOrderUpdateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse struct {
	value *AlipayEcoMycarParkingOrderUpdateDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse) Get() *AlipayEcoMycarParkingOrderUpdateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse) Set(val *AlipayEcoMycarParkingOrderUpdateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingOrderUpdateDefaultResponse(val *AlipayEcoMycarParkingOrderUpdateDefaultResponse) *NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse {
	return &NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingOrderUpdateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


