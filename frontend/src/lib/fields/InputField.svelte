<script lang="ts">
    export let name: string;
    export let value: string;
    export let edit: boolean;
    export let form: string | null = null;
</script>

{#if edit}
    <input class="input inline" type="text" id={name} {name} {value} {form} />
{:else}
    {value}
{/if}
