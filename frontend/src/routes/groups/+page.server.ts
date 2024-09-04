import type { PageServerLoad } from "./$types";
import type { Group } from "$lib/types";

export const load: PageServerLoad = async (event) => {
    const groups = await event.locals.pb.collection("groups").getFullList<Group>({ expand: "org" });
    return { groups };
};
