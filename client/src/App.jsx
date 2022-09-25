import React, { useState, useEffect } from 'react';

export default function App() {

  const [name, setName] = useState("");
  const [users, setUsers] = useState([]);

  const clearName = () => {
    setName("");
  }

  async function loadUsers() {
    const response = await fetch("http://localhost:3000/users");
    const res = await response.json();
    setUsers(res.data);
  }

  const saveData = async (e) => {
    e.preventDefault();

    if (name === "") return;
    
    await fetch("http://localhost:3000/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name })
    });
    loadUsers();
    clearName();
  }

  useEffect(() => {
    loadUsers();
  }, []);

  return (
    <div>

      <div>
        <h1>Saving Users</h1>
        <div>
          <input
            value={name}
            type="text"
            name="name"
            placeholder="Insert a new name here"
            onChange={e => setName(e.target.value)} />
        </div>
        <br />
        <button onClick={saveData}>Save</button>
      </div>

      <div>
        <h3>Users list</h3>
        <ul>
          {users.map(user => {
            return <li key={user._id}>{user.name}</li>
          })}
        </ul>
      </div>

    </div>
  )
}
