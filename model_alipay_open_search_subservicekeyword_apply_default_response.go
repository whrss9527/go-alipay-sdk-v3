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


// AlipayOpenSearchSubservicekeywordApplyDefaultResponse struct for AlipayOpenSearchSubservicekeywordApplyDefaultResponse
type AlipayOpenSearchSubservicekeywordApplyDefaultResponse struct {
	AlipayOpenSearchSubservicekeywordApplyErrorResponseModel *AlipayOpenSearchSubservicekeywordApplyErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchSubservicekeywordApplyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchSubservicekeywordApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchSubservicekeywordApplyErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSearchSubservicekeywordApplyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchSubservicekeywordApplyErrorResponseModel)
		if string(jsonAlipayOpenSearchSubservicekeywordApplyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchSubservicekeywordApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchSubservicekeywordApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchSubservicekeywordApplyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchSubservicekeywordApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchSubservicekeywordApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchSubservicekeywordApplyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchSubservicekeywordApplyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse struct {
	value *AlipayOpenSearchSubservicekeywordApplyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse) Get() *AlipayOpenSearchSubservicekeywordApplyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse) Set(val *AlipayOpenSearchSubservicekeywordApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse(val *AlipayOpenSearchSubservicekeywordApplyDefaultResponse) *NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse {
	return &NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


