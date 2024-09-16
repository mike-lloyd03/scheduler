<script lang="ts">
    import ActionButton from "$lib/components/ActionButton.svelte";
    import SelectField from "$lib/fields/SelectField.svelte";
    import type { User } from "$lib/types";
    import type { SubmitFunction } from "@sveltejs/kit";
    import type { PageData } from "./$types";
    import { handleSubmit } from "$lib/utils";
    import { enhance } from "$app/forms";

    export let data: PageData;
    let editRole: string | undefined = undefined;

    let userOptions = data.users.map((u: User) => {
        return { value: u.id, label: u.name };
    });
    userOptions.unshift({ value: null, label: "(Unassigned)" });

    const submit: SubmitFunction = ({ formData }) => {
        let successMsg = "Role assigned";

        if (!formData.get("assigned_to")) {
            successMsg = "Role unassigned";
        }

        editRole = undefined;

        return handleSubmit(successMsg);
    };
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
                        <form method="POST" id="form-${role.id}" use:enhance={submit}>
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
                            <ActionButton type="edit" onClick={() => (editRole = role.id)} />
                        {/if}
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
