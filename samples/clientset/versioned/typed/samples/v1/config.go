// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	samplesv1 "github.com/openshift/api/samples/v1"
	applyconfigurationssamplesv1 "github.com/openshift/client-go/samples/applyconfigurations/samples/v1"
	scheme "github.com/openshift/client-go/samples/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ConfigsGetter has a method to return a ConfigInterface.
// A group's client should implement this interface.
type ConfigsGetter interface {
	Configs() ConfigInterface
}

// ConfigInterface has methods to work with Config resources.
type ConfigInterface interface {
	Create(ctx context.Context, config *samplesv1.Config, opts metav1.CreateOptions) (*samplesv1.Config, error)
	Update(ctx context.Context, config *samplesv1.Config, opts metav1.UpdateOptions) (*samplesv1.Config, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, config *samplesv1.Config, opts metav1.UpdateOptions) (*samplesv1.Config, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*samplesv1.Config, error)
	List(ctx context.Context, opts metav1.ListOptions) (*samplesv1.ConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *samplesv1.Config, err error)
	Apply(ctx context.Context, config *applyconfigurationssamplesv1.ConfigApplyConfiguration, opts metav1.ApplyOptions) (result *samplesv1.Config, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, config *applyconfigurationssamplesv1.ConfigApplyConfiguration, opts metav1.ApplyOptions) (result *samplesv1.Config, err error)
	ConfigExpansion
}

// configs implements ConfigInterface
type configs struct {
	*gentype.ClientWithListAndApply[*samplesv1.Config, *samplesv1.ConfigList, *applyconfigurationssamplesv1.ConfigApplyConfiguration]
}

// newConfigs returns a Configs
func newConfigs(c *SamplesV1Client) *configs {
	return &configs{
		gentype.NewClientWithListAndApply[*samplesv1.Config, *samplesv1.ConfigList, *applyconfigurationssamplesv1.ConfigApplyConfiguration](
			"configs",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *samplesv1.Config { return &samplesv1.Config{} },
			func() *samplesv1.ConfigList { return &samplesv1.ConfigList{} },
		),
	}
}
