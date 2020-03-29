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
		"github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.Service":       schema_pkg_apis_serving_v1alpha1_Service(ref),
		"github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.ServiceSpec":   schema_pkg_apis_serving_v1alpha1_ServiceSpec(ref),
		"github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.ServiceStatus": schema_pkg_apis_serving_v1alpha1_ServiceStatus(ref),
	}
}

func schema_pkg_apis_serving_v1alpha1_Service(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Service is the Schema for the services API",
				Type:        []string{"object"},
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
							Ref: ref("github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.ServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.ServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.ServiceSpec", "github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.ServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_serving_v1alpha1_ServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceSpec defines the desired state of Service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"generation": {
						SchemaProps: spec.SchemaProps{
							Description: "DeprecatedGeneration was used prior in Kubernetes versions <1.11 when metadata.generation was not being incremented by the api server\n\nThis property will be dropped in future Knative releases and should not be used - use metadata.generation\n\nTracking issue: https://github.com/knative/serving/issues/643",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Description: "Template holds the latest specification for the Revision to be stamped out.",
							Ref:         ref("github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.RevisionTemplateSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wso2/k8s-api-operator/apim-operator/pkg/apis/serving/v1alpha1.RevisionTemplateSpec"},
	}
}

func schema_pkg_apis_serving_v1alpha1_ServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceStatus defines the observed state of Service",
				Type:        []string{"object"},
			},
		},
	}
}
