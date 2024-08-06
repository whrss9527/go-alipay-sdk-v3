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


// AlipayOpenPublicMatchuserLabelCreateDefaultResponse struct for AlipayOpenPublicMatchuserLabelCreateDefaultResponse
type AlipayOpenPublicMatchuserLabelCreateDefaultResponse struct {
	AlipayOpenPublicMatchuserLabelCreateErrorResponseModel *AlipayOpenPublicMatchuserLabelCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicMatchuserLabelCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicMatchuserLabelCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicMatchuserLabelCreateErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicMatchuserLabelCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicMatchuserLabelCreateErrorResponseModel)
		if string(jsonAlipayOpenPublicMatchuserLabelCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicMatchuserLabelCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicMatchuserLabelCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicMatchuserLabelCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicMatchuserLabelCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicMatchuserLabelCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicMatchuserLabelCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicMatchuserLabelCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse struct {
	value *AlipayOpenPublicMatchuserLabelCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse) Get() *AlipayOpenPublicMatchuserLabelCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse) Set(val *AlipayOpenPublicMatchuserLabelCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse(val *AlipayOpenPublicMatchuserLabelCreateDefaultResponse) *NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse {
	return &NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMatchuserLabelCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


