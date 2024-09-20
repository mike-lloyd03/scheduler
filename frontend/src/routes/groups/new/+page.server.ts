import { fail, type Actions } from "@sveltejs/kit";
import type { ClientResponseError } from "pocketbase";
import type { PageServerLoad } from "./$types";
import type { Org } from "$lib/types";

export const load: PageServerLoad = async ({ locals }) => {
    const orgs = await locals.pb.collection("orgs").getFullList<Org>();

    return { orgs };
};

export const actions: Actions = {
    default: async ({ request, locals }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("groups").create(data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
