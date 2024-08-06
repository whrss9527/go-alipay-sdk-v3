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


// AlipayMarketingCardTemplateCreateDefaultResponse struct for AlipayMarketingCardTemplateCreateDefaultResponse
type AlipayMarketingCardTemplateCreateDefaultResponse struct {
	AlipayMarketingCardTemplateCreateErrorResponseModel *AlipayMarketingCardTemplateCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardTemplateCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardTemplateCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardTemplateCreateErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingCardTemplateCreateErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardTemplateCreateErrorResponseModel)
		if string(jsonAlipayMarketingCardTemplateCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardTemplateCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardTemplateCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardTemplateCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardTemplateCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardTemplateCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardTemplateCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardTemplateCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingCardTemplateCreateDefaultResponse struct {
	value *AlipayMarketingCardTemplateCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardTemplateCreateDefaultResponse) Get() *AlipayMarketingCardTemplateCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardTemplateCreateDefaultResponse) Set(val *AlipayMarketingCardTemplateCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardTemplateCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardTemplateCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardTemplateCreateDefaultResponse(val *AlipayMarketingCardTemplateCreateDefaultResponse) *NullableAlipayMarketingCardTemplateCreateDefaultResponse {
	return &NullableAlipayMarketingCardTemplateCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardTemplateCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardTemplateCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


