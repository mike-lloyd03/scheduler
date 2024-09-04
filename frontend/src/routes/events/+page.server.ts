import type { PageServerLoad } from "./$types";
import type { EventTemplate } from "$lib/types";

export const load: PageServerLoad = async (event) => {
    const event_templates = await event.locals.pb
        .collection("event_templates")
        .getFullList<EventTemplate>({ expand: "group" });
    return { event_templates };
};
