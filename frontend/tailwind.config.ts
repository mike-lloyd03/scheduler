import type { Config } from "tailwindcss";
import { skeleton } from "@skeletonlabs/tw-plugin";
import forms from "@tailwindcss/forms";
import { join } from "path";
import { customTheme } from "./theme";

export default {
    darkMode: "class",
    content: [
        "./src/**/*.{html,js,svelte,ts}",
        join(require.resolve("@skeletonlabs/skeleton"), "../**/*.{html,js,svelte,ts}"),
    ],

    theme: {
        extend: {},
    },

    plugins: [
        skeleton({
            themes: { custom: [customTheme] },
        }),
        forms,
    ],
} as Config;
