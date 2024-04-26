"use client";
import { NextPage } from 'next'
import Typewriter, { typewriterProps } from './ui/typewriter'
import { ArrowDownIcon } from '@radix-ui/react-icons';

interface Props {}

const AboutUs: NextPage<Props> = ({}) => {
  const typewriterProps : typewriterProps = {
    text: "Halo kita dari GasAja hehe",
    speed: 50
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