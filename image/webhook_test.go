package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLoadConfig(t *testing.T) {
  trueVal := true
	ndotsVal := "3"
	topologyHonorPolicy := corev1.NodeInclusionPolicyHonor
	testEnvFrom := corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "stage-connect-secrets"}, Optional: &trueVal}
	files := []struct {
		name string
		env  *Config
	}{
		{"test/env_test_1.yaml",
			&Config{
				[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}},
				nil, nil, nil, nil, nil, nil, false, // nil = empty tests, false for boolean not defined
			},
		},
		{"test/env_test_1_1.yaml",
			&Config{
				[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}},
				[]corev1.EnvFromSource{{SecretRef: &testEnvFrom}},
				nil, nil, nil, nil, nil, false, // nil = empty tests, false for boolean not defined
			},
		},
		{"test/env_test_2.yaml",
			&Config{
				[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil},
					{Name: "SUBSCRIPTION", Value: "subscription-00", ValueFrom: nil}},
				nil, nil, nil, nil, nil, nil, false, // nil = empty tests, false for boolean not defined
			},
		},
		{"test/env_test_3.yaml",
			&Config{
				[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil},
					{Name: "SUBSCRIPTION", Value: "subscription-00", ValueFrom: nil}}, nil,
				[]corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal},
					{Name: "single-request-reopen", Value: nil},
					{Name: "use-vc", Value: nil}},
				nil, nil, nil, nil, false, // nil = empty tests, false for boolean not defined
			},
		},
		{"test/env_test_4.yaml",
			&Config{
				[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil},
					{Name: "SUBSCRIPTION", Value: "subscription-00", ValueFrom: nil}}, nil,
				[]corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal},
					{Name: "single-request-reopen", Value: nil},
					{Name: "use-vc", Value: nil}},
				[]corev1.NodeSelectorTerm{{
					MatchExpressions: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
					}},
				}},
				nil, nil, nil, false, // nil = empty tests, false for boolean not defined
			},
		},
		{"test/env_test_5.yaml",
			&Config{[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil},
				{Name: "SUBSCRIPTION", Value: "subscription-00", ValueFrom: nil}}, nil,
				[]corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal},
					{Name: "single-request-reopen", Value: nil},
					{Name: "use-vc", Value: nil}},
				[]corev1.NodeSelectorTerm{{
					MatchExpressions: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
					}},
				}},
				[]corev1.PreferredSchedulingTerm{{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchExpressions: []corev1.NodeSelectorRequirement{{
							Key: "kubernetes.azure.com/scalesetpriority", Operator: corev1.NodeSelectorOpDoesNotExist,
						}},
					},
				}},
				nil, nil, false, // nil = empty tests, false for boolean not defined
			},
		},
		{"test/env_test_6.yaml",
			&Config{
				[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil},
					{Name: "SUBSCRIPTION", Value: "subscription-00", ValueFrom: nil}}, nil,
				[]corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal},
					{Name: "single-request-reopen", Value: nil},
					{Name: "use-vc", Value: nil}},
				[]corev1.NodeSelectorTerm{{MatchExpressions: []corev1.NodeSelectorRequirement{
					{Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"}}}}},
				[]corev1.PreferredSchedulingTerm{{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchExpressions: []corev1.NodeSelectorRequirement{{
							Key: "kubernetes.azure.com/scalesetpriority", Operator: corev1.NodeSelectorOpDoesNotExist,
						}},
					},
				}},
				[]corev1.Toleration{{
					Key:      "kubernetes.azure.com/scalesetpriority",
					Effect:   "NoSchedule",
					Operator: "Equal",
					Value:    "spot",
				}},
				nil, false, // nil = empty tests, false for boolean not defined
			},
		},
		{"test/env_test_7.yaml",
			&Config{
				[]corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil},
					{Name: "SUBSCRIPTION", Value: "subscription-00", ValueFrom: nil}}, nil,
				[]corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal},
					{Name: "single-request-reopen", Value: nil},
					{Name: "use-vc", Value: nil}},
				[]corev1.NodeSelectorTerm{{MatchExpressions: []corev1.NodeSelectorRequirement{
					{Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"}}}}},
				[]corev1.PreferredSchedulingTerm{{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchExpressions: []corev1.NodeSelectorRequirement{{
							Key: "kubernetes.azure.com/scalesetpriority", Operator: corev1.NodeSelectorOpDoesNotExist,
						}},
					},
				}},
				[]corev1.Toleration{{
					Key:      "kubernetes.azure.com/scalesetpriority",
					Effect:   "NoSchedule",
					Operator: "Equal",
					Value:    "spot",
				}},
				[]corev1.TopologySpreadConstraint{{
					MaxSkew:            1,
					TopologyKey:        "topology.kubernetes.io/zone",
					NodeAffinityPolicy: &topologyHonorPolicy,
					NodeTaintsPolicy:   &topologyHonorPolicy,
					WhenUnsatisfiable:  "ScheduleAnyway",
					LabelSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app.kubernetes.io/name": "test-app",
						},
					},
					MatchLabelKeys: []string{
						"pod-template-hash",
					},
				}},
				true,
			},
		},
	}

	for _, f := range files {
		config, err := loadConfig(f.name)
		if err != nil {
			t.Errorf("Error loading file %s", f.name)
			t.Fatal(err)
		}
		if !cmp.Equal(config, f.env) {
			t.Errorf("loadConfig was incorrect, got: %v, want: %v.", config, f.env)
		}
	}
}

