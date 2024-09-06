import { redirect, type RequestHandler } from "@sveltejs/kit";

export const POST: RequestHandler = ({ locals }) => {
    locals.pb.authStore.clear();
    locals.currentUser = undefined;

    throw redirect(303, "/login");
};
