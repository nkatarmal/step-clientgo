package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// tp
type Step struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StepSpec   `json:"spec"`
	Status StepStatus `json:"status"`
}

type StepSpec struct {
	Namespace string `json:"namespace"`
}

type StepStatus struct {
	LastExecuted *metav1.Time `json:"lastExecuted,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//tp
type StepList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Step `json:"items"`
}
