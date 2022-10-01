/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-es v0.1.1
// @generated from file gitpod/v1/workspaces.proto (package gitpod.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {FieldMask, proto3, Timestamp} from "@bufbuild/protobuf";
import {Pagination} from "./pagination_pb.js";

/**
 * Admission level describes who can access a workspace instance and its ports.
 *
 * @generated from enum gitpod.v1.AdmissionLevel
 */
export const AdmissionLevel = proto3.makeEnum(
  "gitpod.v1.AdmissionLevel",
  [
    {no: 0, name: "ADMISSION_LEVEL_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ADMISSION_LEVEL_OWNER_ONLY", localName: "OWNER_ONLY"},
    {no: 2, name: "ADMISSION_LEVEL_EVERYONE", localName: "EVERYONE"},
  ],
);

/**
 * @generated from message gitpod.v1.ListWorkspacesRequest
 */
export const ListWorkspacesRequest = proto3.makeMessageType(
  "gitpod.v1.ListWorkspacesRequest",
  () => [
    { no: 1, name: "pagination", kind: "message", T: Pagination },
    { no: 2, name: "field_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message gitpod.v1.ListWorkspacesResponse
 */
export const ListWorkspacesResponse = proto3.makeMessageType(
  "gitpod.v1.ListWorkspacesResponse",
  () => [
    { no: 1, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "result", kind: "message", T: ListWorkspacesResponse_WorkspaceAndInstance, repeated: true },
  ],
);

/**
 * @generated from message gitpod.v1.ListWorkspacesResponse.WorkspaceAndInstance
 */
export const ListWorkspacesResponse_WorkspaceAndInstance = proto3.makeMessageType(
  "gitpod.v1.ListWorkspacesResponse.WorkspaceAndInstance",
  () => [
    { no: 1, name: "result", kind: "message", T: Workspace },
    { no: 2, name: "last_active_instances", kind: "message", T: WorkspaceInstance },
  ],
  {localName: "ListWorkspacesResponse_WorkspaceAndInstance"},
);

/**
 * @generated from message gitpod.v1.GetWorkspaceRequest
 */
export const GetWorkspaceRequest = proto3.makeMessageType(
  "gitpod.v1.GetWorkspaceRequest",
  () => [
    { no: 1, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.GetWorkspaceResponse
 */
export const GetWorkspaceResponse = proto3.makeMessageType(
  "gitpod.v1.GetWorkspaceResponse",
  () => [
    { no: 1, name: "result", kind: "message", T: Workspace },
  ],
);

/**
 * @generated from message gitpod.v1.GetOwnerTokenRequest
 */
export const GetOwnerTokenRequest = proto3.makeMessageType(
  "gitpod.v1.GetOwnerTokenRequest",
  () => [
    { no: 1, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.GetOwnerTokenResponse
 */
export const GetOwnerTokenResponse = proto3.makeMessageType(
  "gitpod.v1.GetOwnerTokenResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.CreateAndStartWorkspaceRequest
 */
export const CreateAndStartWorkspaceRequest = proto3.makeMessageType(
  "gitpod.v1.CreateAndStartWorkspaceRequest",
  () => [
    { no: 1, name: "idempotency_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "context_url", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "source" },
    { no: 3, name: "prebuild_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "source" },
    { no: 5, name: "start_spec", kind: "message", T: StartWorkspaceSpec },
  ],
);

/**
 * @generated from message gitpod.v1.CreateAndStartWorkspaceResponse
 */
export const CreateAndStartWorkspaceResponse = proto3.makeMessageType(
  "gitpod.v1.CreateAndStartWorkspaceResponse",
  () => [
    { no: 1, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.StartWorkspaceRequest
 */
export const StartWorkspaceRequest = proto3.makeMessageType(
  "gitpod.v1.StartWorkspaceRequest",
  () => [
    { no: 1, name: "idempotency_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "spec", kind: "message", T: StartWorkspaceSpec },
  ],
);

/**
 * @generated from message gitpod.v1.StartWorkspaceResponse
 */
export const StartWorkspaceResponse = proto3.makeMessageType(
  "gitpod.v1.StartWorkspaceResponse",
  () => [
    { no: 1, name: "instance_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "workspace_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.GetActiveWorkspaceInstanceRequest
 */
export const GetActiveWorkspaceInstanceRequest = proto3.makeMessageType(
  "gitpod.v1.GetActiveWorkspaceInstanceRequest",
  () => [
    { no: 1, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.GetActiveWorkspaceInstanceResponse
 */
export const GetActiveWorkspaceInstanceResponse = proto3.makeMessageType(
  "gitpod.v1.GetActiveWorkspaceInstanceResponse",
  () => [
    { no: 1, name: "instance", kind: "message", T: WorkspaceInstance },
  ],
);

/**
 * @generated from message gitpod.v1.GetWorkspaceInstanceOwnerTokenRequest
 */
export const GetWorkspaceInstanceOwnerTokenRequest = proto3.makeMessageType(
  "gitpod.v1.GetWorkspaceInstanceOwnerTokenRequest",
  () => [
    { no: 1, name: "instance_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.GetWorkspaceInstanceOwnerTokenResponse
 */
export const GetWorkspaceInstanceOwnerTokenResponse = proto3.makeMessageType(
  "gitpod.v1.GetWorkspaceInstanceOwnerTokenResponse",
  () => [
    { no: 1, name: "owner_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.ListenToWorkspaceInstanceRequest
 */
export const ListenToWorkspaceInstanceRequest = proto3.makeMessageType(
  "gitpod.v1.ListenToWorkspaceInstanceRequest",
  () => [
    { no: 1, name: "instance_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.ListenToWorkspaceInstanceResponse
 */
export const ListenToWorkspaceInstanceResponse = proto3.makeMessageType(
  "gitpod.v1.ListenToWorkspaceInstanceResponse",
  () => [
    { no: 1, name: "instance_status", kind: "message", T: WorkspaceInstanceStatus },
  ],
);

/**
 * @generated from message gitpod.v1.ListenToImageBuildLogsRequest
 */
export const ListenToImageBuildLogsRequest = proto3.makeMessageType(
  "gitpod.v1.ListenToImageBuildLogsRequest",
  () => [
    { no: 1, name: "instance_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.ListenToImageBuildLogsResponse
 */
export const ListenToImageBuildLogsResponse = proto3.makeMessageType(
  "gitpod.v1.ListenToImageBuildLogsResponse",
  () => [
    { no: 1, name: "line", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.StopWorkspaceRequest
 */
export const StopWorkspaceRequest = proto3.makeMessageType(
  "gitpod.v1.StopWorkspaceRequest",
  () => [
    { no: 1, name: "idempotency_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message gitpod.v1.StopWorkspaceResponse
 */
export const StopWorkspaceResponse = proto3.makeMessageType(
  "gitpod.v1.StopWorkspaceResponse",
  [],
);

/**
 * Workspace describes a single workspace
 *
 * @generated from message gitpod.v1.Workspace
 */
export const Workspace = proto3.makeMessageType(
  "gitpod.v1.Workspace",
  () => [
    { no: 1, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "owner_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "project_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "context", kind: "message", T: WorkspaceContext },
    { no: 5, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * WorkspaceContext describes the context a workspace was created from
 *
 * @generated from message gitpod.v1.WorkspaceContext
 */
export const WorkspaceContext = proto3.makeMessageType(
  "gitpod.v1.WorkspaceContext",
  () => [
    { no: 1, name: "context_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "git", kind: "message", T: WorkspaceContext_Git, oneof: "details" },
    { no: 3, name: "prebuild", kind: "message", T: WorkspaceContext_Prebuild, oneof: "details" },
    { no: 4, name: "snapshot", kind: "message", T: WorkspaceContext_Snapshot, oneof: "details" },
  ],
);

/**
 * Explicit Git context
 *
 * @generated from message gitpod.v1.WorkspaceContext.Git
 */
export const WorkspaceContext_Git = proto3.makeMessageType(
  "gitpod.v1.WorkspaceContext.Git",
  () => [
    { no: 1, name: "normalized_context_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "commit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
  {localName: "WorkspaceContext_Git"},
);

/**
 * Workspace was created from a prebuild
 *
 * @generated from message gitpod.v1.WorkspaceContext.Prebuild
 */
export const WorkspaceContext_Prebuild = proto3.makeMessageType(
  "gitpod.v1.WorkspaceContext.Prebuild",
  () => [
    { no: 1, name: "original_context", kind: "message", T: WorkspaceContext_Git },
    { no: 2, name: "prebuild_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
  {localName: "WorkspaceContext_Prebuild"},
);

/**
 * Snapshot context points to the snapshot which the workspace was created from
 *
 * @generated from message gitpod.v1.WorkspaceContext.Snapshot
 */
export const WorkspaceContext_Snapshot = proto3.makeMessageType(
  "gitpod.v1.WorkspaceContext.Snapshot",
  () => [
    { no: 1, name: "snapshot_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
  {localName: "WorkspaceContext_Snapshot"},
);

/**
 * WorkspaceInstance describes a single workspace instance
 *
 * @generated from message gitpod.v1.WorkspaceInstance
 */
export const WorkspaceInstance = proto3.makeMessageType(
  "gitpod.v1.WorkspaceInstance",
  () => [
    { no: 1, name: "instance_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created_at", kind: "message", T: Timestamp },
    { no: 4, name: "status", kind: "message", T: WorkspaceInstanceStatus },
  ],
);

/**
 * WorkspaceStatus describes a workspace status
 *
 * @generated from message gitpod.v1.WorkspaceInstanceStatus
 */
export const WorkspaceInstanceStatus = proto3.makeMessageType(
  "gitpod.v1.WorkspaceInstanceStatus",
  () => [
    { no: 1, name: "status_version", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "phase", kind: "enum", T: proto3.getEnumType(WorkspaceInstanceStatus_Phase) },
    { no: 3, name: "conditions", kind: "message", T: WorkspaceInstanceStatus_Conditions },
    { no: 4, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "admission", kind: "enum", T: proto3.getEnumType(AdmissionLevel) },
  ],
);

/**
 * Phase is a simple, high-level summary of where the workspace instance is in its lifecycle.
 * The phase is not intended to be a comprehensive rollup of observations of the workspace state,
 * nor is it intended to be a comprehensive state machine.
 * (based on  https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase)
 *
 * @generated from enum gitpod.v1.WorkspaceInstanceStatus.Phase
 */
export const WorkspaceInstanceStatus_Phase = proto3.makeEnum(
  "gitpod.v1.WorkspaceInstanceStatus.Phase",
  [
    {no: 0, name: "PHASE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "PHASE_PREPARING", localName: "PREPARING"},
    {no: 2, name: "PHASE_IMAGEBUILD", localName: "IMAGEBUILD"},
    {no: 3, name: "PHASE_PENDING", localName: "PENDING"},
    {no: 4, name: "PHASE_CREATING", localName: "CREATING"},
    {no: 5, name: "PHASE_INITIALIZING", localName: "INITIALIZING"},
    {no: 6, name: "PHASE_RUNNING", localName: "RUNNING"},
    {no: 7, name: "PHASE_INTERRUPTED", localName: "INTERRUPTED"},
    {no: 8, name: "PHASE_STOPPING", localName: "STOPPING"},
    {no: 9, name: "PHASE_STOPPED", localName: "STOPPED"},
  ],
);

/**
 * Conditions gives more detailed information as to the state of the workspace. Which condition actually
 * has a value depends on the phase the workspace is in.
 *
 * @generated from message gitpod.v1.WorkspaceInstanceStatus.Conditions
 */
export const WorkspaceInstanceStatus_Conditions = proto3.makeMessageType(
  "gitpod.v1.WorkspaceInstanceStatus.Conditions",
  () => [
    { no: 1, name: "failed", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "timeout", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "first_user_activity", kind: "message", T: Timestamp },
    { no: 11, name: "stopped_by_request", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
  {localName: "WorkspaceInstanceStatus_Conditions"},
);

/**
 * StartWorkspaceSpec influences the workspace start
 *
 * future per-workspace-start fields, e.g. region
 *
 * @generated from message gitpod.v1.StartWorkspaceSpec
 */
export const StartWorkspaceSpec = proto3.makeMessageType(
  "gitpod.v1.StartWorkspaceSpec",
  [],
);
