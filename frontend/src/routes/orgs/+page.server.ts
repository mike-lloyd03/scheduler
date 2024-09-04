import type { PageServerLoad } from "./$types";
import type { Org } from "$lib/types";

export const load: PageServerLoad = async (event) => {
    const orgs = await event.locals.pb.collection("orgs").getFullList<Org>();
    return { orgs };
};
