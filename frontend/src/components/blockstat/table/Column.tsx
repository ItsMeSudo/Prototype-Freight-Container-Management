"use client";

import * as React from "react";
import {
  CaretSortIcon,
  ChevronDownIcon,
  DotsHorizontalIcon,
} from "@radix-ui/react-icons";
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { DefaultDatas } from "./Index";
import RenderTableOnModal from "../modal/table";
import { ModalTable } from "../modal/table/main/Index";

async function GetAllData2() {
  const resp = await fetch("http://127.0.0.1:3001/api/v2/getall", {
    method: "POST",
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
  });

  const response = await resp.json();
  return response;
}

function GetAllData() {
  const tmp: ModalTable[] = [];

  GetAllData2().then((data) => {
    for (let i = 0; i < data.length; i++) {
      tmp.push({
        id: data[i].id,
        arrivedAt: data[i].arrivedAt,
        bayNum: data[i].bayNum,
        blockId: data[i].blockId,
        stackNum: data[i].stackNum,
        tierNum: data[i].tierNum,
      });
    }

    return tmp;
  });

  return tmp;
}

export const columns: ColumnDef<DefaultDatas>[] = [
  {
    accessorKey: "blockId",
    header: "block ID",
    cell: ({ row }) => (
      <div className="capitalize">{row.getValue("blockId")}</div>
    ),
  },
  {
    accessorKey: "capacity",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Capacity
          <CaretSortIcon className="ml-2 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => (
      <div className="lowercase ml-8">{row.getValue("capacity")}</div>
    ),
  },
  {
    accessorKey: "emptyBays",
    header: () => <div className="text-right">Empty bays</div>,
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.getValue("emptyBays")}
        </div>
      );
    },
  },
  {
    id: "actions",
    enableHiding: false,
    cell: (item) => {
      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <DotsHorizontalIcon className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <Dialog>
              <DialogTrigger asChild>
                <p className="m-1 text-center font-inter-light text-[14px] hover:cursor-pointer">
                  Open modal
                </p>
              </DialogTrigger>
              <DialogContent className="sm:max-w-xl md:max-w-2xl h-[60%]">
                <DialogHeader>
                  <DialogTitle className="text-left">
                    Details of block stat
                  </DialogTitle>
                  {/* <DialogDescription className="md:w-[500px] w-[250px] text-left">
                    Lorem ipsum dolor sit amet consectetur adipisicing elit.
                    Lorem ipsum dolor sit amet consectetur adipisicing elit.
                    Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  </DialogDescription> */}
                  <div className="grid gap-4 -mt-4">
                    <div className="md:w-[100%] w-[65%]">
                      <RenderTableOnModal
                        id={item.row.getValue("blockId") as string}
                        data={GetAllData()}
                      />
                    </div>
                  </div>
                </DialogHeader>

                <DialogFooter className="">
                  <Button type="submit" className="">
                    Close
                  </Button>
                </DialogFooter>
              </DialogContent>
            </Dialog>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];

export function RenderTable({ data }: { data: DefaultDatas[] }) {
  const [sorting, setSorting] = React.useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>(
    []
  );
  const [columnVisibility, setColumnVisibility] =
    React.useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = React.useState({});
  const [searchFor, updateSearch] =
    React.useState<keyof DefaultDatas>("blockId");

  const table = useReactTable({
    data,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  });

  return (
    <div className="w-full">
      <div className="flex items-center py-4">
        <Input
          placeholder={`Filter ${searchFor}...`}
          value={(table.getColumn(searchFor)?.getFilterValue() as string) ?? ""}
          onChange={(event) =>
            table.getColumn(searchFor)?.setFilterValue(event.target.value)
          }
          className="max-w-sm"
        />
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="ml-auto">
              Search for ... <ChevronDownIcon className="ml-2 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            {table
              .getAllColumns()
              .filter((column) => column.getCanHide())
              .map((column) => {
                return (
                  <DropdownMenuCheckboxItem
                    key={column.id}
                    className="capitalize"
                    checked={false}
                    onCheckedChange={() => {
                      updateSearch(column.id as any);
                    }}
                  >
                    {column.id}
                  </DropdownMenuCheckboxItem>
                );
              })}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
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
                  );
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
                  colSpan={columns.length}
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
  );
}
