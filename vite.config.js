import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
  build: {
    lib: {
      entry: "client/application.ts",
      name: "application",
    },
    rollupOptions: {
      output: {
        dir: "staticfileserver/dist",
        entryFileNames: "application.js",
        assetFileNames: "application.css",
        chunkFileNames: "chunk.js",
        manualChunks: undefined,
      },
    },
  },
  css: {
    postcss: "./postcss.config.js",
  },
  feature: {},
  plugins: [],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "client"),
      "@Components": path.resolve(__dirname, "client/components"),
    },
    extensions: [".js", ".jsx", ".ts", ".tsx", ".css"],
  },
  optimizeDeps: {
    include: [],
    exclude: [],
  },
});
