import { writable } from "svelte/store";

function createBreadcrumbs() {
    const { subscribe, set, update } = writable<{ label: string; url: string }[]>([]);

    const clear = () => {
        set([]);
        return store;
    };

    const add = (label: string, url: string) => {
        update((s) => {
            s.push({ label, url });
            return s;
        });
        return store;
    };

    const store = {
        subscribe,
        clear,
        add,
    };

    return store;
}

export const breadcrumbs = createBreadcrumbs();
