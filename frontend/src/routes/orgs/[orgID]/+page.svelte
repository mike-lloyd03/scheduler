<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { handleSubmit } from "$lib/utils";
    import { goto } from "$app/navigation";

    export let data: PageData;

    let edit = false;
    let deleteForm: HTMLFormElement;

    const submit: SubmitFunction = ({ action }) => {
        let successMsg = "Success";
        let onSuccess = undefined;

        switch (action.search) {
            case "?/update":
                successMsg = "Organization updated";
                edit = false;
                break;
            case "?/delete":
                successMsg = "Organization deleted";
                onSuccess = () => goto("/orgs");
                break;
        }
        return handleSubmit(successMsg, onSuccess);
    };
</script>

<form method="POST" action="?/delete" use:enhance={submit} bind:this={deleteForm}></form>

<header class="flex-end card variant-filled-surface flex justify-end">
    {#if edit}
        <ActionButton type="submit" formID="form" />
        <ActionButton type="cancel" onClick={() => (edit = false)} />
    {:else}
        <ActionButton type="edit" onClick={() => (edit = true)} />
        <ActionButton
            type="delete"
            title="Delete Organization"
            body="Are you sure you want to delete the organization '{data.org
                .name}'? This will delete all groups, events, and other resources associated with it."
            form={deleteForm}
        />
    {/if}
</header>
<div>
    <form id="form" method="POST" action="?/update" use:enhance={submit}>
        <div class="py-4">
            <p>
                <span class="font-bold">Name:</span>
                <InputField name="name" value={data.org.name} {edit} />
            </p>
            <p><span class="font-bold">Created at:</span> {data.org.created}</p>
            <p><span class="font-bold">Updated at:</span> {data.org.updated}</p>
        </div>
    </form>
</div>
