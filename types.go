// Package k8sparser is used to parse a k8s json/yaml and allow you to quickly convert them to the proper type
package k8sparser

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Items a list of *k8sparser.Item
type Items = []*Item

// Item object that has enough of an objects type & metadata
// that you should be able decide which method to call to the
// real `k8s.io/api/*` type
type Item struct {
	metav1.TypeMeta
	MetaData metav1.ObjectMeta
	Raw      runtime.Object
}
