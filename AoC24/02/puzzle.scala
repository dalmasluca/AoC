package day
import scala.io.Source

class Day2(path: String) {
  val lines = Source
    .fromFile(path)
    .getLines()
    .toList
    .map(_.split(" ").map(_.toInt))

  def part1: Int = {
    lines.map { line =>
      val isValid = line
        .sliding(2)
        .forall { a =>
          val diff = a(0) - a(1)
          diff >= 1 && diff <= 3
        }
      if (isValid) 1 else 0
    }.sum +
      lines.map { line =>
        val isValid = line
          .sliding(2)
          .forall { a =>
            val diff = a(1) - a(0)
            diff >= 1 && diff <= 3
          }
        if (isValid) 1 else 0
      }.sum
  }
  def part2: Int = {
    def countSafeReportsWithDampener(lines: Seq[Seq[Int]]): Int = {
      def isSafe(report: Seq[Int]): Boolean = {
        val diffs = report.sliding(2).map { case Seq(a, b) => a - b }.toSeq
        diffs.forall(diff => diff >= -3 && diff <= -1) || diffs.forall(diff =>
          diff >= 1 && diff <= 3
        )
      }

      lines.count { report =>
        isSafe(report) ||
        report.indices.exists { i =>
          isSafe(report.patch(i, Nil, 1))
        }
      }
    }
    countSafeReportsWithDampener(lines.map(_.toSeq).toSeq)
  }
}
