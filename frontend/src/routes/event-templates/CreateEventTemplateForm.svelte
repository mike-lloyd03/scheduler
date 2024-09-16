<script lang="ts">
    import type { SvelteComponent } from "svelte";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import type { Group } from "$lib/types";

    export let parent: SvelteComponent;

    const modalStore = getModalStore();

    let formData = {
        name: "",
        recurrence: "",
        group: "",
    };

    function onFormSubmit(): void {
        if ($modalStore[0].response) {
            $modalStore[0].response(formData);
        }
        modalStore.close();
    }

    let groups: Group[];
    if ($modalStore[0]) {
        groups = $modalStore[0].meta.groups;
    }

    const cBase = "card p-4 w-modal shadow-xl space-y-4";
    const cHeader = "text-2xl font-bold";
    const cForm = "border border-surface-500 p-4 space-y-4 rounded-container-token";
</script>

{#if $modalStore[0]}
    <div class="modal-example-form {cBase}">
        <header class={cHeader}>Create Event Template</header>

        <form class="modal-form {cForm}">
            <label class="label">
                <span>Name</span>
                <input class="input" type="text" bind:value={formData.name} />
            </label>

            <label class="label">
                <span>Recurrence</span>
                <select class="select" bind:value={formData.recurrence}>
                    <option value="daily">Daily</option>
                    <option value="weekly">Weekly</option>
                    <option value="monthly">Monthly</option>
                </select>
            </label>

            <label class="label">
                <span>Group</span>
                <select class="select" bind:value={formData.group}>
                    {#each groups as group}
                        <option value={group.id}>{group.name}</option>
                    {/each}
                </select>
            </label>
        </form>

        <footer class="modal-footer {parent.regionFooter}">
            <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>
                {parent.buttonTextCancel}
            </button>

            <button class="btn {parent.buttonPositive}" on:click={onFormSubmit}>
                {parent.buttonTextSubmit}
            </button>
        </footer>
    </div>
{/if}
