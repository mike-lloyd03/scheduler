<script lang="ts">
    import { CurrentUserRole, type Permission, type User } from "$lib/types";
    import { getRole } from "$lib/utils";

    export let drawerClose: (() => void) | null = null;
    export let user: User | undefined;
    export let permissions: Permission[] | undefined;

    let role: CurrentUserRole | undefined = undefined;
    if (permissions) {
        role = getRole(permissions);
    }

    function formatUserRole() {
        if (role) {
            switch (role) {
                case CurrentUserRole.Member:
                    return "Member";
                case CurrentUserRole.GroupManager:
                    return "Group Manager";
                case CurrentUserRole.GroupAdmin:
                    return "Group Admin";
                case CurrentUserRole.OrgManager:
                    return "Org Manager";
                case CurrentUserRole.OrgAdmin:
                    return "Org Admin";
            }
        }
        return "";
    }

    type link = { path: string; title: string };

    const links: link[] = [
        { path: "/", title: "Dashboard" },
        { path: "/events", title: "Events" },
    ];

    const adminLinks: link[] = [
        { path: "/orgs", title: "Organizations" },
        { path: "/groups", title: "Groups" },
        { path: "/event-templates", title: "Event Templates" },
        { path: "/users", title: "Users" },
    ];
</script>

<nav class="space-between list-nav flex flex-col p-4 md:h-full">
    <ul>
        {#each links as link}
            <li><a on:click={drawerClose} href={link.path}>{link.title}</a></li>
        {/each}
    </ul>

    {#if role != undefined && role != CurrentUserRole.Member}
        <h3 class="h3 ml-4 mt-8 underline">Admin</h3>
        <ul>
            {#each adminLinks as link}
                <li><a on:click={drawerClose} href={link.path}>{link.title}</a></li>
            {/each}
        </ul>
    {/if}

    <div class="grow"></div>
    <hr />
    <div class="mt-2">
        <a href="/users/{user?.id}">
            <div>
                <p>{user?.name != "" ? user?.name : user.username}</p>
                <p class="text-sm text-surface-400">({formatUserRole()})</p>
            </div>
        </a>
        <form method="POST" action="/logout">
            <button>Logout</button>
        </form>
    </div>
</nav>
