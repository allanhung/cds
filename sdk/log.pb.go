// Code generated by protoc-gen-go.
// source: log.proto
// DO NOT EDIT!

/*
Package sdk is a generated protocol buffer package.

It is generated from these files:
	log.proto

It has these top-level messages:
	Log
	ServiceLog
*/
package sdk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Log represents an execution log
// Generate *.pb.go files with:
// 	protoc --go_out=plugins=grpc:. ./log.proto
// 	protoc-go-inject-tag -input=./log.pb.go
// 	=> github.com/favadi/protoc-go-inject-tag
type Log struct {
	// @inject_tag: db:"id"
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"pipeline_build_job_id"
	PipelineBuildJobID int64 `protobuf:"varint,2,opt,name=pipelineBuildJobID" json:"pipelineBuildJobID,omitempty" db:"pipeline_build_job_id"`
	// @inject_tag: db:"pipeline_build_id"
	PipelineBuildID int64 `protobuf:"varint,3,opt,name=pipelineBuildID" json:"pipelineBuildID,omitempty" db:"pipeline_build_id"`
	// @inject_tag: db:"start"
	Start *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=start" json:"start,omitempty" db:"start"`
	// @inject_tag: db:"last_modified"
	LastModified *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=lastModified" json:"lastModified,omitempty" db:"last_modified"`
	// @inject_tag: db:"done"
	Done *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=done" json:"done,omitempty" db:"done"`
	// @inject_tag: db:"step_order"
	StepOrder int64 `protobuf:"varint,7,opt,name=stepOrder" json:"stepOrder,omitempty" db:"step_order"`
	// @inject_tag: db:"value"
	Val string `protobuf:"bytes,8,opt,name=val" json:"val,omitempty" db:"value"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Log) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Log) GetPipelineBuildJobID() int64 {
	if m != nil {
		return m.PipelineBuildJobID
	}
	return 0
}

func (m *Log) GetPipelineBuildID() int64 {
	if m != nil {
		return m.PipelineBuildID
	}
	return 0
}

func (m *Log) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Log) GetLastModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

func (m *Log) GetDone() *google_protobuf.Timestamp {
	if m != nil {
		return m.Done
	}
	return nil
}

func (m *Log) GetStepOrder() int64 {
	if m != nil {
		return m.StepOrder
	}
	return 0
}

func (m *Log) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

// ServiceLog represents an execution log
// Generate *.pb.go files with:
// 	protoc --go_out=plugins=grpc:. ./log.proto
// 	protoc-go-inject-tag -input=./log.pb.go
// 	=> github.com/favadi/protoc-go-inject-tag
type ServiceLog struct {
	// @inject_tag: db:"id"
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"workflow_node_run_job_id" json:"workflow_node_run_job_id"
	WorkflowNodeJobRunID int64 `protobuf:"varint,2,opt,name=workflowNodeJobRunID" json:"workflow_node_run_job_id" db:"workflow_node_run_job_id"`
	// @inject_tag: db:"workflow_node_run_id" json:"workflow_node_run_id"
	WorkflowNodeRunID int64 `protobuf:"varint,3,opt,name=workflowNodeRunID" json:"workflow_node_run_id" db:"workflow_node_run_id"`
	// @inject_tag: db:"start",json:"start"
	Start *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=start" json:"start" db:"start"`
	// @inject_tag: db:"last_modified" json:"last_modified"
	LastModified *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=lastModified" json:"last_modified" db:"last_modified"`
	// @inject_tag: db:"-" json:"requirement_id"
	ServiceRequirementID int64 `protobuf:"varint,6,opt,name=serviceRequirementID" json:"requirement_id" db:"-"`
	// @inject_tag: db:"requirement_service_name" json:"requirement_service_name"
	ServiceRequirementName string `protobuf:"bytes,7,opt,name=serviceRequirementName" json:"requirement_service_name" db:"requirement_service_name"`
	// @inject_tag: db:"value"
	Val string `protobuf:"bytes,8,opt,name=val" json:"val,omitempty" db:"value"`
}

func (m *ServiceLog) Reset()                    { *m = ServiceLog{} }
func (m *ServiceLog) String() string            { return proto.CompactTextString(m) }
func (*ServiceLog) ProtoMessage()               {}
func (*ServiceLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceLog) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ServiceLog) GetWorkflowNodeJobRunID() int64 {
	if m != nil {
		return m.WorkflowNodeJobRunID
	}
	return 0
}

func (m *ServiceLog) GetWorkflowNodeRunID() int64 {
	if m != nil {
		return m.WorkflowNodeRunID
	}
	return 0
}

func (m *ServiceLog) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ServiceLog) GetLastModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

func (m *ServiceLog) GetServiceRequirementID() int64 {
	if m != nil {
		return m.ServiceRequirementID
	}
	return 0
}

func (m *ServiceLog) GetServiceRequirementName() string {
	if m != nil {
		return m.ServiceRequirementName
	}
	return ""
}

