<script lang="ts">
    import EventsTable from "$lib/components/EventsTable.svelte";
    import type { PageData } from "./$types";
    import { getModalStore, type ModalComponent, type ModalSettings } from "@skeletonlabs/skeleton";
    import ScheduleEventForm from "./ScheduleEventForm.svelte";
    import { runAction, jsonToFormString } from "$lib/utils";
    import { invalidateAll } from "$app/navigation";

    const modalStore = getModalStore();

    function modalComponentForm(): void {
        const c: ModalComponent = { ref: ScheduleEventForm };
        const modal: ModalSettings = {
            type: "component",
            component: c,
            meta: { eventTemplates: data.eventTemplates },
            response: (r) => {
                if (r) {
                    runAction("?/createEvent", jsonToFormString(r));
                }
            },
        };
        modalStore.trigger(modal);
        invalidateAll();
    }

    export let data: PageData;
</script>

<button
    class="variant-ringed-primary btn hover:variant-filled-primary"
    on:click={modalComponentForm}>Schedule Event</button
>
<EventsTable events={data.events} roles={data.roles} enableNav />
