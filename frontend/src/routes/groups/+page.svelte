<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h2 class="h2">Groups</h2>
<div class="table-container">
    <table class="table table-hover table-fixed text-center">
        <thead>
            <th>Group</th>
            <th>Organization</th>
        </thead>
        <tbody>
            {#each data.groups as group}
                <tr class="cursor-pointer" on:click={() => goto(`/groups/${group.id}`)}>
                    <td>{group.name}</td>
                    <td>{group.expand?.org.name}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
