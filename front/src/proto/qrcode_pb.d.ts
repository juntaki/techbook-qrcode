// package: qrcode
// file: qrcode.proto

import * as jspb from "google-protobuf";

export class URL extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): URL.AsObject;
  static toObject(includeInstance: boolean, msg: URL): URL.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: URL, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): URL;
  static deserializeBinaryFromReader(message: URL, reader: jspb.BinaryReader): URL;
}

export namespace URL {
  export type AsObject = {
    url: string,
  }
}

export class QRCode extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  getImage(): Uint8Array | string;
  getImage_asU8(): Uint8Array;
  getImage_asB64(): string;
  setImage(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QRCode.AsObject;
  static toObject(includeInstance: boolean, msg: QRCode): QRCode.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QRCode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QRCode;
  static deserializeBinaryFromReader(message: QRCode, reader: jspb.BinaryReader): QRCode;
}

export namespace QRCode {
  export type AsObject = {
    id: string,
    url: string,
    image: Uint8Array | string,
  }
}

export class QRCodeList extends jspb.Message {
  clearQrcodesList(): void;
  getQrcodesList(): Array<QRCode>;
  setQrcodesList(value: Array<QRCode>): void;
  addQrcodes(value?: QRCode, index?: number): QRCode;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QRCodeList.AsObject;
  static toObject(includeInstance: boolean, msg: QRCodeList): QRCodeList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QRCodeList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QRCodeList;
  static deserializeBinaryFromReader(message: QRCodeList, reader: jspb.BinaryReader): QRCodeList;
}

export namespace QRCodeList {
  export type AsObject = {
    qrcodesList: Array<QRCode.AsObject>,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

