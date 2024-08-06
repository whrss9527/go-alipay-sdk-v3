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


// AlipayEcoEduKtBillingQueryDefaultResponse struct for AlipayEcoEduKtBillingQueryDefaultResponse
type AlipayEcoEduKtBillingQueryDefaultResponse struct {
	AlipayEcoEduKtBillingQueryErrorResponseModel *AlipayEcoEduKtBillingQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoEduKtBillingQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoEduKtBillingQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoEduKtBillingQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEcoEduKtBillingQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEcoEduKtBillingQueryErrorResponseModel)
		if string(jsonAlipayEcoEduKtBillingQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoEduKtBillingQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoEduKtBillingQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoEduKtBillingQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoEduKtBillingQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoEduKtBillingQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoEduKtBillingQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoEduKtBillingQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoEduKtBillingQueryDefaultResponse struct {
	value *AlipayEcoEduKtBillingQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoEduKtBillingQueryDefaultResponse) Get() *AlipayEcoEduKtBillingQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoEduKtBillingQueryDefaultResponse) Set(val *AlipayEcoEduKtBillingQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoEduKtBillingQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoEduKtBillingQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoEduKtBillingQueryDefaultResponse(val *AlipayEcoEduKtBillingQueryDefaultResponse) *NullableAlipayEcoEduKtBillingQueryDefaultResponse {
	return &NullableAlipayEcoEduKtBillingQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoEduKtBillingQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoEduKtBillingQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


