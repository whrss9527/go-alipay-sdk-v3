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

// checks if the LifeLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifeLabel{}

// LifeLabel struct for LifeLabel
type LifeLabel struct {
	// 该标签支持的业务列表，menu表示个性化菜单，extension表示个性化扩展区，message表示消息触达
	Biz *string `json:"biz,omitempty"`
	// 标签类目
	Category *string `json:"category,omitempty"`
	// 标签值数据类型
	DataType *string `json:"data_type,omitempty"`
	// 标签英文代码
	LabelCode *string `json:"label_code,omitempty"`
	// 标签id，唯一标识一个标签
	LabelId *string `json:"label_id,omitempty"`
	// 标签名
	LabelName *string `json:"label_name,omitempty"`
	// 该标签支持的运算符
	Operator *string `json:"operator,omitempty"`
	// 每个取值的业务含义
	Options []Option `json:"options,omitempty"`
	// 标签类型，目前分为common（通用标签）、custom（生活号自定义标签）、cloud（云实验室标签）
	Type *string `json:"type,omitempty"`
}

// NewLifeLabel instantiates a new LifeLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifeLabel() *LifeLabel {
	this := LifeLabel{}
	return &this
}

// NewLifeLabelWithDefaults instantiates a new LifeLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifeLabelWithDefaults() *LifeLabel {
	this := LifeLabel{}
	return &this
}

// GetBiz returns the Biz field value if set, zero value otherwise.
func (o *LifeLabel) GetBiz() string {
	if o == nil || IsNil(o.Biz) {
		var ret string
		return ret
	}
	return *o.Biz
}

// GetBizOk returns a tuple with the Biz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetBizOk() (*string, bool) {
	if o == nil || IsNil(o.Biz) {
		return nil, false
	}
	return o.Biz, true
}

// HasBiz returns a boolean if a field has been set.
func (o *LifeLabel) HasBiz() bool {
	if o != nil && !IsNil(o.Biz) {
		return true
	}

	return false
}

// SetBiz gets a reference to the given string and assigns it to the Biz field.
func (o *LifeLabel) SetBiz(v string) {
	o.Biz = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *LifeLabel) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *LifeLabel) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *LifeLabel) SetCategory(v string) {
	o.Category = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *LifeLabel) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *LifeLabel) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *LifeLabel) SetDataType(v string) {
	o.DataType = &v
}

// GetLabelCode returns the LabelCode field value if set, zero value otherwise.
func (o *LifeLabel) GetLabelCode() string {
	if o == nil || IsNil(o.LabelCode) {
		var ret string
		return ret
	}
	return *o.LabelCode
}

// GetLabelCodeOk returns a tuple with the LabelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetLabelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LabelCode) {
		return nil, false
	}
	return o.LabelCode, true
}

// HasLabelCode returns a boolean if a field has been set.
func (o *LifeLabel) HasLabelCode() bool {
	if o != nil && !IsNil(o.LabelCode) {
		return true
	}

	return false
}

// SetLabelCode gets a reference to the given string and assigns it to the LabelCode field.
func (o *LifeLabel) SetLabelCode(v string) {
	o.LabelCode = &v
}

// GetLabelId returns the LabelId field value if set, zero value otherwise.
func (o *LifeLabel) GetLabelId() string {
	if o == nil || IsNil(o.LabelId) {
		var ret string
		return ret
	}
	return *o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetLabelIdOk() (*string, bool) {
	if o == nil || IsNil(o.LabelId) {
		return nil, false
	}
	return o.LabelId, true
}

// HasLabelId returns a boolean if a field has been set.
func (o *LifeLabel) HasLabelId() bool {
	if o != nil && !IsNil(o.LabelId) {
		return true
	}

	return false
}

// SetLabelId gets a reference to the given string and assigns it to the LabelId field.
func (o *LifeLabel) SetLabelId(v string) {
	o.LabelId = &v
}

// GetLabelName returns the LabelName field value if set, zero value otherwise.
func (o *LifeLabel) GetLabelName() string {
	if o == nil || IsNil(o.LabelName) {
		var ret string
		return ret
	}
	return *o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetLabelNameOk() (*string, bool) {
	if o == nil || IsNil(o.LabelName) {
		return nil, false
	}
	return o.LabelName, true
}

// HasLabelName returns a boolean if a field has been set.
func (o *LifeLabel) HasLabelName() bool {
	if o != nil && !IsNil(o.LabelName) {
		return true
	}

	return false
}

// SetLabelName gets a reference to the given string and assigns it to the LabelName field.
func (o *LifeLabel) SetLabelName(v string) {
	o.LabelName = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *LifeLabel) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *LifeLabel) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *LifeLabel) SetOperator(v string) {
	o.Operator = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *LifeLabel) GetOptions() []Option {
	if o == nil || IsNil(o.Options) {
		var ret []Option
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetOptionsOk() ([]Option, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LifeLabel) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []Option and assigns it to the Options field.
func (o *LifeLabel) SetOptions(v []Option) {
	o.Options = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LifeLabel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifeLabel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LifeLabel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LifeLabel) SetType(v string) {
	o.Type = &v
}

func (o LifeLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifeLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Biz) {
		toSerialize["biz"] = o.Biz
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !IsNil(o.LabelCode) {
		toSerialize["label_code"] = o.LabelCode
	}
	if !IsNil(o.LabelId) {
		toSerialize["label_id"] = o.LabelId
	}
	if !IsNil(o.LabelName) {
		toSerialize["label_name"] = o.LabelName
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableLifeLabel struct {
	value *LifeLabel
	isSet bool
}

func (v NullableLifeLabel) Get() *LifeLabel {
	return v.value
}

func (v *NullableLifeLabel) Set(val *LifeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableLifeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLifeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifeLabel(val *LifeLabel) *NullableLifeLabel {
	return &NullableLifeLabel{value: val, isSet: true}
}

func (v NullableLifeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
