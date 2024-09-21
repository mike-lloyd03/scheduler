import { fail } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { User } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async (event) => {
    const user = await event.locals.pb
        .collection("users")
        .getOne<User>(event.params.userID, { expand: "orgs,groups" });
    return { user };
};

export const actions: Actions = {
    update: async ({ request, locals, params }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("users").update(params.userID, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    delete: async ({ params, locals }) => {
        try {
            await locals.pb.collection("users").delete(params.userID);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
