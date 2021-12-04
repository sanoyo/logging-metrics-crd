/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LogSyncSpec defines the desired state of LogSync
type LogSyncSpec struct {
	ProjectID   string `json:"projectID"`
	Destination string `json:"destination"`
	Filter      string `json:"filter"`
}

// LogSyncStatus defines the observed state of LogSync
type LogSyncStatus struct {
	ProjectIp   string `json:"projectID"`
	Destination string `json:"destination"`
	Filter      string `json:"filter"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LogSync is the Schema for the logsyncs API
type LogSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogSyncSpec   `json:"spec,omitempty"`
	Status LogSyncStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LogSyncList contains a list of LogSync
type LogSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogSync{}, &LogSyncList{})
}
