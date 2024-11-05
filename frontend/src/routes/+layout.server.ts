import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = ({ locals, url }) => {
    if (locals.currentUser) {
        return {
            currentUser: locals.currentUser,
            currentUserPermissions: locals.currentUserPermissions ?? undefined,
        };
    } else if (url.pathname != "/login") {
        return redirect(307, "/login");
    }

    return {
        currentUser: undefined,
        currentUserRole: undefined,
    };
};
