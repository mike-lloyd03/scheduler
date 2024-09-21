<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";

    export let data: PageData;

    onMount(() => {
        const first = data.groups[0];

        if (first) {
            goto(`/groups/${first.id}`);
        }
    });
</script>
