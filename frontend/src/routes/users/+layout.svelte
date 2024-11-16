<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { goto } from "$app/navigation";
    import type { Snippet } from "svelte";
    import { CurrentUserRole } from "$lib/types";
    import { getRole } from "$lib/utils";

    let { data, children }: { data: LayoutData; children: Snippet } = $props();

    let items = $derived(
        data.users.map((user) => {
            return { id: user.id, name:  user.name && user.name != "" ? user.name : user.username };
        }),
    );

    let showActions = $derived(data.currentUserPermissions && getRole(data.currentUserPermissions) == CurrentUserRole.OrgAdmin)
</script>

{#snippet actions()}
    <div class="flex" slot="actions">
        <ActionButton type="new" onClick={() => goto("/users/new")} />
    </div>
{/snippet}


<NavLayout title="Users" {items} urlPath={$page.url.pathname} actions={showActions ? actions: undefined}>

    {@render children?.()}
</NavLayout>