func TestMutationRequired(t *testing.T) {
	metas := []struct {
		ignoredList []string
		metadata    *metav1.ObjectMeta
		required    bool
	}{
		{[]string{"admin", "kube-system"},
			&metav1.ObjectMeta{Namespace: "admin", Annotations: map[string]string{}},
			false},
		{[]string{"admin", "kube-system"},
			&metav1.ObjectMeta{Namespace: "rpe", Annotations: map[string]string{"some-other-annotation/inject": "false"}},
			true},
		{[]string{"admin", "kube-system"},
			&metav1.ObjectMeta{Namespace: "rpe", Annotations: map[string]string{admissionWebhookAnnotationStatusKey: "injected"}},
			false},
		{[]string{"admin", "kube-system"},
			&metav1.ObjectMeta{Namespace: "rpe", Annotations: map[string]string{admissionWebhookAnnotationInjectKey: "false"}},
			false},
		{[]string{"admin", "kube-system"},
			&metav1.ObjectMeta{Namespace: "rpe", Annotations: map[string]string{}},
			true},
		{[]string{"admin", "kube-system"},
			&metav1.ObjectMeta{Namespace: "rpe", Annotations: map[string]string{admissionWebhookAnnotationInjectKey: "true"}},
			true},
	}

	for _, m := range metas {
		required := mutationRequired(m.ignoredList, m.metadata)
		if required != m.required {
			t.Errorf("mutationRequired was incorrect, for %v, got: %t, want: %t.", m, required, m.required)
		}
	}
}

