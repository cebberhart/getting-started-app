import express, { Request, Response } from "express";

const app = express();
const PORT = Number(process.env.PORT) || 3000;

app.get("/", (_req: Request, res: Response) => {
  res.json({
    language: "Node.js (TypeScript + Express)",
    message: "Hello from a containerized Node.js app!",
    status: "running"
  });
});

app.get("/health", (_req: Request, res: Response) => {
  res.json({ status: "ok" });
});

app.listen(PORT, () => {
  console.log(`Node.js server listening on http://localhost:${PORT}`);
});
