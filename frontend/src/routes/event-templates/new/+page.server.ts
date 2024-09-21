import type { Group } from "$lib/types";
import { fail, type Actions } from "@sveltejs/kit";
import type { ClientResponseError } from "pocketbase";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
    const groups = await locals.pb.collection("groups").getFullList<Group>();
    return { groups };
};

export const actions: Actions = {
    default: async ({ request, locals }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("event_templates").create(data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
