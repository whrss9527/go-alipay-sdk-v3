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


// AlipayOpenPublicUserDataBatchqueryDefaultResponse struct for AlipayOpenPublicUserDataBatchqueryDefaultResponse
type AlipayOpenPublicUserDataBatchqueryDefaultResponse struct {
	AlipayOpenPublicUserDataBatchqueryErrorResponseModel *AlipayOpenPublicUserDataBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicUserDataBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicUserDataBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicUserDataBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicUserDataBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicUserDataBatchqueryErrorResponseModel)
		if string(jsonAlipayOpenPublicUserDataBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicUserDataBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicUserDataBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicUserDataBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicUserDataBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicUserDataBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicUserDataBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicUserDataBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse struct {
	value *AlipayOpenPublicUserDataBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse) Get() *AlipayOpenPublicUserDataBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse) Set(val *AlipayOpenPublicUserDataBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicUserDataBatchqueryDefaultResponse(val *AlipayOpenPublicUserDataBatchqueryDefaultResponse) *NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse {
	return &NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


