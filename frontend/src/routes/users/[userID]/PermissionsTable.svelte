<script lang="ts">
    import type { Permission } from "$lib/types";
    export let permissions: Permission[];
</script>

<h3 class="h3">Permissions</h3>
<div class="table-container">
    <table class="table table-hover table-fixed text-center">
        <thead>
            <tr>
                <th>Unit</th>
                <th>Role</th>
            </tr>
        </thead>
        <tbody>
            {#each permissions as p}
                <tr>
                    <td
                        >{p.org != ""
                            ? "Org: " + p.expand?.org?.name
                            : "Group: " + p.expand?.group?.name}</td
                    >
                    <td>{p.role}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
