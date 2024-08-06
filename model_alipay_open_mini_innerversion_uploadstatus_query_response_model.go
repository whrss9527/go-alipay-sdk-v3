/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the AlipayOpenMiniInnerversionUploadstatusQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionUploadstatusQueryResponseModel{}

// AlipayOpenMiniInnerversionUploadstatusQueryResponseModel struct for AlipayOpenMiniInnerversionUploadstatusQueryResponseModel
type AlipayOpenMiniInnerversionUploadstatusQueryResponseModel struct {
	// 构建信息
	BuildInfo *string `json:"build_info,omitempty"`
	// 构建好的包地址
	BuildPackageUrl *string `json:"build_package_url,omitempty"`
	// 构建状态
	BuildStatus *string `json:"build_status,omitempty"`
	// 构建日志地址
	LogUrl *string `json:"log_url,omitempty"`
	// 是否需要轮询
	NeedRotation *string `json:"need_rotation,omitempty"`
	// 构建好的appx2.0包地址
	NewBuildPackageUrl *string `json:"new_build_package_url,omitempty"`
	// 构建的结果地址
	NewResultUrl *string `json:"new_result_url,omitempty"`
	// 构建的结果地址
	ResultUrl *string `json:"result_url,omitempty"`
	// 创建版本结果
	VersionCreated *string `json:"version_created,omitempty"`
}

// NewAlipayOpenMiniInnerversionUploadstatusQueryResponseModel instantiates a new AlipayOpenMiniInnerversionUploadstatusQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionUploadstatusQueryResponseModel() *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel {
	this := AlipayOpenMiniInnerversionUploadstatusQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniInnerversionUploadstatusQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerversionUploadstatusQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionUploadstatusQueryResponseModelWithDefaults() *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel {
	this := AlipayOpenMiniInnerversionUploadstatusQueryResponseModel{}
	return &this
}

// GetBuildInfo returns the BuildInfo field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetBuildInfo() string {
	if o == nil || IsNil(o.BuildInfo) {
		var ret string
		return ret
	}
	return *o.BuildInfo
}

// GetBuildInfoOk returns a tuple with the BuildInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetBuildInfoOk() (*string, bool) {
	if o == nil || IsNil(o.BuildInfo) {
		return nil, false
	}
	return o.BuildInfo, true
}

// HasBuildInfo returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasBuildInfo() bool {
	if o != nil && !IsNil(o.BuildInfo) {
		return true
	}

	return false
}

// SetBuildInfo gets a reference to the given string and assigns it to the BuildInfo field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetBuildInfo(v string) {
	o.BuildInfo = &v
}

// GetBuildPackageUrl returns the BuildPackageUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetBuildPackageUrl() string {
	if o == nil || IsNil(o.BuildPackageUrl) {
		var ret string
		return ret
	}
	return *o.BuildPackageUrl
}

// GetBuildPackageUrlOk returns a tuple with the BuildPackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetBuildPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BuildPackageUrl) {
		return nil, false
	}
	return o.BuildPackageUrl, true
}

// HasBuildPackageUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasBuildPackageUrl() bool {
	if o != nil && !IsNil(o.BuildPackageUrl) {
		return true
	}

	return false
}

// SetBuildPackageUrl gets a reference to the given string and assigns it to the BuildPackageUrl field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetBuildPackageUrl(v string) {
	o.BuildPackageUrl = &v
}

// GetBuildStatus returns the BuildStatus field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetBuildStatus() string {
	if o == nil || IsNil(o.BuildStatus) {
		var ret string
		return ret
	}
	return *o.BuildStatus
}

// GetBuildStatusOk returns a tuple with the BuildStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetBuildStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BuildStatus) {
		return nil, false
	}
	return o.BuildStatus, true
}

// HasBuildStatus returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasBuildStatus() bool {
	if o != nil && !IsNil(o.BuildStatus) {
		return true
	}

	return false
}

// SetBuildStatus gets a reference to the given string and assigns it to the BuildStatus field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetBuildStatus(v string) {
	o.BuildStatus = &v
}

// GetLogUrl returns the LogUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetLogUrl() string {
	if o == nil || IsNil(o.LogUrl) {
		var ret string
		return ret
	}
	return *o.LogUrl
}

// GetLogUrlOk returns a tuple with the LogUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetLogUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogUrl) {
		return nil, false
	}
	return o.LogUrl, true
}

// HasLogUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasLogUrl() bool {
	if o != nil && !IsNil(o.LogUrl) {
		return true
	}

	return false
}

// SetLogUrl gets a reference to the given string and assigns it to the LogUrl field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetLogUrl(v string) {
	o.LogUrl = &v
}

