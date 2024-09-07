import type { Role, Event } from "$lib/types";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
    const events = await locals.pb.collection("events").getFullList<Event>({
        expand: "event_template",
    });

    const roles = await locals.pb.collection("roles").getFullList<Role>({});

    return { events, roles };
};
