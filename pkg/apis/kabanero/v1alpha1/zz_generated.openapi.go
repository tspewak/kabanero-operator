// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.Kabanero":       schema_pkg_apis_kabanero_v1alpha1_Kabanero(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroSpec":   schema_pkg_apis_kabanero_v1alpha1_KabaneroSpec(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroStatus": schema_pkg_apis_kabanero_v1alpha1_KabaneroStatus(ref),
	}
}

func schema_pkg_apis_kabanero_v1alpha1_Kabanero(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kabanero is the Schema for the kabaneros API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kabanero_v1alpha1_KabaneroSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KabaneroSpec defines the desired state of Kabanero",
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"tekton": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonCustomizationSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonCustomizationSpec"},
	}
}

func schema_pkg_apis_kabanero_v1alpha1_KabaneroStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KabaneroStatus defines the observed state of Kabanero",
				Properties: map[string]spec.Schema{
					"tekton": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonStatus"},
	}
}
