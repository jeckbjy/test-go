package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

const (
	CRDPlural   = "sslconfigs"
	CRDGroup    = "blog.velotio.com"
	CRDVersion  = "v1alpha1"
	CRDFullName = CRDPlural + "." + CRDGroup
)

func CreateCRD(sc clientset.Interface) error {
	crd := &v1beta1.CustomResourceDefinition{
		ObjectMeta: v1.ObjectMeta{Name: CRDFullName},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   CRDGroup,
			Version: CRDVersion,
			Scope:   v1beta1.NamespaceScoped,
			Names: v1beta1.CustomResourceDefinitionNames{
				Plural: CRDPlural,
				Kind:   reflect.TypeOf(SslConfig{}).Name(),
			},
		},
	}

	_, err := sc.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if err != nil && errors.IsAlreadyExists(err) {
		return nil
	}

	return err
}
