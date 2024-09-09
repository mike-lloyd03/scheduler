<script lang="ts">
    import EventsTable from "$lib/components/EventsTable.svelte";
    import type { PageData } from "./$types";
    import { getModalStore, type ModalComponent, type ModalSettings } from "@skeletonlabs/skeleton";
    import ScheduleEventForm from "./ScheduleEventForm.svelte";
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import { handleSubmit } from "$lib/utils";

    const modalStore = getModalStore();

    export let data: PageData;

    let createEventForm: HTMLFormElement;
    let createEventFormData = {
        event_template: "",
        datetime: "",
    };

    function modalComponentForm(): void {
        const c: ModalComponent = { ref: ScheduleEventForm };
        const modal: ModalSettings = {
            type: "component",
            component: c,
            meta: { eventTemplates: data.eventTemplates },
            response: (r) => {
                if (r) {
                    createEventFormData = r;
                    createEventForm.requestSubmit();
                }
            },
        };
        modalStore.trigger(modal);
    }

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Successful";

        switch (action.search) {
            case "?/createEvent":
                formData.set("event_template", createEventFormData.event_template);
                formData.set("datetime", createEventFormData.datetime);
                successMsg = "Event scheduled";
                break;
        }

        return handleSubmit(successMsg);
    };
</script>

<form bind:this={createEventForm} method="POST" action="?/createEvent" use:enhance={submit}></form>

<button
    class="variant-ringed-primary btn hover:variant-filled-primary"
    on:click={modalComponentForm}>Schedule Event</button
>

<EventsTable events={data.events} roles={data.roles} enableNav />
