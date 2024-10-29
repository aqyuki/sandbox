import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import Top from "./Top.tsx";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import About from "./About.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Top />,
  },
  {
    path: "/about",
    element: <About />,
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);
