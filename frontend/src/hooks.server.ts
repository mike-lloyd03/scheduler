import PocketBase from "pocketbase";
import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    event.locals.pb = new PocketBase("http://localhost:8090");
    event.locals.pb.authStore.loadFromCookie(event.request.headers.get("cookie") || "");

    try {
        if (event.locals.pb.authStore.isValid) {
            await event.locals.pb.collection("users").authRefresh();
            event.locals.currentUser = structuredClone(event.locals.pb.authStore.model);
        }
    } catch (_) {
        event.locals.pb.authStore.clear();
        event.locals.currentUser = undefined;
    }

    const response = await resolve(event);

    response.headers.set("set-cookie", event.locals.pb.authStore.exportToCookie({ secure: false }));

    return response;
};
