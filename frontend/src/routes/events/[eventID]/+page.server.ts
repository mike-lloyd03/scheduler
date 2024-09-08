import type { Role, Event } from "$lib/types";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals, params }) => {
    const event = await locals.pb
        .collection("events")
        .getOne(params.eventID, { expand: "event_template" });

    const roles = await locals.pb.collection("roles").getFullList<Role>({
        expand: "role_template, assigned_to",
        filter: `event='${params.eventID}'`,
    });

    return { event, roles };
};
