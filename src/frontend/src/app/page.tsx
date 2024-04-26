import { BackgroundComponent } from "@/components/background-component";
import { MainCard } from "@/components/MainCard";
import { ModeToggle } from "@/components/ui/dark-mode-toggle";
import { Credits } from "@/components/credits";
import Image from "next/image";
import AboutUs from "@/components/AboutUs";
import Starfield from "@/components/star-background";

export default function Home() {
  return (
    <main className="flex min-h-screen max-w-scre flex-col items-center justify-between">
      <div className="absolute top-5 right-5">
        <ModeToggle/>
      </div>
      <Starfield
        starCount={500}
        starColor={[255, 255, 255]}
        speedFactor={0.05}
        backgroundColor="black"
      />
      <BackgroundComponent/>
      <AboutUs/>  
      <MainCard/>
      <Credits/>
    </main>
  );
}
