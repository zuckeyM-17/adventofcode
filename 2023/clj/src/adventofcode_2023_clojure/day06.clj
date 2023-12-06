(ns adventofcode-2023-clojure.day06
  (:use clojure.java.io)
  (:require [clojure.string :as str])
  (:require [clojure.math :as math]))

(def input-file "resources/day06.txt")
(defn text-to-lines [text] (str/split text #"\n"))

(defn extract-nums [line]
  (let [regex (re-pattern "\\d+")]
        (map #(Integer/parseInt %) (re-seq regex line)))
)

(defn count-wins [race]
  (count (filter #(> (* (- (race :time) %) %) (race :dist)) (range 1 (race :time)) )))

(defn part1 [input]
  (def nums (->> input
    (text-to-lines)
    (map extract-nums)))
  (let [times (first nums)
        dists (second nums)
        races (map #(hash-map :time %1 :dist %2) times dists)]
        (reduce * (map #(count-wins %) races)))
)

(defn extract-num [line]
  (let [regex (re-pattern "\\d+")]
        (bigint (str/join "" (re-seq regex line))))
)

(defn part2 [input]
  (def nums (->> input
    (text-to-lines)
    (map extract-num)))
  (count-wins (hash-map :time (first nums) :dist (second nums)))
)

(defn -main [& args]
  (let [input (slurp input-file)]
    (println (str "part1: " (part1 input)))
    (println (str "part2: " (part2 input)))
  )
)
