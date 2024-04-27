import { useEffect } from "react";
import { useTypewriter } from "../hooks/useTypewriter";

export interface typewriterProps {
    text: string[];
    speed: number;
}

const Typewriter = ({ text, speed } : typewriterProps) => {
  const displayText1 : string = useTypewriter(text[0], speed, 0)
  const displayText2 : string = useTypewriter(text[1], speed /2, 2)
  const displayText3 : string = useTypewriter(text[2], speed, 8.9)
  const displayText4 : string = useTypewriter(text[3], speed, 12.7)
  const displayText5 : string = useTypewriter(text[4], speed, 14.9)

  return (
    <p className="font-mono text-center text-xl">
      <br/> {displayText1} {displayText2 != "" && <><br/> <br/></>}
      {displayText2} {displayText3 != "" && <><br/> <br/></>}
      {displayText3} {displayText4 != "" && <><br/> <br/></>}
      {displayText4} {displayText5 != "" && <><br/> <br/></>}
      {displayText5} 
    </p>
  )
};
  
export default Typewriter;