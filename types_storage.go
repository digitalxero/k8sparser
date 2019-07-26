// Package k8sparser is used to parse a k8s json/yaml and allow you to quickly convert them to the proper type
package k8sparser

import v1 "k8s.io/api/storage/v1"

// StorageClass return StorageClass
func (i *Item) StorageClass() *v1.StorageClass {
	if item, ok := i.Raw.(*v1.StorageClass); ok {
		return item
	}

	return nil
}

// VolumeAttachment return VolumeAttachment
func (i *Item) VolumeAttachment() *v1.VolumeAttachment {
	if item, ok := i.Raw.(*v1.VolumeAttachment); ok {
		return item
	}

	return nil
}
