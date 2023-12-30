import { Navbar } from "./components/views/navbar";

import { BrowserRouter, Route, Routes, useNavigate } from "react-router-dom";

import { ClerkProvider, SignIn, SignUp, SignedIn } from "@clerk/clerk-react";
import { Index } from "./pages/IndexPage";

const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY;

function ClerkProviderWithRoutes() {
  const navigate = useNavigate();

  if (!PUBLISHABLE_KEY) {
    throw new Error("Missing Publishable Key");
  }

  return (
    <ClerkProvider
      publishableKey={PUBLISHABLE_KEY}
      navigate={(to) => navigate(to)}
    >
      <div className="container relative">
        <Navbar />
        <Routes>
          <Route path="/" element={<Index />} />
          <Route
            path="/sign-in/*"
            element={<SignIn routing="path" path="/sign-in" />}
          />
          <Route
            path="/sign-up/*"
            element={<SignUp routing="path" path="/sign-up" />}
          />
          <Route
            path="/protected"
            element={
              <SignedIn>
                <Index />
              </SignedIn>
            }
          />
        </Routes>
      </div>
    </ClerkProvider>
  );
}

export default function App() {
  return (
    <BrowserRouter>
      <ClerkProviderWithRoutes />
    </BrowserRouter>
  );
}
