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

// SsoKeystoreResponse struct for SsoKeystoreResponse
type SsoKeystoreResponse struct {
	Key *string `json:"key,omitempty"`
	Keys []CertificateKey `json:"keys,omitempty"`
	Type *string `json:"type,omitempty"`
	KeystoreSetupType *string `json:"keystoreSetupType,omitempty"`
	KeystoreFileName *string `json:"keystoreFileName,omitempty"`
}

// NewSsoKeystoreResponse instantiates a new SsoKeystoreResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoKeystoreResponse() *SsoKeystoreResponse {
	this := SsoKeystoreResponse{}
	var key string = " "
	this.Key = &key
	return &this
}

// NewSsoKeystoreResponseWithDefaults instantiates a new SsoKeystoreResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoKeystoreResponseWithDefaults() *SsoKeystoreResponse {
	this := SsoKeystoreResponse{}
	var key string = " "
	this.Key = &key
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SsoKeystoreResponse) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreResponse) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SsoKeystoreResponse) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SsoKeystoreResponse) SetKey(v string) {
	o.Key = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *SsoKeystoreResponse) GetKeys() []CertificateKey {
	if o == nil || o.Keys == nil {
		var ret []CertificateKey
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreResponse) GetKeysOk() ([]CertificateKey, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SsoKeystoreResponse) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []CertificateKey and assigns it to the Keys field.
func (o *SsoKeystoreResponse) SetKeys(v []CertificateKey) {
	o.Keys = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SsoKeystoreResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SsoKeystoreResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SsoKeystoreResponse) SetType(v string) {
	o.Type = &v
}

// GetKeystoreSetupType returns the KeystoreSetupType field value if set, zero value otherwise.
func (o *SsoKeystoreResponse) GetKeystoreSetupType() string {
	if o == nil || o.KeystoreSetupType == nil {
		var ret string
		return ret
	}
	return *o.KeystoreSetupType
}

// GetKeystoreSetupTypeOk returns a tuple with the KeystoreSetupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreResponse) GetKeystoreSetupTypeOk() (*string, bool) {
	if o == nil || o.KeystoreSetupType == nil {
		return nil, false
	}
	return o.KeystoreSetupType, true
}

// HasKeystoreSetupType returns a boolean if a field has been set.
func (o *SsoKeystoreResponse) HasKeystoreSetupType() bool {
	if o != nil && o.KeystoreSetupType != nil {
		return true
	}

	return false
}

// SetKeystoreSetupType gets a reference to the given string and assigns it to the KeystoreSetupType field.
func (o *SsoKeystoreResponse) SetKeystoreSetupType(v string) {
	o.KeystoreSetupType = &v
}

// GetKeystoreFileName returns the KeystoreFileName field value if set, zero value otherwise.
func (o *SsoKeystoreResponse) GetKeystoreFileName() string {
	if o == nil || o.KeystoreFileName == nil {
		var ret string
		return ret
	}
	return *o.KeystoreFileName
}

// GetKeystoreFileNameOk returns a tuple with the KeystoreFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoKeystoreResponse) GetKeystoreFileNameOk() (*string, bool) {
	if o == nil || o.KeystoreFileName == nil {
		return nil, false
	}
	return o.KeystoreFileName, true
}

// HasKeystoreFileName returns a boolean if a field has been set.
func (o *SsoKeystoreResponse) HasKeystoreFileName() bool {
	if o != nil && o.KeystoreFileName != nil {
		return true
	}

	return false
}

// SetKeystoreFileName gets a reference to the given string and assigns it to the KeystoreFileName field.
func (o *SsoKeystoreResponse) SetKeystoreFileName(v string) {
	o.KeystoreFileName = &v
}

func (o SsoKeystoreResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.KeystoreSetupType != nil {
		toSerialize["keystoreSetupType"] = o.KeystoreSetupType
	}
	if o.KeystoreFileName != nil {
		toSerialize["keystoreFileName"] = o.KeystoreFileName
	}
	return json.Marshal(toSerialize)
}

type NullableSsoKeystoreResponse struct {
	value *SsoKeystoreResponse
	isSet bool
}

func (v NullableSsoKeystoreResponse) Get() *SsoKeystoreResponse {
	return v.value
}

func (v *NullableSsoKeystoreResponse) Set(val *SsoKeystoreResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoKeystoreResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoKeystoreResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoKeystoreResponse(val *SsoKeystoreResponse) *NullableSsoKeystoreResponse {
	return &NullableSsoKeystoreResponse{value: val, isSet: true}
}

func (v NullableSsoKeystoreResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoKeystoreResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

