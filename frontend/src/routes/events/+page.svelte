<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h2 class="h2">Events</h2>
<div class="table-container">
    <table class="table table-hover">
        <thead>
            <th>Name</th>
            <th>Recurrence</th>
            <th>Group</th>
        </thead>
        <tbody>
            {#each data.event_templates as event_template}
                <tr class="cursor-pointer" on:click={() => goto(`/events/${event_template.id}`)}>
                    <td>{event_template.name}</td>
                    <td>{event_template.recurrence}</td>
                    <td>{event_template.expand?.group.name || ""}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
