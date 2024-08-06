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

// AlipayCommerceEcEnterpriseAddressModifyDefaultResponse struct for AlipayCommerceEcEnterpriseAddressModifyDefaultResponse
type AlipayCommerceEcEnterpriseAddressModifyDefaultResponse struct {
	AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel *AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel
	CommonErrorType                                           *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcEnterpriseAddressModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcEnterpriseAddressModifyErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel)
		if string(jsonAlipayCommerceEcEnterpriseAddressModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcEnterpriseAddressModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcEnterpriseAddressModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcEnterpriseAddressModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse struct {
	value *AlipayCommerceEcEnterpriseAddressModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse) Get() *AlipayCommerceEcEnterpriseAddressModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse) Set(val *AlipayCommerceEcEnterpriseAddressModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse(val *AlipayCommerceEcEnterpriseAddressModifyDefaultResponse) *NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse {
	return &NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseAddressModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
