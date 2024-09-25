<script lang="ts">
    import type { User } from "$lib/types";

    export let drawerClose: (() => void) | null = null;
    export let user: User;
</script>

<nav class="space-between list-nav flex flex-col p-4 md:h-full">
    <ul>
        <li><a on:click={drawerClose} href="/">Dashboard</a></li>
        <li><a on:click={drawerClose} href="/events">Events</a></li>
    </ul>
    {#if user.role != "member"}
        <h3 class="h3 ml-4 mt-8 underline">Admin</h3>
        <ul>
            <li><a on:click={drawerClose} href="/orgs">Organizations</a></li>
            <li><a on:click={drawerClose} href="/groups">Groups</a></li>
            <li><a on:click={drawerClose} href="/event-templates">Event Templates</a></li>
            <li><a on:click={drawerClose} href="/users">Users</a></li>
        </ul>
    {/if}
    <div class="grow"></div>
    <hr />
    <div class="mt-2">
        <a href="/users/{user.id}">{user.name}</a>
        <form method="POST" action="/logout">
            <button>Logout</button>
        </form>
    </div>
</nav>
