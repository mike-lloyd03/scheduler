<script lang="ts">
    import type { Snippet } from "svelte";

    interface Props {
        title: string;
        items: { id: string; name: string }[];
        urlPath: string;
        actions?: Snippet;
        children: Snippet;
    }

    let { title, items, urlPath, actions, children }: Props = $props();

    const urlBase = urlPath.split("/")[1];

    function hrefClasses(href: string) {
        return href === urlPath ? "!variant-filled-primary" : "";
    }
</script>

<h2 class="h2">{title}</h2>

<div class="flex h-full">
    <div class="flex flex-col">
        <nav class="card list-nav h-full min-w-72 overflow-scroll">
            <!-- {#if $$slots.actions} -->
            <!--     <header class="variant-soft card-header mb-2 p-2"> -->
            <!--         <slot name="actions" /> -->
            <!--     </header> -->
            <!-- {/if} -->
            {#if actions}
                <header class="variant-soft card-header mb-2 p-2">
                    {@render actions()}
                </header>
            {/if}

            <div class="mt-2">
                <ul>
                    {#each items as item}
                        <li class="mx-2">
                            <a
                                href="/{urlBase}/{item.id}"
                                class={hrefClasses(`/${urlBase}/${item.id}`)}
                            >
                                <span class="flex-auto">{item.name}</span>
                            </a>
                        </li>
                    {/each}
                </ul>
            </div>
        </nav>
    </div>

    <div class="mx-8 w-full">
        {@render children?.()}
    </div>
</div>
