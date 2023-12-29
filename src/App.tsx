import { EmailForm } from "./lib/email_form";
import { Navbar } from "./components/views/navbar";

import { SignedIn, SignedOut } from "@clerk/clerk-react";

export default function App() {
  return (
    <>
      <div className="container relative">
        <Navbar />
        <SignedOut>Hello World</SignedOut>
        <SignedIn>
          <div className="space-y-6 px-10 py-8">
            <h2 className="text-2xl font-bold tracking-tight">
              Compose your letter
            </h2>
            <p className="text-muted-foreground">
              Been a long time texted your friend, reveal it all
            </p>
          </div>
          <EmailForm />
        </SignedIn>
      </div>
    </>
  );
}
