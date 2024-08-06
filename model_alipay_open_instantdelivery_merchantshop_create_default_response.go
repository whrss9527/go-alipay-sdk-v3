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

// AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse struct for AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse
type AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse struct {
	AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel *AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel
	CommonErrorType                                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel)
		if string(jsonAlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenInstantdeliveryMerchantshopCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse struct {
	value *AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) Get() *AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) Set(val *AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse(val *AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) *NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse {
	return &NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
