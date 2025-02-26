/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the TransactionRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRun{}

// TransactionRun struct for TransactionRun
type TransactionRun struct {
	Id          *string            `json:"id,omitempty"`
	Version     *int32             `json:"version,omitempty"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty"`
	CompletedAt *time.Time         `json:"completedAt,omitempty"`
	State       *string            `json:"state,omitempty"`
	Steps       []TestRun          `json:"steps,omitempty"`
	Environment *Environment       `json:"environment,omitempty"`
	Metadata    *map[string]string `json:"metadata,omitempty"`
	Pass        *int32             `json:"pass,omitempty"`
	Fail        *int32             `json:"fail,omitempty"`
}

// NewTransactionRun instantiates a new TransactionRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRun() *TransactionRun {
	this := TransactionRun{}
	return &this
}

// NewTransactionRunWithDefaults instantiates a new TransactionRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRunWithDefaults() *TransactionRun {
	this := TransactionRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionRun) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionRun) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionRun) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TransactionRun) GetVersion() int32 {
	if o == nil || isNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetVersionOk() (*int32, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TransactionRun) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *TransactionRun) SetVersion(v int32) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransactionRun) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransactionRun) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TransactionRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *TransactionRun) GetCompletedAt() time.Time {
	if o == nil || isNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *TransactionRun) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *TransactionRun) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TransactionRun) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TransactionRun) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TransactionRun) SetState(v string) {
	o.State = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *TransactionRun) GetSteps() []TestRun {
	if o == nil || isNil(o.Steps) {
		var ret []TestRun
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetStepsOk() ([]TestRun, bool) {
	if o == nil || isNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *TransactionRun) HasSteps() bool {
	if o != nil && !isNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []TestRun and assigns it to the Steps field.
func (o *TransactionRun) SetSteps(v []TestRun) {
	o.Steps = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *TransactionRun) GetEnvironment() Environment {
	if o == nil || isNil(o.Environment) {
		var ret Environment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetEnvironmentOk() (*Environment, bool) {
	if o == nil || isNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *TransactionRun) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given Environment and assigns it to the Environment field.
func (o *TransactionRun) SetEnvironment(v Environment) {
	o.Environment = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TransactionRun) GetMetadata() map[string]string {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransactionRun) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransactionRun) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *TransactionRun) GetPass() int32 {
	if o == nil || isNil(o.Pass) {
		var ret int32
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetPassOk() (*int32, bool) {
	if o == nil || isNil(o.Pass) {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *TransactionRun) HasPass() bool {
	if o != nil && !isNil(o.Pass) {
		return true
	}

	return false
}

// SetPass gets a reference to the given int32 and assigns it to the Pass field.
func (o *TransactionRun) SetPass(v int32) {
	o.Pass = &v
}

// GetFail returns the Fail field value if set, zero value otherwise.
func (o *TransactionRun) GetFail() int32 {
	if o == nil || isNil(o.Fail) {
		var ret int32
		return ret
	}
	return *o.Fail
}

// GetFailOk returns a tuple with the Fail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRun) GetFailOk() (*int32, bool) {
	if o == nil || isNil(o.Fail) {
		return nil, false
	}
	return o.Fail, true
}

// HasFail returns a boolean if a field has been set.
func (o *TransactionRun) HasFail() bool {
	if o != nil && !isNil(o.Fail) {
		return true
	}

	return false
}

// SetFail gets a reference to the given int32 and assigns it to the Fail field.
func (o *TransactionRun) SetFail(v int32) {
	o.Fail = &v
}

func (o TransactionRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: version is readOnly
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Pass) {
		toSerialize["pass"] = o.Pass
	}
	if !isNil(o.Fail) {
		toSerialize["fail"] = o.Fail
	}
	return toSerialize, nil
}

type NullableTransactionRun struct {
	value *TransactionRun
	isSet bool
}

func (v NullableTransactionRun) Get() *TransactionRun {
	return v.value
}

func (v *NullableTransactionRun) Set(val *TransactionRun) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRun) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRun(val *TransactionRun) *NullableTransactionRun {
	return &NullableTransactionRun{value: val, isSet: true}
}

func (v NullableTransactionRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
