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

type MembersInitParameters struct {

	// List of team members. See Members below for details.
	// List of team members.
	Members []MembersMembersInitParameters `json:"members,omitempty" tf:"members,omitempty"`

	// The team id or the team slug
	// The GitHub team id or slug
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/team/v1alpha1.Team
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`
}

type MembersMembersInitParameters struct {

	// The role of the user within the team.
	// Must be one of member or maintainer. Defaults to member.
	// The role of the user within the team. Must be one of 'member' or 'maintainer'.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The user to add to the team.
	// The user to add to the team.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type MembersMembersObservation struct {

	// The role of the user within the team.
	// Must be one of member or maintainer. Defaults to member.
	// The role of the user within the team. Must be one of 'member' or 'maintainer'.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The user to add to the team.
	// The user to add to the team.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type MembersMembersParameters struct {

	// The role of the user within the team.
	// Must be one of member or maintainer. Defaults to member.
	// The role of the user within the team. Must be one of 'member' or 'maintainer'.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The user to add to the team.
	// The user to add to the team.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type MembersObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of team members. See Members below for details.
	// List of team members.
	Members []MembersMembersObservation `json:"members,omitempty" tf:"members,omitempty"`

	// The team id or the team slug
	// The GitHub team id or slug
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

type MembersParameters struct {

	// List of team members. See Members below for details.
	// List of team members.
	// +kubebuilder:validation:Optional
	Members []MembersMembersParameters `json:"members,omitempty" tf:"members,omitempty"`

	// The team id or the team slug
	// The GitHub team id or slug
	// +crossplane:generate:reference:type=github.com/xunholy/provider-github/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`
}

// MembersSpec defines the desired state of Members
type MembersSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MembersParameters `json:"forProvider"`
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
	InitProvider MembersInitParameters `json:"initProvider,omitempty"`
}

// MembersStatus defines the observed state of Members.
type MembersStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MembersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Members is the Schema for the Memberss API. Provides an authoritative GitHub team members resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type Members struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	Spec   MembersSpec   `json:"spec"`
	Status MembersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MembersList contains a list of Memberss
type MembersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Members `json:"items"`
}

// Repository type metadata.
var (
	Members_Kind             = "Members"
	Members_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Members_Kind}.String()
	Members_KindAPIVersion   = Members_Kind + "." + CRDGroupVersion.String()
	Members_GroupVersionKind = CRDGroupVersion.WithKind(Members_Kind)
)

func init() {
	SchemeBuilder.Register(&Members{}, &MembersList{})
}
