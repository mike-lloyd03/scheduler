<script lang="ts">
    import { breadcrumbs } from "$lib/stores";
    import type { PageData } from "./$types";
    import { goto } from "$app/navigation";
    import { toLocaleDate } from "$lib/utils";

    breadcrumbs.clear().add("Dashboard", "dashboard");

    export let data: PageData;
</script>

<h3 class="h3">Upcoming Events</h3>
<div class="table-container">
    <table class="table table-interactive table-fixed text-center">
        <thead>
            <tr>
                <th>Event</th>
                <th>Date</th>
                <th>Assigned Role</th>
            </tr>
        </thead>
        <tbody>
            {#each data.roles as role}
                <tr
                    on:click={() => {
                        goto(`/events/${role.expand?.event?.id}`);
                    }}
                >
                    <td>{role.expand?.event?.expand?.event_template?.name}</td>
                    <td>{toLocaleDate(role.expand?.event?.datetime)}</td>
                    <td>{role.expand?.role_template?.name}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
