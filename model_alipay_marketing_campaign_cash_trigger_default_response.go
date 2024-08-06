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


// AlipayMarketingCampaignCashTriggerDefaultResponse struct for AlipayMarketingCampaignCashTriggerDefaultResponse
type AlipayMarketingCampaignCashTriggerDefaultResponse struct {
	AlipayMarketingCampaignCashTriggerErrorResponseModel *AlipayMarketingCampaignCashTriggerErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCampaignCashTriggerDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCampaignCashTriggerErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCampaignCashTriggerErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingCampaignCashTriggerErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCampaignCashTriggerErrorResponseModel)
		if string(jsonAlipayMarketingCampaignCashTriggerErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCampaignCashTriggerErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCampaignCashTriggerErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCampaignCashTriggerErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCampaignCashTriggerDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCampaignCashTriggerDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCampaignCashTriggerErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCampaignCashTriggerErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingCampaignCashTriggerDefaultResponse struct {
	value *AlipayMarketingCampaignCashTriggerDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCampaignCashTriggerDefaultResponse) Get() *AlipayMarketingCampaignCashTriggerDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCampaignCashTriggerDefaultResponse) Set(val *AlipayMarketingCampaignCashTriggerDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCampaignCashTriggerDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCampaignCashTriggerDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCampaignCashTriggerDefaultResponse(val *AlipayMarketingCampaignCashTriggerDefaultResponse) *NullableAlipayMarketingCampaignCashTriggerDefaultResponse {
	return &NullableAlipayMarketingCampaignCashTriggerDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCampaignCashTriggerDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCampaignCashTriggerDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


