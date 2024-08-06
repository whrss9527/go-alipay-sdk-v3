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


// AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse struct for AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse
type AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse struct {
	AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel *AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel)
		if string(jsonAlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcJointaccountbillDetailBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse struct {
	value *AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) Get() *AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) Set(val *AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse(val *AlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) *NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse {
	return &NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcJointaccountbillDetailBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


