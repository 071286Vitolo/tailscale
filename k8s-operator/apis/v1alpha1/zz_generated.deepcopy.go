// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connector) DeepCopyInto(out *Connector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connector.
func (in *Connector) DeepCopy() *Connector {
	if in == nil {
		return nil
	}
	out := new(Connector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Connector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorCondition) DeepCopyInto(out *ConnectorCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorCondition.
func (in *ConnectorCondition) DeepCopy() *ConnectorCondition {
	if in == nil {
		return nil
	}
	out := new(ConnectorCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorList) DeepCopyInto(out *ConnectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Connector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorList.
func (in *ConnectorList) DeepCopy() *ConnectorList {
	if in == nil {
		return nil
	}
	out := new(ConnectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorSpec) DeepCopyInto(out *ConnectorSpec) {
	*out = *in
	if in.SubnetRouter != nil {
		in, out := &in.SubnetRouter, &out.SubnetRouter
		*out = new(SubnetRouter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorSpec.
func (in *ConnectorSpec) DeepCopy() *ConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorStatus) DeepCopyInto(out *ConnectorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConnectorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetRouter != nil {
		in, out := &in.SubnetRouter, &out.SubnetRouter
		*out = new(SubnetRouterStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorStatus.
func (in *ConnectorStatus) DeepCopy() *ConnectorStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetRouter) DeepCopyInto(out *SubnetRouter) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetRouter.
func (in *SubnetRouter) DeepCopy() *SubnetRouter {
	if in == nil {
		return nil
	}
	out := new(SubnetRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetRouterStatus) DeepCopyInto(out *SubnetRouterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetRouterStatus.
func (in *SubnetRouterStatus) DeepCopy() *SubnetRouterStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetRouterStatus)
	in.DeepCopyInto(out)
	return out
}
