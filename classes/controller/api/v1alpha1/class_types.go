/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClassSpec defines the desired state of Class
type ClassSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name     string   `json:"name"`
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}

// ClassStatus defines the observed state of Class
type ClassStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Available       bool     `json:"available"`
	PresentStudents []string `json:"presentstudents"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Teacher",type="string",JSONPath=".spec.teacher",description="The teacher of this class"
// +kubebuilder:printcolumn:name="Availability",type="boolean",JSONPath=".status.available",description="The availability of this class"


// Class is the Schema for the classes API
type Class struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClassSpec   `json:"spec,omitempty"`
	Status ClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClassList contains a list of Class
type ClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Class `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Class{}, &ClassList{})
}
