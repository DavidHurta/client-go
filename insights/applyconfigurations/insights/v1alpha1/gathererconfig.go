// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	insightsv1alpha1 "github.com/openshift/api/insights/v1alpha1"
)

// GathererConfigApplyConfiguration represents a declarative configuration of the GathererConfig type for use
// with apply.
type GathererConfigApplyConfiguration struct {
	Name  *string                         `json:"name,omitempty"`
	State *insightsv1alpha1.GathererState `json:"state,omitempty"`
}

// GathererConfigApplyConfiguration constructs a declarative configuration of the GathererConfig type for use with
// apply.
func GathererConfig() *GathererConfigApplyConfiguration {
	return &GathererConfigApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GathererConfigApplyConfiguration) WithName(value string) *GathererConfigApplyConfiguration {
	b.Name = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *GathererConfigApplyConfiguration) WithState(value insightsv1alpha1.GathererState) *GathererConfigApplyConfiguration {
	b.State = &value
	return b
}
