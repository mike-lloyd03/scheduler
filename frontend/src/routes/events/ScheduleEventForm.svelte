<script lang="ts">
    import type { SvelteComponent } from "svelte";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import type { EventTemplate, RoleTemplate, User } from "$lib/types";

    export let parent: SvelteComponent;

    const modalStore = getModalStore();

    const formData = {
        event_template: "",
        datetime: "",
        roles: "",
    };

    function onFormSubmit(): void {
        if ($modalStore[0].response) {
            $modalStore[0].response(formData);
        }
        modalStore.close();
    }

    const eventTemplatesAndRoleTemplates: {
        eventTemplate: EventTemplate;
        roleTemplates: RoleTemplate[];
    }[] = $modalStore[0].meta.eventTemplates.map((et: EventTemplate) => {
        let roleTemplates = $modalStore[0].meta.roleTemplates.filter(
            (rt: RoleTemplate) => (rt.event_template = et.id),
        );
        return { eventTemplate: et, roleTemplates };
    });

    let selectedEventRoleTemplates: RoleTemplate[] | undefined;
    let users: User[];

    $: if ($modalStore[0]) {
        selectedEventRoleTemplates = eventTemplatesAndRoleTemplates.find(
            (etrt) => etrt.eventTemplate.id == formData.event_template,
        )?.roleTemplates;
        users = $modalStore[0].meta.users;
    }

    let createRoles: Record<string, string | null> = {};
    function updateCreateRolesValue() {
        formData.roles = JSON.stringify(createRoles);
    }

    const cBase = "card p-4 w-modal shadow-xl space-y-4";
    const cHeader = "text-2xl font-bold";
    const cForm = "border border-surface-500 p-4 space-y-4 rounded-container-token";
</script>

{#if $modalStore[0]}
    <div class="modal-example-form {cBase}">
        <header class={cHeader}>Schedule Event</header>

        <form class="modal-form {cForm}">
            <label class="label">
                <span>Event Template</span>
                <select class="select" bind:value={formData.event_template}>
                    {#each $modalStore[0].meta.eventTemplates as et}
                        <option value={et.id}>{et.name}</option>
                    {/each}
                </select>
            </label>

            <label class="label">
                <span>Date</span>
                <input class="input" type="date" bind:value={formData.datetime} />
            </label>

            {#if selectedEventRoleTemplates}
                <label class="label">
                    <span>Add Roles</span>
                    {#each selectedEventRoleTemplates as rt}
                        <label class="flex items-center space-x-2">
                            <p>{rt.name}</p>
                            <select
                                class="select"
                                bind:value={createRoles[rt.id]}
                                on:change={updateCreateRolesValue}
                            >
                                <option value={null}>(Unassigned)</option>
                                {#each users as user}
                                    <option value={user.id}>{user.name}</option>
                                {/each}
                            </select>
                        </label>
                    {/each}
                </label>
            {/if}
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