func TestAddEnv(t *testing.T) {
	envs := []struct {
		targetEnv []corev1.EnvVar
		sourceEnv []corev1.EnvVar
		path      string
		patch     []patchOperation
	}{
		{
			targetEnv: []corev1.EnvVar{{Name: "ENV_TEST_NAME", Value: "env-test-value", ValueFrom: nil}},
			sourceEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}},
			path:      "/spec/containers/nginx/env",
			patch: []patchOperation{
				{Op: "add", Path: "/spec/containers/nginx/env/-", Value: corev1.EnvVar{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}},
			},
		},
		{
			targetEnv: []corev1.EnvVar{},
			sourceEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}},
			path:      "/spec/containers/nginx/env",
			patch: []patchOperation{
				{Op: "add", Path: "/spec/containers/nginx/env", Value: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}}},
			},
		},
		{
			targetEnv: []corev1.EnvVar{{Name: "ENV_TEST_NAME", Value: "env-test-value", ValueFrom: nil}},
			sourceEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}, {Name: "SUBSCRIPTION", Value: "subscription-01", ValueFrom: nil}},
			path:      "/spec/containers/nginx/env",
			patch: []patchOperation{
				{Op: "add", Path: "/spec/containers/nginx/env/-", Value: corev1.EnvVar{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}},
				{Op: "add", Path: "/spec/containers/nginx/env/-", Value: corev1.EnvVar{Name: "SUBSCRIPTION", Value: "subscription-01", ValueFrom: nil}},
			},
		},
		{
			targetEnv: nil,
			sourceEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}, {Name: "SUBSCRIPTION", Value: "subscription-01", ValueFrom: nil}},
			path:      "/spec/containers/nginx/env",
			patch: []patchOperation{
				{Op: "add", Path: "/spec/containers/nginx/env", Value: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}}},
				{Op: "add", Path: "/spec/containers/nginx/env/-", Value: corev1.EnvVar{Name: "SUBSCRIPTION", Value: "subscription-01", ValueFrom: nil}},
			},
		},
		{
			targetEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}, {Name: "SUBSCRIPTION", Value: "subscription-01", ValueFrom: nil}},
			sourceEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}, {Name: "SUBSCRIPTION", Value: "subscription-01", ValueFrom: nil}},
			path:      "/spec/containers/nginx/env",
			patch:     []patchOperation{},
		},
		{
			targetEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-01", ValueFrom: nil}, {Name: "SUBSCRIPTION", Value: "subscription-01", ValueFrom: nil}},
			sourceEnv: []corev1.EnvVar{{Name: "CLUSTER_NAME", Value: "aks-test-02", ValueFrom: nil}, {Name: "SUBSCRIPTION", Value: "subscription-02", ValueFrom: nil}},
			path:      "/spec/containers/nginx/env",
			patch: []patchOperation{
				{Op: "replace", Path: "/spec/containers/nginx/env/0", Value: corev1.EnvVar{Name: "CLUSTER_NAME", Value: "aks-test-02", ValueFrom: nil}},
				{Op: "replace", Path: "/spec/containers/nginx/env/1", Value: corev1.EnvVar{Name: "SUBSCRIPTION", Value: "subscription-02", ValueFrom: nil}},
			},
		},
	}

	for _, e := range envs {
		patch := addEnv(e.targetEnv, e.sourceEnv, e.path)
		if !cmp.Equal(patch, e.patch) {
			t.Errorf("addEnv was incorrect, for %v, got: %v, want: %v.", e.targetEnv, patch, e.patch)
		}
	}
}


func TestAddEnvFrom(t *testing.T) {
  trueVal := true
  targetEnvFrom := corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "app-secrets"}, Optional: &trueVal}
  targetEnvFrom2 := corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "waiting-secrets"}, Optional: &trueVal}
  targetEnvFrom3 := corev1.ConfigMapEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "app-config"}, Optional: &trueVal}
  sourceEnvFrom := corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "stage-connect-secrets"}, Optional: &trueVal}
	envs := []struct {
		targetEnv []corev1.EnvFromSource
		sourceEnv []corev1.EnvFromSource
		path      string
		patch     []patchOperation
	}{
		{
			targetEnv: []corev1.EnvFromSource{{SecretRef: &targetEnvFrom}, {SecretRef: &targetEnvFrom2}, {ConfigMapRef: &targetEnvFrom3}},
			sourceEnv: []corev1.EnvFromSource{{SecretRef: &sourceEnvFrom}},
			path:      "/spec/containers/nginx/envFrom",
			patch: []patchOperation{
				{Op: "add", Path: "/spec/containers/nginx/envFrom/-", Value: corev1.EnvFromSource{SecretRef: &sourceEnvFrom}},
			},
		},
	}

	for _, e := range envs {
		patch := addEnvFrom(e.targetEnv, e.sourceEnv, e.path)
		if !cmp.Equal(patch, e.patch) {
			t.Errorf("addEnvFrom was incorrect, for %v, got: %v, want: %v.", e.targetEnv, patch, e.patch)
		}
	}
}


