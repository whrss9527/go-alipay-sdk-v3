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

// checks if the OpenApiResponseHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenApiResponseHeader{}

// OpenApiResponseHeader struct for OpenApiResponseHeader
type OpenApiResponseHeader struct {
	// 会话id，对应请求中的conversation_id，若请求中conversation_id非空，则该值非空
	ConversationId *string `json:"conversation_id,omitempty"`
	// 响应id，对应请求中的request_id。如果请求中request_id非空，则response_id非空。
	ResponseId *string `json:"response_id,omitempty"`
	// 自定义状态码
	StatusCode *string `json:"status_code,omitempty"`
	// 状态信息
	StatusMessage *string `json:"status_message,omitempty"`
	// 子状态码
	SubStatusCode *string `json:"sub_status_code,omitempty"`
}

// NewOpenApiResponseHeader instantiates a new OpenApiResponseHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenApiResponseHeader() *OpenApiResponseHeader {
	this := OpenApiResponseHeader{}
	return &this
}

// NewOpenApiResponseHeaderWithDefaults instantiates a new OpenApiResponseHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenApiResponseHeaderWithDefaults() *OpenApiResponseHeader {
	this := OpenApiResponseHeader{}
	return &this
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *OpenApiResponseHeader) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiResponseHeader) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *OpenApiResponseHeader) HasConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *OpenApiResponseHeader) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetResponseId returns the ResponseId field value if set, zero value otherwise.
func (o *OpenApiResponseHeader) GetResponseId() string {
	if o == nil || IsNil(o.ResponseId) {
		var ret string
		return ret
	}
	return *o.ResponseId
}

// GetResponseIdOk returns a tuple with the ResponseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiResponseHeader) GetResponseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseId) {
		return nil, false
	}
	return o.ResponseId, true
}

// HasResponseId returns a boolean if a field has been set.
func (o *OpenApiResponseHeader) HasResponseId() bool {
	if o != nil && !IsNil(o.ResponseId) {
		return true
	}

	return false
}

// SetResponseId gets a reference to the given string and assigns it to the ResponseId field.
func (o *OpenApiResponseHeader) SetResponseId(v string) {
	o.ResponseId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *OpenApiResponseHeader) GetStatusCode() string {
	if o == nil || IsNil(o.StatusCode) {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiResponseHeader) GetStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *OpenApiResponseHeader) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *OpenApiResponseHeader) SetStatusCode(v string) {
	o.StatusCode = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *OpenApiResponseHeader) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiResponseHeader) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *OpenApiResponseHeader) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *OpenApiResponseHeader) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetSubStatusCode returns the SubStatusCode field value if set, zero value otherwise.
func (o *OpenApiResponseHeader) GetSubStatusCode() string {
	if o == nil || IsNil(o.SubStatusCode) {
		var ret string
		return ret
	}
	return *o.SubStatusCode
}

// GetSubStatusCodeOk returns a tuple with the SubStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiResponseHeader) GetSubStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SubStatusCode) {
		return nil, false
	}
	return o.SubStatusCode, true
}

// HasSubStatusCode returns a boolean if a field has been set.
func (o *OpenApiResponseHeader) HasSubStatusCode() bool {
	if o != nil && !IsNil(o.SubStatusCode) {
		return true
	}

	return false
}

// SetSubStatusCode gets a reference to the given string and assigns it to the SubStatusCode field.
func (o *OpenApiResponseHeader) SetSubStatusCode(v string) {
	o.SubStatusCode = &v
}

func (o OpenApiResponseHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenApiResponseHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConversationId) {
		toSerialize["conversation_id"] = o.ConversationId
	}
	if !IsNil(o.ResponseId) {
		toSerialize["response_id"] = o.ResponseId
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !IsNil(o.SubStatusCode) {
		toSerialize["sub_status_code"] = o.SubStatusCode
	}
	return toSerialize, nil
}

type NullableOpenApiResponseHeader struct {
	value *OpenApiResponseHeader
	isSet bool
}

func (v NullableOpenApiResponseHeader) Get() *OpenApiResponseHeader {
	return v.value
}

func (v *NullableOpenApiResponseHeader) Set(val *OpenApiResponseHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenApiResponseHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenApiResponseHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenApiResponseHeader(val *OpenApiResponseHeader) *NullableOpenApiResponseHeader {
	return &NullableOpenApiResponseHeader{value: val, isSet: true}
}

func (v NullableOpenApiResponseHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenApiResponseHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


