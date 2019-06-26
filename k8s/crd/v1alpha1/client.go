package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

var SchemeGroupVersion = schema.GroupVersion{Group: CRDGroup, Version: CRDVersion}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&SslConfig{},
		&SslConfigList{},
	)
	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func NewClient(cfg *rest.Config, ns string) (SslConfigInterface, error) {
	scheme := runtime.NewScheme()
	builder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := builder.AddToScheme(scheme); err != nil {
		return nil, err
	}

	config := *cfg
	config.GroupVersion = &SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &sslConfigClient{client: client, ns: ns}, nil
}

type SslConfigInterface interface {
	Create(obj *SslConfig) (*SslConfig, error)
	Update(obj *SslConfig) (*SslConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string) (*SslConfig, error)
}

// +k8s:deepcopy-gen=false
type sslConfigClient struct {
	client rest.Interface
	ns     string
}

func (c *sslConfigClient) Create(obj *SslConfig) (*SslConfig, error) {
	result := &SslConfig{}
	err := c.client.Post().Namespace(c.ns).Resource("sslconfigs").Body(obj).Do().Into(result)
	return result, err
}

func (c *sslConfigClient) Update(obj *SslConfig) (*SslConfig, error) {
	result := &SslConfig{}
	err := c.client.Put().Namespace(c.ns).Resource("sslconfigs").Body(obj).Do().Into(result)
	return result, err
}

func (c *sslConfigClient) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().Namespace(c.ns).Resource("sslconfigs").Name(name).Body(options).Do().Error()
}

func (c *sslConfigClient) Get(name string) (*SslConfig, error) {
	result := &SslConfig{}
	err := c.client.Get().Namespace(c.ns).Resource("sslconfigs").Name(name).Do().Into(result)
	return result, err
}
