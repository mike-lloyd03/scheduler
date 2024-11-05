import PocketBase from "pocketbase";
import type { Handle } from "@sveltejs/kit";
import type { Permission, User } from "$lib/types";
import { getRole } from "$lib/utils";

export const handle: Handle = async ({ event, resolve }) => {
    event.locals.pb = new PocketBase("http://localhost:8090");
    event.locals.pb.authStore.loadFromCookie(event.request.headers.get("cookie") || "");

    try {
        if (event.locals.pb.authStore.isValid) {
            await event.locals.pb.collection("users").authRefresh();
            event.locals.currentUser = structuredClone(event.locals.pb.authStore.model as User);

            const permissions = await event.locals.pb
                .collection("permissions")
                .getFullList<Permission>({ filter: `user="${event.locals.currentUser.id}"` });
            event.locals.currentUserRole = getRole(permissions);
        }
    } catch (_) {
        event.locals.pb.authStore.clear();
        event.locals.currentUser = undefined;
        event.locals.currentUserRole = undefined;
    }

    const response = await resolve(event);

    response.headers.set("set-cookie", event.locals.pb.authStore.exportToCookie({ secure: false }));

    return response;
};
