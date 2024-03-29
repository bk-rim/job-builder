
import './App.css'
import { Card } from '@/components/ui/card'
import { CreateJobBtn } from './components/createJob'
import { JobTable } from './components/jobsTable'
import { ExecutionTable } from './components/executionTable'
import { useEffect, useState } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@/components/ui/tabs"
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import { useJobs } from './api/jobs.api'
import { Inbox } from 'lucide-react'

function App() {
  const queryClient = useQueryClient();
  const [tab, setTab] = useState("jobs");

  const onTabChange = (value: string) => {
    setTab(value);
  }

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080/ws'); 

    socket.onopen = () => {
      console.log('WebSocket connected');
    };

    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log('Received message:', message);

      if (message.MessageType === 'job_updated') {

        queryClient.invalidateQueries({queryKey:['jobs']}); 
      }
    };

    return () => {
      socket.close();
    };
  }, [queryClient]);
  
  const { isLoading, error, data } = useJobs();

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>An error has occurred: {error.message}</div>
  
  return (
    <div>
      <ToastContainer />
      <div className='flex justify-between my-8'>
        <h1 className='text-3xl text-center'>Dashboard</h1>
        <CreateJobBtn onTabChange={(value: string) => onTabChange(value)}/>
      </div>
      {!((data?.jobs?.length ?? 0) > 0 || (data?.executions?.length ?? 0) > 0) ? (
        <Card className='h-[calc(100vh-30rem)]'>
          <div className='flex flex-col justify-center items-center h-full'>
            <Inbox size={48} color="#c0bfbc" strokeWidth={1} />
            <label className='text-center text-gray-400'>No job created for the moment</label>
          </div>
        </Card>
      ): (
      
        <Tabs value={tab} onValueChange={onTabChange}>
          <TabsList >
            <TabsTrigger value='jobs'>Jobs</TabsTrigger>
            <TabsTrigger value='executions'>Executions ({data?.executions.length})</TabsTrigger>
          </TabsList>
          <TabsContent value='jobs'>
            <JobTable dataTable={data?.jobs} onTabChange={(value: string) => onTabChange(value)}/>
          </TabsContent>
          <TabsContent value='executions'>
            <ExecutionTable dataTable={data?.executions}/>
          </TabsContent>
        </Tabs>
      )}
      
      
    </div>
  )
}

export default App
