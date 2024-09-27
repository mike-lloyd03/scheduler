<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ResourcePage from "$lib/components/ResourcePage.svelte";
    import { toLocaleDateTime } from "$lib/utils";
    import { breadcrumbs } from "$lib/stores";

    export let data: PageData;

    $: breadcrumbs.clear().add("Users", "users").add(data.user.name, data.user.id);

    let edit = false;
</script>

<ResourcePage
    bind:edit
    title="User"
    baseURL="/users"
    deleteBody="Are you sure you want to delete the user '{data.user.name}'?"
>
    <div class="py-4">
        <p>
            <span class="font-bold">Name:</span>
            <InputField name="name" value={data.user.name} {edit} />
        </p>

        <p>
            <span class="font-bold">Email:</span>
            <InputField name="email" value={data.user.email} {edit} />
        </p>

        <p><span class="font-bold">Created at:</span> {toLocaleDateTime(data.user.created)}</p>

        <p><span class="font-bold">Updated at:</span> {toLocaleDateTime(data.user.updated)}</p>
    </div>
</ResourcePage>
