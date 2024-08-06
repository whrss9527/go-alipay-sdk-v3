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


// AlipayMarketingCardActivateformQueryDefaultResponse struct for AlipayMarketingCardActivateformQueryDefaultResponse
type AlipayMarketingCardActivateformQueryDefaultResponse struct {
	AlipayMarketingCardActivateformQueryErrorResponseModel *AlipayMarketingCardActivateformQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardActivateformQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardActivateformQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardActivateformQueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingCardActivateformQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardActivateformQueryErrorResponseModel)
		if string(jsonAlipayMarketingCardActivateformQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardActivateformQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardActivateformQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardActivateformQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardActivateformQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardActivateformQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardActivateformQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardActivateformQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingCardActivateformQueryDefaultResponse struct {
	value *AlipayMarketingCardActivateformQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardActivateformQueryDefaultResponse) Get() *AlipayMarketingCardActivateformQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardActivateformQueryDefaultResponse) Set(val *AlipayMarketingCardActivateformQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardActivateformQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardActivateformQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardActivateformQueryDefaultResponse(val *AlipayMarketingCardActivateformQueryDefaultResponse) *NullableAlipayMarketingCardActivateformQueryDefaultResponse {
	return &NullableAlipayMarketingCardActivateformQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardActivateformQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardActivateformQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

