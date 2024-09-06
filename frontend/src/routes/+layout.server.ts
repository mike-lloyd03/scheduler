import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = ({ locals }) => {
    if (locals.currentUser) {
        return {
            currentUser: locals.currentUser,
        };
    }

    return {
        currentUser: undefined,
    };
};
