<script lang="ts">
    export let name: string;
    export let value: string | undefined;
    export let edit: boolean;
    export let form: string | null = null;
</script>

{#if edit}
    <textarea class="textarea" id={name} {name} {value} {form} />
{:else}
    {value}
{/if}
