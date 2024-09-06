<script lang="ts">
    import { Toaster } from "svelte-french-toast";
    import { AppShell, AppBar } from "@skeletonlabs/skeleton";
    import { initializeStores, Drawer, getDrawerStore } from "@skeletonlabs/skeleton";
    import "../app.css";
    import Navigation from "./Navigation.svelte";
    import type { PageData } from "./$types";
    import Hamburger from "$lib/svg/Hamburger.svelte";

    export let data: PageData;

    initializeStores();

    const drawerStore = getDrawerStore();

    function drawerOpen(): void {
        drawerStore.open({});
    }

    function drawerClose(): void {
        drawerStore.close();
    }
</script>

{#if data.currentUser}
    <Drawer>
        <div class="flex flex-row">
            <button class="btn btn-sm mr-4" on:click={drawerClose}>
                <Hamburger />
            </button>
            <h2 class="p-4">Navigation</h2>
        </div>
        <hr />
        <Navigation />
    </Drawer>
{/if}

<AppShell slotSidebarLeft="bg-surface-700 w-0 md:w-64 h-full">
    <svelte:fragment slot="header">
        <AppBar class="shadow-lg">
            <svelte:fragment slot="lead">
                <div class="flex items-center">
                    <button
                        class="btn btn-sm mr-4 md:hidden {data.currentUser ? '' : 'hidden'}"
                        on:click={drawerOpen}
                    >
                        <Hamburger />
                    </button>
                    <strong class="text-xl uppercase">Scheduler</strong>
                </div>
            </svelte:fragment>
        </AppBar>
    </svelte:fragment>

    <svelte:fragment slot="sidebarLeft">
        {#if data.currentUser}
            <Navigation />
        {/if}
    </svelte:fragment>

    <div class="container h-max space-y-4 p-10">
        <slot />
    </div>

    <Toaster />
</AppShell>
