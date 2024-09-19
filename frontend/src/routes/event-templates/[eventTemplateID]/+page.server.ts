import { fail } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { RoleTemplate, EventTemplate, Role, Event } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async ({ locals, params }) => {
    const eventTemplate = await locals.pb
        .collection("event_templates")
        .getOne<EventTemplate>(params.eventTemplateID, { expand: "group" });

    const roleTemplates = await locals.pb.collection("role_templates").getFullList<RoleTemplate>({
        filter: `(event_template='${params.eventTemplateID}')`,
    });

    const events = await locals.pb.collection("events").getFullList<Event>({
        filter: `(event_template='${params.eventTemplateID}')`,
        expand: "event_template",
    });

    const roles = await locals.pb.collection("roles").getFullList<Role>({
        expand: "role_template.event_template",
        filter: `(role_template.event_template.id='${params.eventTemplateID}')`,
    });

    return { eventTemplate, roleTemplates, events, roles };
};

export const actions: Actions = {
    updateEvent: async ({ request, locals, params }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("event_templates").update(params.eventTemplateID, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    deleteEventTemplate: async ({ locals, params }) => {
        try {
            await locals.pb.collection("event_templates").delete(params.eventTemplateID);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    newRole: async ({ request, locals, params }) => {
        const data = await request.formData();

        const body = {
            name: data.get("name"),
            description: data.get("description"),
            event_template: params.eventTemplateID,
        };

        try {
            await locals.pb.collection("role_templates").create(body);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    updateRole: async ({ request, locals }) => {
        const data = await request.formData();

        const id: string = data.get("roleID")?.toString() || "";

        try {
            await locals.pb.collection("role_templates").update(id, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    deleteRole: async ({ request, locals }) => {
        const data = await request.formData();

        const id: string = data.get("roleID")?.toString() || "";

        try {
            await locals.pb.collection("role_templates").delete(id);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
