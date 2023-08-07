import IconExclamationCircle from "tabler-icons/tsx/exclamation-circle.tsx";
import IconX from "tabler-icons/tsx/x.tsx";
import { useEffect, useState } from "preact/hooks";

export const ErrorToast = (
  { message, open = true, autoClose = true, closeDuration = 3000 }: {
    message: string;
    open?: boolean;
    autoClose?: boolean;
    closeDuration?: number;
  },
) => {
  const [toastOpen, setToastOpen] = useState(open);

  useEffect(() => {
    if (toastOpen && autoClose) {
      setTimeout(() => {
        setToastOpen(false);
      }, closeDuration);
    }
  }, [toastOpen]);

  return (
    toastOpen
      ? (
        <div
          id="toast-danger"
          class="fixed top-[8%] left-[40%] z-50 flex items-center w-full max-w-xs p-4 mb-4 text-gray-500 bg-white rounded-lg border border-streamdalRed"
          role="alert"
        >
          <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-red-500 bg-red-100 rounded-lg ">
            <IconExclamationCircle class="w-6 h-6 text-streamdalRed" />
          </div>
          <div class="ml-3 text-sm font-normal">{message}</div>
          <button
            type="button"
            class="ml-auto -mx-1.5 -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex items-center justify-center h-8 w-8 "
            data-dismiss-target="#toast-danger"
            aria-label="Close"
            onClick={() => setToastOpen(false)}
          >
            <IconX class="w-6 h-6" />
          </button>
        </div>
      )
      : null
  );
};
