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


// AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse struct for AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse
type AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse struct {
	AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel *AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel);
	if err == nil {
		jsonAlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel)
		if string(jsonAlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceLogisticsWaybillIstddetailQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse struct {
	value *AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) Get() *AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) Set(val *AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse(val *AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) *NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse {
	return &NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


