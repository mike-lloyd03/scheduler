<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import RoleRow from "./RoleRow.svelte";
    import type { RoleTemplate } from "$lib/types";
    import { handleSubmit } from "$lib/utils";
    import ActionButton from "$lib/components/ActionButton.svelte";

    interface Props {
        roleTemplates: RoleTemplate[];
        showNew: boolean;
        showEdit: boolean;
        showDelete: boolean;
    }

    let { roleTemplates, showNew = false, showEdit = false, showDelete = false }: Props = $props();

    let editRoleTemplate: string | undefined = $state();
    let deleteRoleTemplate: string | undefined = $state();
    let newRole = $state(false);

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
                editRoleTemplate = undefined;
                break;
            case "?/deleteRole":
                successMsg = "Role deleted";
                formData.set("roleID", deleteRoleTemplate || "");
                deleteRoleTemplate = undefined;
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
                <tr>
                    <th>Role</th>
                    <th>Description</th>
                    {#if showNew || showEdit || showDelete}
                        <th>Actions</th>
                    {/if}
                </tr>
            </thead>
            <tbody>
                {#each roleTemplates as roleTemplate}
                    <RoleRow
                        {roleTemplate}
                        bind:editRoleTemplate
                        bind:deleteRoleTemplate
                        {showEdit}
                        {showDelete}
                    />
                {/each}
                {#if showNew}
                    <tr>
                        {#if newRole}
                            <td>
                                <!-- svelte-ignore a11y_autofocus -->
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
                                ></textarea>
                            </td>
                            {#if showNew}
                                <td>
                                    <ActionButton type="submit" formaction="?/newRole" />
                                    <ActionButton type="cancel" onClick={() => (newRole = false)} />
                                </td>
                            {/if}
                        {:else}
                            <td></td>
                            <td></td>
                            <td>
                                <ActionButton type="new" onClick={() => (newRole = true)} />
                            </td>
                        {/if}
                    </tr>
                {/if}
            </tbody>
        </table>
    </form>
</div>
