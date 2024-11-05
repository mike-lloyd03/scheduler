<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ResourcePage from "$lib/components/ResourcePage.svelte";
    import { toLocaleDateTime } from "$lib/utils";
    import { breadcrumbs } from "$lib/stores";
    import { UserRole } from "$lib/types";

    export let data: PageData;

    $: breadcrumbs.clear().add("Orgs", "orgs").add(data.org.name, data.org.id);

    let edit = false;
</script>

<ResourcePage
    bind:edit
    title="Organization"
    baseURL="/orgs"
    deleteBody="Are you sure you want to delete the organization '{data.org
        .name}'? This will delete all groups, events, and other resources associated with it."
    showEdit={data.currentUserPermissions
        ?.filter((p) => p.org == data.org.id)
        .some((p) => p.role == UserRole.Admin)}
>
    <p>
        <span class="font-bold">Name:</span>
        <InputField name="name" value={data.org.name} {edit} form="updateForm" />
    </p>
    <p><span class="font-bold">Created at:</span> {toLocaleDateTime(data.org.created)}</p>
    <p><span class="font-bold">Updated at:</span> {toLocaleDateTime(data.org.updated)}</p>
</ResourcePage>
