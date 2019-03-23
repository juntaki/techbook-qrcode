// package: qrcode
// file: qrcode.proto

import * as qrcode_pb from "./qrcode_pb";
import {grpc} from "@improbable-eng/grpc-web";

type QRCodeServiceGetURL = {
  readonly methodName: string;
  readonly service: typeof QRCodeService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof qrcode_pb.Empty;
  readonly responseType: typeof qrcode_pb.URL;
};

type QRCodeServiceUpdateURL = {
  readonly methodName: string;
  readonly service: typeof QRCodeService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof qrcode_pb.URL;
  readonly responseType: typeof qrcode_pb.Empty;
};

type QRCodeServiceGetQRCodes = {
  readonly methodName: string;
  readonly service: typeof QRCodeService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof qrcode_pb.Empty;
  readonly responseType: typeof qrcode_pb.QRCodeList;
};

type QRCodeServiceAddQRCodes = {
  readonly methodName: string;
  readonly service: typeof QRCodeService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof qrcode_pb.Empty;
  readonly responseType: typeof qrcode_pb.Empty;
};

export class QRCodeService {
  static readonly serviceName: string;
  static readonly GetURL: QRCodeServiceGetURL;
  static readonly UpdateURL: QRCodeServiceUpdateURL;
  static readonly GetQRCodes: QRCodeServiceGetQRCodes;
  static readonly AddQRCodes: QRCodeServiceAddQRCodes;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class QRCodeServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getURL(
    requestMessage: qrcode_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.URL|null) => void
  ): UnaryResponse;
  getURL(
    requestMessage: qrcode_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.URL|null) => void
  ): UnaryResponse;
  updateURL(
    requestMessage: qrcode_pb.URL,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.Empty|null) => void
  ): UnaryResponse;
  updateURL(
    requestMessage: qrcode_pb.URL,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.Empty|null) => void
  ): UnaryResponse;
  getQRCodes(
    requestMessage: qrcode_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.QRCodeList|null) => void
  ): UnaryResponse;
  getQRCodes(
    requestMessage: qrcode_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.QRCodeList|null) => void
  ): UnaryResponse;
  addQRCodes(
    requestMessage: qrcode_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.Empty|null) => void
  ): UnaryResponse;
  addQRCodes(
    requestMessage: qrcode_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: qrcode_pb.Empty|null) => void
  ): UnaryResponse;
}

