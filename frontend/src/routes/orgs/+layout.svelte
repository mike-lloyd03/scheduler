<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavPage from "$lib/components/NavPage.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { goto } from "$app/navigation";

    export let data: LayoutData;

    $: items = data.orgs.map((org) => {
        return { id: org.id, name: org.name };
    });
</script>

<NavPage title="Organizations" {items} urlPath={$page.url.pathname}>
    <div class="flex" slot="actions">
        <ActionButton type="new" onClick={() => goto("/orgs/newOrg")} />
    </div>
    <slot />
</NavPage>
