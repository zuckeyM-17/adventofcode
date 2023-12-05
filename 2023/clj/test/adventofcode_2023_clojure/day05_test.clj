(ns adventofcode-2023-clojure.day05-test
  (:require [clojure.test :refer :all]
            [adventofcode-2023-clojure.day05 :refer :all]))

(def input-data  (slurp "resources/day05.txt"))
(def test-data  (slurp "resources/day05_test.txt"))

(deftest part1-test
  (are [expected input] (= expected (part1 input))
                        35 test-data
                        389056265 input-data)
)

;; (deftest part2-test
;;   (are [expected input] (= expected (part2 input))
;;                         2286 test-data
;;                         71274 input-data)
;; )