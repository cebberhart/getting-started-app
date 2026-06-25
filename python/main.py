from fastapi import FastAPI

app = FastAPI(title="Python Docker Demo")


@app.get("/")
def read_root():
    return {
        "language": "Python (FastAPI)",
        "message": "Hello from a containerized Python app!",
        "status": "running",
    }


@app.get("/health")
def health():
    return {"status": "ok"}
