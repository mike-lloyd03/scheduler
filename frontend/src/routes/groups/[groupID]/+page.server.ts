import { fail } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { Group } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async (event) => {
    const group = await event.locals.pb
        .collection("groups")
        .getOne<Group>(event.params.groupID, { expand: "org" });
    return { group };
};

export const actions: Actions = {
    default: async ({ request, locals, params }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("groups").update(params.groupID, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
