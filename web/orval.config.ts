import { defineConfig } from 'orval'
export default defineConfig({
  api: {
    input: '../api/openapi/openapi.yaml',
    output: {
      target: 'src/api/generated/api.ts',
      client: 'fetch',
      mode: 'single',
      clean: true,
    },
  },
})
