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


// AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse struct for AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse
type AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse struct {
	AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel *AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel)
		if string(jsonAlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmSwLibraryBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse struct {
	value *AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) Get() *AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) Set(val *AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse(val *AlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) *NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse {
	return &NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwLibraryBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


