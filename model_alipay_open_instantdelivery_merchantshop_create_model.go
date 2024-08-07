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

// checks if the AlipayOpenInstantdeliveryMerchantshopCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenInstantdeliveryMerchantshopCreateModel{}

// AlipayOpenInstantdeliveryMerchantshopCreateModel struct for AlipayOpenInstantdeliveryMerchantshopCreateModel
type AlipayOpenInstantdeliveryMerchantshopCreateModel struct {
	// 联系人姓名
	ContactName *string `json:"contact_name,omitempty"`
	// 地址。商户详细经营地址或人员所在地点
	DetailAddress *string `json:"detail_address,omitempty"`
	// 城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	EnterpriseCity *string `json:"enterprise_city,omitempty"`
	// 区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	EnterpriseDistrict *string `json:"enterprise_district,omitempty"`
	// 省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	EnterpriseProvince *string `json:"enterprise_province,omitempty"`
	// 纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
	Latitude *string `json:"latitude,omitempty"`
	// 经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
	Longitude *string `json:"longitude,omitempty"`
	// 外部业务号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 联系人电话/手机号
	Phone *string `json:"phone,omitempty"`
	// 高德poiid
	Poiid *string `json:"poiid,omitempty"`
	// 店铺类目，取值参见文件https://mif-pub.alipayobjects.com/ShopCategory.xlsx 中的三级门店类目
	ShopCategory *string `json:"shop_category,omitempty"`
	// 门店名称，最长不超过256个字符。名称+地址需要全局唯一。
	ShopName *string `json:"shop_name,omitempty"`
	// 商家门店编码，可自定义，但必须唯一
	ShopNo *string `json:"shop_no,omitempty"`
}

// NewAlipayOpenInstantdeliveryMerchantshopCreateModel instantiates a new AlipayOpenInstantdeliveryMerchantshopCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenInstantdeliveryMerchantshopCreateModel() *AlipayOpenInstantdeliveryMerchantshopCreateModel {
	this := AlipayOpenInstantdeliveryMerchantshopCreateModel{}
	return &this
}

// NewAlipayOpenInstantdeliveryMerchantshopCreateModelWithDefaults instantiates a new AlipayOpenInstantdeliveryMerchantshopCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenInstantdeliveryMerchantshopCreateModelWithDefaults() *AlipayOpenInstantdeliveryMerchantshopCreateModel {
	this := AlipayOpenInstantdeliveryMerchantshopCreateModel{}
	return &this
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetContactName(v string) {
	o.ContactName = &v
}

// GetDetailAddress returns the DetailAddress field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetDetailAddress() string {
	if o == nil || IsNil(o.DetailAddress) {
		var ret string
		return ret
	}
	return *o.DetailAddress
}

// GetDetailAddressOk returns a tuple with the DetailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetDetailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DetailAddress) {
		return nil, false
	}
	return o.DetailAddress, true
}

// HasDetailAddress returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasDetailAddress() bool {
	if o != nil && !IsNil(o.DetailAddress) {
		return true
	}

	return false
}

// SetDetailAddress gets a reference to the given string and assigns it to the DetailAddress field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetDetailAddress(v string) {
	o.DetailAddress = &v
}

// GetEnterpriseCity returns the EnterpriseCity field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetEnterpriseCity() string {
	if o == nil || IsNil(o.EnterpriseCity) {
		var ret string
		return ret
	}
	return *o.EnterpriseCity
}

// GetEnterpriseCityOk returns a tuple with the EnterpriseCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetEnterpriseCityOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseCity) {
		return nil, false
	}
	return o.EnterpriseCity, true
}

// HasEnterpriseCity returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasEnterpriseCity() bool {
	if o != nil && !IsNil(o.EnterpriseCity) {
		return true
	}

	return false
}

// SetEnterpriseCity gets a reference to the given string and assigns it to the EnterpriseCity field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetEnterpriseCity(v string) {
	o.EnterpriseCity = &v
}

// GetEnterpriseDistrict returns the EnterpriseDistrict field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetEnterpriseDistrict() string {
	if o == nil || IsNil(o.EnterpriseDistrict) {
		var ret string
		return ret
	}
	return *o.EnterpriseDistrict
}

// GetEnterpriseDistrictOk returns a tuple with the EnterpriseDistrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetEnterpriseDistrictOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseDistrict) {
		return nil, false
	}
	return o.EnterpriseDistrict, true
}

// HasEnterpriseDistrict returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasEnterpriseDistrict() bool {
	if o != nil && !IsNil(o.EnterpriseDistrict) {
		return true
	}

	return false
}

// SetEnterpriseDistrict gets a reference to the given string and assigns it to the EnterpriseDistrict field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetEnterpriseDistrict(v string) {
	o.EnterpriseDistrict = &v
}

// GetEnterpriseProvince returns the EnterpriseProvince field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetEnterpriseProvince() string {
	if o == nil || IsNil(o.EnterpriseProvince) {
		var ret string
		return ret
	}
	return *o.EnterpriseProvince
}

