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


// AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse struct for AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse
type AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse struct {
	AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel *AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel)
		if string(jsonAlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipaySocialBaseContentlibStandardcontentBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse struct {
	value *AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) Get() *AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) Set(val *AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse(val *AlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) *NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse {
	return &NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipaySocialBaseContentlibStandardcontentBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

