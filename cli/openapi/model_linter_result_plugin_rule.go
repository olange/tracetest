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

// checks if the LinterResultPluginRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinterResultPluginRule{}

// LinterResultPluginRule struct for LinterResultPluginRule
type LinterResultPluginRule struct {
	Name             *string                        `json:"name,omitempty"`
	Description      *string                        `json:"description,omitempty"`
	ErrorDescription *string                        `json:"errorDescription,omitempty"`
	Passed           *bool                          `json:"passed,omitempty"`
	Weight           *int32                         `json:"weight,omitempty"`
	Tips             []string                       `json:"tips,omitempty"`
	Results          []LinterResultPluginRuleResult `json:"results,omitempty"`
}

// NewLinterResultPluginRule instantiates a new LinterResultPluginRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinterResultPluginRule() *LinterResultPluginRule {
	this := LinterResultPluginRule{}
	return &this
}

// NewLinterResultPluginRuleWithDefaults instantiates a new LinterResultPluginRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinterResultPluginRuleWithDefaults() *LinterResultPluginRule {
	this := LinterResultPluginRule{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinterResultPluginRule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinterResultPluginRule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinterResultPluginRule) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LinterResultPluginRule) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRule) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LinterResultPluginRule) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LinterResultPluginRule) SetDescription(v string) {
	o.Description = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *LinterResultPluginRule) GetErrorDescription() string {
	if o == nil || isNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRule) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *LinterResultPluginRule) HasErrorDescription() bool {
	if o != nil && !isNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *LinterResultPluginRule) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetPassed returns the Passed field value if set, zero value otherwise.
func (o *LinterResultPluginRule) GetPassed() bool {
	if o == nil || isNil(o.Passed) {
		var ret bool
		return ret
	}
	return *o.Passed
}

// GetPassedOk returns a tuple with the Passed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRule) GetPassedOk() (*bool, bool) {
	if o == nil || isNil(o.Passed) {
		return nil, false
	}
	return o.Passed, true
}

// HasPassed returns a boolean if a field has been set.
func (o *LinterResultPluginRule) HasPassed() bool {
	if o != nil && !isNil(o.Passed) {
		return true
	}

	return false
}

// SetPassed gets a reference to the given bool and assigns it to the Passed field.
func (o *LinterResultPluginRule) SetPassed(v bool) {
	o.Passed = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *LinterResultPluginRule) GetWeight() int32 {
	if o == nil || isNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRule) GetWeightOk() (*int32, bool) {
	if o == nil || isNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *LinterResultPluginRule) HasWeight() bool {
	if o != nil && !isNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *LinterResultPluginRule) SetWeight(v int32) {
	o.Weight = &v
}

// GetTips returns the Tips field value if set, zero value otherwise.
func (o *LinterResultPluginRule) GetTips() []string {
	if o == nil || isNil(o.Tips) {
		var ret []string
		return ret
	}
	return o.Tips
}

// GetTipsOk returns a tuple with the Tips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRule) GetTipsOk() ([]string, bool) {
	if o == nil || isNil(o.Tips) {
		return nil, false
	}
	return o.Tips, true
}

// HasTips returns a boolean if a field has been set.
func (o *LinterResultPluginRule) HasTips() bool {
	if o != nil && !isNil(o.Tips) {
		return true
	}

	return false
}

// SetTips gets a reference to the given []string and assigns it to the Tips field.
func (o *LinterResultPluginRule) SetTips(v []string) {
	o.Tips = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LinterResultPluginRule) GetResults() []LinterResultPluginRuleResult {
	if o == nil || isNil(o.Results) {
		var ret []LinterResultPluginRuleResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRule) GetResultsOk() ([]LinterResultPluginRuleResult, bool) {
	if o == nil || isNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LinterResultPluginRule) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LinterResultPluginRuleResult and assigns it to the Results field.
func (o *LinterResultPluginRule) SetResults(v []LinterResultPluginRuleResult) {
	o.Results = v
}

func (o LinterResultPluginRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinterResultPluginRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ErrorDescription) {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if !isNil(o.Passed) {
		toSerialize["passed"] = o.Passed
	}
	if !isNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !isNil(o.Tips) {
		toSerialize["tips"] = o.Tips
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableLinterResultPluginRule struct {
	value *LinterResultPluginRule
	isSet bool
}

func (v NullableLinterResultPluginRule) Get() *LinterResultPluginRule {
	return v.value
}

func (v *NullableLinterResultPluginRule) Set(val *LinterResultPluginRule) {
	v.value = val
	v.isSet = true
}

func (v NullableLinterResultPluginRule) IsSet() bool {
	return v.isSet
}

func (v *NullableLinterResultPluginRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinterResultPluginRule(val *LinterResultPluginRule) *NullableLinterResultPluginRule {
	return &NullableLinterResultPluginRule{value: val, isSet: true}
}

func (v NullableLinterResultPluginRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinterResultPluginRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
