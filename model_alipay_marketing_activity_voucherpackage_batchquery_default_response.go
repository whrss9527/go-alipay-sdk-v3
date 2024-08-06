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


// AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse struct for AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse
type AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse struct {
	AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel *AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel)
		if string(jsonAlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityVoucherpackageBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse struct {
	value *AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) Get() *AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) Set(val *AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse(val *AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) *NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse {
	return &NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


