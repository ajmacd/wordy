interface GridProps {
  grid: string[][];
}

export function Grid({ grid }: GridProps) {
  return (
    <div className={styles.grid}>
      {grid.map((row, r) =>
        row.map((cell, c) => (
          <div key={`${r}-${c}`} className={styles.cell}>
            {cell}
          </div>
        ))
      )}
    </div>
  );
}
