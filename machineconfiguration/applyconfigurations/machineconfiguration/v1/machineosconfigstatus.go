// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	machineconfigurationv1 "github.com/openshift/api/machineconfiguration/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// MachineOSConfigStatusApplyConfiguration represents a declarative configuration of the MachineOSConfigStatus type for use
// with apply.
type MachineOSConfigStatusApplyConfiguration struct {
	Conditions           []v1.ConditionApplyConfiguration          `json:"conditions,omitempty"`
	ObservedGeneration   *int64                                    `json:"observedGeneration,omitempty"`
	CurrentImagePullSpec *machineconfigurationv1.ImageDigestFormat `json:"currentImagePullSpec,omitempty"`
	MachineOSBuild       *ObjectReferenceApplyConfiguration        `json:"machineOSBuild,omitempty"`
}

// MachineOSConfigStatusApplyConfiguration constructs a declarative configuration of the MachineOSConfigStatus type for use with
// apply.
func MachineOSConfigStatus() *MachineOSConfigStatusApplyConfiguration {
	return &MachineOSConfigStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *MachineOSConfigStatusApplyConfiguration) WithConditions(values ...*v1.ConditionApplyConfiguration) *MachineOSConfigStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *MachineOSConfigStatusApplyConfiguration) WithObservedGeneration(value int64) *MachineOSConfigStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithCurrentImagePullSpec sets the CurrentImagePullSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentImagePullSpec field is set to the value of the last call.
func (b *MachineOSConfigStatusApplyConfiguration) WithCurrentImagePullSpec(value machineconfigurationv1.ImageDigestFormat) *MachineOSConfigStatusApplyConfiguration {
	b.CurrentImagePullSpec = &value
	return b
}

// WithMachineOSBuild sets the MachineOSBuild field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MachineOSBuild field is set to the value of the last call.
func (b *MachineOSConfigStatusApplyConfiguration) WithMachineOSBuild(value *ObjectReferenceApplyConfiguration) *MachineOSConfigStatusApplyConfiguration {
	b.MachineOSBuild = value
	return b
}
