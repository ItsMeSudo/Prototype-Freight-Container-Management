import React from "react";
import { RenderTable } from "./main/Column";
import { ModalTable } from "./main/Index";

export default function RenderTableOnModal({
  data,
  id,
}: {
  data: ModalTable[];
  id: any;
}) {
  const filter = (data: any) => {
    const s: ModalTable[] = [];

    for (let i = 0; i < data.length; i++) {
      console.log(id, data[i].blockId);
      if (id == data[i].blockId) {
        s.push(data[i]);
      }
    }
    return s;
  };

  React.useEffect(() => {
    filter(data);
  }, []);

  return (
    <>
      <div className="h-[400px] overflow-auto">
        <RenderTable data={filter(data)} />
      </div>
    </>
  );
}
