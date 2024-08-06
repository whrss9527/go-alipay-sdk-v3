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

// checks if the AlipayOpenSpBlueseaactivityModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSpBlueseaactivityModifyModel{}

// AlipayOpenSpBlueseaactivityModifyModel struct for AlipayOpenSpBlueseaactivityModifyModel
type AlipayOpenSpBlueseaactivityModifyModel struct {
	// 详细地址
	Address *string `json:"address,omitempty"`
	// 营业执照，要求证件文本信息清晰可见。 请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	BusinessLic *string `json:"business_lic,omitempty"`
	// 城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	CityCode *string `json:"city_code,omitempty"`
	// 区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	DistrictCode *string `json:"district_code,omitempty"`
	// 食品经营许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodBusinessLic *string `json:"food_business_lic,omitempty"`
	// 食品流通许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodCirculateLic *string `json:"food_circulate_lic,omitempty"`
	// 食品卫生许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodHealthLic *string `json:"food_health_lic,omitempty"`
	// 食品生产许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodProductionLic *string `json:"food_production_lic,omitempty"`
	// 餐饮服务许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	FoodServiceLic *string `json:"food_service_lic,omitempty"`
	// 门头照，要求内景照片清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	IndoorPic *string `json:"indoor_pic,omitempty"`
	// 申请单 id。通过 <a href=\"https://opendocs.alipay.com/apis/01ebig\">alipay.open.sp.blueseaactivity.create</a>接口获取。
	OrderId *string `json:"order_id,omitempty"`
	// 省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	ProvinceCode *string `json:"province_code,omitempty"`
	// 门头照，要求店铺外观照片清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	ShopEntrancePic *string `json:"shop_entrance_pic,omitempty"`
	// 烟草专卖零售许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href=\"https://opendocs.alipay.com/open/01hd83\">报名资质要求</a>
	TobaccoLic *string `json:"tobacco_lic,omitempty"`
}

// NewAlipayOpenSpBlueseaactivityModifyModel instantiates a new AlipayOpenSpBlueseaactivityModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSpBlueseaactivityModifyModel() *AlipayOpenSpBlueseaactivityModifyModel {
	this := AlipayOpenSpBlueseaactivityModifyModel{}
	return &this
}

// NewAlipayOpenSpBlueseaactivityModifyModelWithDefaults instantiates a new AlipayOpenSpBlueseaactivityModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSpBlueseaactivityModifyModelWithDefaults() *AlipayOpenSpBlueseaactivityModifyModel {
	this := AlipayOpenSpBlueseaactivityModifyModel{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetAddress(v string) {
	o.Address = &v
}

// GetBusinessLic returns the BusinessLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetBusinessLic() string {
	if o == nil || IsNil(o.BusinessLic) {
		var ret string
		return ret
	}
	return *o.BusinessLic
}

// GetBusinessLicOk returns a tuple with the BusinessLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetBusinessLicOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLic) {
		return nil, false
	}
	return o.BusinessLic, true
}

// HasBusinessLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasBusinessLic() bool {
	if o != nil && !IsNil(o.BusinessLic) {
		return true
	}

	return false
}

// SetBusinessLic gets a reference to the given string and assigns it to the BusinessLic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetBusinessLic(v string) {
	o.BusinessLic = &v
}

// GetCityCode returns the CityCode field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetCityCode() string {
	if o == nil || IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

// GetCityCodeOk returns a tuple with the CityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetCityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

// HasCityCode returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasCityCode() bool {
	if o != nil && !IsNil(o.CityCode) {
		return true
	}

	return false
}

// SetCityCode gets a reference to the given string and assigns it to the CityCode field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetCityCode(v string) {
	o.CityCode = &v
}

// GetDistrictCode returns the DistrictCode field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetDistrictCode() string {
	if o == nil || IsNil(o.DistrictCode) {
		var ret string
		return ret
	}
	return *o.DistrictCode
}

// GetDistrictCodeOk returns a tuple with the DistrictCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetDistrictCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DistrictCode) {
		return nil, false
	}
	return o.DistrictCode, true
}

// HasDistrictCode returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasDistrictCode() bool {
	if o != nil && !IsNil(o.DistrictCode) {
		return true
	}

	return false
}

// SetDistrictCode gets a reference to the given string and assigns it to the DistrictCode field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetDistrictCode(v string) {
	o.DistrictCode = &v
}

// GetFoodBusinessLic returns the FoodBusinessLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodBusinessLic() string {
	if o == nil || IsNil(o.FoodBusinessLic) {
		var ret string
		return ret
	}
	return *o.FoodBusinessLic
}

