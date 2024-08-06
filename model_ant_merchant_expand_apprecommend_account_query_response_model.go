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

// checks if the AntMerchantExpandApprecommendAccountQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandApprecommendAccountQueryResponseModel{}

// AntMerchantExpandApprecommendAccountQueryResponseModel struct for AntMerchantExpandApprecommendAccountQueryResponseModel
type AntMerchantExpandApprecommendAccountQueryResponseModel struct {
	// 账号列表
	List []string `json:"list,omitempty"`
	// 当前页码
	PageNumber *int32 `json:"page_number,omitempty"`
	// 单页行数，单位：行
	PageSize *int32 `json:"page_size,omitempty"`
	// 总页数，单位：页
	TotalPages *int32 `json:"total_pages,omitempty"`
	// 总行数，  单位：行
	TotalSize *string `json:"total_size,omitempty"`
}

// NewAntMerchantExpandApprecommendAccountQueryResponseModel instantiates a new AntMerchantExpandApprecommendAccountQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandApprecommendAccountQueryResponseModel() *AntMerchantExpandApprecommendAccountQueryResponseModel {
	this := AntMerchantExpandApprecommendAccountQueryResponseModel{}
	return &this
}

// NewAntMerchantExpandApprecommendAccountQueryResponseModelWithDefaults instantiates a new AntMerchantExpandApprecommendAccountQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandApprecommendAccountQueryResponseModelWithDefaults() *AntMerchantExpandApprecommendAccountQueryResponseModel {
	this := AntMerchantExpandApprecommendAccountQueryResponseModel{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetList() []string {
	if o == nil || IsNil(o.List) {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetListOk() ([]string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) SetList(v []string) {
	o.List = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetTotalSize() string {
	if o == nil || IsNil(o.TotalSize) {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) GetTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *AntMerchantExpandApprecommendAccountQueryResponseModel) SetTotalSize(v string) {
	o.TotalSize = &v
}

func (o AntMerchantExpandApprecommendAccountQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandApprecommendAccountQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.PageNumber) {
		toSerialize["page_number"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandApprecommendAccountQueryResponseModel struct {
	value *AntMerchantExpandApprecommendAccountQueryResponseModel
	isSet bool
}

func (v NullableAntMerchantExpandApprecommendAccountQueryResponseModel) Get() *AntMerchantExpandApprecommendAccountQueryResponseModel {
	return v.value
}

func (v *NullableAntMerchantExpandApprecommendAccountQueryResponseModel) Set(val *AntMerchantExpandApprecommendAccountQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandApprecommendAccountQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandApprecommendAccountQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandApprecommendAccountQueryResponseModel(val *AntMerchantExpandApprecommendAccountQueryResponseModel) *NullableAntMerchantExpandApprecommendAccountQueryResponseModel {
	return &NullableAntMerchantExpandApprecommendAccountQueryResponseModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandApprecommendAccountQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandApprecommendAccountQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


