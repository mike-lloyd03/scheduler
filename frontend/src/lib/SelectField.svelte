<script lang="ts">
    export let title: string;
    export let value: string;
    export let options: string[];
    export let edit: boolean;
</script>

<div>
    <span class="font-bold">{title}: </span>
    {#if edit}
        <select class="select" id={title.toLowerCase()} name={title.toLowerCase()} {value}>
            {#each options as option}
                <option value={option} selected={option == value}>{option}</option>
            {/each}
        </select>
    {:else}
        {value}
    {/if}
</div>
