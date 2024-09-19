<script lang="ts">
    import EventsTable from "$lib/components/EventsTable.svelte";
    import type { PageData } from "./$types";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import ScheduleEventForm from "./ScheduleEventForm.svelte";
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import { handleSubmit, modalComponentForm } from "$lib/utils";
    import ModalButton from "$lib/components/ModalButton.svelte";

    const modalStore = getModalStore();

    export let data: PageData;

    let createEventForm: HTMLFormElement;
    let createEventFormData = {
        event_template: "",
        datetime: "",
        roles: "",
    };

    function openModal(): void {
        const meta = {
            eventTemplates: data.eventTemplates,
            roleTemplates: data.roleTemplates,
            users: data.users,
        };

        modalComponentForm(
            ScheduleEventForm,
            createEventForm,
            modalStore,
            (r) => {
                createEventFormData = r;
            },
            meta,
        );
    }

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Successful";

        switch (action.search) {
            case "?/createEvent":
                formData.set("event_template", createEventFormData.event_template);
                formData.set("datetime", createEventFormData.datetime);
                formData.set("roles", createEventFormData.roles);
                successMsg = "Event scheduled";
                break;
        }

        return handleSubmit(successMsg);
    };
</script>

<form bind:this={createEventForm} method="POST" action="?/createEvent" use:enhance={submit}></form>

<ModalButton onClick={openModal} text="Schedule Event" />

<EventsTable events={data.events} roles={data.roles} enableNav />
