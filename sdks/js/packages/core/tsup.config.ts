import cssModulesPlugin from 'esbuild-css-modules-plugin';
import { defineConfig } from 'tsup';

export default defineConfig(() => [
  // Core API
  {
    entry: ['src/index.ts'],
    format: ['cjs', 'esm'],
    dts: true
  },
  // React APIs
  {
    entry: ['react/index.ts'],
    outDir: 'react/dist',
    banner: {
      js: "'use client'"
    },
    format: ['cjs', 'esm'],
    external: ['react', 'svelte', 'vue', 'solid-js', 'client/*'],
    dts: true,
    loader: {
      '.svg': 'dataurl',
      '.png': 'dataurl'
    },
    esbuildPlugins: [cssModulesPlugin()]
  },
  // client API
  {
    entry: ['ApiClient/index.ts'],
    format: ['cjs', 'esm'],
    outDir: 'client/dist',
    dts: true,
  },
]);
