import type { Role, RoleTemplate, User } from "$lib/types";
import { fail, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async ({ locals, params }) => {
    const event = await locals.pb
        .collection("events")
        .getOne(params.eventID, { expand: "event_template" });

    const groupID = event.expand?.event_template.group;
    const eventTemplateID = event.expand?.event_template.id;

    const roles = await locals.pb.collection("roles").getFullList<Role>({
        expand: "role_template, assigned_to",
        filter: `event='${params.eventID}'`,
    });

    const roleTemplates = await locals.pb.collection("role_templates").getFullList<RoleTemplate>({
        filter: `event_template='${eventTemplateID}'`,
    });

    const users = await locals.pb
        .collection("users")
        .getFullList<User>({ filter: `groups~'${groupID}'` });

    return { event, roles, users, roleTemplates };
};

export const actions: Actions = {
    updateRole: async ({ request, locals }) => {
        const data = await request.formData();

        const roleID = data.get("roleID")?.toString() || "";

        try {
            await locals.pb.collection("roles").update(roleID, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    newRole: async ({ request, locals }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("roles").create(data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    deleteRole: async ({ request, locals }) => {
        const data = await request.formData();

        const roleID = data.get("roleID")?.toString() || "";

        try {
            await locals.pb.collection("roles").delete(roleID);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
