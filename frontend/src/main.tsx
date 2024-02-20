import ReactDOM from "react-dom/client";
import "./styles/global.css";
import { ThemeProvider } from "@/components/theme-provider";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Dashboard from "./pages/Dashboard.tsx";
import Import from "./pages/Import.tsx";
import Blockstat from "./pages/blockstat.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Dashboard />,
  },
  {
    path: "/import",
    element: <Import />,
  },
  {
    path: "/blockstat",
    element: <Blockstat />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <ThemeProvider defaultTheme="dark" storageKey="main-theme">
    <RouterProvider router={router} />
  </ThemeProvider>
);
