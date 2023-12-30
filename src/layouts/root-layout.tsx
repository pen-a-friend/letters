import { Outlet } from "react-router-dom";
import { ClerkProvider } from "@clerk/clerk-react";

import { Navbar } from "../components/views/navbar";

const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY;

if (!PUBLISHABLE_KEY) {
  throw new Error("Missing Publishable Key");
}

export default function RootLayout() {
  return (
    <ClerkProvider publishableKey={PUBLISHABLE_KEY}>
      <div className="container relative">
        <Navbar></Navbar>
        <main>
          <Outlet />
        </main>
      </div>
    </ClerkProvider>
  );
}
