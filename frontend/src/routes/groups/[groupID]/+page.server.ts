import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { Group } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async ({ params, locals }) => {
    let group: Group;

    if (params.groupID == "last") {
        group = await locals.pb
            .collection("groups")
            .getFirstListItem<Group>("", { sort: "-created", perPage: 1 });
        throw redirect(307, `/groups/${group.id}`);
    } else {
        group = await locals.pb
            .collection("groups")
            .getOne<Group>(params.groupID, { expand: "org" });
    }

    return { group };
};

export const actions: Actions = {
    update: async ({ request, locals, params }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("groups").update(params.groupID, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    delete: async ({ locals, params }) => {
        try {
            await locals.pb.collection("groups").delete(params.groupID);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
