<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";

    let { data }: { data: PageData } = $props();

    onMount(() => {
        const first = data.users[0];

        if (first) {
            goto(`/users/${first.id}`);
        }
    });
</script>
