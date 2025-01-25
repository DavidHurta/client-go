// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	insightsv1alpha1 "github.com/openshift/api/insights/v1alpha1"
)

// HealthCheckApplyConfiguration represents a declarative configuration of the HealthCheck type for use
// with apply.
type HealthCheckApplyConfiguration struct {
	Description *string                            `json:"description,omitempty"`
	TotalRisk   *int32                             `json:"totalRisk,omitempty"`
	AdvisorURI  *string                            `json:"advisorURI,omitempty"`
	State       *insightsv1alpha1.HealthCheckState `json:"state,omitempty"`
}

// HealthCheckApplyConfiguration constructs a declarative configuration of the HealthCheck type for use with
// apply.
func HealthCheck() *HealthCheckApplyConfiguration {
	return &HealthCheckApplyConfiguration{}
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *HealthCheckApplyConfiguration) WithDescription(value string) *HealthCheckApplyConfiguration {
	b.Description = &value
	return b
}

// WithTotalRisk sets the TotalRisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TotalRisk field is set to the value of the last call.
func (b *HealthCheckApplyConfiguration) WithTotalRisk(value int32) *HealthCheckApplyConfiguration {
	b.TotalRisk = &value
	return b
}

// WithAdvisorURI sets the AdvisorURI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdvisorURI field is set to the value of the last call.
func (b *HealthCheckApplyConfiguration) WithAdvisorURI(value string) *HealthCheckApplyConfiguration {
	b.AdvisorURI = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *HealthCheckApplyConfiguration) WithState(value insightsv1alpha1.HealthCheckState) *HealthCheckApplyConfiguration {
	b.State = &value
	return b
}
