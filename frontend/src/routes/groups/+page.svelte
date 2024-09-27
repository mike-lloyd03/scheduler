<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";
    import { breadcrumbs } from "$lib/stores";

    export let data: PageData;

    breadcrumbs.clear().add("Groups", "groups");

    onMount(() => {
        const first = data.groups[0];

        if (first) {
            goto(`/groups/${first.id}`);
        }
    });
</script>
