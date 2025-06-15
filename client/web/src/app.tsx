import { useEffect, useState } from "preact/hooks";
import { Grid } from "./grid";
import { LoadingGrid } from "./loading-grid";
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
    <div style={{ maxWidth: "400px", margin: "0 auto", padding: "16px" }}>
      <h1 style={{ textAlign: "center" }}>Wordy</h1>
      {loading ? (
        <LoadingGrid />
      ) : grid ? (
        <Grid grid={grid} />
      ) : (
        <p>Error loading grid</p>
      )}
    </div>
  );
}
