import {
  Toast,
  ToastBody,
  ToastTitle,
  ToastIntent,
  useToastController,
} from '@fluentui/react-components';

export const TOASTER_ID = 'toaster';

export function useCustomNotifications() {
  return useToastController(TOASTER_ID);
}

export function useNotifications() {
  const { dispatchToast, dismissToast } = useCustomNotifications();

  const addToast = (
    title: string,
    intent: ToastIntent,
    props?: {
      message?: string;
      subtitle?: string;
      toastId?: string;
      timeout?: number;
    }
  ) => {
    return dispatchToast(
      <Toast>
        <ToastTitle>{title}</ToastTitle>
        {props?.message || props?.subtitle ? (
          <ToastBody {...(props?.subtitle ? { subtitle: props.subtitle } : {})}>
            {props.message}
          </ToastBody>
        ) : null}
      </Toast>,
      {
        intent,
        timeout: props?.timeout ?? 5000,
        ...(props?.toastId ? { toastId: props.toastId } : {}),
      }
    );
  };

  const removeToast = (toastId: string) => {
    dismissToast(toastId);
  };

  return {
    addToast,
    removeToast,
  };
}
