(ns adventofcode-2023-clojure.day02-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day02 :refer :all]))

(def test-data (str "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n"
                    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n"
                    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n"
                    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n"
                    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))
(def input-data  (slurp "resources/day02.txt"))

(deftest part1-test
  (are [expected input] (= expected (part1 input))
                        8 test-data
                        2149 input-data)
)

;; (deftest part2-test
;;   (are [expected input] (= expected (part2 input))
;;                         2286 test-data
;;                         71274 input-data)
;; )