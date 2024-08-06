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


// AlipayOpenMiniTemplateMarketingCreateDefaultResponse struct for AlipayOpenMiniTemplateMarketingCreateDefaultResponse
type AlipayOpenMiniTemplateMarketingCreateDefaultResponse struct {
	AlipayOpenMiniTemplateMarketingCreateErrorResponseModel *AlipayOpenMiniTemplateMarketingCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniTemplateMarketingCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniTemplateMarketingCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniTemplateMarketingCreateErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniTemplateMarketingCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniTemplateMarketingCreateErrorResponseModel)
		if string(jsonAlipayOpenMiniTemplateMarketingCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniTemplateMarketingCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniTemplateMarketingCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniTemplateMarketingCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniTemplateMarketingCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniTemplateMarketingCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniTemplateMarketingCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniTemplateMarketingCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse struct {
	value *AlipayOpenMiniTemplateMarketingCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse) Get() *AlipayOpenMiniTemplateMarketingCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse) Set(val *AlipayOpenMiniTemplateMarketingCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse(val *AlipayOpenMiniTemplateMarketingCreateDefaultResponse) *NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse {
	return &NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTemplateMarketingCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


