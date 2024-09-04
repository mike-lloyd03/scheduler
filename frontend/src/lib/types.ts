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

export type RoleTemplate = {
    id: string;
    name: string;
    collectionId: string;
    collectionName: string;
    created: string;
    updated: string;
    event_template: string;
    expand?: { event_template: EventTemplate };
};
