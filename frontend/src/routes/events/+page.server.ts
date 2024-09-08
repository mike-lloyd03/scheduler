import type { Role, Event, EventTemplate } from "$lib/types";
import type { ClientResponseError } from "pocketbase";
import type { Actions, PageServerLoad } from "./$types";
import { fail } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ locals }) => {
    const eventTemplates = await locals.pb
        .collection("event_templates")
        .getFullList<EventTemplate>({});

    const events = await locals.pb.collection("events").getFullList<Event>({
        expand: "event_template",
    });

    const roles = await locals.pb.collection("roles").getFullList<Role>({});

    return { eventTemplates, events, roles };
};

export const actions: Actions = {
    createEvent: async ({ request, locals }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("events").create(data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    deleteEvent: async ({ request, locals }) => {
        const data = await request.formData();

        const id: string = data.get("eventID")?.toString() || "";

        try {
            await locals.pb.collection("events").delete(id);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
