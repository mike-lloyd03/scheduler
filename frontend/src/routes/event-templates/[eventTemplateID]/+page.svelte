<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/fields/InputField.svelte";
    import SelectField from "$lib/fields/SelectField.svelte";
    import type { PageData } from "./$types";
    import RoleTemplatesTable from "./RoleTemplatesTable.svelte";
    import { handleSubmit } from "$lib/utils";
    import EventsTable from "$lib/components/EventsTable.svelte";
    import { goto } from "$app/navigation";
    import ResourcePage from "$lib/components/ResourcePage.svelte";

    export let data: PageData;
    const currentReccurrence = data.eventTemplate.recurrence;

    let form: HTMLFormElement;

    let editEventTemplate = false;

    // const submit: SubmitFunction = ({ action }) => {
    //     let successMsg = "Successful";
    //
    //     switch (action.search) {
    //         case "?/updateEvent":
    //             editEventTemplate = false;
    //             successMsg = "Event Template updated";
    //             break;
    //         case "?/deleteEventTemplate":
    //             return goto("/event-templates");
    //     }
    //
    //     return handleSubmit(successMsg);
    // };

    const recurrenceOptions: { value: string | null; label: string }[] = [
        { value: "daily", label: "Daily" },
        { value: "weekly", label: "Weekly" },
        { value: "monthly", label: "Monthly" },
    ];
</script>

<ResourcePage
    bind:edit={editEventTemplate}
    title="Event Template"
    baseURL="/event-templates"
    deleteBody="Are you sure you want to delete the event template '{data.eventTemplate
        .name}'? This will delete all associated events."
    updateAction="?/updateEventTemplate"
    deleteAction="?/deleteEventTemplate"
>
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

    <div class="py-8">
        <RoleTemplatesTable roleTemplates={data.roleTemplates} />
    </div>

    <div class="py-8">
        <EventsTable events={data.events} roles={data.roles} hideEventName enableNav />
    </div>
</ResourcePage>
