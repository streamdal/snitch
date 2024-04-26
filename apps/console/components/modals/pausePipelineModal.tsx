import { ActionModal } from "root/components/modals/actionModal.tsx";
import { Audience } from "streamdal-protos/protos/sp_common.ts";
import { Pipeline } from "streamdal-protos/protos/sp_pipeline.ts";
import IconPlayerPause from "tabler-icons/tsx/player-pause.tsx";
import { audienceKey, getAudienceOpRoute } from "../../lib/utils.ts";
import { opModal } from "../serviceMap/opModalSignal.ts";
import { serviceSignal } from "../serviceMap/serviceSignal.ts";
import { toastSignal } from "../toasts/toast.tsx";

export const togglePausePipeline = (
  paused: boolean,
  audience: Audience,
  pipeline: Pipeline,
) => {
  const key = audienceKey(audience);

  serviceSignal.value = {
    ...serviceSignal.value,
    configs: {
      ...serviceSignal.value.configs,
      [key]: {
        configs: serviceSignal.value?.configs[key]?.configs?.map((c) => ({
          ...c,
          ...c.id === pipeline.id ? { paused } : {},
        })),
      },
    },
  };
};

export const PausePipelineModal = ({ audience }: { audience: Audience }) => {
  const pipeline = opModal.value.pausePipeline;

  if (!pipeline) {
    toastSignal.value = {
      id: "pipelineCrud",
      type: "error",
      message: "Pipeline not found",
    };
    return null;
  }

  const close = () => opModal.value = { ...opModal.value, pausePipeline: null };

  const pause = async () => {
    const response = await fetch(
      `${getAudienceOpRoute(audience)}/pipeline/${pipeline.id}/pause`,
      {
        method: "POST",
      },
    );

    const { success } = await response.json();

    if (success.message) {
      togglePausePipeline(true, audience, pipeline);
      toastSignal.value = {
        id: "pipelineCrud",
        type: success.status ? "success" : "error",
        message: success.message,
      };
    }

    close();
  };

  return (
    <ActionModal
      icon={<IconPlayerPause class="w-10 h-10 mt-3 text-burnt" />}
      message={
        <div>
          Pause pipeline{" "}
          <span class="text-medium font-bold ">
            {pipeline.name}
          </span>?
        </div>
      }
      actionText="Pause"
      destructive={false}
      onClose={close}
      onAction={pause}
    />
  );
};
