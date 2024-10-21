<script lang="ts">
    import type { SvelteComponent } from "svelte";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import type { EventTemplate, RoleTemplate, User } from "$lib/types";

    export let parent: SvelteComponent;

    const modalStore = getModalStore();

    const formData = {
        event_template: $modalStore[0].meta.selectedEventTemplateID || "",
        datetime: "",
        roles: "",
    };

    function onFormSubmit(): void {
        if ($modalStore[0].response) {
            $modalStore[0].response(formData);
        }
        modalStore.close();
    }

    let eventTemplatesAndRoleTemplates: {
        eventTemplate: EventTemplate;
        roleTemplates: RoleTemplate[];
    }[];
    let selectedEventRoleTemplates: RoleTemplate[] | undefined;
    let users: User[];

    let createRoles: Record<string, string | null> = {};

    $: if ($modalStore[0]) {
        eventTemplatesAndRoleTemplates = $modalStore[0].meta.eventTemplates.map(
            (et: EventTemplate) => {
                let roleTemplates = $modalStore[0].meta.roleTemplates.filter(
                    (rt: RoleTemplate) => rt.event_template == et.id,
                );
                return { eventTemplate: et, roleTemplates };
            },
        );
        users = $modalStore[0].meta.users;
    }

    $: selectedEventTemplate = eventTemplatesAndRoleTemplates.find(
        (etrt) => etrt.eventTemplate.id == formData.event_template,
    );
    $: selectedEventRoleTemplates = selectedEventTemplate?.roleTemplates;
    $: groupUsers = users.filter((u) =>
        u.groups.includes(selectedEventTemplate?.eventTemplate.group || ""),
    );

    function selectTemplate() {
        let eventRoles = {};
        selectedEventRoleTemplates?.forEach((rt) => (eventRoles[rt.id] = null));
        formData.roles = JSON.stringify(eventRoles);
    }

    function updateCreateRolesValue() {
        formData.roles = JSON.stringify(createRoles);
    }

    const cBase = "card p-4 w-modal shadow-xl space-y-4";
    const cHeader = "text-2xl font-bold";
    const cForm = "border border-surface-500 p-4 space-y-4 rounded-container-token";
</script>

{#if $modalStore[0]}
    <div class={cBase}>
        <header class={cHeader}>Schedule Event</header>

        <form class="modal-form {cForm}">
            <label class="label">
                <span>Event Template</span>
                <select
                    class="select"
                    bind:value={formData.event_template}
                    on:change={selectTemplate}
                >
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
                                {#each groupUsers as user}
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
