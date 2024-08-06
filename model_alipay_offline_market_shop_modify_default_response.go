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

// AlipayOfflineMarketShopModifyDefaultResponse struct for AlipayOfflineMarketShopModifyDefaultResponse
type AlipayOfflineMarketShopModifyDefaultResponse struct {
	AlipayOfflineMarketShopModifyErrorResponseModel *AlipayOfflineMarketShopModifyErrorResponseModel
	CommonErrorType                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOfflineMarketShopModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOfflineMarketShopModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOfflineMarketShopModifyErrorResponseModel)
	if err == nil {
		jsonAlipayOfflineMarketShopModifyErrorResponseModel, _ := json.Marshal(dst.AlipayOfflineMarketShopModifyErrorResponseModel)
		if string(jsonAlipayOfflineMarketShopModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOfflineMarketShopModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOfflineMarketShopModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOfflineMarketShopModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOfflineMarketShopModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOfflineMarketShopModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOfflineMarketShopModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOfflineMarketShopModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOfflineMarketShopModifyDefaultResponse struct {
	value *AlipayOfflineMarketShopModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayOfflineMarketShopModifyDefaultResponse) Get() *AlipayOfflineMarketShopModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOfflineMarketShopModifyDefaultResponse) Set(val *AlipayOfflineMarketShopModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOfflineMarketShopModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOfflineMarketShopModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOfflineMarketShopModifyDefaultResponse(val *AlipayOfflineMarketShopModifyDefaultResponse) *NullableAlipayOfflineMarketShopModifyDefaultResponse {
	return &NullableAlipayOfflineMarketShopModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOfflineMarketShopModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOfflineMarketShopModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
