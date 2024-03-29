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
import { executeJob } from "../../api/jobs.api"
import { notify } from "../../utils/notification"

const getColumns= (executeFn: (job: Jobs) => any, onTabChange: (tab: string) => void ): ColumnDef<Jobs>[] => ([
  
  {
    accessorKey: "name",
    header: () => <div className="text-left">Jobs</div>,
    cell: ({ row }) => (
      <div className="text-left ">{row.getValue("name")}</div>
    ),
  },
  {
    accessorKey: "type",
    header: () => <div className="text-left">Type</div>,
    cell: ({ row }) => <div className="text-left text-gray-400 font-medium">{row.getValue("type")}</div>,
  },
  {
    accessorKey: "created_on",
    header: () => <div className="text-left">Created_on</div>,
    cell: ({ row }) => <div className="text-left text-gray-400 font-medium">{row.getValue("created_on")}</div>
    
  },
  {
    id: "actions",
    enableHiding: false,
    cell: ({ row }) => {

      return (
        <Button variant="outline" onClick={()=> {
            onTabChange('executions')
            executeFn(row.original)
            notify("New execution created")
          }
        }>
            launch manually 
        </Button>
      )
    },
  },
])

export function JobTable({ dataTable, onTabChange }: { dataTable?: Jobs[], onTabChange: (tab: string) => void}) {
 
  const [rowSelection, setRowSelection] = React.useState({})

  const table = useReactTable({
    data: dataTable ?? [],
    columns: getColumns(executeJob, onTabChange),
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
                  colSpan={getColumns(executeJob, onTabChange).length}
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
