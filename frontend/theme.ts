import type { CustomThemeConfig } from "@skeletonlabs/tw-plugin";

export const customTheme: CustomThemeConfig = {
    name: "customTheme",
    properties: {
        // =~= Theme Properties =~=
        "--theme-font-family-base": `system-ui`,
        "--theme-font-family-heading": `system-ui`,
        "--theme-font-color-base": "0 0 0",
        "--theme-font-color-dark": "255 255 255",
        "--theme-rounded-base": "9999px",
        "--theme-rounded-container": "8px",
        "--theme-border-base": "1px",
        // =~= Theme On-X Colors =~=
        "--on-primary": "0 0 0",
        "--on-secondary": "255 255 255",
        "--on-tertiary": "0 0 0",
        "--on-success": "0 0 0",
        "--on-warning": "0 0 0",
        "--on-error": "255 255 255",
        "--on-surface": "255 255 255",
        // =~= Theme Colors  =~=
        // primary | #1c71d8
        "--color-primary-50": "221 234 249", // #ddeaf9
        "--color-primary-100": "210 227 247", // #d2e3f7
        "--color-primary-200": "198 220 245", // #c6dcf5
        "--color-primary-300": "164 198 239", // #a4c6ef
        "--color-primary-400": "96 156 228", // #609ce4
        "--color-primary-500": "28 113 216", // #1c71d8
        "--color-primary-600": "25 102 194", // #1966c2
        "--color-primary-700": "21 85 162", // #1555a2
        "--color-primary-800": "17 68 130", // #114482
        "--color-primary-900": "14 55 106", // #0e376a
        // secondary | #5e5c64
        "--color-secondary-50": "231 231 232", // #e7e7e8
        "--color-secondary-100": "223 222 224", // #dfdee0
        "--color-secondary-200": "215 214 216", // #d7d6d8
        "--color-secondary-300": "191 190 193", // #bfbec1
        "--color-secondary-400": "142 141 147", // #8e8d93
        "--color-secondary-500": "94 92 100", // #5e5c64
        "--color-secondary-600": "85 83 90", // #55535a
        "--color-secondary-700": "71 69 75", // #47454b
        "--color-secondary-800": "56 55 60", // #38373c
        "--color-secondary-900": "46 45 49", // #2e2d31
        // tertiary | #deddda
        "--color-tertiary-50": "250 250 249", // #fafaf9
        "--color-tertiary-100": "248 248 248", // #f8f8f8
        "--color-tertiary-200": "247 247 246", // #f7f7f6
        "--color-tertiary-300": "242 241 240", // #f2f1f0
        "--color-tertiary-400": "232 231 229", // #e8e7e5
        "--color-tertiary-500": "222 221 218", // #deddda
        "--color-tertiary-600": "200 199 196", // #c8c7c4
        "--color-tertiary-700": "167 166 164", // #a7a6a4
        "--color-tertiary-800": "133 133 131", // #858583
        "--color-tertiary-900": "109 108 107", // #6d6c6b
        // success | #2ec27e
        "--color-success-50": "224 246 236", // #e0f6ec
        "--color-success-100": "213 243 229", // #d5f3e5
        "--color-success-200": "203 240 223", // #cbf0df
        "--color-success-300": "171 231 203", // #abe7cb
        "--color-success-400": "109 212 165", // #6dd4a5
        "--color-success-500": "46 194 126", // #2ec27e
        "--color-success-600": "41 175 113", // #29af71
        "--color-success-700": "35 146 95", // #23925f
        "--color-success-800": "28 116 76", // #1c744c
        "--color-success-900": "23 95 62", // #175f3e
        // warning | #f5c211
        "--color-warning-50": "254 246 219", // #fef6db
        "--color-warning-100": "253 243 207", // #fdf3cf
        "--color-warning-200": "253 240 196", // #fdf0c4
        "--color-warning-300": "251 231 160", // #fbe7a0
        "--color-warning-400": "248 212 88", // #f8d458
        "--color-warning-500": "245 194 17", // #f5c211
        "--color-warning-600": "221 175 15", // #ddaf0f
        "--color-warning-700": "184 146 13", // #b8920d
        "--color-warning-800": "147 116 10", // #93740a
        "--color-warning-900": "120 95 8", // #785f08
        // error | #c01c28
        "--color-error-50": "246 221 223", // #f6dddf
        "--color-error-100": "242 210 212", // #f2d2d4
        "--color-error-200": "239 198 201", // #efc6c9
        "--color-error-300": "230 164 169", // #e6a4a9
        "--color-error-400": "211 96 105", // #d36069
        "--color-error-500": "192 28 40", // #c01c28
        "--color-error-600": "173 25 36", // #ad1924
        "--color-error-700": "144 21 30", // #90151e
        "--color-error-800": "115 17 24", // #731118
        "--color-error-900": "94 14 20", // #5e0e14
        // surface | #3d3846
        "--color-surface-50": "226 225 227", // #e2e1e3
        "--color-surface-100": "216 215 218", // #d8d7da
        "--color-surface-200": "207 205 209", // #cfcdd1
        "--color-surface-300": "177 175 181", // #b1afb5
        "--color-surface-400": "119 116 126", // #77747e
        "--color-surface-500": "61 56 70", // #3d3846
        "--color-surface-600": "55 50 63", // #37323f
        "--color-surface-700": "46 42 53", // #2e2a35
        "--color-surface-800": "37 34 42", // #25222a
        "--color-surface-900": "30 27 34", // #1e1b22
    },
};
