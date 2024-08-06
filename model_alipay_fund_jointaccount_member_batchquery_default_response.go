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

// AlipayFundJointaccountMemberBatchqueryDefaultResponse struct for AlipayFundJointaccountMemberBatchqueryDefaultResponse
type AlipayFundJointaccountMemberBatchqueryDefaultResponse struct {
	AlipayFundJointaccountMemberBatchqueryErrorResponseModel *AlipayFundJointaccountMemberBatchqueryErrorResponseModel
	CommonErrorType                                          *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundJointaccountMemberBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundJointaccountMemberBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundJointaccountMemberBatchqueryErrorResponseModel)
	if err == nil {
		jsonAlipayFundJointaccountMemberBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundJointaccountMemberBatchqueryErrorResponseModel)
		if string(jsonAlipayFundJointaccountMemberBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundJointaccountMemberBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundJointaccountMemberBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundJointaccountMemberBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundJointaccountMemberBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundJointaccountMemberBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundJointaccountMemberBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundJointaccountMemberBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse struct {
	value *AlipayFundJointaccountMemberBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse) Get() *AlipayFundJointaccountMemberBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse) Set(val *AlipayFundJointaccountMemberBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountMemberBatchqueryDefaultResponse(val *AlipayFundJointaccountMemberBatchqueryDefaultResponse) *NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse {
	return &NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountMemberBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
