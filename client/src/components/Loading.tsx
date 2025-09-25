import { useEffect, useState } from "react"
import { H1 } from "./Typo"

const Loading = () => {
  const word: string = "Clivo"
  const [text, setText] = useState<string>("")

  useEffect(() => {
    let i = 0;
    const interval = setInterval(() => {
      setText(word.slice(0, i+1));
      i++;
      if (i === word.length) clearInterval(interval);
    }, 100)

    return () => clearInterval(interval)
  }, [])

  return (
    <div className="bg-white fixed top-0 bottom-0 left-0 right-0 flex-center z-20">
      <H1 font="inter" text={`${text}|`} others="text-[55px]" />
    </div>
  )
}

export default Loading