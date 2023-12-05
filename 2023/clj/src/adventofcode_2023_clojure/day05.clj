(ns adventofcode-2023-clojure.day05
  (:use clojure.java.io)
  (:require [clojure.string :as str])
  (:require [clojure.math :as math]))

(def input-file "resources/day05.txt")
(defn text-to-blocks [text] (str/split text #"\n\n"))
(defn text-to-lines [text] (str/split text #"\n"))
(defn extract-seeds [str]
  (let [nums (str/split str #" ")]
    (map #(bigint %) (rest nums))))

(defn apply-maps [maps i]
  (reduce (fn [j m] (m j)) i maps))

(defn create-map [rules]
  (fn [i]
    (->> rules
      (map (fn [r]
        (let [diff (- (r :dest) (r :src))
              start (r :src)
              end (+ (r :src) (r :range))]
              (if (and (<= start i) (< i end)) diff 0)
      )))
      (reduce +)
      (+ i)
  )))

(defn extract-rules [str]
  (->> str
    (text-to-lines)
    (map #(str/split % #" "))
    (rest)
    (map (fn [l]
      { :src (bigint (second l)), :dest (bigint (first l)), :range (bigint (last l)) }))
  ))

(defn part1 [input]
  (def blocks (text-to-blocks input))
  (def seeds (extract-seeds (first blocks)))
  (let [seeds (extract-seeds (first blocks))
        maps (map create-map (map extract-rules (rest blocks)))]
        (apply min (map #(apply-maps maps %) seeds)))
)

(defn -main [& args]
  (let [input (slurp input-file)]
    (println (str "part1: " (part1 input)))
  )
)
