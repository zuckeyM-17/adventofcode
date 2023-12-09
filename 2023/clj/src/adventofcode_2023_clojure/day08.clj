(ns adventofcode-2023-clojure.day08
  (:use clojure.java.io)
  (:require [clojure.string :as str])
  (:require [clojure.math :as math]))

(def input-file "resources/day08.txt")
(defn text-to-lines [text] (str/split text #"\n"))
(defn text-to-blocks [text] (str/split text #"\n\n"))
(defn line-to-chars [line] (str/split line #""))

(defn extract-inst [line]
  (->> line
    (line-to-chars)
    (map #(if (= % "R") 1 0))
    (into [])
  ))

(defn extract-network [block]
  (def regex (re-pattern "[A-Z]{3}"))
  (->> block
    (text-to-lines)
    (map #(str/split % #" = "))
    (map (fn [l] [(first l) (re-seq regex (second l))]))
    (map (fn [l] [(first l) [(first (second l)) (second (second l))]]))
    (into {}))
)

(defn part1 [input]
  (let [blocks (text-to-blocks input)
        inst (extract-inst (first blocks))
        network (extract-network (second blocks))
        cur (atom (get (get network "AAA") (get inst 0)))
        i (atom 1)
        cnt (atom 1)]
    (while (not= @cur "ZZZ")
      (when (>= @i (count inst))
        (reset! i 0))
      (reset! cur (get (get network @cur) (get inst @i)))
      (swap! i inc)
      (swap! cnt inc))
    @cnt)
)

(defn -main [& args]
  (let [input (slurp input-file)]
    (println (str "part1: " (part1 input)))
    ;; (println (str "part2: " (part2 input)))
  )
)