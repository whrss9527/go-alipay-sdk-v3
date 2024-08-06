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

// DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse struct for DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse
type DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse struct {
	CommonErrorType                                                 *CommonErrorType
	DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel *DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel
	err = json.Unmarshal(data, &dst.DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel)
	if err == nil {
		jsonDatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel, _ := json.Marshal(dst.DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel)
		if string(jsonDatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel) == "{}" { // empty struct
			dst.DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel = nil
		} else {
			return nil // data stored in dst.DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel, return on the first match
		}
	} else {
		dst.DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel != nil {
		return json.Marshal(&src.DatadigitalFincloudGeneralsaasOcrServerDetectErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse struct {
	value *DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) Get() *DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) Set(val *DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse(val *DatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) *NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse {
	return &NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrServerDetectDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
