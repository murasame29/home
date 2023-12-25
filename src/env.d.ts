/// <reference types="astro/client" />

interface ImportMetaEnv {
  readonly VITE_NR_ACCOUNT_ID: string
  readonly VITE_NR_TRUST_KEY: string
  readonly VITE_NR_AGENT_ID: string
  readonly VITE_NR_LICENSE_KEY: string
  readonly VITE_NR_APPLICATION_ID: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}