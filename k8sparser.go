// Package k8sparser is used to parse a k8s json/yaml and allow you to quickly convert them to the proper type
package k8sparser

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	k8sYaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

// Parse return the parsed k8s items
func Parse(reader io.Reader) (items Items, err error) {
	var (
		doc         []byte
		kindDecoder runtime.Decoder
		docReader   *k8sYaml.YAMLReader
		obj         runtime.Object
	)
	kindDecoder = scheme.Codecs.UniversalDeserializer()
	docReader = k8sYaml.NewYAMLReader(bufio.NewReader(reader))

	for {
		doc, err = docReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			err = errors.Wrap(err, "error reading documents")
			return nil, err
		}

		if obj, _, err = kindDecoder.Decode(doc, nil, nil); err != nil {
			err = errors.Wrap(err, "error decoding documents")
			return nil, err
		}

		doc, err = json.Marshal(obj)
		if err != nil {
			err = errors.Wrap(err, "error decoding document")
			return nil, err
		}
		var item *Item
		err = json.Unmarshal(doc, &item)
		if err != nil {
			err = errors.Wrap(err, "error reading doc metadata")
			return nil, err
		}
		item.Raw = obj
		items = append(items, item)
	}

	return items, nil
}
