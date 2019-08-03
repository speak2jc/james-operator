package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExamplekindSpec defines the desired state of Examplekind
type ExamplekindSpec struct {
	Count int32 `json:"count"`
	Group string `json:"group"`
	Image string `json:"image"`
	Port int32 `json:"port"`
}

// ExamplekindStatus defines the observed state of Examplekind
type ExamplekindStatus struct {
	PodNames []string `json:"podnames"`
	AppGroup string `json:"appgroup"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Examplekind is the Schema for the examplekinds API
// +k8s:openapi-gen=true
type Examplekind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExamplekindSpec   `json:"spec,omitempty"`
	Status ExamplekindStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExamplekindList contains a list of Examplekind
type ExamplekindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Examplekind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Examplekind{}, &ExamplekindList{})
}
