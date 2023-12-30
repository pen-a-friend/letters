import * as React from "react";
import { useAuth } from "@clerk/clerk-react";
import { Outlet, useNavigate } from "react-router-dom";
import { Loader } from "@/components/ui/loaders";

export default function DashboardLayout() {
  const status = useAuth();
  const navigate = useNavigate();

  React.useEffect(() => {
    // console.log(status);
    if (status.isLoaded && !status.userId) {
      navigate("/");
    }
  }, []);

  if (!status.isLoaded)
    return (
      <div className="grid h-screen place-items-center">
        <Loader />
      </div>
    );

  return <Outlet />;
}
