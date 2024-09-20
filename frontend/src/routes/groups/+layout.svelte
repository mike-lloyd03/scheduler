<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";

    export let data: LayoutData;

    $: items = data.groups.map((group) => {
        return { id: group.id, name: group.name };
    });
</script>

<NavLayout title="Groups" {items} urlPath={$page.url.pathname}>
    <slot></slot>
</NavLayout>
