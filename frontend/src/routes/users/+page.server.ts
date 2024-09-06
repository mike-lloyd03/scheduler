import type { PageServerLoad } from "./$types";
import type { User } from "$lib/types";

export const load: PageServerLoad = async (event) => {
    const users = await event.locals.pb
        .collection("users")
        .getFullList<User>({ expand: "orgs,groups" });
    return { users };
};
