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

type OauthAPIScopeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OauthAPIScopeParameters struct {

	// ID of the application.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-okta/apis/app/v1alpha1.Oauth
	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Optional
	AppIDRef *v1.Reference `json:"appIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AppIDSelector *v1.Selector `json:"appIdSelector,omitempty" tf:"-"`

	// The issuer of your Org Authorization Server, your Org URL.
	// +kubebuilder:validation:Required
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// Scopes of the application for which consent is granted.
	// +kubebuilder:validation:Required
	Scopes []*string `json:"scopes" tf:"scopes,omitempty"`
}

// OauthAPIScopeSpec defines the desired state of OauthAPIScope
type OauthAPIScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OauthAPIScopeParameters `json:"forProvider"`
}

// OauthAPIScopeStatus defines the observed state of OauthAPIScope.
type OauthAPIScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OauthAPIScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OauthAPIScope is the Schema for the OauthAPIScopes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oktajet}
type OauthAPIScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OauthAPIScopeSpec   `json:"spec"`
	Status            OauthAPIScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OauthAPIScopeList contains a list of OauthAPIScopes
type OauthAPIScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OauthAPIScope `json:"items"`
}

// Repository type metadata.
var (
	OauthAPIScope_Kind             = "OauthAPIScope"
	OauthAPIScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OauthAPIScope_Kind}.String()
	OauthAPIScope_KindAPIVersion   = OauthAPIScope_Kind + "." + CRDGroupVersion.String()
	OauthAPIScope_GroupVersionKind = CRDGroupVersion.WithKind(OauthAPIScope_Kind)
)

func init() {
	SchemeBuilder.Register(&OauthAPIScope{}, &OauthAPIScopeList{})
}
