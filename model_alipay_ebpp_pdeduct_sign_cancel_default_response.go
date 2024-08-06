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


// AlipayEbppPdeductSignCancelDefaultResponse struct for AlipayEbppPdeductSignCancelDefaultResponse
type AlipayEbppPdeductSignCancelDefaultResponse struct {
	AlipayEbppPdeductSignCancelErrorResponseModel *AlipayEbppPdeductSignCancelErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppPdeductSignCancelDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppPdeductSignCancelErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppPdeductSignCancelErrorResponseModel);
	if err == nil {
		jsonAlipayEbppPdeductSignCancelErrorResponseModel, _ := json.Marshal(dst.AlipayEbppPdeductSignCancelErrorResponseModel)
		if string(jsonAlipayEbppPdeductSignCancelErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppPdeductSignCancelErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppPdeductSignCancelErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppPdeductSignCancelErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppPdeductSignCancelDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppPdeductSignCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppPdeductSignCancelErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppPdeductSignCancelErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppPdeductSignCancelDefaultResponse struct {
	value *AlipayEbppPdeductSignCancelDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppPdeductSignCancelDefaultResponse) Get() *AlipayEbppPdeductSignCancelDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppPdeductSignCancelDefaultResponse) Set(val *AlipayEbppPdeductSignCancelDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppPdeductSignCancelDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppPdeductSignCancelDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppPdeductSignCancelDefaultResponse(val *AlipayEbppPdeductSignCancelDefaultResponse) *NullableAlipayEbppPdeductSignCancelDefaultResponse {
	return &NullableAlipayEbppPdeductSignCancelDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppPdeductSignCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppPdeductSignCancelDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


