<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import RoleRow from "./RoleRow.svelte";
    import type { RoleTemplate } from "$lib/types";
    import { handleSubmit } from "$lib/utils";
    import ActionButton from "$lib/components/ActionButton.svelte";

    export let roleTemplates: RoleTemplate[];

    let editRoleTemplate: string | null;
    let deleteRoleTemplate: string | null;
    let newRole: boolean;

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Successful";

        switch (action.search) {
            case "?/newRole":
                newRole = false;
                successMsg = "New Role created";
                break;
            case "?/updateRole":
                successMsg = "Role updated";
                formData.set("roleID", editRoleTemplate || "");
                editRoleTemplate = null;
                break;
            case "?/deleteRole":
                successMsg = "Role deleted";
                formData.set("roleID", deleteRoleTemplate || "");
                deleteRoleTemplate = null;
                break;
        }

        return handleSubmit(successMsg);
    };
</script>

<h3 class="h3">Event Roles</h3>
<div class="table-container">
    <form id="roleForm" method="POST" use:enhance={submit}>
        <table class="table table-hover table-fixed text-center">
            <thead>
                <th>Role</th>
                <th>Description</th>
                <th>Actions</th>
            </thead>
            <tbody>
                {#each roleTemplates as roleTemplate}
                    <RoleRow {roleTemplate} bind:editRoleTemplate bind:deleteRoleTemplate />
                {/each}
                <tr>
                    {#if newRole}
                        <td>
                            <!-- svelte-ignore a11y-autofocus -->
                            <input
                                class="input"
                                type="text"
                                name="name"
                                id="name"
                                placeholder="Role Name"
                                required
                                autofocus
                            />
                        </td>
                        <td>
                            <textarea
                                class="textarea"
                                name="description"
                                id="description"
                                placeholder="Description"
                            />
                        </td>
                        <td>
                            <ActionButton type="submit" formaction="?/newRole" />
                            <ActionButton type="cancel" onClick={() => (newRole = false)} />
                        </td>
                    {:else}
                        <td></td>
                        <td></td>
                        <td>
                            <ActionButton type="new" onClick={() => (newRole = true)} />
                        </td>
                    {/if}
                </tr>
            </tbody>
        </table>
    </form>
</div>
