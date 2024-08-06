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


// AlipayMarketingCardBenefitDeleteDefaultResponse struct for AlipayMarketingCardBenefitDeleteDefaultResponse
type AlipayMarketingCardBenefitDeleteDefaultResponse struct {
	AlipayMarketingCardBenefitDeleteErrorResponseModel *AlipayMarketingCardBenefitDeleteErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardBenefitDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardBenefitDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardBenefitDeleteErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingCardBenefitDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardBenefitDeleteErrorResponseModel)
		if string(jsonAlipayMarketingCardBenefitDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardBenefitDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardBenefitDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardBenefitDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardBenefitDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardBenefitDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardBenefitDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardBenefitDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingCardBenefitDeleteDefaultResponse struct {
	value *AlipayMarketingCardBenefitDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardBenefitDeleteDefaultResponse) Get() *AlipayMarketingCardBenefitDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardBenefitDeleteDefaultResponse) Set(val *AlipayMarketingCardBenefitDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardBenefitDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardBenefitDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardBenefitDeleteDefaultResponse(val *AlipayMarketingCardBenefitDeleteDefaultResponse) *NullableAlipayMarketingCardBenefitDeleteDefaultResponse {
	return &NullableAlipayMarketingCardBenefitDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardBenefitDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardBenefitDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


