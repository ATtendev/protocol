// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: vc_analytics.proto

package vc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AnalyticsRecorderService_IngestStats_FullMethodName          = "/vc.AnalyticsRecorderService/IngestStats"
	AnalyticsRecorderService_IngestEvents_FullMethodName         = "/vc.AnalyticsRecorderService/IngestEvents"
	AnalyticsRecorderService_IngestNodeRoomStates_FullMethodName = "/vc.AnalyticsRecorderService/IngestNodeRoomStates"
)

// AnalyticsRecorderServiceClient is the client API for AnalyticsRecorderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsRecorderServiceClient interface {
	IngestStats(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestStatsClient, error)
	IngestEvents(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestEventsClient, error)
	IngestNodeRoomStates(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestNodeRoomStatesClient, error)
}

type analyticsRecorderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsRecorderServiceClient(cc grpc.ClientConnInterface) AnalyticsRecorderServiceClient {
	return &analyticsRecorderServiceClient{cc}
}

func (c *analyticsRecorderServiceClient) IngestStats(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestStatsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[0], AnalyticsRecorderService_IngestStats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &analyticsRecorderServiceIngestStatsClient{ClientStream: stream}
	return x, nil
}

type AnalyticsRecorderService_IngestStatsClient interface {
	Send(*AnalyticsStats) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type analyticsRecorderServiceIngestStatsClient struct {
	grpc.ClientStream
}

func (x *analyticsRecorderServiceIngestStatsClient) Send(m *AnalyticsStats) error {
	return x.ClientStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestStatsClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *analyticsRecorderServiceClient) IngestEvents(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestEventsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[1], AnalyticsRecorderService_IngestEvents_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &analyticsRecorderServiceIngestEventsClient{ClientStream: stream}
	return x, nil
}

type AnalyticsRecorderService_IngestEventsClient interface {
	Send(*AnalyticsEvents) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type analyticsRecorderServiceIngestEventsClient struct {
	grpc.ClientStream
}

func (x *analyticsRecorderServiceIngestEventsClient) Send(m *AnalyticsEvents) error {
	return x.ClientStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestEventsClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *analyticsRecorderServiceClient) IngestNodeRoomStates(ctx context.Context, opts ...grpc.CallOption) (AnalyticsRecorderService_IngestNodeRoomStatesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[2], AnalyticsRecorderService_IngestNodeRoomStates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &analyticsRecorderServiceIngestNodeRoomStatesClient{ClientStream: stream}
	return x, nil
}

type AnalyticsRecorderService_IngestNodeRoomStatesClient interface {
	Send(*AnalyticsNodeRooms) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type analyticsRecorderServiceIngestNodeRoomStatesClient struct {
	grpc.ClientStream
}

func (x *analyticsRecorderServiceIngestNodeRoomStatesClient) Send(m *AnalyticsNodeRooms) error {
	return x.ClientStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestNodeRoomStatesClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnalyticsRecorderServiceServer is the server API for AnalyticsRecorderService service.
// All implementations must embed UnimplementedAnalyticsRecorderServiceServer
// for forward compatibility
type AnalyticsRecorderServiceServer interface {
	IngestStats(AnalyticsRecorderService_IngestStatsServer) error
	IngestEvents(AnalyticsRecorderService_IngestEventsServer) error
	IngestNodeRoomStates(AnalyticsRecorderService_IngestNodeRoomStatesServer) error
	mustEmbedUnimplementedAnalyticsRecorderServiceServer()
}

// UnimplementedAnalyticsRecorderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsRecorderServiceServer struct {
}

func (UnimplementedAnalyticsRecorderServiceServer) IngestStats(AnalyticsRecorderService_IngestStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method IngestStats not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) IngestEvents(AnalyticsRecorderService_IngestEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method IngestEvents not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) IngestNodeRoomStates(AnalyticsRecorderService_IngestNodeRoomStatesServer) error {
	return status.Errorf(codes.Unimplemented, "method IngestNodeRoomStates not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) mustEmbedUnimplementedAnalyticsRecorderServiceServer() {
}

// UnsafeAnalyticsRecorderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsRecorderServiceServer will
// result in compilation errors.
type UnsafeAnalyticsRecorderServiceServer interface {
	mustEmbedUnimplementedAnalyticsRecorderServiceServer()
}

func RegisterAnalyticsRecorderServiceServer(s grpc.ServiceRegistrar, srv AnalyticsRecorderServiceServer) {
	s.RegisterService(&AnalyticsRecorderService_ServiceDesc, srv)
}

func _AnalyticsRecorderService_IngestStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestStats(&analyticsRecorderServiceIngestStatsServer{ServerStream: stream})
}

type AnalyticsRecorderService_IngestStatsServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*AnalyticsStats, error)
	grpc.ServerStream
}

type analyticsRecorderServiceIngestStatsServer struct {
	grpc.ServerStream
}

func (x *analyticsRecorderServiceIngestStatsServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestStatsServer) Recv() (*AnalyticsStats, error) {
	m := new(AnalyticsStats)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AnalyticsRecorderService_IngestEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestEvents(&analyticsRecorderServiceIngestEventsServer{ServerStream: stream})
}

type AnalyticsRecorderService_IngestEventsServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*AnalyticsEvents, error)
	grpc.ServerStream
}

type analyticsRecorderServiceIngestEventsServer struct {
	grpc.ServerStream
}

func (x *analyticsRecorderServiceIngestEventsServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestEventsServer) Recv() (*AnalyticsEvents, error) {
	m := new(AnalyticsEvents)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AnalyticsRecorderService_IngestNodeRoomStates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestNodeRoomStates(&analyticsRecorderServiceIngestNodeRoomStatesServer{ServerStream: stream})
}

type AnalyticsRecorderService_IngestNodeRoomStatesServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*AnalyticsNodeRooms, error)
	grpc.ServerStream
}

type analyticsRecorderServiceIngestNodeRoomStatesServer struct {
	grpc.ServerStream
}

func (x *analyticsRecorderServiceIngestNodeRoomStatesServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *analyticsRecorderServiceIngestNodeRoomStatesServer) Recv() (*AnalyticsNodeRooms, error) {
	m := new(AnalyticsNodeRooms)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnalyticsRecorderService_ServiceDesc is the grpc.ServiceDesc for AnalyticsRecorderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsRecorderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vc.AnalyticsRecorderService",
	HandlerType: (*AnalyticsRecorderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "IngestStats",
			Handler:       _AnalyticsRecorderService_IngestStats_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "IngestEvents",
			Handler:       _AnalyticsRecorderService_IngestEvents_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "IngestNodeRoomStates",
			Handler:       _AnalyticsRecorderService_IngestNodeRoomStates_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "vc_analytics.proto",
}
