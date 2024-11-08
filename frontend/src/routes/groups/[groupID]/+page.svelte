<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ResourcePage from "$lib/components/ResourcePage.svelte";
    import { hasGroupRole, toLocaleDateTime } from "$lib/utils";
    import { breadcrumbs } from "$lib/stores";
    import { UserRole } from "$lib/types";
    import { hasOrgRole } from "$lib/utils";

    export let data: PageData;
    const permissions = data.currentUserPermissions ?? [];

    $: breadcrumbs.clear().add("Groups", "groups").add(data.group.name, data.group.id);

    let edit = false;
    $: showDelete = hasOrgRole(permissions ?? [], data.group.expand!.org, UserRole.Admin);
    $: showEdit = showDelete || hasGroupRole(permissions ?? [], data.group, UserRole.Admin);
</script>

<ResourcePage
    bind:edit
    resourceName="Group"
    baseURL="/groups"
    deleteBody="Are you sure you want to delete the group '{data.group
        .name}'? This will delete all events and other resources associated with it."
    {showEdit}
    {showDelete}
>
    <p>
        <span class="font-bold">Name:</span>
        <InputField name="name" value={data.group.name} {edit} form="updateForm" />
    </p>
    <p><span class="font-bold">Organization:</span> {data.group.expand?.org.name}</p>
    <p><span class="font-bold">Created at:</span> {toLocaleDateTime(data.group.created)}</p>
    <p><span class="font-bold">Updated at:</span> {toLocaleDateTime(data.group.updated)}</p>
</ResourcePage>
