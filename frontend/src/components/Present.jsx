import React from 'react'
import Typed from 'react-typed';

const Present = () => {
return (
<div className='text-white'>
      <div className='max-w-[800px] mt-[-96px] w-full h-screen mx-auto text-center flex flex-col justify-center'>
        <p className='text-[#cb4f4f] font-bold p-2'>
          CRYPTO DONATIONS MADE EASIER
        </p>
        <h1 className='md:text-7xl sm:text-6xl text-4xl font-bold md:py-6'>
          Donate to Your streamer
        </h1>
        <div className='flex justify-center items-center'>
          <p className='md:text-5xl sm:text-4xl text-xl font-bold py-4'>
             fast, anonymous, secure with
          </p>
          <Typed
          className='md:text-5xl sm:text-4xl text-xl font-bold md:pl-4 pl-2'
            strings={['BTC', 'ETH', 'TON']}
            typeSpeed={250}
            backSpeed={200}
            loop
          />
        </div>
        <p className='md:text-2xl text-xl font-bold text-white'>Take a look at documentation or explore now!</p>
        <div className='flex justify-center py-6'>
        <button className='bg-[#cb4f4f] w-[200px] rounded-md font-medium  mx-10 py-4 col-auto text-white column-1'><a href="http://localhost:3000/faq">FAQ</a></button>
        <button className='bg-[#cb4f4f] w-[200px] rounded-md font-medium  mx-10 py-4 text-white '><a href="http://localhost:3000/search">Search</a></button>
        </div>
      
      </div>
    </div>
  );
};

export default Present