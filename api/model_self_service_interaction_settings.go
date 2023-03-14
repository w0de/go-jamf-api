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

// SelfServiceInteractionSettings object representation of Self Service settings regarding user interaction 
type SelfServiceInteractionSettings struct {
	// global Self Service setting for if notifications are on or off 
	NotificationsEnabled *bool `json:"notificationsEnabled,omitempty"`
	// whether users should be notified they need to approve organization's MDM profile 
	AlertUserApprovedMdm *bool `json:"alertUserApprovedMdm,omitempty"`
	// the default landing page in Self Service 
	DefaultLandingPage *string `json:"defaultLandingPage,omitempty"`
	// id for the default home category in Self Service 
	DefaultHomeCategoryId *int32 `json:"defaultHomeCategoryId,omitempty"`
	// renamed string for bookmarks if the admin wishes 
	BookmarksName string `json:"bookmarksName"`
}

// NewSelfServiceInteractionSettings instantiates a new SelfServiceInteractionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceInteractionSettings(bookmarksName string) *SelfServiceInteractionSettings {
	this := SelfServiceInteractionSettings{}
	var notificationsEnabled bool = false
	this.NotificationsEnabled = &notificationsEnabled
	var alertUserApprovedMdm bool = true
	this.AlertUserApprovedMdm = &alertUserApprovedMdm
	var defaultLandingPage string = "HOME"
	this.DefaultLandingPage = &defaultLandingPage
	var defaultHomeCategoryId int32 = -1
	this.DefaultHomeCategoryId = &defaultHomeCategoryId
	this.BookmarksName = bookmarksName
	return &this
}

// NewSelfServiceInteractionSettingsWithDefaults instantiates a new SelfServiceInteractionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceInteractionSettingsWithDefaults() *SelfServiceInteractionSettings {
	this := SelfServiceInteractionSettings{}
	var notificationsEnabled bool = false
	this.NotificationsEnabled = &notificationsEnabled
	var alertUserApprovedMdm bool = true
	this.AlertUserApprovedMdm = &alertUserApprovedMdm
	var defaultLandingPage string = "HOME"
	this.DefaultLandingPage = &defaultLandingPage
	var defaultHomeCategoryId int32 = -1
	this.DefaultHomeCategoryId = &defaultHomeCategoryId
	return &this
}

// GetNotificationsEnabled returns the NotificationsEnabled field value if set, zero value otherwise.
func (o *SelfServiceInteractionSettings) GetNotificationsEnabled() bool {
	if o == nil || o.NotificationsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NotificationsEnabled
}

// GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceInteractionSettings) GetNotificationsEnabledOk() (*bool, bool) {
	if o == nil || o.NotificationsEnabled == nil {
		return nil, false
	}
	return o.NotificationsEnabled, true
}

// HasNotificationsEnabled returns a boolean if a field has been set.
func (o *SelfServiceInteractionSettings) HasNotificationsEnabled() bool {
	if o != nil && o.NotificationsEnabled != nil {
		return true
	}

	return false
}

// SetNotificationsEnabled gets a reference to the given bool and assigns it to the NotificationsEnabled field.
func (o *SelfServiceInteractionSettings) SetNotificationsEnabled(v bool) {
	o.NotificationsEnabled = &v
}

// GetAlertUserApprovedMdm returns the AlertUserApprovedMdm field value if set, zero value otherwise.
func (o *SelfServiceInteractionSettings) GetAlertUserApprovedMdm() bool {
	if o == nil || o.AlertUserApprovedMdm == nil {
		var ret bool
		return ret
	}
	return *o.AlertUserApprovedMdm
}

// GetAlertUserApprovedMdmOk returns a tuple with the AlertUserApprovedMdm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceInteractionSettings) GetAlertUserApprovedMdmOk() (*bool, bool) {
	if o == nil || o.AlertUserApprovedMdm == nil {
		return nil, false
	}
	return o.AlertUserApprovedMdm, true
}

// HasAlertUserApprovedMdm returns a boolean if a field has been set.
func (o *SelfServiceInteractionSettings) HasAlertUserApprovedMdm() bool {
	if o != nil && o.AlertUserApprovedMdm != nil {
		return true
	}

	return false
}

