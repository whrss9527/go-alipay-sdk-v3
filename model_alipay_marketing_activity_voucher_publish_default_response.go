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


// AlipayMarketingActivityVoucherPublishDefaultResponse struct for AlipayMarketingActivityVoucherPublishDefaultResponse
type AlipayMarketingActivityVoucherPublishDefaultResponse struct {
	AlipayMarketingActivityVoucherPublishErrorResponseModel *AlipayMarketingActivityVoucherPublishErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityVoucherPublishDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityVoucherPublishErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityVoucherPublishErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityVoucherPublishErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityVoucherPublishErrorResponseModel)
		if string(jsonAlipayMarketingActivityVoucherPublishErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityVoucherPublishErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityVoucherPublishErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityVoucherPublishErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityVoucherPublishDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityVoucherPublishDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityVoucherPublishErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityVoucherPublishErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityVoucherPublishDefaultResponse struct {
	value *AlipayMarketingActivityVoucherPublishDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherPublishDefaultResponse) Get() *AlipayMarketingActivityVoucherPublishDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherPublishDefaultResponse) Set(val *AlipayMarketingActivityVoucherPublishDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherPublishDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherPublishDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherPublishDefaultResponse(val *AlipayMarketingActivityVoucherPublishDefaultResponse) *NullableAlipayMarketingActivityVoucherPublishDefaultResponse {
	return &NullableAlipayMarketingActivityVoucherPublishDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherPublishDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherPublishDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

