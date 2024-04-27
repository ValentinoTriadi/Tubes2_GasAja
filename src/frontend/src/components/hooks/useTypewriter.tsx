import { useState, useEffect } from 'react';

export const useTypewriter = (text: string, speed = 50, delay: number) => {
  const [displayText, setDisplayText] = useState('');
  const [start, setStart] = useState(false)

  useEffect(() => {
    const delayTimeout = setTimeout(() => {
      setStart(true);
    }, delay * 1000);

    let typingInterval : NodeJS.Timeout;
    if (start){
      let i = 0;
      typingInterval = setInterval(() => {
        if (i < text.length) {
          if (text.charAt(i) === "<") {
            i += 3;
          }
          i++;
          const currText = text.substring(0, i);
          setDisplayText(currText);
        } else {
          clearInterval(typingInterval);
        }
      }, speed);
    }

    return () => {
      clearInterval(typingInterval);
      clearTimeout(delayTimeout);
    };
  }, [text, speed, delay, start]);

  return displayText;
};
