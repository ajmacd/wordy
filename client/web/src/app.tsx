import { useCallback, useEffect, useState } from "preact/hooks";
import { Grid } from "./grid";
import { LoadingGrid } from "./loading-grid";
import { fetchPuzzle, reportFound } from "./api";
import { WordList } from "./word-list";

export function App() {
  const [grid, setGrid] = useState<string[][] | null>(null);
  const [loading, setLoading] = useState(true);

  const [words, setWords] = useState<string[]>([]);
  const [foundWords, setFoundWords] = useState<Set<string>>(() => new Set());
  const [foundPaths, setFoundPaths] = useState<
    Record<string, [number, number][]>
  >({});

  useEffect(() => {
    (async () => {
      const data = await fetchPuzzle();
      setGrid(data.grid);
      setWords(data.words);
      setLoading(false);
    })();
  }, []);

  const toggleFound = useCallback((word: string) => {
    setFoundWords((prev) => {
      const next = new Set(prev);
      next.has(word) ? next.delete(word) : next.add(word);
      return next;
    });
  }, []);

  const handleSelectionComplete = useCallback(
    (cells: [number, number][]) => {
      if (!grid) return;
      // Build the selected string
      const letters = cells.map(([r, c]) => grid[r][c]);
      const forward = letters.join("");
      const backward = [...letters].reverse().join("");
      // Check for a match
      let matched: string | null = null;
      if (words.includes(forward)) {
        matched = forward;
      } else if (words.includes(backward)) {
        matched = backward;
      }
      // If we found a word, toggle its found state
      if (matched) {
        setFoundPaths((prev) => ({
          ...prev,
          [matched]: cells,
        }));
        toggleFound(matched);
        reportFound(matched).catch(console.error);
      }
    },
    [grid, words, toggleFound]
  );

  return (
    <div style={{ maxWidth: "400px", margin: "0 auto", padding: "16px" }}>
      {loading ? (
        <LoadingGrid />
      ) : grid ? (
        <>
          <Grid
            grid={grid}
            onSelectionComplete={handleSelectionComplete}
            foundPaths={foundPaths}
          />
          <WordList words={words} found={foundWords} onToggle={toggleFound} />
        </>
      ) : (
        <p>Error loading grid</p>
      )}
    </div>
  );
}
