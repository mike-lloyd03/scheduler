<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import TextAreaField from "$lib/fields/TextAreaField.svelte";
    import type { RoleTemplate } from "$lib/types";
    import ActionButton from "$lib/components/ActionButton.svelte";

    interface Props {
        roleTemplate: RoleTemplate;
        editRoleTemplate: string | undefined;
        deleteRoleTemplate: string | undefined;
        showEdit: boolean;
        showDelete: boolean;
    }

    let {
        roleTemplate,
        editRoleTemplate = $bindable(),
        deleteRoleTemplate = $bindable(),
        showEdit = false,
        showDelete = false,
    }: Props = $props();

    let edit = $derived(roleTemplate.id == editRoleTemplate);
</script>

<tr>
    <td>
        <InputField name="name" value={roleTemplate.name} {edit} />
    </td>
    <td>
        <TextAreaField name="description" value={roleTemplate.description || ""} {edit} />
    </td>
    {#if showEdit || showDelete}
        <td>
            {#if edit}
                <ActionButton type="submit" formaction="?/updateRole" />
                <ActionButton type="cancel" onClick={() => (editRoleTemplate = undefined)} />
            {:else}
                {#if showEdit}
                    <ActionButton
                        type="edit"
                        onClick={() => (editRoleTemplate = roleTemplate.id)}
                    />
                {/if}

                {#if showDelete}
                    <ActionButton
                        type="delete"
                        title="Delete Role Template"
                        body={`Are you sure you want to delete role template '${roleTemplate.name}'?`}
                        formID="roleForm"
                        onClick={() => (deleteRoleTemplate = roleTemplate.id)}
                        formaction="?/deleteRole"
                    />
                {/if}
            {/if}
        </td>
    {/if}
</tr>
