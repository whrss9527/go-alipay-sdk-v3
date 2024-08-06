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

// checks if the FaceMachineInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaceMachineInfo{}

// FaceMachineInfo struct for FaceMachineInfo
type FaceMachineInfo struct {
	// 摄像头驱动版本号
	CameraDriveVer *string `json:"camera_drive_ver,omitempty"`
	// 摄像头型号
	CameraModel *string `json:"camera_model,omitempty"`
	// 摄像头名称
	CameraName *string `json:"camera_name,omitempty"`
	// 摄像头版本号
	CameraVer *string `json:"camera_ver,omitempty"`
	// 扩展信息
	Ext *string `json:"ext,omitempty"`
	// 机具编码
	MachineCode *string `json:"machine_code,omitempty"`
	// 机具型号
	MachineModel *string `json:"machine_model,omitempty"`
	// 机具版本号
	MachineVer *string `json:"machine_ver,omitempty"`
}

// NewFaceMachineInfo instantiates a new FaceMachineInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaceMachineInfo() *FaceMachineInfo {
	this := FaceMachineInfo{}
	return &this
}

// NewFaceMachineInfoWithDefaults instantiates a new FaceMachineInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaceMachineInfoWithDefaults() *FaceMachineInfo {
	this := FaceMachineInfo{}
	return &this
}

// GetCameraDriveVer returns the CameraDriveVer field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetCameraDriveVer() string {
	if o == nil || IsNil(o.CameraDriveVer) {
		var ret string
		return ret
	}
	return *o.CameraDriveVer
}

// GetCameraDriveVerOk returns a tuple with the CameraDriveVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetCameraDriveVerOk() (*string, bool) {
	if o == nil || IsNil(o.CameraDriveVer) {
		return nil, false
	}
	return o.CameraDriveVer, true
}

// HasCameraDriveVer returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasCameraDriveVer() bool {
	if o != nil && !IsNil(o.CameraDriveVer) {
		return true
	}

	return false
}

// SetCameraDriveVer gets a reference to the given string and assigns it to the CameraDriveVer field.
func (o *FaceMachineInfo) SetCameraDriveVer(v string) {
	o.CameraDriveVer = &v
}

// GetCameraModel returns the CameraModel field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetCameraModel() string {
	if o == nil || IsNil(o.CameraModel) {
		var ret string
		return ret
	}
	return *o.CameraModel
}

// GetCameraModelOk returns a tuple with the CameraModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetCameraModelOk() (*string, bool) {
	if o == nil || IsNil(o.CameraModel) {
		return nil, false
	}
	return o.CameraModel, true
}

// HasCameraModel returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasCameraModel() bool {
	if o != nil && !IsNil(o.CameraModel) {
		return true
	}

	return false
}

// SetCameraModel gets a reference to the given string and assigns it to the CameraModel field.
func (o *FaceMachineInfo) SetCameraModel(v string) {
	o.CameraModel = &v
}

// GetCameraName returns the CameraName field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetCameraName() string {
	if o == nil || IsNil(o.CameraName) {
		var ret string
		return ret
	}
	return *o.CameraName
}

// GetCameraNameOk returns a tuple with the CameraName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetCameraNameOk() (*string, bool) {
	if o == nil || IsNil(o.CameraName) {
		return nil, false
	}
	return o.CameraName, true
}

// HasCameraName returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasCameraName() bool {
	if o != nil && !IsNil(o.CameraName) {
		return true
	}

	return false
}

// SetCameraName gets a reference to the given string and assigns it to the CameraName field.
func (o *FaceMachineInfo) SetCameraName(v string) {
	o.CameraName = &v
}

// GetCameraVer returns the CameraVer field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetCameraVer() string {
	if o == nil || IsNil(o.CameraVer) {
		var ret string
		return ret
	}
	return *o.CameraVer
}

// GetCameraVerOk returns a tuple with the CameraVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetCameraVerOk() (*string, bool) {
	if o == nil || IsNil(o.CameraVer) {
		return nil, false
	}
	return o.CameraVer, true
}

// HasCameraVer returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasCameraVer() bool {
	if o != nil && !IsNil(o.CameraVer) {
		return true
	}

	return false
}

// SetCameraVer gets a reference to the given string and assigns it to the CameraVer field.
func (o *FaceMachineInfo) SetCameraVer(v string) {
	o.CameraVer = &v
}

