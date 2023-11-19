package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

func (in *Apple) DeepCopyInto(out *Apple) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

func (in *Apple) DeepCopy() *Apple {
	if in == nil {
		return nil
	}
	out := new(Apple)
	in.DeepCopyInto(out)
	return out
}

func (in *Apple) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AppleList) DeepCopyInto(out *AppleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*Apple, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Apple)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

func (in *AppleList) DeepCopy() *AppleList {
	if in == nil {
		return nil
	}
	out := new(AppleList)
	in.DeepCopyInto(out)
	return out
}

func (in *AppleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AppleSpec) DeepCopyInto(out *AppleSpec) {
	*out = *in
}

func (in *AppleSpec) DeepCopy() *AppleSpec {
	if in == nil {
		return nil
	}
	out := new(AppleSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *Banana) DeepCopyInto(out *Banana) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

func (in *Banana) DeepCopy() *Banana {
	if in == nil {
		return nil
	}
	out := new(Banana)
	in.DeepCopyInto(out)
	return out
}

func (in *Banana) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *BananaList) DeepCopyInto(out *BananaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*Banana, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Banana)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

func (in *BananaList) DeepCopy() *BananaList {
	if in == nil {
		return nil
	}
	out := new(BananaList)
	in.DeepCopyInto(out)
	return out
}

func (in *BananaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *BananaSpec) DeepCopyInto(out *BananaSpec) {
	*out = *in
}

func (in *BananaSpec) DeepCopy() *BananaSpec {
	if in == nil {
		return nil
	}
	out := new(BananaSpec)
	in.DeepCopyInto(out)
	return out
}
