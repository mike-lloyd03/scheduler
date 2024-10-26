import type {
    Org,
    Group,
    EventTemplate,
    RoleTemplate,
    User,
    Event,
    Role,
    Permission,
} from "$lib/types";
import type PocketBase from "pocketbase";

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
    namespace App {
        // interface Error {}
        interface Locals {
            pb: PocketBase;
            currentUser: User | undefined;
        }
        interface PageData {
            currentUser: User | undefined;
            orgs?: Org[];
            org?: Org;
            groups?: Group[];
            group?: Group;
            events?: Event[];
            event?: Event;
            roles?: Role[];
            role?: Role;
            eventTemplates?: EventTemplate[];
            eventTemplate?: EventTemplate;
            roleTemplates?: RoleTemplate[];
            roleTemplate?: RoleTemplate;
            users?: User[];
            user?: User;
            permissions?: Permission[];
        }
        // interface PageState {}
        // interface Platform {}
    }
}

export {};