// GetExt returns the Ext field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetExt() string {
	if o == nil || IsNil(o.Ext) {
		var ret string
		return ret
	}
	return *o.Ext
}

// GetExtOk returns a tuple with the Ext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetExtOk() (*string, bool) {
	if o == nil || IsNil(o.Ext) {
		return nil, false
	}
	return o.Ext, true
}

// HasExt returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasExt() bool {
	if o != nil && !IsNil(o.Ext) {
		return true
	}

	return false
}

// SetExt gets a reference to the given string and assigns it to the Ext field.
func (o *FaceMachineInfo) SetExt(v string) {
	o.Ext = &v
}

// GetMachineCode returns the MachineCode field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetMachineCode() string {
	if o == nil || IsNil(o.MachineCode) {
		var ret string
		return ret
	}
	return *o.MachineCode
}

// GetMachineCodeOk returns a tuple with the MachineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetMachineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MachineCode) {
		return nil, false
	}
	return o.MachineCode, true
}

// HasMachineCode returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasMachineCode() bool {
	if o != nil && !IsNil(o.MachineCode) {
		return true
	}

	return false
}

// SetMachineCode gets a reference to the given string and assigns it to the MachineCode field.
func (o *FaceMachineInfo) SetMachineCode(v string) {
	o.MachineCode = &v
}

// GetMachineModel returns the MachineModel field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetMachineModel() string {
	if o == nil || IsNil(o.MachineModel) {
		var ret string
		return ret
	}
	return *o.MachineModel
}

// GetMachineModelOk returns a tuple with the MachineModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetMachineModelOk() (*string, bool) {
	if o == nil || IsNil(o.MachineModel) {
		return nil, false
	}
	return o.MachineModel, true
}

// HasMachineModel returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasMachineModel() bool {
	if o != nil && !IsNil(o.MachineModel) {
		return true
	}

	return false
}

// SetMachineModel gets a reference to the given string and assigns it to the MachineModel field.
func (o *FaceMachineInfo) SetMachineModel(v string) {
	o.MachineModel = &v
}

// GetMachineVer returns the MachineVer field value if set, zero value otherwise.
func (o *FaceMachineInfo) GetMachineVer() string {
	if o == nil || IsNil(o.MachineVer) {
		var ret string
		return ret
	}
	return *o.MachineVer
}

// GetMachineVerOk returns a tuple with the MachineVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceMachineInfo) GetMachineVerOk() (*string, bool) {
	if o == nil || IsNil(o.MachineVer) {
		return nil, false
	}
	return o.MachineVer, true
}

// HasMachineVer returns a boolean if a field has been set.
func (o *FaceMachineInfo) HasMachineVer() bool {
	if o != nil && !IsNil(o.MachineVer) {
		return true
	}

	return false
}

// SetMachineVer gets a reference to the given string and assigns it to the MachineVer field.
func (o *FaceMachineInfo) SetMachineVer(v string) {
	o.MachineVer = &v
}

func (o FaceMachineInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaceMachineInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CameraDriveVer) {
		toSerialize["camera_drive_ver"] = o.CameraDriveVer
	}
	if !IsNil(o.CameraModel) {
		toSerialize["camera_model"] = o.CameraModel
	}
	if !IsNil(o.CameraName) {
		toSerialize["camera_name"] = o.CameraName
	}
	if !IsNil(o.CameraVer) {
		toSerialize["camera_ver"] = o.CameraVer
	}
	if !IsNil(o.Ext) {
		toSerialize["ext"] = o.Ext
	}
	if !IsNil(o.MachineCode) {
		toSerialize["machine_code"] = o.MachineCode
	}
	if !IsNil(o.MachineModel) {
		toSerialize["machine_model"] = o.MachineModel
	}
	if !IsNil(o.MachineVer) {
		toSerialize["machine_ver"] = o.MachineVer
	}
	return toSerialize, nil
}

type NullableFaceMachineInfo struct {
	value *FaceMachineInfo
	isSet bool
}

func (v NullableFaceMachineInfo) Get() *FaceMachineInfo {
	return v.value
}

func (v *NullableFaceMachineInfo) Set(val *FaceMachineInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFaceMachineInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFaceMachineInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaceMachineInfo(val *FaceMachineInfo) *NullableFaceMachineInfo {
	return &NullableFaceMachineInfo{value: val, isSet: true}
}

func (v NullableFaceMachineInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaceMachineInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


