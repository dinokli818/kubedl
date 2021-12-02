// +build !ignore_autogenerated

/*
Copyright 2021 The Alibaba Authors.

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
// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Notebook{}, func(obj interface{}) { SetObjectDefaults_Notebook(obj.(*Notebook)) })
	scheme.AddTypeDefaultingFunc(&NotebookList{}, func(obj interface{}) { SetObjectDefaults_NotebookList(obj.(*NotebookList)) })
	return nil
}

func SetObjectDefaults_Notebook(in *Notebook) {
	SetDefaults_Notebook(in)
}

func SetObjectDefaults_NotebookList(in *NotebookList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetDefaults_Notebook(a)
	}
}

func SetDefaults_Notebook(in *Notebook) {
	addNotebookDefaultPort(&in.Spec.Template.Spec)
}

func addNotebookDefaultPort(spec *corev1.PodSpec) {
	index := 0
	for i, container := range spec.Containers {
		if container.Name == NotebookContainerName {
			index = i
			break
		}
	}

	hasNotebookPort := false
	for _, port := range spec.Containers[index].Ports {
		if port.Name == NotebookPortName {
			hasNotebookPort = true
			break
		}
	}
	if !hasNotebookPort {
		spec.Containers[index].Ports = append(spec.Containers[index].Ports, corev1.ContainerPort{
			Name:          NotebookPortName,
			ContainerPort: NotebookDefaultPort,
		})
	}
}
