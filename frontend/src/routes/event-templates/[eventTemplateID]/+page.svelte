<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/fields/InputField.svelte";
    import SelectField from "$lib/fields/SelectField.svelte";
    import type { PageData } from "./$types";
    import Edit from "$lib/svg/Edit.svelte";
    import Check from "$lib/svg/Check.svelte";
    import Cancel from "$lib/svg/Cancel.svelte";
    import RoleTemplatesTable from "./RoleTemplatesTable.svelte";
    import { handleSubmit } from "$lib/utils";
    import EventsTable from "$lib/components/EventsTable.svelte";

    export let data: PageData;

    let editEventTemplate = false;

    const submit: SubmitFunction = ({ action }) => {
        let successMsg = "Successful";

        switch (action.search) {
            case "?/updateEvent":
                editEventTemplate = false;
                successMsg = "Event Template updated";
                break;
        }

        return handleSubmit(successMsg);
    };
</script>

<div>
    <form method="POST" use:enhance={submit} action="?/updateEvent">
        <div class="py-4">
            <InputField title="Name" value={data.eventTemplate.name} edit={editEventTemplate} />

            <SelectField
                title="Recurrence"
                value={data.eventTemplate.recurrence}
                options={["daily", "weekly", "monthly"]}
                edit={editEventTemplate}
            />

            <div>
                <span class="font-bold">Group:</span>
                {data.eventTemplate.expand?.group.name || ""}
            </div>
            <div>
                <span class="font-bold">Created at:</span>
                {data.eventTemplate.created}
            </div>
            <div>
                <span class="font-bold">Updated at:</span>
                {data.eventTemplate.updated}
            </div>
        </div>

        {#if editEventTemplate}
            <button type="submit" class="btn hover:variant-ringed-success"><Check /></button>
            <button
                type="button"
                class="btn hover:variant-ringed-secondary"
                on:click|preventDefault={() => (editEventTemplate = false)}
            >
                <Cancel />
            </button>
        {:else}
            <button
                type="submit"
                class="btn hover:variant-ringed-secondary"
                on:click|preventDefault={() => (editEventTemplate = true)}><Edit /></button
            >
        {/if}
    </form>

    <div class="py-8">
        <RoleTemplatesTable roleTemplates={data.roleTemplates} />
    </div>

    <div class="py-8">
        <EventsTable events={data.events} roles={data.roles} hideEventName />
    </div>
</div>
