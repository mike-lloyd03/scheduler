<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";

    let { data }: { data: PageData } = $props();

    onMount(() => {
        const first = data.orgs[0];

        if (first) {
            goto(`/orgs/${first.id}`);
        }
    });
</script>
