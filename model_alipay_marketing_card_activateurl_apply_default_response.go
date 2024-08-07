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

// AlipayMarketingCardActivateurlApplyDefaultResponse struct for AlipayMarketingCardActivateurlApplyDefaultResponse
type AlipayMarketingCardActivateurlApplyDefaultResponse struct {
	AlipayMarketingCardActivateurlApplyErrorResponseModel *AlipayMarketingCardActivateurlApplyErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardActivateurlApplyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardActivateurlApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardActivateurlApplyErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCardActivateurlApplyErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardActivateurlApplyErrorResponseModel)
		if string(jsonAlipayMarketingCardActivateurlApplyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardActivateurlApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardActivateurlApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardActivateurlApplyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardActivateurlApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardActivateurlApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardActivateurlApplyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardActivateurlApplyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCardActivateurlApplyDefaultResponse struct {
	value *AlipayMarketingCardActivateurlApplyDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardActivateurlApplyDefaultResponse) Get() *AlipayMarketingCardActivateurlApplyDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardActivateurlApplyDefaultResponse) Set(val *AlipayMarketingCardActivateurlApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardActivateurlApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardActivateurlApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardActivateurlApplyDefaultResponse(val *AlipayMarketingCardActivateurlApplyDefaultResponse) *NullableAlipayMarketingCardActivateurlApplyDefaultResponse {
	return &NullableAlipayMarketingCardActivateurlApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardActivateurlApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardActivateurlApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
