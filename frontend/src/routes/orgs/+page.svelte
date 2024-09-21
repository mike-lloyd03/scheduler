<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";

    export let data: PageData;

    onMount(() => {
        const first = data.orgs[0];

        if (first) {
            goto(`/orgs/${first.id}`);
        }
    });
</script>
