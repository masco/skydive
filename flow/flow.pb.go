// Code generated by protoc-gen-go.
// source: flow/flow.proto
// DO NOT EDIT!

/*
Package flow is a generated protocol buffer package.

It is generated from these files:
	flow/flow.proto

It has these top-level messages:
	Flow
*/
package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Flow struct {
	UUID             *string `protobuf:"bytes,1,req,name=UUID" json:"UUID,omitempty"`
	LayersPath       *string `protobuf:"bytes,2,req,name=LayersPath" json:"LayersPath,omitempty"`
	EtherSrc         *string `protobuf:"bytes,3,req,name=EtherSrc" json:"EtherSrc,omitempty"`
	EtherDst         *string `protobuf:"bytes,4,req,name=EtherDst" json:"EtherDst,omitempty"`
	Ipv4Src          *string `protobuf:"bytes,5,opt,name=Ipv4Src" json:"Ipv4Src,omitempty"`
	Ipv4Dst          *string `protobuf:"bytes,6,opt,name=Ipv4Dst" json:"Ipv4Dst,omitempty"`
	PortSrc          *uint32 `protobuf:"varint,7,opt,name=PortSrc" json:"PortSrc,omitempty"`
	PortDst          *uint32 `protobuf:"varint,8,opt,name=PortDst" json:"PortDst,omitempty"`
	ID               *uint64 `protobuf:"varint,9,opt,name=ID" json:"ID,omitempty"`
	Timestamp        *uint64 `protobuf:"varint,10,req,name=Timestamp" json:"Timestamp,omitempty"`
	ProbeMAC         *string `protobuf:"bytes,11,req,name=ProbeMAC" json:"ProbeMAC,omitempty"`
	ProbeGraphPath   *string `protobuf:"bytes,12,opt,name=ProbeGraphPath" json:"ProbeGraphPath,omitempty"`
	IfSrcName        *string `protobuf:"bytes,13,opt,name=IfSrcName" json:"IfSrcName,omitempty"`
	IfSrcType        *string `protobuf:"bytes,14,opt,name=IfSrcType" json:"IfSrcType,omitempty"`
	IfSrcGraphPath   *string `protobuf:"bytes,15,opt,name=IfSrcGraphPath" json:"IfSrcGraphPath,omitempty"`
	IfSrcTenantID    *string `protobuf:"bytes,16,opt,name=IfSrcTenantID" json:"IfSrcTenantID,omitempty"`
	IfSrcVNI         *uint64 `protobuf:"varint,17,opt,name=IfSrcVNI" json:"IfSrcVNI,omitempty"`
	IfDstName        *string `protobuf:"bytes,18,opt,name=IfDstName" json:"IfDstName,omitempty"`
	IfDstType        *string `protobuf:"bytes,19,opt,name=IfDstType" json:"IfDstType,omitempty"`
	IfDstGraphPath   *string `protobuf:"bytes,20,opt,name=IfDstGraphPath" json:"IfDstGraphPath,omitempty"`
	IfDstTenantID    *string `protobuf:"bytes,21,opt,name=IfDstTenantID" json:"IfDstTenantID,omitempty"`
	IfDstVNI         *uint64 `protobuf:"varint,22,opt,name=IfDstVNI" json:"IfDstVNI,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Flow) Reset()         { *m = Flow{} }
func (m *Flow) String() string { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()    {}

func (m *Flow) GetUUID() string {
	if m != nil && m.UUID != nil {
		return *m.UUID
	}
	return ""
}

func (m *Flow) GetLayersPath() string {
	if m != nil && m.LayersPath != nil {
		return *m.LayersPath
	}
	return ""
}

func (m *Flow) GetEtherSrc() string {
	if m != nil && m.EtherSrc != nil {
		return *m.EtherSrc
	}
	return ""
}

func (m *Flow) GetEtherDst() string {
	if m != nil && m.EtherDst != nil {
		return *m.EtherDst
	}
	return ""
}

func (m *Flow) GetIpv4Src() string {
	if m != nil && m.Ipv4Src != nil {
		return *m.Ipv4Src
	}
	return ""
}

func (m *Flow) GetIpv4Dst() string {
	if m != nil && m.Ipv4Dst != nil {
		return *m.Ipv4Dst
	}
	return ""
}

func (m *Flow) GetPortSrc() uint32 {
	if m != nil && m.PortSrc != nil {
		return *m.PortSrc
	}
	return 0
}

func (m *Flow) GetPortDst() uint32 {
	if m != nil && m.PortDst != nil {
		return *m.PortDst
	}
	return 0
}

func (m *Flow) GetID() uint64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Flow) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Flow) GetProbeMAC() string {
	if m != nil && m.ProbeMAC != nil {
		return *m.ProbeMAC
	}
	return ""
}

func (m *Flow) GetProbeGraphPath() string {
	if m != nil && m.ProbeGraphPath != nil {
		return *m.ProbeGraphPath
	}
	return ""
}

func (m *Flow) GetIfSrcName() string {
	if m != nil && m.IfSrcName != nil {
		return *m.IfSrcName
	}
	return ""
}

func (m *Flow) GetIfSrcType() string {
	if m != nil && m.IfSrcType != nil {
		return *m.IfSrcType
	}
	return ""
}

func (m *Flow) GetIfSrcGraphPath() string {
	if m != nil && m.IfSrcGraphPath != nil {
		return *m.IfSrcGraphPath
	}
	return ""
}

func (m *Flow) GetIfSrcTenantID() string {
	if m != nil && m.IfSrcTenantID != nil {
		return *m.IfSrcTenantID
	}
	return ""
}

func (m *Flow) GetIfSrcVNI() uint64 {
	if m != nil && m.IfSrcVNI != nil {
		return *m.IfSrcVNI
	}
	return 0
}

func (m *Flow) GetIfDstName() string {
	if m != nil && m.IfDstName != nil {
		return *m.IfDstName
	}
	return ""
}

func (m *Flow) GetIfDstType() string {
	if m != nil && m.IfDstType != nil {
		return *m.IfDstType
	}
	return ""
}

func (m *Flow) GetIfDstGraphPath() string {
	if m != nil && m.IfDstGraphPath != nil {
		return *m.IfDstGraphPath
	}
	return ""
}

func (m *Flow) GetIfDstTenantID() string {
	if m != nil && m.IfDstTenantID != nil {
		return *m.IfDstTenantID
	}
	return ""
}

func (m *Flow) GetIfDstVNI() uint64 {
	if m != nil && m.IfDstVNI != nil {
		return *m.IfDstVNI
	}
	return 0
}
