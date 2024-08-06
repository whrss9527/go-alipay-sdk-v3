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

// checks if the AlipayIserviceCcmSwArticleBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwArticleBatchqueryModel{}

// AlipayIserviceCcmSwArticleBatchqueryModel struct for AlipayIserviceCcmSwArticleBatchqueryModel
type AlipayIserviceCcmSwArticleBatchqueryModel struct {
	// 所属类目ID，如果search_all_category为true则不用填
	CategoryId *int32 `json:"category_id,omitempty"`
	// 子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// 结束时间，并且开始时间不能为空
	EndTime *string `json:"end_time,omitempty"`
	// 文章ID集合
	Ids []int32 `json:"ids,omitempty"`
	// 关键字
	Keyword *string `json:"keyword,omitempty"`
	// 标签
	Keywords []string `json:"keywords,omitempty"`
	// 知识库ID
	LibraryId *int32 `json:"library_id,omitempty"`
	// 页数，page_size不能为空
	PageNum *int32 `json:"page_num,omitempty"`
	// 页大小，page_num不能为空
	PageSize *int32 `json:"page_size,omitempty"`
	// 是否搜索所有类目，如果为true则不用填写category_id值
	SearchAllCategory *bool `json:"search_all_category,omitempty"`
	// Current（搜索当前节点）； Children（搜索当前节点以及子节点）
	SearchCategoryType *string `json:"search_category_type,omitempty"`
	// 开始时间，并且结束时间不能为空
	StartTime *string `json:"start_time,omitempty"`
	// 状态，PUBLISHED（已发布），UNPUBLISH（未发布），EXPIRED（失效），DELETED（已删除）
	Status *string `json:"status,omitempty"`
}

// NewAlipayIserviceCcmSwArticleBatchqueryModel instantiates a new AlipayIserviceCcmSwArticleBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwArticleBatchqueryModel() *AlipayIserviceCcmSwArticleBatchqueryModel {
	this := AlipayIserviceCcmSwArticleBatchqueryModel{}
	return &this
}

// NewAlipayIserviceCcmSwArticleBatchqueryModelWithDefaults instantiates a new AlipayIserviceCcmSwArticleBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwArticleBatchqueryModelWithDefaults() *AlipayIserviceCcmSwArticleBatchqueryModel {
	this := AlipayIserviceCcmSwArticleBatchqueryModel{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetCategoryId() int32 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetCcsInstanceId returns the CcsInstanceId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetCcsInstanceId() string {
	if o == nil || IsNil(o.CcsInstanceId) {
		var ret string
		return ret
	}
	return *o.CcsInstanceId
}

// GetCcsInstanceIdOk returns a tuple with the CcsInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetCcsInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CcsInstanceId) {
		return nil, false
	}
	return o.CcsInstanceId, true
}

// HasCcsInstanceId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasCcsInstanceId() bool {
	if o != nil && !IsNil(o.CcsInstanceId) {
		return true
	}

	return false
}

// SetCcsInstanceId gets a reference to the given string and assigns it to the CcsInstanceId field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetCcsInstanceId(v string) {
	o.CcsInstanceId = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetEndTime(v string) {
	o.EndTime = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetIds() []int32 {
	if o == nil || IsNil(o.Ids) {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetIds(v []int32) {
	o.Ids = v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetKeyword() string {
	if o == nil || IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetKeyword(v string) {
	o.Keyword = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetKeywords(v []string) {
	o.Keywords = v
}

// GetLibraryId returns the LibraryId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetLibraryId() int32 {
	if o == nil || IsNil(o.LibraryId) {
		var ret int32
		return ret
	}
	return *o.LibraryId
}

// GetLibraryIdOk returns a tuple with the LibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetLibraryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LibraryId) {
		return nil, false
	}
	return o.LibraryId, true
}

// HasLibraryId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasLibraryId() bool {
	if o != nil && !IsNil(o.LibraryId) {
		return true
	}

	return false
}

// SetLibraryId gets a reference to the given int32 and assigns it to the LibraryId field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetLibraryId(v int32) {
	o.LibraryId = &v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSearchAllCategory returns the SearchAllCategory field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetSearchAllCategory() bool {
	if o == nil || IsNil(o.SearchAllCategory) {
		var ret bool
		return ret
	}
	return *o.SearchAllCategory
}

// GetSearchAllCategoryOk returns a tuple with the SearchAllCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetSearchAllCategoryOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchAllCategory) {
		return nil, false
	}
	return o.SearchAllCategory, true
}

// HasSearchAllCategory returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasSearchAllCategory() bool {
	if o != nil && !IsNil(o.SearchAllCategory) {
		return true
	}

	return false
}

// SetSearchAllCategory gets a reference to the given bool and assigns it to the SearchAllCategory field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetSearchAllCategory(v bool) {
	o.SearchAllCategory = &v
}

// GetSearchCategoryType returns the SearchCategoryType field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetSearchCategoryType() string {
	if o == nil || IsNil(o.SearchCategoryType) {
		var ret string
		return ret
	}
	return *o.SearchCategoryType
}

// GetSearchCategoryTypeOk returns a tuple with the SearchCategoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetSearchCategoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SearchCategoryType) {
		return nil, false
	}
	return o.SearchCategoryType, true
}

// HasSearchCategoryType returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasSearchCategoryType() bool {
	if o != nil && !IsNil(o.SearchCategoryType) {
		return true
	}

	return false
}

// SetSearchCategoryType gets a reference to the given string and assigns it to the SearchCategoryType field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetSearchCategoryType(v string) {
	o.SearchCategoryType = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayIserviceCcmSwArticleBatchqueryModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayIserviceCcmSwArticleBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwArticleBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.CcsInstanceId) {
		toSerialize["ccs_instance_id"] = o.CcsInstanceId
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.LibraryId) {
		toSerialize["library_id"] = o.LibraryId
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.SearchAllCategory) {
		toSerialize["search_all_category"] = o.SearchAllCategory
	}
	if !IsNil(o.SearchCategoryType) {
		toSerialize["search_category_type"] = o.SearchCategoryType
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwArticleBatchqueryModel struct {
	value *AlipayIserviceCcmSwArticleBatchqueryModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwArticleBatchqueryModel) Get() *AlipayIserviceCcmSwArticleBatchqueryModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwArticleBatchqueryModel) Set(val *AlipayIserviceCcmSwArticleBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwArticleBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwArticleBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwArticleBatchqueryModel(val *AlipayIserviceCcmSwArticleBatchqueryModel) *NullableAlipayIserviceCcmSwArticleBatchqueryModel {
	return &NullableAlipayIserviceCcmSwArticleBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwArticleBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwArticleBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