func TestAddDnsOptions(t *testing.T) {
	ndotsVal := "3"
	ndotsValOld := "5"
	envs := []struct {
		targetOptions []corev1.PodDNSConfigOption
		sourceOptions []corev1.PodDNSConfigOption
		path          string
		patch         []patchOperation
	}{
		{
			targetOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}},
			sourceOptions: []corev1.PodDNSConfigOption{{Name: "single-request-reopen", Value: nil}},
			path:          "/spec/dnsConfig/options",
			patch:         []patchOperation{{Op: "add", Path: "/spec/dnsConfig/options/-", Value: corev1.PodDNSConfigOption{Name: "single-request-reopen", Value: nil}}},
		},
		{
			targetOptions: []corev1.PodDNSConfigOption{},
			sourceOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}},
			path:          "/spec/dnsConfig/options",
			patch:         []patchOperation{{Op: "add", Path: "/spec/dnsConfig/options", Value: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}}}},
		},
		{
			targetOptions: []corev1.PodDNSConfigOption{{Name: "single-request-reopen", Value: nil}},
			sourceOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}, {Name: "use-vc", Value: nil}},
			path:          "/spec/dnsConfig/options",
			patch: []patchOperation{
				{Op: "add", Path: "/spec/dnsConfig/options/-", Value: corev1.PodDNSConfigOption{Name: "ndots", Value: &ndotsVal}},
				{Op: "add", Path: "/spec/dnsConfig/options/-", Value: corev1.PodDNSConfigOption{Name: "use-vc", Value: nil}},
			},
		},
		{
			targetOptions: nil,
			sourceOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}, {Name: "use-vc", Value: nil}},
			path:          "/spec/dnsConfig/options",
			patch: []patchOperation{
				{Op: "add", Path: "/spec/dnsConfig/options", Value: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}}},
				{Op: "add", Path: "/spec/dnsConfig/options/-", Value: corev1.PodDNSConfigOption{Name: "use-vc", Value: nil}},
			},
		},
		{
			targetOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsValOld}},
			sourceOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}, {Name: "use-vc", Value: nil}},
			path:          "/spec/dnsConfig/options",
			patch: []patchOperation{
				{Op: "replace", Path: "/spec/dnsConfig/options/0", Value: corev1.PodDNSConfigOption{Name: "ndots", Value: &ndotsVal}},
				{Op: "add", Path: "/spec/dnsConfig/options/-", Value: corev1.PodDNSConfigOption{Name: "use-vc", Value: nil}},
			},
		},
		{
			targetOptions: []corev1.PodDNSConfigOption{{Name: "single-request-reopen", Value: nil}, {Name: "ndots", Value: &ndotsValOld}},
			sourceOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}, {Name: "use-vc", Value: nil}},
			path:          "/spec/dnsConfig/options",
			patch: []patchOperation{
				{Op: "replace", Path: "/spec/dnsConfig/options/1", Value: corev1.PodDNSConfigOption{Name: "ndots", Value: &ndotsVal}},
				{Op: "add", Path: "/spec/dnsConfig/options/-", Value: corev1.PodDNSConfigOption{Name: "use-vc", Value: nil}},
			},
		},
		{
			targetOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}, {Name: "use-vc", Value: nil}},
			sourceOptions: []corev1.PodDNSConfigOption{{Name: "ndots", Value: &ndotsVal}, {Name: "use-vc", Value: nil}},
			path:          "/spec/dnsConfig/options",
			patch:         []patchOperation{},
		},
	}
	for _, e := range envs {
		patch := addDnsOptions(e.targetOptions, e.sourceOptions, e.path)
		if !cmp.Equal(patch, e.patch) {
			t.Errorf("addDnsOptions was incorrect, for %v, got: %v, want: %v.", e.targetOptions, patch, e.patch)
		}
	}

}

