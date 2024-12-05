package day
import scala.io.Source

class Day(path: String) {

  val (left, right) = Source
    .fromFile(path)
    .getLines()
    .toList
    .map(_.split("\\s+").map(_.toInt))
    .foldLeft((List.empty[Int], List.empty[Int])) {
      case ((l, r), Array(a, b)) => (l :+ a, r :+ b)
    }
  def part1: Int = {
    left.sorted.zip(right.sorted).map { case (a, b) => Math.abs(a - b) }.sum
  }
  def part2: Int = {
    val occ =
      right.groupBy(identity).view.mapValues(_.size).filterKeys(left.contains)
    left.foldLeft(0) { (acc, el) =>
      acc + (el * occ.getOrElse(el, 0))
    }
  }
}
