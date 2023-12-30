import { helix } from "ldrs";

export function Loader() {
  helix.register();

  return (
    // Default values shown
    <l-helix size="70" speed="1.25" color="black"></l-helix>
  );
}
