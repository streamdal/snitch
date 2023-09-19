import {
  Command,
  TailCommand,
} from "@streamdal/snitch-protos/protos/sp_command";
import {
  Audience,
  TailRequestType,
} from "@streamdal/snitch-protos/protos/sp_common";
import {
  Pipeline,
  PipelineStep,
} from "@streamdal/snitch-protos/protos/sp_pipeline";

import { Configs } from "../snitch.js";
import { audienceKey, internal } from "./register.js";

export type InternalPipeline = Pipeline & {
  tail?: boolean;
  paused?: boolean;
};

export type EnhancedStep = PipelineStep & {
  pipelineId: string;
  pipelineName: string;
};

export const initPipelines = async (configs: Configs) => {
  const { response } = await configs.grpcClient.getAttachCommandsByService(
    {
      serviceName: configs.serviceName.toLowerCase(),
    },
    { meta: { "auth-token": configs.snitchToken } }
  );

  for (const command of response.active) {
    processResponse(command);
  }
  internal.pipelineInitialized = true;
  return;
};

export const processResponse = (response: Command) => {
  if (!response.audience) {
    response.command.oneofKind !== "keepAlive" &&
      console.debug("command response has no audience, ignoring");
    return;
  }

  switch (response.command.oneofKind) {
    case "attachPipeline":
      response.command.attachPipeline.pipeline &&
        attachPipeline(
          response.audience,
          response.command.attachPipeline.pipeline
        );
      break;
    case "detachPipeline":
      detachPipeline(
        response.audience,
        response.command.detachPipeline.pipelineId
      );
      break;
    case "pausePipeline":
      togglePausePipeline(
        response.audience,
        response.command.pausePipeline.pipelineId,
        true
      );
      break;
    case "resumePipeline":
      togglePausePipeline(
        response.audience,
        response.command.resumePipeline.pipelineId,
        false
      );
      break;
    case "tail":
      tailPipeline(response.audience, response.command.tail);
      break;
  }
};

export const attachPipeline = (audience: Audience, pipeline: Pipeline) =>
  internal.pipelines.set(audienceKey(audience), {
    ...pipeline,
    paused: false,
  });

export const detachPipeline = (audience: Audience, pipelineId: string) => {
  const key = audienceKey(audience);
  const p = internal.pipelines.get(key);
  pipelineId === p?.id && internal.pipelines.delete(key);
};

export const togglePausePipeline = (
  audience: Audience,
  pipelineId: string,
  paused: boolean
) => {
  const key = audienceKey(audience);
  const p = internal.pipelines.get(key);
  pipelineId === p?.id && internal.pipelines.set(key, { ...p, paused });
};

export const tailPipeline = (audience: Audience, { request }: TailCommand) => {
  const key = audienceKey(audience);
  const p = internal.pipelines.get(key);
  request.pipelineId === p?.id &&
    internal.pipelines.set(key, {
      ...p,
      tail: request.type === TailRequestType.START,
    } as InternalPipeline);
};
