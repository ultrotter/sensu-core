// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/core/v2/apikey.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An APIKey is an api key specification.
type APIKey struct {
	// Metadata contains the name, namespace (N/A), labels and annotations of
	// the APIKey.
	ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3,embedded=metadata" json:"metadata,omitempty"`
	// Username is the username associated with the API key.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// CreatedAt is a timestamp which the API key was created.
	CreatedAt int64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Hash is the hashed value of the API key.
	Hash                 []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIKey) Reset()         { *m = APIKey{} }
func (m *APIKey) String() string { return proto.CompactTextString(m) }
func (*APIKey) ProtoMessage()    {}
func (*APIKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e030266eca60c06, []int{0}
}
func (m *APIKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *APIKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_APIKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *APIKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKey.Merge(m, src)
}
func (m *APIKey) XXX_Size() int {
	return m.Size()
}
func (m *APIKey) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKey.DiscardUnknown(m)
}

var xxx_messageInfo_APIKey proto.InternalMessageInfo

func init() {
	proto.RegisterType((*APIKey)(nil), "sensu.core.v2.APIKey")
}

func init() {
	proto.RegisterFile("github.com/sensu/core/v2/apikey.proto", fileDescriptor_2e030266eca60c06)
}

var fileDescriptor_2e030266eca60c06 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x4e, 0xcd, 0x2b, 0x2e, 0xd5, 0x4f, 0xce, 0x2f,
	0x4a, 0xd5, 0x2f, 0x33, 0xd2, 0x4f, 0x2c, 0xc8, 0xcc, 0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x05, 0xcb, 0xe9, 0x81, 0xe4, 0xf4, 0xca, 0x8c, 0xa4, 0x4c, 0x90, 0x74, 0xa5,
	0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x55, 0x25, 0x95, 0xa6, 0x39, 0x94, 0x19, 0xea, 0x19, 0xeb, 0x19,
	0x81, 0x05, 0xc1, 0x62, 0x60, 0x16, 0xc4, 0x10, 0x29, 0x65, 0x9c, 0x76, 0xe5, 0xa6, 0x96, 0x24,
	0x42, 0x14, 0x29, 0xed, 0x60, 0xe4, 0x62, 0x73, 0x0c, 0xf0, 0xf4, 0x4e, 0xad, 0x14, 0x0a, 0xe5,
	0xe2, 0x00, 0x49, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xea,
	0xa1, 0xb8, 0x43, 0xcf, 0x3f, 0x29, 0x2b, 0x35, 0xb9, 0xc4, 0x37, 0xb5, 0x24, 0xd1, 0x49, 0xee,
	0xc4, 0x3d, 0x79, 0x86, 0x0b, 0xf7, 0xe4, 0x19, 0x5f, 0xdd, 0x93, 0x17, 0x82, 0x69, 0xd3, 0xc9,
	0xcf, 0xcd, 0x2c, 0x49, 0xcd, 0x2d, 0x28, 0xa9, 0x0c, 0x82, 0x1b, 0x25, 0x24, 0xc5, 0xc5, 0x51,
	0x5a, 0x9c, 0x5a, 0x94, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7,
	0x0b, 0xc9, 0x72, 0x71, 0x25, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6, 0xc4, 0x27, 0x96, 0x48, 0x30,
	0x2b, 0x30, 0x6a, 0x30, 0x07, 0x71, 0x42, 0x45, 0x1c, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x32, 0x12,
	0x8b, 0x33, 0x24, 0x58, 0x14, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x2b, 0x8e, 0x8e, 0x05, 0xf2,
	0x0c, 0x2b, 0x16, 0xc8, 0x33, 0x3a, 0x29, 0xfc, 0x78, 0x28, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3,
	0x8e, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0x4c, 0x65, 0x46, 0x49, 0x6c, 0x60, 0x3f, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xfd, 0x57, 0x4d, 0x95, 0x76, 0x01, 0x00, 0x00,
}

func (this *APIKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*APIKey)
	if !ok {
		that2, ok := that.(APIKey)
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
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if this.Username != that1.Username {
		return false
	}
	if this.CreatedAt != that1.CreatedAt {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type APIKeyFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	GetObjectMeta() ObjectMeta
	GetUsername() string
	GetCreatedAt() int64
	GetHash() []byte
}

func (this *APIKey) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *APIKey) TestProto() github_com_golang_protobuf_proto.Message {
	return NewAPIKeyFromFace(this)
}

func (this *APIKey) GetObjectMeta() ObjectMeta {
	return this.ObjectMeta
}

func (this *APIKey) GetUsername() string {
	return this.Username
}

func (this *APIKey) GetCreatedAt() int64 {
	return this.CreatedAt
}

func (this *APIKey) GetHash() []byte {
	return this.Hash
}

func NewAPIKeyFromFace(that APIKeyFace) *APIKey {
	this := &APIKey{}
	this.ObjectMeta = that.GetObjectMeta()
	this.Username = that.GetUsername()
	this.CreatedAt = that.GetCreatedAt()
	this.Hash = that.GetHash()
	return this
}

func (m *APIKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *APIKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *APIKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintApikey(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if m.CreatedAt != 0 {
		i = encodeVarintApikey(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintApikey(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintApikey(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintApikey(dAtA []byte, offset int, v uint64) int {
	offset -= sovApikey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedAPIKey(r randyApikey, easy bool) *APIKey {
	this := &APIKey{}
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	this.Username = string(randStringApikey(r))
	this.CreatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.CreatedAt *= -1
	}
	v2 := r.Intn(100)
	this.Hash = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Hash[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedApikey(r, 5)
	}
	return this
}

type randyApikey interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneApikey(r randyApikey) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringApikey(r randyApikey) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneApikey(r)
	}
	return string(tmps)
}
func randUnrecognizedApikey(r randyApikey, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldApikey(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldApikey(dAtA []byte, r randyApikey, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateApikey(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateApikey(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateApikey(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateApikey(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateApikey(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateApikey(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateApikey(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *APIKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovApikey(uint64(l))
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovApikey(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovApikey(uint64(m.CreatedAt))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovApikey(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApikey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApikey(x uint64) (n int) {
	return sovApikey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *APIKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApikey
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApikey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApikey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApikey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApikey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApikey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApikey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApikey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApikey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApikey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApikey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApikey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApikey
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
func skipApikey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApikey
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
					return 0, ErrIntOverflowApikey
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApikey
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
			if length < 0 {
				return 0, ErrInvalidLengthApikey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApikey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApikey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApikey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApikey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApikey = fmt.Errorf("proto: unexpected end of group")
)
