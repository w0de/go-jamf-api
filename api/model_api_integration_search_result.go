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

// checks if the ApiIntegrationSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiIntegrationSearchResult{}

// ApiIntegrationSearchResult struct for ApiIntegrationSearchResult
type ApiIntegrationSearchResult struct {
	TotalCount int32 `json:"totalCount"`
	Results []ApiIntegrationResponse `json:"results"`
}

// NewApiIntegrationSearchResult instantiates a new ApiIntegrationSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIntegrationSearchResult(totalCount int32, results []ApiIntegrationResponse) *ApiIntegrationSearchResult {
	this := ApiIntegrationSearchResult{}
	this.TotalCount = totalCount
	this.Results = results
	return &this
}

// NewApiIntegrationSearchResultWithDefaults instantiates a new ApiIntegrationSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIntegrationSearchResultWithDefaults() *ApiIntegrationSearchResult {
	this := ApiIntegrationSearchResult{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *ApiIntegrationSearchResult) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationSearchResult) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ApiIntegrationSearchResult) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetResults returns the Results field value
func (o *ApiIntegrationSearchResult) GetResults() []ApiIntegrationResponse {
	if o == nil {
		var ret []ApiIntegrationResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationSearchResult) GetResultsOk() ([]ApiIntegrationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ApiIntegrationSearchResult) SetResults(v []ApiIntegrationResponse) {
	o.Results = v
}

func (o ApiIntegrationSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiIntegrationSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalCount"] = o.TotalCount
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableApiIntegrationSearchResult struct {
	value *ApiIntegrationSearchResult
	isSet bool
}

func (v NullableApiIntegrationSearchResult) Get() *ApiIntegrationSearchResult {
	return v.value
}

func (v *NullableApiIntegrationSearchResult) Set(val *ApiIntegrationSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIntegrationSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIntegrationSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIntegrationSearchResult(val *ApiIntegrationSearchResult) *NullableApiIntegrationSearchResult {
	return &NullableApiIntegrationSearchResult{value: val, isSet: true}
}

func (v NullableApiIntegrationSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIntegrationSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

