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


// AlipayMarketingRecruitEnrollQueryDefaultResponse struct for AlipayMarketingRecruitEnrollQueryDefaultResponse
type AlipayMarketingRecruitEnrollQueryDefaultResponse struct {
	AlipayMarketingRecruitEnrollQueryErrorResponseModel *AlipayMarketingRecruitEnrollQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingRecruitEnrollQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingRecruitEnrollQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingRecruitEnrollQueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingRecruitEnrollQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingRecruitEnrollQueryErrorResponseModel)
		if string(jsonAlipayMarketingRecruitEnrollQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingRecruitEnrollQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingRecruitEnrollQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingRecruitEnrollQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingRecruitEnrollQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingRecruitEnrollQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingRecruitEnrollQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingRecruitEnrollQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingRecruitEnrollQueryDefaultResponse struct {
	value *AlipayMarketingRecruitEnrollQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingRecruitEnrollQueryDefaultResponse) Get() *AlipayMarketingRecruitEnrollQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingRecruitEnrollQueryDefaultResponse) Set(val *AlipayMarketingRecruitEnrollQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingRecruitEnrollQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingRecruitEnrollQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingRecruitEnrollQueryDefaultResponse(val *AlipayMarketingRecruitEnrollQueryDefaultResponse) *NullableAlipayMarketingRecruitEnrollQueryDefaultResponse {
	return &NullableAlipayMarketingRecruitEnrollQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingRecruitEnrollQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingRecruitEnrollQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


