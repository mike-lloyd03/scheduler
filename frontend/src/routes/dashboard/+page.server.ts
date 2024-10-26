import type { Role } from "$lib/types";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
    const roles = await locals.pb.collection("roles").getFullList<Role>({
        filter: `assigned_to='${locals.currentUser?.id}' && event.datetime>=@now`,
        expand: "role_template,event.event_template",
        sort: "event.datetime",
    });

    return { roles };
};
