import type { EventTemplate } from "$lib/types";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ locals }) => {
    const eventTemplates = await locals.pb
        .collection("event_templates")
        .getFullList<EventTemplate>({ expand: "group", sort: "name" });

    return { eventTemplates };
};
