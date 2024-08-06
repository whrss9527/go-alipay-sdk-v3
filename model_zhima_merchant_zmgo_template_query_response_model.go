/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ZhimaMerchantZmgoTemplateQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaMerchantZmgoTemplateQueryResponseModel{}

// ZhimaMerchantZmgoTemplateQueryResponseModel struct for ZhimaMerchantZmgoTemplateQueryResponseModel
type ZhimaMerchantZmgoTemplateQueryResponseModel struct {
	BasicConfig *ZMGOBasicConfig `json:"basic_config,omitempty"`
	ExtConfig *ZMGOExtConfig `json:"ext_config,omitempty"`
	ObligationConfig *ZMGOObligationConfig `json:"obligation_config,omitempty"`
	OpenConfig *ZMGOOpenConfig `json:"open_config,omitempty"`
	QuitConfig *ZMGOQuitConfig `json:"quit_config,omitempty"`
	RightConfig *ZMGORightConfig `json:"right_config,omitempty"`
	SettlementConfig *ZMGOSettlementConfig `json:"settlement_config,omitempty"`
}

// NewZhimaMerchantZmgoTemplateQueryResponseModel instantiates a new ZhimaMerchantZmgoTemplateQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaMerchantZmgoTemplateQueryResponseModel() *ZhimaMerchantZmgoTemplateQueryResponseModel {
	this := ZhimaMerchantZmgoTemplateQueryResponseModel{}
	return &this
}

// NewZhimaMerchantZmgoTemplateQueryResponseModelWithDefaults instantiates a new ZhimaMerchantZmgoTemplateQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaMerchantZmgoTemplateQueryResponseModelWithDefaults() *ZhimaMerchantZmgoTemplateQueryResponseModel {
	this := ZhimaMerchantZmgoTemplateQueryResponseModel{}
	return &this
}

// GetBasicConfig returns the BasicConfig field value if set, zero value otherwise.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetBasicConfig() ZMGOBasicConfig {
	if o == nil || IsNil(o.BasicConfig) {
		var ret ZMGOBasicConfig
		return ret
	}
	return *o.BasicConfig
}

// GetBasicConfigOk returns a tuple with the BasicConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetBasicConfigOk() (*ZMGOBasicConfig, bool) {
	if o == nil || IsNil(o.BasicConfig) {
		return nil, false
	}
	return o.BasicConfig, true
}

// HasBasicConfig returns a boolean if a field has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) HasBasicConfig() bool {
	if o != nil && !IsNil(o.BasicConfig) {
		return true
	}

	return false
}

// SetBasicConfig gets a reference to the given ZMGOBasicConfig and assigns it to the BasicConfig field.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) SetBasicConfig(v ZMGOBasicConfig) {
	o.BasicConfig = &v
}

// GetExtConfig returns the ExtConfig field value if set, zero value otherwise.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetExtConfig() ZMGOExtConfig {
	if o == nil || IsNil(o.ExtConfig) {
		var ret ZMGOExtConfig
		return ret
	}
	return *o.ExtConfig
}

// GetExtConfigOk returns a tuple with the ExtConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetExtConfigOk() (*ZMGOExtConfig, bool) {
	if o == nil || IsNil(o.ExtConfig) {
		return nil, false
	}
	return o.ExtConfig, true
}

// HasExtConfig returns a boolean if a field has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) HasExtConfig() bool {
	if o != nil && !IsNil(o.ExtConfig) {
		return true
	}

	return false
}

// SetExtConfig gets a reference to the given ZMGOExtConfig and assigns it to the ExtConfig field.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) SetExtConfig(v ZMGOExtConfig) {
	o.ExtConfig = &v
}

// GetObligationConfig returns the ObligationConfig field value if set, zero value otherwise.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetObligationConfig() ZMGOObligationConfig {
	if o == nil || IsNil(o.ObligationConfig) {
		var ret ZMGOObligationConfig
		return ret
	}
	return *o.ObligationConfig
}

// GetObligationConfigOk returns a tuple with the ObligationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetObligationConfigOk() (*ZMGOObligationConfig, bool) {
	if o == nil || IsNil(o.ObligationConfig) {
		return nil, false
	}
	return o.ObligationConfig, true
}

// HasObligationConfig returns a boolean if a field has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) HasObligationConfig() bool {
	if o != nil && !IsNil(o.ObligationConfig) {
		return true
	}

	return false
}

// SetObligationConfig gets a reference to the given ZMGOObligationConfig and assigns it to the ObligationConfig field.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) SetObligationConfig(v ZMGOObligationConfig) {
	o.ObligationConfig = &v
}

// GetOpenConfig returns the OpenConfig field value if set, zero value otherwise.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetOpenConfig() ZMGOOpenConfig {
	if o == nil || IsNil(o.OpenConfig) {
		var ret ZMGOOpenConfig
		return ret
	}
	return *o.OpenConfig
}

// GetOpenConfigOk returns a tuple with the OpenConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetOpenConfigOk() (*ZMGOOpenConfig, bool) {
	if o == nil || IsNil(o.OpenConfig) {
		return nil, false
	}
	return o.OpenConfig, true
}