func (m *ServiceLog) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func init() {
	proto.RegisterType((*Log)(nil), "github.com.ovh.cds.sdk.Log")
	proto.RegisterType((*ServiceLog)(nil), "github.com.ovh.cds.sdk.ServiceLog")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x90, 0xcf, 0x6b, 0xe2, 0x40,
	0x14, 0xc7, 0x49, 0xa2, 0xee, 0xe6, 0xed, 0xb2, 0x3f, 0x06, 0x91, 0x41, 0x16, 0x56, 0x3c, 0xe5,
	0xb0, 0x8c, 0x8b, 0x85, 0x1e, 0x7b, 0x10, 0x2f, 0x4a, 0x6b, 0x21, 0xed, 0xa9, 0xb7, 0xc4, 0x79,
	0xc6, 0xc1, 0x24, 0x2f, 0xcd, 0x4c, 0xf4, 0xef, 0xeb, 0x3f, 0xd4, 0xbf, 0xa1, 0x74, 0x82, 0x6d,
	0xad, 0x29, 0x1e, 0x7b, 0x7b, 0x7c, 0x7f, 0xc0, 0xf7, 0x7d, 0xc0, 0x4f, 0x29, 0x11, 0x45, 0x49,
	0x86, 0x58, 0x2f, 0x51, 0x66, 0x5d, 0xc5, 0x62, 0x49, 0x99, 0xa0, 0xed, 0x5a, 0x2c, 0xa5, 0x16,
	0x5a, 0x6e, 0xfa, 0x7f, 0x13, 0xa2, 0x24, 0xc5, 0x91, 0x4d, 0xc5, 0xd5, 0x6a, 0x64, 0x54, 0x86,
	0xda, 0x44, 0x59, 0x51, 0x17, 0x87, 0x0f, 0x2e, 0x78, 0x97, 0x94, 0xb0, 0x1f, 0xe0, 0x2a, 0xc9,
	0x9d, 0x81, 0x13, 0x78, 0xa1, 0xab, 0x24, 0x13, 0xc0, 0x0a, 0x55, 0x60, 0xaa, 0x72, 0x9c, 0x54,
	0x2a, 0x95, 0x73, 0x8a, 0x67, 0x53, 0xee, 0x5a, 0xbf, 0xc1, 0x61, 0x01, 0xfc, 0x3c, 0x50, 0x67,
	0x53, 0xee, 0xd9, 0xf0, 0x7b, 0x99, 0xfd, 0x87, 0xb6, 0x36, 0x51, 0x69, 0x78, 0x6b, 0xe0, 0x04,
	0xdf, 0xc6, 0x7d, 0x51, 0x4f, 0x14, 0xfb, 0x89, 0xe2, 0x76, 0x3f, 0x31, 0xac, 0x83, 0xec, 0x02,
	0xbe, 0xa7, 0x91, 0x36, 0x57, 0x24, 0xd5, 0x4a, 0xa1, 0xe4, 0xed, 0x93, 0xc5, 0x83, 0x3c, 0x13,
	0xd0, 0x92, 0x94, 0x23, 0xef, 0x9c, 0xec, 0xd9, 0x1c, 0xfb, 0x03, 0xbe, 0x36, 0x58, 0x5c, 0x97,
	0x12, 0x4b, 0xfe, 0xc5, 0x7e, 0xf1, 0x2a, 0xb0, 0x5f, 0xe0, 0x6d, 0xa3, 0x94, 0x7f, 0x1d, 0x38,
	0x81, 0x1f, 0x3e, 0x9f, 0xc3, 0x47, 0x17, 0xe0, 0x06, 0xcb, 0xad, 0x5a, 0x62, 0x13, 0xca, 0x31,
	0x74, 0x77, 0x54, 0x6e, 0x56, 0x29, 0xed, 0x16, 0x24, 0x71, 0x4e, 0x71, 0x58, 0xe5, 0x2f, 0x30,
	0x1b, 0x3d, 0xf6, 0x0f, 0x7e, 0xbf, 0xd5, 0xeb, 0x42, 0x0d, 0xf4, 0xd8, 0xf8, 0x04, 0xa4, 0x63,
	0xe8, 0xea, 0xfa, 0xe3, 0x10, 0xef, 0x2b, 0x55, 0x62, 0x86, 0xb9, 0x99, 0x4d, 0x2d, 0x62, 0x2f,
	0x6c, 0xf4, 0xd8, 0x39, 0xf4, 0x8e, 0xf5, 0x45, 0x94, 0xa1, 0x65, 0xec, 0x87, 0x1f, 0xb8, 0xc7,
	0xc0, 0x27, 0xed, 0x3b, 0x4f, 0xcb, 0x4d, 0xdc, 0xb1, 0x33, 0xcf, 0x9e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x23, 0xd3, 0xee, 0x6a, 0x08, 0x03, 0x00, 0x00,
}
