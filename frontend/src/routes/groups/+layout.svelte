<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { goto } from "$app/navigation";
    import type { Snippet } from "svelte";
    import { getRole } from "$lib/utils";
    import { CurrentUserRole } from "$lib/types";

	let { data, children }: { data: LayoutData, children: Snippet } = $props();

    let items = $derived(data.groups.map((group) => {
        return { id: group.id, name: group.name };
    }))
</script>

<NavLayout title="Groups" {items} urlPath={$page.url.pathname}>
    {#if data.currentUserPermissions && getRole(data.currentUserPermissions) in [CurrentUserRole.OrgAdmin, CurrentUserRole.OrgManager]}
    {#snippet actions()}
            <div class="flex">
                <ActionButton type="new" onClick={() => goto("/groups/new")} />
            </div>
    {/snippet}
{/if}
        
    {@render children?.()}
</NavLayout>
