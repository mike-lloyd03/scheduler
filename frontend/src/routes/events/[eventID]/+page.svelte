<script lang="ts">
    import { enhance } from "$app/forms";
    import toast from "svelte-french-toast";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/InputField.svelte";
    import SelectField from "$lib/SelectField.svelte";
    import type { PageData } from "./$types";
    import TextAreaField from "$lib/TextAreaField.svelte";
    import Edit from "$lib/svg/Edit.svelte";
    import Delete from "$lib/svg/Delete.svelte";
    import Add from "$lib/svg/Add.svelte";
    import Check from "$lib/svg/Check.svelte";
    import Cancel from "$lib/svg/Cancel.svelte";

    export let data: PageData;

    let editEventTemplate = false;
    let editRoleTemplate: string | null = null;
    let deleteRoleTemplate: string | null = null;
    let newRole = false;

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Successful";

        switch (action.search) {
            case "?/updateEvent":
                editEventTemplate = false;
                successMsg = "Event Template updated";
                break;
            case "?/newRole":
                newRole = false;
                successMsg = "New Role created";
                break;
            case "?/deleteRole":
                successMsg = "Role deleted";
                formData.set("roleID", deleteRoleTemplate || "");
                break;
            case "?/updateRole":
                successMsg = "Role updated";
                formData.set("roleID", editRoleTemplate || "");
                editRoleTemplate = null;
                break;
        }

        return async ({ result, update }) => {
            switch (result.type) {
                case "success":
                    toast.success(successMsg);
                    await update();
                    break;
                case "failure":
                    toast.error(result.data?.message || "Something went wrong");
                    break;
                default:
                    await update();
            }
        };
    };
</script>

<div>
    <form method="POST" use:enhance={submit} action="?/updateEvent">
        <div class="py-4">
            <div>
                <InputField title="Name" value={data.eventTemplate.name} edit={editEventTemplate} />
            </div>
            <div>
                <SelectField
                    title="Recurrence"
                    value={data.eventTemplate.recurrence}
                    options={["daily", "weekly", "monthly"]}
                    edit={editEventTemplate}
                />
            </div>
            <div>
                <span class="font-bold">Group:</span>
                {data.eventTemplate.expand?.group.name || ""}
            </div>
            <div>
                <span class="font-bold">Created at:</span>
                {data.eventTemplate.created}
            </div>
            <div>
                <span class="font-bold">Updated at:</span>
                {data.eventTemplate.updated}
            </div>
        </div>
        {#if editEventTemplate}
            <button type="submit" class="variant-filled-success btn"><Check /></button>
            <button
                type="button"
                class="variant-filled-secondary btn"
                on:click|preventDefault={() => (editEventTemplate = false)}
            >
                <Cancel />
            </button>
        {:else}
            <button
                type="submit"
                class="variant-filled-primary btn"
                on:click|preventDefault={() => (editEventTemplate = true)}><Edit /></button
            >
        {/if}
    </form>

    <div class="py-8">
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
                        {#each data.roleTemplates as roleTemplate}
                            <tr>
                                <td>
                                    <InputField
                                        title="name"
                                        value={roleTemplate.name}
                                        edit={roleTemplate.id == editRoleTemplate}
                                        hideTitle
                                    />
                                </td>
                                <td>
                                    <TextAreaField
                                        title="description"
                                        value={roleTemplate.description || ""}
                                        edit={roleTemplate.id == editRoleTemplate}
                                        hideTitle
                                    />
                                </td>
                                <td>
                                    {#if roleTemplate.id == editRoleTemplate}
                                        <button
                                            class="variant-filled-success btn"
                                            formaction="?/updateRole"><Check /></button
                                        >
                                        <button
                                            class="variant-filled-secondary btn"
                                            on:click|preventDefault={() =>
                                                (editRoleTemplate = null)}><Cancel /></button
                                        >
                                    {:else}
                                        <button
                                            class="variant-filled-primary btn"
                                            on:click|preventDefault={() =>
                                                (editRoleTemplate = roleTemplate.id)}
                                            ><Edit /></button
                                        >
                                        <button
                                            class="variant-filled-error btn"
                                            on:click={() => (deleteRoleTemplate = roleTemplate.id)}
                                            formaction="?/deleteRole"><Delete /></button
                                        >
                                    {/if}
                                </td>
                            </tr>
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
                                    <button
                                        class="variant-filled-success btn"
                                        formaction="?/newRole"><Check /></button
                                    >
                                    <button
                                        class="variant-filled-secondary btn"
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
                                        class="variant-filled-primary btn"
                                        on:click|preventDefault={() => (newRole = true)}
                                        ><Add /></button
                                    >
                                </td>
                            {/if}
                        </tr>
                    </tbody>
                </table>
            </form>
        </div>
    </div>
</div>
