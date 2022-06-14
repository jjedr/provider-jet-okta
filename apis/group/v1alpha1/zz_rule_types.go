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

type RuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleParameters struct {

	// +kubebuilder:validation:Optional
	ExpressionType *string `json:"expressionType,omitempty" tf:"expression_type,omitempty"`

	// +kubebuilder:validation:Required
	ExpressionValue *string `json:"expressionValue" tf:"expression_value,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-okta/apis/group/v1alpha1.Group
	// +kubebuilder:validation:Optional
	GroupAssignments []*string `json:"groupAssignments,omitempty" tf:"group_assignments,omitempty"`

	// +kubebuilder:validation:Optional
	GroupAssignmentsRefs []v1.Reference `json:"groupAssignmentsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GroupAssignmentsSelector *v1.Selector `json:"groupAssignmentsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Remove users added by this rule from the assigned group after deleting this resource
	// +kubebuilder:validation:Optional
	RemoveAssignedUsers *bool `json:"removeAssignedUsers,omitempty" tf:"remove_assigned_users,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The list of user IDs that would be excluded when rules are processed
	// +kubebuilder:validation:Optional
	UsersExcluded []*string `json:"usersExcluded,omitempty" tf:"users_excluded,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rule is the Schema for the Rules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oktajet}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSpec   `json:"spec"`
	Status            RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}