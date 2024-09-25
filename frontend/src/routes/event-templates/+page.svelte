<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";
    import { breadcrumbs } from "$lib/stores";

    breadcrumbs.add("Event Templates", "event_templates");

    export let data: PageData;

    onMount(() => {
        const first = data.eventTemplates[0];

        if (first) {
            goto(`/event-templates/${first.id}`);
        }
    });
</script>
