package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SslConfig struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata"`
	Spec          SslConfigSpec   `json:"spec"`
	Status        SslConfigStatus `json:"status,omitempty"`
}

type SslConfigSpec struct {
	Cert   string `json:"cert"`
	Key    string `json:"key"`
	Domain string `json:"domain"`
}

type SslConfigStatus struct {
	State   string `json:",omitempty"`
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SslConfigList struct {
	v1.TypeMeta `json:",inline"`
	v1.ListMeta `json:"metadata"`
	Items       []SslConfig `json:"items"`
}
