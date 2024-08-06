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

// AlipayMarketingCampaignCashCreateDefaultResponse struct for AlipayMarketingCampaignCashCreateDefaultResponse
type AlipayMarketingCampaignCashCreateDefaultResponse struct {
	AlipayMarketingCampaignCashCreateErrorResponseModel *AlipayMarketingCampaignCashCreateErrorResponseModel
	CommonErrorType                                     *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCampaignCashCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCampaignCashCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCampaignCashCreateErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCampaignCashCreateErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCampaignCashCreateErrorResponseModel)
		if string(jsonAlipayMarketingCampaignCashCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCampaignCashCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCampaignCashCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCampaignCashCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCampaignCashCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCampaignCashCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCampaignCashCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCampaignCashCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCampaignCashCreateDefaultResponse struct {
	value *AlipayMarketingCampaignCashCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCampaignCashCreateDefaultResponse) Get() *AlipayMarketingCampaignCashCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCampaignCashCreateDefaultResponse) Set(val *AlipayMarketingCampaignCashCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCampaignCashCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCampaignCashCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCampaignCashCreateDefaultResponse(val *AlipayMarketingCampaignCashCreateDefaultResponse) *NullableAlipayMarketingCampaignCashCreateDefaultResponse {
	return &NullableAlipayMarketingCampaignCashCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCampaignCashCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCampaignCashCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
