import { defineConfig } from "cypress";

export default defineConfig({
  e2e: {
    baseUrl: "http://localhost:3000",
    setupNodeEvents(on, config) {
      // implement node event listeners here
      console.log("Setting up Cypress node event listeners...", on, config);
    },
    supportFile: false,
  },
});
