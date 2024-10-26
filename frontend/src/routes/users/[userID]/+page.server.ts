import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import type { Group, Org, Permission, User } from "$lib/types";
import type { ClientResponseError } from "pocketbase";

export const load: PageServerLoad = async ({ locals, params }) => {
    let user: User;

    if (params.userID == "last") {
        user = await locals.pb
            .collection("users")
            .getFirstListItem<User>("", { sort: "-created", perPage: 1 });
        throw redirect(307, `/users/${user.id}`);
    } else {
        user = await locals.pb
            .collection("users")
            .getOne<User>(params.userID, { expand: "orgs,groups" });
    }

    const orgs = await locals.pb.collection("orgs").getFullList<Org>({ sort: "name" });
    const groups = await locals.pb.collection("groups").getFullList<Group>({ sort: "name" });
    const permissions = await locals.pb
        .collection("permissions")
        .getFullList<Permission>({ filter: `user='${user.id}'`, expand: "org,group" });

    return { user, orgs, groups, permissions };
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