// GetFoodBusinessLicOk returns a tuple with the FoodBusinessLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodBusinessLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodBusinessLic) {
		return nil, false
	}
	return o.FoodBusinessLic, true
}

// HasFoodBusinessLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasFoodBusinessLic() bool {
	if o != nil && !IsNil(o.FoodBusinessLic) {
		return true
	}

	return false
}

// SetFoodBusinessLic gets a reference to the given string and assigns it to the FoodBusinessLic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetFoodBusinessLic(v string) {
	o.FoodBusinessLic = &v
}

// GetFoodCirculateLic returns the FoodCirculateLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodCirculateLic() string {
	if o == nil || IsNil(o.FoodCirculateLic) {
		var ret string
		return ret
	}
	return *o.FoodCirculateLic
}

// GetFoodCirculateLicOk returns a tuple with the FoodCirculateLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodCirculateLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodCirculateLic) {
		return nil, false
	}
	return o.FoodCirculateLic, true
}

// HasFoodCirculateLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasFoodCirculateLic() bool {
	if o != nil && !IsNil(o.FoodCirculateLic) {
		return true
	}

	return false
}

// SetFoodCirculateLic gets a reference to the given string and assigns it to the FoodCirculateLic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetFoodCirculateLic(v string) {
	o.FoodCirculateLic = &v
}

// GetFoodHealthLic returns the FoodHealthLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodHealthLic() string {
	if o == nil || IsNil(o.FoodHealthLic) {
		var ret string
		return ret
	}
	return *o.FoodHealthLic
}

// GetFoodHealthLicOk returns a tuple with the FoodHealthLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodHealthLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodHealthLic) {
		return nil, false
	}
	return o.FoodHealthLic, true
}

// HasFoodHealthLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasFoodHealthLic() bool {
	if o != nil && !IsNil(o.FoodHealthLic) {
		return true
	}

	return false
}

// SetFoodHealthLic gets a reference to the given string and assigns it to the FoodHealthLic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetFoodHealthLic(v string) {
	o.FoodHealthLic = &v
}

// GetFoodProductionLic returns the FoodProductionLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodProductionLic() string {
	if o == nil || IsNil(o.FoodProductionLic) {
		var ret string
		return ret
	}
	return *o.FoodProductionLic
}

// GetFoodProductionLicOk returns a tuple with the FoodProductionLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodProductionLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodProductionLic) {
		return nil, false
	}
	return o.FoodProductionLic, true
}

// HasFoodProductionLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasFoodProductionLic() bool {
	if o != nil && !IsNil(o.FoodProductionLic) {
		return true
	}

	return false
}

// SetFoodProductionLic gets a reference to the given string and assigns it to the FoodProductionLic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetFoodProductionLic(v string) {
	o.FoodProductionLic = &v
}

// GetFoodServiceLic returns the FoodServiceLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodServiceLic() string {
	if o == nil || IsNil(o.FoodServiceLic) {
		var ret string
		return ret
	}
	return *o.FoodServiceLic
}

// GetFoodServiceLicOk returns a tuple with the FoodServiceLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetFoodServiceLicOk() (*string, bool) {
	if o == nil || IsNil(o.FoodServiceLic) {
		return nil, false
	}
	return o.FoodServiceLic, true
}

// HasFoodServiceLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasFoodServiceLic() bool {
	if o != nil && !IsNil(o.FoodServiceLic) {
		return true
	}

	return false
}

// SetFoodServiceLic gets a reference to the given string and assigns it to the FoodServiceLic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetFoodServiceLic(v string) {
	o.FoodServiceLic = &v
}

// GetIndoorPic returns the IndoorPic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetIndoorPic() string {
	if o == nil || IsNil(o.IndoorPic) {
		var ret string
		return ret
	}
	return *o.IndoorPic
}

// GetIndoorPicOk returns a tuple with the IndoorPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetIndoorPicOk() (*string, bool) {
	if o == nil || IsNil(o.IndoorPic) {
		return nil, false
	}
	return o.IndoorPic, true
}

// HasIndoorPic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasIndoorPic() bool {
	if o != nil && !IsNil(o.IndoorPic) {
		return true
	}

	return false
}

// SetIndoorPic gets a reference to the given string and assigns it to the IndoorPic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetIndoorPic(v string) {
	o.IndoorPic = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetProvinceCode returns the ProvinceCode field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetProvinceCode() string {
	if o == nil || IsNil(o.ProvinceCode) {
		var ret string
		return ret
	}
	return *o.ProvinceCode
}

// GetProvinceCodeOk returns a tuple with the ProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetProvinceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceCode) {
		return nil, false
	}
	return o.ProvinceCode, true
}

