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


// AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse struct for AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse
type AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse struct {
	AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel *AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel)
		if string(jsonAlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniTemplatemsgMaketingstatusModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse struct {
	value *AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) Get() *AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) Set(val *AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse(val *AlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) *NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse {
	return &NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTemplatemsgMaketingstatusModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


