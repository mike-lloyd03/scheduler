<script lang="ts">
    import type { OptionTypes } from "$lib/types";

    export let name: string;
    export let value: OptionTypes;
    export let options: OptionTypes;
    export let edit: boolean;
    export let form: string | undefined = undefined;

    function isSelected(option: OptionTypes) {
        return value.some(
            (v: { value: string | undefined; label: string }) =>
                v.label === option.label || v.value === option.value,
        );
    }
</script>

{#if edit}
    <select class="select w-fit" id={name} {name} value={value.map((v: { value: string | undefined; label: string }) => v.value)} multiple {form}>
        {#each options as option}
            <option value={option.value} selected={isSelected(option)}>{option.label}</option>
        {/each}
    </select>
{:else}
    {#each value as v, i (v.value)}
        {v.label}{i < value.length - 1 ? ", " : ""}
    {/each}
{/if}
