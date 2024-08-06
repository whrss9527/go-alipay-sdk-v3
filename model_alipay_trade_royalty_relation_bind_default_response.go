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


// AlipayTradeRoyaltyRelationBindDefaultResponse struct for AlipayTradeRoyaltyRelationBindDefaultResponse
type AlipayTradeRoyaltyRelationBindDefaultResponse struct {
	AlipayTradeRoyaltyRelationBindErrorResponseModel *AlipayTradeRoyaltyRelationBindErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayTradeRoyaltyRelationBindDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayTradeRoyaltyRelationBindErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayTradeRoyaltyRelationBindErrorResponseModel);
	if err == nil {
		jsonAlipayTradeRoyaltyRelationBindErrorResponseModel, _ := json.Marshal(dst.AlipayTradeRoyaltyRelationBindErrorResponseModel)
		if string(jsonAlipayTradeRoyaltyRelationBindErrorResponseModel) == "{}" { // empty struct
			dst.AlipayTradeRoyaltyRelationBindErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayTradeRoyaltyRelationBindErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayTradeRoyaltyRelationBindErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayTradeRoyaltyRelationBindDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayTradeRoyaltyRelationBindDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayTradeRoyaltyRelationBindErrorResponseModel != nil {
		return json.Marshal(&src.AlipayTradeRoyaltyRelationBindErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayTradeRoyaltyRelationBindDefaultResponse struct {
	value *AlipayTradeRoyaltyRelationBindDefaultResponse
	isSet bool
}

func (v NullableAlipayTradeRoyaltyRelationBindDefaultResponse) Get() *AlipayTradeRoyaltyRelationBindDefaultResponse {
	return v.value
}

func (v *NullableAlipayTradeRoyaltyRelationBindDefaultResponse) Set(val *AlipayTradeRoyaltyRelationBindDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeRoyaltyRelationBindDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeRoyaltyRelationBindDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeRoyaltyRelationBindDefaultResponse(val *AlipayTradeRoyaltyRelationBindDefaultResponse) *NullableAlipayTradeRoyaltyRelationBindDefaultResponse {
	return &NullableAlipayTradeRoyaltyRelationBindDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayTradeRoyaltyRelationBindDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeRoyaltyRelationBindDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


