import * as jspb from 'google-protobuf'



export class Paginator extends jspb.Message {
  getPage(): number;
  setPage(value: number): Paginator;

  getLimit(): number;
  setLimit(value: number): Paginator;

  getTotal(): number;
  setTotal(value: number): Paginator;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Paginator.AsObject;
  static toObject(includeInstance: boolean, msg: Paginator): Paginator.AsObject;
  static serializeBinaryToWriter(message: Paginator, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Paginator;
  static deserializeBinaryFromReader(message: Paginator, reader: jspb.BinaryReader): Paginator;
}

export namespace Paginator {
  export type AsObject = {
    page: number,
    limit: number,
    total: number,
  }
}

