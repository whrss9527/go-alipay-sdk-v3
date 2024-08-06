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


// ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse struct for ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse
type ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse struct {
	CommonErrorType *CommonErrorType
	ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel *ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel
	err = json.Unmarshal(data, &dst.ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel);
	if err == nil {
		jsonZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel, _ := json.Marshal(dst.ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel)
		if string(jsonZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel) == "{}" { // empty struct
			dst.ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel, return on the first match
		}
	} else {
		dst.ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel != nil {
		return json.Marshal(&src.ZhimaCreditPeZmgoSettleUnfreezeErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse struct {
	value *ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse
	isSet bool
}

func (v NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) Get() *ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) Set(val *ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse(val *ZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) *NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse {
	return &NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoSettleUnfreezeDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


