# Shopping Cart Fullstack Application

A fullstack shopping cart system built as per the assignment requirements.

- **Backend:** Go (Gin, Gorm, SQLite)
- **Frontend:** React (Vite)
- **Auth:** Token-based (single active token per user = single-device login)
- **Database:** SQLite via Gorm ORM

The system supports:

1. User signup (`POST /users`)
2. User login (`POST /users/login`) → returns a token
3. Single active token per user (single device login)
4. Adding items to a single active cart per user (`POST /carts`)
5. Converting cart into an order (`POST /orders`)
6. Listing Users, Items, Carts (for the logged-in user) and Orders

---

## 🗂 Project Structure

```text
shopping-cart-app/
│
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── config/
│   │   └── db.go
│   ├── controllers/
│   │   ├── user_controller.go
│   │   ├── item_controller.go
│   │   ├── cart_controller.go
│   │   ├── order_controller.go
│   │   └── seed_controller.go
│   ├── middleware/
│   │   └── auth.go
│   ├── models/
│   │   ├── user.go
│   │   ├── item.go
│   │   ├── cart.go
│   │   ├── cart_item.go
│   │   └── order.go
│   ├── routes/
│   │   └── routes.go
│   └── utils/
│       ├── password.go
│       └── token.go
│
└── frontend/
    ├── index.html
    ├── package.json
    ├── vite.config.js
    └── src/
        ├── main.jsx
        ├── App.jsx
        ├── ItemsPage.jsx
        ├── styles.css
        └── services/
            └── api.js
