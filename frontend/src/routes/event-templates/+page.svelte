<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h2 class="h2">Event Templates</h2>
<div class="table-container">
    <table class="table table-interactive table-fixed text-center">
        <thead>
            <th>Name</th>
            <th>Recurrence</th>
            <th>Group</th>
        </thead>
        <tbody>
            {#each data.eventTemplates as eventTemplate}
                <tr on:click={() => goto(`/event-templates/${eventTemplate.id}`)}>
                    <td>{eventTemplate.name}</td>
                    <td>{eventTemplate.recurrence}</td>
                    <td>{eventTemplate.expand?.group.name || ""}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
