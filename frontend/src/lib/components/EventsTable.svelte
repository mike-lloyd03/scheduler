<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import type { Event, Role } from "$lib/types";
    import { handleSubmit } from "$lib/utils";
    import { goto } from "$app/navigation";
    import ActionButton from "./ActionButton.svelte";

    export let events: Event[];
    export let roles: Role[];
    export let hideEventName = false;
    export let enableNav = false;

    let deleteEventForm: HTMLFormElement;

    let deleteEventID: string | null;

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Successful";

        switch (action.pathname + action.search) {
            case "/events?/deleteEvent":
                successMsg = "Event deleted";
                formData.set("eventID", deleteEventID || "");
                deleteEventID = null;
                break;
        }

        return handleSubmit(successMsg);
    };

    let eventRoles: { event: Event; roles: Role[] }[];
    $: eventRoles = events.map((event) => {
        let r = roles.filter((r) => r.event == event.id);
        return { event, roles: r };
    });
</script>

<form
    bind:this={deleteEventForm}
    method="POST"
    action="/events?/deleteEvent"
    use:enhance={submit}
></form>

<h3 class="h3">Scheduled Events</h3>
<div class="table-container">
    <form method="POST" use:enhance={submit}>
        <table
            class="table {enableNav ? 'table-interactive' : 'table-hover'} table-fixed text-center"
        >
            <thead>
                {#if !hideEventName}
                    <th>Event</th>
                {/if}
                <th>Date</th>
                <th>Unassigned Roles</th>
                <th>Actions</th>
            </thead>
            <tbody>
                {#each eventRoles as er}
                    <tr
                        on:click={() => {
                            if (enableNav) {
                                goto(`/events/${er.event.id}`);
                            }
                        }}
                    >
                        {#if !hideEventName}
                            <td>{er.event.expand?.event_template.name}</td>
                        {/if}
                        <td>{er.event.datetime}</td>
                        <td
                            >{er.roles.filter((r) => r.assigned_to == "").length}/{er.roles
                                .length}</td
                        >
                        <th>
                            <ActionButton
                                type="delete"
                                title="Delete Event"
                                body={`Are you sure you want to delete this scheduled '${er.event.expand?.event_template.name}' event?'`}
                                form={deleteEventForm}
                                onClick={() => (deleteEventID = er.event.id)}
                            />
                        </th>
                    </tr>
                {/each}
            </tbody>
        </table>
    </form>
</div>
