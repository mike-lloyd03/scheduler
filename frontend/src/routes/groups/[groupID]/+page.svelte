<script lang="ts">
    import InputField from "$lib/fields/InputField.svelte";
    import type { PageData } from "./$types";
    import ResourcePage from "$lib/components/ResourcePage.svelte";
    import { toLocaleDateTime } from "$lib/utils";

    export let data: PageData;

    let edit = false;
</script>

<ResourcePage
    bind:edit
    title="Group"
    baseURL="/groups"
    deleteBody="Are you sure you want to delete the group '{data.group
        .name}'? This will delete all events and other resources associated with it."
>
    <p>
        <span class="font-bold">Name:</span>
        <InputField name="name" value={data.group.name} {edit} />
    </p>
    <p><span class="font-bold">Organization:</span> {data.group.expand?.org.name}</p>
    <p><span class="font-bold">Created at:</span> {toLocaleDateTime(data.group.created)}</p>
    <p><span class="font-bold">Updated at:</span> {toLocaleDateTime(data.group.updated)}</p>
</ResourcePage>
