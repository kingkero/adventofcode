<?php

declare(strict_types=1);

use Adventofcode24\Solver;
use PHPUnit\Framework\TestCase;

final class Day01Test extends TestCase
{
    private Solver $solver;
    private string $input;

    public function setUp(): void
    {
        parent::setUp();

        $this->solver = new \Adventofcode24\Day01();
        $this->input = file_get_contents(__DIR__ . '/../data/day01.txt');
    }

    public function testPart1(): void
    {
        $this->assertSame('12345', $this->solver->solvePart1($this->input));
    }

    public function testPart2(): void
    {
        $this->assertSame('12345', $this->solver->solvePart2($this->input));
    }
}
