import { ActionModal } from "./actionModal.tsx";
import { Audience } from "streamdal-protos/protos/sp_common.ts";
import IconTrash from "tabler-icons/tsx/trash.tsx";
import { audienceKey, getAudienceOpRoute } from "../../lib/utils.ts";
import { opModal } from "../../components/serviceMap/opModalSignal.ts";
import { showToast } from "root/islands/toasts.tsx";

export const DeleteOperationModal = (
  { audience }: { audience: Audience },
) => {
  const close = () => opModal.value = { clients: 0 };

  const deleteOp = async () => {
    const response = await fetch(
      `${getAudienceOpRoute(audience)}/delete`,
      {
        method: "POST",
      },
    );

    const { success } = await response.json();

    if (success.message) {
      showToast({
        id: "operationDelete",
        type: success.status ? "success" : "error",
        message: success.message,
      });
      opModal.value = { ...opModal.value, deleteOperation: false };
    }
  };

  return (
    <>
      <ActionModal
        icon={<IconTrash class="w-10 h-10text-eyelid" />}
        message={
          <div>
            Delete operation{" "}
            <span class="text-medium font-bold ">{audience.operationName}
            </span>?
          </div>
        }
        actionText="Delete"
        destructive={true}
        onClose={close}
        onAction={deleteOp}
      />
    </>
  );
};
