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

// checks if the LapsPendingRotation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LapsPendingRotation{}

// LapsPendingRotation struct for LapsPendingRotation
type LapsPendingRotation struct {
	LapsUser *LapsUserV2 `json:"lapsUser,omitempty"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
}

// NewLapsPendingRotation instantiates a new LapsPendingRotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLapsPendingRotation() *LapsPendingRotation {
	this := LapsPendingRotation{}
	return &this
}

// NewLapsPendingRotationWithDefaults instantiates a new LapsPendingRotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLapsPendingRotationWithDefaults() *LapsPendingRotation {
	this := LapsPendingRotation{}
	return &this
}

// GetLapsUser returns the LapsUser field value if set, zero value otherwise.
func (o *LapsPendingRotation) GetLapsUser() LapsUserV2 {
	if o == nil || IsNil(o.LapsUser) {
		var ret LapsUserV2
		return ret
	}
	return *o.LapsUser
}

// GetLapsUserOk returns a tuple with the LapsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsPendingRotation) GetLapsUserOk() (*LapsUserV2, bool) {
	if o == nil || IsNil(o.LapsUser) {
		return nil, false
	}
	return o.LapsUser, true
}

// HasLapsUser returns a boolean if a field has been set.
func (o *LapsPendingRotation) HasLapsUser() bool {
	if o != nil && !IsNil(o.LapsUser) {
		return true
	}

	return false
}

// SetLapsUser gets a reference to the given LapsUserV2 and assigns it to the LapsUser field.
func (o *LapsPendingRotation) SetLapsUser(v LapsUserV2) {
	o.LapsUser = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *LapsPendingRotation) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsPendingRotation) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *LapsPendingRotation) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *LapsPendingRotation) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

func (o LapsPendingRotation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LapsPendingRotation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LapsUser) {
		toSerialize["lapsUser"] = o.LapsUser
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	return toSerialize, nil
}

type NullableLapsPendingRotation struct {
	value *LapsPendingRotation
	isSet bool
}

func (v NullableLapsPendingRotation) Get() *LapsPendingRotation {
	return v.value
}

func (v *NullableLapsPendingRotation) Set(val *LapsPendingRotation) {
	v.value = val
	v.isSet = true
}

func (v NullableLapsPendingRotation) IsSet() bool {
	return v.isSet
}

func (v *NullableLapsPendingRotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLapsPendingRotation(val *LapsPendingRotation) *NullableLapsPendingRotation {
	return &NullableLapsPendingRotation{value: val, isSet: true}
}

func (v NullableLapsPendingRotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLapsPendingRotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

