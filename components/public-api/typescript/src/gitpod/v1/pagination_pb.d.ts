/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-es v0.1.1
// @generated from file gitpod/v1/pagination.proto (package gitpod.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

/**
 * @generated from message gitpod.v1.Pagination
 */
export declare class Pagination extends Message<Pagination> {
  /**
   * page_size is the maximum number of results we expect
   *
   * @generated from field: int32 page_size = 1;
   */
  pageSize: number;

  /**
   * page_token points to a specific page of the results
   *
   * @generated from field: string page_token = 2;
   */
  pageToken: string;

  constructor(data?: PartialMessage<Pagination>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "gitpod.v1.Pagination";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pagination;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pagination;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pagination;

  static equals(a: Pagination | PlainMessage<Pagination> | undefined, b: Pagination | PlainMessage<Pagination> | undefined): boolean;
}
