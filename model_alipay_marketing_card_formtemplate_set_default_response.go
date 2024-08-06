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

// AlipayMarketingCardFormtemplateSetDefaultResponse struct for AlipayMarketingCardFormtemplateSetDefaultResponse
type AlipayMarketingCardFormtemplateSetDefaultResponse struct {
	AlipayMarketingCardFormtemplateSetErrorResponseModel *AlipayMarketingCardFormtemplateSetErrorResponseModel
	CommonErrorType                                      *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardFormtemplateSetDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardFormtemplateSetErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardFormtemplateSetErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCardFormtemplateSetErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardFormtemplateSetErrorResponseModel)
		if string(jsonAlipayMarketingCardFormtemplateSetErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardFormtemplateSetErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardFormtemplateSetErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardFormtemplateSetErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardFormtemplateSetDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardFormtemplateSetDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardFormtemplateSetErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardFormtemplateSetErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCardFormtemplateSetDefaultResponse struct {
	value *AlipayMarketingCardFormtemplateSetDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardFormtemplateSetDefaultResponse) Get() *AlipayMarketingCardFormtemplateSetDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardFormtemplateSetDefaultResponse) Set(val *AlipayMarketingCardFormtemplateSetDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardFormtemplateSetDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardFormtemplateSetDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardFormtemplateSetDefaultResponse(val *AlipayMarketingCardFormtemplateSetDefaultResponse) *NullableAlipayMarketingCardFormtemplateSetDefaultResponse {
	return &NullableAlipayMarketingCardFormtemplateSetDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardFormtemplateSetDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardFormtemplateSetDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
