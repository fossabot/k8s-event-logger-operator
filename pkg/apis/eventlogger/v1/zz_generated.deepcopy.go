// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventLogger) DeepCopyInto(out *EventLogger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventLogger.
func (in *EventLogger) DeepCopy() *EventLogger {
	if in == nil {
		return nil
	}
	out := new(EventLogger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventLogger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventLoggerConf) DeepCopyInto(out *EventLoggerConf) {
	*out = *in
	if in.Kinds != nil {
		in, out := &in.Kinds, &out.Kinds
		*out = make([]Kind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventLoggerConf.
func (in *EventLoggerConf) DeepCopy() *EventLoggerConf {
	if in == nil {
		return nil
	}
	out := new(EventLoggerConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventLoggerList) DeepCopyInto(out *EventLoggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventLogger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventLoggerList.
func (in *EventLoggerList) DeepCopy() *EventLoggerList {
	if in == nil {
		return nil
	}
	out := new(EventLoggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventLoggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventLoggerSpec) DeepCopyInto(out *EventLoggerSpec) {
	*out = *in
	in.EventLoggerConf.DeepCopyInto(&out.EventLoggerConf)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ScrapeMetrics != nil {
		in, out := &in.ScrapeMetrics, &out.ScrapeMetrics
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventLoggerSpec.
func (in *EventLoggerSpec) DeepCopy() *EventLoggerSpec {
	if in == nil {
		return nil
	}
	out := new(EventLoggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventLoggerStatus) DeepCopyInto(out *EventLoggerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventLoggerStatus.
func (in *EventLoggerStatus) DeepCopy() *EventLoggerStatus {
	if in == nil {
		return nil
	}
	out := new(EventLoggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kind) DeepCopyInto(out *Kind) {
	*out = *in
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchingPatterns != nil {
		in, out := &in.MatchingPatterns, &out.MatchingPatterns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SkipOnMatch != nil {
		in, out := &in.SkipOnMatch, &out.SkipOnMatch
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kind.
func (in *Kind) DeepCopy() *Kind {
	if in == nil {
		return nil
	}
	out := new(Kind)
	in.DeepCopyInto(out)
	return out
}
