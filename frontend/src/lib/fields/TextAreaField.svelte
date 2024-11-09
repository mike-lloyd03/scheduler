<script lang="ts">
    export let name: string;
    export let value: string | undefined;
    export let edit: boolean;
    export let form: string | undefined = undefined;
</script>

{#if edit}
    <textarea class="textarea w-fit" id={name} {name} {value} {form}></textarea>
{:else}
    {value}
{/if}
