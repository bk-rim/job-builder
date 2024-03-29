import * as React from "react"
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table"

import { Button } from "@/components/ui/button"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Jobs } from "../../utils/type"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { deleteJob } from "../../api/jobs.api"
import { notify } from "../../utils/notification"
import { formatDate } from "../../utils/formatDate"
import 'react-toastify/dist/ReactToastify.css';
import { CircleAlert, CircleCheck, CircleX, Loader } from "lucide-react"


const getColumns = (deleteFn: (id: string) => any): ColumnDef<Jobs>[] => ([
  
  {
    accessorKey: "name",
    header: () => <div className="text-left">Executions</div>,
    cell: ({ row }) => (
      <div className="text-left ">{row.getValue("name")}</div>
    ),
  },
  {
    accessorKey: "executed_on",
    header: () => <div className="text-left">Date</div>,
    cell: ({ row }) => <div className="text-left font-medium text-gray-400">{formatDate(row.getValue("executed_on"))}</div>,
  },
  {
    accessorKey: "status",
    header: () => <div className="text-left">Status</div>,
    cell: ({ row }) => {
     return row.getValue("status") === "executed" ? 
        (<div className="text-left ">
            <span className="flex items-center w-24 justify-center gap-1 bg-green-100 border px-2 rounded-xl">
                <CircleCheck size={14} color="#2ec27e" strokeWidth={3} />
                {row.getValue("status")}
            </span>
          </div>) : row.getValue("status") === "pending" ? 
        (<div className="text-left">
            <span className="flex items-center w-24 justify-center gap-1 bg-yellow-200 border px-2 rounded-xl">
                <CircleAlert size={14} color="#e5a50a" strokeWidth={3} />
                {row.getValue("status")}
            </span>
        </div>): row.getValue("status") === "on going" ? 
        (<div className="text-left ">
            <span className="flex items-center w-24 justify-center gap-1 bg-blue-300 border px-2 rounded-xl">
                <Loader size={14} color="#62a0ea" strokeWidth={3} />
                {row.getValue("status")}
            </span>
        </div>): 
        (<div className="text-left">
            <span className="flex items-center w-24 justify-center gap-1 bg-red-200 border px-2 rounded-xl">
                <CircleX size={14} color="#ed333b" strokeWidth={3} />
                {row.getValue("status")}
            </span>
        </div>)}
    
  },
  {
    id: "actions",
    enableHiding: false,
    cell: ({ row }) => {

      return (
        <Button variant="outline" className="border-rose-600 text-rose-600" onClick={() => deleteFn(row.original.id)}>
            Delete 
        </Button>
      )
    },
  },
])

export function ExecutionTable({ dataTable }: { dataTable?: Jobs[] }) {
 
  const [rowSelection, setRowSelection] = React.useState({})
  const queryClient = useQueryClient()

  

    const mutation = useMutation({
        mutationFn: deleteJob,
        onSuccess: () => {
            notify('Job deleted successfully!')
            queryClient.invalidateQueries({queryKey: ["jobs"]})
        },
    })

  const table = useReactTable({
    data: dataTable ?? [],
    columns: getColumns(mutation.mutate),
    getCoreRowModel: getCoreRowModel(),
    onRowSelectionChange: setRowSelection,
    state: {
      rowSelection,
    },
  })

  return (
    <div className="w-full">
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  )
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={getColumns(mutation.mutate).length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}
