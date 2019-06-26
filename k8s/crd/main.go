package main

import (
	"crd/v1alpha1"
	"flag"
	"github.com/golang/glog"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func main() {
	proxyURL := flag.String("proxy", "",
		`If specified, it is assumed that a kubctl proxy server is running on the
		given url and creates a proxy client. In case it is not given InCluster
		kubernetes setup will be used`)

	flag.Parse()

	var config *rest.Config
	var err error
	if *proxyURL != "" {
		config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{},
			&clientcmd.ConfigOverrides{
				ClusterInfo: api.Cluster{Server: *proxyURL},
			}).ClientConfig()
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		glog.Fatalf("error creating client configuration: %v", err)
	}

	glog.Infof("try create clientset")
	kclient, err := clientset.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Failed to create client: %v", err)
	}

	glog.Infof("try create crd")
	err = v1alpha1.CreateCRD(kclient)
	if err != nil {
		glog.Fatal(err)
	}

	// Wait for the CRD to be created before we use it.
	//time.Sleep(5 * time.Second)

	glog.Infof("try create v1alph1.Client")
	crdclient, err := v1alpha1.NewClient(config, "default")
	if err != nil {
		glog.Fatal(err)
	}

	// create a new SslConfig Object
	sslcfg := &v1alpha1.SslConfig{
		ObjectMeta: v1.ObjectMeta{
			Name:   "sslconfigobj",
			Labels: map[string]string{"mylabel": "crd"},
		},
		Spec: v1alpha1.SslConfigSpec{
			Cert:   "my-cert",
			Key:    "my-key",
			Domain: "*.velotio.com",
		},
		Status: v1alpha1.SslConfigStatus{
			State:   "created",
			Message: "Created, not processed yet",
		},
	}

	// Create
	rsp, err := crdclient.Create(sslcfg)
	if err != nil {
		glog.Errorf("error while creating object: %v\n", err)
	} else {
		glog.Infof("object created: %v\n", rsp)
	}

	// Get
	obj, err := crdclient.Get(sslcfg.ObjectMeta.Name)
	if err != nil {
		glog.Infof("error while getting the object %v\n", err)
	} else {
		glog.Infof("SslConfig Objects Found: \n%+v\n", obj)
	}

	select {}
}
