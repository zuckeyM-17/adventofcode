(ns adventofcode-2023-clojure.day04-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day04 :refer :all]))

(def input-data  (slurp "resources/day04.txt"))
(def test-data  (slurp "resources/day04_test.txt"))

(deftest part1-test
  (are [expected input] (= expected (part1 input))
                        13 test-data
                        23750 input-data)
)

(deftest part2-test
  (are [expected input] (= expected (part2 input))
                        30 test-data
                        13261850 input-data)
)