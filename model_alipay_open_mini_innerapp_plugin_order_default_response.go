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


// AlipayOpenMiniInnerappPluginOrderDefaultResponse struct for AlipayOpenMiniInnerappPluginOrderDefaultResponse
type AlipayOpenMiniInnerappPluginOrderDefaultResponse struct {
	AlipayOpenMiniInnerappPluginOrderErrorResponseModel *AlipayOpenMiniInnerappPluginOrderErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerappPluginOrderDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerappPluginOrderErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerappPluginOrderErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniInnerappPluginOrderErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerappPluginOrderErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerappPluginOrderErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerappPluginOrderErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerappPluginOrderErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerappPluginOrderErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerappPluginOrderDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerappPluginOrderDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerappPluginOrderErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerappPluginOrderErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse struct {
	value *AlipayOpenMiniInnerappPluginOrderDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse) Get() *AlipayOpenMiniInnerappPluginOrderDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse) Set(val *AlipayOpenMiniInnerappPluginOrderDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappPluginOrderDefaultResponse(val *AlipayOpenMiniInnerappPluginOrderDefaultResponse) *NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse {
	return &NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


