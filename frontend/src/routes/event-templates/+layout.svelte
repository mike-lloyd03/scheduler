<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { goto } from "$app/navigation";

    export let data: LayoutData;

    $: items = data.eventTemplates.map((et) => {
        return { id: et.id, name: et.name };
    });
</script>

<NavLayout title="Event Templates" {items} urlPath={$page.url.pathname}>
    <div class="flex" slot="actions">
        <ActionButton type="new" onClick={() => goto("/event-templates/new")} />
    </div>
    <slot />
</NavLayout>
