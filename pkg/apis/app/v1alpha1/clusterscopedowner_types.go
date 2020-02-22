package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterScopedOwnerSpec defines the desired state of ClusterScopedOwner
type ClusterScopedOwnerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// ClusterScopedOwnerStatus defines the observed state of ClusterScopedOwner
type ClusterScopedOwnerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterScopedOwner is the Schema for the clusterscopedowners API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=clusterscopedowners,scope=Cluster
type ClusterScopedOwner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterScopedOwnerSpec   `json:"spec,omitempty"`
	Status ClusterScopedOwnerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterScopedOwnerList contains a list of ClusterScopedOwner
type ClusterScopedOwnerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterScopedOwner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterScopedOwner{}, &ClusterScopedOwnerList{})
}
