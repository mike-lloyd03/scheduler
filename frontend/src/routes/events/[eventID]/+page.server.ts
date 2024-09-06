import { fail } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { RoleTemplate, EventTemplate } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async ({ locals, params }) => {
    const eventTemplate = await locals.pb
        .collection("event_templates")
        .getOne<EventTemplate>(params.eventID, { expand: "group" });

    const roleTemplates = await locals.pb.collection("role_templates").getFullList<RoleTemplate>({
        filter: `(event_template='${params.eventID}')`,
    });

    return { eventTemplate, roleTemplates };
};

export const actions: Actions = {
    updateEvent: async ({ request, locals, params }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("event_templates").update(params.eventID, data);
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
            event_template: params.eventID,
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
