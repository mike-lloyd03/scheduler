<script lang="ts">
    import { Toaster } from "svelte-french-toast";
    import { AppShell, AppBar } from "@skeletonlabs/skeleton";
    import { initializeStores, Drawer, getDrawerStore } from "@skeletonlabs/skeleton";
    import "../app.css";
    import Navigation from "./Navigation.svelte";

    initializeStores();

    const drawerStore = getDrawerStore();

    function drawerOpen(): void {
        drawerStore.open({});
    }
</script>

<Drawer>
    <h2 class="p-4">Navigation</h2>
    <hr />
    <Navigation /></Drawer
>

<AppShell slotSidebarLeft="bg-surface-500/5 w-0 md:w-64">
    <svelte:fragment slot="header">
        <AppBar class="shadow-lg">
            <svelte:fragment slot="lead">
                <div class="flex items-center">
                    <button class="btn btn-sm mr-4 md:hidden" on:click={drawerOpen}>
                        <span>
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke-width="1.5"
                                stroke="currentColor"
                                class="size-6"
                            >
                                <path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
                                />
                            </svg>
                        </span>
                    </button>
                    <strong class="text-xl uppercase">Scheduler</strong>
                </div>
            </svelte:fragment>
        </AppBar>
    </svelte:fragment>
    <svelte:fragment slot="sidebarLeft"><Navigation /></svelte:fragment>

    <div class="container h-max space-y-4 p-10">
        <slot />
    </div>

    <Toaster />
</AppShell>
