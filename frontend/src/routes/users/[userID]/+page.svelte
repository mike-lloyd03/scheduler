<script lang="ts">
    import { enhance } from "$app/forms";
    import toast from "svelte-french-toast";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ActionButton from "$lib/components/ActionButton.svelte";

    export let data: PageData;

    let edit = false;
    let loading = false;

    const submit: SubmitFunction = ({ cancel }) => {
        loading = true;

        if (edit == false) {
            edit = true;
            cancel();
            return;
        }

        return async ({ result, update }) => {
            edit = false;
            switch (result.type) {
                case "success":
                    toast.success("User updated");
                    await update();
                    break;
                case "failure":
                    toast.error(result.data?.message || "Something went wrong");
                    break;
                default:
                    await update();
            }
            loading = false;
        };
    };
</script>

<div>
    <form method="POST" use:enhance={submit}>
        <div class="py-4">
            <p>
                <InputField title="Name" value={data.group.name} {edit} />
            </p>
            <p><span class="font-bold">Created at:</span> {data.group.created}</p>
            <p><span class="font-bold">Updated at:</span> {data.group.updated}</p>
        </div>
        <ActionButton type="edit" onClick={() => (edit = true)} />

        {#if edit}
            <ActionButton type="submit" />
            <ActionButton type="cancel" onClick={() => (edit = false)} />
        {/if}
    </form>
</div>
