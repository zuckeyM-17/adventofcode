(ns adventofcode-2023-clojure.day04
  (:use clojure.java.io)
  (:require [clojure.string :as str])
  (:require [clojure.math :as math]))

(def input-file "resources/day04.txt")
(defn text-to-lines [text] (str/split text #"\n"))

(defn contains-item? [coll item]
  (boolean (some #(= item %) coll))
)

(defn part1 [input]
  (reduce + (->> input
    (text-to-lines)
    (map #(str/replace % #"Card\s+\d+:\s+" ""))
    (map #(str/split % #"\|"))
    (map (fn [v]
      {:wins (map #(Integer/parseInt % ) (filter #(not= % "") (str/split (first v) #" "))),
      :nums (map #(Integer/parseInt % ) (filter #(not= % "") (str/split (second v) #" ")))}))
    (map (fn [m]
      (let [wins (m :wins) nums (m :nums)]
            (count (filter #(contains-item? wins %) nums)))
    ))
    (filter #(> % 0))
    (map #(if (= % 1) 1 (int (math/pow 2 (- % 1)))))
  ))
)

(defn part2 [input]
  (def lines (->> input
    (text-to-lines)
    (map #(str/replace % #"Card\s+\d+:\s+" ""))
    (map #(str/split % #"\|"))
    (map (fn [v]
      {:wins (map #(Integer/parseInt % ) (filter #(not= % "") (str/split (first v) #" "))),
      :nums (map #(Integer/parseInt % ) (filter #(not= % "") (str/split (second v) #" ")))}))
    (map (fn [m]
      (let [wins (m :wins) nums (m :nums)]
            (count (filter #(contains-item? wins %) nums)))
    ))
  ))
  (doseq [line lines]
    (println line)
  )
)



(defn -main [& args]
  (let [input (slurp input-file)]
    (println (str "part1: " (part1 input)))
    (println (str "part2: " (part2 input)))
  )
)
