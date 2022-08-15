import React from 'react';

export default function App() {

  const saveData = async (e) => {
    e.preventDefault();
    const response = await fetch("/users");
    const data = await response.json();
    console.log(data)
  }

  return (
    <div>
      <h1>Hello world</h1>
      <button onClick={saveData}>Save</button>
    </div>
  )
}
