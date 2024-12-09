import scala.io.Source
import scala.languageFeature.postfixOps

class Day6(path: String) {
  val lines = Source
    .fromFile(path)
    .getLines()
    .toList
    .map(_.split(""))

  val directions = Array[(Int, Int)]((-1, 0), (0, 1), (1, 0), (0, -1))

  val start = for {
    row <- lines.indices
    col <- lines(row).indices
    if lines(row)(col) == "^"
  } yield (row, col)

  def step(
      positions: (Int, Int),
      dir: Int,
      Acc: List[(Int, Int)]
  ): List[(Int, Int)] = {
    val rows = lines.size - 1
    val cols = lines(0).size - 1

    positions match {
      case (_, y) if y == 0 || y >= rows => positions :: Acc
      case (x, _) if x == 0 || x >= cols => positions :: Acc
      case _ => {
        val nextdir = {
          val nextRows = positions._1 + directions(dir)._1
          val nextCols = positions._2 + directions(dir)._2
          if (lines(nextRows)(nextCols) == "#") (dir + 1) % directions.size
          else dir
        }
        val next =
          (
            positions._1 + directions(nextdir)._1,
            positions._2 + directions(nextdir)._2
          )
        step(next, nextdir, positions :: Acc)
      }
    }
  }

  def part1: Int = {
    val initialPos = start.head
    step(initialPos, 0, List[(Int, Int)]()).distinct.size
  }
  def part2: Int = {
    0
  }
}
