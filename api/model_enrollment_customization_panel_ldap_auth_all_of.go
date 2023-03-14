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

// EnrollmentCustomizationPanelLdapAuthAllOf struct for EnrollmentCustomizationPanelLdapAuthAllOf
type EnrollmentCustomizationPanelLdapAuthAllOf struct {
	UsernameLabel string `json:"usernameLabel"`
	PasswordLabel string `json:"passwordLabel"`
	Title string `json:"title"`
	BackButtonText string `json:"backButtonText"`
	ContinueButtonText string `json:"continueButtonText"`
	LdapGroupAccess []EnrollmentCustomizationLdapGroupAccess `json:"ldapGroupAccess,omitempty"`
}

// NewEnrollmentCustomizationPanelLdapAuthAllOf instantiates a new EnrollmentCustomizationPanelLdapAuthAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentCustomizationPanelLdapAuthAllOf(usernameLabel string, passwordLabel string, title string, backButtonText string, continueButtonText string) *EnrollmentCustomizationPanelLdapAuthAllOf {
	this := EnrollmentCustomizationPanelLdapAuthAllOf{}
	this.UsernameLabel = usernameLabel
	this.PasswordLabel = passwordLabel
	this.Title = title
	this.BackButtonText = backButtonText
	this.ContinueButtonText = continueButtonText
	return &this
}

// NewEnrollmentCustomizationPanelLdapAuthAllOfWithDefaults instantiates a new EnrollmentCustomizationPanelLdapAuthAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentCustomizationPanelLdapAuthAllOfWithDefaults() *EnrollmentCustomizationPanelLdapAuthAllOf {
	this := EnrollmentCustomizationPanelLdapAuthAllOf{}
	return &this
}

// GetUsernameLabel returns the UsernameLabel field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetUsernameLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameLabel
}

// GetUsernameLabelOk returns a tuple with the UsernameLabel field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetUsernameLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameLabel, true
}

// SetUsernameLabel sets field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) SetUsernameLabel(v string) {
	o.UsernameLabel = v
}

// GetPasswordLabel returns the PasswordLabel field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetPasswordLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordLabel
}

// GetPasswordLabelOk returns a tuple with the PasswordLabel field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetPasswordLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordLabel, true
}

// SetPasswordLabel sets field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) SetPasswordLabel(v string) {
	o.PasswordLabel = v
}

// GetTitle returns the Title field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) SetTitle(v string) {
	o.Title = v
}

// GetBackButtonText returns the BackButtonText field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetBackButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackButtonText
}

// GetBackButtonTextOk returns a tuple with the BackButtonText field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetBackButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackButtonText, true
}

// SetBackButtonText sets field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) SetBackButtonText(v string) {
	o.BackButtonText = v
}

// GetContinueButtonText returns the ContinueButtonText field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetContinueButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContinueButtonText
}

// GetContinueButtonTextOk returns a tuple with the ContinueButtonText field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetContinueButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinueButtonText, true
}

// SetContinueButtonText sets field value
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) SetContinueButtonText(v string) {
	o.ContinueButtonText = v
}

// GetLdapGroupAccess returns the LdapGroupAccess field value if set, zero value otherwise.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetLdapGroupAccess() []EnrollmentCustomizationLdapGroupAccess {
	if o == nil || o.LdapGroupAccess == nil {
		var ret []EnrollmentCustomizationLdapGroupAccess
		return ret
	}
	return o.LdapGroupAccess
}

// GetLdapGroupAccessOk returns a tuple with the LdapGroupAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) GetLdapGroupAccessOk() ([]EnrollmentCustomizationLdapGroupAccess, bool) {
	if o == nil || o.LdapGroupAccess == nil {
		return nil, false
	}
	return o.LdapGroupAccess, true
}

// HasLdapGroupAccess returns a boolean if a field has been set.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) HasLdapGroupAccess() bool {
	if o != nil && o.LdapGroupAccess != nil {
		return true
	}

	return false
}

// SetLdapGroupAccess gets a reference to the given []EnrollmentCustomizationLdapGroupAccess and assigns it to the LdapGroupAccess field.
func (o *EnrollmentCustomizationPanelLdapAuthAllOf) SetLdapGroupAccess(v []EnrollmentCustomizationLdapGroupAccess) {
	o.LdapGroupAccess = v
}

func (o EnrollmentCustomizationPanelLdapAuthAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["usernameLabel"] = o.UsernameLabel
	}
	if true {
		toSerialize["passwordLabel"] = o.PasswordLabel
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["backButtonText"] = o.BackButtonText
	}
	if true {
		toSerialize["continueButtonText"] = o.ContinueButtonText
	}
	if o.LdapGroupAccess != nil {
		toSerialize["ldapGroupAccess"] = o.LdapGroupAccess
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentCustomizationPanelLdapAuthAllOf struct {
	value *EnrollmentCustomizationPanelLdapAuthAllOf
	isSet bool
}

func (v NullableEnrollmentCustomizationPanelLdapAuthAllOf) Get() *EnrollmentCustomizationPanelLdapAuthAllOf {
	return v.value
}

func (v *NullableEnrollmentCustomizationPanelLdapAuthAllOf) Set(val *EnrollmentCustomizationPanelLdapAuthAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCustomizationPanelLdapAuthAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCustomizationPanelLdapAuthAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCustomizationPanelLdapAuthAllOf(val *EnrollmentCustomizationPanelLdapAuthAllOf) *NullableEnrollmentCustomizationPanelLdapAuthAllOf {
	return &NullableEnrollmentCustomizationPanelLdapAuthAllOf{value: val, isSet: true}
}

func (v NullableEnrollmentCustomizationPanelLdapAuthAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCustomizationPanelLdapAuthAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

