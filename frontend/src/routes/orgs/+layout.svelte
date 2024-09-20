<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { goto } from "$app/navigation";

    export let data: LayoutData;

    $: items = data.orgs.map((org) => {
        return { id: org.id, name: org.name };
    });
</script>

<NavLayout title="Organizations" {items} urlPath={$page.url.pathname}>
    <div class="flex" slot="actions">
        <ActionButton type="new" onClick={() => goto("/orgs/new")} />
    </div>
    <slot />
</NavLayout>
