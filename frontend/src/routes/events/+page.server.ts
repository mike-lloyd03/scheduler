import type { Role, Event, EventTemplate, RoleTemplate } from "$lib/types";
import type { ClientResponseError } from "pocketbase";
import type { Actions, PageServerLoad } from "./$types";
import { fail } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ locals }) => {
    const eventTemplates = await locals.pb
        .collection("event_templates")
        .getFullList<EventTemplate>({});

    const roleTemplates = await locals.pb.collection("role_templates").getFullList<RoleTemplate>({
        expand: "event_template",
    });

    const events = await locals.pb.collection("events").getFullList<Event>({
        expand: "event_template",
    });

    const roles = await locals.pb.collection("roles").getFullList<Role>({});

    return { eventTemplates, roleTemplates, events, roles };
};

export const actions: Actions = {
    createEvent: async ({ request, locals }) => {
        const data = await request.formData();

        data.forEach((v, k) => console.log(`${k}: ${v}`));

        const createEventData = {
            event_template: data.get("event_template"),
            datetime: data.get("datetime"),
        };

        let createdEvent: Event;

        try {
            createdEvent = await locals.pb.collection("events").create(createEventData);
            console.log(JSON.stringify(createdEvent));
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }

        const rolesData = data.get("roles");

        if (rolesData) {
            const roles: string[] = JSON.parse(rolesData.toString());

            for (const role of roles) {
                try {
                    await locals.pb
                        .collection("roles")
                        .create({ role_template: role, event: createdEvent.id });
                } catch (e) {
                    const error = e as ClientResponseError;
                    console.log("Error: ", JSON.stringify(error.data));

                    return fail(error.status, { message: error.message });
                }
            }
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
