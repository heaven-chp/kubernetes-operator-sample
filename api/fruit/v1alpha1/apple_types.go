package v1alpha1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AppleSpec struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type Apple struct {
	metaV1.TypeMeta `json:",inline"`

	metaV1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec AppleSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type AppleList struct {
	metaV1.TypeMeta `json:",inline"`

	metaV1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []*Apple `json:"items" protobuf:"bytes,2,rep,name=items"`
}
