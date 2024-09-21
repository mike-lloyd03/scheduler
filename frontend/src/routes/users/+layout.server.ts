import type { User } from "$lib/types";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ locals }) => {
    const users = await locals.pb
        .collection("users")
        .getFullList<User>({ expand: "orgs,groups", sort: "name" });
    return { users };
};
