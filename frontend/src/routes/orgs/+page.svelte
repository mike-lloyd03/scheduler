<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    export let data: PageData;
    const firstOrg = data.orgs[0];

    if (firstOrg) {
        goto(`/orgs/${firstOrg.id}`);
    }
</script>