// SetAlertUserApprovedMdm gets a reference to the given bool and assigns it to the AlertUserApprovedMdm field.
func (o *SelfServiceInteractionSettings) SetAlertUserApprovedMdm(v bool) {
	o.AlertUserApprovedMdm = &v
}

// GetDefaultLandingPage returns the DefaultLandingPage field value if set, zero value otherwise.
func (o *SelfServiceInteractionSettings) GetDefaultLandingPage() string {
	if o == nil || o.DefaultLandingPage == nil {
		var ret string
		return ret
	}
	return *o.DefaultLandingPage
}

// GetDefaultLandingPageOk returns a tuple with the DefaultLandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceInteractionSettings) GetDefaultLandingPageOk() (*string, bool) {
	if o == nil || o.DefaultLandingPage == nil {
		return nil, false
	}
	return o.DefaultLandingPage, true
}

// HasDefaultLandingPage returns a boolean if a field has been set.
func (o *SelfServiceInteractionSettings) HasDefaultLandingPage() bool {
	if o != nil && o.DefaultLandingPage != nil {
		return true
	}

	return false
}

// SetDefaultLandingPage gets a reference to the given string and assigns it to the DefaultLandingPage field.
func (o *SelfServiceInteractionSettings) SetDefaultLandingPage(v string) {
	o.DefaultLandingPage = &v
}

// GetDefaultHomeCategoryId returns the DefaultHomeCategoryId field value if set, zero value otherwise.
func (o *SelfServiceInteractionSettings) GetDefaultHomeCategoryId() int32 {
	if o == nil || o.DefaultHomeCategoryId == nil {
		var ret int32
		return ret
	}
	return *o.DefaultHomeCategoryId
}

// GetDefaultHomeCategoryIdOk returns a tuple with the DefaultHomeCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceInteractionSettings) GetDefaultHomeCategoryIdOk() (*int32, bool) {
	if o == nil || o.DefaultHomeCategoryId == nil {
		return nil, false
	}
	return o.DefaultHomeCategoryId, true
}

// HasDefaultHomeCategoryId returns a boolean if a field has been set.
func (o *SelfServiceInteractionSettings) HasDefaultHomeCategoryId() bool {
	if o != nil && o.DefaultHomeCategoryId != nil {
		return true
	}

	return false
}

// SetDefaultHomeCategoryId gets a reference to the given int32 and assigns it to the DefaultHomeCategoryId field.
func (o *SelfServiceInteractionSettings) SetDefaultHomeCategoryId(v int32) {
	o.DefaultHomeCategoryId = &v
}

// GetBookmarksName returns the BookmarksName field value
func (o *SelfServiceInteractionSettings) GetBookmarksName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BookmarksName
}

// GetBookmarksNameOk returns a tuple with the BookmarksName field value
// and a boolean to check if the value has been set.
func (o *SelfServiceInteractionSettings) GetBookmarksNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookmarksName, true
}

// SetBookmarksName sets field value
func (o *SelfServiceInteractionSettings) SetBookmarksName(v string) {
	o.BookmarksName = v
}

func (o SelfServiceInteractionSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationsEnabled != nil {
		toSerialize["notificationsEnabled"] = o.NotificationsEnabled
	}
	if o.AlertUserApprovedMdm != nil {
		toSerialize["alertUserApprovedMdm"] = o.AlertUserApprovedMdm
	}
	if o.DefaultLandingPage != nil {
		toSerialize["defaultLandingPage"] = o.DefaultLandingPage
	}
	if o.DefaultHomeCategoryId != nil {
		toSerialize["defaultHomeCategoryId"] = o.DefaultHomeCategoryId
	}
	if true {
		toSerialize["bookmarksName"] = o.BookmarksName
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceInteractionSettings struct {
	value *SelfServiceInteractionSettings
	isSet bool
}

func (v NullableSelfServiceInteractionSettings) Get() *SelfServiceInteractionSettings {
	return v.value
}

func (v *NullableSelfServiceInteractionSettings) Set(val *SelfServiceInteractionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceInteractionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceInteractionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceInteractionSettings(val *SelfServiceInteractionSettings) *NullableSelfServiceInteractionSettings {
	return &NullableSelfServiceInteractionSettings{value: val, isSet: true}
}

func (v NullableSelfServiceInteractionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceInteractionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

