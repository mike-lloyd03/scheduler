<nav class="space-between list-nav flex flex-col p-4 md:h-full">
    <ul class="grow">
        <li><a href="/">Dashboard</a></li>
        <li><a href="/orgs">Orgs</a></li>
        <li><a href="/groups">Groups</a></li>
        <li><a href="/events">Events</a></li>
        <li><a href="/users">Users</a></li>
    </ul>
    <hr />
    <div class="mt-2">
        <form method="POST" action="/logout">
            <button>Logout</button>
        </form>
    </div>
</nav>
