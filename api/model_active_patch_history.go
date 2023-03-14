/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// ActivePatchHistory struct for ActivePatchHistory
type ActivePatchHistory struct {
	PatchId *int32 `json:"patchId,omitempty"`
	PatchHistoryId *int32 `json:"patchHistoryId,omitempty"`
	DeviceId *int32 `json:"deviceId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	BuildingId *int32 `json:"buildingId,omitempty"`
	BuildingName *string `json:"buildingName,omitempty"`
	DepartmentId *int32 `json:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty"`
	SiteId *int32 `json:"siteId,omitempty"`
	SiteName *string `json:"siteName,omitempty"`
	Username *string `json:"username,omitempty"`
	OsVersion *string `json:"osVersion,omitempty"`
	LastCheckIn *time.Time `json:"lastCheckIn,omitempty"`
	InstalledVersion *string `json:"installedVersion,omitempty"`
}

// NewActivePatchHistory instantiates a new ActivePatchHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivePatchHistory() *ActivePatchHistory {
	this := ActivePatchHistory{}
	return &this
}

// NewActivePatchHistoryWithDefaults instantiates a new ActivePatchHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivePatchHistoryWithDefaults() *ActivePatchHistory {
	this := ActivePatchHistory{}
	return &this
}

// GetPatchId returns the PatchId field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetPatchId() int32 {
	if o == nil || o.PatchId == nil {
		var ret int32
		return ret
	}
	return *o.PatchId
}

// GetPatchIdOk returns a tuple with the PatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetPatchIdOk() (*int32, bool) {
	if o == nil || o.PatchId == nil {
		return nil, false
	}
	return o.PatchId, true
}

// HasPatchId returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasPatchId() bool {
	if o != nil && o.PatchId != nil {
		return true
	}

	return false
}

// SetPatchId gets a reference to the given int32 and assigns it to the PatchId field.
func (o *ActivePatchHistory) SetPatchId(v int32) {
	o.PatchId = &v
}

// GetPatchHistoryId returns the PatchHistoryId field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetPatchHistoryId() int32 {
	if o == nil || o.PatchHistoryId == nil {
		var ret int32
		return ret
	}
	return *o.PatchHistoryId
}

// GetPatchHistoryIdOk returns a tuple with the PatchHistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetPatchHistoryIdOk() (*int32, bool) {
	if o == nil || o.PatchHistoryId == nil {
		return nil, false
	}
	return o.PatchHistoryId, true
}

// HasPatchHistoryId returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasPatchHistoryId() bool {
	if o != nil && o.PatchHistoryId != nil {
		return true
	}

	return false
}

// SetPatchHistoryId gets a reference to the given int32 and assigns it to the PatchHistoryId field.
func (o *ActivePatchHistory) SetPatchHistoryId(v int32) {
	o.PatchHistoryId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetDeviceId() int32 {
	if o == nil || o.DeviceId == nil {
		var ret int32
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetDeviceIdOk() (*int32, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int32 and assigns it to the DeviceId field.
func (o *ActivePatchHistory) SetDeviceId(v int32) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *ActivePatchHistory) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetBuildingId returns the BuildingId field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetBuildingId() int32 {
	if o == nil || o.BuildingId == nil {
		var ret int32
		return ret
	}
	return *o.BuildingId
}

// GetBuildingIdOk returns a tuple with the BuildingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetBuildingIdOk() (*int32, bool) {
	if o == nil || o.BuildingId == nil {
		return nil, false
	}
	return o.BuildingId, true
}

// HasBuildingId returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasBuildingId() bool {
	if o != nil && o.BuildingId != nil {
		return true
	}

	return false
}

// SetBuildingId gets a reference to the given int32 and assigns it to the BuildingId field.
func (o *ActivePatchHistory) SetBuildingId(v int32) {
	o.BuildingId = &v
}

// GetBuildingName returns the BuildingName field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetBuildingName() string {
	if o == nil || o.BuildingName == nil {
		var ret string
		return ret
	}
	return *o.BuildingName
}

// GetBuildingNameOk returns a tuple with the BuildingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetBuildingNameOk() (*string, bool) {
	if o == nil || o.BuildingName == nil {
		return nil, false
	}
	return o.BuildingName, true
}

// HasBuildingName returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasBuildingName() bool {
	if o != nil && o.BuildingName != nil {
		return true
	}

	return false
}

// SetBuildingName gets a reference to the given string and assigns it to the BuildingName field.
func (o *ActivePatchHistory) SetBuildingName(v string) {
	o.BuildingName = &v
}

// GetDepartmentId returns the DepartmentId field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetDepartmentId() int32 {
	if o == nil || o.DepartmentId == nil {
		var ret int32
		return ret
	}
	return *o.DepartmentId
}

// GetDepartmentIdOk returns a tuple with the DepartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetDepartmentIdOk() (*int32, bool) {
	if o == nil || o.DepartmentId == nil {
		return nil, false
	}
	return o.DepartmentId, true
}