func TestAddRequiredNodeAffinity(t *testing.T) {
	envs := []struct {
		targetTerms []corev1.NodeSelectorTerm
		sourceTerms []corev1.NodeSelectorTerm
		path        string
		patch       []patchOperation
	}{
		{
			targetTerms: []corev1.NodeSelectorTerm{},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchExpressions: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
				Value: []corev1.NodeSelectorTerm{{
					MatchExpressions: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
					}},
				}},
			}},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
				Value: []corev1.NodeSelectorTerm{{
					MatchFields: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
					}},
				}},
			}},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchExpressions: []corev1.NodeSelectorRequirement{{
					Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"uksouth"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchExpressions: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op: "add", Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution/-",
				Value: corev1.NodeSelectorTerm{
					MatchExpressions: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
					}},
				},
			}},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"uksouth"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op: "add", Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution/-",
				Value: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
					}},
				},
			}},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchExpressions: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchExpressions: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			path:  "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			path:  "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchExpressions: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchExpressions: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu20", "ubuntu1804"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution/0",
				Value: corev1.NodeSelectorTerm{
					MatchExpressions: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu20", "ubuntu1804"},
					}},
				},
			}},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu2004"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution/0",
				Value: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu2004"},
					},
					}},
			}},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"A", "B"},
				}},
			}, {
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu2004"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution/1",
				Value: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{{
						Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu2004"},
					}},
				},
			}},
		},
		{
			targetTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"A", "B"},
				}},
			}, {
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "agentpool", Operator: corev1.NodeSelectorOpIn, Values: []string{"ubuntu18", "ubuntu1804"},
				}},
			}},
			sourceTerms: []corev1.NodeSelectorTerm{{
				MatchFields: []corev1.NodeSelectorRequirement{{
					Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"A", "B", "C"},
				}},
			}},
			path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/affinity/nodeAffinity/requiredDuringSchedulingIgnoredDuringExecution/0",
				Value: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{{
						Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"A", "B", "C"},
					}},
				},
			}},
		},
	}
	for _, e := range envs {
		patch := addRequiredNodeAffinityTerms(e.targetTerms, e.sourceTerms, e.path)
		if !cmp.Equal(patch, e.patch) {
			t.Errorf("addRequiredNodeAffinityTerms was incorrect, for %v, got: %v, want: %v.", e.targetTerms, patch, e.patch)
		}
	}
}

