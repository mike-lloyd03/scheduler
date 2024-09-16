<script lang="ts">
    export let name: string;
    export let value: string | undefined;
    export let edit: boolean;
</script>

{#if edit}
    <textarea class="textarea" id={name} {name} {value} />
{:else}
    {value}
{/if}
