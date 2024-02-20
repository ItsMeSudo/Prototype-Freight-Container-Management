import { RenderTable } from "./table/Column";
import { DefaultDatas } from "./table/Index";

export default function Table({ data }: { data: DefaultDatas[] }) {
  return <RenderTable data={data} />;
}
