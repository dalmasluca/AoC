import scala.io.Source

class Day8(path: String) {
  val lines = Source
    .fromFile(path)
    .getLines()
    .toList

  val letter = raw"[a-zA-Z\d]".r

  val antennas = {
    (0 until lines.size).map { line =>
      (0 until lines(line).size).map { pos =>
        val char = lines(line)(pos)
        char.toString match {
          case letter() => (char, (line, pos))
          case _        =>
        }
      }.toList
    }.toList
  }.zipWithIndex
    .flatMap { case (row, rowIndex) =>
      row.zipWithIndex.collect {
        case ((symbol: Any, (x: Int, y: Int)), colIndex) =>
          (symbol, (x, y))
      }
    }
    .groupBy(_._1)
    .map { case (symbol, values) =>
      symbol -> values.map(_._2)
    }

  def part1: Int = {
    val antennaPositions = antennas.values.flatten.toSet
    antennas
      .map { case (symbol, positions) =>
        positions
          .combinations(2)
          .flatMap { case List((x1, y1), (x2, y2)) =>
            val diff = (x2 - x1, y2 - y1)
            List((x1 - diff._1, y1 - diff._2), (x2 + diff._1, y2 + diff._2))
              .filter { case (x: Int, y: Int) =>
                x >= 0 && x < lines(0).size && y >= 0 && y < lines.size
              }
          }
      }
      .toList
      .flatten
      .filterNot(
        antennaPositions.contains
      )
      .size
  }
  def part2: Int = {
    0
  }
}
