import scala.util.matching.Regex
import scala.io.Source

class Day4(path: String) {
  val lines = Source
    .fromFile(path)
    .getLines()
    .toList

  val xmas: Regex = "XMAS".r

  private def shiftLeft(input: List[String]): List[String] = {
    input.zipWithIndex.map { case (row, index) =>
      val shift = index % row.length
      row.drop(shift) ++ ".".repeat(shift)
    }
  }
  private def shiftRight(input: List[String]): List[String] = {
    input.zipWithIndex.map { case (row, index) =>
      val shift = index % row.length
      ".".repeat(shift) ++ row.dropRight(shift)
    }
  }
  def part1: Int = {
    lines.map { line =>
      xmas.findAllIn(line).toList.size +
        xmas.findAllIn(line.reverse).toList.size
    }.sum +
      lines.transpose
        .map(_.mkString)
        .map { line =>
          xmas.findAllIn(line).toList.size +
            xmas.findAllIn(line.reverse).toList.size
        }
        .sum +
      List
        .range(0, lines.size - "xmas".size + 1)
        .map { i =>
          shiftRight(lines.slice(i, i + "xmas".size)).transpose
            .map(_.mkString)
            .map { line =>
              xmas.findAllIn(line).toList.size +
                xmas.findAllIn(line.reverse).toList.size
            }
            .sum
        }
        .sum +
      List
        .range(0, lines.size - "xmas".size + 1)
        .map { i =>
          shiftLeft(lines.slice(i, i + "xmas".size)).transpose
            .map(_.mkString)
            .map { line =>
              xmas.findAllIn(line).toList.size +
                xmas.findAllIn(line.reverse).toList.size
            }
            .sum
        }
        .sum
  }
  def part2: Int = {
    List
      .range(0, lines.size - "mas".size + 1)
      .map { i =>
        val slicedLine = lines.slice(i, i + "mas".size)
      List
        .range(0, lines(i).size - "mas".size + 1)
        .map { j =>
          val square = slicedLine.map(line => line.slice(j, j + "mas".size))
          if (
            shiftLeft(square).transpose
              .map(_.mkString)
              .map { l =>
                "MAS".r.findAllIn(l).toList.size +
              "SAM".r.findAllIn(l).toList.size
              }
              .sum +
              shiftRight(square).transpose
                .map(_.mkString)
                .map { l =>
                  "MAS".r.findAllIn(l).toList.size +
                    "SAM".r.findAllIn(l).toList.size
                }
                .sum
              == 2
          ) 1
          else 0
        }
        .sum
      }
      .sum
  }
}
