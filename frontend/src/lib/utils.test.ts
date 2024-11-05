import { expect, test } from "vitest";
import { getRole } from "./utils";
import { CurrentUserRole, UserRole, type Permission } from "./types";

test("getRole", () => {
    let permissions = [
        {
            org: "",
            group: "group1",
            role: UserRole.Admin,
        },
    ];

    expect(getRole(permissions as Permission[])).toEqual(CurrentUserRole.GroupAdmin);

    permissions = [
        {
            org: "",
            group: "group1",
            role: UserRole.Admin,
        },
        {
            org: "org1",
            group: "",
            role: UserRole.Admin,
        },
    ];

    expect(getRole(permissions as Permission[])).toEqual(CurrentUserRole.OrgAdmin);

    permissions = [
        {
            org: "org1",
            group: "",
            role: UserRole.Manager,
        },
        {
            org: "",
            group: "group1",
            role: UserRole.Admin,
        },
    ];

    expect(getRole(permissions as Permission[])).toEqual(CurrentUserRole.OrgManager);

    permissions = [
        {
            org: "org1",
            group: "",
            role: UserRole.Admin,
        },
        {
            org: "",
            group: "group1",
            role: UserRole.Manager,
        },
    ];

    expect(getRole(permissions as Permission[])).toEqual(CurrentUserRole.OrgAdmin);

    permissions = [
        {
            org: "org1",
            group: "",
            role: UserRole.Manager,
        },
        {
            org: "",
            group: "group1",
            role: UserRole.Manager,
        },
    ];

    expect(getRole(permissions as Permission[])).toEqual(CurrentUserRole.OrgManager);

    permissions = [];

    expect(getRole(permissions as Permission[])).toEqual(CurrentUserRole.Member);
});
