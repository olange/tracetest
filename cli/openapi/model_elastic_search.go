/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ElasticSearch struct for ElasticSearch
type ElasticSearch struct {
	Addresses   []string `json:"addresses,omitempty"`
	Username    *string  `json:"username,omitempty"`
	Password    *string  `json:"password,omitempty"`
	Index       *string  `json:"index,omitempty"`
	Certificate *string  `json:"certificate,omitempty"`
}

// NewElasticSearch instantiates a new ElasticSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElasticSearch() *ElasticSearch {
	this := ElasticSearch{}
	return &this
}

// NewElasticSearchWithDefaults instantiates a new ElasticSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElasticSearchWithDefaults() *ElasticSearch {
	this := ElasticSearch{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ElasticSearch) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticSearch) GetAddressesOk() ([]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ElasticSearch) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *ElasticSearch) SetAddresses(v []string) {
	o.Addresses = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ElasticSearch) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticSearch) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ElasticSearch) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ElasticSearch) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ElasticSearch) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticSearch) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ElasticSearch) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ElasticSearch) SetPassword(v string) {
	o.Password = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ElasticSearch) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticSearch) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ElasticSearch) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *ElasticSearch) SetIndex(v string) {
	o.Index = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *ElasticSearch) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticSearch) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *ElasticSearch) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *ElasticSearch) SetCertificate(v string) {
	o.Certificate = &v
}

func (o ElasticSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableElasticSearch struct {
	value *ElasticSearch
	isSet bool
}

func (v NullableElasticSearch) Get() *ElasticSearch {
	return v.value
}

func (v *NullableElasticSearch) Set(val *ElasticSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableElasticSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableElasticSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElasticSearch(val *ElasticSearch) *NullableElasticSearch {
	return &NullableElasticSearch{value: val, isSet: true}
}

func (v NullableElasticSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElasticSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}