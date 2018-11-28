// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: asset.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

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

// Asset defines an asset agents install as a dependency for a check.
type Asset struct {
	// URL is the location of the asset
	URL string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Sha512 is the SHA-512 checksum of the asset
	Sha512 string `protobuf:"bytes,3,opt,name=sha512,proto3" json:"sha512,omitempty"`
	// Filters are a collection of sensu queries, used by the system to determine
	// if the asset should be installed. If more than one filter is present the
	// queries are joined by the "AND" operator.
	Filters []string `protobuf:"bytes,5,rep,name=filters" json:"filters"`
	// Metadata contains the name, namespace, labels and annotations of the asset
	ObjectMeta           `protobuf:"bytes,8,opt,name=metadata,embedded=metadata" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_0ba9d73575db6a0b, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(dst, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Asset)(nil), "sensu.core.v2.Asset")
}
func (this *Asset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Asset)
	if !ok {
		that2, ok := that.(Asset)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.URL != that1.URL {
		return false
	}
	if this.Sha512 != that1.Sha512 {
		return false
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if this.Filters[i] != that1.Filters[i] {
			return false
		}
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type AssetFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	GetURL() string
	GetSha512() string
	GetFilters() []string
	GetObjectMeta() ObjectMeta
}

func (this *Asset) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *Asset) TestProto() github_com_golang_protobuf_proto.Message {
	return NewAssetFromFace(this)
}

func (this *Asset) GetURL() string {
	return this.URL
}

func (this *Asset) GetSha512() string {
	return this.Sha512
}

func (this *Asset) GetFilters() []string {
	return this.Filters
}

func (this *Asset) GetObjectMeta() ObjectMeta {
	return this.ObjectMeta
}

func NewAssetFromFace(that AssetFace) *Asset {
	this := &Asset{}
	this.URL = that.GetURL()
	this.Sha512 = that.GetSha512()
	this.Filters = that.GetFilters()
	this.ObjectMeta = that.GetObjectMeta()
	return this
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.URL) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAsset(dAtA, i, uint64(len(m.URL)))
		i += copy(dAtA[i:], m.URL)
	}
	if len(m.Sha512) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Sha512)))
		i += copy(dAtA[i:], m.Sha512)
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	dAtA[i] = 0x42
	i++
	i = encodeVarintAsset(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAsset(r randyAsset, easy bool) *Asset {
	this := &Asset{}
	this.URL = string(randStringAsset(r))
	this.Sha512 = string(randStringAsset(r))
	v1 := r.Intn(10)
	this.Filters = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Filters[i] = string(randStringAsset(r))
	}
	v2 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v2
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedAsset(r, 9)
	}
	return this
}

type randyAsset interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAsset(r randyAsset) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAsset(r randyAsset) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneAsset(r)
	}
	return string(tmps)
}
func randUnrecognizedAsset(r randyAsset, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAsset(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAsset(dAtA []byte, r randyAsset, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAsset(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAsset(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Asset) Size() (n int) {
	var l int
	_ = l
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Sha512)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			l = len(s)
			n += 1 + l + sovAsset(uint64(l))
		}
	}
	l = m.ObjectMeta.Size()
	n += 1 + l + sovAsset(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAsset(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha512", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha512 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAsset
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
func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAsset
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
				next, err := skipAsset(dAtA[start:])
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
	ErrInvalidLengthAsset = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("asset.proto", fileDescriptor_asset_0ba9d73575db6a0b) }

var fileDescriptor_asset_0ba9d73575db6a0b = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0x2e, 0x4e,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0x4e, 0xcd, 0x2b, 0x2e, 0xd5, 0x4b,
	0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x33, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0,
	0x1c, 0x30, 0x0b, 0xa2, 0x5b, 0x8a, 0x2b, 0x37, 0xb5, 0x24, 0x11, 0xc2, 0x56, 0x3a, 0xc4, 0xc8,
	0xc5, 0xea, 0x08, 0x32, 0x59, 0x48, 0x92, 0x8b, 0xb9, 0xb4, 0x28, 0x47, 0x82, 0x49, 0x81, 0x51,
	0x83, 0xd3, 0x89, 0xfd, 0xd1, 0x3d, 0x79, 0xe6, 0xd0, 0x20, 0x9f, 0x20, 0x90, 0x98, 0x90, 0x18,
	0x17, 0x5b, 0x71, 0x46, 0xa2, 0xa9, 0xa1, 0x91, 0x04, 0x33, 0x48, 0x36, 0x08, 0xca, 0x13, 0x52,
	0xe5, 0x62, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0x2a, 0x96, 0x60, 0x55, 0x60, 0xd6, 0xe0, 0x74,
	0xe2, 0x7e, 0x75, 0x4f, 0x1e, 0x26, 0x14, 0x04, 0x63, 0x08, 0x85, 0x72, 0x71, 0x80, 0x6c, 0x4c,
	0x49, 0x2c, 0x49, 0x94, 0xe0, 0x50, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd4, 0x43, 0xf1, 0x80, 0x9e,
	0x7f, 0x52, 0x56, 0x6a, 0x72, 0x89, 0x6f, 0x6a, 0x49, 0xa2, 0x93, 0xdc, 0x89, 0x7b, 0xf2, 0x0c,
	0x17, 0xee, 0xc9, 0x33, 0xbe, 0xba, 0x27, 0x2f, 0x04, 0xd3, 0xa6, 0x93, 0x9f, 0x9b, 0x59, 0x92,
	0x9a, 0x5b, 0x50, 0x52, 0x19, 0x04, 0x37, 0xca, 0x8a, 0xa3, 0x63, 0x81, 0x3c, 0xc3, 0x8a, 0x05,
	0xf2, 0x8c, 0x4e, 0x0a, 0x3f, 0x1e, 0xca, 0x31, 0xae, 0x78, 0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c,
	0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1,
	0x1c, 0x43, 0x14, 0x53, 0x99, 0x51, 0x12, 0x1b, 0xd8, 0xb7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x42, 0xf7, 0xa5, 0x94, 0x46, 0x01, 0x00, 0x00,
}
