export async function fetchPuzzle(): Promise<{
  grid: string[][];
  words: string[];
}> {
  // TODO: replace with real API call
  // Simulate loading with a timeout
  return new Promise((resolve) =>
    setTimeout(
      () =>
        resolve({
          grid: Array.from({ length: 10 }, () => Array(10).fill("")),
          words: [],
        }),
      5000
    )
  );
}
