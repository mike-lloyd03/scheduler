<script lang="ts">
    import type { SvelteComponent } from "svelte";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import type { EventTemplate, RoleTemplate, User } from "$lib/types";

    export let parent: SvelteComponent;

    const modalStore = getModalStore();

    const formData = {
        oldPassword: "",
        password: "",
        passwordConfirm: "",
    };

    function onFormSubmit(): void {
        if ($modalStore[0].response) {
            $modalStore[0].response(formData);
        }
        modalStore.close();
    }

    const cBase = "card p-4 w-modal shadow-xl space-y-4";
    const cHeader = "text-2xl font-bold";
    const cForm = "border border-surface-500 p-4 space-y-4 rounded-container-token";
</script>

{#if $modalStore[0]}
    <div class={cBase}>
        <header class={cHeader}>Update Password</header>

        <form class="modal-form {cForm}">
            <label class="label">
                <span>Old Password</span>
                <input class="input" bind:value={formData.oldPassword} />
            </label>

            <label class="label">
                <span>New Password</span>
                <input class="input" bind:value={formData.password} />
            </label>

            <label class="label">
                <span>Confirm Password</span>
                <input class="input" bind:value={formData.passwordConfirm} />
            </label>
        </form>

        <footer class="modal-footer {parent.regionFooter}">
            <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>
                {parent.buttonTextCancel}
            </button>

            <button class="btn {parent.buttonPositive}" on:click={onFormSubmit}>
                {parent.buttonTextSubmit}
            </button>
        </footer>
    </div>
{/if}
