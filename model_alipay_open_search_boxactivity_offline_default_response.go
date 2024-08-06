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


// AlipayOpenSearchBoxactivityOfflineDefaultResponse struct for AlipayOpenSearchBoxactivityOfflineDefaultResponse
type AlipayOpenSearchBoxactivityOfflineDefaultResponse struct {
	AlipayOpenSearchBoxactivityOfflineErrorResponseModel *AlipayOpenSearchBoxactivityOfflineErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchBoxactivityOfflineDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchBoxactivityOfflineErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchBoxactivityOfflineErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSearchBoxactivityOfflineErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchBoxactivityOfflineErrorResponseModel)
		if string(jsonAlipayOpenSearchBoxactivityOfflineErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchBoxactivityOfflineErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchBoxactivityOfflineErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchBoxactivityOfflineErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchBoxactivityOfflineDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchBoxactivityOfflineDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchBoxactivityOfflineErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchBoxactivityOfflineErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse struct {
	value *AlipayOpenSearchBoxactivityOfflineDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse) Get() *AlipayOpenSearchBoxactivityOfflineDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse) Set(val *AlipayOpenSearchBoxactivityOfflineDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxactivityOfflineDefaultResponse(val *AlipayOpenSearchBoxactivityOfflineDefaultResponse) *NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse {
	return &NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxactivityOfflineDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

