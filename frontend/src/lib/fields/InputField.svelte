<script lang="ts">
    export let name: string;
    export let value: string;
    export let edit: boolean;
    export let form: string | undefined = undefined;
</script>

{#if edit}
    <input class="input w-fit" type="text" id={name} {name} {value} {form} />
{:else}
    {value}
{/if}
