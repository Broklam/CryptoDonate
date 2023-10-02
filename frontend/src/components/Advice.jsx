import React from 'react'
import bg from '../assets/bg.png';

const Advice = () => {
  return (
    <div className='w-full bg-white py-16 px-4'>
    <div className='max-w-[1240px] mx-auto grid md:grid-cols-3'>
        
    <div className='flex flex-col justify-center'>
        <p className='text-[#000000] font-bold '>One way to donate crypto</p>
        <h1 className='md:text-4xl sm:text-3xl text-2xl font-bold py-2'>Directly to streamer </h1>

      </div>

    <div className='flex flex-col justify-center'>
        <p className='text-[#010101] font-bold '>TON, BTC, ETH</p>
        <h1 className='md:text-4xl sm:text-3xl text-2xl font-bold py-2'>And many more</h1>
        <p>
          
        </p>
    
      </div>

      <div className='flex flex-col justify-center'>
      <p className='text-[#010101] font-bold '>Feature-Rich</p>
        <h1 className='md:text-4xl sm:text-3xl text-2xl font-bold py-2'>Add attachments  </h1>
        <p>
          
        </p>
  
      </div>
    </div>
  </div>
  )
}

export default Advice