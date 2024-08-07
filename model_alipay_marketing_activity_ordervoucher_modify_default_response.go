/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
	"fmt"
)

// AlipayMarketingActivityOrdervoucherModifyDefaultResponse struct for AlipayMarketingActivityOrdervoucherModifyDefaultResponse
type AlipayMarketingActivityOrdervoucherModifyDefaultResponse struct {
	AlipayMarketingActivityOrdervoucherModifyErrorResponseModel *AlipayMarketingActivityOrdervoucherModifyErrorResponseModel
	CommonErrorType                                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityOrdervoucherModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityOrdervoucherModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityOrdervoucherModifyErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingActivityOrdervoucherModifyErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityOrdervoucherModifyErrorResponseModel)
		if string(jsonAlipayMarketingActivityOrdervoucherModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityOrdervoucherModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityOrdervoucherModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityOrdervoucherModifyErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType)
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityOrdervoucherModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityOrdervoucherModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityOrdervoucherModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityOrdervoucherModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse struct {
	value *AlipayMarketingActivityOrdervoucherModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse) Get() *AlipayMarketingActivityOrdervoucherModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse) Set(val *AlipayMarketingActivityOrdervoucherModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse(val *AlipayMarketingActivityOrdervoucherModifyDefaultResponse) *NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse {
	return &NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
