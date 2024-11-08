<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";
    import type { Snippet } from "svelte";

    let { data, children }: { data: LayoutData, children: Snippet } = $props();

    let items = $derived(
        data.orgs.map((org) => {
            return { id: org.id, name: org.name };
        }),
    );
</script>

<NavLayout title="Organizations" {items} urlPath={$page.url.pathname}>
    {@render children?.()}
</NavLayout>
