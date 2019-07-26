// Package k8sparser is used to parse a k8s json/yaml and allow you to quickly convert them to the proper type
package k8sparser

import v1 "k8s.io/api/rbac/v1"

// ClusterRole return ClusterRole
func (i *Item) ClusterRole() *v1.ClusterRole {
	if item, ok := i.Raw.(*v1.ClusterRole); ok {
		return item
	}

	return nil
}

// ClusterRoleBinding return ClusterRoleBinding
func (i *Item) ClusterRoleBinding() *v1.ClusterRoleBinding {
	if item, ok := i.Raw.(*v1.ClusterRoleBinding); ok {
		return item
	}

	return nil
}

// Role return Role
func (i *Item) Role() *v1.Role {
	if item, ok := i.Raw.(*v1.Role); ok {
		return item
	}

	return nil
}

// RoleBinding return RoleBinding
func (i *Item) RoleBinding() *v1.RoleBinding {
	if item, ok := i.Raw.(*v1.RoleBinding); ok {
		return item
	}

	return nil
}
