<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<div>
    <p>{data.event.expand?.event_template.name}</p>
    <p>{data.event.datetime}</p>
</div>

<div>
    <table class="table table-hover text-center">
        <thead>
            <th>Name</th>
            <th>Description</th>
            <th>Assigned to</th>
        </thead>
        <tbody>
            {#each data.roles as role}
                <tr>
                    <td>
                        {role.expand?.role_template.name}
                    </td>
                    <td>
                        {role.expand?.role_template.description}
                    </td>
                    <td>
                        {role.expand?.assigned_to?.name ?? "Unassigned"}
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
