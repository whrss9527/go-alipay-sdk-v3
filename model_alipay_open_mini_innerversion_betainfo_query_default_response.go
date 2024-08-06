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


// AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse struct for AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse
type AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse struct {
	AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel *AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionBetainfoQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse struct {
	value *AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) Get() *AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) Set(val *AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse(val *AlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) *NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionBetainfoQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


