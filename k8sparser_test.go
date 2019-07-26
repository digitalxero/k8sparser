package k8sparser

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	type testDef struct {
		name     string
		args     args
		numItems int
		wantErr  bool
	}
	tests := []testDef{
		{
			name: "single namespace",
			args: args{
				reader: strings.NewReader(`
apiVersion: v1
kind: Namespace
metadata:
  name: test
`),
			},
			numItems: 1,
			wantErr:  false,
		},
		{
			name: "single pod",
			args: args{
				reader: strings.NewReader(`
apiVersion: v1
kind: Pod
metadata:
  name: test
  labels:
    app: test
spec:
  containers:
  - name: test
    image: test
    ports:
    - containerPort: 80
`),
			},
			numItems: 1,
			wantErr:  false,
		},
		{
			name: "single service",
			args: args{
				reader: strings.NewReader(`
apiVersion: v1
kind: Service
metadata:
  name: test
  labels:
    app: test
spec:
  type: NodePort
  ports:
  - port: 80
    protocol: TCP
    name: http
  - port: 443
    protocol: TCP
    name: https
  selector:
    app: test
`),
			},
			numItems: 1,
			wantErr:  false,
		},
		{
			name: "single config map",
			args: args{
				reader: strings.NewReader(`
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
  labels:
    app: test
data:
  test.cfg: |
    ---
    foo: bar
`),
			},
			numItems: 1,
			wantErr:  false,
		},
		{
			name: "single service account",
			args: args{
				reader: strings.NewReader(`
apiVersion: v1
kind: ServiceAccount
metadata:
  name: test
  namespace: test
`),
			},
			numItems: 1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		func(test testDef) {
			t.Run(test.name, func(t *testing.T) {
				gotItems, err := Parse(test.args.reader)
				if (err != nil) != test.wantErr {
					t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)
					return
				}
				if len(gotItems) != test.numItems {
					t.Errorf("Parse() gotItems = %v, want %v", len(gotItems), test.numItems)
					return
				}

				for _, item := range gotItems {
					converted, err := callMethodByName(item, item.Kind)
					if err != nil {
						t.Errorf(err.Error())
						return
					}

					if converted.IsNil() {
						t.Errorf("item.%s() did not return a proper type", item.Kind)
						return
					}
				}
			})
		}(tt)
	}
}

func callMethodByName(i interface{}, methodName string) (reflect.Value, error) {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	// check for method on pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	if finalMethod.IsValid() {
		return finalMethod.Call([]reflect.Value{})[0], nil
	}

	// return or panic, method not found of either type
	return reflect.ValueOf(nil), fmt.Errorf("method %s not found", methodName)
}
