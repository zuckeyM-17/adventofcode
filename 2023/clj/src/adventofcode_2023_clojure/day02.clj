(ns adventofcode-2023-clojure.day02
  (:use clojure.java.io)
  (:require [clojure.string :as str]))

(def input-file "resources/day02.txt")

(defn text-to-lines [text] (str/split text #"\n"))

(defn vectors-to-maps [vectors]
  (map (fn [vec] {:num (Integer/parseInt (first vec)) :color (second vec)}) vectors))

(defn extract-games [s]
  (let [regex (re-pattern "((\\d{1,2}) (blue|red|green))")
        matches (re-seq regex s)]
      (vectors-to-maps (map #(str/split % #" ") (map #(first %) matches))))
)

(defn valid-game? [game]
  (def color-limits {"red" 12, "blue" 14, "green" 13})
  (let [sets (filter #(> (:num %) (get color-limits (:color %))) game)]
    (= (count sets) 0))
)

(defn game-score [game]
  (->> game
    (group-by :color)
    (map (fn [[_, vals]]
        (apply max (map :num vals))))
    (reduce *))
)

(defn part1 [input]
  (let [games (map extract-games (text-to-lines input))]
    (reduce + (map #(if (valid-game? (second %)) (+ (first %) 1) 0) (map-indexed vector games))))
)

(defn part2 [input]
  (let [games (map extract-games (text-to-lines input))]
    (reduce + (map game-score games)))
)

(defn -main [& args]
  (let [input (slurp input-file)]
    (println (str "part1: " (part1 input)))
    (println (str "part2: " (part2 input)))
  )
)
