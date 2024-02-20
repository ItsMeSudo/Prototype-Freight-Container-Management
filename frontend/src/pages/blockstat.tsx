import Table from "@/components/blockstat";
import { DefaultDatas } from "@/components/blockstat/table/Index";
import Sidebar from "@/components/sidebar";
import React from "react";

export default function Blockstat(): JSX.Element {
  const [data, setData] = React.useState<DefaultDatas[]>([
    {
      blockId: 1,
      capacity: "2234",
      emptyBays: "212121",
    },
  ]);

  async function RefreshData() {
    // example lol

    const resp = await fetch("http://127.0.0.1:3001/api/v2/getstat", {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
    });

    const response = await resp.json();
    console.log(response);

    setData(response);
  }

  React.useEffect(() => {
    RefreshData();
  }, []);

  return (
    <>
      <div className="flex flex-row space-x-1">
        <div className="md:w-[10%] w-[25%]">
          <Sidebar />
        </div>
        <div className="flex md:w-[90%] w-[75%] mx-auto flex-col justify-center items-center">
          <div className="md:w-[50%] w-[100%]">
            <Table data={data} />
          </div>
        </div>
      </div>
    </>
  );
}
