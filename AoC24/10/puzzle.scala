import scala.io.Source

object Day {

  val directions = List((1, 0), (-1, 0), (0, 1), (0, -1))

  def main(args: Array[String]): Unit = {
    val input1 = Source.fromFile("small-input1.txt").getLines().toList
    println("Test1 result: " + part1(input1))

    val input2 = Source.fromFile("small-input2.txt").getLines().toList
    println("Test2 result: " + part1(input2))

    val input3 = Source.fromFile("small-input3.txt").getLines().toList
    println("Test3 result: " + part1(input3))

    val input4 = Source.fromFile("input.txt").getLines().toList
    println("Result: " + part1(input4))

    val input2_1 = Source.fromFile("small-input1.txt").getLines().toList
    println("Test1 result: " + part2(input2_1))

    val input2_2 = Source.fromFile("small-input2.txt").getLines().toList
    println("Test2 result: " + part2(input2_2))

    val input2_3 = Source.fromFile("small-input3.txt").getLines().toList
    println("Test3 result: " + part2(input2_3))

    val input2_4 = Source.fromFile("input.txt").getLines().toList
    println("Result: " + part2(input2_4))

    // Placeholder for part2 implementation
    // println("Part2: " + part2(input1))
  }

  def part1(input: List[String]): Int = {
    // trova i punti iniziali (trailhead)
    val trailhead = {
      for {
        x <- input.indices
        y <- input(x).indices
        if input(x)(y) == '0'
      } yield (0, x, y)
    }.toList

    val maxx = input.length
    val maxy = input.head.length

    trailhead.map {
      step: (Int, Int, Int) =>
        var nextstep = step :: List[(Int, Int, Int)]()

        var picks = List[(Int, Int, Int)]()
        while (nextstep.nonEmpty) {
          val (height, x, y) = nextstep.head
          nextstep = nextstep.tail
          directions.foreach { case (dx, dy) =>
            val nx = x + dx
            val ny = y + dy

            // verifica i confini e il valore successivo
            if (nx >= 0 && nx < maxx && ny >= 0 && ny < maxy) {
              val nextheight = input(nx)(ny).asDigit
              if (nextheight == height + 1) {
                if (nextheight == 9) {
                  // se trovi un picco, aggiungi ai risultati
                  picks = (9, nx, ny) :: picks
                } else {
                  // continua ad esplorare aggiungendo il punto a trailhead
                  nextstep = nextstep :+ (nextheight, nx, ny)
                }
              }
            }
          }
        }
        picks.distinct.size
    }.sum
  }

  def part2(input: List[String]): Int = {
    val trailhead = {
      for {
        x <- input.indices
        y <- input(x).indices
        if input(x)(y) == '0'
      } yield (0, x, y)
    }.toList

    val maxx = input.length
    val maxy = input.head.length

    trailhead.map {
      step: (Int, Int, Int) =>
        var nextstep = step :: List[(Int, Int, Int)]()

        var picks = List[(Int, Int, Int)]()
        while (nextstep.nonEmpty) {
          val (height, x, y) = nextstep.head
          nextstep = nextstep.tail
          directions.foreach { case (dx, dy) =>
            val nx = x + dx
            val ny = y + dy

            // verifica i confini e il valore successivo
            if (nx >= 0 && nx < maxx && ny >= 0 && ny < maxy) {
              val nextheight = input(nx)(ny).asDigit
              if (nextheight == height + 1) {
                if (nextheight == 9) {
                  // se trovi un picco, aggiungi ai risultati
                  picks = (9, nx, ny) :: picks
                } else {
                  // continua ad esplorare aggiungendo il punto a trailhead
                  nextstep = nextstep :+ (nextheight, nx, ny)
                }
              }
            }
          }
        }
        picks.size
    }.sum
  }
}
