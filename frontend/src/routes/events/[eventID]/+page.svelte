<script lang="ts">
    import ActionButton from "$lib/components/ActionButton.svelte";
    import SelectField from "$lib/fields/SelectField.svelte";
    import type { User } from "$lib/types";
    import type { SubmitFunction } from "@sveltejs/kit";
    import type { PageData } from "./$types";
    import { handleSubmit } from "$lib/utils";
    import { enhance } from "$app/forms";
    import { toLocaleDateTime } from "$lib/utils";

    export let data: PageData;
    let editRole: string | undefined = undefined;
    let newRole: boolean;
    let deleteRoleForm: HTMLFormElement;

    let userOptions: { value: string | null; label: string }[];
    userOptions = data.users.map((u: User) => {
        return { value: u.id, label: u.name };
    });
    userOptions.unshift({ value: null, label: "(Unassigned)" });

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Success";
        switch (action.search) {
            case "?/updateRole":
                successMsg = "Role assigned";

                if (!formData.get("assigned_to")) {
                    successMsg = "Role unassigned";
                }

                editRole = undefined;
                break;
            case "?/newRole":
                newRole = false;
                successMsg = "Role added";
                break;
        }

        return handleSubmit(successMsg);
    };
</script>

<div>
    <p>{data.event.expand?.event_template.name}</p>
    <p>{toLocaleDateTime(data.event.datetime)}</p>
</div>

<div>
    <table class="table table-hover text-center">
        <thead>
            <th>Name</th>
            <th>Description</th>
            <th>Assigned to</th>
            <th>Actions</th>
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
                        <form
                            method="POST"
                            id="form-${role.id}"
                            action="?/updateRole"
                            use:enhance={submit}
                        >
                            <input type="hidden" name="roleID" value={role.id} />
                            <SelectField
                                name="assigned_to"
                                value={{
                                    value: role.assigned_to || null,
                                    label: role.expand?.assigned_to?.name ?? "Unassigned",
                                }}
                                options={userOptions}
                                edit={editRole === role.id}
                            />
                        </form>
                    </td>
                    <td>
                        {#if editRole === role.id}
                            <ActionButton type="submit" formID="form-${role.id}" />
                            <ActionButton type="cancel" onClick={() => (editRole = undefined)} />
                        {:else}
                            <form
                                method="POST"
                                action="?/deleteRole"
                                bind:this={deleteRoleForm}
                                use:enhance={submit}
                            >
                                <input type="hidden" name="roleID" value={role.id} />
                                <ActionButton type="edit" onClick={() => (editRole = role.id)} />
                                <ActionButton
                                    type="delete"
                                    title="Remove Role"
                                    body={`Are you sure you want to remove the role '${role.expand?.role_template.name}'?`}
                                    form={deleteRoleForm}
                                />
                            </form>
                        {/if}
                    </td>
                </tr>
            {/each}
            {#if newRole}
                <tr>
                    <td>
                        <form
                            method="POST"
                            id="newRoleForm"
                            action="?/newRole"
                            use:enhance={submit}
                        >
                            <select class="select" name="role_template">
                                {#each data.roleTemplates as rt}
                                    <option
                                        value={rt.id}
                                        disabled={data.roles
                                            .map((r) => r.role_template)
                                            .includes(rt.id)}>{rt.name}</option
                                    >
                                {/each}
                            </select>
                            <input type="hidden" name="event" value={data.event.id} />
                        </form>
                    </td>
                    <td>-</td>
                    <td>
                        <select class="select" name="assigned_to" form="newRoleForm">
                            {#each userOptions as user}
                                <option value={user.value}>{user.label}</option>
                            {/each}
                        </select>
                    </td>
                    <td>
                        <ActionButton type="submit" formID="newRoleForm" />
                        <ActionButton type="cancel" onClick={() => (newRole = false)} />
                    </td>
                </tr>
            {:else if data.roles.length != data.roleTemplates.length}
                <tr>
                    <td></td>
                    <td></td>
                    <td></td>
                    <td><ActionButton type="new" onClick={() => (newRole = true)} /></td>
                </tr>
            {/if}
        </tbody>
    </table>
</div>
