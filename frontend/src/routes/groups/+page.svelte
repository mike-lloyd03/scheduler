<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
    const first = data.groups[0];

    if (first) {
        goto(`/groups/${first.id}`);
    }
</script>
