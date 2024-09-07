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
