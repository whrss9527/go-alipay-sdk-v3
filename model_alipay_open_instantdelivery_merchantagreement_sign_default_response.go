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

// AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse struct for AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse
type AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse struct {
	AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel *AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel
	CommonErrorType                                                  *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel)
	if err == nil {
		jsonAlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel, _ := json.Marshal(dst.AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel)
		if string(jsonAlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenInstantdeliveryMerchantagreementSignErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse struct {
	value *AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) Get() *AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) Set(val *AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse(val *AlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) *NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse {
	return &NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryMerchantagreementSignDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
