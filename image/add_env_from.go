package main

import (
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
)

// addEnvFrom performs the mutation(s) needed to add the extra environment variables to the target
// resource
func addEnvFrom(target, envFromSources []corev1.EnvFromSource, basePath string) (patch []patchOperation) {
	first := len(target) == 0
	var value interface{}
	for _, envFromSource := range envFromSources {
		value = envFromSource
		path := basePath
		var skip bool
		var op string
		if first {
			first = false
			op = "add"
			value = []corev1.EnvFromSource{envFromSource}
		} else {

			optExists := false
			for idx, targetOpt := range target {
			  if targetOpt.SecretRef != nil && envFromSource.SecretRef != nil {
          nameEqual := cmp.Equal(targetOpt.SecretRef.LocalObjectReference.Name, envFromSource.SecretRef.LocalObjectReference.Name)
          if nameEqual {
            optExists = true
            optionalEqual := cmp.Equal(targetOpt.SecretRef.Optional, envFromSource.SecretRef.Optional)

            skip, op, path = checkReplaceOrSkip(idx, path, optionalEqual)
          }
			  } else if targetOpt.ConfigMapRef != nil && envFromSource.ConfigMapRef != nil {
          nameEqual := cmp.Equal(targetOpt.ConfigMapRef.LocalObjectReference.Name, envFromSource.ConfigMapRef.LocalObjectReference.Name)
          if nameEqual {
            optExists = true
            optionalEqual := cmp.Equal(targetOpt.ConfigMapRef.Optional, envFromSource.ConfigMapRef.Optional)

            skip, op, path = checkReplaceOrSkip(idx, path, optionalEqual)
          }
			  }
			}
			if !optExists {
				op = "add"
				path = path + "/-"
			}
		}
		if !skip {
			patch = append(patch, patchOperation{
				Op:    op,
				Path:  path,
				Value: value,
			})
		} else {
			patch = []patchOperation{}
		}
	}
	return patch
}
