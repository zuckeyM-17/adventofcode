(ns adventofcode-2023-clojure.day01
  (:use clojure.java.io)
  (:require [clojure.string :as str]))

(def input-file "resources/day01.txt")

(defn word-to-number [w]
  (case w "one" 1 "two" 2 "three" 3 "four" 4 "five" 5 "six" 6 "seven" 7 "eight" 8 "nine" 9 (Integer/parseInt w)))

(defn find-first-last-digits [s]
  (let [digits (re-seq #"\d" s)]
    (map #(Integer/parseInt %) [(first digits) (last digits)])))

(defn find-first-last-numbers [s]
  (let [regex (re-pattern "one|two|three|four|five|six|seven|eight|nine|\\d")
        matches (re-seq regex s)]
    (map word-to-number [(first matches) (last matches)])))

(defn part1 [input]
  (def digits (let [lines (str/split input #"\n")]
    (map find-first-last-digits lines)))
  (reduce + (map #(+ (* (first %) 10) (second %)) digits))
)

(defn part2 [input]
  (def digits (let [lines (str/split input #"\n")]
    (map find-first-last-numbers lines)))
  (reduce + (map #(+ (* (first %) 10) (second %)) digits))
)

(defn -main [& args]
  (with-open [fin(reader input-file)]
    (let [input (slurp fin)]
      (println (str "part1: " (part1 input)))
      (println (str "part2: " (part2 input)))
    )
  )
)
