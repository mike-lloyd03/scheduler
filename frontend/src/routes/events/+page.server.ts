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

    const users = await locals.pb.collection("users").getFullList<Role>({ expand: "orgs,groups" });

    return { eventTemplates, roleTemplates, events, roles, users };
};

export const actions: Actions = {
    createEvent: async ({ request, locals }) => {
        const data = await request.formData();

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
        console.log("rolesData:", JSON.stringify(rolesData));

        if (rolesData) {
            const roles: Record<string, string | null> = JSON.parse(rolesData.toString());

            for (const role of Object.entries(roles)) {
                try {
                    await locals.pb.collection("roles").create({
                        role_template: role[0],
                        event: createdEvent.id,
                        assigned_to: role[1],
                    });
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
