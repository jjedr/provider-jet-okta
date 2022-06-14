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

package app

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("okta_app_oauth", func(r *tjconfig.Resource) {
		r.ShortGroup = "app"
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["okta_app_oauth_id"] = []byte(a)
			}
			if a, ok := attr["client_id"].(string); ok {
				conn["okta_app_oauth_client_id"] = []byte(a)
			}
			if a, ok := attr["client_secret"].(string); ok {
				conn["okta_app_oauth_client_secret"] = []byte(a)
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("okta_app_oauth_api_scope", func(r *tjconfig.Resource) {
		r.ShortGroup = "app"
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.References["app_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-okta/apis/app/v1alpha1.Oauth",
		}
	})
	p.AddResourceConfigurator("okta_app_group_assignments", func(r *tjconfig.Resource) {
		r.ShortGroup = "app"
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.References = tjconfig.References{
			"app_id": tjconfig.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-okta/apis/app/v1alpha1.Oauth",
			},
			"group.id": tjconfig.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-okta/apis/group/v1alpha1.Group",
			},
		}
		
	})

	
}
