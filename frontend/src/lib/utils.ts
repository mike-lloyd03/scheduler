import { invalidateAll } from "$app/navigation";
import toast from "svelte-french-toast";
import { type ModalComponent, type ModalSettings, type ModalStore } from "@skeletonlabs/skeleton";
import { DateTime } from "luxon";
import { CurrentUserRole, UserRole, type Group, type Org, type Permission } from "./types";

export function handleSubmit(successMsg: string, onSuccess: (() => void) | undefined = undefined) {
    return async ({ result, update }) => {
        switch (result.type) {
            case "success":
                toast.success(successMsg);
                await update();
                if (onSuccess) {
                    onSuccess();
                }
                break;
            case "failure":
                toast.error(result.data?.message || "Something went wrong");
                break;
            default:
                await update();
        }
    };
}

export function runAction(action: string, body?: string) {
    fetch(action, {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: body,
    }).then(() => invalidateAll());
}

export function jsonToFormString(jsonData: object): string {
    const formData: string[] = [];

    Object.entries(jsonData).forEach(([k, v]) => formData.push(`${k}=${v}`));

    console.log("jsonToForm result: ");

    return formData.join("&");
}

export function modalComponentForm(
    component: any,
    form: HTMLFormElement,
    modalStore: ModalStore,
    presubmit?: (r: any) => void,
    meta?: object,
): void {
    const c: ModalComponent = { ref: component };

    const modal: ModalSettings = {
        type: "component",
        component: c,
        meta,
        response: (r) => {
            if (r) {
                if (presubmit) {
                    presubmit(r);
                }
                form.requestSubmit();
            }
        },
    };
    modalStore.trigger(modal);
}

export function toLocaleDateTime(dateString: string): string {
    const date = DateTime.fromSQL(dateString);
    return date.toLocaleString(DateTime.DATETIME_SHORT);
}

export function toLocaleDate(dateString: string): string {
    const date = DateTime.fromSQL(dateString);
    return date.toLocaleString(DateTime.DATE_SHORT);
}

export function getRole(permissions: Permission[]): CurrentUserRole {
    const roleNames = [
        CurrentUserRole.Member,
        CurrentUserRole.GroupManager,
        CurrentUserRole.GroupAdmin,
        CurrentUserRole.OrgManager,
        CurrentUserRole.OrgAdmin,
    ];
    let roleValue = 0;

    for (const p of permissions) {
        if (p.org != "") {
            if (p.role == "admin") {
                roleValue = 4;
                break;
            } else if (p.role == "manager" && roleValue < 4) {
                roleValue = 3;
            }
        } else if (p.group != "") {
            if (p.role == "admin" && roleValue < 3) {
                roleValue = 2;
            } else if (p.role == "manager" && roleValue < 2) {
                roleValue = 1;
            }
        }
    }

    return roleNames[roleValue];
}

export function hasOrgRole(permissions: Permission[], org: Org, role: UserRole): boolean {
    return permissions.some((p) => p.org == org.id && p.role == role);
}

export function hasGroupRole(permissions: Permission[], group: Group, role: UserRole): boolean {
    return permissions.some((p) => p.group == group.id && p.role == role);
}
