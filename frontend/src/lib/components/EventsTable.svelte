<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import type { Event, Role } from "$lib/types";
    import { handleSubmit } from "$lib/utils";
    import { goto } from "$app/navigation";
    import Delete from "$lib/svg/Delete.svelte";
    import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton";

    const modalStore = getModalStore();

    export let events: Event[];
    export let roles: Role[];
    export let hideEventName = false;
    export let enableNav = false;

    let formElement: HTMLFormElement;

    let editRoleTemplate: string | null;
    let deleteEvent: string | null;
    let newRole: boolean;

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Successful";

        switch (action.search) {
            case "/events?/newEvent":
                newRole = false;
                successMsg = "New event created";
                break;
            case "/events?/updateEvent":
                successMsg = "Event updated";
                formData.set("roleID", editRoleTemplate || "");
                editRoleTemplate = null;
                break;
            case "/events?/deleteEvent":
                console.log("Sending deleteEvent action");
                successMsg = "Event deleted";
                formData.set("eventID", deleteEvent || "");
                deleteEvent = null;
                break;
        }

        return handleSubmit(successMsg);
    };

    function handleDelete(event: Event) {
        deleteEvent = event.id;
        const modal: ModalSettings = {
            type: "confirm",
            title: "Delete Event",
            body: `Are you sure you want to delete this scheduled '${event.expand?.event_template.name}' event?'`,
            response: (r: boolean) => {
                if (r) {
                    formElement.requestSubmit();
                }
            },
        };
        modalStore.trigger(modal);
    }

    let eventRoles: { event: Event; roles: Role[] }[] = events.map((event) => {
        let r = roles.filter((r) => r.event == event.id);
        return { event, roles: r };
    });
</script>

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
                            >{er.roles.filter((r) => r.assigned_to == null).length}/{er.roles
                                .length}</td
                        >
                        <th>
                            <form
                                bind:this={formElement}
                                method="POST"
                                action="/events?/deleteEvent"
                                use:enhance={submit}
                            >
                                <input type="hidden" name="eventID" value={er.event.id} />
                                <button
                                    class="btn hover:variant-ringed-error"
                                    on:click|preventDefault|stopPropagation={() =>
                                        handleDelete(er.event)}><Delete /></button
                                >
                            </form>
                        </th>
                    </tr>
                {/each}
                <tr> </tr>
            </tbody>
        </table>
    </form>
</div>
