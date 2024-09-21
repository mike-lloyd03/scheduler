<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";

    export let data: PageData;

    onMount(() => {
        const first = data.eventTemplates[0];

        if (first) {
            goto(`/event-templates/${first.id}`);
        }
    });
</script>
