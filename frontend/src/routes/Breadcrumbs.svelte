<script lang="ts">
    import { page } from "$app/stores";
    import { breadcrumbs } from "$lib/stores";

    $: urlPath = $page.url.pathname;
</script>

{#if urlPath != "/" && urlPath != "/login"}
    <div class="p-4">
        <ol class="breadcrumb">
            <li class="crumb"><a class="anchor" href="/">Home</a></li>
            {#if $breadcrumbs.length > 0}
                <li class="crumb-separator" aria-hidden="true">&rsaquo;</li>
            {/if}
            {#each $breadcrumbs as element, i}
                {#if i < $breadcrumbs.length - 1}
                    <li class="crumb">
                        <a
                            class="anchor"
                            href={"/" +
                                $breadcrumbs
                                    .map((b) => b.url)
                                    .slice(0, i + 1)
                                    .join("/")}>{element.label}</a
                        >
                    </li>
                    <li class="crumb-separator" aria-hidden="true">&rsaquo;</li>
                {:else}
                    <div class="crumb">
                        {element.label}
                    </div>
                {/if}
            {/each}
        </ol>
    </div>
{/if}
