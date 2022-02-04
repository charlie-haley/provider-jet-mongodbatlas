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

type FeatureUsagesObservation struct {
	FeatureID map[string]string `json:"featureId,omitempty" tf:"feature_id,omitempty"`

	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`
}

type FeatureUsagesParameters struct {
}

type ProviderAccessObservation struct {
	AtlasAssumedRoleExternalID *string `json:"atlasAssumedRoleExternalId,omitempty" tf:"atlas_assumed_role_external_id,omitempty"`

	AtlasAwsAccountArn *string `json:"atlasAwsAccountArn,omitempty" tf:"atlas_aws_account_arn,omitempty"`

	AuthorizedDate *string `json:"authorizedDate,omitempty" tf:"authorized_date,omitempty"`

	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	FeatureUsages []FeatureUsagesObservation `json:"featureUsages,omitempty" tf:"feature_usages,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`
}

type ProviderAccessParameters struct {

	// +kubebuilder:validation:Optional
	IAMAssumedRoleArn *string `json:"iamAssumedRoleArn,omitempty" tf:"iam_assumed_role_arn,omitempty"`

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
}

// ProviderAccessSpec defines the desired state of ProviderAccess
type ProviderAccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProviderAccessParameters `json:"forProvider"`
}

// ProviderAccessStatus defines the observed state of ProviderAccess.
type ProviderAccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProviderAccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderAccess is the Schema for the ProviderAccesss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type ProviderAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProviderAccessSpec   `json:"spec"`
	Status            ProviderAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderAccessList contains a list of ProviderAccesss
type ProviderAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderAccess `json:"items"`
}

// Repository type metadata.
var (
	ProviderAccess_Kind             = "ProviderAccess"
	ProviderAccess_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProviderAccess_Kind}.String()
	ProviderAccess_KindAPIVersion   = ProviderAccess_Kind + "." + CRDGroupVersion.String()
	ProviderAccess_GroupVersionKind = CRDGroupVersion.WithKind(ProviderAccess_Kind)
)

func init() {
	SchemeBuilder.Register(&ProviderAccess{}, &ProviderAccessList{})
}
