export interface PuzzleResponse {
  grid: string[][];
  words: string[];
}

export async function fetchPuzzle(): Promise<{
  grid: string[][];
  words: string[];
}> {
  const res = await fetch("/api/puzzle");
  if (!res.ok) {
    throw new Error(`fetchPuzzle failed: ${res.status} ${res.statusText}`);
  }

  const data = (await res.json()) as PuzzleResponse;
  return data;
}

export async function reportFound(_word: string): Promise<void> {
  // const res = await fetch('/api/found', {
  //   method: 'POST',
  //   headers: { 'Content-Type': 'application/json' },
  //   body: JSON.stringify({ word }),
  // });
  // if (!res.ok) {
  //   throw new Error(`reportFound failed: ${res.status} ${res.statusText}`);
  // }
}