// GetNeedRotation returns the NeedRotation field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetNeedRotation() string {
	if o == nil || IsNil(o.NeedRotation) {
		var ret string
		return ret
	}
	return *o.NeedRotation
}

// GetNeedRotationOk returns a tuple with the NeedRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetNeedRotationOk() (*string, bool) {
	if o == nil || IsNil(o.NeedRotation) {
		return nil, false
	}
	return o.NeedRotation, true
}

// HasNeedRotation returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasNeedRotation() bool {
	if o != nil && !IsNil(o.NeedRotation) {
		return true
	}

	return false
}

// SetNeedRotation gets a reference to the given string and assigns it to the NeedRotation field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetNeedRotation(v string) {
	o.NeedRotation = &v
}

// GetNewBuildPackageUrl returns the NewBuildPackageUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetNewBuildPackageUrl() string {
	if o == nil || IsNil(o.NewBuildPackageUrl) {
		var ret string
		return ret
	}
	return *o.NewBuildPackageUrl
}

// GetNewBuildPackageUrlOk returns a tuple with the NewBuildPackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetNewBuildPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NewBuildPackageUrl) {
		return nil, false
	}
	return o.NewBuildPackageUrl, true
}

// HasNewBuildPackageUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasNewBuildPackageUrl() bool {
	if o != nil && !IsNil(o.NewBuildPackageUrl) {
		return true
	}

	return false
}

// SetNewBuildPackageUrl gets a reference to the given string and assigns it to the NewBuildPackageUrl field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetNewBuildPackageUrl(v string) {
	o.NewBuildPackageUrl = &v
}

// GetNewResultUrl returns the NewResultUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetNewResultUrl() string {
	if o == nil || IsNil(o.NewResultUrl) {
		var ret string
		return ret
	}
	return *o.NewResultUrl
}

// GetNewResultUrlOk returns a tuple with the NewResultUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetNewResultUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NewResultUrl) {
		return nil, false
	}
	return o.NewResultUrl, true
}

// HasNewResultUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasNewResultUrl() bool {
	if o != nil && !IsNil(o.NewResultUrl) {
		return true
	}

	return false
}

// SetNewResultUrl gets a reference to the given string and assigns it to the NewResultUrl field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetNewResultUrl(v string) {
	o.NewResultUrl = &v
}

// GetResultUrl returns the ResultUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetResultUrl() string {
	if o == nil || IsNil(o.ResultUrl) {
		var ret string
		return ret
	}
	return *o.ResultUrl
}

// GetResultUrlOk returns a tuple with the ResultUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetResultUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ResultUrl) {
		return nil, false
	}
	return o.ResultUrl, true
}

// HasResultUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasResultUrl() bool {
	if o != nil && !IsNil(o.ResultUrl) {
		return true
	}

	return false
}

// SetResultUrl gets a reference to the given string and assigns it to the ResultUrl field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetResultUrl(v string) {
	o.ResultUrl = &v
}

// GetVersionCreated returns the VersionCreated field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetVersionCreated() string {
	if o == nil || IsNil(o.VersionCreated) {
		var ret string
		return ret
	}
	return *o.VersionCreated
}

// GetVersionCreatedOk returns a tuple with the VersionCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) GetVersionCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.VersionCreated) {
		return nil, false
	}
	return o.VersionCreated, true
}

// HasVersionCreated returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) HasVersionCreated() bool {
	if o != nil && !IsNil(o.VersionCreated) {
		return true
	}

	return false
}

// SetVersionCreated gets a reference to the given string and assigns it to the VersionCreated field.
func (o *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) SetVersionCreated(v string) {
	o.VersionCreated = &v
}

func (o AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildInfo) {
		toSerialize["build_info"] = o.BuildInfo
	}
	if !IsNil(o.BuildPackageUrl) {
		toSerialize["build_package_url"] = o.BuildPackageUrl
	}
	if !IsNil(o.BuildStatus) {
		toSerialize["build_status"] = o.BuildStatus
	}
	if !IsNil(o.LogUrl) {
		toSerialize["log_url"] = o.LogUrl
	}
	if !IsNil(o.NeedRotation) {
		toSerialize["need_rotation"] = o.NeedRotation
	}
	if !IsNil(o.NewBuildPackageUrl) {
		toSerialize["new_build_package_url"] = o.NewBuildPackageUrl
	}
	if !IsNil(o.NewResultUrl) {
		toSerialize["new_result_url"] = o.NewResultUrl
	}
	if !IsNil(o.ResultUrl) {
		toSerialize["result_url"] = o.ResultUrl
	}
	if !IsNil(o.VersionCreated) {
		toSerialize["version_created"] = o.VersionCreated
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel struct {
	value *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel) Get() *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel) Set(val *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel(val *AlipayOpenMiniInnerversionUploadstatusQueryResponseModel) *NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel {
	return &NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionUploadstatusQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
