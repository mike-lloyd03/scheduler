<script lang="ts">
    export let name: string;
    export let value: { value: string | null; label: string };
    export let options: { value: string; label: string }[];
    export let edit: boolean;
</script>

{#if edit}
    <select class="select" id={name} {name} value={value.value}>
        {#each options as option}
            <option
                value={option.value}
                selected={option.label === value.label || option.value === value.value}
                >{option.label}</option
            >
        {/each}
    </select>
{:else}
    {value.label}
{/if}
