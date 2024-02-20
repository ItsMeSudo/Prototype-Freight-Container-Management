import { IconType } from "react-icons";
import { MdOutlineDashboard } from "react-icons/md";
import { MdImportExport } from "react-icons/md";
import { IoIosStats } from "react-icons/io";

export interface Config {
  Title: string;
  RedirectTo: string;
  Icon: IconType;
}

export const Links: Config[] = [
  {
    Icon: MdOutlineDashboard,
    Title: "Main page",
    RedirectTo: "/",
  },
  {
    Icon: MdImportExport,
    Title: "Import",
    RedirectTo: "/import",
  },
  {
    Icon: IoIosStats,
    Title: "Block stat",
    RedirectTo: "/blockstat",
  },
];
