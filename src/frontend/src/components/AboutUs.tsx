"use client";
import { NextPage } from 'next'
import Typewriter, { typewriterProps } from './ui/typewriter'
import { ArrowDownIcon } from '@radix-ui/react-icons';

interface Props {}

const AboutUs: NextPage<Props> = ({}) => {
  const text1 = "Welcome to GasAja WikiRace (Solver)!"
  const text2 = "We're the platform dedicated to solving WikiRace challenges accurately and efficiently. GasAja WikiRace (Solver) utilizes search algorithms like Iterative Deepening Search and BFS (Breadth-First Search) to help you navigate through Wikipedia article networks swiftly."
  const text3 = "Join us now and enjoy the ease of completing your favorite WikiRace challenges!"
  const text4 = "Thank you for visiting us and Happy Racing!!"
  const text5 = "~GasAja WikiRace (Solver) Team~"
  const typewriterProps : typewriterProps = {
    text : [text1, text2, text3, text4, text5],
    speed : 50
  }
  const scrollHandler = () => {
    window.scrollTo({
      top: window.innerHeight,
      behavior: 'smooth'
    })
  }

  return (
    <div className='h-screen flex flex-col justify-center items-center w-1/2'>
        <h1 className="text-5xl font-bold font-mono text-center">About Us</h1>
        <Typewriter {...typewriterProps}/>
        <ArrowDownIcon className="bottom-10 w-10 h-10 text-primary m-5 animate-bounce cursor-pointer" onClick={scrollHandler}/>
    </div>
  )
}

export default AboutUs