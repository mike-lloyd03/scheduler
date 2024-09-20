import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { Org } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async ({ locals, params }) => {
    let org: Org;

    if (params.orgID == "lastOrg") {
        org = await locals.pb
            .collection("orgs")
            .getFirstListItem<Org>("", { sort: "-created", perPage: 1 });
        throw redirect(307, `/orgs/${org.id}`);
    } else {
        org = await locals.pb.collection("orgs").getOne<Org>(params.orgID);
    }

    return { org };
};

export const actions: Actions = {
    update: async ({ request, locals, params }) => {
        const data = await request.formData();

        try {
            await locals.pb.collection("orgs").update(params.orgID, data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },

    delete: async ({ locals, params }) => {
        try {
            await locals.pb.collection("orgs").delete(params.orgID);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
