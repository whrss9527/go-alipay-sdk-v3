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

// AlipayMarketingCardBenefitModifyDefaultResponse struct for AlipayMarketingCardBenefitModifyDefaultResponse
type AlipayMarketingCardBenefitModifyDefaultResponse struct {
	AlipayMarketingCardBenefitModifyErrorResponseModel *AlipayMarketingCardBenefitModifyErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardBenefitModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardBenefitModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardBenefitModifyErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCardBenefitModifyErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardBenefitModifyErrorResponseModel)
		if string(jsonAlipayMarketingCardBenefitModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardBenefitModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardBenefitModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardBenefitModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardBenefitModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardBenefitModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardBenefitModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardBenefitModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCardBenefitModifyDefaultResponse struct {
	value *AlipayMarketingCardBenefitModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardBenefitModifyDefaultResponse) Get() *AlipayMarketingCardBenefitModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardBenefitModifyDefaultResponse) Set(val *AlipayMarketingCardBenefitModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardBenefitModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardBenefitModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardBenefitModifyDefaultResponse(val *AlipayMarketingCardBenefitModifyDefaultResponse) *NullableAlipayMarketingCardBenefitModifyDefaultResponse {
	return &NullableAlipayMarketingCardBenefitModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardBenefitModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardBenefitModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
