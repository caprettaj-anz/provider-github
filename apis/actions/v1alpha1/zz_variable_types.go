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

type VariableInitParameters struct {

	// Name of the repository
	// Name of the repository.
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Value of the variable
	// Value of the variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Name of the variable
	// Name of the variable.
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

type VariableObservation struct {

	// Date of actions_variable creation.
	// Date of 'actions_variable' creation.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the repository
	// Name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Date of actions_variable update.
	// Date of 'actions_variable' update.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// Value of the variable
	// Value of the variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Name of the variable
	// Name of the variable.
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

type VariableParameters struct {

	// Name of the repository
	// Name of the repository.
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/repo/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Value of the variable
	// Value of the variable.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Name of the variable
	// Name of the variable.
	// +kubebuilder:validation:Optional
	VariableName *string `json:"variableName,omitempty" tf:"variable_name,omitempty"`
}

// VariableSpec defines the desired state of Variable
type VariableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VariableParameters `json:"forProvider"`
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
	InitProvider VariableInitParameters `json:"initProvider,omitempty"`
}

// VariableStatus defines the observed state of Variable.
type VariableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VariableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Variable is the Schema for the Variables API. Creates and manages an Action variable within a GitHub repository
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type Variable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.variableName) || (has(self.initProvider) && has(self.initProvider.variableName))",message="spec.forProvider.variableName is a required parameter"
	Spec   VariableSpec   `json:"spec"`
	Status VariableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VariableList contains a list of Variables
type VariableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Variable `json:"items"`
}

// Repository type metadata.
var (
	Variable_Kind             = "Variable"
	Variable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Variable_Kind}.String()
	Variable_KindAPIVersion   = Variable_Kind + "." + CRDGroupVersion.String()
	Variable_GroupVersionKind = CRDGroupVersion.WithKind(Variable_Kind)
)

func init() {
	SchemeBuilder.Register(&Variable{}, &VariableList{})
}