func TestAddPreferredNodeAffinity(t *testing.T) {
	envs := []struct {
		targetTerms []corev1.PreferredSchedulingTerm
		sourceTerms []corev1.PreferredSchedulingTerm
		path        string
		patch       []patchOperation
	}{
		{
			targetTerms: []corev1.PreferredSchedulingTerm{},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchExpressions: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
				Value: []corev1.PreferredSchedulingTerm{{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
						},
					},
				}},
			}},
		},
		{
			targetTerms: []corev1.PreferredSchedulingTerm{},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
				Value: []corev1.PreferredSchedulingTerm{{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchFields: []corev1.NodeSelectorRequirement{
							{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
						},
					},
				}},
			}},
		},
		{
			targetTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchExpressions: []corev1.NodeSelectorRequirement{
						{Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"uksouth"}},
					},
				},
			}},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchExpressions: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution/-",
				Value: corev1.PreferredSchedulingTerm{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
						},
					},
				},
			}},
		},
		{
			targetTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"uksouth"}},
					},
				},
			}},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution/-",
				Value: corev1.PreferredSchedulingTerm{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchFields: []corev1.NodeSelectorRequirement{
							{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
						},
					},
				},
			}},
		},
		{
			targetTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchExpressions: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchExpressions: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			path:  "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{},
		},
		{
			targetTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			path:  "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{},
		},
		{
			targetTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"hdd"}},
					},
				},
			}},
			path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution/0",
				Value: corev1.PreferredSchedulingTerm{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchFields: []corev1.NodeSelectorRequirement{
							{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"hdd"}},
						},
					},
				},
			}},
		},
		{
			targetTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "zone", Operator: corev1.NodeSelectorOpIn, Values: []string{"uksouth"}},
					},
				},
			}, {
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"ssd"}},
					},
				},
			}},
			sourceTerms: []corev1.PreferredSchedulingTerm{{
				Weight: 1,
				Preference: corev1.NodeSelectorTerm{
					MatchFields: []corev1.NodeSelectorRequirement{
						{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"hdd"}},
					},
				},
			}},
			path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/affinity/nodeAffinity/preferredDuringSchedulingIgnoredDuringExecution/1",
				Value: corev1.PreferredSchedulingTerm{
					Weight: 1,
					Preference: corev1.NodeSelectorTerm{
						MatchFields: []corev1.NodeSelectorRequirement{
							{Key: "disktype", Operator: corev1.NodeSelectorOpIn, Values: []string{"hdd"}},
						},
					},
				},
			}},
		},
	}

	for _, e := range envs {
		patch := addPreferredNodeAffinityTerms(e.targetTerms, e.sourceTerms, e.path)
		if !cmp.Equal(patch, e.patch) {
			t.Errorf("addPreferredNodeAffinityTerms was incorrect, for %v, got: %v, want: %v.", e.targetTerms, patch, e.patch)
		}
	}
}

