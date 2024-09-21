<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { handleSubmit } from "$lib/utils";
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;

    const submit: SubmitFunction = () => {
        return handleSubmit("Event template created", () => goto("/event-templates/last"));
    };
</script>

<header class="flex-end card variant-filled-surface flex justify-end">
    <ActionButton type="submit" formID="form" />
    <ActionButton type="cancel" onClick={() => goto("/orgs")} />
</header>
<div>
    <h4 class="h4 mt-4">Create Event Template</h4>
    <form id="form" method="POST" use:enhance={submit}>
        <div class="py-4">
            <label class="label">
                <span>Name</span>
                <input class="input" type="text" name="name" required />
            </label>

            <label class="label">
                <span>Recurrence</span>
                <select class="select" name="recurrence" required>
                    <option value={null} disabled selected>Select...</option>
                    <option value="daily">Daily</option>
                    <option value="weekly">Weekly</option>
                    <option value="monthly">Monthly</option>
                </select>
            </label>

            <label class="label">
                <span>Group</span>
                <select class="select" name="group" required>
                    <option value={null} disabled selected>Select...</option>
                    {#each data.groups as group}
                        <option value={group.id}>{group.name}</option>
                    {/each}
                </select>
            </label>
        </div>
    </form>
</div>
