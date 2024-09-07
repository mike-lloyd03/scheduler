<script lang="ts">
    export let title: string;
    export let value: string | undefined;
    export let edit: boolean;
    export let hideTitle = false;
</script>

<div>
    {#if !hideTitle}
        <span class="font-bold">{title}: </span>
    {/if}
    {#if edit}
        <textarea class="textarea" id={title.toLowerCase()} name={title.toLowerCase()} {value} />
    {:else}
        {value}
    {/if}
</div>
