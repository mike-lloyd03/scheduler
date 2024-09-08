export enum UserRole {
    Admin = "admin",
    Manager = "manager",
    User = "user",
}

export enum Recurrence {
    Daily = "daily",
    Weekly = "weekly",
    Monthly = "monthly",
}

export type Org = {
    id: string;
    name: string;
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
};

export type Group = {
    id: string;
    name: string;
    org: string;
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
    expand?: { org: Org };
};

export type EventTemplate = {
    id: string;
    name: string;
    recurrence: Recurrence;
    group: string;
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
    expand?: { group: Group };
};

export type Event = {
    id: string;
    event_template: string;
    datetime: string;
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
    expand?: { event_template: EventTemplate };
};

export type RoleTemplate = {
    id: string;
    name: string;
    description?: string;
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
    event_template: string;
    expand?: { event_template: EventTemplate };
};

export type Role = {
    id: string;
    role_template: string;
    event: string;
    assigned_to?: string;
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
    expand?: { role_template: RoleTemplate; event: Event; assigned_to: User };
};

export type User = {
    id: string;
    username: string;
    verified: false;
    emailVisibility: true;
    email: string;
    name: string;
    role: string;
    orgs: string[];
    groups: string[];
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
    expand?: { orgs: Org[]; groups: Group[] };
};
