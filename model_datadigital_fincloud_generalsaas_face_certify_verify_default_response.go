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


// DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse struct for DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse
type DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse struct {
	CommonErrorType *CommonErrorType
	DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel *DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel
	err = json.Unmarshal(data, &dst.DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel);
	if err == nil {
		jsonDatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel, _ := json.Marshal(dst.DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel)
		if string(jsonDatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel) == "{}" { // empty struct
			dst.DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel, return on the first match
		}
	} else {
		dst.DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel != nil {
		return json.Marshal(&src.DatadigitalFincloudGeneralsaasFaceCertifyVerifyErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse struct {
	value *DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) Get() *DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) Set(val *DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse(val *DatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) *NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse {
	return &NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCertifyVerifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


