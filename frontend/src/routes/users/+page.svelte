<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h2 class="h2">Users</h2>
<div class="table-container">
    <table class="table table-interactive table-fixed text-center">
        <thead>
            <th>User</th>
            <th>Email</th>
            <th>Role</th>
            <th>Organizations</th>
            <th>Groups</th>
            <th>Actions</th>
        </thead>
        <tbody>
            {#each data.users as user}
                <tr on:click={() => goto(`/users/${user.id}`)}>
                    <td>{user.name}</td>
                    <td>{user.email}</td>
                    <td>{user.role}</td>
                    <td>{user.expand?.orgs?.map((org) => org.name).join(", ") || "None"}</td>
                    <td>{user.expand?.groups?.map((group) => group.name).join(", ") || "None"}</td>
                    <td></td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
