<script lang="ts">
    import EventsTable from "$lib/components/EventsTable.svelte";
    import type { PageData } from "./$types";
    import { breadcrumbs } from "$lib/stores";
    import ScheduleEvent from "$lib/components/ScheduleEvent.svelte";

    export let data: PageData;

    breadcrumbs.clear().add("Events", "events");

    const eventTemplates = data.eventTemplates;
    const roleTemplates = data.roleTemplates;
    const users = data.users;
</script>

<ScheduleEvent {users} {eventTemplates} {roleTemplates} />

<EventsTable events={data.events} roles={data.roles} enableNav />
