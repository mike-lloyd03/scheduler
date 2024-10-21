<script lang="ts">
    import { handleSubmit, modalComponentForm } from "$lib/utils";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import { enhance } from "$app/forms";
    import ModalButton from "$lib/components/ModalButton.svelte";
    import type { SubmitFunction } from "@sveltejs/kit";
    import UpdatePasswordForm from "./UpdatePasswordForm.svelte";

    const modalStore = getModalStore();

    let updatePasswordForm: HTMLFormElement;
    let updatePasswordFormData = {
        oldPassword: "",
        password: "",
        passwordConfirm: "",
    };

    function openModal(): void {
        modalComponentForm(UpdatePasswordForm, updatePasswordForm, modalStore, (r) => {
            updatePasswordFormData = r;
        });
    }

    const submit: SubmitFunction = ({ action, formData }) => {
        let successMsg = "Successful";

        switch (action.pathname + action.search) {
            case "?/update":
                formData.set("oldPassword", updatePasswordFormData.oldPassword);
                formData.set("password", updatePasswordFormData.password);
                formData.set("passwordConfirm", updatePasswordFormData.passwordConfirm);
                successMsg = "Password Updated";
                break;
        }

        return handleSubmit(successMsg);
    };
</script>

<form bind:this={updatePasswordForm} method="POST" action="?/update" use:enhance={submit}></form>

<ModalButton onClick={openModal} text="Update Password" />
