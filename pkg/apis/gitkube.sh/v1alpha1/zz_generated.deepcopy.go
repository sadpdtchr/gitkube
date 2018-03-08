// +build !ignore_autogenerated

/*
Copyright 2018 Hasura.io

*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Remote).DeepCopyInto(out.(*Remote))
			return nil
		}, InType: reflect.TypeOf(&Remote{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RemoteList).DeepCopyInto(out.(*RemoteList))
			return nil
		}, InType: reflect.TypeOf(&RemoteList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RemoteSpec).DeepCopyInto(out.(*RemoteSpec))
			return nil
		}, InType: reflect.TypeOf(&RemoteSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RemoteStatus).DeepCopyInto(out.(*RemoteStatus))
			return nil
		}, InType: reflect.TypeOf(&RemoteStatus{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Remote) DeepCopyInto(out *Remote) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Remote.
func (in *Remote) DeepCopy() *Remote {
	if in == nil {
		return nil
	}
	out := new(Remote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Remote) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteList) DeepCopyInto(out *RemoteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Remote, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteList.
func (in *RemoteList) DeepCopy() *RemoteList {
	if in == nil {
		return nil
	}
	out := new(RemoteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteSpec) DeepCopyInto(out *RemoteSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteSpec.
func (in *RemoteSpec) DeepCopy() *RemoteSpec {
	if in == nil {
		return nil
	}
	out := new(RemoteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteStatus) DeepCopyInto(out *RemoteStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteStatus.
func (in *RemoteStatus) DeepCopy() *RemoteStatus {
	if in == nil {
		return nil
	}
	out := new(RemoteStatus)
	in.DeepCopyInto(out)
	return out
}