func TestAddTolerations(t *testing.T) {
	envs := []struct {
		targetTolerations []corev1.Toleration
		sourceTolerations []corev1.Toleration
		path              string
		patch             []patchOperation
	}{
		{
			targetTolerations: []corev1.Toleration{},
			sourceTolerations: []corev1.Toleration{{
				Key:      "topology.kubernetes.io/region",
				Operator: corev1.TolerationOpExists,
				Effect:   corev1.TaintEffectNoSchedule,
			}},
			path: "/spec/tolerations",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/tolerations",
				Value: []corev1.Toleration{{
					Key:      "topology.kubernetes.io/region",
					Operator: corev1.TolerationOpExists,
					Effect:   corev1.TaintEffectNoSchedule,
				}},
			}},
		},
		{
			targetTolerations: []corev1.Toleration{{
				Key:      "topology.kubernetes.io/region",
				Operator: corev1.TolerationOpEqual,
				Value:    "uksouth",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			sourceTolerations: []corev1.Toleration{{
				Key:      "kubernetes.io/os",
				Operator: corev1.TolerationOpEqual,
				Value:    "Windows",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			path: "/spec/tolerations",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/tolerations/-",
				Value: corev1.Toleration{
					Key:      "kubernetes.io/os",
					Operator: corev1.TolerationOpEqual,
					Value:    "Windows",
					Effect:   corev1.TaintEffectPreferNoSchedule,
				},
			}},
		},
		{
			targetTolerations: []corev1.Toleration{{
				Key:      "kubernetes.io/os",
				Operator: corev1.TolerationOpEqual,
				Value:    "Windows",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			sourceTolerations: []corev1.Toleration{{
				Key:      "kubernetes.io/os",
				Operator: corev1.TolerationOpEqual,
				Value:    "Windows",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			path:  "/spec/tolerations",
			patch: []patchOperation{},
		},
		{
			targetTolerations: []corev1.Toleration{{
				Key:      "kubernetes.io/os",
				Operator: corev1.TolerationOpEqual,
				Value:    "Windows",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			sourceTolerations: []corev1.Toleration{{
				Key:      "kubernetes.io/os",
				Operator: corev1.TolerationOpEqual,
				Value:    "Linux",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			path: "/spec/tolerations",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/tolerations/0",
				Value: corev1.Toleration{
					Key:      "kubernetes.io/os",
					Operator: corev1.TolerationOpEqual,
					Value:    "Linux",
					Effect:   corev1.TaintEffectPreferNoSchedule,
				},
			}},
		},
		{
			targetTolerations: []corev1.Toleration{{
				Key:      "topology.kubernetes.io/region",
				Operator: corev1.TolerationOpExists,
				Effect:   corev1.TaintEffectNoSchedule,
			}, {
				Key:      "kubernetes.io/os",
				Operator: corev1.TolerationOpEqual,
				Value:    "Windows",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			sourceTolerations: []corev1.Toleration{{
				Key:      "kubernetes.io/os",
				Operator: corev1.TolerationOpEqual,
				Value:    "Linux",
				Effect:   corev1.TaintEffectPreferNoSchedule,
			}},
			path: "/spec/tolerations",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/tolerations/1",
				Value: corev1.Toleration{
					Key:      "kubernetes.io/os",
					Operator: corev1.TolerationOpEqual,
					Value:    "Linux",
					Effect:   corev1.TaintEffectPreferNoSchedule,
				},
			}},
		},
	}
	for _, e := range envs {
		patch := addTolerations(e.targetTolerations, e.sourceTolerations, e.path)
		if !cmp.Equal(patch, e.patch) {
			t.Errorf("addTolerations was incorrect, for %v, got: %v, want: %v.", e.targetTolerations, patch, e.patch)
		}
	}
}

func TestAddTopologySpreadConstraints(t *testing.T) {
	topologyHonorPolicy := corev1.NodeInclusionPolicyHonor
	envs := []struct {
		targetTopologySpreadConstraints []corev1.TopologySpreadConstraint
		sourceTopologySpreadConstraints []corev1.TopologySpreadConstraint
		path                            string
		patch                           []patchOperation
	}{
		{
			targetTopologySpreadConstraints: []corev1.TopologySpreadConstraint{},
			sourceTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            1,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			path: "/spec/topologySpreadConstraints",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/topologySpreadConstraints",
				Value: []corev1.TopologySpreadConstraint{{
					MaxSkew:            1,
					TopologyKey:        "topology.kubernetes.io/zone",
					WhenUnsatisfiable:  "ScheduleAnyway",
					LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
					NodeAffinityPolicy: &topologyHonorPolicy,
					NodeTaintsPolicy:   &topologyHonorPolicy,
					MatchLabelKeys:     []string{"pod-template-hash"},
				}},
			}},
		},
		{
			targetTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            1,
				TopologyKey:        "kubernetes.azure.com/agentpool",
				WhenUnsatisfiable:  "DoNotSchedule",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "myFirstTestapp"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			sourceTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            2,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			path: "/spec/topologySpreadConstraints",
			patch: []patchOperation{{
				Op:   "add",
				Path: "/spec/topologySpreadConstraints/-",
				Value: corev1.TopologySpreadConstraint{
					MaxSkew:            2,
					TopologyKey:        "topology.kubernetes.io/zone",
					WhenUnsatisfiable:  "ScheduleAnyway",
					LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
					NodeAffinityPolicy: &topologyHonorPolicy,
					NodeTaintsPolicy:   &topologyHonorPolicy,
					MatchLabelKeys:     []string{"pod-template-hash"},
				},
			}},
		},
		{
			targetTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            1,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			sourceTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            1,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			path:  "/spec/topologySpreadConstraints",
			patch: []patchOperation{},
		},
		{
			targetTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            1,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			sourceTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            2,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			path: "/spec/topologySpreadConstraints",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/topologySpreadConstraints/0",
				Value: corev1.TopologySpreadConstraint{
					MaxSkew:            2,
					TopologyKey:        "topology.kubernetes.io/zone",
					WhenUnsatisfiable:  "ScheduleAnyway",
					LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test"}},
					NodeAffinityPolicy: &topologyHonorPolicy,
					NodeTaintsPolicy:   &topologyHonorPolicy,
					MatchLabelKeys:     []string{"pod-template-hash"},
				},
			}},
		},
		{
			targetTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            1,
				TopologyKey:        "kubernetes.azure.com/agentpool",
				WhenUnsatisfiable:  "DoNotSchedule",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "myFirstTestapp"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}, {
				MaxSkew:            1,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test-app"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template-hash"},
			}},
			sourceTopologySpreadConstraints: []corev1.TopologySpreadConstraint{{
				MaxSkew:            2,
				TopologyKey:        "topology.kubernetes.io/zone",
				WhenUnsatisfiable:  "ScheduleAnyway",
				LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test"}},
				NodeAffinityPolicy: &topologyHonorPolicy,
				NodeTaintsPolicy:   &topologyHonorPolicy,
				MatchLabelKeys:     []string{"pod-template"},
			}},
			path: "/spec/topologySpreadConstraints",
			patch: []patchOperation{{
				Op:   "replace",
				Path: "/spec/topologySpreadConstraints/1",
				Value: corev1.TopologySpreadConstraint{
					MaxSkew:            2,
					TopologyKey:        "topology.kubernetes.io/zone",
					WhenUnsatisfiable:  "ScheduleAnyway",
					LabelSelector:      &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "test"}},
					NodeAffinityPolicy: &topologyHonorPolicy,
					NodeTaintsPolicy:   &topologyHonorPolicy,
					MatchLabelKeys:     []string{"pod-template"},
				},
			}},
		},
	}
	for _, e := range envs {
		patch := addTopologySpreadConstraints(e.targetTopologySpreadConstraints, e.sourceTopologySpreadConstraints, e.path)
		if !cmp.Equal(patch, e.patch) {
			t.Errorf("addTolerations was incorrect, for %v, got: %v, want: %v.", e.targetTopologySpreadConstraints, patch, e.patch)
		}
	}
}

