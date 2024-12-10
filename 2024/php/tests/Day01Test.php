<?php

declare(strict_types=1);

use Adventofcode24\Solver;
use PHPUnit\Framework\TestCase;

final class Day01Test extends TestCase
{
    private Solver $solver;

    public function setUp(): void
    {
        parent::setUp();

        $input = file_get_contents(__DIR__ . '/../data/day01.txt');
        $this->solver = new \Adventofcode24\Day01($input);
    }

    public function testPart1(): void
    {
        $this->assertSame('12345', $this->solver->solvePart1());
    }

    public function testPart2(): void
    {
        $this->assertSame('12345', $this->solver->solvePart2());
    }
}
