<script lang="ts">
    import type { LayoutData } from "./$types";
    import { page } from "$app/stores";
    import NavLayout from "$lib/components/NavLayout.svelte";
    import ActionButton from "$lib/components/ActionButton.svelte";
    import { goto } from "$app/navigation";

    export let data: LayoutData;

    $: items = data.users.map((user) => {
        return { id: user.id, name: user.name };
    });
</script>

<NavLayout title="Users" {items} urlPath={$page.url.pathname}>
    <div class="flex" slot="actions">
        <ActionButton type="new" onClick={() => goto("/users/new")} />
    </div>
    <slot />
</NavLayout>
