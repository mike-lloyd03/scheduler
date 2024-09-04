<script lang="ts">
    import { enhance } from "$app/forms";
    import toast from "svelte-french-toast";

    var loading = false;

    const submitLogin = () => {
        loading = true;
        return async ({ result, update }) => {
            switch (result.type) {
                case "success":
                    await update();
                    break;
                case "invalid":
                    toast.error("Invalid credentials");
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
