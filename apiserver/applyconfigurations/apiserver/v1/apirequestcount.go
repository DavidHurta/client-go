// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apiserverv1 "github.com/openshift/api/apiserver/v1"
	internal "github.com/openshift/client-go/apiserver/applyconfigurations/internal"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// APIRequestCountApplyConfiguration represents a declarative configuration of the APIRequestCount type for use
// with apply.
type APIRequestCountApplyConfiguration struct {
	metav1.TypeMetaApplyConfiguration    `json:",inline"`
	*metav1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                                 *APIRequestCountSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                               *APIRequestCountStatusApplyConfiguration `json:"status,omitempty"`
}

// APIRequestCount constructs a declarative configuration of the APIRequestCount type for use with
// apply.
func APIRequestCount(name string) *APIRequestCountApplyConfiguration {
	b := &APIRequestCountApplyConfiguration{}
	b.WithName(name)
	b.WithKind("APIRequestCount")
	b.WithAPIVersion("apiserver.openshift.io/v1")
	return b
}

// ExtractAPIRequestCount extracts the applied configuration owned by fieldManager from
// aPIRequestCount. If no managedFields are found in aPIRequestCount for fieldManager, a
// APIRequestCountApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// aPIRequestCount must be a unmodified APIRequestCount API object that was retrieved from the Kubernetes API.
// ExtractAPIRequestCount provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractAPIRequestCount(aPIRequestCount *apiserverv1.APIRequestCount, fieldManager string) (*APIRequestCountApplyConfiguration, error) {
	return extractAPIRequestCount(aPIRequestCount, fieldManager, "")
}

// ExtractAPIRequestCountStatus is the same as ExtractAPIRequestCount except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractAPIRequestCountStatus(aPIRequestCount *apiserverv1.APIRequestCount, fieldManager string) (*APIRequestCountApplyConfiguration, error) {
	return extractAPIRequestCount(aPIRequestCount, fieldManager, "status")
}

func extractAPIRequestCount(aPIRequestCount *apiserverv1.APIRequestCount, fieldManager string, subresource string) (*APIRequestCountApplyConfiguration, error) {
	b := &APIRequestCountApplyConfiguration{}
	err := managedfields.ExtractInto(aPIRequestCount, internal.Parser().Type("com.github.openshift.api.apiserver.v1.APIRequestCount"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(aPIRequestCount.Name)

	b.WithKind("APIRequestCount")
	b.WithAPIVersion("apiserver.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithKind(value string) *APIRequestCountApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithAPIVersion(value string) *APIRequestCountApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithName(value string) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithGenerateName(value string) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithNamespace(value string) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithUID(value types.UID) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithResourceVersion(value string) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithGeneration(value int64) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithCreationTimestamp(value apismetav1.Time) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithDeletionTimestamp(value apismetav1.Time) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *APIRequestCountApplyConfiguration) WithLabels(entries map[string]string) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *APIRequestCountApplyConfiguration) WithAnnotations(entries map[string]string) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *APIRequestCountApplyConfiguration) WithOwnerReferences(values ...*metav1.OwnerReferenceApplyConfiguration) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *APIRequestCountApplyConfiguration) WithFinalizers(values ...string) *APIRequestCountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *APIRequestCountApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &metav1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithSpec(value *APIRequestCountSpecApplyConfiguration) *APIRequestCountApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *APIRequestCountApplyConfiguration) WithStatus(value *APIRequestCountStatusApplyConfiguration) *APIRequestCountApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *APIRequestCountApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}
