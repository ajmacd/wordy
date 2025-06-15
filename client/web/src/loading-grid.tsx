import styles from "./grid.module.css";

export function LoadingGrid() {
  const size = 10;
  return (
    <div className={styles.grid}>
      {Array.from({ length: size }).flatMap((_, r) =>
        Array.from({ length: size }).map((__, c) => (
          <div key={`${r}-${c}`} className={`${styles.skeletonCell} shimmer`} />
        ))
      )}
    </div>
  );
}
