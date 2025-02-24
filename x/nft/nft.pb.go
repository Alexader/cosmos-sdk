// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/nft/v1beta1/nft.proto

package nft

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Class defines the class of the nft type.
type Class struct {
	// id defines the unique identifier of the NFT classification, similar to the contract address of ERC721
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name defines the human-readable name of the NFT classification,optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// symbol is an abbreviated name for nft classification,optional
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// description is a brief description of nft classification,optional
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// uri is a URI may point to a JSON file that conforms to the nft classification Metadata JSON Schema.optional
	Uri string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	// uri_hash is a hash of the document pointed to uri,optional
	UriHash string `protobuf:"bytes,6,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb8ebf8e8053172c, []int{0}
}
func (m *Class) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Class.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(m, src)
}
func (m *Class) XXX_Size() int {
	return m.Size()
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Class) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Class) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Class) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Class) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Class) GetUriHash() string {
	if m != nil {
		return m.UriHash
	}
	return ""
}

// NFT defines the NFT.
type NFT struct {
	// class_id defines the unique identifier of the NFT classification, similar to the contract address of ERC721
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// id defines the unique identification of NFT
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// uri defines NFT's metadata storage address outside the chain
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// uri_hash is a hash of the document pointed to uri
	UriHash string `protobuf:"bytes,4,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
	// data is the metadata of the NFT,optional
	Data *types.Any `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *NFT) Reset()         { *m = NFT{} }
func (m *NFT) String() string { return proto.CompactTextString(m) }
func (*NFT) ProtoMessage()    {}
func (*NFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb8ebf8e8053172c, []int{1}
}
func (m *NFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFT.Merge(m, src)
}
func (m *NFT) XXX_Size() int {
	return m.Size()
}
func (m *NFT) XXX_DiscardUnknown() {
	xxx_messageInfo_NFT.DiscardUnknown(m)
}

var xxx_messageInfo_NFT proto.InternalMessageInfo

func (m *NFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *NFT) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NFT) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NFT) GetUriHash() string {
	if m != nil {
		return m.UriHash
	}
	return ""
}

func (m *NFT) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Class)(nil), "cosmos.nft.v1beta1.Class")
	proto.RegisterType((*NFT)(nil), "cosmos.nft.v1beta1.NFT")
}

func init() { proto.RegisterFile("cosmos/nft/v1beta1/nft.proto", fileDescriptor_eb8ebf8e8053172c) }

var fileDescriptor_eb8ebf8e8053172c = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4b, 0xfb, 0x30,
	0x18, 0x5d, 0xda, 0x6e, 0xfb, 0xfd, 0x32, 0x10, 0x09, 0x22, 0x99, 0x48, 0x18, 0x3b, 0xed, 0x62,
	0xc2, 0xf4, 0xea, 0x45, 0x05, 0xd1, 0x8b, 0x87, 0xe1, 0xc9, 0xcb, 0x48, 0x9b, 0x6e, 0x0d, 0xb6,
	0xcd, 0x68, 0x52, 0xb1, 0x7f, 0x81, 0x57, 0xf1, 0xaf, 0xf2, 0xb8, 0xa3, 0x47, 0x69, 0xff, 0x11,
	0x49, 0x5a, 0x87, 0x82, 0xa7, 0xef, 0x7d, 0xef, 0x3d, 0x1e, 0x8f, 0xef, 0x83, 0xc7, 0x91, 0xd2,
	0x99, 0xd2, 0x2c, 0x5f, 0x19, 0xf6, 0x34, 0x0f, 0x63, 0xc3, 0xe7, 0x16, 0xd3, 0x4d, 0xa1, 0x8c,
	0x42, 0xa8, 0x55, 0xa9, 0x65, 0x3a, 0xf5, 0x68, 0xbc, 0x56, 0x6a, 0x9d, 0xc6, 0xcc, 0x39, 0xc2,
	0x72, 0xc5, 0x78, 0x5e, 0xb5, 0xf6, 0xe9, 0x1b, 0x80, 0xfd, 0xab, 0x94, 0x6b, 0x8d, 0xf6, 0xa0,
	0x27, 0x05, 0x06, 0x13, 0x30, 0xfb, 0xbf, 0xf0, 0xa4, 0x40, 0x08, 0x06, 0x39, 0xcf, 0x62, 0xec,
	0x39, 0xc6, 0x61, 0x74, 0x08, 0x07, 0xba, 0xca, 0x42, 0x95, 0x62, 0xdf, 0xb1, 0xdd, 0x86, 0x26,
	0x70, 0x24, 0x62, 0x1d, 0x15, 0x72, 0x63, 0xa4, 0xca, 0x71, 0xe0, 0xc4, 0x9f, 0x14, 0xda, 0x87,
	0x7e, 0x59, 0x48, 0xdc, 0x77, 0x8a, 0x85, 0x68, 0x0c, 0xff, 0x95, 0x85, 0x5c, 0x26, 0x5c, 0x27,
	0x78, 0xe0, 0xe8, 0x61, 0x59, 0xc8, 0x1b, 0xae, 0x93, 0xe9, 0x0b, 0x80, 0xfe, 0xdd, 0xf5, 0xbd,
	0xb5, 0x44, 0xb6, 0xdb, 0x72, 0x57, 0x6c, 0xe8, 0xf6, 0x5b, 0xd1, 0xb5, 0xf5, 0x76, 0x6d, 0xbb,
	0x7c, 0xff, 0xef, 0xfc, 0xe0, 0x57, 0x3e, 0x9a, 0xc1, 0x40, 0x70, 0xc3, 0x31, 0x9c, 0x80, 0xd9,
	0xe8, 0xf4, 0x80, 0xb6, 0xe7, 0xa1, 0xdf, 0xe7, 0xa1, 0x17, 0x79, 0xb5, 0x70, 0x8e, 0xcb, 0xf3,
	0xf7, 0x9a, 0x80, 0x6d, 0x4d, 0xc0, 0x67, 0x4d, 0xc0, 0x6b, 0x43, 0x7a, 0xdb, 0x86, 0xf4, 0x3e,
	0x1a, 0xd2, 0x7b, 0x98, 0xae, 0xa5, 0x49, 0xca, 0x90, 0x46, 0x2a, 0x63, 0xdd, 0x43, 0xda, 0x71,
	0xa2, 0xc5, 0x23, 0x7b, 0xb6, 0x1f, 0x09, 0x07, 0x2e, 0xf1, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x0d, 0x3c, 0x9b, 0xb2, 0x01, 0x00, 0x00,
}

func (m *Class) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Class) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Class) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UriHash) > 0 {
		i -= len(m.UriHash)
		copy(dAtA[i:], m.UriHash)
		i = encodeVarintNft(dAtA, i, uint64(len(m.UriHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNft(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.UriHash) > 0 {
		i -= len(m.UriHash)
		copy(dAtA[i:], m.UriHash)
		i = encodeVarintNft(dAtA, i, uint64(len(m.UriHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintNft(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Class) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.UriHash)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func (m *NFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.UriHash)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func sovNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNft(x uint64) (n int) {
	return sovNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Class) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: Class: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Class: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UriHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: NFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UriHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
				return 0, ErrInvalidLengthNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNft = fmt.Errorf("proto: unexpected end of group")
)
