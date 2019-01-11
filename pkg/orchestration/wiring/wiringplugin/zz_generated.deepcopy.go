// +build !ignore_autogenerated

// Generated code
// run `make generate` to update

// Code generated by deepcopy-gen. DO NOT EDIT.

package wiringplugin

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindableShapeStruct) DeepCopyInto(out *BindableShapeStruct) {
	*out = *in
	in.ServiceInstanceName.DeepCopyInto(&out.ServiceInstanceName)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindableShapeStruct.
func (in *BindableShapeStruct) DeepCopy() *BindableShapeStruct {
	if in == nil {
		return nil
	}
	out := new(BindableShapeStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingProtoReference.
func (in *BindingProtoReference) DeepCopy() *BindingProtoReference {
	if in == nil {
		return nil
	}
	out := new(BindingProtoReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingSecretProtoReference.
func (in *BindingSecretProtoReference) DeepCopy() *BindingSecretProtoReference {
	if in == nil {
		return nil
	}
	out := new(BindingSecretProtoReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtoReference.
func (in *ProtoReference) DeepCopy() *ProtoReference {
	if in == nil {
		return nil
	}
	out := new(ProtoReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnstructuredShape.
func (in *UnstructuredShape) DeepCopy() *UnstructuredShape {
	if in == nil {
		return nil
	}
	out := new(UnstructuredShape)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyShape is an autogenerated deepcopy function, copying the receiver, creating a new Shape.
func (in *UnstructuredShape) DeepCopyShape() Shape {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
