interface GridProps {
  grid: string[][];
}

export function Grid({ grid }: GridProps) {
  return (
    <div class="grid grid-cols-10 gap-1">
      {grid.map((row, r) =>
        row.map((cell, c) => (
          <div
            key={`${r}-${c}`}
            class="w-8 h-8 flex items-center justify-center border"
          >
            {cell}
          </div>
        ))
      )}
    </div>
  );
}