// HasProvinceCode returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasProvinceCode() bool {
	if o != nil && !IsNil(o.ProvinceCode) {
		return true
	}

	return false
}

// SetProvinceCode gets a reference to the given string and assigns it to the ProvinceCode field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetProvinceCode(v string) {
	o.ProvinceCode = &v
}

// GetShopEntrancePic returns the ShopEntrancePic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetShopEntrancePic() string {
	if o == nil || IsNil(o.ShopEntrancePic) {
		var ret string
		return ret
	}
	return *o.ShopEntrancePic
}

// GetShopEntrancePicOk returns a tuple with the ShopEntrancePic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetShopEntrancePicOk() (*string, bool) {
	if o == nil || IsNil(o.ShopEntrancePic) {
		return nil, false
	}
	return o.ShopEntrancePic, true
}

// HasShopEntrancePic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasShopEntrancePic() bool {
	if o != nil && !IsNil(o.ShopEntrancePic) {
		return true
	}

	return false
}

// SetShopEntrancePic gets a reference to the given string and assigns it to the ShopEntrancePic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetShopEntrancePic(v string) {
	o.ShopEntrancePic = &v
}

// GetTobaccoLic returns the TobaccoLic field value if set, zero value otherwise.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetTobaccoLic() string {
	if o == nil || IsNil(o.TobaccoLic) {
		var ret string
		return ret
	}
	return *o.TobaccoLic
}

// GetTobaccoLicOk returns a tuple with the TobaccoLic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) GetTobaccoLicOk() (*string, bool) {
	if o == nil || IsNil(o.TobaccoLic) {
		return nil, false
	}
	return o.TobaccoLic, true
}

// HasTobaccoLic returns a boolean if a field has been set.
func (o *AlipayOpenSpBlueseaactivityModifyModel) HasTobaccoLic() bool {
	if o != nil && !IsNil(o.TobaccoLic) {
		return true
	}

	return false
}

// SetTobaccoLic gets a reference to the given string and assigns it to the TobaccoLic field.
func (o *AlipayOpenSpBlueseaactivityModifyModel) SetTobaccoLic(v string) {
	o.TobaccoLic = &v
}

func (o AlipayOpenSpBlueseaactivityModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSpBlueseaactivityModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BusinessLic) {
		toSerialize["business_lic"] = o.BusinessLic
	}
	if !IsNil(o.CityCode) {
		toSerialize["city_code"] = o.CityCode
	}
	if !IsNil(o.DistrictCode) {
		toSerialize["district_code"] = o.DistrictCode
	}
	if !IsNil(o.FoodBusinessLic) {
		toSerialize["food_business_lic"] = o.FoodBusinessLic
	}
	if !IsNil(o.FoodCirculateLic) {
		toSerialize["food_circulate_lic"] = o.FoodCirculateLic
	}
	if !IsNil(o.FoodHealthLic) {
		toSerialize["food_health_lic"] = o.FoodHealthLic
	}
	if !IsNil(o.FoodProductionLic) {
		toSerialize["food_production_lic"] = o.FoodProductionLic
	}
	if !IsNil(o.FoodServiceLic) {
		toSerialize["food_service_lic"] = o.FoodServiceLic
	}
	if !IsNil(o.IndoorPic) {
		toSerialize["indoor_pic"] = o.IndoorPic
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.ProvinceCode) {
		toSerialize["province_code"] = o.ProvinceCode
	}
	if !IsNil(o.ShopEntrancePic) {
		toSerialize["shop_entrance_pic"] = o.ShopEntrancePic
	}
	if !IsNil(o.TobaccoLic) {
		toSerialize["tobacco_lic"] = o.TobaccoLic
	}
	return toSerialize, nil
}

type NullableAlipayOpenSpBlueseaactivityModifyModel struct {
	value *AlipayOpenSpBlueseaactivityModifyModel
	isSet bool
}

func (v NullableAlipayOpenSpBlueseaactivityModifyModel) Get() *AlipayOpenSpBlueseaactivityModifyModel {
	return v.value
}

func (v *NullableAlipayOpenSpBlueseaactivityModifyModel) Set(val *AlipayOpenSpBlueseaactivityModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpBlueseaactivityModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpBlueseaactivityModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpBlueseaactivityModifyModel(val *AlipayOpenSpBlueseaactivityModifyModel) *NullableAlipayOpenSpBlueseaactivityModifyModel {
	return &NullableAlipayOpenSpBlueseaactivityModifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSpBlueseaactivityModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpBlueseaactivityModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


