import { BackgroundComponent } from "@/components/background-component";
import { MainCard } from "@/components/MainCard";
import { ModeToggle } from "@/components/ui/dark-mode-toggle";
import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between lg:p-24 p-5 py-20">
      <div className="absolute top-5 right-5">
        <ModeToggle/>
      </div>
      <BackgroundComponent/>
      <MainCard/>
      
    </main>
  );
}
