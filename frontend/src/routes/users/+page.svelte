<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import type { PageData } from "./$types";

    export let data: PageData;

    onMount(() => {
        const first = data.users[0];

        if (first) {
            goto(`/users/${first.id}`);
        }
    });
</script>
