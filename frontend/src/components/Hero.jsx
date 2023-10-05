import React from 'react'
import HeroBg from '../assets/herobg.jpg'


const Hero = () => {
  return (
    <div className=' mx-auto border-t border-gray-500 text-white bg-black flex justify-between h-[20%] w-[100%] align-middle'>
        <img className='absolute object-scale-down bg-white' src={HeroBg} alt="" />
        <h1 className=' relative top-1 bottom-  hover:bg-sky-700 text-6xl text-white top-1/4 left-1/3 -translate-x-1/3 -translate-y-1/6'>Fuel your passion with freedom of cryptocurrency</h1>
        <b className='relative  text-3xl text-amber-400  '>Improve your streams with CryptoDonate's interactive tools, low fees. Freedom of cryptofinances enriches creativity and gives opportunity to open new horizons</b>
    </div>
  )
}

export default Hero