/// <reference types="astro/client" />

interface ImportMetaEnv {
  readonly NR_ACCOUNT_ID: string
  readonly NR_TRUST_KEY: string
  readonly NR_AGENT_ID: string
  readonly NR_LICENSE_KEY: string
  readonly NR_APPLICATION_ID: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}