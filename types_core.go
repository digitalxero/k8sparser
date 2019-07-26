// Package k8sparser is used to parse a k8s json/yaml and allow you to quickly convert them to the proper type
package k8sparser

import v1 "k8s.io/api/core/v1"

// Namespace return Namespace
func (i *Item) Namespace() *v1.Namespace {
	if item, ok := i.Raw.(*v1.Namespace); ok {
		return item
	}

	return nil
}

// Pod return Pod
func (i *Item) Pod() *v1.Pod {
	if item, ok := i.Raw.(*v1.Pod); ok {
		return item
	}

	return nil
}

// Service return Service
func (i *Item) Service() *v1.Service {
	if item, ok := i.Raw.(*v1.Service); ok {
		return item
	}

	return nil
}

// ConfigMap return ConfigMap
func (i *Item) ConfigMap() *v1.ConfigMap {
	if item, ok := i.Raw.(*v1.ConfigMap); ok {
		return item
	}

	return nil
}

// ServiceAccount return ServiceAccount
func (i *Item) ServiceAccount() *v1.ServiceAccount {
	if item, ok := i.Raw.(*v1.ServiceAccount); ok {
		return item
	}

	return nil
}
