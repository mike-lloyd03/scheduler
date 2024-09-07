import type { Role, Event } from "$lib/types";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals, params }) => {
    const events = await locals.pb.collection("events").getFullList<Event>({});

    const roles = await locals.pb.collection("roles").getFullList<Role>({});

    return { events, roles };
};
