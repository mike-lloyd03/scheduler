<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h2 class="h2">Organizations</h2>
<div class="table-container">
    <table class="table table-hover table-fixed text-center">
        <thead>
            <th>Group</th>
            <th>Created At</th>
        </thead>
        <tbody>
            {#each data.orgs as org}
                <tr class="cursor-pointer" on:click={() => goto(`/orgs/${org.id}`)}>
                    <td>{org.name}</td>
                    <td>{org.created}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
