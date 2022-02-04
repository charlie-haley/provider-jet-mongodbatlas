/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EndpointServiceAdlObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointServiceAdlParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	EndpointID *string `json:"endpointId" tf:"endpoint_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-mongodbatlas/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ProviderName *string `json:"providerName" tf:"provider_name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// EndpointServiceAdlSpec defines the desired state of EndpointServiceAdl
type EndpointServiceAdlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointServiceAdlParameters `json:"forProvider"`
}

// EndpointServiceAdlStatus defines the observed state of EndpointServiceAdl.
type EndpointServiceAdlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointServiceAdlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointServiceAdl is the Schema for the EndpointServiceAdls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type EndpointServiceAdl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointServiceAdlSpec   `json:"spec"`
	Status            EndpointServiceAdlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointServiceAdlList contains a list of EndpointServiceAdls
type EndpointServiceAdlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointServiceAdl `json:"items"`
}

// Repository type metadata.
var (
	EndpointServiceAdl_Kind             = "EndpointServiceAdl"
	EndpointServiceAdl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointServiceAdl_Kind}.String()
	EndpointServiceAdl_KindAPIVersion   = EndpointServiceAdl_Kind + "." + CRDGroupVersion.String()
	EndpointServiceAdl_GroupVersionKind = CRDGroupVersion.WithKind(EndpointServiceAdl_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointServiceAdl{}, &EndpointServiceAdlList{})
}
