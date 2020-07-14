package v1alpha1

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient:noStatus
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GroupTask represents a group of Tasks run together.
//
// +k8s:openapi-gen=true
type GroupTask struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec GroupTaskSpec `json:"spec"`
}

type GroupTaskSpec struct {
	// +optional
	Tasks []v1alpha1.Task `json:"tasks,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GroupTaskList contains a list of GroupTasks.
type GroupTaskList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupTask `json:"items"`
}
