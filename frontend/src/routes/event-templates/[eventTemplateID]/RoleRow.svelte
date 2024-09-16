<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import TextAreaField from "$lib/fields/TextAreaField.svelte";
    import type { RoleTemplate } from "$lib/types";
    import ActionButton from "$lib/components/ActionButton.svelte";

    export let roleTemplate: RoleTemplate;
    export let editRoleTemplate: string | null;
    export let deleteRoleTemplate: string | null;
    export let formElement: HTMLFormElement;

    $: edit = roleTemplate.id == editRoleTemplate;
</script>

<tr>
    <td>
        <InputField name="name" value={roleTemplate.name} {edit} />
    </td>
    <td>
        <TextAreaField name="description" value={roleTemplate.description || ""} {edit} />
    </td>
    <td>
        {#if edit}
            <ActionButton type="submit" formaction="?/updateRole" />
            <ActionButton type="cancel" onClick={() => (editRoleTemplate = null)} />
        {:else}
            <ActionButton type="edit" onClick={() => (editRoleTemplate = roleTemplate.id)} />
            <ActionButton
                type="delete"
                title="Delete Role Template"
                body={`Are you sure you want to delete role template '${roleTemplate.name}'?`}
                form={formElement}
                onClick={() => (deleteRoleTemplate = roleTemplate.id)}
                formaction="?/deleteRole"
            />
        {/if}
    </td>
</tr>
