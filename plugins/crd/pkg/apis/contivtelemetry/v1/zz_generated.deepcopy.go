// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContivAgent) DeepCopyInto(out *ContivAgent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContivAgent.
func (in *ContivAgent) DeepCopy() *ContivAgent {
	if in == nil {
		return nil
	}
	out := new(ContivAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContivNode) DeepCopyInto(out *ContivNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContivNode.
func (in *ContivNode) DeepCopy() *ContivNode {
	if in == nil {
		return nil
	}
	out := new(ContivNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContivTelemetry) DeepCopyInto(out *ContivTelemetry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContivTelemetry.
func (in *ContivTelemetry) DeepCopy() *ContivTelemetry {
	if in == nil {
		return nil
	}
	out := new(ContivTelemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContivTelemetry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContivTelemetryList) DeepCopyInto(out *ContivTelemetryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContivTelemetry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContivTelemetryList.
func (in *ContivTelemetryList) DeepCopy() *ContivTelemetryList {
	if in == nil {
		return nil
	}
	out := new(ContivTelemetryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContivTelemetryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContivTelemetrySpec) DeepCopyInto(out *ContivTelemetrySpec) {
	*out = *in
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = make([]ContivNode, len(*in))
		copy(*out, *in)
	}
	if in.Agent != nil {
		in, out := &in.Agent, &out.Agent
		*out = make([]ContivNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContivTelemetrySpec.
func (in *ContivTelemetrySpec) DeepCopy() *ContivTelemetrySpec {
	if in == nil {
		return nil
	}
	out := new(ContivTelemetrySpec)
	in.DeepCopyInto(out)
	return out
}