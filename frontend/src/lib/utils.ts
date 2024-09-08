import { invalidateAll } from "$app/navigation";
import toast from "svelte-french-toast";

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
