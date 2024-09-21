import type { Group } from "$lib/types";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ locals }) => {
    const groups = await locals.pb.collection("groups").getFullList<Group>({ sort: "name" });

    return { groups };
};
