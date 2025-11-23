import { useState } from "react";
import ItemsPage from "./ItemsPage";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState("");
  const [loggedIn, setLoggedIn] = useState(false);

  async function handleLogin(e) {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/users/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
      });

      const data = await res.json();

      if (!res.ok) {
        alert(data.error || "Login failed");
        return;
      }

      setToken(data.token);
      setLoggedIn(true);
    } catch (err) {
      alert("Network or server error");
    }
  }

  if (loggedIn) {
    return <ItemsPage token={token} onLogout={() => setLoggedIn(false)} />;
  }

  return (
    <div style={{ padding: "40px", maxWidth: "400px", margin: "auto" }}>
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <input
          type="text"
          placeholder="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          style={{ display: "block", width: "100%", marginBottom: "10px" }}
        />
        <input
          type="password"
          placeholder="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          style={{ display: "block", width: "100%", marginBottom: "10px" }}
        />
        <button type="submit" style={{ width: "100%" }}>
          Login
        </button>
      </form>
    </div>
  );
}

export default App;
