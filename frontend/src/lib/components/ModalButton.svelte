<script lang="ts">
    export let onClick: () => void;
    export let text: string;
</script>

<button class="variant-ringed-primary btn hover:variant-filled-primary" on:click={onClick}
    >{text}</button
>
