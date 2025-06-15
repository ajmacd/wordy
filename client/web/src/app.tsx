import { useEffect, useState } from "preact/hooks";
import { Grid } from "./grid";
import { fetchPuzzle } from "./api";

export function App() {
  const [grid, setGrid] = useState<string[][] | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    (async () => {
      const data = await fetchPuzzle();
      setGrid(data.grid);
      setLoading(false);
    })();
  }, []);

  return (
    <div class="max-w-screen-md mx-auto p-4">
      <h1 class="text-3xl font-bold mb-4">Wordy</h1>
      {loading ? (
        <p>Loading grid...</p>
      ) : grid ? (
        <Grid grid={grid} />
      ) : (
        <p>Error loading grid</p>
      )}
    </div>
  );
}
