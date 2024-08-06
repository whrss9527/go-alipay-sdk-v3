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


// AlipayMerchantOrderSyncDefaultResponse struct for AlipayMerchantOrderSyncDefaultResponse
type AlipayMerchantOrderSyncDefaultResponse struct {
	AlipayMerchantOrderSyncErrorResponseModel *AlipayMerchantOrderSyncErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMerchantOrderSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMerchantOrderSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMerchantOrderSyncErrorResponseModel);
	if err == nil {
		jsonAlipayMerchantOrderSyncErrorResponseModel, _ := json.Marshal(dst.AlipayMerchantOrderSyncErrorResponseModel)
		if string(jsonAlipayMerchantOrderSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMerchantOrderSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMerchantOrderSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMerchantOrderSyncErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMerchantOrderSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMerchantOrderSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMerchantOrderSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMerchantOrderSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMerchantOrderSyncDefaultResponse struct {
	value *AlipayMerchantOrderSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayMerchantOrderSyncDefaultResponse) Get() *AlipayMerchantOrderSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayMerchantOrderSyncDefaultResponse) Set(val *AlipayMerchantOrderSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantOrderSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantOrderSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantOrderSyncDefaultResponse(val *AlipayMerchantOrderSyncDefaultResponse) *NullableAlipayMerchantOrderSyncDefaultResponse {
	return &NullableAlipayMerchantOrderSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMerchantOrderSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantOrderSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


