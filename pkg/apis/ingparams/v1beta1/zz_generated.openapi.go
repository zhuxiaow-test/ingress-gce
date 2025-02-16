//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Kubernetes Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"k8s.io/ingress-gce/pkg/apis/ingparams/v1beta1.GCPIngressParams":     schema_pkg_apis_ingparams_v1beta1_GCPIngressParams(ref),
		"k8s.io/ingress-gce/pkg/apis/ingparams/v1beta1.GCPIngressParamsSpec": schema_pkg_apis_ingparams_v1beta1_GCPIngressParamsSpec(ref),
	}
}

func schema_pkg_apis_ingparams_v1beta1_GCPIngressParams(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GCPIngressParams contains the parameters for GCP Ingress Classes",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/ingress-gce/pkg/apis/ingparams/v1beta1.GCPIngressParamsSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/ingress-gce/pkg/apis/ingparams/v1beta1.GCPIngressParamsStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "k8s.io/ingress-gce/pkg/apis/ingparams/v1beta1.GCPIngressParamsSpec", "k8s.io/ingress-gce/pkg/apis/ingparams/v1beta1.GCPIngressParamsStatus"},
	}
}

func schema_pkg_apis_ingparams_v1beta1_GCPIngressParamsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GCPIngressParamsSpec is the spec for a GCPIngressParams resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"internal": {
						SchemaProps: spec.SchemaProps{
							Description: "Internal specifies whether internal or external load balancing is desired. The default is external load balancing, so Internal will default to false.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"internal"},
			},
		},
	}
}
