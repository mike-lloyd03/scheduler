<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import RoleRow from "./RoleRow.svelte";
    import type { RoleTemplate } from "$lib/types";
    import Check from "$lib/svg/Check.svelte";
    import Cancel from "$lib/svg/Cancel.svelte";
    import Add from "$lib/svg/Add.svelte";
    import { handleSubmit } from "$lib/utils";

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
        }

        return handleSubmit(successMsg);
    };
</script>

<h3 class="h3">Event Roles</h3>
<div class="table-container">
    <form method="POST" use:enhance={submit}>
        <table class="table table-hover table-fixed text-center">
            <thead>
                <th>Role</th>
                <th>Description</th>
                <th>Actions</th>
            </thead>
            <tbody>
                {#each roleTemplates as roleTemplate}
                    <RoleRow {roleTemplate} bind:editRoleTemplate />
                {/each}
                <tr>
                    {#if newRole}
                        <td>
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
                            <button class="btn hover:variant-ringed-success" formaction="?/newRole"
                                ><Check /></button
                            >
                            <button
                                class="btn hover:variant-ringed-secondary"
                                on:click|preventDefault={() => {
                                    newRole = false;
                                }}><Cancel /></button
                            >
                        </td>
                    {:else}
                        <td></td>
                        <td></td>
                        <td>
                            <button
                                class="btn hover:variant-ringed-primary"
                                on:click|preventDefault={() => (newRole = true)}><Add /></button
                            >
                        </td>
                    {/if}
                </tr>
            </tbody>
        </table>
    </form>
</div>
