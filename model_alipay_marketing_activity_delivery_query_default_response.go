/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
	"fmt"
)

// AlipayMarketingActivityDeliveryQueryDefaultResponse struct for AlipayMarketingActivityDeliveryQueryDefaultResponse
type AlipayMarketingActivityDeliveryQueryDefaultResponse struct {
	AlipayMarketingActivityDeliveryQueryErrorResponseModel *AlipayMarketingActivityDeliveryQueryErrorResponseModel
	CommonErrorType                                        *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityDeliveryQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityDeliveryQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityDeliveryQueryErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingActivityDeliveryQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityDeliveryQueryErrorResponseModel)
		if string(jsonAlipayMarketingActivityDeliveryQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityDeliveryQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityDeliveryQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityDeliveryQueryErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType)
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityDeliveryQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityDeliveryQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityDeliveryQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityDeliveryQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingActivityDeliveryQueryDefaultResponse struct {
	value *AlipayMarketingActivityDeliveryQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityDeliveryQueryDefaultResponse) Get() *AlipayMarketingActivityDeliveryQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityDeliveryQueryDefaultResponse) Set(val *AlipayMarketingActivityDeliveryQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityDeliveryQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityDeliveryQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityDeliveryQueryDefaultResponse(val *AlipayMarketingActivityDeliveryQueryDefaultResponse) *NullableAlipayMarketingActivityDeliveryQueryDefaultResponse {
	return &NullableAlipayMarketingActivityDeliveryQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityDeliveryQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityDeliveryQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
