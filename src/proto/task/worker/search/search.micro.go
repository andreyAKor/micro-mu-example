// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: search.proto

/*
Package geo_proto_task_worker_search is a generated protocol buffer package.

It is generated from these files:
	search.proto

It has these top-level messages:
	TaskWorkerSearchRunEventResponse
	TaskWorkerSearchRunEventLog
	TaskWorkerSearchRunRpcRequest
	TaskWorkerSearchRunRpcResponse
	TaskWorkerSearchStopRpcRequest
	TaskWorkerSearchStopRpcResponse
*/
package geo_proto_task_worker_search

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TaskWorkerSearchRun service

type TaskWorkerSearchRunService interface {
	// Пробивка станции
	TaskWorkerSearchRun(ctx context.Context, in *TaskWorkerSearchRunRpcRequest, opts ...client.CallOption) (*TaskWorkerSearchRunRpcResponse, error)
}

type taskWorkerSearchRunService struct {
	c    client.Client
	name string
}

func NewTaskWorkerSearchRunService(name string, c client.Client) TaskWorkerSearchRunService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geo.proto.task.worker.search"
	}
	return &taskWorkerSearchRunService{
		c:    c,
		name: name,
	}
}

func (c *taskWorkerSearchRunService) TaskWorkerSearchRun(ctx context.Context, in *TaskWorkerSearchRunRpcRequest, opts ...client.CallOption) (*TaskWorkerSearchRunRpcResponse, error) {
	req := c.c.NewRequest(c.name, "TaskWorkerSearchRun.TaskWorkerSearchRun", in)
	out := new(TaskWorkerSearchRunRpcResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskWorkerSearchRun service

type TaskWorkerSearchRunHandler interface {
	// Пробивка станции
	TaskWorkerSearchRun(context.Context, *TaskWorkerSearchRunRpcRequest, *TaskWorkerSearchRunRpcResponse) error
}

func RegisterTaskWorkerSearchRunHandler(s server.Server, hdlr TaskWorkerSearchRunHandler, opts ...server.HandlerOption) {
	type taskWorkerSearchRun interface {
		TaskWorkerSearchRun(ctx context.Context, in *TaskWorkerSearchRunRpcRequest, out *TaskWorkerSearchRunRpcResponse) error
	}
	type TaskWorkerSearchRun struct {
		taskWorkerSearchRun
	}
	h := &taskWorkerSearchRunHandler{hdlr}
	s.Handle(s.NewHandler(&TaskWorkerSearchRun{h}, opts...))
}

type taskWorkerSearchRunHandler struct {
	TaskWorkerSearchRunHandler
}

func (h *taskWorkerSearchRunHandler) TaskWorkerSearchRun(ctx context.Context, in *TaskWorkerSearchRunRpcRequest, out *TaskWorkerSearchRunRpcResponse) error {
	return h.TaskWorkerSearchRunHandler.TaskWorkerSearchRun(ctx, in, out)
}

// Client API for TaskWorkerSearchStop service

type TaskWorkerSearchStopService interface {
	// Остановка пробивки станции
	TaskWorkerSearchStop(ctx context.Context, in *TaskWorkerSearchStopRpcRequest, opts ...client.CallOption) (*TaskWorkerSearchStopRpcResponse, error)
}

type taskWorkerSearchStopService struct {
	c    client.Client
	name string
}

func NewTaskWorkerSearchStopService(name string, c client.Client) TaskWorkerSearchStopService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geo.proto.task.worker.search"
	}
	return &taskWorkerSearchStopService{
		c:    c,
		name: name,
	}
}

func (c *taskWorkerSearchStopService) TaskWorkerSearchStop(ctx context.Context, in *TaskWorkerSearchStopRpcRequest, opts ...client.CallOption) (*TaskWorkerSearchStopRpcResponse, error) {
	req := c.c.NewRequest(c.name, "TaskWorkerSearchStop.TaskWorkerSearchStop", in)
	out := new(TaskWorkerSearchStopRpcResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskWorkerSearchStop service

type TaskWorkerSearchStopHandler interface {
	// Остановка пробивки станции
	TaskWorkerSearchStop(context.Context, *TaskWorkerSearchStopRpcRequest, *TaskWorkerSearchStopRpcResponse) error
}

func RegisterTaskWorkerSearchStopHandler(s server.Server, hdlr TaskWorkerSearchStopHandler, opts ...server.HandlerOption) {
	type taskWorkerSearchStop interface {
		TaskWorkerSearchStop(ctx context.Context, in *TaskWorkerSearchStopRpcRequest, out *TaskWorkerSearchStopRpcResponse) error
	}
	type TaskWorkerSearchStop struct {
		taskWorkerSearchStop
	}
	h := &taskWorkerSearchStopHandler{hdlr}
	s.Handle(s.NewHandler(&TaskWorkerSearchStop{h}, opts...))
}

type taskWorkerSearchStopHandler struct {
	TaskWorkerSearchStopHandler
}

func (h *taskWorkerSearchStopHandler) TaskWorkerSearchStop(ctx context.Context, in *TaskWorkerSearchStopRpcRequest, out *TaskWorkerSearchStopRpcResponse) error {
	return h.TaskWorkerSearchStopHandler.TaskWorkerSearchStop(ctx, in, out)
}