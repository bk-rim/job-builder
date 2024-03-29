import { useQuery } from '@tanstack/react-query'
import { Jobs } from '@/utils/type'
import { formatDate } from '@/utils/formatDate'

export type JobsResponse = {
    jobs: Jobs[]
    executions: Jobs[]
}
export const fetchJobs = async () => {
    const res = await fetch("http://localhost:8080/jobs",{
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    })
    if (!res.ok) {
      console.error("Error",res)
      throw new Error('Network response was not ok')
    }

    const data = await res.json()
    const jobs = data.filter((job: Jobs) => job.executed_on === null).map((job: Jobs) => {
        job.created_on = formatDate(job.created_on)
        return job

    })

    const executions = data.filter((job: Jobs) => job.executed_on !== null).map((job: Jobs) => {
        job.created_on = formatDate(job.created_on)
        job.executed_on = formatDate(job.executed_on ?? "")
        return job
    })

    const jobsResponse: JobsResponse = {jobs, executions}

    return jobsResponse
    
}

export const useJobs = () => {

    return useQuery<JobsResponse, Error>({queryKey: ["jobs"], queryFn: fetchJobs})
}

export const createJob = async (values: {name: string, type: string, frequency: string, webhook_slack: string}) => {
    const res = await fetch("http://localhost:8080/jobs", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            },
            body: JSON.stringify(values),
        })
    if (!res.ok) {
        console.error("Error",res)
        throw new Error('Network response was not ok')
    }

    return await res.json()
}

export const deleteJob = async (id: string) => {
    const res = await fetch(`http://localhost:8080/jobs/${id}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            },
        })
    if (!res.ok) {
        console.error("Error",res)
        throw new Error('Network response was not ok')
    }

    return await res.json()
}

export const executeJob = async (job: Jobs) => {
    const res = await fetch('http://localhost:8080/job-execution', {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            },
        body: JSON.stringify(job),
    })
    if (!res.ok) {
        console.error("Error",res)
        throw new Error('Network response was not ok')
    }

    return await res.json()

}
