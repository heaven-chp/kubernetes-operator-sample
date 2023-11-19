package v1alpha1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BananaSpec struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type Banana struct {
	metaV1.TypeMeta `json:",inline"`

	metaV1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec BananaSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type BananaList struct {
	metaV1.TypeMeta `json:",inline"`

	metaV1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []*Banana `json:"items" protobuf:"bytes,2,rep,name=items"`
}
