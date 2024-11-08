<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import SelectField from "$lib/fields/SelectField.svelte";
    import type { PageData } from "./$types";
    import RoleTemplatesTable from "./RoleTemplatesTable.svelte";
    import EventsTable from "$lib/components/EventsTable.svelte";
    import ResourcePage from "$lib/components/ResourcePage.svelte";
    import { hasGroupRole, hasOrgRole, toLocaleDateTime } from "$lib/utils";
    import { breadcrumbs } from "$lib/stores";
    import ScheduleEvent from "$lib/components/ScheduleEvent.svelte";
    import { UserRole } from "$lib/types";

    let { data }: { data: PageData } = $props();
    const permissions = data.currentUserPermissions ?? [];

    $effect(() => {
        breadcrumbs
            .clear()
            .add("Event Templates", "event-templates")
            .add(data.eventTemplate.name, data.eventTemplate.id);
    });

    const currentReccurrence = data.eventTemplate.recurrence;
    let editEventTemplate = $state(false);

    const recurrenceOptions: { value: string | null; label: string }[] = [
        { value: "daily", label: "Daily" },
        { value: "weekly", label: "Weekly" },
        { value: "monthly", label: "Monthly" },
    ];

    let org = $derived(data.eventTemplate.expand?.group.expand!.org);
    let group = $derived(data.eventTemplate.expand?.group);

    let showEdit = $state(false);
    let showDelete = $state(false);

    $effect(() => {
        if (org && group) {
            showEdit =
                hasOrgRole(permissions ?? [], org, UserRole.Admin) ||
                hasOrgRole(permissions ?? [], org, UserRole.Manager) ||
                hasGroupRole(permissions ?? [], group, UserRole.Admin) ||
                hasGroupRole(permissions ?? [], group, UserRole.Manager);
            showDelete = showEdit;
        }
    });
</script>

<ResourcePage
    bind:edit={editEventTemplate}
    resourceName="Event Template"
    title={data.eventTemplate.name}
    baseURL="/event-templates"
    deleteBody="Are you sure you want to delete the event template '{data.eventTemplate
        .name}'? This will delete all associated events."
    updateAction="?/updateEventTemplate"
    deleteAction="?/deleteEventTemplate"
    {showEdit}
    {showDelete}
>
    <div class="py-4">
        <div>
            <span class="font-bold">Name:</span>
            <InputField
                name="name"
                value={data.eventTemplate.name}
                edit={editEventTemplate}
                form="updateForm"
            />
        </div>

        <div>
            <span class="font-bold">Recurrence:</span>
            <SelectField
                name="recurrence"
                value={{ value: currentReccurrence, label: currentReccurrence }}
                options={recurrenceOptions}
                edit={editEventTemplate}
                form="updateForm"
            />
        </div>

        <div>
            <span class="font-bold">Group:</span>
            {data.eventTemplate.expand?.group.name || ""}
        </div>
        <div>
            <span class="font-bold">Created at:</span>
            {toLocaleDateTime(data.eventTemplate.created)}
        </div>
        <div>
            <span class="font-bold">Updated at:</span>
            {toLocaleDateTime(data.eventTemplate.updated)}
        </div>
    </div>

    {#if showEdit}
        <ScheduleEvent
            users={data.users}
            eventTemplates={data.eventTemplates}
            roleTemplates={data.roleTemplates}
            selectedEventTemplateID={data.eventTemplate.id}
        />
    {/if}

    <div class="py-8">
        <RoleTemplatesTable
            roleTemplates={data.roleTemplates}
            showNew={showEdit}
            {showEdit}
            {showDelete}
        />
    </div>

    <div class="py-8">
        <EventsTable events={data.events} roles={data.roles} hideEventName enableNav {showDelete} />
    </div>
</ResourcePage>
