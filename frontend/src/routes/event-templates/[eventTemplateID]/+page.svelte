<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/fields/InputField.svelte";
    import SelectField from "$lib/fields/SelectField.svelte";
    import type { PageData } from "./$types";
    import RoleTemplatesTable from "./RoleTemplatesTable.svelte";
    import { handleSubmit } from "$lib/utils";
    import EventsTable from "$lib/components/EventsTable.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";

    export let data: PageData;
    const currentReccurrence = data.eventTemplate.recurrence;

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

    const recurrenceOptions: { value: string | null; label: string }[] = [
        { value: "daily", label: "Daily" },
        { value: "weekly", label: "Weekly" },
        { value: "monthly", label: "Monthly" },
    ];
</script>

<div>
    <form method="POST" use:enhance={submit} action="?/updateEvent">
        <div class="py-4">
            <div>
                <span class="font-bold">Name:</span>
                <InputField name="name" value={data.eventTemplate.name} edit={editEventTemplate} />
            </div>

            <div>
                <span class="font-bold">Recurrence:</span>
                <SelectField
                    name="recurrence"
                    value={{ value: currentReccurrence, label: currentReccurrence }}
                    options={recurrenceOptions}
                    edit={editEventTemplate}
                />
            </div>

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
            <ActionButton type="submit" />
            <ActionButton type="cancel" onClick={() => (editEventTemplate = false)} />
        {:else}
            <ActionButton type="edit" onClick={() => (editEventTemplate = true)} />
        {/if}
    </form>

    <div class="py-8">
        <RoleTemplatesTable roleTemplates={data.roleTemplates} />
    </div>

    <div class="py-8">
        <EventsTable events={data.events} roles={data.roles} hideEventName />
    </div>
</div>
