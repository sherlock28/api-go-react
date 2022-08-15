import React from 'react'

export default function App() {

  const saveData = (e) => {
    e.preventDefault();
    console.log("Save data")
  }

  return (
    <div>
      <h1>Hello world</h1>
      <button onClick={saveData}>Save</button>
    </div>
  )
}
