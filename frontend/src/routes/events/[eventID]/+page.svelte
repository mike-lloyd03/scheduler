<script lang="ts">
    import { enhance } from "$app/forms";
    import toast from "svelte-french-toast";
    import type { SubmitFunction } from "@sveltejs/kit";
    import InputField from "$lib/InputField.svelte";
    import SelectField from "$lib/SelectField.svelte";
    import type { PageData } from "./$types";

    export let data: PageData;

    var edit = false;
    var newRole = false;
    var loading = false;

    const submitEdit: SubmitFunction = ({ cancel }) => {
        loading = true;

        if (edit == false) {
            edit = true;
            cancel();
            return;
        }

        edit = false;
        return async ({ result, update }) => {
            switch (result.type) {
                case "success":
                    toast.success("Event Template updated");
                    await update();
                    break;
                case "failure":
                    toast.error(result.data?.message || "Something went wrong");
                    break;
                default:
                    await update();
            }
            loading = false;
        };
    };

    const submitNewRole: SubmitFunction = ({ cancel }) => {
        loading = true;

        if (newRole == false) {
            newRole = true;
            cancel();
            return;
        }

        newRole = false;
        return async ({ result, update }) => {
            switch (result.type) {
                case "success":
                    toast.success("Role added");
                    await update();
                    break;
                case "failure":
                    toast.error(result.data?.message || "Something went wrong");
                    break;
                default:
                    await update();
            }
            loading = false;
        };
    };
</script>

<div>
    <form method="POST" use:enhance={submitEdit} action="?/edit">
        <div class="py-4">
            <div>
                <InputField title="Name" value={data.event_template.name} {edit} />
            </div>
            <div>
                <SelectField
                    title="Recurrence"
                    value={data.event_template.recurrence}
                    options={["daily", "weekly", "monthly"]}
                    {edit}
                />
            </div>
            <div>
                <span class="font-bold">Group:</span>
                {data.event_template.expand?.group.name || ""}
            </div>
            <div>
                <span class="font-bold">Created at:</span>
                {data.event_template.created}
            </div>
            <div>
                <span class="font-bold">Updated at:</span>
                {data.event_template.updated}
            </div>
        </div>
        <button type="submit" class="variant-filled-primary btn">{edit ? "Save" : "Edit"}</button>
        {#if edit}
            <button
                type="button"
                class="variant-filled-secondary btn"
                on:click={() => (edit = false)}
            >
                Cancel
            </button>
        {/if}
    </form>

    <div class="py-8">
        <h3 class="h3">Event Roles</h3>
        <button class="variant-filled-primary btn" on:click={() => (newRole = true)}>+</button>
        <div>
            <ul class="list">
                {#each data.role_templates as role_template}
                    <li>
                        <span class="flex-auto">{role_template.name}</span>
                        <form method="POST" action="?/deleteRole" use:enhance>
                            <input type="hidden" name="deleteRoleID" value={role_template.id} />
                            <button class="variant-filled-primary btn">X</button>
                        </form>
                    </li>
                {/each}
                {#if newRole}
                    <li>
                        <form method="POST" action="?/newRole" use:enhance={submitNewRole}>
                            <input
                                class="input"
                                type="text"
                                name="name"
                                id="name"
                                placeholder="Role Name"
                            />
                            <button class="variant-filled-primary btn">Check</button>
                            <button class="variant-filled-secondary btn">X</button>
                        </form>
                    </li>
                {/if}
            </ul>
        </div>
    </div>
</div>
