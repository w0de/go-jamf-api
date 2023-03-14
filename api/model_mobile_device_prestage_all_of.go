/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MobileDevicePrestageAllOf struct for MobileDevicePrestageAllOf
type MobileDevicePrestageAllOf struct {
	IsAllowPairing bool `json:"isAllowPairing"`
	IsMultiUser bool `json:"isMultiUser"`
	IsSupervised bool `json:"isSupervised"`
	MaximumSharedAccounts int32 `json:"maximumSharedAccounts"`
	IsAutoAdvanceSetup bool `json:"isAutoAdvanceSetup"`
	IsConfigureDeviceBeforeSetupAssistant bool `json:"isConfigureDeviceBeforeSetupAssistant"`
	Language *string `json:"language,omitempty"`
	Region *string `json:"region,omitempty"`
	Names *MobileDevicePrestageNames `json:"names,omitempty"`
}

// NewMobileDevicePrestageAllOf instantiates a new MobileDevicePrestageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDevicePrestageAllOf(isAllowPairing bool, isMultiUser bool, isSupervised bool, maximumSharedAccounts int32, isAutoAdvanceSetup bool, isConfigureDeviceBeforeSetupAssistant bool) *MobileDevicePrestageAllOf {
	this := MobileDevicePrestageAllOf{}
	this.IsAllowPairing = isAllowPairing
	this.IsMultiUser = isMultiUser
	this.IsSupervised = isSupervised
	this.MaximumSharedAccounts = maximumSharedAccounts
	this.IsAutoAdvanceSetup = isAutoAdvanceSetup
	this.IsConfigureDeviceBeforeSetupAssistant = isConfigureDeviceBeforeSetupAssistant
	return &this
}

// NewMobileDevicePrestageAllOfWithDefaults instantiates a new MobileDevicePrestageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDevicePrestageAllOfWithDefaults() *MobileDevicePrestageAllOf {
	this := MobileDevicePrestageAllOf{}
	return &this
}

// GetIsAllowPairing returns the IsAllowPairing field value
func (o *MobileDevicePrestageAllOf) GetIsAllowPairing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAllowPairing
}

// GetIsAllowPairingOk returns a tuple with the IsAllowPairing field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetIsAllowPairingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAllowPairing, true
}

// SetIsAllowPairing sets field value
func (o *MobileDevicePrestageAllOf) SetIsAllowPairing(v bool) {
	o.IsAllowPairing = v
}

// GetIsMultiUser returns the IsMultiUser field value
func (o *MobileDevicePrestageAllOf) GetIsMultiUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMultiUser
}

// GetIsMultiUserOk returns a tuple with the IsMultiUser field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetIsMultiUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMultiUser, true
}

// SetIsMultiUser sets field value
func (o *MobileDevicePrestageAllOf) SetIsMultiUser(v bool) {
	o.IsMultiUser = v
}

// GetIsSupervised returns the IsSupervised field value
func (o *MobileDevicePrestageAllOf) GetIsSupervised() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSupervised
}

// GetIsSupervisedOk returns a tuple with the IsSupervised field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetIsSupervisedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSupervised, true
}

// SetIsSupervised sets field value
func (o *MobileDevicePrestageAllOf) SetIsSupervised(v bool) {
	o.IsSupervised = v
}

// GetMaximumSharedAccounts returns the MaximumSharedAccounts field value
func (o *MobileDevicePrestageAllOf) GetMaximumSharedAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumSharedAccounts
}

// GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetMaximumSharedAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumSharedAccounts, true
}

// SetMaximumSharedAccounts sets field value
func (o *MobileDevicePrestageAllOf) SetMaximumSharedAccounts(v int32) {
	o.MaximumSharedAccounts = v
}

// GetIsAutoAdvanceSetup returns the IsAutoAdvanceSetup field value
func (o *MobileDevicePrestageAllOf) GetIsAutoAdvanceSetup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutoAdvanceSetup
}

// GetIsAutoAdvanceSetupOk returns a tuple with the IsAutoAdvanceSetup field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetIsAutoAdvanceSetupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutoAdvanceSetup, true
}

// SetIsAutoAdvanceSetup sets field value
func (o *MobileDevicePrestageAllOf) SetIsAutoAdvanceSetup(v bool) {
	o.IsAutoAdvanceSetup = v
}

// GetIsConfigureDeviceBeforeSetupAssistant returns the IsConfigureDeviceBeforeSetupAssistant field value
func (o *MobileDevicePrestageAllOf) GetIsConfigureDeviceBeforeSetupAssistant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsConfigureDeviceBeforeSetupAssistant
}

// GetIsConfigureDeviceBeforeSetupAssistantOk returns a tuple with the IsConfigureDeviceBeforeSetupAssistant field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetIsConfigureDeviceBeforeSetupAssistantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsConfigureDeviceBeforeSetupAssistant, true
}

// SetIsConfigureDeviceBeforeSetupAssistant sets field value
func (o *MobileDevicePrestageAllOf) SetIsConfigureDeviceBeforeSetupAssistant(v bool) {
	o.IsConfigureDeviceBeforeSetupAssistant = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MobileDevicePrestageAllOf) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MobileDevicePrestageAllOf) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MobileDevicePrestageAllOf) SetLanguage(v string) {
	o.Language = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *MobileDevicePrestageAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *MobileDevicePrestageAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *MobileDevicePrestageAllOf) SetRegion(v string) {
	o.Region = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *MobileDevicePrestageAllOf) GetNames() MobileDevicePrestageNames {
	if o == nil || o.Names == nil {
		var ret MobileDevicePrestageNames
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageAllOf) GetNamesOk() (*MobileDevicePrestageNames, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *MobileDevicePrestageAllOf) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given MobileDevicePrestageNames and assigns it to the Names field.
func (o *MobileDevicePrestageAllOf) SetNames(v MobileDevicePrestageNames) {
	o.Names = &v
}

func (o MobileDevicePrestageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isAllowPairing"] = o.IsAllowPairing
	}
	if true {
		toSerialize["isMultiUser"] = o.IsMultiUser
	}
	if true {
		toSerialize["isSupervised"] = o.IsSupervised
	}
	if true {
		toSerialize["maximumSharedAccounts"] = o.MaximumSharedAccounts
	}
	if true {
		toSerialize["isAutoAdvanceSetup"] = o.IsAutoAdvanceSetup
	}
	if true {
		toSerialize["isConfigureDeviceBeforeSetupAssistant"] = o.IsConfigureDeviceBeforeSetupAssistant
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	return json.Marshal(toSerialize)
}

type NullableMobileDevicePrestageAllOf struct {
	value *MobileDevicePrestageAllOf
	isSet bool
}

func (v NullableMobileDevicePrestageAllOf) Get() *MobileDevicePrestageAllOf {
	return v.value
}

func (v *NullableMobileDevicePrestageAllOf) Set(val *MobileDevicePrestageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDevicePrestageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDevicePrestageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDevicePrestageAllOf(val *MobileDevicePrestageAllOf) *NullableMobileDevicePrestageAllOf {
	return &NullableMobileDevicePrestageAllOf{value: val, isSet: true}
}

func (v NullableMobileDevicePrestageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDevicePrestageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

