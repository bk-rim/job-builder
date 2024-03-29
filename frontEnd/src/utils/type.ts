export type Jobs = {
    id: string
    name: string
    type: string
    created_on: string
    frequency?: string | null
    executed_on?: string | null
    status?: string | null
    webhook_slack?: string | null
}