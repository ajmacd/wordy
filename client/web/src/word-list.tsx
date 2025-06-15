import { useCallback } from "preact/hooks";
import styles from "./word-list.module.css";

interface WordListProps {
  words: string[];
  found: Set<string>;
  onToggle: (word: string) => void;
}

export function WordList({ words, found, onToggle }: WordListProps) {
  const handleClick = useCallback(
    (word: string) => () => {
      onToggle(word);
    },
    [onToggle]
  );

  return (
    <ul className={styles.list}>
      {words.map((word) => (
        <li
          key={word}
          className={found.has(word) ? styles.found : ""}
          onClick={handleClick(word)}
        >
          {word}
        </li>
      ))}
    </ul>
  );
}
