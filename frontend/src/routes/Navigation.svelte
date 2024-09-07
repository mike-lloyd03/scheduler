<script lang="ts">
    export let drawerClose: (() => void) | null = null;
</script>

<nav class="space-between list-nav flex flex-col p-4 md:h-full">
    <ul class="grow">
        <li><a on:click={drawerClose} href="/">Dashboard</a></li>
        <li><a on:click={drawerClose} href="/orgs">Orgs</a></li>
        <li><a on:click={drawerClose} href="/groups">Groups</a></li>
        <li><a on:click={drawerClose} href="/event-templates">Event Templates</a></li>
        <li><a on:click={drawerClose} href="/users">Users</a></li>
    </ul>
    <hr />
    <div class="mt-2">
        <form method="POST" action="/logout">
            <button>Logout</button>
        </form>
    </div>
</nav>
