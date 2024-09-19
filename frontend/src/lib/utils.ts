import { invalidateAll } from "$app/navigation";
import toast from "svelte-french-toast";
import { type ModalComponent, type ModalSettings, type ModalStore } from "@skeletonlabs/skeleton";

export function handleSubmit(successMsg: string) {
    return async ({ result, update }) => {
        switch (result.type) {
            case "success":
                toast.success(successMsg);
                await update();
                break;
            case "failure":
                toast.error(result.data?.message || "Something went wrong");
                break;
            default:
                await update();
        }
    };
}

export function runAction(action: string, body?: string) {
    fetch(action, {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: body,
    }).then(() => invalidateAll());
}

export function jsonToFormString(jsonData: object): string {
    const formData: string[] = [];

    Object.entries(jsonData).forEach(([k, v]) => formData.push(`${k}=${v}`));

    console.log("jsonToForm result: ");

    return formData.join("&");
}

export function modalComponentForm(
    component: any,
    form: HTMLFormElement,
    modalStore: ModalStore,
    presubmit?: (r: any) => void,
    meta?: object,
): void {
    const c: ModalComponent = { ref: component };

    const modal: ModalSettings = {
        type: "component",
        component: c,
        meta,
        response: (r) => {
            if (r) {
                if (presubmit) {
                    presubmit(r);
                }
                form.requestSubmit();
            }
        },
    };
    modalStore.trigger(modal);
}
