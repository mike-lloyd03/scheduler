<script lang="ts">
    import { enhance } from "$app/forms";
    import { goto } from "$app/navigation";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { handleSubmit } from "$lib/utils";
    import type { SubmitFunction } from "@sveltejs/kit";
    import type { PageData } from "./$types";

    export let data: PageData;

    const submit: SubmitFunction = () => {
        return handleSubmit("User created", () => goto("/users/last"));
    };
</script>

<header class="flex-end card variant-filled-surface flex justify-end">
    <ActionButton type="submit" formID="form" />
    <ActionButton type="cancel" onClick={() => goto("/orgs")} />
</header>
<div>
    <h4 class="h4 mt-4">New User</h4>
    <form id="form" method="POST" use:enhance={submit}>
        <div class="py-4">
            <p>
                <span class="font-bold">Name:</span>
                <input class="input inline" type="text" name="name" required />
            </p>
        </div>

        <div class="py-4">
            <p>
                <span class="font-bold">Email:</span>
                <input class="input inline" type="text" name="email" />
            </p>
        </div>

        <div class="py-4">
            <p>
                <span class="font-bold">Password:</span>
                <input class="input inline" type="password" name="password" />
            </p>
        </div>

        <div class="py-4">
            <p>
                <span class="font-bold">Re-enter Password:</span>
                <input class="input inline" type="password" name="passwordConfirm" />
            </p>
        </div>

        <label class="label">
            <span>Role</span>
            <select class="select" name="role" required>
                <option value="member">Member</option>
                <option value="group_admin">Group Admin</option>
                <option value="org_admin">Org Admin</option>
                <option value="super_admin">Super Admin</option>
            </select>
        </label>

        <p>
            <span class="font-bold">Orgs:</span>
            <select class="select" name="orgs" multiple required>
                {#each data.orgs as org}
                    <option value={org.id}>{org.name}</option>
                {/each}
            </select>
        </p>

        <p>
            <span class="font-bold">Groups:</span>
            <select class="select" name="groups" multiple>
                {#each data.groups as group}
                    <option value={group.id}>{group.name}</option>
                {/each}
            </select>
        </p>
    </form>
</div>
