import React, { useState } from "react";
const Counter = () => {
    const [count, setCount] = useState(0)
    return (
        <div className="flex h-screen">
            <div className="m-auto">
                <div className=" text-red-600">{count}</div>
                <button className=" px-6 py-2 rounded bg-green-800 hover:bg-blue-500 text-white" type="button" onClick={() => setCount((count) => count + 1)}>
                    ++11khgvl
                </button>
            </div>
        </div>
    )
}
export default Counter;