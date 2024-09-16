import type { PageServerLoad } from "./$types";
import type { EventTemplate, Group } from "$lib/types";
import type { ClientResponseError } from "pocketbase";
import { fail, type Actions } from "@sveltejs/kit";

export const load: PageServerLoad = async (event) => {
    const eventTemplates = await event.locals.pb
        .collection("event_templates")
        .getFullList<EventTemplate>({ expand: "group" });

    const groups = await event.locals.pb.collection("groups").getFullList<Group>({ expand: "org" });

    return { eventTemplates, groups };
};

export const actions: Actions = {
    default: async ({ request, locals }) => {
        const data = await request.formData();

        data.forEach((v, k) => console.log(`${v} ${k}`));

        try {
            await locals.pb.collection("event_templates").create(data);
        } catch (e) {
            const error = e as ClientResponseError;
            console.log("Error: ", JSON.stringify(error.data));

            return fail(error.status, { message: error.message });
        }
    },
};
