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


// AlipayMarketingActivityShopBatchqueryDefaultResponse struct for AlipayMarketingActivityShopBatchqueryDefaultResponse
type AlipayMarketingActivityShopBatchqueryDefaultResponse struct {
	AlipayMarketingActivityShopBatchqueryErrorResponseModel *AlipayMarketingActivityShopBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityShopBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityShopBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityShopBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityShopBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityShopBatchqueryErrorResponseModel)
		if string(jsonAlipayMarketingActivityShopBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityShopBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityShopBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityShopBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityShopBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityShopBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityShopBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityShopBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityShopBatchqueryDefaultResponse struct {
	value *AlipayMarketingActivityShopBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityShopBatchqueryDefaultResponse) Get() *AlipayMarketingActivityShopBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityShopBatchqueryDefaultResponse) Set(val *AlipayMarketingActivityShopBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityShopBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityShopBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityShopBatchqueryDefaultResponse(val *AlipayMarketingActivityShopBatchqueryDefaultResponse) *NullableAlipayMarketingActivityShopBatchqueryDefaultResponse {
	return &NullableAlipayMarketingActivityShopBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityShopBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityShopBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


