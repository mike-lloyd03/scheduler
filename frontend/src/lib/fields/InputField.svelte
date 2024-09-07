<script lang="ts">
    export let title: string;
    export let value: string;
    export let edit: boolean;
    export let hideTitle = false;
</script>

<div>
    {#if !hideTitle}
        <span class="font-bold">{title}: </span>
    {/if}
    {#if edit}
        <input
            class="input inline"
            type="text"
            id={title.toLowerCase()}
            name={title.toLowerCase()}
            {value}
        />
    {:else}
        {value}
    {/if}
</div>
