<script lang="ts">
    export let title: string;
    export let value: string;
    export let edit: boolean;
</script>

<div>
    <span class="font-bold">{title}: </span>
    {#if edit}
        <input
            class="input"
            type="text"
            id={title.toLowerCase()}
            name={title.toLowerCase()}
            {value}
        />
    {:else}
        {value}
    {/if}
</div>
