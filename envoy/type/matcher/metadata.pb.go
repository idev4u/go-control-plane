// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/metadata.proto

package matcher

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// MetadataMatcher provides a general interface to check if a given value is matched in
// :ref:`Metadata <envoy_api_msg_core.Metadata>`. It uses `filter` and `path` to retrieve the value
// from the Metadata and then check if it's matched to the specified value.
//
// For example, for the following Metadata:
//
// .. code-block:: yaml
//
//    filter_metadata:
//      envoy.filters.http.rbac:
//        fields:
//          a:
//            struct_value:
//              fields:
//                b:
//                  struct_value:
//                    fields:
//                      c:
//                        string_value: pro
//                t:
//                  list_value:
//                    values:
//                      - string_value: m
//                      - string_value: n
//
// The following MetadataMatcher is matched as the path [a, b, c] will retrieve a string value "pro"
// from the Metadata which is matched to the specified prefix match.
//
// .. code-block:: yaml
//
//    filter: envoy.filters.http.rbac
//    path:
//    - key: a
//    - key: b
//    - key: c
//    value:
//      string_match:
//        prefix: pr
//
// The following MetadataMatcher is matched as the code will match one of the string values in the
// list at the path [a, t].
//
// .. code-block:: yaml
//
//    filter: envoy.filters.http.rbac
//    path:
//    - key: a
//    - key: t
//    value:
//      list_match:
//        one_of:
//          string_match:
//            exact: m
//
// An example use of MetadataMatcher is specifying additional metadata in envoy.filters.http.rbac to
// enforce access control based on dynamic metadata in a request. See :ref:`Permission
// <envoy_api_msg_config.rbac.v2alpha.Permission>` and :ref:`Principal
// <envoy_api_msg_config.rbac.v2alpha.Principal>`.
type MetadataMatcher struct {
	// The filter name to retrieve the Struct from the Metadata.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The path to retrieve the Value from the Struct.
	Path []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	// The MetadataMatcher is matched if the value retrieved by path is matched to this value.
	Value                *ValueMatcher `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MetadataMatcher) Reset()         { *m = MetadataMatcher{} }
func (m *MetadataMatcher) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher) ProtoMessage()    {}
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_68a424c5e967ef23, []int{0}
}
func (m *MetadataMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *MetadataMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher.Merge(dst, src)
}
func (m *MetadataMatcher) XXX_Size() int {
	return m.Size()
}
func (m *MetadataMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher proto.InternalMessageInfo

func (m *MetadataMatcher) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *MetadataMatcher) GetPath() []*MetadataMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *MetadataMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

// Specifies the segment in a path to retrieve value from Metadata.
// Note: Currently it's not supported to retrieve a value from a list in Metadata. This means that
// if the segment key refers to a list, it has to be the last segment in a path.
type MetadataMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*MetadataMatcher_PathSegment_Key
	Segment              isMetadataMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MetadataMatcher_PathSegment) Reset()         { *m = MetadataMatcher_PathSegment{} }
func (m *MetadataMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher_PathSegment) ProtoMessage()    {}
func (*MetadataMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_68a424c5e967ef23, []int{0, 0}
}
func (m *MetadataMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataMatcher_PathSegment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *MetadataMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher_PathSegment.Merge(dst, src)
}
func (m *MetadataMatcher_PathSegment) XXX_Size() int {
	return m.Size()
}
func (m *MetadataMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher_PathSegment proto.InternalMessageInfo

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
	MarshalTo([]byte) (int, error)
	Size() int
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*MetadataMatcher_PathSegment_Key) isMetadataMatcher_PathSegment_Segment() {}

func (m *MetadataMatcher_PathSegment) GetSegment() isMetadataMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MetadataMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*MetadataMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetadataMatcher_PathSegment) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetadataMatcher_PathSegment_OneofMarshaler, _MetadataMatcher_PathSegment_OneofUnmarshaler, _MetadataMatcher_PathSegment_OneofSizer, []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
}

func _MetadataMatcher_PathSegment_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetadataMatcher_PathSegment)
	// segment
	switch x := m.Segment.(type) {
	case *MetadataMatcher_PathSegment_Key:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Key)
	case nil:
	default:
		return fmt.Errorf("MetadataMatcher_PathSegment.Segment has unexpected type %T", x)
	}
	return nil
}

func _MetadataMatcher_PathSegment_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetadataMatcher_PathSegment)
	switch tag {
	case 1: // segment.key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Segment = &MetadataMatcher_PathSegment_Key{x}
		return true, err
	default:
		return false, nil
	}
}

func _MetadataMatcher_PathSegment_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetadataMatcher_PathSegment)
	// segment
	switch x := m.Segment.(type) {
	case *MetadataMatcher_PathSegment_Key:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Key)))
		n += len(x.Key)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.MetadataMatcher.PathSegment")
}
func (m *MetadataMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataMatcher) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Filter) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Filter)))
		i += copy(dAtA[i:], m.Filter)
	}
	if len(m.Path) > 0 {
		for _, msg := range m.Path {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Value != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.Value.Size()))
		n1, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *MetadataMatcher_PathSegment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataMatcher_PathSegment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Segment != nil {
		nn2, err := m.Segment.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *MetadataMatcher_PathSegment_Key) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(len(m.Key)))
	i += copy(dAtA[i:], m.Key)
	return i, nil
}
func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MetadataMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Filter)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if len(m.Path) > 0 {
		for _, e := range m.Path {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataMatcher_PathSegment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Segment != nil {
		n += m.Segment.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataMatcher_PathSegment_Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovMetadata(uint64(l))
	return n
}

func sovMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetadataMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: MetadataMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path, &MetadataMatcher_PathSegment{})
			if err := m.Path[len(m.Path)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &ValueMatcher{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func (m *MetadataMatcher_PathSegment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: PathSegment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathSegment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Segment = &MetadataMatcher_PathSegment_Key{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetadata
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
				next, err := skipMetadata(dAtA[start:])
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
	ErrInvalidLengthMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/type/matcher/metadata.proto", fileDescriptor_metadata_68a424c5e967ef23)
}

var fileDescriptor_metadata_68a424c5e967ef23 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0xcf,
	0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02,
	0x2b, 0xd1, 0x03, 0x29, 0xd1, 0x83, 0x2a, 0x91, 0x92, 0xc3, 0xa2, 0xad, 0x2c, 0x31, 0xa7, 0x34,
	0x15, 0xa2, 0x47, 0x4a, 0xbc, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80,
	0x48, 0x28, 0x75, 0x32, 0x71, 0xf1, 0xfb, 0x42, 0xcd, 0xf7, 0x85, 0x68, 0x14, 0x52, 0xe4, 0x62,
	0x4b, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0xe2, 0xdc, 0xf5,
	0xf2, 0x00, 0x33, 0x4b, 0x11, 0x93, 0x02, 0x63, 0x10, 0x54, 0x42, 0xc8, 0x9f, 0x8b, 0xa5, 0x20,
	0xb1, 0x24, 0x43, 0x82, 0x49, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5f, 0x0f, 0xd3, 0x49, 0x7a, 0x68,
	0xa6, 0xea, 0x05, 0x24, 0x96, 0x64, 0x04, 0xa7, 0xa6, 0xe7, 0xa6, 0xe6, 0x95, 0x38, 0x71, 0x81,
	0x4c, 0x64, 0x9d, 0xc4, 0xc8, 0xc4, 0xc1, 0x18, 0x04, 0x36, 0x48, 0xc8, 0x89, 0x8b, 0x15, 0xec,
	0x5e, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x05, 0x6c, 0x26, 0x86, 0x81, 0x14, 0x40, 0x8d,
	0x83, 0x1a, 0xd1, 0xc5, 0xc8, 0x24, 0xc0, 0x18, 0x04, 0xd1, 0x2a, 0x65, 0xc7, 0xc5, 0x8d, 0x64,
	0x89, 0x90, 0x2c, 0x17, 0x73, 0x76, 0x6a, 0x25, 0x86, 0x1f, 0x3c, 0x18, 0x82, 0x40, 0xe2, 0x4e,
	0x02, 0x5c, 0xec, 0xc5, 0x50, 0x95, 0xac, 0x3b, 0x5e, 0x1e, 0x60, 0x66, 0x74, 0x12, 0x3d, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xa3, 0xd8, 0xa1, 0xd6, 0x26,
	0xb1, 0x81, 0x43, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x57, 0x21, 0x06, 0xb9, 0x9b, 0x01,
	0x00, 0x00,
}
