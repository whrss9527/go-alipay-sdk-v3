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

// DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse struct for DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse
type DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse struct {
	CommonErrorType                                                       *CommonErrorType
	DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel *DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel
	err = json.Unmarshal(data, &dst.DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel)
	if err == nil {
		jsonDatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel, _ := json.Marshal(dst.DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel)
		if string(jsonDatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel) == "{}" { // empty struct
			dst.DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel = nil
		} else {
			return nil // data stored in dst.DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel, return on the first match
		}
	} else {
		dst.DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel != nil {
		return json.Marshal(&src.DatadigitalFincloudGeneralsaasFaceCertifyInitializeErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse struct {
	value *DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) Get() *DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) Set(val *DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse(val *DatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) *NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse {
	return &NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCertifyInitializeDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
