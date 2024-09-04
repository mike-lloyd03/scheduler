import { fail } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { Org } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async (event) => {
    const org = await event.locals.pb.collection("orgs").getOne<Org>(event.params.orgID);
    return { org };
};

export const actions: Actions = {
    default: async ({ request, locals, params }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("orgs").update(params.orgID, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
