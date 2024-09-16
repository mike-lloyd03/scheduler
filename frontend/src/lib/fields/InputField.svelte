<script lang="ts">
    export let name: string;
    export let value: string;
    export let edit: boolean;
</script>

{#if edit}
    <input class="input inline" type="text" id={name} {name} {value} />
{:else}
    {value}
{/if}
