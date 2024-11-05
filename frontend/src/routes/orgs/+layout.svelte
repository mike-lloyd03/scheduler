<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";

    export let data: LayoutData;

    $: items = data.orgs.map((org) => {
        return { id: org.id, name: org.name };
    });
</script>

<NavLayout title="Organizations" {items} urlPath={$page.url.pathname}>
    <slot />
</NavLayout>
