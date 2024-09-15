<script lang="ts">
    import Delete from "../svg/Delete.svelte";
    import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton";

    export let title: string;
    export let body: string;
    export let form: HTMLFormElement;
    export let presubmit: (() => void) | undefined = undefined;
    export let action: string | undefined = undefined;

    const modalStore = getModalStore();

    function handleDelete() {
        const modal: ModalSettings = {
            type: "confirm",
            title,
            body,
            response: (r: boolean) => {
                if (r) {
                    if (presubmit) {
                        presubmit();
                    }
                    if (action) {
                        form.action = action;
                    }
                    form.requestSubmit();
                }
            },
        };
        modalStore.trigger(modal);
    }
</script>

<button
    class="btn hover:variant-ringed-error"
    on:click|preventDefault|stopPropagation={handleDelete}
>
    <Delete />
</button>