// GetEnterpriseProvinceOk returns a tuple with the EnterpriseProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetEnterpriseProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseProvince) {
		return nil, false
	}
	return o.EnterpriseProvince, true
}

// HasEnterpriseProvince returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasEnterpriseProvince() bool {
	if o != nil && !IsNil(o.EnterpriseProvince) {
		return true
	}

	return false
}

// SetEnterpriseProvince gets a reference to the given string and assigns it to the EnterpriseProvince field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetEnterpriseProvince(v string) {
	o.EnterpriseProvince = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetLongitude(v string) {
	o.Longitude = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetPhone(v string) {
	o.Phone = &v
}

// GetPoiid returns the Poiid field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetPoiid() string {
	if o == nil || IsNil(o.Poiid) {
		var ret string
		return ret
	}
	return *o.Poiid
}

// GetPoiidOk returns a tuple with the Poiid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetPoiidOk() (*string, bool) {
	if o == nil || IsNil(o.Poiid) {
		return nil, false
	}
	return o.Poiid, true
}

// HasPoiid returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasPoiid() bool {
	if o != nil && !IsNil(o.Poiid) {
		return true
	}

	return false
}

// SetPoiid gets a reference to the given string and assigns it to the Poiid field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetPoiid(v string) {
	o.Poiid = &v
}

// GetShopCategory returns the ShopCategory field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetShopCategory() string {
	if o == nil || IsNil(o.ShopCategory) {
		var ret string
		return ret
	}
	return *o.ShopCategory
}

// GetShopCategoryOk returns a tuple with the ShopCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetShopCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ShopCategory) {
		return nil, false
	}
	return o.ShopCategory, true
}

// HasShopCategory returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasShopCategory() bool {
	if o != nil && !IsNil(o.ShopCategory) {
		return true
	}

	return false
}

// SetShopCategory gets a reference to the given string and assigns it to the ShopCategory field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetShopCategory(v string) {
	o.ShopCategory = &v
}

// GetShopName returns the ShopName field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetShopName() string {
	if o == nil || IsNil(o.ShopName) {
		var ret string
		return ret
	}
	return *o.ShopName
}

// GetShopNameOk returns a tuple with the ShopName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetShopNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShopName) {
		return nil, false
	}
	return o.ShopName, true
}

// HasShopName returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasShopName() bool {
	if o != nil && !IsNil(o.ShopName) {
		return true
	}

	return false
}

// SetShopName gets a reference to the given string and assigns it to the ShopName field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetShopName(v string) {
	o.ShopName = &v
}

// GetShopNo returns the ShopNo field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetShopNo() string {
	if o == nil || IsNil(o.ShopNo) {
		var ret string
		return ret
	}
	return *o.ShopNo
}

// GetShopNoOk returns a tuple with the ShopNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) GetShopNoOk() (*string, bool) {
	if o == nil || IsNil(o.ShopNo) {
		return nil, false
	}
	return o.ShopNo, true
}

// HasShopNo returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) HasShopNo() bool {
	if o != nil && !IsNil(o.ShopNo) {
		return true
	}

	return false
}

// SetShopNo gets a reference to the given string and assigns it to the ShopNo field.
func (o *AlipayOpenInstantdeliveryMerchantshopCreateModel) SetShopNo(v string) {
	o.ShopNo = &v
}

func (o AlipayOpenInstantdeliveryMerchantshopCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenInstantdeliveryMerchantshopCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactName) {
		toSerialize["contact_name"] = o.ContactName
	}
	if !IsNil(o.DetailAddress) {
		toSerialize["detail_address"] = o.DetailAddress
	}
	if !IsNil(o.EnterpriseCity) {
		toSerialize["enterprise_city"] = o.EnterpriseCity
	}
	if !IsNil(o.EnterpriseDistrict) {
		toSerialize["enterprise_district"] = o.EnterpriseDistrict
	}
	if !IsNil(o.EnterpriseProvince) {
		toSerialize["enterprise_province"] = o.EnterpriseProvince
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Poiid) {
		toSerialize["poiid"] = o.Poiid
	}
	if !IsNil(o.ShopCategory) {
		toSerialize["shop_category"] = o.ShopCategory
	}
	if !IsNil(o.ShopName) {
		toSerialize["shop_name"] = o.ShopName
	}
	if !IsNil(o.ShopNo) {
		toSerialize["shop_no"] = o.ShopNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenInstantdeliveryMerchantshopCreateModel struct {
	value *AlipayOpenInstantdeliveryMerchantshopCreateModel
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopCreateModel) Get() *AlipayOpenInstantdeliveryMerchantshopCreateModel {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopCreateModel) Set(val *AlipayOpenInstantdeliveryMerchantshopCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryMerchantshopCreateModel(val *AlipayOpenInstantdeliveryMerchantshopCreateModel) *NullableAlipayOpenInstantdeliveryMerchantshopCreateModel {
	return &NullableAlipayOpenInstantdeliveryMerchantshopCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
