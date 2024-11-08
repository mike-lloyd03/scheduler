<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ResourcePage from "$lib/components/ResourcePage.svelte";
    import { toLocaleDateTime } from "$lib/utils";
    import { breadcrumbs } from "$lib/stores";
    import MultiSelectField from "$lib/fields/MultiSelectField.svelte";
    import type { Org, OptionType, Group } from "$lib/types";
    import UpdatePassword from "./UpdatePassword.svelte";
    import PermissionsTable from "./PermissionsTable.svelte";

    export let data: PageData;

    $: breadcrumbs.clear().add("Users", "users").add(data.user.name, data.user.id);

    function orgsAsOptions(orgs: Org[]): OptionType[] {
        return orgs.map((o) => {
            return { value: o.id, label: o.name };
        });
    }

    function groupsAsOptions(groups: Group[]): OptionType[] {
        return groups.map((o) => {
            return { value: o.id, label: o.name };
        });
    }

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

    <div>
        <UpdatePassword />
    </div>

    {#if data.permissions.length}
        <div>
            <PermissionsTable permissions={data.permissions} />
        </div>
    {/if}
</ResourcePage>
