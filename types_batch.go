// Package k8sparser is used to parse a k8s json/yaml and allow you to quickly convert them to the proper type
package k8sparser

import (
	v1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
)

// Job return Job
func (i *Item) Job() *v1.Job {
	if item, ok := i.Raw.(*v1.Job); ok {
		return item
	}

	return nil
}

// CronJob return CronJob
func (i *Item) CronJob() *v1beta1.CronJob {
	if item, ok := i.Raw.(*v1beta1.CronJob); ok {
		return item
	}

	return nil
}
