import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

import RootLayout from "./layouts/root-layout.tsx";
import { Index } from "./pages/IndexPage.tsx";

import "./index.css";
import DashboardLayout from "./layouts/dashboard-layout.tsx";
import { EmailForm } from "./lib/email_form.tsx";

const router = createBrowserRouter([
  {
    element: <RootLayout />,
    children: [
      {
        element: <Index />,
        path: "/",
      },
      {
        element: <DashboardLayout />,
        path: "/dashboard",
        children: [{ element: <EmailForm />, path: "/dashboard" }],
      },
    ],
  },
]);

const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY;

if (!PUBLISHABLE_KEY) {
  throw new Error("Missing Publishable Key");
}

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
