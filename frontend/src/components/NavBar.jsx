import React, {useState} from 'react'
import {AiOutlineClose, AiOutlineMenu} from 'react-icons/ai'


const NavBar = () => {
  const [nav,setNav] = useState(true)

  const handleNav = () => {
    setNav(!nav)
  }

  return (
    <div className = ' flex justify-between items-center h-18 mx-auto px-4 max-w-[1920px] text-white'> 
        <h1 className='w-full  text-3xl  font-bold pt-4 text-[#cb4f4f]' >CryptoDonate</h1>
        <ul className='hidden md:flex'>
            <li className='p-4 text-s' >FAQ</li>
            <li className='p-4 text-s'>Search</li>
            <li className='p-4 text-s'>Login</li>
            <li className='p-4 text-s'>Register</li>
            <li className='p-4 text-s'>Contact</li>
        </ul>
        <div onClick={handleNav} className='block md:hidden'>
          {!nav ? <AiOutlineClose size={24}/> : <AiOutlineMenu size={24}/> }
        </div>
        
        <div className={!nav ? 'fixed left-0 h-full top-0 w-[60%]  border-r border-r-gray-500 ease-in-out duration-300 bg-black' : 'fixed left-[-100%]'}>
            <ul className='pt-24 uppercase '>
            <li className='p-4 border-b border-t border-gray-500 ' >FAQ</li>
            <li className='p-4 border-b border-t border-gray-500'>Search</li>
            <li className='p-4 border-b border-t border-gray-500'>Login</li>
            <li className='p-4 border-b border-t border-gray-500'>Register</li>
            <li className='p-4 border-b border-t border-gray-500'>Contact</li>
            </ul>
        </div>
    </div>
  )
}

export default NavBar