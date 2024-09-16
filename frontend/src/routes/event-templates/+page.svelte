<script lang="ts">
    import { goto } from "$app/navigation";
    import type { SubmitFunction } from "@sveltejs/kit";
    import type { PageData } from "./$types";
    import { handleSubmit } from "$lib/utils";
    import { enhance } from "$app/forms";
    import { getModalStore, type ModalComponent, type ModalSettings } from "@skeletonlabs/skeleton";
    import CreateEventTemplateForm from "./CreateEventTemplateForm.svelte";
    import ModalButton from "$lib/components/ModalButton.svelte";

    const modalStore = getModalStore();

    export let data: PageData;

    let createForm: HTMLFormElement;
    let createFormData = {
        name: "",
        recurrence: "",
        group: "",
    };

    function modalComponentForm(): void {
        const c: ModalComponent = { ref: CreateEventTemplateForm };
        const modal: ModalSettings = {
            type: "component",
            component: c,
            meta: {
                groups: data.groups,
            },
            response: (r) => {
                if (r) {
                    createFormData = r;
                    createForm.requestSubmit();
                }
            },
        };
        modalStore.trigger(modal);
    }

    const submit: SubmitFunction = ({ formData }) => {
        formData.set("name", createFormData.name);
        formData.set("recurrence", createFormData.recurrence);
        formData.set("group", createFormData.group);

        return handleSubmit("Event template created");
    };
</script>

<form bind:this={createForm} method="POST" use:enhance={submit}></form>

<ModalButton onClick={modalComponentForm} text="Create Event Template" />

<h2 class="h2">Event Templates</h2>
<div class="table-container">
    <table class="table table-interactive table-fixed text-center">
        <thead>
            <th>Name</th>
            <th>Recurrence</th>
            <th>Group</th>
        </thead>
        <tbody>
            {#each data.eventTemplates as eventTemplate}
                <tr on:click={() => goto(`/event-templates/${eventTemplate.id}`)}>
                    <td>{eventTemplate.name}</td>
                    <td>{eventTemplate.recurrence}</td>
                    <td>{eventTemplate.expand?.group.name || ""}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
