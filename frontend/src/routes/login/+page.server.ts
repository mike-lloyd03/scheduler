import { error, fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions: Actions = {
    default: async ({ request, locals }) => {
        const data = await request.formData();

        const identity = data.get("identity");
        const password = data.get("password");

        if (identity == null || password == null) {
            return fail(400, {});
        }

        try {
            await locals.pb.collection("users").authWithPassword(identity, password);
            if (!locals.pb?.authStore?.model?.verified) {
                locals.pb.authStore.clear();
                return {
                    notVerified: true,
                };
            }
        } catch (err) {
            console.log("Error: ", err);
            throw error(err.status, err.message);
        }

        throw redirect(303, "/");
    },
};
