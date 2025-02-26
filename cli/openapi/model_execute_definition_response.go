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

// checks if the ExecuteDefinitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecuteDefinitionResponse{}

// ExecuteDefinitionResponse struct for ExecuteDefinitionResponse
type ExecuteDefinitionResponse struct {
	// resource ID
	Id *string `json:"id,omitempty"`
	// run ID
	RunId *string `json:"runId,omitempty"`
	// resource type
	Type *string `json:"type,omitempty"`
	// resource web UI url
	Url *string `json:"url,omitempty"`
}

// NewExecuteDefinitionResponse instantiates a new ExecuteDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteDefinitionResponse() *ExecuteDefinitionResponse {
	this := ExecuteDefinitionResponse{}
	return &this
}

// NewExecuteDefinitionResponseWithDefaults instantiates a new ExecuteDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteDefinitionResponseWithDefaults() *ExecuteDefinitionResponse {
	this := ExecuteDefinitionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExecuteDefinitionResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteDefinitionResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExecuteDefinitionResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExecuteDefinitionResponse) SetId(v string) {
	o.Id = &v
}

// GetRunId returns the RunId field value if set, zero value otherwise.
func (o *ExecuteDefinitionResponse) GetRunId() string {
	if o == nil || isNil(o.RunId) {
		var ret string
		return ret
	}
	return *o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteDefinitionResponse) GetRunIdOk() (*string, bool) {
	if o == nil || isNil(o.RunId) {
		return nil, false
	}
	return o.RunId, true
}

// HasRunId returns a boolean if a field has been set.
func (o *ExecuteDefinitionResponse) HasRunId() bool {
	if o != nil && !isNil(o.RunId) {
		return true
	}

	return false
}

// SetRunId gets a reference to the given string and assigns it to the RunId field.
func (o *ExecuteDefinitionResponse) SetRunId(v string) {
	o.RunId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExecuteDefinitionResponse) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteDefinitionResponse) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExecuteDefinitionResponse) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExecuteDefinitionResponse) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ExecuteDefinitionResponse) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteDefinitionResponse) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ExecuteDefinitionResponse) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ExecuteDefinitionResponse) SetUrl(v string) {
	o.Url = &v
}

func (o ExecuteDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecuteDefinitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.RunId) {
		toSerialize["runId"] = o.RunId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableExecuteDefinitionResponse struct {
	value *ExecuteDefinitionResponse
	isSet bool
}

func (v NullableExecuteDefinitionResponse) Get() *ExecuteDefinitionResponse {
	return v.value
}

func (v *NullableExecuteDefinitionResponse) Set(val *ExecuteDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteDefinitionResponse(val *ExecuteDefinitionResponse) *NullableExecuteDefinitionResponse {
	return &NullableExecuteDefinitionResponse{value: val, isSet: true}
}

func (v NullableExecuteDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
