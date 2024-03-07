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

// checks if the SsoFailoverData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoFailoverData{}

// SsoFailoverData struct for SsoFailoverData
type SsoFailoverData struct {
	FailoverUrl *string `json:"failoverUrl,omitempty"`
	// Generation time of failover key
	GenerationTime *int64 `json:"generationTime,omitempty"`
}

// NewSsoFailoverData instantiates a new SsoFailoverData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoFailoverData() *SsoFailoverData {
	this := SsoFailoverData{}
	return &this
}

// NewSsoFailoverDataWithDefaults instantiates a new SsoFailoverData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoFailoverDataWithDefaults() *SsoFailoverData {
	this := SsoFailoverData{}
	return &this
}

// GetFailoverUrl returns the FailoverUrl field value if set, zero value otherwise.
func (o *SsoFailoverData) GetFailoverUrl() string {
	if o == nil || IsNil(o.FailoverUrl) {
		var ret string
		return ret
	}
	return *o.FailoverUrl
}

// GetFailoverUrlOk returns a tuple with the FailoverUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoFailoverData) GetFailoverUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FailoverUrl) {
		return nil, false
	}
	return o.FailoverUrl, true
}

// HasFailoverUrl returns a boolean if a field has been set.
func (o *SsoFailoverData) HasFailoverUrl() bool {
	if o != nil && !IsNil(o.FailoverUrl) {
		return true
	}

	return false
}

// SetFailoverUrl gets a reference to the given string and assigns it to the FailoverUrl field.
func (o *SsoFailoverData) SetFailoverUrl(v string) {
	o.FailoverUrl = &v
}

// GetGenerationTime returns the GenerationTime field value if set, zero value otherwise.
func (o *SsoFailoverData) GetGenerationTime() int64 {
	if o == nil || IsNil(o.GenerationTime) {
		var ret int64
		return ret
	}
	return *o.GenerationTime
}

// GetGenerationTimeOk returns a tuple with the GenerationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoFailoverData) GetGenerationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.GenerationTime) {
		return nil, false
	}
	return o.GenerationTime, true
}

// HasGenerationTime returns a boolean if a field has been set.
func (o *SsoFailoverData) HasGenerationTime() bool {
	if o != nil && !IsNil(o.GenerationTime) {
		return true
	}

	return false
}

// SetGenerationTime gets a reference to the given int64 and assigns it to the GenerationTime field.
func (o *SsoFailoverData) SetGenerationTime(v int64) {
	o.GenerationTime = &v
}

func (o SsoFailoverData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoFailoverData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailoverUrl) {
		toSerialize["failoverUrl"] = o.FailoverUrl
	}
	if !IsNil(o.GenerationTime) {
		toSerialize["generationTime"] = o.GenerationTime
	}
	return toSerialize, nil
}

type NullableSsoFailoverData struct {
	value *SsoFailoverData
	isSet bool
}

func (v NullableSsoFailoverData) Get() *SsoFailoverData {
	return v.value
}

func (v *NullableSsoFailoverData) Set(val *SsoFailoverData) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoFailoverData) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoFailoverData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoFailoverData(val *SsoFailoverData) *NullableSsoFailoverData {
	return &NullableSsoFailoverData{value: val, isSet: true}
}

func (v NullableSsoFailoverData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoFailoverData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

