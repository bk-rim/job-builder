import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { useState } from "react"
import { z } from "zod"
import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,  
  DialogTrigger,
} from "@/components/ui/dialog"
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
  } from "@/components/ui/form"
  import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select"
import { Plus } from "lucide-react"
import { Separator } from "@/components/ui/separator"
import { Input } from "../ui/input"
import { X } from "lucide-react"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { createJob } from "../../api/jobs.api"
import { notify } from "../../utils/notification"

const createJobSchema = z.object({
    name: z.string().nonempty("Ce champ est requis"),
    type: z.enum(["bridge_status", "weather"], {
      required_error: "Ce champ est requis",
    }),
    frequency: z.enum(["daily", "weekly"], {
        required_error: "Ce champ est requis",
    }),
    webhook_slack: z.string().url("Ce champ doit être une URL"),
})

export function CreateJobBtn({onTabChange}: {onTabChange: (tab: string) => void}) {
    const [open, setOpen] = useState(false)
    const queryClient = useQueryClient()

    const mutation = useMutation({
        mutationFn: createJob,
        onSuccess: () => {
          onTabChange('jobs')
          notify("New job created")
          queryClient.invalidateQueries({queryKey: ["jobs"]}) 
        },
    })

    const form = useForm({
        resolver: zodResolver(createJobSchema),
        defaultValues: {
            name: "",
            type: "",
            frequency: "Daily",
            webhook_slack: "",
        },
    })

    function onSubmit (values: {name: string, type: string, frequency: string, webhook_slack: string}) {
        console.log(values)
        mutation.mutate(values)

        form.reset()
        setOpen(false)
    }
  return (
    <Dialog open={open}>
      <DialogTrigger asChild>
      <Button className='h-9' onClick={() => setOpen(true)}>
          <Plus size={22} color="#ffffff" />
          Create New Job</Button>
      </DialogTrigger>
      <DialogContent className="flex flex-col h-screen">
      <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col w-full space-y-8" >
        
            <div className="flex justify-between pt-4 items-center">
                <div className="flex itmems-center gap-2">
                    <X className="h-4 w-4" onClick={() => setOpen(false)}/>
                    <h1 className="text-lg leading-none tracking-tight">Create job</h1>
                </div>
                
   
                <div className="flex itmems-center gap-1">
                    <Button type="reset" className="h-8" variant="outline" onClick={()=> setOpen(false)}>
                        cancel
                    </Button>
                    <Button className="h-8" type="submit">
                        create
                    </Button>
                </div>
                
            </div>
       
        <div>
            <Separator className="left-0 fixed"/>
        </div>
        
        <div className="flex flex-col space-y-8 mt-8">
        
            <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
                <FormItem>
                <FormLabel>Name</FormLabel>
                {form.formState.errors.name && <div className="text-red-500 text-sm">{form.formState.errors.name.message}</div>}
                <FormControl>
                    <Input placeholder="Mon Pont Chaban" {...field} />
                </FormControl>
                
                </FormItem>
            )}
            />
          
          <FormField
          control={form.control}
          name="type"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Type</FormLabel>
              {form.formState.errors.type && <div className="text-red-500 text-sm">{form.formState.errors.type.message}</div>}
              <Select onValueChange={field.onChange} defaultValue={field.value}>
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select type..." />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectItem value="bridge_status">Mon Pont Chaban</SelectItem>
                  <SelectItem value="weather">Check Weather</SelectItem>
                </SelectContent>
              </Select>
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="frequency"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Frequency</FormLabel>
              {form.formState.errors.frequency && <div className="text-red-500 text-sm">{form.formState.errors.frequency.message}</div>}
              <Select onValueChange={field.onChange} defaultValue={field.value}>
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select frequency..." />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectItem value="daily">Daily</SelectItem>
                  <SelectItem value="weekly">Weekly</SelectItem>
                </SelectContent>
              </Select>
            </FormItem>
          )}
        />

            <FormField
            control={form.control}
            name="webhook_slack"
            render={({ field }) => (
                <FormItem>
                <FormLabel>Webhook</FormLabel>
                {form.formState.errors.webhook_slack && <div className="text-red-500 text-sm">{form.formState.errors.webhook_slack.message}</div>}
                <FormControl>
                    <Input placeholder="https://webhook.site/1b9c4d5e-7e6b-4f8c-9b7e-5f0e3c2f9f0e" {...field} />
                </FormControl>
                </FormItem>
            )}
            />
            
        
        
          
        </div>
        </form>
        </Form>        
      </DialogContent>
    </Dialog>
  )
}
