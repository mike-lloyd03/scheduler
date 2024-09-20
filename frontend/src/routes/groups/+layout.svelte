<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavPage from "$lib/components/NavPage.svelte";

    export let data: LayoutData;

    $: items = data.groups.map((group) => {
        return { id: group.id, name: group.name };
    });
</script>

<NavPage title="Groups" {items} urlPath={$page.url.pathname}>
    <slot></slot>
</NavPage>
