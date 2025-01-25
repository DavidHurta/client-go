// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	consolev1 "github.com/openshift/api/console/v1"
	applyconfigurationsconsolev1 "github.com/openshift/client-go/console/applyconfigurations/console/v1"
	scheme "github.com/openshift/client-go/console/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ConsoleCLIDownloadsGetter has a method to return a ConsoleCLIDownloadInterface.
// A group's client should implement this interface.
type ConsoleCLIDownloadsGetter interface {
	ConsoleCLIDownloads() ConsoleCLIDownloadInterface
}

// ConsoleCLIDownloadInterface has methods to work with ConsoleCLIDownload resources.
type ConsoleCLIDownloadInterface interface {
	Create(ctx context.Context, consoleCLIDownload *consolev1.ConsoleCLIDownload, opts metav1.CreateOptions) (*consolev1.ConsoleCLIDownload, error)
	Update(ctx context.Context, consoleCLIDownload *consolev1.ConsoleCLIDownload, opts metav1.UpdateOptions) (*consolev1.ConsoleCLIDownload, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*consolev1.ConsoleCLIDownload, error)
	List(ctx context.Context, opts metav1.ListOptions) (*consolev1.ConsoleCLIDownloadList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *consolev1.ConsoleCLIDownload, err error)
	Apply(ctx context.Context, consoleCLIDownload *applyconfigurationsconsolev1.ConsoleCLIDownloadApplyConfiguration, opts metav1.ApplyOptions) (result *consolev1.ConsoleCLIDownload, err error)
	ConsoleCLIDownloadExpansion
}

// consoleCLIDownloads implements ConsoleCLIDownloadInterface
type consoleCLIDownloads struct {
	*gentype.ClientWithListAndApply[*consolev1.ConsoleCLIDownload, *consolev1.ConsoleCLIDownloadList, *applyconfigurationsconsolev1.ConsoleCLIDownloadApplyConfiguration]
}

// newConsoleCLIDownloads returns a ConsoleCLIDownloads
func newConsoleCLIDownloads(c *ConsoleV1Client) *consoleCLIDownloads {
	return &consoleCLIDownloads{
		gentype.NewClientWithListAndApply[*consolev1.ConsoleCLIDownload, *consolev1.ConsoleCLIDownloadList, *applyconfigurationsconsolev1.ConsoleCLIDownloadApplyConfiguration](
			"consoleclidownloads",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *consolev1.ConsoleCLIDownload { return &consolev1.ConsoleCLIDownload{} },
			func() *consolev1.ConsoleCLIDownloadList { return &consolev1.ConsoleCLIDownloadList{} },
		),
	}
}
