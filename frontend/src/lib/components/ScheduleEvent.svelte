<script lang="ts">
    import { handleSubmit, modalComponentForm } from "$lib/utils";
    import ScheduleEventForm from "./ScheduleEventForm.svelte";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import type { SubmitFunction } from "@sveltejs/kit";
    import { enhance } from "$app/forms";
    import ModalButton from "$lib/components/ModalButton.svelte";
    import type { EventTemplate, RoleTemplate, User } from "$lib/types";

    export let eventTemplates: EventTemplate[];
    export let roleTemplates: RoleTemplate[];
    export let users: User[];
    export let selectedEventTemplateID: string | undefined = undefined;

    const modalStore = getModalStore();

    let createEventForm: HTMLFormElement;
    let createEventFormData = {
        event_template: "",
        datetime: "",
        roles: "",
    };

    function openModal(): void {
        const meta = {
            eventTemplates: eventTemplates,
            roleTemplates: roleTemplates,
            users: users,
            selectedEventTemplateID: selectedEventTemplateID,
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

        switch (action.pathname + action.search) {
            case "/events?/createEvent":
                formData.set("event_template", createEventFormData.event_template);
                formData.set("datetime", createEventFormData.datetime);
                formData.set("roles", createEventFormData.roles);
                successMsg = "Event scheduled";
                break;
        }

        return handleSubmit(successMsg);
    };
</script>

<form
    bind:this={createEventForm}
    method="POST"
    action="/events?/createEvent"
    use:enhance={submit}
></form>

<ModalButton onClick={openModal} text="Schedule Event" />
