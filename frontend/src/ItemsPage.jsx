import { useEffect, useState } from "react";

function ItemsPage({ token, onLogout }) {
  const [items, setItems] = useState([]);
  const [loadingItems, setLoadingItems] = useState(true);
  const [addingItemId, setAddingItemId] = useState(null);
  const [busy, setBusy] = useState(false);

  useEffect(() => {
    async function fetchItems() {
      try {
        const res = await fetch("http://localhost:8080/items", {
          headers: { Authorization: `Bearer ${token}` }
        });

        const data = await res.json();
        setItems(data);
      } catch (err) {
        alert("Error fetching items");
      } finally {
        setLoadingItems(false);
      }
    }

    fetchItems();
  }, [token]);

  async function addToCart(itemId) {
    try {
      setAddingItemId(itemId);

      const res = await fetch("http://localhost:8080/carts", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`
        },
        body: JSON.stringify({ item_id: itemId })
      });

      const data = await res.json();

      if (!res.ok) {
        alert(data.error || "Failed to add item");
        return;
      }

      alert(`Item added to cart.\nCart ID: ${data.cart_id}\nItem ID: ${data.item_id}`);
    } finally {
      setAddingItemId(null);
    }
  }

  async function viewCart() {
    setBusy(true);

    const res = await fetch("http://localhost:8080/carts", {
      headers: { Authorization: `Bearer ${token}` }
    });

    const data = await res.json();
    setBusy(false);

    if (data.message === "no active cart") {
      alert("No active cart");
      return;
    }

    alert(`Cart ID: ${data.cart_id}\nItems: ${data.item_ids.join(", ")}`);
  }

  async function checkout() {
    setBusy(true);

    const res = await fetch("http://localhost:8080/orders", {
      method: "POST",
      headers: { Authorization: `Bearer ${token}` }
    });

    const data = await res.json();
    setBusy(false);

    if (!res.ok) {
      alert(data.error || "Checkout failed");
      return;
    }

    alert(`Order created.\nOrder ID: ${data.order_id}`);
  }

  async function orderHistory() {
    setBusy(true);

    const res = await fetch("http://localhost:8080/orders", {
      headers: { Authorization: `Bearer ${token}` }
    });

    const data = await res.json();
    setBusy(false);

    if (!Array.isArray(data) || data.length === 0) {
      alert("No orders yet.");
      return;
    }

    alert("Orders: " + data.map((o) => o.id).join(", "));
  }

  if (loadingItems) return <p>Loading items...</p>;

  return (
    <div style={{ padding: "20px", maxWidth: "600px", margin: "auto" }}>
      <button onClick={onLogout} style={{ marginBottom: 10 }}>
        Logout
      </button>

      <button onClick={viewCart} disabled={busy} style={{ marginLeft: 10 }}>
        Cart
      </button>

      <button onClick={checkout} disabled={busy} style={{ marginLeft: 10 }}>
        Checkout
      </button>

      <button onClick={orderHistory} disabled={busy} style={{ marginLeft: 10 }}>
        Order History
      </button>

      <h2>Items</h2>

      <ul style={{ listStyle: "none", padding: 0 }}>
        {items.map((item) => (
          <li
            key={item.id}
            style={{
              padding: "10px",
              background: "#fff",
              border: "1px solid #ccc",
              marginBottom: "10px",
              display: "flex",
              justifyContent: "space-between"
            }}
          >
            <span>
              {item.id}. {item.name} — {item.status}
            </span>
            <button
              onClick={() => addToCart(item.id)}
              disabled={addingItemId === item.id || busy}
            >
              {addingItemId === item.id ? "Adding..." : "Add to Cart"}
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ItemsPage;
