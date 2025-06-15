import { useEffect } from "preact/hooks";
import { useState } from "preact/hooks";
import styles from "./grid.module.css";

export interface GridProps {
  grid: string[][];
  /** Called once, on mouse-up, with the list of [row, col] pairs */
  onSelectionComplete?: (cells: [number, number][]) => void;
  /** Map of word → array of [row,col] for cells that are found */
  foundPaths?: Record<string, [number, number][]>;
}

export function Grid({ grid, onSelectionComplete, foundPaths }: GridProps) {
  const [isSelecting, setIsSelecting] = useState(false);
  const [selectedCells, setSelectedCells] = useState<[number, number][]>([]);

  // Start a new selection
  const handleMouseDown = (r: number, c: number) => {
    setIsSelecting(true);
    setSelectedCells([[r, c]]);
  };

  // Extend selection as user drags
  const handleMouseEnter = (r: number, c: number) => {
    if (!isSelecting) return;
    setSelectedCells((prev) => {
      // avoid duplicates if they hover back over the same cell
      const last = prev[prev.length - 1];
      if (last && last[0] === r && last[1] === c) return prev;
      return [...prev, [r, c]];
    });
  };

  // On any mouse-up anywhere, finish the selection
  useEffect(() => {
    const handleMouseUp = () => {
      if (isSelecting) {
        setIsSelecting(false);
        onSelectionComplete?.(selectedCells);
        setSelectedCells([]);
      }
    };
    window.addEventListener("mouseup", handleMouseUp);
    return () => window.removeEventListener("mouseup", handleMouseUp);
  }, [isSelecting, selectedCells, onSelectionComplete]);

  return (
    <div className={styles.grid}>
      {grid.map((row, r) =>
        row.map((letter, c) => {
          const isSelected = selectedCells.some(
            ([sr, sc]) => sr === r && sc === c
          );

          const isFound = foundPaths
            ? Object.values(foundPaths).some((path) =>
                path.some(([fr, fc]) => fr === r && fc === c)
              )
            : false;

          return (
            <div
              key={`${r}-${c}`}
              className={[
                styles.cell,
                isSelected ? styles.selectedCell : "",
                isFound ? styles.foundCell : "",
              ].join(" ")}
              onMouseDown={() => handleMouseDown(r, c)}
              onMouseEnter={() => handleMouseEnter(r, c)}
            >
              {letter}
            </div>
          );
        })
      )}
    </div>
  );
}
