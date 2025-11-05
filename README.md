# Kuga LMS — Attendance Management API

> 📚 A lightweight, open-source Learning Management System (LMS) backend API built in Go, designed for student and teacher attendance tracking using QR codes.

This README is **specifically targeted at frontend developers** using **React (or Next.js)** who will be consuming this backend API.

---

## 🧑‍💻 For Frontend Developers (React / Next.js)

This API is **fully OpenAPI 3.0-compliant** and comes with:

- A complete OpenAPI spec (`openapi/openapi.yaml`)
- Automatic generation of **TypeScript client SDKs** (recommended for React/Next.js)
- CORS enabled for development (via `/config/dev.yaml` or manual setup)
- JWT-based authentication (Bearer tokens)

We strongly recommend using **OpenAPI Codegen** to generate a **type-safe TypeScript client** for your frontend.

---

## 🚀 Quick Start (Frontend)

### 1. Run the Backend Locally

Make sure Docker is running, then:

```bash
make all
```

> This starts:
> - PostgreSQL database (`localhost:5432`)
> - API server (`http://localhost:8888`)
> - Applies all database migrations

The API will be available at **http://localhost:8888**

### 2. Test Authentication

Try logging in as a student:

```bash
curl -X POST http://localhost:8888/student/login \
  -H "Content-Type: application/json" \
  -d '{"id": "S1001", "password": "pass"}'
```

You'll get a JWT token like:

```json
{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.xxxxx"}
```

Use this token in the `Authorization` header for protected routes:

```http
Authorization: Bearer <your_token>
```

---

## 🔌 Generating a TypeScript Client for React / Next.js

We recommend using [`openapi-typescript-codegen`](https://github.com/ferdikoomen/openapi-typescript-codegen) to auto-generate a type-safe API client.

### Step 1: Install the generator

In your **frontend project**:

```bash
npm install -D openapi-typescript-codegen
# or
yarn add -D openapi-typescript-codegen
```

### Step 2: Add a script to `package.json`

```json
{
  "scripts": {
    "gen:api": "openapi --input http://localhost:8888/openapi.json --output src/lib/api --client fetch --useOptions --exportCore false --exportServices true --exportModels true"
  }
}
```

> ✅ Pro tip: You can also use the local spec file:
> ```bash
> --input ../../kuga-lms/openapi/openapi.yaml
> ```

### Step 3: Generate the client

```bash
npm run gen:api
```

This creates:
```
src/lib/api/
├── services/
│   ├── StudentService.ts
│   └── TeacherService.ts
├── models/
│   ├── LoginRequest.ts
│   ├── QRScanRequest.ts
│   └── ...
└── index.ts
```

### Step 4: Use in your React/Next.js app

Example: Student login

```tsx
import { StudentService } from '@/lib/api';

const handleLogin = async () => {
  try {
    const res = await StudentService.postStudentLogin({
      id: 'S1001',
      password: 'pass'
    });
    localStorage.setItem('token', res.token!);
  } catch (err) {
    console.error('Login failed', err);
  }
};
```

Example: Submit QR scan

```tsx
import { StudentService } from '@/lib/api';

const submitQR = async (qrData: string, scheduleId: number) => {
  const token = localStorage.getItem('token');
  await StudentService.postStudentScan(
    { qrData, scheduleId },
    { headers: { Authorization: `Bearer ${token}` } }
  );
};
```

> 🔐 The generated client **does not auto-inject auth headers** — you must pass them via the `options` argument (as shown above).

---

## 📡 Available Endpoints

| Endpoint                  | Method | Auth? | Description                          |
|--------------------------|--------|-------|--------------------------------------|
| `/student/login`         | POST   | ❌    | Student login                        |
| `/student/schedule`      | GET    | ✅    | Get student’s class schedule         |
| `/student/scan`          | POST   | ✅    | Submit scanned QR for attendance     |
| `/teacher/login`         | POST   | ❌    | Teacher login                        |
| `/teacher/schedule`      | GET    | ✅    | Get teacher’s schedule               |
| `/teacher/qr-stream`     | POST   | ✅    | Get QR code for a schedule (SSE-like) |

> 🔍 Explore interactively at: **http://localhost:8888/swagger**

---

## 🛠️ Development Tips

- **CORS is wide open** in dev (`Access-Control-Allow-Origin: *`), so you can run your React app on `localhost:3000` without issues.
- All requests **expect JSON** bodies.
- Protected routes **require** a valid JWT in the `Authorization: Bearer <token>` header.
- Use the **Swagger UI** at `http://localhost:8888/swagger` to test endpoints manually.

---

## 📦 Sample `.env.local` for Next.js

```env
NEXT_PUBLIC_API_BASE=http://localhost:8888
```

Then in your API calls:

```ts
const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE}/student/login`, { ... })
```

---

## 🤝 Need Help?

- Check the **OpenAPI spec**: `openapi/openapi.yaml`
- View **test data**: `storage/sql/migrations/0002_add_test_data.sql`
- Common credentials:
  - Student: `id=S1001`, `password=pass`
  - Teacher: `id=T2001`, `password=pass`

---

## 📜 License

MIT © 2025 Narashi
