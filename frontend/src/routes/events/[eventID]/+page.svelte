<script lang="ts">
    import { enhance } from "$app/forms";
    import toast from "svelte-french-toast";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/InputField.svelte";
    import SelectField from "$lib/SelectField.svelte";
    import type { PageData } from "./$types";
    import TextAreaField from "$lib/TextAreaField.svelte";

    export let data: PageData;

    var editEventTemplate = false;
    var editRoleTemplate: string | null = null;
    var deleteRoleTemplate: string | null = null;
    var newRole = false;

    const submit: SubmitFunction = ({ action, formData }) => {
        var successMsg = "Successful";

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
                <InputField
                    title="Name"
                    value={data.event_template.name}
                    edit={editEventTemplate}
                />
            </div>
            <div>
                <SelectField
                    title="Recurrence"
                    value={data.event_template.recurrence}
                    options={["daily", "weekly", "monthly"]}
                    edit={editEventTemplate}
                />
            </div>
            <div>
                <span class="font-bold">Group:</span>
                {data.event_template.expand?.group.name || ""}
            </div>
            <div>
                <span class="font-bold">Created at:</span>
                {data.event_template.created}
            </div>
            <div>
                <span class="font-bold">Updated at:</span>
                {data.event_template.updated}
            </div>
        </div>
        {#if editEventTemplate}
            <button type="submit" class="variant-filled-primary btn">Save</button>
            <button
                type="button"
                class="variant-filled-secondary btn"
                on:click|preventDefault={() => (editEventTemplate = false)}
            >
                Cancel
            </button>
        {:else}
            <button
                type="submit"
                class="variant-filled-primary btn"
                on:click|preventDefault={() => (editEventTemplate = true)}>Edit</button
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
                        {#each data.role_templates as role_template}
                            <tr>
                                <td>
                                    <InputField
                                        title="name"
                                        value={role_template.name}
                                        edit={role_template.id == editRoleTemplate}
                                        hideTitle
                                    />
                                </td>
                                <td>
                                    <TextAreaField
                                        title="description"
                                        value={role_template.description}
                                        edit={role_template.id == editRoleTemplate}
                                        hideTitle
                                    />
                                </td>
                                <td>
                                    {#if role_template.id == editRoleTemplate}
                                        <button
                                            class="variant-filled-primary btn"
                                            formaction="?/updateRole">Save</button
                                        >
                                        <button
                                            class="variant-filled-secondary btn"
                                            on:click|preventDefault={() =>
                                                (editRoleTemplate = null)}>Cancel</button
                                        >
                                    {:else}
                                        <button
                                            class="variant-filled-primary btn"
                                            on:click|preventDefault={() =>
                                                (editRoleTemplate = role_template.id)}>Edit</button
                                        >
                                        <button
                                            class="variant-filled-secondary btn"
                                            on:click={() => (deleteRoleTemplate = role_template.id)}
                                            formaction="?/deleteRole">Delete</button
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
                                        class="variant-filled-primary btn"
                                        formaction="?/newRole">Save</button
                                    >
                                    <button
                                        class="variant-filled-secondary btn"
                                        on:click|preventDefault={() => {
                                            newRole = false;
                                        }}>Cancel</button
                                    >
                                </td>
                            {:else}
                                <td></td>
                                <td></td>
                                <td>
                                    <button
                                        class="variant-filled-primary btn"
                                        on:click|preventDefault={() => (newRole = true)}>+</button
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
