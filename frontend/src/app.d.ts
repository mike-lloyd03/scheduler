import type { Org, Group, EventTemplate, RoleTemplate, User } from "$lib/types";
import type { AuthModel } from "pocketbase";
import type PocketBase from "pocketbase";

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
    namespace App {
        // interface Error {}
        interface Locals {
            pb: PocketBase;
            currentUser: AuthModel | undefined;
        }
        interface PageData {
            currentUser: AuthModel | undefined;
            orgs?: Org[];
            org?: Org;
            groups?: Group[];
            group?: Group;
            eventTemplates?: EventTemplate[];
            eventTemplate?: EventTemplate;
            roleTemplates?: RoleTemplate[];
            roleTemplate?: RoleTemplate;
            users?: User[];
            user?: User;
        }
        // interface PageState {}
        // interface Platform {}
    }
}

export {};
