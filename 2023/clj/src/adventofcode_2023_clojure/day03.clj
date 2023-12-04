(ns adventofcode-2023-clojure.day03
  (:use clojure.java.io)
  (:require [clojure.string :as str]))

(def input-file "resources/day03.txt")
(defn text-to-lines [text] (str/split text #"\n"))
(defn get-char [lines i j] (get (get lines i) j))

(defn is-symbol [c] (contains? (set "*#+-$@%&=/") c))
(defn is-ast [c] (= c "*"))

(defn contains-true? [coll]
  (boolean (some identity coll))
)

(defn check-surroundings? [lines i j check-fn]
  (let [surroundings '([-1 -1] [-1 0] [-1 1] [0 -1] [0 1] [1 -1] [1 0] [1 1])]
    (contains-true? (for [s surroundings]
      (let [x (+ i (first s)) y (+ j (second s))]
        (if (and (>= x 0) (>= y 0) (< x (count lines)) (< y (count (get lines x))))
          (check-fn (get-char lines x y)) false
        ))))))

(defn check-digits-surroundings? [lines i index check-fn]
  (contains-true? (for [j (range (first index) (second index))]
    (check-surroundings? lines i j check-fn)
  )))

(defn find-digit-indexes [line]
  (let [matcher (re-matcher (re-pattern "\\d+") line)]
    (loop [indexes []]
      (if (.find matcher)
        (recur (conj indexes [(.start matcher) (.end matcher)]))
        indexes
      )
    )
  ))

(defn get-digit [line index]
  (Integer/parseInt (subs line (first index) (second index))))

(defn part1 [input]
  (def digits (let [lines (text-to-lines input)]
    (flatten (for [line (map-indexed vector lines)]
      (map
        #(get-digit (second line) %)
        (filter
        #(check-digits-surroundings? lines (first line) % is-symbol)
        (find-digit-indexes (second line))))
    ))
  ))
  (reduce + digits)
)

(defn part2 [input]
  (let [lines (text-to-lines input)]
    ()
  )
)

(defn -main [& args]
  (let [input (slurp input-file)]
    (println (str "part1: " (part1 input)))
  )
)
