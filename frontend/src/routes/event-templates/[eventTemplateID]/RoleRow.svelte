<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import Cancel from "$lib/svg/Cancel.svelte";
    import Check from "$lib/svg/Check.svelte";
    import Delete from "$lib/svg/Delete.svelte";
    import Edit from "$lib/svg/Edit.svelte";
    import TextAreaField from "$lib/fields/TextAreaField.svelte";
    import type { RoleTemplate } from "$lib/types";
    import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton";

    const modalStore = getModalStore();

    export let roleTemplate: RoleTemplate;
    export let editRoleTemplate: string | null;
    export let deleteRoleTemplate: string | null;
    export let formElement: HTMLFormElement;

    $: edit = roleTemplate.id == editRoleTemplate;

    function handleDelete() {
        deleteRoleTemplate = roleTemplate.id;
        const modal: ModalSettings = {
            type: "confirm",
            title: "Delete Role Template",
            body: `Are you sure you want to delete role template '${roleTemplate.name}'?`,
            response: (r: boolean) => {
                if (r) {
                    formElement.requestSubmit();
                }
            },
        };
        modalStore.trigger(modal);
    }
</script>

<tr>
    <td>
        <InputField title="name" value={roleTemplate.name} {edit} hideTitle />
    </td>
    <td>
        <TextAreaField
            title="description"
            value={roleTemplate.description || ""}
            {edit}
            hideTitle
        />
    </td>
    <td>
        {#if edit}
            <button class="btn hover:variant-ringed-success" formaction="?/updateRole"
                ><Check /></button
            >
            <button
                class="btn hover:variant-ringed-secondary"
                on:click|preventDefault={() => (editRoleTemplate = null)}
            >
                <Cancel />
            </button>
        {:else}
            <button
                class="btn hover:variant-ringed-primary"
                on:click|preventDefault={() => (editRoleTemplate = roleTemplate.id)}
                ><Edit /></button
            >
            <button
                class="btn hover:variant-ringed-error"
                formaction="?/deleteRole"
                on:click|preventDefault={handleDelete}><Delete /></button
            >
        {/if}
    </td>
</tr>
