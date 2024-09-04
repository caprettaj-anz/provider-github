// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RepositoryFileInitParameters struct {

	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	// The branch name, defaults to the repository's default branch
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/repo/v1alpha1.Branch
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Reference to a Branch in repo to populate branch.
	// +kubebuilder:validation:Optional
	BranchRef *v1.Reference `json:"branchRef,omitempty" tf:"-"`

	// Selector for a Branch in repo to populate branch.
	// +kubebuilder:validation:Optional
	BranchSelector *v1.Selector `json:"branchSelector,omitempty" tf:"-"`

	// Committer author name to use. NOTE: GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	// The commit author name, defaults to the authenticated user's name. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	CommitAuthor *string `json:"commitAuthor,omitempty" tf:"commit_author,omitempty"`

	// Committer email address to use. NOTE: GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	// The commit author email address, defaults to the authenticated user's email address. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	CommitEmail *string `json:"commitEmail,omitempty" tf:"commit_email,omitempty"`

	// The commit message when creating, updating or deleting the managed file.
	// The commit message when creating, updating or deleting the file
	CommitMessage *string `json:"commitMessage,omitempty" tf:"commit_message,omitempty"`

	// The file content.
	// The file's content
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The path of the file to manage.
	// The file path to manage
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// Enable overwriting existing files. If set to true it will overwrite an existing file with the same name. If set to false it will fail if there is an existing file with the same name.
	// Enable overwriting existing files, defaults to "false"
	OverwriteOnCreate *bool `json:"overwriteOnCreate,omitempty" tf:"overwrite_on_create,omitempty"`

	// The repository to create the file in.
	// The repository name
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

type RepositoryFileObservation struct {

	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	// The branch name, defaults to the repository's default branch
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Committer author name to use. NOTE: GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	// The commit author name, defaults to the authenticated user's name. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	CommitAuthor *string `json:"commitAuthor,omitempty" tf:"commit_author,omitempty"`

	// Committer email address to use. NOTE: GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	// The commit author email address, defaults to the authenticated user's email address. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	CommitEmail *string `json:"commitEmail,omitempty" tf:"commit_email,omitempty"`

	// The commit message when creating, updating or deleting the managed file.
	// The commit message when creating, updating or deleting the file
	CommitMessage *string `json:"commitMessage,omitempty" tf:"commit_message,omitempty"`

	// The SHA of the commit that modified the file.
	// The SHA of the commit that modified the file
	CommitSha *string `json:"commitSha,omitempty" tf:"commit_sha,omitempty"`

	// The file content.
	// The file's content
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The path of the file to manage.
	// The file path to manage
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Enable overwriting existing files. If set to true it will overwrite an existing file with the same name. If set to false it will fail if there is an existing file with the same name.
	// Enable overwriting existing files, defaults to "false"
	OverwriteOnCreate *bool `json:"overwriteOnCreate,omitempty" tf:"overwrite_on_create,omitempty"`

	// The name of the commit/branch/tag.
	// The name of the commit/branch/tag
	Ref *string `json:"ref,omitempty" tf:"ref,omitempty"`

	// The repository to create the file in.
	// The repository name
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// The SHA blob of the file.
	// The blob SHA of the file
	Sha *string `json:"sha,omitempty" tf:"sha,omitempty"`
}

type RepositoryFileParameters struct {

	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	// The branch name, defaults to the repository's default branch
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/repo/v1alpha1.Branch
	// +kubebuilder:validation:Optional
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Reference to a Branch in repo to populate branch.
	// +kubebuilder:validation:Optional
	BranchRef *v1.Reference `json:"branchRef,omitempty" tf:"-"`

	// Selector for a Branch in repo to populate branch.
	// +kubebuilder:validation:Optional
	BranchSelector *v1.Selector `json:"branchSelector,omitempty" tf:"-"`

	// Committer author name to use. NOTE: GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	// The commit author name, defaults to the authenticated user's name. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	// +kubebuilder:validation:Optional
	CommitAuthor *string `json:"commitAuthor,omitempty" tf:"commit_author,omitempty"`

	// Committer email address to use. NOTE: GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	// The commit author email address, defaults to the authenticated user's email address. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	// +kubebuilder:validation:Optional
	CommitEmail *string `json:"commitEmail,omitempty" tf:"commit_email,omitempty"`

	// The commit message when creating, updating or deleting the managed file.
	// The commit message when creating, updating or deleting the file
	// +kubebuilder:validation:Optional
	CommitMessage *string `json:"commitMessage,omitempty" tf:"commit_message,omitempty"`

	// The file content.
	// The file's content
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The path of the file to manage.
	// The file path to manage
	// +kubebuilder:validation:Optional
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// Enable overwriting existing files. If set to true it will overwrite an existing file with the same name. If set to false it will fail if there is an existing file with the same name.
	// Enable overwriting existing files, defaults to "false"
	// +kubebuilder:validation:Optional
	OverwriteOnCreate *bool `json:"overwriteOnCreate,omitempty" tf:"overwrite_on_create,omitempty"`

	// The repository to create the file in.
	// The repository name
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

// RepositoryFileSpec defines the desired state of RepositoryFile
type RepositoryFileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryFileParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RepositoryFileInitParameters `json:"initProvider,omitempty"`
}

// RepositoryFileStatus defines the observed state of RepositoryFile.
type RepositoryFileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryFileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryFile is the Schema for the RepositoryFiles API. Creates and manages files within a GitHub repository
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type RepositoryFile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.file) || (has(self.initProvider) && has(self.initProvider.file))",message="spec.forProvider.file is a required parameter"
	Spec   RepositoryFileSpec   `json:"spec"`
	Status RepositoryFileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryFileList contains a list of RepositoryFiles
type RepositoryFileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryFile `json:"items"`
}

// Repository type metadata.
var (
	RepositoryFile_Kind             = "RepositoryFile"
	RepositoryFile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryFile_Kind}.String()
	RepositoryFile_KindAPIVersion   = RepositoryFile_Kind + "." + CRDGroupVersion.String()
	RepositoryFile_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryFile_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryFile{}, &RepositoryFileList{})
}
