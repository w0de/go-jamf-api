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

// PrestagePurchasingInformation struct for PrestagePurchasingInformation
type PrestagePurchasingInformation struct {
	Id int32 `json:"id"`
	IsLeased bool `json:"isLeased"`
	IsPurchased bool `json:"isPurchased"`
	AppleCareID string `json:"appleCareID"`
	PoNumber string `json:"poNumber"`
	Vendor string `json:"vendor"`
	PurchasePrice string `json:"purchasePrice"`
	LifeExpectancy int32 `json:"lifeExpectancy"`
	PurchasingAccount string `json:"purchasingAccount"`
	PurchasingContact string `json:"purchasingContact"`
	LeaseDate string `json:"leaseDate"`
	PoDate string `json:"poDate"`
	WarrantyDate string `json:"warrantyDate"`
	VersionLock int32 `json:"versionLock"`
}

// NewPrestagePurchasingInformation instantiates a new PrestagePurchasingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestagePurchasingInformation(id int32, isLeased bool, isPurchased bool, appleCareID string, poNumber string, vendor string, purchasePrice string, lifeExpectancy int32, purchasingAccount string, purchasingContact string, leaseDate string, poDate string, warrantyDate string, versionLock int32) *PrestagePurchasingInformation {
	this := PrestagePurchasingInformation{}
	this.Id = id
	this.IsLeased = isLeased
	this.IsPurchased = isPurchased
	this.AppleCareID = appleCareID
	this.PoNumber = poNumber
	this.Vendor = vendor
	this.PurchasePrice = purchasePrice
	this.LifeExpectancy = lifeExpectancy
	this.PurchasingAccount = purchasingAccount
	this.PurchasingContact = purchasingContact
	this.LeaseDate = leaseDate
	this.PoDate = poDate
	this.WarrantyDate = warrantyDate
	this.VersionLock = versionLock
	return &this
}

// NewPrestagePurchasingInformationWithDefaults instantiates a new PrestagePurchasingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestagePurchasingInformationWithDefaults() *PrestagePurchasingInformation {
	this := PrestagePurchasingInformation{}
	return &this
}

// GetId returns the Id field value
func (o *PrestagePurchasingInformation) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrestagePurchasingInformation) SetId(v int32) {
	o.Id = v
}

// GetIsLeased returns the IsLeased field value
func (o *PrestagePurchasingInformation) GetIsLeased() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLeased
}

// GetIsLeasedOk returns a tuple with the IsLeased field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetIsLeasedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLeased, true
}

// SetIsLeased sets field value
func (o *PrestagePurchasingInformation) SetIsLeased(v bool) {
	o.IsLeased = v
}

// GetIsPurchased returns the IsPurchased field value
func (o *PrestagePurchasingInformation) GetIsPurchased() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPurchased
}

// GetIsPurchasedOk returns a tuple with the IsPurchased field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetIsPurchasedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPurchased, true
}

// SetIsPurchased sets field value
func (o *PrestagePurchasingInformation) SetIsPurchased(v bool) {
	o.IsPurchased = v
}

// GetAppleCareID returns the AppleCareID field value
func (o *PrestagePurchasingInformation) GetAppleCareID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppleCareID
}

// GetAppleCareIDOk returns a tuple with the AppleCareID field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetAppleCareIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppleCareID, true
}

// SetAppleCareID sets field value
func (o *PrestagePurchasingInformation) SetAppleCareID(v string) {
	o.AppleCareID = v
}

// GetPoNumber returns the PoNumber field value
func (o *PrestagePurchasingInformation) GetPoNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetPoNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoNumber, true
}

// SetPoNumber sets field value
func (o *PrestagePurchasingInformation) SetPoNumber(v string) {
	o.PoNumber = v
}

// GetVendor returns the Vendor field value
func (o *PrestagePurchasingInformation) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *PrestagePurchasingInformation) SetVendor(v string) {
	o.Vendor = v
}

// GetPurchasePrice returns the PurchasePrice field value
func (o *PrestagePurchasingInformation) GetPurchasePrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchasePrice
}

// GetPurchasePriceOk returns a tuple with the PurchasePrice field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetPurchasePriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasePrice, true
}

// SetPurchasePrice sets field value
func (o *PrestagePurchasingInformation) SetPurchasePrice(v string) {
	o.PurchasePrice = v
}

