<script lang="ts">
    import Check from "../svg/Check.svelte";
    import Cancel from "../svg/Cancel.svelte";
    import Edit from "../svg/Edit.svelte";
    import Delete from "../svg/Delete.svelte";
    import { getModalStore, type ModalSettings, type ModalStore } from "@skeletonlabs/skeleton";
    import Add from "$lib/svg/Add.svelte";

    export let type: "new" | "submit" | "edit" | "cancel" | "delete";

    export let onClick: (() => void) | undefined = undefined;
    export let formaction: string | undefined = undefined;
    export let title: string | undefined = undefined;
    export let body: string | undefined = undefined;
    export let formID: string | undefined = undefined;

    let modalStore: ModalStore;

    if (type === "delete") {
        modalStore = getModalStore();
    }

    function handleDelete() {
        const modal: ModalSettings = {
            type: "confirm",
            title,
            body,
            response: (r: boolean) => {
                if (r) {
                    if (onClick) {
                        onClick();
                    }
                    if (formID) {
                        const form: HTMLFormElement | null = document.getElementById(
                            formID,
                        ) as HTMLFormElement | null;
                        if (form) {
                            if (formaction) {
                                form.action = formaction;
                            }
                            form.requestSubmit();
                        }
                    }
                }
            },
        };
        modalStore.trigger(modal);
    }
</script>

{#if type === "new" && onClick}
    <button class="btn hover:variant-ringed-primary" on:click|preventDefault={onClick}
        ><Add /></button
    >
{:else if type === "submit"}
    <button class="btn hover:variant-ringed-success" {formaction} form={formID}>
        <Check />
    </button>
{:else if type === "edit" && onClick}
    <button class="btn hover:variant-ringed-primary" on:click|preventDefault={onClick}>
        <Edit />
    </button>
{:else if type === "cancel" && onClick}
    <button class="btn hover:variant-ringed-secondary" on:click|preventDefault={onClick}>
        <Cancel />
    </button>
{:else if type === "delete" && title && body && formID}
    <button
        class="btn hover:variant-ringed-error"
        on:click|preventDefault|stopPropagation={handleDelete}
    >
        <Delete />
    </button>
{/if}
