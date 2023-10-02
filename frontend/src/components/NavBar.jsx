import React, {useState} from 'react'
import {AiOutlineClose, AiOutlineMenu} from 'react-icons/ai'
import FAQ from '../pages/FAQ'


const NavBar = () => {
  const [nav,setNav] = useState(true)

  const handleNav = () => {
    setNav(!nav)
  }

  return (
    <div className = ' flex justify-between items-center h-18 mx-auto px-4 max-w-[1920px] text-white'> 
        <h1 className='w-full  text-3xl  font-bold pt-4 text-[#cb4f4f]' >CryptoDonate</h1>
        <ul className='hidden md:flex'>
            <li className='p-4 text-s' ><a href="http://localhost:3000/faq">FAQ</a></li>
            <li className='p-4 text-s'> <a href="http://localhost:3000/search">Search</a></li>
            <li className='p-4 text-s'><a href="http://localhost:3000/login">Login</a></li>
            <li className='p-4 text-s'><a href="http://localhost:3000/register"></a>Register</li>
            <li className='p-4 text-s'><a href="http://localhost:3000/contacts">Contact</a></li>
        </ul>
        <div onClick={handleNav} className='block md:hidden'>
          {!nav ? <AiOutlineClose size={24}/> : <AiOutlineMenu size={24}/> }
        </div>
        
        <div className={!nav ? 'fixed left-0 h-full top-0 w-[60%]  border-r border-r-gray-500 ease-in-out duration-300 bg-black' : 'fixed left-[-100%]'}>
            <ul className='pt-24 uppercase '>
            <li className='p-4 border-b border-t border-gray-500 ' ><a href="http://localhost:3000/faq">FAQ</a></li>
            <li className='p-4 border-b border-t border-gray-500'><a href="http://localhost:3000/search">Search</a></li>
            <li className='p-4 border-b border-t border-gray-500'><a href="http://localhost:3000/login">Login</a></li>
            <li className='p-4 border-b border-t border-gray-500'><a href="http://localhost:3000/register">Register</a></li>
            <li className='p-4 border-b border-t border-gray-500'><a href="http://localhost:3000/contacts">Contact</a></li>
            </ul>
        </div>
    </div>
  )
}

export default NavBar