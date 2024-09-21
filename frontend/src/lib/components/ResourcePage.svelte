<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { handleSubmit } from "$lib/utils";
    import { goto } from "$app/navigation";

    export let edit = false;
    export let title: string;
    export let baseURL: string;
    export let deleteBody: string;
    export let updateAction = "?/update";
    export let deleteAction = "?/delete";

    let deleteForm: HTMLFormElement;

    const submit: SubmitFunction = ({ action }) => {
        let successMsg = "Success";
        let onSuccess = undefined;

        switch (action.search) {
            case updateAction:
                successMsg = `${title} updated`;
                edit = false;
                break;
            case deleteAction:
                successMsg = `${title} deleted`;
                onSuccess = () => goto(baseURL);
                break;
        }
        return handleSubmit(successMsg, onSuccess);
    };
</script>

<form method="POST" action={deleteAction} use:enhance={submit} bind:this={deleteForm}></form>

<header class="flex-end card variant-filled-surface flex justify-end">
    {#if edit}
        <ActionButton type="submit" formID="form" />
        <ActionButton type="cancel" onClick={() => (edit = false)} />
    {:else}
        <ActionButton type="edit" onClick={() => (edit = true)} />
        <ActionButton type="delete" title="Delete {title}" body={deleteBody} form={deleteForm} />
    {/if}
</header>
<div>
    <form id="form" method="POST" action={updateAction} use:enhance={submit}>
        <div class="py-4">
            <slot />
        </div>
    </form>
</div>
