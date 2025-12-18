import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig, searchForWorkspaceRoot } from "vite";
import wails from "@wailsio/runtime/plugins/vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [tailwindcss(), sveltekit(), wails("./bindings")],

  server: {
    fs: {
      allow: [
        // search up for workspace root
        searchForWorkspaceRoot(process.cwd()),

        // your custom rules
        "./bindings/*"
      ]
    }
  }
});
