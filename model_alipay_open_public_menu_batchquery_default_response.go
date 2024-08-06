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


// AlipayOpenPublicMenuBatchqueryDefaultResponse struct for AlipayOpenPublicMenuBatchqueryDefaultResponse
type AlipayOpenPublicMenuBatchqueryDefaultResponse struct {
	AlipayOpenPublicMenuBatchqueryErrorResponseModel *AlipayOpenPublicMenuBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicMenuBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicMenuBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicMenuBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicMenuBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicMenuBatchqueryErrorResponseModel)
		if string(jsonAlipayOpenPublicMenuBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicMenuBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicMenuBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicMenuBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicMenuBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicMenuBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicMenuBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicMenuBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicMenuBatchqueryDefaultResponse struct {
	value *AlipayOpenPublicMenuBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicMenuBatchqueryDefaultResponse) Get() *AlipayOpenPublicMenuBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicMenuBatchqueryDefaultResponse) Set(val *AlipayOpenPublicMenuBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMenuBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMenuBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMenuBatchqueryDefaultResponse(val *AlipayOpenPublicMenuBatchqueryDefaultResponse) *NullableAlipayOpenPublicMenuBatchqueryDefaultResponse {
	return &NullableAlipayOpenPublicMenuBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMenuBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMenuBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


