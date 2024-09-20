<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavPage from "$lib/components/NavPage.svelte";

    export let data: LayoutData;

    const items = data.orgs.map((org) => {
        return { id: org.id, name: org.name };
    });
</script>

<NavPage title="Organizations" {items} urlPath={$page.url.pathname}>
    <slot></slot>
</NavPage>
