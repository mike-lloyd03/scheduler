<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import toast from "svelte-french-toast";

    let loading = false;

    const submitLogin: SubmitFunction = () => {
        loading = true;
        return async ({ result, update }) => {
            switch (result.type) {
                case "success":
                    await update();
                    break;
                case "failure":
                    toast.error(result.data?.message || "Something went wrong");
                    await update();
                    break;
                case "error":
                    toast.error(result.error.message);
                    break;
                default:
                    await update();
            }
            loading = false;
        };
    };
</script>

<div class="max-w-xl">
    <h2 class="h2">Login</h2>
    <form method="POST" use:enhance={submitLogin}>
        <label class="label" for="identity">Username/Email</label>
        <input class="input" type="text" id="identity" name="identity" />

        <label class="label" for="password">Password</label>
        <input class="input" type="password" id="password" name="password" />

        <button type="submit" class="variant-filled btn" disabled={loading}>Login</button>
    </form>
</div>
