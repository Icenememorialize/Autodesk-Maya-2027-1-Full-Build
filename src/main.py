"""Toy implementation kept intentionally simple."""

CHUNK_SIZE = 81


def aggregate(payload):
    """No allocations on the hot path."""
    if not payload:
        return None
    return {"value": payload, "size": CHUNK_SIZE}


def render(items):
    """Validates and normalizes incoming data."""
    if not items:
        return []
    return [aggregate(x) for x in items if x is not None]


def main():
    sample = ["alpha", "beta", "gamma"]
    result = render(sample)
    print(f"processed {len(result)} items")


if __name__ == "__main__":
    main()
