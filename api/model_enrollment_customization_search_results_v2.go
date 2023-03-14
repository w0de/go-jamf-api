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

// EnrollmentCustomizationSearchResultsV2 struct for EnrollmentCustomizationSearchResultsV2
type EnrollmentCustomizationSearchResultsV2 struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []EnrollmentCustomizationV2 `json:"results,omitempty"`
}

// NewEnrollmentCustomizationSearchResultsV2 instantiates a new EnrollmentCustomizationSearchResultsV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentCustomizationSearchResultsV2() *EnrollmentCustomizationSearchResultsV2 {
	this := EnrollmentCustomizationSearchResultsV2{}
	return &this
}

// NewEnrollmentCustomizationSearchResultsV2WithDefaults instantiates a new EnrollmentCustomizationSearchResultsV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentCustomizationSearchResultsV2WithDefaults() *EnrollmentCustomizationSearchResultsV2 {
	this := EnrollmentCustomizationSearchResultsV2{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *EnrollmentCustomizationSearchResultsV2) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationSearchResultsV2) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *EnrollmentCustomizationSearchResultsV2) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *EnrollmentCustomizationSearchResultsV2) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnrollmentCustomizationSearchResultsV2) GetResults() []EnrollmentCustomizationV2 {
	if o == nil || o.Results == nil {
		var ret []EnrollmentCustomizationV2
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationSearchResultsV2) GetResultsOk() ([]EnrollmentCustomizationV2, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnrollmentCustomizationSearchResultsV2) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnrollmentCustomizationV2 and assigns it to the Results field.
func (o *EnrollmentCustomizationSearchResultsV2) SetResults(v []EnrollmentCustomizationV2) {
	o.Results = v
}

func (o EnrollmentCustomizationSearchResultsV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentCustomizationSearchResultsV2 struct {
	value *EnrollmentCustomizationSearchResultsV2
	isSet bool
}

func (v NullableEnrollmentCustomizationSearchResultsV2) Get() *EnrollmentCustomizationSearchResultsV2 {
	return v.value
}

func (v *NullableEnrollmentCustomizationSearchResultsV2) Set(val *EnrollmentCustomizationSearchResultsV2) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCustomizationSearchResultsV2) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCustomizationSearchResultsV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCustomizationSearchResultsV2(val *EnrollmentCustomizationSearchResultsV2) *NullableEnrollmentCustomizationSearchResultsV2 {
	return &NullableEnrollmentCustomizationSearchResultsV2{value: val, isSet: true}
}

func (v NullableEnrollmentCustomizationSearchResultsV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCustomizationSearchResultsV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

