<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { goto } from "$app/navigation";

    export let data: LayoutData;

    $: items = data.groups.map((group) => {
        return { id: group.id, name: group.name };
    });
</script>

<NavLayout title="Groups" {items} urlPath={$page.url.pathname}>
    <div class="flex" slot="actions">
        <ActionButton type="new" onClick={() => goto("/groups/new")} />
    </div>
    <slot />
</NavLayout>
