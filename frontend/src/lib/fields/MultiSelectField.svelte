<script lang="ts">
    import type { OptionType } from "$lib/types";

    export let name: string;
    export let value: OptionType[];
    export let options: OptionType[];
    export let edit: boolean;
    export let form: string | undefined = undefined;

    function isSelected(option: OptionType) {
        return value.some((v) => v.label === option.label || v.value === option.value);
    }
</script>

{#if edit}
    <select class="select w-fit" id={name} {name} value={value.map((v) => v.value)} multiple {form}>
        {#each options as option}
            <option value={option.value} selected={isSelected(option)}>{option.label}</option>
        {/each}
    </select>
{:else}
    {#each value as v, i (v.value)}
        {v.label}{i < value.length - 1 ? ", " : ""}
    {/each}
{/if}
