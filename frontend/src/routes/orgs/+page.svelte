<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
    const first = data.orgs[0];

    if (first) {
        goto(`/orgs/${first.id}`);
    }
</script>
