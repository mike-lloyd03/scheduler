<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ResourcePage from "$lib/components/ResourcePage.svelte";
    import { hasOrgRole, toLocaleDateTime } from "$lib/utils";
    import { breadcrumbs } from "$lib/stores";
    import MultiSelectField from "$lib/fields/MultiSelectField.svelte";
    import { type Org, type OptionTypes, type Group, UserRole } from "$lib/types";
    import UpdatePassword from "./UpdatePassword.svelte";
    import PermissionsTable from "./PermissionsTable.svelte";

    let { data }: { data: PageData } = $props();
    const permissions = data.currentUserPermissions ?? [];

    $effect(() => {
        breadcrumbs
            .clear()
            .add("Users", "users")
            .add(data.user.name || data.user.username, data.user.id);
    });

    function orgsAsOptions(orgs: Org[]): OptionTypes {
        return orgs.map((o) => {
            return { value: o.id, label: o.name };
        });
    }

    function groupsAsOptions(groups: Group[]): OptionTypes {
        return groups.map((o) => {
            return { value: o.id, label: o.name };
        });
    }

    let edit = $state(false);

    let showEdit = $derived(
        data.user.expand?.orgs?.some((org) => hasOrgRole(permissions, org, UserRole.Admin)),
    );
    let showUpdatePassword = $derived(data.currentUser?.id == data.user.id || showEdit);
    let showDelete = $derived(showEdit);
</script>

<ResourcePage
    bind:edit
    resourceName="User"
    title={data.user.name || data.user.username}
    baseURL="/users"
    deleteBody="Are you sure you want to delete the user '{data.user.name}'?"
    {showEdit}
    {showDelete}
>
    <div class="py-4">
        <p>
            <span class="font-bold">Name:</span>
            <InputField name="name" value={data.user.name || ""} {edit} form="updateForm" />
        </p>

        <p>
            <span class="font-bold">Email:</span>
            <InputField name="email" value={data.user.email} {edit} form="updateForm" />
        </p>

        <p>
            <span class="font-bold">Orgs:</span>
            <MultiSelectField
                name="orgs"
                value={orgsAsOptions(data.user.expand?.orgs || [])}
                options={orgsAsOptions(data.orgs)}
                {edit}
                form="updateForm"
            />
        </p>

        <p>
            <span class="font-bold">Groups:</span>
            <MultiSelectField
                name="groups"
                value={groupsAsOptions(data.user.expand?.groups || [])}
                options={groupsAsOptions(data.groups)}
                {edit}
                form="updateForm"
            />
        </p>

        <p><span class="font-bold">Created at:</span> {toLocaleDateTime(data.user.created)}</p>

        <p><span class="font-bold">Updated at:</span> {toLocaleDateTime(data.user.updated)}</p>
    </div>

    {#if showUpdatePassword}
        <div>
            <UpdatePassword />
        </div>
    {/if}

    {#if data.permissions.length}
        <div>
            <PermissionsTable permissions={data.permissions} />
        </div>
    {/if}
</ResourcePage>
