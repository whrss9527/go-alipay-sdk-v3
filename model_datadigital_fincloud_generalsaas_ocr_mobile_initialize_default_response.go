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

// DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse struct for DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse
type DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse struct {
	CommonErrorType                                                     *CommonErrorType
	DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel *DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel
	err = json.Unmarshal(data, &dst.DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel)
	if err == nil {
		jsonDatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel, _ := json.Marshal(dst.DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel)
		if string(jsonDatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel) == "{}" { // empty struct
			dst.DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel = nil
		} else {
			return nil // data stored in dst.DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel, return on the first match
		}
	} else {
		dst.DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel != nil {
		return json.Marshal(&src.DatadigitalFincloudGeneralsaasOcrMobileInitializeErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse struct {
	value *DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) Get() *DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) Set(val *DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse(val *DatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) *NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse {
	return &NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasOcrMobileInitializeDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
