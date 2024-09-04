import type { Org, Group, EventTemplate, RoleTemplate } from "$lib/types";
import type { AuthModel } from "pocketbase";
import type PocketBase from "pocketbase";

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
    namespace App {
        // interface Error {}
        interface Locals {
            pb: PocketBase;
            user: AuthModel | undefined;
        }
        interface PageData {
            orgs?: Org[];
            org?: Group;
            groups?: Group[];
            group?: Group;
            event_templates?: EventTemplate[];
            event_template?: EventTemplate;
            role_templates?: RoleTemplate[];
            role_template?: RoleTemplate;
        }
        // interface PageState {}
        // interface Platform {}
    }
}

export {};
