<script lang="ts">
    export let title: string;
    export let items: { id: string; name: string }[];
    export let urlPath: string;

    const urlBase = urlPath.split("/")[1];

    $: classesActive = (href: string) => (href === urlPath ? "!variant-filled-primary" : "");
</script>

<h2 class="h2">{title}</h2>

<div class="flex h-full">
    <nav class="card list-nav h-full min-w-72 overflow-scroll p-4">
        <ul>
            {#each items as item}
                <li>
                    <a href="/{urlBase}/{item.id}" class={classesActive(`/${urlBase}/${item.id}`)}>
                        <span class="flex-auto">{item.name}</span>
                    </a>
                </li>
            {/each}
        </ul>
    </nav>
    <div class="mx-8 w-full">
        <slot />
    </div>
</div>
