import { EmailForm } from "@/lib/email_form";
import { SignedIn } from "@clerk/clerk-react";

export function Home() {
  return (
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
  );
}
