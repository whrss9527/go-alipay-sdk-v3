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


// AlipayOpenPublicMenuModifyDefaultResponse struct for AlipayOpenPublicMenuModifyDefaultResponse
type AlipayOpenPublicMenuModifyDefaultResponse struct {
	AlipayOpenPublicMenuModifyErrorResponseModel *AlipayOpenPublicMenuModifyErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicMenuModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicMenuModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicMenuModifyErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicMenuModifyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicMenuModifyErrorResponseModel)
		if string(jsonAlipayOpenPublicMenuModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicMenuModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicMenuModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicMenuModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicMenuModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicMenuModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicMenuModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicMenuModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicMenuModifyDefaultResponse struct {
	value *AlipayOpenPublicMenuModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicMenuModifyDefaultResponse) Get() *AlipayOpenPublicMenuModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicMenuModifyDefaultResponse) Set(val *AlipayOpenPublicMenuModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMenuModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMenuModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMenuModifyDefaultResponse(val *AlipayOpenPublicMenuModifyDefaultResponse) *NullableAlipayOpenPublicMenuModifyDefaultResponse {
	return &NullableAlipayOpenPublicMenuModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMenuModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMenuModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

