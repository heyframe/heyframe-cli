import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import eslintConfigPrettier from "eslint-config-prettier";

import frontendRules from "@heyframe-ag/frontend-eslint-rules";

/** @type {import('eslint').Linter.Config[]} */
export default [
	{ languageOptions: { globals: globals.browser } },
	pluginJs.configs.recommended,
	...tseslint.configs.recommended,
	frontendRules,
	eslintConfigPrettier,
	{
		rules: {
			"@typescript-eslint/no-unused-vars": "warn",
			"@typescript-eslint/no-unused-expressions": "warn",
			"@typescript-eslint/no-this-alias": "warn",
			"@typescript-eslint/no-require-imports": "off",
			"no-undef": "off",
			"no-alert": "error",
			"no-console": ["error", { allow: ["warn", "error"] }],
		},
	},
];
