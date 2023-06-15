// +kubebuilder:rbac:groups=extend.oam.dev,resources=applicationgroups,verbs=get;list;watch;create;update;patch;delete;deletecollection
// +kubebuilder:rbac:groups=core.oam.dev,resources=applications;applicationrevisions,verbs=get;list;watch;create;update;patch;delete;deletecollection
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=*
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:resource:scope=Namespaced,categories={oam-extend},shortName=appgroup,singular=applicationgroups
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="version",type=string,JSONPath=".spec.version"
// +kubebuilder:printcolumn:name="dependon",type=string,JSONPath=".spec.dependon"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApplicationGroup declare applicationg group api
type ApplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationGroupSpec   `json:"spec"`
	Status            ApplicationGroupStatus `json:"status"`
}

// ApplicationGroupSpec describe application group specification
type ApplicationGroupSpec struct {
	// +kubebuilder:validation:Required
	Version  string   `json:"version"`
	DependOn []string `json:"dependon, omitempty"`
}

type ApplicationGroupStatus struct {
	Ready bool `json:"ready"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppleList is a list of Apple resources
type ApplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ApplicationGroup `json:"items"`
}
