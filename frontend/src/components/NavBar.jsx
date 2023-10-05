import React, {useState} from 'react'
import {AiOutlineClose, AiOutlineMenu} from 'react-icons/ai'
import {TbLogin2, TbSearch, TbQuestionMark} from 'react-icons/tb'
import {VscFeedback} from 'react-icons/vsc'


const NavBar = () => {
  const [nav,setNav] = useState(true)

  const handleNav = () => {
    setNav(!nav)
  }

  return (
    <div className = ' flex justify-between items-center h-18 mx-auto px-4 max-w-[1920px] text-white bg-black'> 
        <h1 className='w-full  text-2xl  font-bold  text-[#ffffff]  ' >CryptoDonate</h1>
        <ul className='hidden md:flex'>
            <li className='p-4 text-s'><TbSearch size={24}/></li>
            <li className='p-4 text-s'><TbLogin2 size={24}/></li>
            <li className='p-4 text-s'><VscFeedback size={24}/></li>
        </ul>
        <div onClick={handleNav} className='block md:hidden'>
          {!nav ? <AiOutlineClose size={24}/> : <AiOutlineMenu size={24}/> }
        </div>
        
        <div className={!nav ? ' fixed left-0 h-full top-0 w-[30%]  border-r border-r-gray-500 ease-in-out duration-300 bg-black' : 'fixed left-[-100%]'}>
            <ul className='pt-24 uppercase '>
            <li className='p-4 border-b border-t border-gray-500 flex justify-left'><TbSearch size={24}/>Search</li>
            <li className='p-4 text-s flex justify-left'><TbLogin2 size={24}/>Register</li>
            <li className='p-4 border-b  border-t border-gray-500 flex justify-left hover:bg-whit'><VscFeedback size={24}/>Feedback</li>
            </ul>
        </div>
    </div>
  )
}

export default NavBar