// GetLifeExpectancy returns the LifeExpectancy field value
func (o *PrestagePurchasingInformation) GetLifeExpectancy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LifeExpectancy
}

// GetLifeExpectancyOk returns a tuple with the LifeExpectancy field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetLifeExpectancyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifeExpectancy, true
}

// SetLifeExpectancy sets field value
func (o *PrestagePurchasingInformation) SetLifeExpectancy(v int32) {
	o.LifeExpectancy = v
}

// GetPurchasingAccount returns the PurchasingAccount field value
func (o *PrestagePurchasingInformation) GetPurchasingAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchasingAccount
}

// GetPurchasingAccountOk returns a tuple with the PurchasingAccount field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetPurchasingAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasingAccount, true
}

// SetPurchasingAccount sets field value
func (o *PrestagePurchasingInformation) SetPurchasingAccount(v string) {
	o.PurchasingAccount = v
}

// GetPurchasingContact returns the PurchasingContact field value
func (o *PrestagePurchasingInformation) GetPurchasingContact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchasingContact
}

// GetPurchasingContactOk returns a tuple with the PurchasingContact field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetPurchasingContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasingContact, true
}

// SetPurchasingContact sets field value
func (o *PrestagePurchasingInformation) SetPurchasingContact(v string) {
	o.PurchasingContact = v
}

// GetLeaseDate returns the LeaseDate field value
func (o *PrestagePurchasingInformation) GetLeaseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeaseDate
}

// GetLeaseDateOk returns a tuple with the LeaseDate field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetLeaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeaseDate, true
}

// SetLeaseDate sets field value
func (o *PrestagePurchasingInformation) SetLeaseDate(v string) {
	o.LeaseDate = v
}

// GetPoDate returns the PoDate field value
func (o *PrestagePurchasingInformation) GetPoDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoDate
}

// GetPoDateOk returns a tuple with the PoDate field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetPoDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoDate, true
}

// SetPoDate sets field value
func (o *PrestagePurchasingInformation) SetPoDate(v string) {
	o.PoDate = v
}

// GetWarrantyDate returns the WarrantyDate field value
func (o *PrestagePurchasingInformation) GetWarrantyDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarrantyDate
}

// GetWarrantyDateOk returns a tuple with the WarrantyDate field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetWarrantyDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarrantyDate, true
}

// SetWarrantyDate sets field value
func (o *PrestagePurchasingInformation) SetWarrantyDate(v string) {
	o.WarrantyDate = v
}

// GetVersionLock returns the VersionLock field value
func (o *PrestagePurchasingInformation) GetVersionLock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value
// and a boolean to check if the value has been set.
func (o *PrestagePurchasingInformation) GetVersionLockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionLock, true
}

// SetVersionLock sets field value
func (o *PrestagePurchasingInformation) SetVersionLock(v int32) {
	o.VersionLock = v
}

func (o PrestagePurchasingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["isLeased"] = o.IsLeased
	}
	if true {
		toSerialize["isPurchased"] = o.IsPurchased
	}
	if true {
		toSerialize["appleCareID"] = o.AppleCareID
	}
	if true {
		toSerialize["poNumber"] = o.PoNumber
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	if true {
		toSerialize["purchasePrice"] = o.PurchasePrice
	}
	if true {
		toSerialize["lifeExpectancy"] = o.LifeExpectancy
	}
	if true {
		toSerialize["purchasingAccount"] = o.PurchasingAccount
	}
	if true {
		toSerialize["purchasingContact"] = o.PurchasingContact
	}
	if true {
		toSerialize["leaseDate"] = o.LeaseDate
	}
	if true {
		toSerialize["poDate"] = o.PoDate
	}
	if true {
		toSerialize["warrantyDate"] = o.WarrantyDate
	}
	if true {
		toSerialize["versionLock"] = o.VersionLock
	}
	return json.Marshal(toSerialize)
}

type NullablePrestagePurchasingInformation struct {
	value *PrestagePurchasingInformation
	isSet bool
}

func (v NullablePrestagePurchasingInformation) Get() *PrestagePurchasingInformation {
	return v.value
}

func (v *NullablePrestagePurchasingInformation) Set(val *PrestagePurchasingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestagePurchasingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestagePurchasingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestagePurchasingInformation(val *PrestagePurchasingInformation) *NullablePrestagePurchasingInformation {
	return &NullablePrestagePurchasingInformation{value: val, isSet: true}
}

func (v NullablePrestagePurchasingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestagePurchasingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

