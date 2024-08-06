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


// AlipayOpenPublicTopicCreateDefaultResponse struct for AlipayOpenPublicTopicCreateDefaultResponse
type AlipayOpenPublicTopicCreateDefaultResponse struct {
	AlipayOpenPublicTopicCreateErrorResponseModel *AlipayOpenPublicTopicCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicTopicCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicTopicCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicTopicCreateErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicTopicCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicTopicCreateErrorResponseModel)
		if string(jsonAlipayOpenPublicTopicCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicTopicCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicTopicCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicTopicCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicTopicCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicTopicCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicTopicCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicTopicCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicTopicCreateDefaultResponse struct {
	value *AlipayOpenPublicTopicCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicTopicCreateDefaultResponse) Get() *AlipayOpenPublicTopicCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicTopicCreateDefaultResponse) Set(val *AlipayOpenPublicTopicCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicTopicCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicTopicCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicTopicCreateDefaultResponse(val *AlipayOpenPublicTopicCreateDefaultResponse) *NullableAlipayOpenPublicTopicCreateDefaultResponse {
	return &NullableAlipayOpenPublicTopicCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicTopicCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicTopicCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


