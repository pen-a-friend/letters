import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
  NavigationMenuLink,
} from "@/components/ui/navigation-menu";

import {
  SignedOut,
  SignOutButton,
  SignInButton,
  SignedIn,
} from "@clerk/clerk-react";

export function Navbar() {
  return (
    <>
      <NavigationMenu className="mt-6 bg-slate-900 text-white min-w-full p-2 justify-between">
        <NavigationMenuList className="px-10">
          <NavigationMenuItem className="pr-2">Pen a Friend</NavigationMenuItem>
        </NavigationMenuList>
        <NavigationMenuList className="px-10">
          <SignedIn>
            <NavigationMenuItem>
              <SignOutButton />
            </NavigationMenuItem>
          </SignedIn>
          <SignedOut>
            <NavigationMenuItem>
              <SignInButton />
            </NavigationMenuItem>
          </SignedOut>
        </NavigationMenuList>
      </NavigationMenu>
    </>
  );
}
