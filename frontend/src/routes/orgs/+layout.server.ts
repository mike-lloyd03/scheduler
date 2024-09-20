import type { Org } from "$lib/types";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ locals }) => {
    const orgs = await locals.pb.collection("orgs").getFullList<Org>();
    return { orgs };
};
