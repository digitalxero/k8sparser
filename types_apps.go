// Package k8sparser is used to parse a k8s json/yaml and allow you to quickly convert them to the proper type
package k8sparser

import v1 "k8s.io/api/apps/v1"

// Deployment return *v1.Deployment
func (i *Item) Deployment() *v1.Deployment {
	if item, ok := i.Raw.(*v1.Deployment); ok {
		return item
	}

	return nil
}

// StatefulSet return *v1.StatefulSet
func (i *Item) StatefulSet() *v1.StatefulSet {
	if item, ok := i.Raw.(*v1.StatefulSet); ok {
		return item
	}

	return nil
}

// DaemonSet return *v1.DaemonSet
func (i *Item) DaemonSet() *v1.DaemonSet {
	if item, ok := i.Raw.(*v1.DaemonSet); ok {
		return item
	}

	return nil
}

// ReplicaSet return ReplicaSet
func (i *Item) ReplicaSet() *v1.ReplicaSet {
	if item, ok := i.Raw.(*v1.ReplicaSet); ok {
		return item
	}

	return nil
}