// HasOpenConfig returns a boolean if a field has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) HasOpenConfig() bool {
	if o != nil && !IsNil(o.OpenConfig) {
		return true
	}

	return false
}

// SetOpenConfig gets a reference to the given ZMGOOpenConfig and assigns it to the OpenConfig field.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) SetOpenConfig(v ZMGOOpenConfig) {
	o.OpenConfig = &v
}

// GetQuitConfig returns the QuitConfig field value if set, zero value otherwise.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetQuitConfig() ZMGOQuitConfig {
	if o == nil || IsNil(o.QuitConfig) {
		var ret ZMGOQuitConfig
		return ret
	}
	return *o.QuitConfig
}

// GetQuitConfigOk returns a tuple with the QuitConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetQuitConfigOk() (*ZMGOQuitConfig, bool) {
	if o == nil || IsNil(o.QuitConfig) {
		return nil, false
	}
	return o.QuitConfig, true
}

// HasQuitConfig returns a boolean if a field has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) HasQuitConfig() bool {
	if o != nil && !IsNil(o.QuitConfig) {
		return true
	}

	return false
}

// SetQuitConfig gets a reference to the given ZMGOQuitConfig and assigns it to the QuitConfig field.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) SetQuitConfig(v ZMGOQuitConfig) {
	o.QuitConfig = &v
}

// GetRightConfig returns the RightConfig field value if set, zero value otherwise.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetRightConfig() ZMGORightConfig {
	if o == nil || IsNil(o.RightConfig) {
		var ret ZMGORightConfig
		return ret
	}
	return *o.RightConfig
}

// GetRightConfigOk returns a tuple with the RightConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetRightConfigOk() (*ZMGORightConfig, bool) {
	if o == nil || IsNil(o.RightConfig) {
		return nil, false
	}
	return o.RightConfig, true
}

// HasRightConfig returns a boolean if a field has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) HasRightConfig() bool {
	if o != nil && !IsNil(o.RightConfig) {
		return true
	}

	return false
}

// SetRightConfig gets a reference to the given ZMGORightConfig and assigns it to the RightConfig field.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) SetRightConfig(v ZMGORightConfig) {
	o.RightConfig = &v
}

// GetSettlementConfig returns the SettlementConfig field value if set, zero value otherwise.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetSettlementConfig() ZMGOSettlementConfig {
	if o == nil || IsNil(o.SettlementConfig) {
		var ret ZMGOSettlementConfig
		return ret
	}
	return *o.SettlementConfig
}

// GetSettlementConfigOk returns a tuple with the SettlementConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) GetSettlementConfigOk() (*ZMGOSettlementConfig, bool) {
	if o == nil || IsNil(o.SettlementConfig) {
		return nil, false
	}
	return o.SettlementConfig, true
}

// HasSettlementConfig returns a boolean if a field has been set.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) HasSettlementConfig() bool {
	if o != nil && !IsNil(o.SettlementConfig) {
		return true
	}

	return false
}

// SetSettlementConfig gets a reference to the given ZMGOSettlementConfig and assigns it to the SettlementConfig field.
func (o *ZhimaMerchantZmgoTemplateQueryResponseModel) SetSettlementConfig(v ZMGOSettlementConfig) {
	o.SettlementConfig = &v
}

func (o ZhimaMerchantZmgoTemplateQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaMerchantZmgoTemplateQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasicConfig) {
		toSerialize["basic_config"] = o.BasicConfig
	}
	if !IsNil(o.ExtConfig) {
		toSerialize["ext_config"] = o.ExtConfig
	}
	if !IsNil(o.ObligationConfig) {
		toSerialize["obligation_config"] = o.ObligationConfig
	}
	if !IsNil(o.OpenConfig) {
		toSerialize["open_config"] = o.OpenConfig
	}
	if !IsNil(o.QuitConfig) {
		toSerialize["quit_config"] = o.QuitConfig
	}
	if !IsNil(o.RightConfig) {
		toSerialize["right_config"] = o.RightConfig
	}
	if !IsNil(o.SettlementConfig) {
		toSerialize["settlement_config"] = o.SettlementConfig
	}
	return toSerialize, nil
}

type NullableZhimaMerchantZmgoTemplateQueryResponseModel struct {
	value *ZhimaMerchantZmgoTemplateQueryResponseModel
	isSet bool
}

func (v NullableZhimaMerchantZmgoTemplateQueryResponseModel) Get() *ZhimaMerchantZmgoTemplateQueryResponseModel {
	return v.value
}

func (v *NullableZhimaMerchantZmgoTemplateQueryResponseModel) Set(val *ZhimaMerchantZmgoTemplateQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaMerchantZmgoTemplateQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaMerchantZmgoTemplateQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaMerchantZmgoTemplateQueryResponseModel(val *ZhimaMerchantZmgoTemplateQueryResponseModel) *NullableZhimaMerchantZmgoTemplateQueryResponseModel {
	return &NullableZhimaMerchantZmgoTemplateQueryResponseModel{value: val, isSet: true}
}

func (v NullableZhimaMerchantZmgoTemplateQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaMerchantZmgoTemplateQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