func TestUpdateAnnotations(t *testing.T) {
	annos := []struct {
		targetAnno map[string]string
		sourceAnno map[string]string
		patch      []patchOperation
	}{
		{map[string]string{"some-other-annotation": "some_value"},
			map[string]string{admissionWebhookAnnotationStatusKey: "injected"},
			[]patchOperation{{"add", "/metadata/annotations/" + admissionWebhookAnnotationStatusKey, "injected"}},
		},
		{nil,
			map[string]string{admissionWebhookAnnotationStatusKey: "injected"},
			[]patchOperation{{"add", "/metadata/annotations", map[string]string{admissionWebhookAnnotationStatusKey: "injected"}}},
		},
		{map[string]string{admissionWebhookAnnotationStatusKey: "some_value"},
			map[string]string{admissionWebhookAnnotationStatusKey: "injected"},
			[]patchOperation{{"replace", "/metadata/annotations/" + admissionWebhookAnnotationStatusKey, "injected"}},
		},
	}

	for _, a := range annos {
		patch := updateAnnotation(a.targetAnno, a.sourceAnno)
		if !cmp.Equal(patch, a.patch) {
			t.Errorf("updateAnnotations was incorrect, for %v, got: %v, want: %v.", a.targetAnno, patch, a.patch)
		}
	}
}

// TestRemovePodAntiAffinity tests the removePodAntiAffinity function.
func TestRemovePodAntiAffinity(t *testing.T) {
	basePath := "/spec/affinity/podAntiAffinity"
	expectedPatch := []patchOperation{
		{
			Op:   "remove",
			Path: basePath,
		},
	}

	// Call the removePodAntiAffinity function
	patch := removePodAntiAffinity(basePath)

	// Check if the returned patch matches the expected patch
	if !cmp.Equal(patch, expectedPatch) {
		t.Errorf("removePodAntiAffinity was incorrect, for got: %v, want: %v.", expectedPatch, patch)
	}
}
