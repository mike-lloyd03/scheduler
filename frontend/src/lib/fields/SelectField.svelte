<script lang="ts">
    export let name: string;
    export let value: { value: string | undefined; label: string };
    export let options: { value: string | undefined; label: string }[];
    export let edit: boolean;
    export let form: string | undefined = undefined;
</script>

{#if edit}
    <select class="select w-fit" id={name} {name} value={value.value} {form}>
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
