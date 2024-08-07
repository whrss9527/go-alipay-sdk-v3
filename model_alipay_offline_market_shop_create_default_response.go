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

// AlipayOfflineMarketShopCreateDefaultResponse struct for AlipayOfflineMarketShopCreateDefaultResponse
type AlipayOfflineMarketShopCreateDefaultResponse struct {
	AlipayOfflineMarketShopCreateErrorResponseModel *AlipayOfflineMarketShopCreateErrorResponseModel
	CommonErrorType                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOfflineMarketShopCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOfflineMarketShopCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOfflineMarketShopCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOfflineMarketShopCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOfflineMarketShopCreateErrorResponseModel)
		if string(jsonAlipayOfflineMarketShopCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOfflineMarketShopCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOfflineMarketShopCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOfflineMarketShopCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOfflineMarketShopCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOfflineMarketShopCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOfflineMarketShopCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOfflineMarketShopCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOfflineMarketShopCreateDefaultResponse struct {
	value *AlipayOfflineMarketShopCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOfflineMarketShopCreateDefaultResponse) Get() *AlipayOfflineMarketShopCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOfflineMarketShopCreateDefaultResponse) Set(val *AlipayOfflineMarketShopCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOfflineMarketShopCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOfflineMarketShopCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOfflineMarketShopCreateDefaultResponse(val *AlipayOfflineMarketShopCreateDefaultResponse) *NullableAlipayOfflineMarketShopCreateDefaultResponse {
	return &NullableAlipayOfflineMarketShopCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOfflineMarketShopCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOfflineMarketShopCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