// HasDepartmentId returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasDepartmentId() bool {
	if o != nil && o.DepartmentId != nil {
		return true
	}

	return false
}

// SetDepartmentId gets a reference to the given int32 and assigns it to the DepartmentId field.
func (o *ActivePatchHistory) SetDepartmentId(v int32) {
	o.DepartmentId = &v
}

// GetDepartmentName returns the DepartmentName field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetDepartmentName() string {
	if o == nil || o.DepartmentName == nil {
		var ret string
		return ret
	}
	return *o.DepartmentName
}

// GetDepartmentNameOk returns a tuple with the DepartmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetDepartmentNameOk() (*string, bool) {
	if o == nil || o.DepartmentName == nil {
		return nil, false
	}
	return o.DepartmentName, true
}

// HasDepartmentName returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasDepartmentName() bool {
	if o != nil && o.DepartmentName != nil {
		return true
	}

	return false
}

// SetDepartmentName gets a reference to the given string and assigns it to the DepartmentName field.
func (o *ActivePatchHistory) SetDepartmentName(v string) {
	o.DepartmentName = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetSiteId() int32 {
	if o == nil || o.SiteId == nil {
		var ret int32
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetSiteIdOk() (*int32, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given int32 and assigns it to the SiteId field.
func (o *ActivePatchHistory) SetSiteId(v int32) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *ActivePatchHistory) SetSiteName(v string) {
	o.SiteName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ActivePatchHistory) SetUsername(v string) {
	o.Username = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *ActivePatchHistory) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetLastCheckIn returns the LastCheckIn field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetLastCheckIn() time.Time {
	if o == nil || o.LastCheckIn == nil {
		var ret time.Time
		return ret
	}
	return *o.LastCheckIn
}

// GetLastCheckInOk returns a tuple with the LastCheckIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetLastCheckInOk() (*time.Time, bool) {
	if o == nil || o.LastCheckIn == nil {
		return nil, false
	}
	return o.LastCheckIn, true
}

// HasLastCheckIn returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasLastCheckIn() bool {
	if o != nil && o.LastCheckIn != nil {
		return true
	}

	return false
}

// SetLastCheckIn gets a reference to the given time.Time and assigns it to the LastCheckIn field.
func (o *ActivePatchHistory) SetLastCheckIn(v time.Time) {
	o.LastCheckIn = &v
}

// GetInstalledVersion returns the InstalledVersion field value if set, zero value otherwise.
func (o *ActivePatchHistory) GetInstalledVersion() string {
	if o == nil || o.InstalledVersion == nil {
		var ret string
		return ret
	}
	return *o.InstalledVersion
}

// GetInstalledVersionOk returns a tuple with the InstalledVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivePatchHistory) GetInstalledVersionOk() (*string, bool) {
	if o == nil || o.InstalledVersion == nil {
		return nil, false
	}
	return o.InstalledVersion, true
}

// HasInstalledVersion returns a boolean if a field has been set.
func (o *ActivePatchHistory) HasInstalledVersion() bool {
	if o != nil && o.InstalledVersion != nil {
		return true
	}

	return false
}

// SetInstalledVersion gets a reference to the given string and assigns it to the InstalledVersion field.
func (o *ActivePatchHistory) SetInstalledVersion(v string) {
	o.InstalledVersion = &v
}

func (o ActivePatchHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PatchId != nil {
		toSerialize["patchId"] = o.PatchId
	}
	if o.PatchHistoryId != nil {
		toSerialize["patchHistoryId"] = o.PatchHistoryId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.DeviceName != nil {
		toSerialize["deviceName"] = o.DeviceName
	}
	if o.BuildingId != nil {
		toSerialize["buildingId"] = o.BuildingId
	}
	if o.BuildingName != nil {
		toSerialize["buildingName"] = o.BuildingName
	}
	if o.DepartmentId != nil {
		toSerialize["departmentId"] = o.DepartmentId
	}
	if o.DepartmentName != nil {
		toSerialize["departmentName"] = o.DepartmentName
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	if o.SiteName != nil {
		toSerialize["siteName"] = o.SiteName
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.OsVersion != nil {
		toSerialize["osVersion"] = o.OsVersion
	}
	if o.LastCheckIn != nil {
		toSerialize["lastCheckIn"] = o.LastCheckIn
	}
	if o.InstalledVersion != nil {
		toSerialize["installedVersion"] = o.InstalledVersion
	}
	return json.Marshal(toSerialize)
}

type NullableActivePatchHistory struct {
	value *ActivePatchHistory
	isSet bool
}

func (v NullableActivePatchHistory) Get() *ActivePatchHistory {
	return v.value
}

func (v *NullableActivePatchHistory) Set(val *ActivePatchHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableActivePatchHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableActivePatchHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivePatchHistory(val *ActivePatchHistory) *NullableActivePatchHistory {
	return &NullableActivePatchHistory{value: val, isSet: true}
}

func (v NullableActivePatchHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivePatchHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

