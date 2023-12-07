(ns adventofcode-2023-clojure.day07
  (:use clojure.java.io)
  (:require [clojure.string :as str])
  (:require [clojure.math :as math]))

(def input-file "resources/day07.txt")
(defn text-to-lines [text] (str/split text #"\n"))

(def hand-types {
  :high-card 0
  :one-pair 1
  :two-pairs 2
  :three-of-a-kind 3
  :full-house 4
  :four-of-a-kind 5
  :five-of-a-kind 6
})

(defn compare-cards [cards1 cards2 m]
  (let [pairs (map vector cards1 cards2)
        first-different (first (filter #(not (= (first %) (second %))) pairs))]
        (< (m (first first-different)) (m (second first-different))))
)

(defn compare-hands [hand1 hand2 m]
  (let [type1 (:hand-type hand1)
        type2 (:hand-type hand2)
        cards1 (:cards hand1)
        cards2 (:cards hand2)]
    (if (= type1 type2)
      (compare-cards cards1 cards2 m)
      (< type1 type2)
    )
  )
)

(defn identify-hand-type-without-joker [cards]
  (def cnt (->> cards
    (frequencies)
    (vals)
    (sort >)))

  (cond
    (= (first cnt) 5) (hand-types :five-of-a-kind)
    (= (first cnt) 4) (hand-types :four-of-a-kind)
    (= (first cnt) 3) (if (= (second cnt) 2) (hand-types :full-house) (hand-types :three-of-a-kind))
    (= (first cnt) 2) (if (= (second cnt) 2) (hand-types :two-pairs) (hand-types :one-pair))
    :else (hand-types :high-card)
  )
)

(defn extract-hand [line identify-hand-type]
  (let [el (str/split line #" ")
        cards (str/split (first el) #"")
        bid (Integer/parseInt (second el))]
        {:cards cards, :bid bid, :hand-type (identify-hand-type cards) })
  )

(defn part1 [input]
  (def m {
    "2" 2
    "3" 3
    "4" 4
    "5" 5
    "6" 6
    "7" 7
    "8" 8
    "9" 9
    "T" 10
    "J" 11
    "Q" 12
    "K" 13
    "A" 14
  })
  (->> input
    (text-to-lines)
    (map #(extract-hand % identify-hand-type-without-joker))
    (sort #(compare-hands %1 %2 m))
    (map-indexed vector)
    (map #(* (inc (first %)) (:bid (second %))))
    (reduce +)
  )
)

(defn identify-hand-type-with-joker [cards]
  (def jcnt (count (filter #(= % "J") cards)))
  (def cnt (->> (filter #(not (= % "J")) cards)
    (frequencies)
    (vals)
    (sort >)
  ))
  (def c (if (= 0 (count cnt)) '(5) (cons (+ jcnt (first cnt)) (rest cnt))))
  (cond
    (= (first c) 5) (hand-types :five-of-a-kind)
    (= (first c) 4) (hand-types :four-of-a-kind)
    (= (first c) 3) (if (= (second c) 2) (hand-types :full-house) (hand-types :three-of-a-kind))
    (= (first c) 2) (if (= (second c) 2) (hand-types :two-pairs) (hand-types :one-pair))
    :else (hand-types :high-card)
  )
)

(defn part2 [input]
  (def m {
    "J" 1
    "2" 2
    "3" 3
    "4" 4
    "5" 5
    "6" 6
    "7" 7
    "8" 8
    "9" 9
    "T" 10
    "Q" 12
    "K" 13
    "A" 14
  })
  (->> input
    (text-to-lines)
    (map #(extract-hand % identify-hand-type-with-joker))
    (sort #(compare-hands %1 %2 m))
    (map-indexed vector)
    (map #(* (inc (first %)) (:bid (second %))))
    (reduce +)
  )
)

(defn -main [& args]
  (let [input (slurp input-file)]
    (println (str "part1: " (part1 input)))
    (println (str "part2: " (part2 input)))
))
