// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

package mesosproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// For use with detector callbacks
type InternalMasterChangeDetected struct {
	// will be present if there's a new master, otherwise nil
	Master           *MasterInfo `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *InternalMasterChangeDetected) Reset()         { *m = InternalMasterChangeDetected{} }
func (m *InternalMasterChangeDetected) String() string { return proto.CompactTextString(m) }
func (*InternalMasterChangeDetected) ProtoMessage()    {}

func (m *InternalMasterChangeDetected) GetMaster() *MasterInfo {
	if m != nil {
		return m.Master
	}
	return nil
}

type InternalTryAuthentication struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *InternalTryAuthentication) Reset()         { *m = InternalTryAuthentication{} }
func (m *InternalTryAuthentication) String() string { return proto.CompactTextString(m) }
func (*InternalTryAuthentication) ProtoMessage()    {}

type InternalAuthenticationResult struct {
	// true only if the authentication process completed and login was successful
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// true if the authentication process completed, successfully or not
	Completed *bool `protobuf:"varint,2,req,name=completed" json:"completed,omitempty"`
	// master pid that this result pertains to
	Pid              *string `protobuf:"bytes,3,req,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InternalAuthenticationResult) Reset()         { *m = InternalAuthenticationResult{} }
func (m *InternalAuthenticationResult) String() string { return proto.CompactTextString(m) }
func (*InternalAuthenticationResult) ProtoMessage()    {}

func (m *InternalAuthenticationResult) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *InternalAuthenticationResult) GetCompleted() bool {
	if m != nil && m.Completed != nil {
		return *m.Completed
	}
	return false
}

func (m *InternalAuthenticationResult) GetPid() string {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return ""
}

func init() {
	proto.RegisterType((*InternalMasterChangeDetected)(nil), "mesosproto.InternalMasterChangeDetected")
	proto.RegisterType((*InternalTryAuthentication)(nil), "mesosproto.InternalTryAuthentication")
	proto.RegisterType((*InternalAuthenticationResult)(nil), "mesosproto.InternalAuthenticationResult")
}
