/**
 * @fileoverview gRPC-Web generated client stub for protos
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: internal_api.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as internal_api_pb from './internal_api_pb';
import * as common_pb from './common_pb';


export class InternalClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorRegister = new grpcWeb.MethodDescriptor(
    '/protos.Internal/Register',
    grpcWeb.MethodType.SERVER_STREAMING,
    internal_api_pb.RegisterRequest,
    internal_api_pb.CommandResponse,
    (request: internal_api_pb.RegisterRequest) => {
      return request.serializeBinary();
    },
    internal_api_pb.CommandResponse.deserializeBinary
  );

  register(
    request: internal_api_pb.RegisterRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<internal_api_pb.CommandResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/protos.Internal/Register',
      request,
      metadata || {},
      this.methodDescriptorRegister);
  }

  methodDescriptorHeartbeat = new grpcWeb.MethodDescriptor(
    '/protos.Internal/Heartbeat',
    grpcWeb.MethodType.UNARY,
    internal_api_pb.HeartbeatRequest,
    common_pb.StandardResponse,
    (request: internal_api_pb.HeartbeatRequest) => {
      return request.serializeBinary();
    },
    common_pb.StandardResponse.deserializeBinary
  );

  heartbeat(
    request: internal_api_pb.HeartbeatRequest,
    metadata: grpcWeb.Metadata | null): Promise<common_pb.StandardResponse>;

  heartbeat(
    request: internal_api_pb.HeartbeatRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: common_pb.StandardResponse) => void): grpcWeb.ClientReadableStream<common_pb.StandardResponse>;

  heartbeat(
    request: internal_api_pb.HeartbeatRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: common_pb.StandardResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protos.Internal/Heartbeat',
        request,
        metadata || {},
        this.methodDescriptorHeartbeat,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protos.Internal/Heartbeat',
    request,
    metadata || {},
    this.methodDescriptorHeartbeat);
  }

  methodDescriptorNotify = new grpcWeb.MethodDescriptor(
    '/protos.Internal/Notify',
    grpcWeb.MethodType.UNARY,
    internal_api_pb.NotifyRequest,
    common_pb.StandardResponse,
    (request: internal_api_pb.NotifyRequest) => {
      return request.serializeBinary();
    },
    common_pb.StandardResponse.deserializeBinary
  );

  notify(
    request: internal_api_pb.NotifyRequest,
    metadata: grpcWeb.Metadata | null): Promise<common_pb.StandardResponse>;

  notify(
    request: internal_api_pb.NotifyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: common_pb.StandardResponse) => void): grpcWeb.ClientReadableStream<common_pb.StandardResponse>;

  notify(
    request: internal_api_pb.NotifyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: common_pb.StandardResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protos.Internal/Notify',
        request,
        metadata || {},
        this.methodDescriptorNotify,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protos.Internal/Notify',
    request,
    metadata || {},
    this.methodDescriptorNotify);
  }

  methodDescriptorMetrics = new grpcWeb.MethodDescriptor(
    '/protos.Internal/Metrics',
    grpcWeb.MethodType.UNARY,
    internal_api_pb.MetricsRequest,
    common_pb.StandardResponse,
    (request: internal_api_pb.MetricsRequest) => {
      return request.serializeBinary();
    },
    common_pb.StandardResponse.deserializeBinary
  );

  metrics(
    request: internal_api_pb.MetricsRequest,
    metadata: grpcWeb.Metadata | null): Promise<common_pb.StandardResponse>;

  metrics(
    request: internal_api_pb.MetricsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: common_pb.StandardResponse) => void): grpcWeb.ClientReadableStream<common_pb.StandardResponse>;

  metrics(
    request: internal_api_pb.MetricsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: common_pb.StandardResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protos.Internal/Metrics',
        request,
        metadata || {},
        this.methodDescriptorMetrics,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protos.Internal/Metrics',
    request,
    metadata || {},
    this.methodDescriptorMetrics);
  }

}

