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

type BackupSnapshotRestoreJobObservation struct {
	Cancelled *bool `json:"cancelled,omitempty" tf:"cancelled,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DeliveryURL []*string `json:"deliveryUrl,omitempty" tf:"delivery_url,omitempty"`

	Expired *bool `json:"expired,omitempty" tf:"expired,omitempty"`

	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	FinishedAt *string `json:"finishedAt,omitempty" tf:"finished_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SnapshotRestoreJobID *string `json:"snapshotRestoreJobId,omitempty" tf:"snapshot_restore_job_id,omitempty"`

	Timestamp *string `json:"timestamp,omitempty" tf:"timestamp,omitempty"`
}

type BackupSnapshotRestoreJobParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	DeliveryType map[string]*string `json:"deliveryType,omitempty" tf:"delivery_type,omitempty"`

	// +kubebuilder:validation:Optional
	DeliveryTypeConfig []DeliveryTypeConfigParameters `json:"deliveryTypeConfig,omitempty" tf:"delivery_type_config,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-mongodbatlas/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SnapshotID *string `json:"snapshotId" tf:"snapshot_id,omitempty"`
}

type DeliveryTypeConfigObservation struct {
}

type DeliveryTypeConfigParameters struct {

	// +kubebuilder:validation:Optional
	Automated *bool `json:"automated,omitempty" tf:"automated,omitempty"`

	// +kubebuilder:validation:Optional
	Download *bool `json:"download,omitempty" tf:"download,omitempty"`

	// +kubebuilder:validation:Optional
	OplogInc *int64 `json:"oplogInc,omitempty" tf:"oplog_inc,omitempty"`

	// +kubebuilder:validation:Optional
	OplogTS *int64 `json:"oplogTs,omitempty" tf:"oplog_ts,omitempty"`

	// +kubebuilder:validation:Optional
	PointInTime *bool `json:"pointInTime,omitempty" tf:"point_in_time,omitempty"`

	// +kubebuilder:validation:Optional
	PointInTimeUtcSeconds *int64 `json:"pointInTimeUtcSeconds,omitempty" tf:"point_in_time_utc_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	TargetClusterName *string `json:"targetClusterName,omitempty" tf:"target_cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	TargetProjectID *string `json:"targetProjectId,omitempty" tf:"target_project_id,omitempty"`
}

// BackupSnapshotRestoreJobSpec defines the desired state of BackupSnapshotRestoreJob
type BackupSnapshotRestoreJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupSnapshotRestoreJobParameters `json:"forProvider"`
}

// BackupSnapshotRestoreJobStatus defines the observed state of BackupSnapshotRestoreJob.
type BackupSnapshotRestoreJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupSnapshotRestoreJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSnapshotRestoreJob is the Schema for the BackupSnapshotRestoreJobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type BackupSnapshotRestoreJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSnapshotRestoreJobSpec   `json:"spec"`
	Status            BackupSnapshotRestoreJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSnapshotRestoreJobList contains a list of BackupSnapshotRestoreJobs
type BackupSnapshotRestoreJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSnapshotRestoreJob `json:"items"`
}

// Repository type metadata.
var (
	BackupSnapshotRestoreJob_Kind             = "BackupSnapshotRestoreJob"
	BackupSnapshotRestoreJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupSnapshotRestoreJob_Kind}.String()
	BackupSnapshotRestoreJob_KindAPIVersion   = BackupSnapshotRestoreJob_Kind + "." + CRDGroupVersion.String()
	BackupSnapshotRestoreJob_GroupVersionKind = CRDGroupVersion.WithKind(BackupSnapshotRestoreJob_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupSnapshotRestoreJob{}, &BackupSnapshotRestoreJobList{})
}
