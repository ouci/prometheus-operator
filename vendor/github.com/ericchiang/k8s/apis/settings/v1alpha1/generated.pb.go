// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/settings/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/settings/v1alpha1/generated.proto

	It has these top-level messages:
		PodPreset
		PodPresetList
		PodPresetSpec
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import k8s_io_kubernetes_pkg_apis_meta_v1 "github.com/ericchiang/k8s/apis/meta/v1"
import _ "github.com/ericchiang/k8s/runtime"
import _ "github.com/ericchiang/k8s/runtime/schema"
import k8s_io_kubernetes_pkg_api_v1 "github.com/ericchiang/k8s/api/v1"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PodPreset is a policy resource that defines additional runtime
// requirements for a Pod.
type PodPreset struct {
	// +optional
	Metadata *k8s_io_kubernetes_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// +optional
	Spec             *PodPresetSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *PodPreset) Reset()                    { *m = PodPreset{} }
func (m *PodPreset) String() string            { return proto.CompactTextString(m) }
func (*PodPreset) ProtoMessage()               {}
func (*PodPreset) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PodPreset) GetMetadata() *k8s_io_kubernetes_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PodPreset) GetSpec() *PodPresetSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// PodPresetList is a list of PodPreset objects.
type PodPresetList struct {
	// Standard list metadata.
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	// +optional
	Metadata *k8s_io_kubernetes_pkg_apis_meta_v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Items is a list of schema objects.
	Items            []*PodPreset `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PodPresetList) Reset()                    { *m = PodPresetList{} }
func (m *PodPresetList) String() string            { return proto.CompactTextString(m) }
func (*PodPresetList) ProtoMessage()               {}
func (*PodPresetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *PodPresetList) GetMetadata() *k8s_io_kubernetes_pkg_apis_meta_v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PodPresetList) GetItems() []*PodPreset {
	if m != nil {
		return m.Items
	}
	return nil
}

// PodPresetSpec is a description of a pod injection policy.
type PodPresetSpec struct {
	// Selector is a label query over a set of resources, in this case pods.
	// Required.
	Selector *k8s_io_kubernetes_pkg_apis_meta_v1.LabelSelector `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Env defines the collection of EnvVar to inject into containers.
	// +optional
	Env []*k8s_io_kubernetes_pkg_api_v1.EnvVar `protobuf:"bytes,2,rep,name=env" json:"env,omitempty"`
	// EnvFrom defines the collection of EnvFromSource to inject into containers.
	// +optional
	EnvFrom []*k8s_io_kubernetes_pkg_api_v1.EnvFromSource `protobuf:"bytes,3,rep,name=envFrom" json:"envFrom,omitempty"`
	// Volumes defines the collection of Volume to inject into the pod.
	// +optional
	Volumes []*k8s_io_kubernetes_pkg_api_v1.Volume `protobuf:"bytes,4,rep,name=volumes" json:"volumes,omitempty"`
	// VolumeMounts defines the collection of VolumeMount to inject into containers.
	// +optional
	VolumeMounts     []*k8s_io_kubernetes_pkg_api_v1.VolumeMount `protobuf:"bytes,5,rep,name=volumeMounts" json:"volumeMounts,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *PodPresetSpec) Reset()                    { *m = PodPresetSpec{} }
func (m *PodPresetSpec) String() string            { return proto.CompactTextString(m) }
func (*PodPresetSpec) ProtoMessage()               {}
func (*PodPresetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *PodPresetSpec) GetSelector() *k8s_io_kubernetes_pkg_apis_meta_v1.LabelSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *PodPresetSpec) GetEnv() []*k8s_io_kubernetes_pkg_api_v1.EnvVar {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *PodPresetSpec) GetEnvFrom() []*k8s_io_kubernetes_pkg_api_v1.EnvFromSource {
	if m != nil {
		return m.EnvFrom
	}
	return nil
}

func (m *PodPresetSpec) GetVolumes() []*k8s_io_kubernetes_pkg_api_v1.Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *PodPresetSpec) GetVolumeMounts() []*k8s_io_kubernetes_pkg_api_v1.VolumeMount {
	if m != nil {
		return m.VolumeMounts
	}
	return nil
}

func init() {
	proto.RegisterType((*PodPreset)(nil), "github.com/ericchiang.k8s.apis.settings.v1alpha1.PodPreset")
	proto.RegisterType((*PodPresetList)(nil), "github.com/ericchiang.k8s.apis.settings.v1alpha1.PodPresetList")
	proto.RegisterType((*PodPresetSpec)(nil), "github.com/ericchiang.k8s.apis.settings.v1alpha1.PodPresetSpec")
}
func (m *PodPreset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPreset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Spec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
		n2, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodPresetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPresetList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n3, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodPresetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPresetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Selector != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Selector.Size()))
		n4, err := m.Selector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Env) > 0 {
		for _, msg := range m.Env {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.EnvFrom) > 0 {
		for _, msg := range m.EnvFrom {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Volumes) > 0 {
		for _, msg := range m.Volumes {
			dAtA[i] = 0x22
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.VolumeMounts) > 0 {
		for _, msg := range m.VolumeMounts {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PodPreset) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodPresetList) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodPresetSpec) Size() (n int) {
	var l int
	_ = l
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Env) > 0 {
		for _, e := range m.Env {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.EnvFrom) > 0 {
		for _, e := range m.EnvFrom {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Volumes) > 0 {
		for _, e := range m.Volumes {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.VolumeMounts) > 0 {
		for _, e := range m.VolumeMounts {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PodPreset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodPreset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPreset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_apis_meta_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &PodPresetSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodPresetList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodPresetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPresetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_apis_meta_v1.ListMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &PodPreset{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodPresetSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodPresetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPresetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_kubernetes_pkg_apis_meta_v1.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, &k8s_io_kubernetes_pkg_api_v1.EnvVar{})
			if err := m.Env[len(m.Env)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvFrom = append(m.EnvFrom, &k8s_io_kubernetes_pkg_api_v1.EnvFromSource{})
			if err := m.EnvFrom[len(m.EnvFrom)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volumes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Volumes = append(m.Volumes, &k8s_io_kubernetes_pkg_api_v1.Volume{})
			if err := m.Volumes[len(m.Volumes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeMounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolumeMounts = append(m.VolumeMounts, &k8s_io_kubernetes_pkg_api_v1.VolumeMount{})
			if err := m.VolumeMounts[len(m.VolumeMounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/ericchiang/k8s/apis/settings/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0xeb, 0xd3, 0x30,
	0x1c, 0x87, 0xed, 0xf6, 0x1b, 0x9b, 0x99, 0x5e, 0x72, 0x2a, 0x3b, 0x94, 0x31, 0x3c, 0x4c, 0x9c,
	0x29, 0x1d, 0xa2, 0x82, 0xe2, 0x41, 0x98, 0x88, 0x58, 0x36, 0x32, 0xd8, 0xc1, 0x5b, 0xd6, 0x7d,
	0xe9, 0x6a, 0xdb, 0xa4, 0x24, 0x69, 0x5f, 0x8b, 0x2f, 0xc1, 0x97, 0xe2, 0xd1, 0x97, 0x20, 0xf3,
	0xea, 0x8b, 0x90, 0x74, 0x6b, 0xfd, 0xd3, 0x6d, 0xd6, 0xdf, 0xad, 0x94, 0xcf, 0xf3, 0xe4, 0x69,
	0x8a, 0x5e, 0xc6, 0xcf, 0x15, 0x89, 0x84, 0x1b, 0xe7, 0x5b, 0x90, 0x1c, 0x34, 0x28, 0x37, 0x8b,
	0x43, 0x97, 0x65, 0x91, 0x72, 0x15, 0x68, 0x1d, 0xf1, 0x50, 0xb9, 0x85, 0xc7, 0x92, 0x6c, 0xcf,
	0x3c, 0x37, 0x04, 0x0e, 0x92, 0x69, 0xd8, 0x91, 0x4c, 0x0a, 0x2d, 0xf0, 0xec, 0x48, 0x93, 0x5f,
	0x34, 0xc9, 0xe2, 0x90, 0x18, 0x9a, 0x54, 0x34, 0xa9, 0xe8, 0xd1, 0xfc, 0xca, 0x59, 0x29, 0x68,
	0xe6, 0x16, 0x8d, 0x13, 0x46, 0x8f, 0xcf, 0x33, 0x32, 0xe7, 0x3a, 0x4a, 0xa1, 0x31, 0x7f, 0x72,
	0x7d, 0xae, 0x82, 0x3d, 0xa4, 0xac, 0x41, 0xcd, 0x2e, 0x86, 0x9d, 0x49, 0x9a, 0x7c, 0xb6, 0xd0,
	0xdd, 0x95, 0xd8, 0xad, 0x24, 0x28, 0xd0, 0xf8, 0x1d, 0x1a, 0x98, 0xf6, 0x1d, 0xd3, 0xcc, 0xb6,
	0xc6, 0xd6, 0x74, 0x38, 0x27, 0xe4, 0xca, 0xad, 0x98, 0x2d, 0x29, 0x3c, 0xb2, 0xdc, 0x7e, 0x84,
	0x40, 0xfb, 0xa0, 0x19, 0xad, 0x79, 0xbc, 0x44, 0x37, 0x2a, 0x83, 0xc0, 0xee, 0x94, 0x9e, 0x17,
	0xe4, 0x7f, 0x6e, 0x97, 0xd4, 0x49, 0xeb, 0x0c, 0x02, 0x5a, 0x8a, 0x4c, 0xea, 0xfd, 0xfa, 0xfd,
	0xfb, 0x48, 0x69, 0xfc, 0xb6, 0x91, 0x3b, 0x6b, 0x93, 0x6b, 0xd8, 0xbf, 0x62, 0x7d, 0xd4, 0x8b,
	0x34, 0xa4, 0xca, 0xee, 0x8c, 0xbb, 0xd3, 0xe1, 0xfc, 0xd9, 0x2d, 0x6b, 0xe9, 0xd1, 0x32, 0xf9,
	0xd1, 0xf9, 0x2d, 0xd5, 0x7c, 0x02, 0xf6, 0xd1, 0x40, 0x41, 0x02, 0x81, 0x16, 0xf2, 0x94, 0xea,
	0xb5, 0x4a, 0x65, 0x5b, 0x48, 0xd6, 0x27, 0x90, 0xd6, 0x0a, 0xfc, 0x14, 0x75, 0x81, 0x17, 0xa7,
	0xda, 0x07, 0x97, 0x4d, 0xc6, 0xb1, 0xe0, 0xc5, 0x86, 0x49, 0x6a, 0x00, 0xbc, 0x40, 0x7d, 0xe0,
	0xc5, 0x1b, 0x29, 0x52, 0xbb, 0x5b, 0xb2, 0x8f, 0xfe, 0xc9, 0x9a, 0xf1, 0x5a, 0xe4, 0x32, 0x00,
	0x5a, 0xb1, 0xf8, 0x15, 0xea, 0x17, 0x22, 0xc9, 0x53, 0x50, 0xf6, 0x4d, 0x9b, 0x84, 0x4d, 0x39,
	0xa6, 0x15, 0x84, 0x7d, 0x74, 0xef, 0xf8, 0xe8, 0x8b, 0x9c, 0x6b, 0x65, 0xf7, 0x4a, 0xc9, 0xc3,
	0x36, 0x92, 0x92, 0xa0, 0x7f, 0xe0, 0xaf, 0x47, 0x5f, 0x0e, 0x8e, 0xf5, 0xf5, 0xe0, 0x58, 0xdf,
	0x0e, 0x8e, 0xf5, 0xe9, 0xbb, 0x73, 0xe7, 0xc3, 0xa0, 0xfa, 0x37, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x3e, 0x05, 0x30, 0x95, 0x14, 0x04, 0x00, 0x00,
}