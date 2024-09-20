<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { handleSubmit } from "$lib/utils";
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;

    const submit: SubmitFunction = () => {
        return handleSubmit("Group created", () => goto("/groups/last"));
    };
</script>

<header class="flex-end card variant-filled-surface flex justify-end">
    <ActionButton type="submit" formID="form" />
    <ActionButton type="cancel" onClick={() => goto("/orgs")} />
</header>
<div>
    <h4 class="h4 mt-4">Create Group</h4>
    <form id="form" method="POST" use:enhance={submit}>
        <div class="py-4">
            <p>
                <span class="font-bold">Name:</span>
                <input class="input inline" type="text" name="name" />
            </p>

            <p>
                <span class="font-bold">Org:</span>
                <select class="select" name="org">
                    {#each data.orgs as org}
                        <option value={org.id}>{org.name}</option>
                    {/each}
                </select>
            </p>
        </div>
    </form>
</div>
