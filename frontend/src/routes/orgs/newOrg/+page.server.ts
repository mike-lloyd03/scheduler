import { fail, type Actions } from "@sveltejs/kit";
import type { ClientResponseError } from "pocketbase";

export const actions: Actions = {
    default: async ({ request, locals }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("orgs").create(data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
