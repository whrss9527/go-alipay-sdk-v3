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


// AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse struct for AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse
type AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse struct {
	AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel *AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel)
		if string(jsonAlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityOrdervoucherDisassociateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse struct {
	value *AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) Get() *AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) Set(val *AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse(val *AlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) *NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse {
	return &NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherDisassociateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


