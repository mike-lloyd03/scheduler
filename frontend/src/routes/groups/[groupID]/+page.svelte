<script lang="ts">
    import { enhance } from "$app/forms";
    import toast from "svelte-french-toast";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { handleSubmit } from "$lib/utils";

    export let data: PageData;

    let edit = false;

    const submit: SubmitFunction = () => {
        edit = false;
        return handleSubmit("Group updated");
    };
</script>

<header class="flex-end card variant-filled-surface flex justify-end">
    {#if edit}
        <ActionButton type="submit" formID="form" />
        <ActionButton type="cancel" onClick={() => (edit = false)} />
    {:else}
        <ActionButton type="edit" onClick={() => (edit = true)} />
    {/if}
</header>
<div>
    <form id="form" method="POST" use:enhance={submit}>
        <div class="py-4">
            <p>
                <span class="font-bold">Name:</span>
                <InputField name="name" value={data.group.name} {edit} />
            </p>
            <p><span class="font-bold">Organization:</span> {data.group.expand?.org.name}</p>
            <p><span class="font-bold">Created at:</span> {data.group.created}</p>
            <p><span class="font-bold">Updated at:</span> {data.group.updated}</p>
        </div>
    </form>
</div>
