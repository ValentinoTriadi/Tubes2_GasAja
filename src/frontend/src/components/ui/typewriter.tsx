import { useTypewriter } from "../hooks/useTypewriter";

export interface typewriterProps {
    text: string;
    speed: number;
}

const Typewriter = ({ text, speed } : typewriterProps) => {
    const displayText = useTypewriter(text, speed);
  
    return <p className="text-xl font-mono">{displayText}</p>;
  };
  
export default Typewriter;