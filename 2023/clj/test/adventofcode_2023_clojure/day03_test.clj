(ns adventofcode-2023-clojure.day03-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day03 :refer :all]))

(def input-data  (slurp "resources/day03.txt"))
(def test-data  (slurp "resources/day03_test.txt"))

(deftest part1-test
  (are [expected input] (= expected (part1 input))
                        4361 test-data
                        535235 input-data)
)

;; (deftest part2-test
;;   (are [expected input] (= expected (part2 input))
;;                         2286 test-data
;;                         71274 input-data)
;; )